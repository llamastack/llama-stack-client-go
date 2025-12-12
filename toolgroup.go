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
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ToolgroupService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolgroupService] method instead.
type ToolgroupService struct {
	Options []option.RequestOption
}

// NewToolgroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolgroupService(opts ...option.RequestOption) (r ToolgroupService) {
	r = ToolgroupService{}
	r.Options = opts
	return
}

// List tool groups with optional provider.
//
// Deprecated: deprecated
func (r *ToolgroupService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ToolGroup, err error) {
	var env ListToolGroupsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/toolgroups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a tool group by its ID.
//
// Deprecated: deprecated
func (r *ToolgroupService) Get(ctx context.Context, toolgroupID string, opts ...option.RequestOption) (res *ToolGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolgroupID == "" {
		err = errors.New("missing required toolgroup_id parameter")
		return
	}
	path := fmt.Sprintf("v1/toolgroups/%s", toolgroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Register a tool group.
//
// Deprecated: deprecated
func (r *ToolgroupService) Register(ctx context.Context, body ToolgroupRegisterParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/toolgroups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Unregister a tool group.
//
// Deprecated: deprecated
func (r *ToolgroupService) Unregister(ctx context.Context, toolgroupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if toolgroupID == "" {
		err = errors.New("missing required toolgroup_id parameter")
		return
	}
	path := fmt.Sprintf("v1/toolgroups/%s", toolgroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response containing a list of tool groups.
type ListToolGroupsResponse struct {
	Data []ToolGroup `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListToolGroupsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListToolGroupsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A group of related tools managed together.
type ToolGroup struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string         `json:"provider_id,required"`
	Args       map[string]any `json:"args,nullable"`
	// A URL reference to external content.
	McpEndpoint ToolGroupMcpEndpoint `json:"mcp_endpoint,nullable"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Any of "tool_group".
	Type ToolGroupType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Args               respjson.Field
		McpEndpoint        respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolGroup) RawJSON() string { return r.JSON.raw }
func (r *ToolGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
type ToolGroupMcpEndpoint struct {
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolGroupMcpEndpoint) RawJSON() string { return r.JSON.raw }
func (r *ToolGroupMcpEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolGroupType string

const (
	ToolGroupTypeToolGroup ToolGroupType = "tool_group"
)

type ToolgroupRegisterParams struct {
	ProviderID  string         `json:"provider_id,required"`
	ToolgroupID string         `json:"toolgroup_id,required"`
	Args        map[string]any `json:"args,omitzero"`
	// A URL reference to external content.
	McpEndpoint ToolgroupRegisterParamsMcpEndpoint `json:"mcp_endpoint,omitzero"`
	paramObj
}

func (r ToolgroupRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolgroupRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolgroupRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type ToolgroupRegisterParamsMcpEndpoint struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r ToolgroupRegisterParamsMcpEndpoint) MarshalJSON() (data []byte, err error) {
	type shadow ToolgroupRegisterParamsMcpEndpoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolgroupRegisterParamsMcpEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
