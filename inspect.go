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
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// InspectService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInspectService] method instead.
type InspectService struct {
	Options []option.RequestOption
}

// NewInspectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInspectService(opts ...option.RequestOption) (r InspectService) {
	r = InspectService{}
	r.Options = opts
	return
}

// Get the current health status of the service.
func (r *InspectService) Health(ctx context.Context, opts ...option.RequestOption) (res *HealthInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the version of the service.
func (r *InspectService) Version(ctx context.Context, opts ...option.RequestOption) (res *VersionInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Health status information for the service.
type HealthInfo struct {
	// The health status of the service
	//
	// Any of "OK", "Error", "Not Implemented".
	Status HealthInfoStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthInfo) RawJSON() string { return r.JSON.raw }
func (r *HealthInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The health status of the service
type HealthInfoStatus string

const (
	HealthInfoStatusOk             HealthInfoStatus = "OK"
	HealthInfoStatusError          HealthInfoStatus = "Error"
	HealthInfoStatusNotImplemented HealthInfoStatus = "Not Implemented"
)

// Information about a registered provider including its configuration and health
// status.
type ProviderInfo struct {
	// The API name this provider implements
	API string `json:"api,required"`
	// Configuration parameters for the provider
	Config map[string]any `json:"config,required"`
	// Current health status of the provider
	Health map[string]any `json:"health,required"`
	// Unique identifier for the provider
	ProviderID string `json:"provider_id,required"`
	// The type of provider implementation
	ProviderType string `json:"provider_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		API          respjson.Field
		Config       respjson.Field
		Health       respjson.Field
		ProviderID   respjson.Field
		ProviderType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderInfo) RawJSON() string { return r.JSON.raw }
func (r *ProviderInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about an API route including its path, method, and implementing
// providers.
type RouteInfo struct {
	// The HTTP method for the route
	Method string `json:"method,required"`
	// List of provider types implementing this route
	ProviderTypes []string `json:"provider_types,required"`
	// The API route path
	Route string `json:"route,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		ProviderTypes respjson.Field
		Route         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteInfo) RawJSON() string { return r.JSON.raw }
func (r *RouteInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version information for the service.
type VersionInfo struct {
	// The version string of the service
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionInfo) RawJSON() string { return r.JSON.raw }
func (r *VersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
