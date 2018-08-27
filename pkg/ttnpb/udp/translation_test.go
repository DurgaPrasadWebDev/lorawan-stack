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

package udp_test

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"time"

	"github.com/kr/pretty"
	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/ttnpb/udp"
	"go.thethings.network/lorawan-stack/pkg/types"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
	"go.thethings.network/lorawan-stack/pkg/version"
)

var ids = ttnpb.GatewayIdentifiers{GatewayID: "test-gateway"}

func TestStatusRaw(t *testing.T) {
	a := assertions.New(t)

	raw := []byte(`{"stat":{"rxfw":0,"hal":"5.1.0","fpga":2,"dsp":31,"lpps":2,"lmnw":3,"lmst":1,"lmok":3,"temp":30,"lati":52.34223,"long":5.29685,"txnb":0,"dwnb":0,"alti":66,"rxok":0,"boot":"2017-06-07 09:40:42 GMT","time":"2017-06-08 09:40:42 GMT","rxnb":0,"ackr":0.0}}`)
	var statusData udp.Data
	err := json.Unmarshal(raw, &statusData)
	a.So(err, should.BeNil)

	upstream, err := udp.ToGatewayUp(statusData, udp.UpstreamMetadata{
		IP: "127.0.0.1",
		ID: ids,
	})
	a.So(err, should.BeNil)

	status := upstream.GatewayStatus
	a.So(status, should.NotBeNil)

	a.So(status.AntennaLocations, should.NotBeNil)
	a.So(len(status.AntennaLocations), should.Equal, 1)
	a.So(status.AntennaLocations[0].Longitude, should.AlmostEqual, 5.29685, 0.0001)
	a.So(status.AntennaLocations[0].Latitude, should.AlmostEqual, 52.34223, 0.0001)
	a.So(status.AntennaLocations[0].Altitude, should.AlmostEqual, 66)

	a.So(status.Versions, should.NotBeNil)
	a.So(status.Metrics, should.NotBeNil)

	a.So(status.Versions["ttn-lw-gateway-server"], should.Equal, version.TTN)
	a.So(status.Versions["hal"], should.Equal, "5.1.0")
	a.So(status.Versions["fpga"], should.Equal, "2")
	a.So(status.Versions["dsp"], should.Equal, "31")

	a.So(status.Metrics["rxfw"], should.AlmostEqual, 0)
	a.So(status.Metrics["txnb"], should.AlmostEqual, 0)
	a.So(status.Metrics["dwnb"], should.AlmostEqual, 0)
	a.So(status.Metrics["rxok"], should.AlmostEqual, 0)
	a.So(status.Metrics["rxnb"], should.AlmostEqual, 0)
	a.So(status.Metrics["ackr"], should.AlmostEqual, 0)
	a.So(status.Metrics["temp"], should.AlmostEqual, 30)
	a.So(status.Metrics["lpps"], should.AlmostEqual, 2)
	a.So(status.Metrics["lmnw"], should.AlmostEqual, 3)
	a.So(status.Metrics["lmst"], should.AlmostEqual, 1)
	a.So(status.Metrics["lmok"], should.AlmostEqual, 3)

	a.So(status.BootTime, should.NotBeNil)
	a.So(status.Time, should.NotBeNil)
	currentTime := time.Date(2017, 06, 8, 9, 40, 42, 0, time.UTC)
	a.So(status.Time, should.Equal, currentTime)
	bootTime := time.Date(2017, 06, 7, 9, 40, 42, 0, time.UTC)
	a.So(status.BootTime, should.Equal, bootTime)
}

