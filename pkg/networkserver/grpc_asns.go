// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package networkserver provides a LoRaWAN 1.1-compliant Network Server implementation.
package networkserver

import (
	"context"

	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/pkg/events"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
)

type applicationUpStream struct {
	ttnpb.AsNs_LinkApplicationServer
	closeCh chan struct{}
}

func (s applicationUpStream) Close() error {
	close(s.closeCh)
	return nil
}

// LinkApplication is called by the Application Server to subscribe to application events.
func (ns *NetworkServer) LinkApplication(id *ttnpb.ApplicationIdentifiers, stream ttnpb.AsNs_LinkApplicationServer) (err error) {
	ws := &applicationUpStream{
		AsNs_LinkApplicationServer: stream,
		closeCh:                    make(chan struct{}),
	}

	ctx := stream.Context()
	uid := unique.ID(ctx, id)

	events.Publish(evtBeginApplicationLink(ctx, id, nil))
	defer events.Publish(evtEndApplicationLink(ctx, id, err))

	ns.applicationServersMu.Lock()
	cl, ok := ns.applicationServers[uid]
	ns.applicationServers[uid] = ws
	if ok {
		if err := cl.Close(); err != nil {
			ns.applicationServersMu.Unlock()
			return err
		}
	}
	ns.applicationServersMu.Unlock()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		ns.applicationServersMu.Lock()
		cl, ok := ns.applicationServers[uid]
		if ok && cl == ws {
			delete(ns.applicationServers, uid)
		}
		ns.applicationServersMu.Unlock()
		return err
	case <-ws.closeCh:
		return errDuplicateSubscription
	}
}

// DownlinkQueueReplace is called by the Application Server to completely replace the downlink queue for a device.
func (ns *NetworkServer) DownlinkQueueReplace(ctx context.Context, req *ttnpb.DownlinkQueueRequest) (*pbtypes.Empty, error) {
	dev, err := ns.devices.SetByID(ctx, req.EndDeviceIdentifiers.ApplicationIdentifiers, req.EndDeviceIdentifiers.DeviceID, func(dev *ttnpb.EndDevice) (*ttnpb.EndDevice, error) {
		if dev == nil {
			return nil, errDeviceNotFound
		}
		dev.QueuedApplicationDownlinks = req.Downlinks
		return dev, nil
	})
	if err != nil {
		return nil, err
	}

	if dev.MACState != nil && dev.MACState.DeviceClass == ttnpb.CLASS_C {
		// TODO: Schedule the next downlink (https://github.com/TheThingsIndustries/ttn/issues/728)
	}
	return ttnpb.Empty, nil
}

// DownlinkQueuePush is called by the Application Server to push a downlink to queue for a device.
func (ns *NetworkServer) DownlinkQueuePush(ctx context.Context, req *ttnpb.DownlinkQueueRequest) (*pbtypes.Empty, error) {
	dev, err := ns.devices.SetByID(ctx, req.EndDeviceIdentifiers.ApplicationIdentifiers, req.EndDeviceIdentifiers.DeviceID, func(dev *ttnpb.EndDevice) (*ttnpb.EndDevice, error) {
		if dev == nil {
			return nil, errDeviceNotFound
		}
		dev.QueuedApplicationDownlinks = append(dev.QueuedApplicationDownlinks, req.Downlinks...)
		return dev, nil
	})
	if err != nil {
		return nil, err
	}

	if dev.MACState != nil && dev.MACState.DeviceClass == ttnpb.CLASS_C {
		// TODO: Schedule the next downlink (https://github.com/TheThingsIndustries/ttn/issues/728)
	}
	return ttnpb.Empty, nil
}

// DownlinkQueueList is called by the Application Server to get the current state of the downlink queue for a device.
func (ns *NetworkServer) DownlinkQueueList(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) (*ttnpb.ApplicationDownlinks, error) {
	dev, err := ns.devices.GetByID(ctx, ids.ApplicationIdentifiers, ids.DeviceID)
	if err != nil {
		return nil, err
	}
	return &ttnpb.ApplicationDownlinks{Downlinks: dev.QueuedApplicationDownlinks}, nil
}