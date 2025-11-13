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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/toolgroups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Unregister a tool group.
//
// Deprecated: deprecated
func (r *ToolgroupService) Unregister(ctx context.Context, toolgroupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	// List of tool groups
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
	Identifier string `json:"identifier,required"`
	ProviderID string `json:"provider_id,required"`
	// Type of resource, always 'tool_group'
	Type constant.ToolGroup `json:"type,required"`
	// (Optional) Additional arguments for the tool group
	Args map[string]ToolGroupArgUnion `json:"args"`
	// (Optional) Model Context Protocol endpoint for remote tools
	McpEndpoint        ToolGroupMcpEndpoint `json:"mcp_endpoint"`
	ProviderResourceID string               `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Type               respjson.Field
		Args               respjson.Field
		McpEndpoint        respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolGroup) RawJSON() string { return r.JSON.raw }
func (r *ToolGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolGroupArgUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolGroupArgUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ToolGroupArgUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolGroupArgUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolGroupArgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolGroupArgUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolGroupArgUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolGroupArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Model Context Protocol endpoint for remote tools
type ToolGroupMcpEndpoint struct {
	// The URL string pointing to the resource
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

type ToolgroupRegisterParams struct {
	// The ID of the provider to use for the tool group.
	ProviderID string `json:"provider_id,required"`
	// The ID of the tool group to register.
	ToolgroupID string `json:"toolgroup_id,required"`
	// A dictionary of arguments to pass to the tool group.
	Args map[string]ToolgroupRegisterParamsArgUnion `json:"args,omitzero"`
	// The MCP endpoint to use for the tool group.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolgroupRegisterParamsArgUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolgroupRegisterParamsArgUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolgroupRegisterParamsArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolgroupRegisterParamsArgUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// The MCP endpoint to use for the tool group.
//
// The property Uri is required.
type ToolgroupRegisterParamsMcpEndpoint struct {
	// The URL string pointing to the resource
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
