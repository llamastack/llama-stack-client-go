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
}

// NewToolRuntimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolRuntimeService(opts ...option.RequestOption) (r ToolRuntimeService) {
	r = ToolRuntimeService{}
	r.Options = opts
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
	Name         string         `json:"name,required"`
	Description  string         `json:"description,nullable"`
	InputSchema  map[string]any `json:"input_schema,nullable"`
	Metadata     map[string]any `json:"metadata,nullable"`
	OutputSchema map[string]any `json:"output_schema,nullable"`
	ToolgroupID  string         `json:"toolgroup_id,nullable"`
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

// Result of a tool invocation.
type ToolInvocationResult struct {
	// A image content item
	Content      ToolInvocationResultContentUnion `json:"content,nullable"`
	ErrorCode    int64                            `json:"error_code,nullable"`
	ErrorMessage string                           `json:"error_message,nullable"`
	Metadata     map[string]any                   `json:"metadata,nullable"`
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

// ToolInvocationResultContentUnion contains all possible properties and values
// from [string], [ToolInvocationResultContentImageContentItemOutput],
// [ToolInvocationResultContentTextContentItem],
// [[]ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListImageContentItemOutputTextContentItem]
type ToolInvocationResultContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion]
	// instead of an object.
	OfListImageContentItemOutputTextContentItem []ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion `json:",inline"`
	// This field is from variant [ToolInvocationResultContentImageContentItemOutput].
	Image ToolInvocationResultContentImageContentItemOutputImage `json:"image"`
	Type  string                                                 `json:"type"`
	// This field is from variant [ToolInvocationResultContentTextContentItem].
	Text string `json:"text"`
	JSON struct {
		OfString                                    respjson.Field
		OfListImageContentItemOutputTextContentItem respjson.Field
		Image                                       respjson.Field
		Type                                        respjson.Field
		Text                                        respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u ToolInvocationResultContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultContentUnion) AsImageContentItemOutput() (v ToolInvocationResultContentImageContentItemOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultContentUnion) AsTextContentItem() (v ToolInvocationResultContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultContentUnion) AsListImageContentItemOutputTextContentItem() (v []ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolInvocationResultContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolInvocationResultContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A image content item
type ToolInvocationResultContentImageContentItemOutput struct {
	// A URL or a base64 encoded string
	Image ToolInvocationResultContentImageContentItemOutputImage `json:"image,required"`
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentImageContentItemOutput) RawJSON() string { return r.JSON.raw }
func (r *ToolInvocationResultContentImageContentItemOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL or a base64 encoded string
type ToolInvocationResultContentImageContentItemOutputImage struct {
	Data string `json:"data,nullable"`
	// A URL reference to external content.
	URL ToolInvocationResultContentImageContentItemOutputImageURL `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentImageContentItemOutputImage) RawJSON() string { return r.JSON.raw }
func (r *ToolInvocationResultContentImageContentItemOutputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
type ToolInvocationResultContentImageContentItemOutputImageURL struct {
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentImageContentItemOutputImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ToolInvocationResultContentImageContentItemOutputImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type ToolInvocationResultContentTextContentItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *ToolInvocationResultContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion
// contains all possible properties and values from
// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage],
// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText].
//
// Use the
// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion struct {
	// This field is from variant
	// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage].
	Image ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImage `json:"image"`
	// Any of "image", "text".
	Type string `json:"type"`
	// This field is from variant
	// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText].
	Text string `json:"text"`
	JSON struct {
		Image respjson.Field
		Type  respjson.Field
		Text  respjson.Field
		raw   string
	} `json:"-"`
}

// anyToolInvocationResultContentListImageContentItemOutputTextContentItemItem is
// implemented by each variant of
// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion]
// to add type safety for the return type of
// [ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion.AsAny]
type anyToolInvocationResultContentListImageContentItemOutputTextContentItemItem interface {
	implToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion()
}

func (ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage) implToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion() {
}
func (ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText) implToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion.AsAny().(type) {
//	case llamastackclient.ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage:
//	case llamastackclient.ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion) AsAny() anyToolInvocationResultContentListImageContentItemOutputTextContentItemItem {
	switch u.Type {
	case "image":
		return u.AsImage()
	case "text":
		return u.AsText()
	}
	return nil
}

func (u ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion) AsImage() (v ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion) AsText() (v ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ToolInvocationResultContentListImageContentItemOutputTextContentItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A image content item
type ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage struct {
	// A URL or a base64 encoded string
	Image ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImage `json:"image,required"`
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL or a base64 encoded string
type ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImage struct {
	Data string `json:"data,nullable"`
	// A URL reference to external content.
	URL ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImageURL `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
type ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImageURL struct {
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ToolInvocationResultContentListImageContentItemOutputTextContentItemItemImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ToolInvocationResultContentListImageContentItemOutputTextContentItemItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeInvokeToolParams struct {
	Kwargs        map[string]any    `json:"kwargs,omitzero,required"`
	ToolName      string            `json:"tool_name,required"`
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

type ToolRuntimeListToolsParams struct {
	Authorization param.Opt[string] `query:"authorization,omitzero" json:"-"`
	ToolGroupID   param.Opt[string] `query:"tool_group_id,omitzero" json:"-"`
	// A URL reference to external content.
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

// A URL reference to external content.
//
// The property Uri is required.
type ToolRuntimeListToolsParamsMcpEndpoint struct {
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
