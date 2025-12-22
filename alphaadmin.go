// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// AlphaAdminService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaAdminService] method instead.
type AlphaAdminService struct {
	Options []option.RequestOption
}

// NewAlphaAdminService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaAdminService(opts ...option.RequestOption) (r AlphaAdminService) {
	r = AlphaAdminService{}
	r.Options = opts
	return
}

// Get the current health status of the service.
func (r *AlphaAdminService) Health(ctx context.Context, opts ...option.RequestOption) (res *HealthInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/admin/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get detailed information about a specific provider.
func (r *AlphaAdminService) InspectProvider(ctx context.Context, providerID string, opts ...option.RequestOption) (res *ProviderInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/admin/providers/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available providers with their configuration and health status.
func (r *AlphaAdminService) ListProviders(ctx context.Context, opts ...option.RequestOption) (res *[]ProviderInfo, err error) {
	var env ListProvidersResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/admin/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all available API routes with their methods and implementing providers.
func (r *AlphaAdminService) ListRoutes(ctx context.Context, query AlphaAdminListRoutesParams, opts ...option.RequestOption) (res *[]RouteInfo, err error) {
	var env ListRoutesResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/admin/inspect/routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get the version of the service.
func (r *AlphaAdminService) Version(ctx context.Context, opts ...option.RequestOption) (res *VersionInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/admin/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing a list of all available providers.
type ListProvidersResponse struct {
	// List of provider information objects
	Data []ProviderInfo `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListProvidersResponse) RawJSON() string { return r.JSON.raw }
func (r *ListProvidersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaAdminListRoutesParams struct {
	// Filter to control which routes are returned. Can be an API level ('v1',
	// 'v1alpha', 'v1beta') to show non-deprecated routes at that level, or
	// 'deprecated' to show deprecated routes across all levels. If not specified,
	// returns all non-deprecated routes.
	//
	// Any of "v1", "v1alpha", "v1beta", "deprecated".
	APIFilter AlphaAdminListRoutesParamsAPIFilter `query:"api_filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaAdminListRoutesParams]'s query parameters as
// `url.Values`.
func (r AlphaAdminListRoutesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to control which routes are returned. Can be an API level ('v1',
// 'v1alpha', 'v1beta') to show non-deprecated routes at that level, or
// 'deprecated' to show deprecated routes across all levels. If not specified,
// returns all non-deprecated routes.
type AlphaAdminListRoutesParamsAPIFilter string

const (
	AlphaAdminListRoutesParamsAPIFilterV1         AlphaAdminListRoutesParamsAPIFilter = "v1"
	AlphaAdminListRoutesParamsAPIFilterV1alpha    AlphaAdminListRoutesParamsAPIFilter = "v1alpha"
	AlphaAdminListRoutesParamsAPIFilterV1beta     AlphaAdminListRoutesParamsAPIFilter = "v1beta"
	AlphaAdminListRoutesParamsAPIFilterDeprecated AlphaAdminListRoutesParamsAPIFilter = "deprecated"
)

// Response containing a list of all available API routes.
type ListRoutesResponse struct {
	// List of available API routes
	Data []RouteInfo `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListRoutesResponse) RawJSON() string { return r.JSON.raw }
func (r *ListRoutesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