func TestToGatewayUp(t *testing.T) {
	a := assertions.New(t)

	p := udp.Packet{
		GatewayEUI:      &types.EUI64{0xAA, 0xEE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ProtocolVersion: udp.Version1,
		Token:           [2]byte{0x11, 0x00},
		Data: &udp.Data{
			RxPacket: []*udp.RxPacket{
				{
					Freq: 868.0,
					Chan: 2,
					Modu: "LORA",
					DatR: udp.DataRate{DataRate: types.DataRate{LoRa: "SF10BW125"}},
					CodR: "4/7",
					Data: "QCkuASaAAAAByFaF53Iu+vzmwQ==",
					Size: 19,
					Tmst: 1000,
				},
			},
		},
		PacketType: udp.PushData,
	}

	upstream, err := udp.ToGatewayUp(*p.Data, udp.UpstreamMetadata{ID: ids})
	a.So(err, should.BeNil)

	msg := upstream.UplinkMessages[0]
	a.So(msg.Settings.CodingRate, should.Equal, "4/7")
	a.So(msg.Settings.SpreadingFactor, should.Equal, 10)
	a.So(msg.Settings.Bandwidth, should.Equal, 125000)
	a.So(msg.Settings.Frequency, should.Equal, 868000000)
	a.So(msg.Settings.Modulation, should.Equal, ttnpb.Modulation_LORA)
	a.So(msg.RxMetadata[0].Timestamp, should.Equal, 1000000)
	a.So(msg.RawPayload, should.Resemble, []byte{0x40, 0x29, 0x2e, 0x01, 0x26, 0x80, 0x00, 0x00, 0x01, 0xc8, 0x56, 0x85, 0xe7, 0x72, 0x2e, 0xfa, 0xfc, 0xe6, 0xc1})
}

func TestToGatewayUpRoundtrip(t *testing.T) {
	a := assertions.New(t)

	expected := udp.Packet{
		ProtocolVersion: udp.Version1,
		Token:           [2]byte{0x11, 0x00},
		Data: &udp.Data{
			RxPacket: []*udp.RxPacket{
				{
					Freq: 868.0,
					Chan: 2,
					Modu: "LORA",
					DatR: udp.DataRate{DataRate: types.DataRate{LoRa: "SF10BW125"}},
					CodR: "4/7",
					Data: "QCkuASaAAAAByFaF53Iu+vzmwQ==",
					Size: 19,
					Tmst: 1000,
				},
			},
		},
		PacketType: udp.PushData,
	}
	expectedMd := udp.UpstreamMetadata{
		ID: ttnpb.GatewayIdentifiers{
			EUI: &types.EUI64{0xAA, 0xEE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		IP: "1.1.1.1",
	}

	up, err := udp.ToGatewayUp(*expected.Data, expectedMd)
	a.So(err, should.BeNil)

	actual := udp.Packet{
		ProtocolVersion: udp.Version1,
		Token:           [2]byte{0x11, 0x00},
		Data:            &udp.Data{},
	}
	actual.Data.RxPacket, actual.Data.Stat = udp.FromGatewayUp(up)

	a.So(pretty.Diff(actual, expected), should.BeEmpty)
}

func TestToGatewayUpRaw(t *testing.T) {
	a := assertions.New(t)

	raw := []byte(`{"rxpk":[{"tmst":368384825,"chan":0,"rfch":0,"freq":868.100000,"stat":1,"modu":"LORA","datr":"SF7BW125","codr":"4/5","lsnr":-11,"rssi":-107,"size":108,"data":"Wqish6GVYpKy6o9WFHingeTJ1oh+ABc8iALBvwz44yxZP+BKDocaC5VQT5Y6dDdUaBILVjRMz0Ynzow1U/Kkts9AoZh3Ja3DX+DyY27exB+BKpSx2rXJ2vs9svm/EKYIsPF0RG1E+7lBYaD9"}]}`)
	var rxData udp.Data
	err := json.Unmarshal(raw, &rxData)
	a.So(err, should.BeNil)

	upstream, err := udp.ToGatewayUp(rxData, udp.UpstreamMetadata{ID: ids})
	a.So(err, should.BeNil)

	a.So(len(upstream.UplinkMessages), should.Equal, 1)
	msg := upstream.UplinkMessages[0]
	a.So(msg.Settings.CodingRate, should.Equal, "4/5")
	a.So(msg.Settings.SpreadingFactor, should.Equal, 7)
	a.So(msg.Settings.Bandwidth, should.Equal, 125000)
	a.So(msg.Settings.Frequency, should.Equal, uint64(868100000))
	a.So(msg.Settings.Modulation, should.Equal, ttnpb.Modulation_LORA)
	a.So(msg.RxMetadata[0].Timestamp, should.Equal, uint64(368384825000))
	a.So(len(msg.RawPayload), should.Equal, base64.StdEncoding.DecodedLen(len("Wqish6GVYpKy6o9WFHingeTJ1oh+ABc8iALBvwz44yxZP+BKDocaC5VQT5Y6dDdUaBILVjRMz0Ynzow1U/Kkts9AoZh3Ja3DX+DyY27exB+BKpSx2rXJ2vs9svm/EKYIsPF0RG1E+7lBYaD9")))
}

func TestToGatewayUpRawMultiAntenna(t *testing.T) {
	a := assertions.New(t)

	rx := []byte(`{
		"rxpk": [{
			"tmst": 879148780,
			"time": "2017-07-04T13:51:17.997099Z",
			"rfch": 0,
			"freq": 868.500000,
			"stat": 1,
			"modu": "LORA",
			"datr": "SF7BW125",
			"codr": "4/5",
			"size": 24,
			"data": "gM+AMQcAvgQBlohnlJqUGOJKTDuTscQD",
			"rsig": [{
				"ant": 0,
				"chan": 7,
				"rssic": -95,
				"lsnr": 14.0,
				"etime": "42QMzOlYSSPMMeqVPrY0fQ==",
				"rssis": -95,
				"rssisd": 0,
				"ftime": 1255738435,
				"foff": -8898,
				"ft2d": -251,
				"rfbsb": 100,
				"rs2s1": 97
			}, {
				"ant": 1,
				"chan": 23,
				"rssic": -88,
				"lsnr": 14.0,
				"etime": "djGiSzOC+gCT7vRPv7+Asw==",
				"rssis": -88,
				"rssisd": 0,
				"ftime": -1252538435,
				"foff": -8898,
				"ft2d": -187,
				"rfbsb": 100,
				"rs2s1": 104
			}]
		}]
	}`)
	var rxData udp.Data
	err := json.Unmarshal(rx, &rxData)
	a.So(err, should.BeNil)

	_, err = udp.ToGatewayUp(rxData, udp.UpstreamMetadata{ID: ids})
	a.So(err, should.BeNil)
}

func TestFromDownlinkMessage(t *testing.T) {
	a := assertions.New(t)

	msg := &ttnpb.DownlinkMessage{
		TxMetadata: ttnpb.TxMetadata{
			Timestamp: 1886440700000,
		},
		Settings: ttnpb.TxSettings{
			Frequency:             925700000,
			Modulation:            ttnpb.Modulation_LORA,
			TxPower:               20,
			SpreadingFactor:       10,
			Bandwidth:             500000,
			PolarizationInversion: true,
		},
		RawPayload: []byte{0x7d, 0xf3, 0x8e},
	}
	tx, err := udp.FromDownlinkMessage(msg)
	a.So(err, should.BeNil)
	a.So(tx.DatR.LoRa, should.Equal, "SF10BW500")
	a.So(tx.Tmst, should.Equal, 1886440700)
	a.So(tx.NCRC, should.Equal, true)
	a.So(tx.Data, should.Equal, "ffOO")
}

func TestDownlinkRoundtrip(t *testing.T) {
	a := assertions.New(t)
	expected := &ttnpb.DownlinkMessage{
		TxMetadata: ttnpb.TxMetadata{
			Timestamp: 188700000,
		},
		Settings: ttnpb.TxSettings{
			Frequency:             925700000,
			Modulation:            ttnpb.Modulation_LORA,
			TxPower:               20,
			SpreadingFactor:       10,
			Bandwidth:             500000,
			PolarizationInversion: true,
		},
		RawPayload: []byte{0x7d, 0xf3, 0x8e},
	}
	tx, err := udp.FromDownlinkMessage(expected)
	a.So(err, should.BeNil)

	actual, err := udp.ToDownlinkMessage(tx)
	a.So(err, should.BeNil)

	a.So(pretty.Diff(actual, expected), should.BeEmpty)
}

func TestFromDownlinkMessageDummy(t *testing.T) {
	a := assertions.New(t)

	msg := ttnpb.DownlinkMessage{Settings: ttnpb.TxSettings{Modulation: 3939}} // Dummy modulation set
	_, err := udp.FromDownlinkMessage(&msg)
	a.So(err, should.NotBeNil)
}