// Copyright (c) The OGX Contributors.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ogxclient

import (
	"github.com/ogx-ai/ogx-client-go/internal/apijson"
	"github.com/ogx-ai/ogx-client-go/packages/respjson"
)

// Health status information for the service.
type HealthInfo struct {
	// The health status of the service
	//
	// Any of "OK", "Error", "Not Implemented".
	Status HealthInfoStatus `json:"status" api:"required"`
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

// Response containing a list of all available providers.
type ListProvidersResponse struct {
	// List of provider information objects
	Data []ProviderInfo `json:"data" api:"required"`
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

// Response containing a list of all available API routes.
type ListRoutesResponse struct {
	// List of available API routes
	Data []RouteInfo `json:"data" api:"required"`
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

// Information about a registered provider including its configuration and health
// status.
type ProviderInfo struct {
	// The API name this provider implements
	API string `json:"api" api:"required"`
	// Configuration parameters for the provider
	Config map[string]any `json:"config" api:"required"`
	// Current health status of the provider
	Health map[string]any `json:"health" api:"required"`
	// Unique identifier for the provider
	ProviderID string `json:"provider_id" api:"required"`
	// The type of provider implementation
	ProviderType string `json:"provider_type" api:"required"`
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
	Method string `json:"method" api:"required"`
	// List of provider types implementing this route
	ProviderTypes []string `json:"provider_types" api:"required"`
	// The API route path
	Route string `json:"route" api:"required"`
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
	Version string `json:"version" api:"required"`
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
