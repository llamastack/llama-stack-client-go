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
	"net/http"
	"net/url"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ToolRuntimeService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolRuntimeService] method instead.
type ToolRuntimeService struct {
	Options []option.RequestOption
	RagTool ToolRuntimeRagToolService
}

// NewToolRuntimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolRuntimeService(opts ...option.RequestOption) (r ToolRuntimeService) {
	r = ToolRuntimeService{}
	r.Options = opts
	r.RagTool = NewToolRuntimeRagToolService(opts...)
	return
}

// Run a tool with the given arguments.
func (r *ToolRuntimeService) InvokeTool(ctx context.Context, body ToolRuntimeInvokeToolParams, opts ...option.RequestOption) (res *ToolInvocationResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tool-runtime/invoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all tools in the runtime.
func (r *ToolRuntimeService) ListTools(ctx context.Context, query ToolRuntimeListToolsParams, opts ...option.RequestOption) (res *[]ToolDef, err error) {
	var env ToolRuntimeListToolsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v1/tool-runtime/list-tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Tool definition used in runtime contexts.
type ToolDef struct {
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Human-readable description of what the tool does
	Description string `json:"description"`
	// (Optional) JSON Schema for tool inputs (MCP inputSchema)
	InputSchema map[string]ToolDefInputSchemaUnion `json:"input_schema"`
	// (Optional) Additional metadata about the tool
	Metadata map[string]ToolDefMetadataUnion `json:"metadata"`
	// (Optional) JSON Schema for tool outputs (MCP outputSchema)
	OutputSchema map[string]ToolDefOutputSchemaUnion `json:"output_schema"`
	// (Optional) ID of the tool group this tool belongs to
	ToolgroupID string `json:"toolgroup_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		Description  respjson.Field
		InputSchema  respjson.Field
		Metadata     respjson.Field
		OutputSchema respjson.Field
		ToolgroupID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDef) RawJSON() string { return r.JSON.raw }
func (r *ToolDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolDefInputSchemaUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolDefInputSchemaUnion struct {
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

func (u ToolDefInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolDefInputSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolDefInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolDefMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolDefMetadataUnion struct {
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

func (u ToolDefMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolDefMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolDefMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolDefOutputSchemaUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolDefOutputSchemaUnion struct {
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

func (u ToolDefOutputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefOutputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefOutputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefOutputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolDefOutputSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolDefOutputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a tool invocation.
type ToolInvocationResult struct {
	// (Optional) The output content from the tool execution
	Content InterleavedContentUnion `json:"content"`
	// (Optional) Numeric error code if the tool execution failed
	ErrorCode int64 `json:"error_code"`
	// (Optional) Error message if the tool execution failed
	ErrorMessage string `json:"error_message"`
	// (Optional) Additional metadata about the tool execution
	Metadata map[string]ToolInvocationResultMetadataUnion `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		Metadata     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResult) RawJSON() string { return r.JSON.raw }
func (r *ToolInvocationResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolInvocationResultMetadataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolInvocationResultMetadataUnion struct {
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

func (u ToolInvocationResultMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolInvocationResultMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolInvocationResultMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeInvokeToolParams struct {
	// A dictionary of arguments to pass to the tool.
	Kwargs map[string]ToolRuntimeInvokeToolParamsKwargUnion `json:"kwargs,omitzero,required"`
	// The name of the tool to invoke.
	ToolName string `json:"tool_name,required"`
	// (Optional) OAuth access token for authenticating with the MCP server.
	Authorization param.Opt[string] `json:"authorization,omitzero"`
	paramObj
}

func (r ToolRuntimeInvokeToolParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeInvokeToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeInvokeToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRuntimeInvokeToolParamsKwargUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRuntimeInvokeToolParamsKwargUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolRuntimeInvokeToolParamsKwargUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRuntimeInvokeToolParamsKwargUnion) asAny() any {
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

type ToolRuntimeListToolsParams struct {
	// (Optional) OAuth access token for authenticating with the MCP server.
	Authorization param.Opt[string] `query:"authorization,omitzero" json:"-"`
	// The ID of the tool group to list tools for.
	ToolGroupID param.Opt[string] `query:"tool_group_id,omitzero" json:"-"`
	// The MCP endpoint to use for the tool group.
	McpEndpoint ToolRuntimeListToolsParamsMcpEndpoint `query:"mcp_endpoint,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolRuntimeListToolsParams]'s query parameters as
// `url.Values`.
func (r ToolRuntimeListToolsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The MCP endpoint to use for the tool group.
//
// The property Uri is required.
type ToolRuntimeListToolsParamsMcpEndpoint struct {
	// The URL string pointing to the resource
	Uri string `query:"uri,required" json:"-"`
	paramObj
}

// URLQuery serializes [ToolRuntimeListToolsParamsMcpEndpoint]'s query parameters
// as `url.Values`.
func (r ToolRuntimeListToolsParamsMcpEndpoint) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response containing a list of tool definitions.
type ToolRuntimeListToolsResponseEnvelope struct {
	// List of tool definitions
	Data []ToolDef `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolRuntimeListToolsResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ToolRuntimeListToolsResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
