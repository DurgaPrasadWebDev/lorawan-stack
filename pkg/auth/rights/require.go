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

package rights

import (
	"context"

	errors "go.thethings.network/lorawan-stack/pkg/errorsv3"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
)

var errMissingApplicationRights = errors.DefinePermissionDenied(
	"missing_application_rights",
	"missing rights for application `{uid}`",
)

var errMissingGatewayRights = errors.DefinePermissionDenied(
	"missing_gateway_rights",
	"missing rights for gateway `{uid}`",
)

var errMissingOrganizationRights = errors.DefinePermissionDenied(
	"missing_organization_rights",
	"missing rights for organization `{uid}`",
)

// RequireApplication checks that context contains the required rights for the
// given Application ID.
func RequireApplication(ctx context.Context, appID ttnpb.ApplicationIdentifiers, required ...ttnpb.Right) error {
	appUID := unique.ID(ctx, appID)
	rights, ok := FromContext(ctx)
	if !ok {
		fetcher, ok := fetcherFromContext(ctx)
		if !ok {
			panic(errNoFetcher)
		}
		appRights, err := fetcher.ApplicationRights(ctx, appID)
		switch {
		case err == nil, errors.IsPermissionDenied(err):
			break
		default:
			return err
		}
		rights.ApplicationRights = map[string]*ttnpb.Rights{appUID: appRights}
	}
	if !rights.IncludesApplicationRights(appUID, required...) {
		return errMissingApplicationRights.WithAttributes("uid", appUID)
	}
	return nil
}

// RequireGateway checks that context contains the required rights for the
// given Gateway ID.
func RequireGateway(ctx context.Context, gtwID ttnpb.GatewayIdentifiers, required ...ttnpb.Right) error {
	gtwUID := unique.ID(ctx, gtwID)
	rights, ok := FromContext(ctx)
	if !ok {
		fetcher, ok := fetcherFromContext(ctx)
		if !ok {
			panic(errNoFetcher)
		}
		gtwRights, err := fetcher.GatewayRights(ctx, gtwID)
		switch {
		case err == nil, errors.IsPermissionDenied(err):
			break
		default:
			return err
		}
		rights.GatewayRights = map[string]*ttnpb.Rights{gtwUID: gtwRights}
	}
	if !rights.IncludesGatewayRights(gtwUID, required...) {
		return errMissingGatewayRights.WithAttributes("uid", gtwUID)
	}
	return nil
}

// RequireOrganization checks that context contains the required rights for the
// given organization ID.
func RequireOrganization(ctx context.Context, orgID ttnpb.OrganizationIdentifiers, required ...ttnpb.Right) error {
	orgUID := unique.ID(ctx, orgID)
	rights, ok := FromContext(ctx)
	if !ok {
		fetcher, ok := fetcherFromContext(ctx)
		if !ok {
			panic(errNoFetcher)
		}
		orgRights, err := fetcher.OrganizationRights(ctx, orgID)
		switch {
		case err == nil, errors.IsPermissionDenied(err):
			break
		default:
			return err
		}
		rights.OrganizationRights = map[string]*ttnpb.Rights{orgUID: orgRights}
	}
	if !rights.IncludesOrganizationRights(orgUID, required...) {
		return errMissingOrganizationRights.WithAttributes("uid", orgUID)
	}
	return nil
}
