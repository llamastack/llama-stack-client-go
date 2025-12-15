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
	"net/url"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/pagination"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
)

// ResponseService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseService] method instead.
type ResponseService struct {
	Options    []option.RequestOption
	InputItems ResponseInputItemService
}

// NewResponseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResponseService(opts ...option.RequestOption) (r ResponseService) {
	r = ResponseService{}
	r.Options = opts
	r.InputItems = NewResponseInputItemService(opts...)
	return
}

// Create a model response.
func (r *ResponseService) New(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a model response.
func (r *ResponseService) NewStreaming(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ResponseObjectStreamUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ResponseObjectStreamUnion](ssestream.NewDecoder(raw), err)
}

// Get a model response.
func (r *ResponseService) Get(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all responses.
func (r *ResponseService) List(ctx context.Context, query ResponseListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[ResponseListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/responses"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all responses.
func (r *ResponseService) ListAutoPaging(ctx context.Context, query ResponseListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ResponseListResponse] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a response.
func (r *ResponseService) Delete(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Complete OpenAI response object containing generation results and metadata.
type ResponseObject struct {
	ID        string                      `json:"id,required"`
	CreatedAt int64                       `json:"created_at,required"`
	Model     string                      `json:"model,required"`
	Output    []ResponseObjectOutputUnion `json:"output,required"`
	Status    string                      `json:"status,required"`
	// Error details for failed OpenAI response requests.
	Error        ResponseObjectError `json:"error,nullable"`
	Instructions string              `json:"instructions,nullable"`
	MaxToolCalls int64               `json:"max_tool_calls,nullable"`
	Metadata     map[string]string   `json:"metadata,nullable"`
	// Any of "response".
	Object             ResponseObjectObject `json:"object"`
	ParallelToolCalls  bool                 `json:"parallel_tool_calls,nullable"`
	PreviousResponseID string               `json:"previous_response_id,nullable"`
	// OpenAI compatible Prompt object that is used in OpenAI responses.
	Prompt      ResponseObjectPrompt `json:"prompt,nullable"`
	Temperature float64              `json:"temperature,nullable"`
	// Text response configuration for OpenAI responses.
	Text ResponseObjectText `json:"text"`
	// Constrains the tools available to the model to a pre-defined set.
	ToolChoice ResponseObjectToolChoiceUnion `json:"tool_choice,nullable"`
	Tools      []ResponseObjectToolUnion     `json:"tools,nullable"`
	TopP       float64                       `json:"top_p,nullable"`
	Truncation string                        `json:"truncation,nullable"`
	// Usage information for OpenAI response.
	Usage ResponseObjectUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Model              respjson.Field
		Output             respjson.Field
		Status             respjson.Field
		Error              respjson.Field
		Instructions       respjson.Field
		MaxToolCalls       respjson.Field
		Metadata           respjson.Field
		Object             respjson.Field
		ParallelToolCalls  respjson.Field
		PreviousResponseID respjson.Field
		Prompt             respjson.Field
		Temperature        respjson.Field
		Text               respjson.Field
		ToolChoice         respjson.Field
		Tools              respjson.Field
		TopP               respjson.Field
		Truncation         respjson.Field
		Usage              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputUnion contains all possible properties and values from
// [ResponseObjectOutputMessage], [ResponseObjectOutputWebSearchCall],
// [ResponseObjectOutputFileSearchCall], [ResponseObjectOutputFunctionCall],
// [ResponseObjectOutputMcpCall], [ResponseObjectOutputMcpListTools],
// [ResponseObjectOutputMcpApprovalRequest].
//
// Use the [ResponseObjectOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputUnion struct {
	// This field is from variant [ResponseObjectOutputMessage].
	Content ResponseObjectOutputMessageContentUnion `json:"content"`
	// This field is from variant [ResponseObjectOutputMessage].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type string `json:"type"`
	// This field is from variant [ResponseObjectOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseObjectOutputFileSearchCall].
	Results   []ResponseObjectOutputFileSearchCallResult `json:"results"`
	Arguments string                                     `json:"arguments"`
	// This field is from variant [ResponseObjectOutputFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseObjectOutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [ResponseObjectOutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [ResponseObjectOutputMcpListTools].
	Tools []ResponseObjectOutputMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectOutput is implemented by each variant of
// [ResponseObjectOutputUnion] to add type safety for the return type of
// [ResponseObjectOutputUnion.AsAny]
type anyResponseObjectOutput interface {
	implResponseObjectOutputUnion()
}

func (ResponseObjectOutputMessage) implResponseObjectOutputUnion()            {}
func (ResponseObjectOutputWebSearchCall) implResponseObjectOutputUnion()      {}
func (ResponseObjectOutputFileSearchCall) implResponseObjectOutputUnion()     {}
func (ResponseObjectOutputFunctionCall) implResponseObjectOutputUnion()       {}
func (ResponseObjectOutputMcpCall) implResponseObjectOutputUnion()            {}
func (ResponseObjectOutputMcpListTools) implResponseObjectOutputUnion()       {}
func (ResponseObjectOutputMcpApprovalRequest) implResponseObjectOutputUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessage:
//	case llamastackclient.ResponseObjectOutputWebSearchCall:
//	case llamastackclient.ResponseObjectOutputFileSearchCall:
//	case llamastackclient.ResponseObjectOutputFunctionCall:
//	case llamastackclient.ResponseObjectOutputMcpCall:
//	case llamastackclient.ResponseObjectOutputMcpListTools:
//	case llamastackclient.ResponseObjectOutputMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputUnion) AsAny() anyResponseObjectOutput {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseObjectOutputUnion) AsMessage() (v ResponseObjectOutputMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsWebSearchCall() (v ResponseObjectOutputWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsFileSearchCall() (v ResponseObjectOutputFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsFunctionCall() (v ResponseObjectOutputFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsMcpCall() (v ResponseObjectOutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsMcpListTools() (v ResponseObjectOutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsMcpApprovalRequest() (v ResponseObjectOutputMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseObjectOutputMessage struct {
	Content ResponseObjectOutputMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string `json:"role,required"`
	ID     string `json:"id,nullable"`
	Status string `json:"status,nullable"`
	// Any of "message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentUnion contains all possible properties and
// values from [string],
// [[]ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ResponseObjectOutputMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseObjectOutputMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                  `json:"text,required"`
	Annotations []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                    `json:"token,required"`
	Logprob     float64                                                                                                                                                   `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                   `json:"bytes,nullable"`
	TopLogprobs []ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseObjectOutputWebSearchCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseObjectOutputFileSearchCall struct {
	ID      string                                     `json:"id,required"`
	Queries []string                                   `json:"queries,required"`
	Status  string                                     `json:"status,required"`
	Results []ResponseObjectOutputFileSearchCallResult `json:"results,nullable"`
	// Any of "file_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Results     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseObjectOutputFileSearchCallResult struct {
	Attributes map[string]any `json:"attributes,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseObjectOutputFunctionCall struct {
	Arguments string `json:"arguments,required"`
	CallID    string `json:"call_id,required"`
	Name      string `json:"name,required"`
	ID        string `json:"id,nullable"`
	Status    string `json:"status,nullable"`
	// Any of "function_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseObjectOutputMcpCall struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	Error       string `json:"error,nullable"`
	Output      string `json:"output,nullable"`
	// Any of "mcp_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseObjectOutputMcpListTools struct {
	ID          string                                 `json:"id,required"`
	ServerLabel string                                 `json:"server_label,required"`
	Tools       []ResponseObjectOutputMcpListToolsTool `json:"tools,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseObjectOutputMcpListToolsTool struct {
	InputSchema map[string]any `json:"input_schema,required"`
	Name        string         `json:"name,required"`
	Description string         `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseObjectOutputMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectOutputMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for failed OpenAI response requests.
type ResponseObjectError struct {
	Code    string `json:"code,required"`
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectError) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectObject string

const (
	ResponseObjectObjectResponse ResponseObjectObject = "response"
)

// OpenAI compatible Prompt object that is used in OpenAI responses.
type ResponseObjectPrompt struct {
	ID        string                                       `json:"id,required"`
	Variables map[string]ResponseObjectPromptVariableUnion `json:"variables,nullable"`
	Version   string                                       `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Variables   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectPrompt) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectPromptVariableUnion contains all possible properties and values
// from [ResponseObjectPromptVariableInputText],
// [ResponseObjectPromptVariableInputImage],
// [ResponseObjectPromptVariableInputFile].
//
// Use the [ResponseObjectPromptVariableUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectPromptVariableUnion struct {
	// This field is from variant [ResponseObjectPromptVariableInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant [ResponseObjectPromptVariableInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant [ResponseObjectPromptVariableInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant [ResponseObjectPromptVariableInputFile].
	FileData string `json:"file_data"`
	// This field is from variant [ResponseObjectPromptVariableInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant [ResponseObjectPromptVariableInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectPromptVariable is implemented by each variant of
// [ResponseObjectPromptVariableUnion] to add type safety for the return type of
// [ResponseObjectPromptVariableUnion.AsAny]
type anyResponseObjectPromptVariable interface {
	implResponseObjectPromptVariableUnion()
}

func (ResponseObjectPromptVariableInputText) implResponseObjectPromptVariableUnion()  {}
func (ResponseObjectPromptVariableInputImage) implResponseObjectPromptVariableUnion() {}
func (ResponseObjectPromptVariableInputFile) implResponseObjectPromptVariableUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectPromptVariableUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectPromptVariableInputText:
//	case llamastackclient.ResponseObjectPromptVariableInputImage:
//	case llamastackclient.ResponseObjectPromptVariableInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectPromptVariableUnion) AsAny() anyResponseObjectPromptVariable {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseObjectPromptVariableUnion) AsInputText() (v ResponseObjectPromptVariableInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectPromptVariableUnion) AsInputImage() (v ResponseObjectPromptVariableInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectPromptVariableUnion) AsInputFile() (v ResponseObjectPromptVariableInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectPromptVariableUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectPromptVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectPromptVariableInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseObjectPromptVariableInputText) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectPromptVariableInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectPromptVariableInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectPromptVariableInputImage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectPromptVariableInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseObjectPromptVariableInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectPromptVariableInputFile) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectPromptVariableInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text response configuration for OpenAI responses.
type ResponseObjectText struct {
	// Configuration for Responses API text format.
	Format ResponseObjectTextFormat `json:"format,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectText) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Responses API text format.
type ResponseObjectTextFormat struct {
	Description string         `json:"description,nullable"`
	Name        string         `json:"name,nullable"`
	Schema      map[string]any `json:"schema,nullable"`
	Strict      bool           `json:"strict,nullable"`
	// Any of "text", "json_schema", "json_object".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Schema      respjson.Field
		Strict      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolChoiceUnion contains all possible properties and values from
// [string], [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools],
// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFileSearch],
// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceWebSearch],
// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFunctionTool],
// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMcpTool],
// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceCustomTool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOpenAIResponseInputToolChoiceMode]
type ResponseObjectToolChoiceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfOpenAIResponseInputToolChoiceMode string `json:",inline"`
	// This field is from variant
	// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools].
	Tools []map[string]string `json:"tools"`
	// This field is from variant
	// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools].
	Mode string `json:"mode"`
	Type string `json:"type"`
	Name string `json:"name"`
	// This field is from variant
	// [ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMcpTool].
	ServerLabel string `json:"server_label"`
	JSON        struct {
		OfOpenAIResponseInputToolChoiceMode respjson.Field
		Tools                               respjson.Field
		Mode                                respjson.Field
		Type                                respjson.Field
		Name                                respjson.Field
		ServerLabel                         respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceMode() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceAllowedTools() (v ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceFileSearch() (v ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFileSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceWebSearch() (v ResponseObjectToolChoiceOpenAIResponseInputToolChoiceWebSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceFunctionTool() (v ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFunctionTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceMcpTool() (v ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMcpTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolChoiceUnion) AsOpenAIResponseInputToolChoiceCustomTool() (v ResponseObjectToolChoiceOpenAIResponseInputToolChoiceCustomTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolChoiceUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMode string

const (
	ResponseObjectToolChoiceOpenAIResponseInputToolChoiceModeAuto     ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMode = "auto"
	ResponseObjectToolChoiceOpenAIResponseInputToolChoiceModeRequired ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMode = "required"
	ResponseObjectToolChoiceOpenAIResponseInputToolChoiceModeNone     ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMode = "none"
)

// Constrains the tools available to the model to a pre-defined set.
type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools struct {
	Tools []map[string]string `json:"tools,required"`
	// Any of "auto", "required".
	Mode string `json:"mode"`
	// Any of "allowed_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tools       respjson.Field
		Mode        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectToolChoiceOpenAIResponseInputToolChoiceAllowedTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates that the model should use file search to generate a response.
type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFileSearch struct {
	// Any of "file_search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFileSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates that the model should use web search to generate a response
type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceWebSearch struct {
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11",
	// "web_search_2025_08_26".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolChoiceOpenAIResponseInputToolChoiceWebSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectToolChoiceOpenAIResponseInputToolChoiceWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forces the model to call a specific function.
type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFunctionTool struct {
	Name string `json:"name,required"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFunctionTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectToolChoiceOpenAIResponseInputToolChoiceFunctionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forces the model to call a specific tool on a remote MCP server
type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMcpTool struct {
	ServerLabel string `json:"server_label,required"`
	Name        string `json:"name,nullable"`
	// Any of "mcp".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServerLabel respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMcpTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectToolChoiceOpenAIResponseInputToolChoiceMcpTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forces the model to call a custom tool.
type ResponseObjectToolChoiceOpenAIResponseInputToolChoiceCustomTool struct {
	Name string `json:"name,required"`
	// Any of "custom".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolChoiceOpenAIResponseInputToolChoiceCustomTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectToolChoiceOpenAIResponseInputToolChoiceCustomTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolUnion contains all possible properties and values from
// [ResponseObjectToolOpenAIResponseInputToolWebSearch],
// [ResponseObjectToolFileSearch], [ResponseObjectToolFunction],
// [ResponseObjectToolMcp].
//
// Use the [ResponseObjectToolUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectToolUnion struct {
	// This field is from variant [ResponseObjectToolOpenAIResponseInputToolWebSearch].
	SearchContextSize string `json:"search_context_size"`
	// Any of nil, "file_search", "function", "mcp".
	Type string `json:"type"`
	// This field is from variant [ResponseObjectToolFileSearch].
	VectorStoreIDs []string `json:"vector_store_ids"`
	// This field is from variant [ResponseObjectToolFileSearch].
	Filters map[string]any `json:"filters"`
	// This field is from variant [ResponseObjectToolFileSearch].
	MaxNumResults int64 `json:"max_num_results"`
	// This field is from variant [ResponseObjectToolFileSearch].
	RankingOptions ResponseObjectToolFileSearchRankingOptions `json:"ranking_options"`
	// This field is from variant [ResponseObjectToolFunction].
	Name string `json:"name"`
	// This field is from variant [ResponseObjectToolFunction].
	Parameters map[string]any `json:"parameters"`
	// This field is from variant [ResponseObjectToolFunction].
	Description string `json:"description"`
	// This field is from variant [ResponseObjectToolFunction].
	Strict bool `json:"strict"`
	// This field is from variant [ResponseObjectToolMcp].
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseObjectToolMcp].
	AllowedTools ResponseObjectToolMcpAllowedToolsUnion `json:"allowed_tools"`
	JSON         struct {
		SearchContextSize respjson.Field
		Type              respjson.Field
		VectorStoreIDs    respjson.Field
		Filters           respjson.Field
		MaxNumResults     respjson.Field
		RankingOptions    respjson.Field
		Name              respjson.Field
		Parameters        respjson.Field
		Description       respjson.Field
		Strict            respjson.Field
		ServerLabel       respjson.Field
		AllowedTools      respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseObjectToolUnion) AsOpenAIResponseInputToolWebSearch() (v ResponseObjectToolOpenAIResponseInputToolWebSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolUnion) AsFileSearch() (v ResponseObjectToolFileSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolUnion) AsFunction() (v ResponseObjectToolFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolUnion) AsMcp() (v ResponseObjectToolMcp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool configuration for OpenAI response inputs.
type ResponseObjectToolOpenAIResponseInputToolWebSearch struct {
	SearchContextSize string `json:"search_context_size,nullable"`
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11",
	// "web_search_2025_08_26".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SearchContextSize respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolOpenAIResponseInputToolWebSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool configuration for OpenAI response inputs.
type ResponseObjectToolFileSearch struct {
	VectorStoreIDs []string       `json:"vector_store_ids,required"`
	Filters        map[string]any `json:"filters,nullable"`
	MaxNumResults  int64          `json:"max_num_results,nullable"`
	// Options for ranking and filtering search results.
	RankingOptions ResponseObjectToolFileSearchRankingOptions `json:"ranking_options,nullable"`
	// Any of "file_search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VectorStoreIDs respjson.Field
		Filters        respjson.Field
		MaxNumResults  respjson.Field
		RankingOptions respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolFileSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for ranking and filtering search results.
type ResponseObjectToolFileSearchRankingOptions struct {
	Ranker         string  `json:"ranker,nullable"`
	ScoreThreshold float64 `json:"score_threshold,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ranker         respjson.Field
		ScoreThreshold respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolFileSearchRankingOptions) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool configuration for OpenAI response inputs.
type ResponseObjectToolFunction struct {
	Name        string         `json:"name,required"`
	Parameters  map[string]any `json:"parameters,required"`
	Description string         `json:"description,nullable"`
	Strict      bool           `json:"strict,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Parameters  respjson.Field
		Description respjson.Field
		Strict      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolFunction) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) tool configuration for OpenAI response object.
type ResponseObjectToolMcp struct {
	ServerLabel string `json:"server_label,required"`
	// Filter configuration for restricting which MCP tools can be used.
	AllowedTools ResponseObjectToolMcpAllowedToolsUnion `json:"allowed_tools,nullable"`
	// Any of "mcp".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServerLabel  respjson.Field
		AllowedTools respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolMcp) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolMcpAllowedToolsUnion contains all possible properties and
// values from [[]string], [ResponseObjectToolMcpAllowedToolsAllowedToolsFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfListString]
type ResponseObjectToolMcpAllowedToolsUnion struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfListString []string `json:",inline"`
	// This field is from variant
	// [ResponseObjectToolMcpAllowedToolsAllowedToolsFilter].
	ToolNames []string `json:"tool_names"`
	JSON      struct {
		OfListString respjson.Field
		ToolNames    respjson.Field
		raw          string
	} `json:"-"`
}

func (u ResponseObjectToolMcpAllowedToolsUnion) AsListString() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolMcpAllowedToolsUnion) AsAllowedToolsFilter() (v ResponseObjectToolMcpAllowedToolsAllowedToolsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolMcpAllowedToolsUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter configuration for restricting which MCP tools can be used.
type ResponseObjectToolMcpAllowedToolsAllowedToolsFilter struct {
	ToolNames []string `json:"tool_names,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolNames   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolMcpAllowedToolsAllowedToolsFilter) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage information for OpenAI response.
type ResponseObjectUsage struct {
	InputTokens  int64 `json:"input_tokens,required"`
	OutputTokens int64 `json:"output_tokens,required"`
	TotalTokens  int64 `json:"total_tokens,required"`
	// Token details for input tokens in OpenAI response usage.
	InputTokensDetails ResponseObjectUsageInputTokensDetails `json:"input_tokens_details,nullable"`
	// Token details for output tokens in OpenAI response usage.
	OutputTokensDetails ResponseObjectUsageOutputTokensDetails `json:"output_tokens_details,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		OutputTokens        respjson.Field
		TotalTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokensDetails respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectUsage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for input tokens in OpenAI response usage.
type ResponseObjectUsageInputTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for output tokens in OpenAI response usage.
type ResponseObjectUsageOutputTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnion contains all possible properties and values from
// [ResponseObjectStreamResponseCreated], [ResponseObjectStreamResponseInProgress],
// [ResponseObjectStreamResponseOutputItemAdded],
// [ResponseObjectStreamResponseOutputItemDone],
// [ResponseObjectStreamResponseOutputTextDelta],
// [ResponseObjectStreamResponseOutputTextDone],
// [ResponseObjectStreamResponseFunctionCallArgumentsDelta],
// [ResponseObjectStreamResponseFunctionCallArgumentsDone],
// [ResponseObjectStreamResponseWebSearchCallInProgress],
// [ResponseObjectStreamResponseWebSearchCallSearching],
// [ResponseObjectStreamResponseWebSearchCallCompleted],
// [ResponseObjectStreamResponseMcpListToolsInProgress],
// [ResponseObjectStreamResponseMcpListToolsFailed],
// [ResponseObjectStreamResponseMcpListToolsCompleted],
// [ResponseObjectStreamResponseMcpCallArgumentsDelta],
// [ResponseObjectStreamResponseMcpCallArgumentsDone],
// [ResponseObjectStreamResponseMcpCallInProgress],
// [ResponseObjectStreamResponseMcpCallFailed],
// [ResponseObjectStreamResponseMcpCallCompleted],
// [ResponseObjectStreamResponseContentPartAdded],
// [ResponseObjectStreamResponseContentPartDone],
// [ResponseObjectStreamResponseReasoningTextDelta],
// [ResponseObjectStreamResponseReasoningTextDone],
// [ResponseObjectStreamResponseReasoningSummaryPartAdded],
// [ResponseObjectStreamResponseReasoningSummaryPartDone],
// [ResponseObjectStreamResponseReasoningSummaryTextDelta],
// [ResponseObjectStreamResponseReasoningSummaryTextDone],
// [ResponseObjectStreamResponseRefusalDelta],
// [ResponseObjectStreamResponseRefusalDone],
// [ResponseObjectStreamResponseOutputTextAnnotationAdded],
// [ResponseObjectStreamResponseFileSearchCallInProgress],
// [ResponseObjectStreamResponseFileSearchCallSearching],
// [ResponseObjectStreamResponseFileSearchCallCompleted],
// [ResponseObjectStreamResponseIncomplete], [ResponseObjectStreamResponseFailed],
// [ResponseObjectStreamResponseCompleted].
//
// Use the [ResponseObjectStreamUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamUnion struct {
	// This field is from variant [ResponseObjectStreamResponseCreated].
	Response ResponseObject `json:"response"`
	// Any of "response.created", "response.in_progress", "response.output_item.added",
	// "response.output_item.done", "response.output_text.delta",
	// "response.output_text.done", "response.function_call_arguments.delta",
	// "response.function_call_arguments.done", "response.web_search_call.in_progress",
	// "response.web_search_call.searching", "response.web_search_call.completed",
	// "response.mcp_list_tools.in_progress", "response.mcp_list_tools.failed",
	// "response.mcp_list_tools.completed", "response.mcp_call.arguments.delta",
	// "response.mcp_call.arguments.done", "response.mcp_call.in_progress",
	// "response.mcp_call.failed", "response.mcp_call.completed",
	// "response.content_part.added", "response.content_part.done",
	// "response.reasoning_text.delta", "response.reasoning_text.done",
	// "response.reasoning_summary_part.added", "response.reasoning_summary_part.done",
	// "response.reasoning_summary_text.delta", "response.reasoning_summary_text.done",
	// "response.refusal.delta", "response.refusal.done",
	// "response.output_text.annotation.added",
	// "response.file_search_call.in_progress", "response.file_search_call.searching",
	// "response.file_search_call.completed", "response.incomplete", "response.failed",
	// "response.completed".
	Type           string `json:"type"`
	SequenceNumber int64  `json:"sequence_number"`
	// This field is a union of [ResponseObjectStreamResponseOutputItemAddedItemUnion],
	// [ResponseObjectStreamResponseOutputItemDoneItemUnion]
	Item         ResponseObjectStreamUnionItem `json:"item"`
	OutputIndex  int64                         `json:"output_index"`
	ResponseID   string                        `json:"response_id"`
	ContentIndex int64                         `json:"content_index"`
	Delta        string                        `json:"delta"`
	ItemID       string                        `json:"item_id"`
	// This field is from variant [ResponseObjectStreamResponseOutputTextDelta].
	Logprobs  []ResponseObjectStreamResponseOutputTextDeltaLogprob `json:"logprobs"`
	Text      string                                               `json:"text"`
	Arguments string                                               `json:"arguments"`
	// This field is a union of
	// [ResponseObjectStreamResponseContentPartAddedPartUnion],
	// [ResponseObjectStreamResponseContentPartDonePartUnion],
	// [ResponseObjectStreamResponseReasoningSummaryPartAddedPart],
	// [ResponseObjectStreamResponseReasoningSummaryPartDonePart]
	Part         ResponseObjectStreamUnionPart `json:"part"`
	SummaryIndex int64                         `json:"summary_index"`
	// This field is from variant [ResponseObjectStreamResponseRefusalDone].
	Refusal string `json:"refusal"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputTextAnnotationAdded].
	Annotation ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion `json:"annotation"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputTextAnnotationAdded].
	AnnotationIndex int64 `json:"annotation_index"`
	JSON            struct {
		Response        respjson.Field
		Type            respjson.Field
		SequenceNumber  respjson.Field
		Item            respjson.Field
		OutputIndex     respjson.Field
		ResponseID      respjson.Field
		ContentIndex    respjson.Field
		Delta           respjson.Field
		ItemID          respjson.Field
		Logprobs        respjson.Field
		Text            respjson.Field
		Arguments       respjson.Field
		Part            respjson.Field
		SummaryIndex    respjson.Field
		Refusal         respjson.Field
		Annotation      respjson.Field
		AnnotationIndex respjson.Field
		raw             string
	} `json:"-"`
}

// anyResponseObjectStream is implemented by each variant of
// [ResponseObjectStreamUnion] to add type safety for the return type of
// [ResponseObjectStreamUnion.AsAny]
type anyResponseObjectStream interface {
	implResponseObjectStreamUnion()
}

func (ResponseObjectStreamResponseCreated) implResponseObjectStreamUnion()                    {}
func (ResponseObjectStreamResponseInProgress) implResponseObjectStreamUnion()                 {}
func (ResponseObjectStreamResponseOutputItemAdded) implResponseObjectStreamUnion()            {}
func (ResponseObjectStreamResponseOutputItemDone) implResponseObjectStreamUnion()             {}
func (ResponseObjectStreamResponseOutputTextDelta) implResponseObjectStreamUnion()            {}
func (ResponseObjectStreamResponseOutputTextDone) implResponseObjectStreamUnion()             {}
func (ResponseObjectStreamResponseFunctionCallArgumentsDelta) implResponseObjectStreamUnion() {}
func (ResponseObjectStreamResponseFunctionCallArgumentsDone) implResponseObjectStreamUnion()  {}
func (ResponseObjectStreamResponseWebSearchCallInProgress) implResponseObjectStreamUnion()    {}
func (ResponseObjectStreamResponseWebSearchCallSearching) implResponseObjectStreamUnion()     {}
func (ResponseObjectStreamResponseWebSearchCallCompleted) implResponseObjectStreamUnion()     {}
func (ResponseObjectStreamResponseMcpListToolsInProgress) implResponseObjectStreamUnion()     {}
func (ResponseObjectStreamResponseMcpListToolsFailed) implResponseObjectStreamUnion()         {}
func (ResponseObjectStreamResponseMcpListToolsCompleted) implResponseObjectStreamUnion()      {}
func (ResponseObjectStreamResponseMcpCallArgumentsDelta) implResponseObjectStreamUnion()      {}
func (ResponseObjectStreamResponseMcpCallArgumentsDone) implResponseObjectStreamUnion()       {}
func (ResponseObjectStreamResponseMcpCallInProgress) implResponseObjectStreamUnion()          {}
func (ResponseObjectStreamResponseMcpCallFailed) implResponseObjectStreamUnion()              {}
func (ResponseObjectStreamResponseMcpCallCompleted) implResponseObjectStreamUnion()           {}
func (ResponseObjectStreamResponseContentPartAdded) implResponseObjectStreamUnion()           {}
func (ResponseObjectStreamResponseContentPartDone) implResponseObjectStreamUnion()            {}
func (ResponseObjectStreamResponseReasoningTextDelta) implResponseObjectStreamUnion()         {}
func (ResponseObjectStreamResponseReasoningTextDone) implResponseObjectStreamUnion()          {}
func (ResponseObjectStreamResponseReasoningSummaryPartAdded) implResponseObjectStreamUnion()  {}
func (ResponseObjectStreamResponseReasoningSummaryPartDone) implResponseObjectStreamUnion()   {}
func (ResponseObjectStreamResponseReasoningSummaryTextDelta) implResponseObjectStreamUnion()  {}
func (ResponseObjectStreamResponseReasoningSummaryTextDone) implResponseObjectStreamUnion()   {}
func (ResponseObjectStreamResponseRefusalDelta) implResponseObjectStreamUnion()               {}
func (ResponseObjectStreamResponseRefusalDone) implResponseObjectStreamUnion()                {}
func (ResponseObjectStreamResponseOutputTextAnnotationAdded) implResponseObjectStreamUnion()  {}
func (ResponseObjectStreamResponseFileSearchCallInProgress) implResponseObjectStreamUnion()   {}
func (ResponseObjectStreamResponseFileSearchCallSearching) implResponseObjectStreamUnion()    {}
func (ResponseObjectStreamResponseFileSearchCallCompleted) implResponseObjectStreamUnion()    {}
func (ResponseObjectStreamResponseIncomplete) implResponseObjectStreamUnion()                 {}
func (ResponseObjectStreamResponseFailed) implResponseObjectStreamUnion()                     {}
func (ResponseObjectStreamResponseCompleted) implResponseObjectStreamUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseCreated:
//	case llamastackclient.ResponseObjectStreamResponseInProgress:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAdded:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDone:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextDelta:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextDone:
//	case llamastackclient.ResponseObjectStreamResponseFunctionCallArgumentsDelta:
//	case llamastackclient.ResponseObjectStreamResponseFunctionCallArgumentsDone:
//	case llamastackclient.ResponseObjectStreamResponseWebSearchCallInProgress:
//	case llamastackclient.ResponseObjectStreamResponseWebSearchCallSearching:
//	case llamastackclient.ResponseObjectStreamResponseWebSearchCallCompleted:
//	case llamastackclient.ResponseObjectStreamResponseMcpListToolsInProgress:
//	case llamastackclient.ResponseObjectStreamResponseMcpListToolsFailed:
//	case llamastackclient.ResponseObjectStreamResponseMcpListToolsCompleted:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallArgumentsDelta:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallArgumentsDone:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallInProgress:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallFailed:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallCompleted:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAdded:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDone:
//	case llamastackclient.ResponseObjectStreamResponseReasoningTextDelta:
//	case llamastackclient.ResponseObjectStreamResponseReasoningTextDone:
//	case llamastackclient.ResponseObjectStreamResponseReasoningSummaryPartAdded:
//	case llamastackclient.ResponseObjectStreamResponseReasoningSummaryPartDone:
//	case llamastackclient.ResponseObjectStreamResponseReasoningSummaryTextDelta:
//	case llamastackclient.ResponseObjectStreamResponseReasoningSummaryTextDone:
//	case llamastackclient.ResponseObjectStreamResponseRefusalDelta:
//	case llamastackclient.ResponseObjectStreamResponseRefusalDone:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextAnnotationAdded:
//	case llamastackclient.ResponseObjectStreamResponseFileSearchCallInProgress:
//	case llamastackclient.ResponseObjectStreamResponseFileSearchCallSearching:
//	case llamastackclient.ResponseObjectStreamResponseFileSearchCallCompleted:
//	case llamastackclient.ResponseObjectStreamResponseIncomplete:
//	case llamastackclient.ResponseObjectStreamResponseFailed:
//	case llamastackclient.ResponseObjectStreamResponseCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamUnion) AsAny() anyResponseObjectStream {
	switch u.Type {
	case "response.created":
		return u.AsResponseCreated()
	case "response.in_progress":
		return u.AsResponseInProgress()
	case "response.output_item.added":
		return u.AsResponseOutputItemAdded()
	case "response.output_item.done":
		return u.AsResponseOutputItemDone()
	case "response.output_text.delta":
		return u.AsResponseOutputTextDelta()
	case "response.output_text.done":
		return u.AsResponseOutputTextDone()
	case "response.function_call_arguments.delta":
		return u.AsResponseFunctionCallArgumentsDelta()
	case "response.function_call_arguments.done":
		return u.AsResponseFunctionCallArgumentsDone()
	case "response.web_search_call.in_progress":
		return u.AsResponseWebSearchCallInProgress()
	case "response.web_search_call.searching":
		return u.AsResponseWebSearchCallSearching()
	case "response.web_search_call.completed":
		return u.AsResponseWebSearchCallCompleted()
	case "response.mcp_list_tools.in_progress":
		return u.AsResponseMcpListToolsInProgress()
	case "response.mcp_list_tools.failed":
		return u.AsResponseMcpListToolsFailed()
	case "response.mcp_list_tools.completed":
		return u.AsResponseMcpListToolsCompleted()
	case "response.mcp_call.arguments.delta":
		return u.AsResponseMcpCallArgumentsDelta()
	case "response.mcp_call.arguments.done":
		return u.AsResponseMcpCallArgumentsDone()
	case "response.mcp_call.in_progress":
		return u.AsResponseMcpCallInProgress()
	case "response.mcp_call.failed":
		return u.AsResponseMcpCallFailed()
	case "response.mcp_call.completed":
		return u.AsResponseMcpCallCompleted()
	case "response.content_part.added":
		return u.AsResponseContentPartAdded()
	case "response.content_part.done":
		return u.AsResponseContentPartDone()
	case "response.reasoning_text.delta":
		return u.AsResponseReasoningTextDelta()
	case "response.reasoning_text.done":
		return u.AsResponseReasoningTextDone()
	case "response.reasoning_summary_part.added":
		return u.AsResponseReasoningSummaryPartAdded()
	case "response.reasoning_summary_part.done":
		return u.AsResponseReasoningSummaryPartDone()
	case "response.reasoning_summary_text.delta":
		return u.AsResponseReasoningSummaryTextDelta()
	case "response.reasoning_summary_text.done":
		return u.AsResponseReasoningSummaryTextDone()
	case "response.refusal.delta":
		return u.AsResponseRefusalDelta()
	case "response.refusal.done":
		return u.AsResponseRefusalDone()
	case "response.output_text.annotation.added":
		return u.AsResponseOutputTextAnnotationAdded()
	case "response.file_search_call.in_progress":
		return u.AsResponseFileSearchCallInProgress()
	case "response.file_search_call.searching":
		return u.AsResponseFileSearchCallSearching()
	case "response.file_search_call.completed":
		return u.AsResponseFileSearchCallCompleted()
	case "response.incomplete":
		return u.AsResponseIncomplete()
	case "response.failed":
		return u.AsResponseFailed()
	case "response.completed":
		return u.AsResponseCompleted()
	}
	return nil
}

func (u ResponseObjectStreamUnion) AsResponseCreated() (v ResponseObjectStreamResponseCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseInProgress() (v ResponseObjectStreamResponseInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputItemAdded() (v ResponseObjectStreamResponseOutputItemAdded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputItemDone() (v ResponseObjectStreamResponseOutputItemDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputTextDelta() (v ResponseObjectStreamResponseOutputTextDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputTextDone() (v ResponseObjectStreamResponseOutputTextDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFunctionCallArgumentsDelta() (v ResponseObjectStreamResponseFunctionCallArgumentsDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFunctionCallArgumentsDone() (v ResponseObjectStreamResponseFunctionCallArgumentsDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseWebSearchCallInProgress() (v ResponseObjectStreamResponseWebSearchCallInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseWebSearchCallSearching() (v ResponseObjectStreamResponseWebSearchCallSearching) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseWebSearchCallCompleted() (v ResponseObjectStreamResponseWebSearchCallCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpListToolsInProgress() (v ResponseObjectStreamResponseMcpListToolsInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpListToolsFailed() (v ResponseObjectStreamResponseMcpListToolsFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpListToolsCompleted() (v ResponseObjectStreamResponseMcpListToolsCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallArgumentsDelta() (v ResponseObjectStreamResponseMcpCallArgumentsDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallArgumentsDone() (v ResponseObjectStreamResponseMcpCallArgumentsDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallInProgress() (v ResponseObjectStreamResponseMcpCallInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallFailed() (v ResponseObjectStreamResponseMcpCallFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallCompleted() (v ResponseObjectStreamResponseMcpCallCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseContentPartAdded() (v ResponseObjectStreamResponseContentPartAdded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseContentPartDone() (v ResponseObjectStreamResponseContentPartDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseReasoningTextDelta() (v ResponseObjectStreamResponseReasoningTextDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseReasoningTextDone() (v ResponseObjectStreamResponseReasoningTextDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseReasoningSummaryPartAdded() (v ResponseObjectStreamResponseReasoningSummaryPartAdded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseReasoningSummaryPartDone() (v ResponseObjectStreamResponseReasoningSummaryPartDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseReasoningSummaryTextDelta() (v ResponseObjectStreamResponseReasoningSummaryTextDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseReasoningSummaryTextDone() (v ResponseObjectStreamResponseReasoningSummaryTextDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseRefusalDelta() (v ResponseObjectStreamResponseRefusalDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseRefusalDone() (v ResponseObjectStreamResponseRefusalDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputTextAnnotationAdded() (v ResponseObjectStreamResponseOutputTextAnnotationAdded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFileSearchCallInProgress() (v ResponseObjectStreamResponseFileSearchCallInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFileSearchCallSearching() (v ResponseObjectStreamResponseFileSearchCallSearching) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFileSearchCallCompleted() (v ResponseObjectStreamResponseFileSearchCallCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseIncomplete() (v ResponseObjectStreamResponseIncomplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFailed() (v ResponseObjectStreamResponseFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseCompleted() (v ResponseObjectStreamResponseCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItem is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItem provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
type ResponseObjectStreamUnionItem struct {
	// This field is a union of
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion],
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion]
	Content ResponseObjectStreamUnionItemContent `json:"content"`
	Role    string                               `json:"role"`
	ID      string                               `json:"id"`
	Status  string                               `json:"status"`
	Type    string                               `json:"type"`
	Queries []string                             `json:"queries"`
	// This field is a union of
	// [[]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult],
	// [[]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult]
	Results     ResponseObjectStreamUnionItemResults `json:"results"`
	Arguments   string                               `json:"arguments"`
	CallID      string                               `json:"call_id"`
	Name        string                               `json:"name"`
	ServerLabel string                               `json:"server_label"`
	Error       string                               `json:"error"`
	Output      string                               `json:"output"`
	// This field is a union of
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool],
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool]
	Tools ResponseObjectStreamUnionItemTools `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItemContent is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItemContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ResponseObjectStreamUnionItemContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItemContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItemResults is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItemResults provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResults
// OfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResults]
type ResponseObjectStreamUnionItemResults struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult] instead
	// of an object.
	OfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResults []ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult] instead
	// of an object.
	OfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResults []ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult `json:",inline"`
	JSON                                                                  struct {
		OfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResults respjson.Field
		OfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResults  respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItemResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItemTools is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItemTools provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTools]
type ResponseObjectStreamUnionItemTools struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool] instead of
	// an object.
	OfTools []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:",inline"`
	JSON    struct {
		OfTools respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItemTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionPart is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionPart provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
type ResponseObjectStreamUnionPart struct {
	Text string `json:"text"`
	// This field is a union of
	// [[]ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion],
	// [[]ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion]
	Annotations ResponseObjectStreamUnionPartAnnotations `json:"annotations"`
	// This field is a union of
	// [[]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob],
	// [[]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob]
	Logprobs ResponseObjectStreamUnionPartLogprobs `json:"logprobs"`
	Type     string                                `json:"type"`
	Refusal  string                                `json:"refusal"`
	JSON     struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionPartAnnotations is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionPartAnnotations provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnnotations]
type ResponseObjectStreamUnionPartAnnotations struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion]
	// instead of an object.
	OfAnnotations []ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion `json:",inline"`
	JSON          struct {
		OfAnnotations respjson.Field
		raw           string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPartAnnotations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionPartLogprobs is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionPartLogprobs provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobs
// OfResponseObjectStreamResponseContentPartDonePartOutputTextLogprobs]
type ResponseObjectStreamUnionPartLogprobs struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob] instead of
	// an object.
	OfResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobs []ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob] instead of
	// an object.
	OfResponseObjectStreamResponseContentPartDonePartOutputTextLogprobs []ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob `json:",inline"`
	JSON                                                                struct {
		OfResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobs respjson.Field
		OfResponseObjectStreamResponseContentPartDonePartOutputTextLogprobs  respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPartLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating a new response has been created.
type ResponseObjectStreamResponseCreated struct {
	// Complete OpenAI response object containing generation results and metadata.
	Response ResponseObject `json:"response,required"`
	// Any of "response.created".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseCreated) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating the response remains in progress.
type ResponseObjectStreamResponseInProgress struct {
	// Complete OpenAI response object containing generation results and metadata.
	Response       ResponseObject `json:"response,required"`
	SequenceNumber int64          `json:"sequence_number,required"`
	// Any of "response.in_progress".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response       respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a new output item is added to the response.
type ResponseObjectStreamResponseOutputItemAdded struct {
	// Corresponds to the various Message types in the Responses API. They are all
	// under one type because the Responses API gives them all the same "type" value,
	// and there is no way to tell them apart in certain scenarios.
	Item           ResponseObjectStreamResponseOutputItemAddedItemUnion `json:"item,required"`
	OutputIndex    int64                                                `json:"output_index,required"`
	ResponseID     string                                               `json:"response_id,required"`
	SequenceNumber int64                                                `json:"sequence_number,required"`
	// Any of "response.output_item.added".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item           respjson.Field
		OutputIndex    respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAdded) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessage],
// [ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall],
// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall],
// [ResponseObjectStreamResponseOutputItemAddedItemFunctionCall],
// [ResponseObjectStreamResponseOutputItemAddedItemMcpCall],
// [ResponseObjectStreamResponseOutputItemAddedItemMcpListTools],
// [ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest].
//
// Use the [ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessage].
	Content ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion `json:"content"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessage].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall].
	Results   []ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult `json:"results"`
	Arguments string                                                                `json:"arguments"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMcpCall].
	Error string `json:"error"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMcpCall].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMcpListTools].
	Tools []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemAddedItem is implemented by each
// variant of [ResponseObjectStreamResponseOutputItemAddedItemUnion] to add type
// safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItem interface {
	implResponseObjectStreamResponseOutputItemAddedItemUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessage) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMcpCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemFunctionCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpListTools:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItem {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMessage() (v ResponseObjectStreamResponseOutputItemAddedItemMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsWebSearchCall() (v ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsFileSearchCall() (v ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsFunctionCall() (v ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMcpCall() (v ResponseObjectStreamResponseOutputItemAddedItemMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMcpListTools() (v ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMcpApprovalRequest() (v ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseOutputItemAddedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseObjectStreamResponseOutputItemAddedItemMessage struct {
	Content ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string `json:"role,required"`
	ID     string `json:"id,nullable"`
	Status string `json:"status,nullable"`
	// Any of "message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal() (v []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem interface {
	implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                                       `json:"text,required"`
	Annotations []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                                         `json:"token,required"`
	Logprob     float64                                                                                                                                                                        `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                                        `json:"bytes,nullable"`
	TopLogprobs []ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall struct {
	ID      string                                                                `json:"id,required"`
	Queries []string                                                              `json:"queries,required"`
	Status  string                                                                `json:"status,required"`
	Results []ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult `json:"results,nullable"`
	// Any of "file_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Results     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult struct {
	Attributes map[string]any `json:"attributes,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemFunctionCall struct {
	Arguments string `json:"arguments,required"`
	CallID    string `json:"call_id,required"`
	Name      string `json:"name,required"`
	ID        string `json:"id,nullable"`
	Status    string `json:"status,nullable"`
	// Any of "function_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemMcpCall struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	Error       string `json:"error,nullable"`
	Output      string `json:"output,nullable"`
	// Any of "mcp_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseObjectStreamResponseOutputItemAddedItemMcpListTools struct {
	ID          string                                                            `json:"id,required"`
	ServerLabel string                                                            `json:"server_label,required"`
	Tools       []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:"tools,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool struct {
	InputSchema map[string]any `json:"input_schema,required"`
	Name        string         `json:"name,required"`
	Description string         `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when an output item is completed.
type ResponseObjectStreamResponseOutputItemDone struct {
	// Corresponds to the various Message types in the Responses API. They are all
	// under one type because the Responses API gives them all the same "type" value,
	// and there is no way to tell them apart in certain scenarios.
	Item           ResponseObjectStreamResponseOutputItemDoneItemUnion `json:"item,required"`
	OutputIndex    int64                                               `json:"output_index,required"`
	ResponseID     string                                              `json:"response_id,required"`
	SequenceNumber int64                                               `json:"sequence_number,required"`
	// Any of "response.output_item.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item           respjson.Field
		OutputIndex    respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessage],
// [ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall],
// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall],
// [ResponseObjectStreamResponseOutputItemDoneItemFunctionCall],
// [ResponseObjectStreamResponseOutputItemDoneItemMcpCall],
// [ResponseObjectStreamResponseOutputItemDoneItemMcpListTools],
// [ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest].
//
// Use the [ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessage].
	Content ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion `json:"content"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessage].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall].
	Results   []ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult `json:"results"`
	Arguments string                                                               `json:"arguments"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMcpCall].
	Error string `json:"error"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMcpCall].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMcpListTools].
	Tools []ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemDoneItem is implemented by each variant
// of [ResponseObjectStreamResponseOutputItemDoneItemUnion] to add type safety for
// the return type of [ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItem interface {
	implResponseObjectStreamResponseOutputItemDoneItemUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessage) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMcpCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemFunctionCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpListTools:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItem {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMessage() (v ResponseObjectStreamResponseOutputItemDoneItemMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsWebSearchCall() (v ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsFileSearchCall() (v ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsFunctionCall() (v ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMcpCall() (v ResponseObjectStreamResponseOutputItemDoneItemMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMcpListTools() (v ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMcpApprovalRequest() (v ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseOutputItemDoneItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseObjectStreamResponseOutputItemDoneItemMessage struct {
	Content ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string `json:"role,required"`
	ID     string `json:"id,nullable"`
	Status string `json:"status,nullable"`
	// Any of "message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal() (v []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem interface {
	implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                                      `json:"text,required"`
	Annotations []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                                        `json:"token,required"`
	Logprob     float64                                                                                                                                                                       `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                                       `json:"bytes,nullable"`
	TopLogprobs []ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall struct {
	ID      string                                                               `json:"id,required"`
	Queries []string                                                             `json:"queries,required"`
	Status  string                                                               `json:"status,required"`
	Results []ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult `json:"results,nullable"`
	// Any of "file_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Results     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult struct {
	Attributes map[string]any `json:"attributes,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemFunctionCall struct {
	Arguments string `json:"arguments,required"`
	CallID    string `json:"call_id,required"`
	Name      string `json:"name,required"`
	ID        string `json:"id,nullable"`
	Status    string `json:"status,nullable"`
	// Any of "function_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemMcpCall struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	Error       string `json:"error,nullable"`
	Output      string `json:"output,nullable"`
	// Any of "mcp_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseObjectStreamResponseOutputItemDoneItemMcpListTools struct {
	ID          string                                                           `json:"id,required"`
	ServerLabel string                                                           `json:"server_label,required"`
	Tools       []ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool `json:"tools,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool struct {
	InputSchema map[string]any `json:"input_schema,required"`
	Name        string         `json:"name,required"`
	Description string         `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for incremental text content updates.
type ResponseObjectStreamResponseOutputTextDelta struct {
	ContentIndex   int64                                                `json:"content_index,required"`
	Delta          string                                               `json:"delta,required"`
	ItemID         string                                               `json:"item_id,required"`
	OutputIndex    int64                                                `json:"output_index,required"`
	SequenceNumber int64                                                `json:"sequence_number,required"`
	Logprobs       []ResponseObjectStreamResponseOutputTextDeltaLogprob `json:"logprobs,nullable"`
	// Any of "response.output_text.delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Logprobs       respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputTextDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseObjectStreamResponseOutputTextDeltaLogprob struct {
	Token       string                                                         `json:"token,required"`
	Logprob     float64                                                        `json:"logprob,required"`
	Bytes       []int64                                                        `json:"bytes,nullable"`
	TopLogprobs []ResponseObjectStreamResponseOutputTextDeltaLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextDeltaLogprob) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputTextDeltaLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseObjectStreamResponseOutputTextDeltaLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextDeltaLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputTextDeltaLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when text output is completed.
type ResponseObjectStreamResponseOutputTextDone struct {
	ContentIndex   int64  `json:"content_index,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	Text           string `json:"text,required"`
	// Any of "response.output_text.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Text           respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputTextDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for incremental function call argument updates.
type ResponseObjectStreamResponseFunctionCallArgumentsDelta struct {
	Delta          string `json:"delta,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.function_call_arguments.delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFunctionCallArgumentsDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFunctionCallArgumentsDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when function call arguments are completed.
type ResponseObjectStreamResponseFunctionCallArgumentsDone struct {
	Arguments      string `json:"arguments,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.function_call_arguments.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments      respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFunctionCallArgumentsDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFunctionCallArgumentsDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for web search calls in progress.
type ResponseObjectStreamResponseWebSearchCallInProgress struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.web_search_call.in_progress".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseWebSearchCallInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseWebSearchCallInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseWebSearchCallSearching struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.web_search_call.searching".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseWebSearchCallSearching) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseWebSearchCallSearching) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for completed web search calls.
type ResponseObjectStreamResponseWebSearchCallCompleted struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.web_search_call.completed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseWebSearchCallCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseWebSearchCallCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpListToolsInProgress struct {
	SequenceNumber int64 `json:"sequence_number,required"`
	// Any of "response.mcp_list_tools.in_progress".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpListToolsInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpListToolsInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpListToolsFailed struct {
	SequenceNumber int64 `json:"sequence_number,required"`
	// Any of "response.mcp_list_tools.failed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpListToolsFailed) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpListToolsFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpListToolsCompleted struct {
	SequenceNumber int64 `json:"sequence_number,required"`
	// Any of "response.mcp_list_tools.completed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpListToolsCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpListToolsCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpCallArgumentsDelta struct {
	Delta          string `json:"delta,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.mcp_call.arguments.delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallArgumentsDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallArgumentsDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpCallArgumentsDone struct {
	Arguments      string `json:"arguments,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.mcp_call.arguments.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments      respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallArgumentsDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallArgumentsDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for MCP calls in progress.
type ResponseObjectStreamResponseMcpCallInProgress struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.mcp_call.in_progress".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for failed MCP calls.
type ResponseObjectStreamResponseMcpCallFailed struct {
	SequenceNumber int64 `json:"sequence_number,required"`
	// Any of "response.mcp_call.failed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallFailed) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for completed MCP calls.
type ResponseObjectStreamResponseMcpCallCompleted struct {
	SequenceNumber int64 `json:"sequence_number,required"`
	// Any of "response.mcp_call.completed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a new content part is added to a response item.
type ResponseObjectStreamResponseContentPartAdded struct {
	ContentIndex int64  `json:"content_index,required"`
	ItemID       string `json:"item_id,required"`
	OutputIndex  int64  `json:"output_index,required"`
	// Text content within a streamed response part.
	Part           ResponseObjectStreamResponseContentPartAddedPartUnion `json:"part,required"`
	ResponseID     string                                                `json:"response_id,required"`
	SequenceNumber int64                                                 `json:"sequence_number,required"`
	// Any of "response.content_part.added".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Part           respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAdded) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartAddedPartUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseContentPartAddedPartOutputText],
// [ResponseObjectStreamResponseContentPartAddedPartRefusal],
// [ResponseObjectStreamResponseContentPartAddedPartReasoningText].
//
// Use the [ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartAddedPartUnion struct {
	Text string `json:"text"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputText].
	Annotations []ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputText].
	Logprobs []ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal", "reasoning_text".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseContentPartAddedPart is implemented by each
// variant of [ResponseObjectStreamResponseContentPartAddedPartUnion] to add type
// safety for the return type of
// [ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny]
type anyResponseObjectStreamResponseContentPartAddedPart interface {
	implResponseObjectStreamResponseContentPartAddedPartUnion()
}

func (ResponseObjectStreamResponseContentPartAddedPartOutputText) implResponseObjectStreamResponseContentPartAddedPartUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartRefusal) implResponseObjectStreamResponseContentPartAddedPartUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartReasoningText) implResponseObjectStreamResponseContentPartAddedPartUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputText:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartRefusal:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartReasoningText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsAny() anyResponseObjectStreamResponseContentPartAddedPart {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	case "reasoning_text":
		return u.AsReasoningText()
	}
	return nil
}

func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsOutputText() (v ResponseObjectStreamResponseContentPartAddedPartOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsRefusal() (v ResponseObjectStreamResponseContentPartAddedPartRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsReasoningText() (v ResponseObjectStreamResponseContentPartAddedPartReasoningText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartAddedPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseContentPartAddedPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content within a streamed response part.
type ResponseObjectStreamResponseContentPartAddedPartOutputText struct {
	Text        string                                                                      `json:"text,required"`
	Annotations []ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation],
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation],
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotation is
// implemented by each variant of
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion] to
// add type safety for the return type of
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotation interface {
	implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion()
}

func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsAny() anyResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob struct {
	Token       string                                                                        `json:"token,required"`
	Logprob     float64                                                                       `json:"logprob,required"`
	Bytes       []int64                                                                       `json:"bytes,nullable"`
	TopLogprobs []ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectStreamResponseContentPartAddedPartRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartRefusal) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartAddedPartRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text emitted as part of a streamed response.
type ResponseObjectStreamResponseContentPartAddedPartReasoningText struct {
	Text string `json:"text,required"`
	// Any of "reasoning_text".
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
func (r ResponseObjectStreamResponseContentPartAddedPartReasoningText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartReasoningText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a content part is completed.
type ResponseObjectStreamResponseContentPartDone struct {
	ContentIndex int64  `json:"content_index,required"`
	ItemID       string `json:"item_id,required"`
	OutputIndex  int64  `json:"output_index,required"`
	// Text content within a streamed response part.
	Part           ResponseObjectStreamResponseContentPartDonePartUnion `json:"part,required"`
	ResponseID     string                                               `json:"response_id,required"`
	SequenceNumber int64                                                `json:"sequence_number,required"`
	// Any of "response.content_part.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Part           respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartDonePartUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseContentPartDonePartOutputText],
// [ResponseObjectStreamResponseContentPartDonePartRefusal],
// [ResponseObjectStreamResponseContentPartDonePartReasoningText].
//
// Use the [ResponseObjectStreamResponseContentPartDonePartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartDonePartUnion struct {
	Text string `json:"text"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputText].
	Annotations []ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputText].
	Logprobs []ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal", "reasoning_text".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseContentPartDonePart is implemented by each
// variant of [ResponseObjectStreamResponseContentPartDonePartUnion] to add type
// safety for the return type of
// [ResponseObjectStreamResponseContentPartDonePartUnion.AsAny]
type anyResponseObjectStreamResponseContentPartDonePart interface {
	implResponseObjectStreamResponseContentPartDonePartUnion()
}

func (ResponseObjectStreamResponseContentPartDonePartOutputText) implResponseObjectStreamResponseContentPartDonePartUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartRefusal) implResponseObjectStreamResponseContentPartDonePartUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartReasoningText) implResponseObjectStreamResponseContentPartDonePartUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartDonePartUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputText:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartRefusal:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartReasoningText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsAny() anyResponseObjectStreamResponseContentPartDonePart {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	case "reasoning_text":
		return u.AsReasoningText()
	}
	return nil
}

func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsOutputText() (v ResponseObjectStreamResponseContentPartDonePartOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsRefusal() (v ResponseObjectStreamResponseContentPartDonePartRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsReasoningText() (v ResponseObjectStreamResponseContentPartDonePartReasoningText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartDonePartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseContentPartDonePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content within a streamed response part.
type ResponseObjectStreamResponseContentPartDonePartOutputText struct {
	Text        string                                                                     `json:"text,required"`
	Annotations []ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation],
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation],
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseContentPartDonePartOutputTextAnnotation is
// implemented by each variant of
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion] to
// add type safety for the return type of
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseContentPartDonePartOutputTextAnnotation interface {
	implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion()
}

func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsAny() anyResponseObjectStreamResponseContentPartDonePartOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob struct {
	Token       string                                                                       `json:"token,required"`
	Logprob     float64                                                                      `json:"logprob,required"`
	Bytes       []int64                                                                      `json:"bytes,nullable"`
	TopLogprobs []ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectStreamResponseContentPartDonePartRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartRefusal) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartDonePartRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text emitted as part of a streamed response.
type ResponseObjectStreamResponseContentPartDonePartReasoningText struct {
	Text string `json:"text,required"`
	// Any of "reasoning_text".
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
func (r ResponseObjectStreamResponseContentPartDonePartReasoningText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartReasoningText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for incremental reasoning text updates.
type ResponseObjectStreamResponseReasoningTextDelta struct {
	ContentIndex   int64  `json:"content_index,required"`
	Delta          string `json:"delta,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.reasoning_text.delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseReasoningTextDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningTextDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when reasoning text is completed.
type ResponseObjectStreamResponseReasoningTextDone struct {
	ContentIndex   int64  `json:"content_index,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	Text           string `json:"text,required"`
	// Any of "response.reasoning_text.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Text           respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseReasoningTextDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningTextDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a new reasoning summary part is added.
type ResponseObjectStreamResponseReasoningSummaryPartAdded struct {
	ItemID      string `json:"item_id,required"`
	OutputIndex int64  `json:"output_index,required"`
	// Reasoning summary part in a streamed response.
	Part           ResponseObjectStreamResponseReasoningSummaryPartAddedPart `json:"part,required"`
	SequenceNumber int64                                                     `json:"sequence_number,required"`
	SummaryIndex   int64                                                     `json:"summary_index,required"`
	// Any of "response.reasoning_summary_part.added".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Part           respjson.Field
		SequenceNumber respjson.Field
		SummaryIndex   respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseReasoningSummaryPartAdded) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningSummaryPartAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning summary part in a streamed response.
type ResponseObjectStreamResponseReasoningSummaryPartAddedPart struct {
	Text string `json:"text,required"`
	// Any of "summary_text".
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
func (r ResponseObjectStreamResponseReasoningSummaryPartAddedPart) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseReasoningSummaryPartAddedPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a reasoning summary part is completed.
type ResponseObjectStreamResponseReasoningSummaryPartDone struct {
	ItemID      string `json:"item_id,required"`
	OutputIndex int64  `json:"output_index,required"`
	// Reasoning summary part in a streamed response.
	Part           ResponseObjectStreamResponseReasoningSummaryPartDonePart `json:"part,required"`
	SequenceNumber int64                                                    `json:"sequence_number,required"`
	SummaryIndex   int64                                                    `json:"summary_index,required"`
	// Any of "response.reasoning_summary_part.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Part           respjson.Field
		SequenceNumber respjson.Field
		SummaryIndex   respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseReasoningSummaryPartDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningSummaryPartDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning summary part in a streamed response.
type ResponseObjectStreamResponseReasoningSummaryPartDonePart struct {
	Text string `json:"text,required"`
	// Any of "summary_text".
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
func (r ResponseObjectStreamResponseReasoningSummaryPartDonePart) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningSummaryPartDonePart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for incremental reasoning summary text updates.
type ResponseObjectStreamResponseReasoningSummaryTextDelta struct {
	Delta          string `json:"delta,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	SummaryIndex   int64  `json:"summary_index,required"`
	// Any of "response.reasoning_summary_text.delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		SummaryIndex   respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseReasoningSummaryTextDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningSummaryTextDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when reasoning summary text is completed.
type ResponseObjectStreamResponseReasoningSummaryTextDone struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	SummaryIndex   int64  `json:"summary_index,required"`
	Text           string `json:"text,required"`
	// Any of "response.reasoning_summary_text.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		SummaryIndex   respjson.Field
		Text           respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseReasoningSummaryTextDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseReasoningSummaryTextDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for incremental refusal text updates.
type ResponseObjectStreamResponseRefusalDelta struct {
	ContentIndex   int64  `json:"content_index,required"`
	Delta          string `json:"delta,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.refusal.delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseRefusalDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseRefusalDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when refusal text is completed.
type ResponseObjectStreamResponseRefusalDone struct {
	ContentIndex   int64  `json:"content_index,required"`
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	Refusal        string `json:"refusal,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.refusal.done".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Refusal        respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseRefusalDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseRefusalDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when an annotation is added to output text.
type ResponseObjectStreamResponseOutputTextAnnotationAdded struct {
	// File citation annotation for referencing specific files in response content.
	Annotation      ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion `json:"annotation,required"`
	AnnotationIndex int64                                                                `json:"annotation_index,required"`
	ContentIndex    int64                                                                `json:"content_index,required"`
	ItemID          string                                                               `json:"item_id,required"`
	OutputIndex     int64                                                                `json:"output_index,required"`
	SequenceNumber  int64                                                                `json:"sequence_number,required"`
	// Any of "response.output_text.annotation.added".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotation      respjson.Field
		AnnotationIndex respjson.Field
		ContentIndex    respjson.Field
		ItemID          respjson.Field
		OutputIndex     respjson.Field
		SequenceNumber  respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextAnnotationAdded) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputTextAnnotationAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion contains
// all possible properties and values from
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation],
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation],
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputTextAnnotationAddedAnnotation is
// implemented by each variant of
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion] to add
// type safety for the return type of
// [ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseOutputTextAnnotationAddedAnnotation interface {
	implResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion()
}

func (ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation) implResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation) implResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation) implResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath) implResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) AsAny() anyResponseObjectStreamResponseOutputTextAnnotationAddedAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputTextAnnotationAddedAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for file search calls in progress.
type ResponseObjectStreamResponseFileSearchCallInProgress struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.file_search_call.in_progress".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFileSearchCallInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFileSearchCallInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for file search currently searching.
type ResponseObjectStreamResponseFileSearchCallSearching struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.file_search_call.searching".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFileSearchCallSearching) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFileSearchCallSearching) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for completed file search calls.
type ResponseObjectStreamResponseFileSearchCallCompleted struct {
	ItemID         string `json:"item_id,required"`
	OutputIndex    int64  `json:"output_index,required"`
	SequenceNumber int64  `json:"sequence_number,required"`
	// Any of "response.file_search_call.completed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFileSearchCallCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFileSearchCallCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event emitted when a response ends in an incomplete state.
type ResponseObjectStreamResponseIncomplete struct {
	// Complete OpenAI response object containing generation results and metadata.
	Response       ResponseObject `json:"response,required"`
	SequenceNumber int64          `json:"sequence_number,required"`
	// Any of "response.incomplete".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response       respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseIncomplete) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseIncomplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event emitted when a response fails.
type ResponseObjectStreamResponseFailed struct {
	// Complete OpenAI response object containing generation results and metadata.
	Response       ResponseObject `json:"response,required"`
	SequenceNumber int64          `json:"sequence_number,required"`
	// Any of "response.failed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response       respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFailed) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating a response has been completed.
type ResponseObjectStreamResponseCompleted struct {
	// Complete OpenAI response object containing generation results and metadata.
	Response ResponseObject `json:"response,required"`
	// Any of "response.completed".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI response object extended with input context information.
type ResponseListResponse struct {
	ID        string                            `json:"id,required"`
	CreatedAt int64                             `json:"created_at,required"`
	Input     []ResponseListResponseInputUnion  `json:"input,required"`
	Model     string                            `json:"model,required"`
	Output    []ResponseListResponseOutputUnion `json:"output,required"`
	Status    string                            `json:"status,required"`
	// Error details for failed OpenAI response requests.
	Error        ResponseListResponseError `json:"error,nullable"`
	Instructions string                    `json:"instructions,nullable"`
	MaxToolCalls int64                     `json:"max_tool_calls,nullable"`
	Metadata     map[string]string         `json:"metadata,nullable"`
	// Any of "response".
	Object             ResponseListResponseObject `json:"object"`
	ParallelToolCalls  bool                       `json:"parallel_tool_calls,nullable"`
	PreviousResponseID string                     `json:"previous_response_id,nullable"`
	// OpenAI compatible Prompt object that is used in OpenAI responses.
	Prompt      ResponseListResponsePrompt `json:"prompt,nullable"`
	Temperature float64                    `json:"temperature,nullable"`
	// Text response configuration for OpenAI responses.
	Text ResponseListResponseText `json:"text"`
	// Constrains the tools available to the model to a pre-defined set.
	ToolChoice ResponseListResponseToolChoiceUnion `json:"tool_choice,nullable"`
	Tools      []ResponseListResponseToolUnion     `json:"tools,nullable"`
	TopP       float64                             `json:"top_p,nullable"`
	Truncation string                              `json:"truncation,nullable"`
	// Usage information for OpenAI response.
	Usage ResponseListResponseUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Input              respjson.Field
		Model              respjson.Field
		Output             respjson.Field
		Status             respjson.Field
		Error              respjson.Field
		Instructions       respjson.Field
		MaxToolCalls       respjson.Field
		Metadata           respjson.Field
		Object             respjson.Field
		ParallelToolCalls  respjson.Field
		PreviousResponseID respjson.Field
		Prompt             respjson.Field
		Temperature        respjson.Field
		Text               respjson.Field
		ToolChoice         respjson.Field
		Tools              respjson.Field
		TopP               respjson.Field
		Truncation         respjson.Field
		Usage              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputUnion contains all possible properties and values from
// [ResponseListResponseInputOpenAIResponseMessageOutput],
// [ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseListResponseInputOpenAIResponseOutputMessageMcpCall],
// [ResponseListResponseInputOpenAIResponseOutputMessageMcpListTools],
// [ResponseListResponseInputOpenAIResponseMcpApprovalRequest],
// [ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput],
// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse],
// [ResponseListResponseInputOpenAIResponseMessageOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputUnion struct {
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutput].
	Content ResponseListResponseInputOpenAIResponseMessageOutputContentUnion `json:"content"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutput].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall].
	Results     []ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	Arguments   string                                                                         `json:"arguments"`
	CallID      string                                                                         `json:"call_id"`
	Name        string                                                                         `json:"name"`
	ServerLabel string                                                                         `json:"server_label"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseOutputMessageMcpCall].
	Error  string `json:"error"`
	Output string `json:"output"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseOutputMessageMcpListTools].
	Tools []ResponseListResponseInputOpenAIResponseOutputMessageMcpListToolsTool `json:"tools"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	JSON   struct {
		Content           respjson.Field
		Role              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		Queries           respjson.Field
		Results           respjson.Field
		Arguments         respjson.Field
		CallID            respjson.Field
		Name              respjson.Field
		ServerLabel       respjson.Field
		Error             respjson.Field
		Output            respjson.Field
		Tools             respjson.Field
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Reason            respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseMessageOutput() (v ResponseListResponseInputOpenAIResponseMessageOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageMcpCall() (v ResponseListResponseInputOpenAIResponseOutputMessageMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageMcpListTools() (v ResponseListResponseInputOpenAIResponseOutputMessageMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseMcpApprovalRequest() (v ResponseListResponseInputOpenAIResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseMcpApprovalResponse() (v ResponseListResponseInputOpenAIResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsResponseListResponseInputOpenAIResponseMessageOutput() (v ResponseListResponseInputOpenAIResponseMessageOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseListResponseInputOpenAIResponseMessageOutput struct {
	Content ResponseListResponseInputOpenAIResponseMessageOutputContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string `json:"role,required"`
	ID     string `json:"id,nullable"`
	Status string `json:"status,nullable"`
	// Any of "message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutput) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseInputOpenAIResponseMessageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageOutputContentUnion contains all
// possible properties and values from [string],
// [[]ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ResponseListResponseInputOpenAIResponseMessageOutputContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                                           `json:"text,required"`
	Annotations []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                                             `json:"token,required"`
	Logprob     float64                                                                                                                                                                            `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                                            `json:"bytes,nullable"`
	TopLogprobs []ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall struct {
	ID      string                                                                         `json:"id,required"`
	Queries []string                                                                       `json:"queries,required"`
	Status  string                                                                         `json:"status,required"`
	Results []ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results,nullable"`
	// Any of "file_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Results     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult struct {
	Attributes map[string]any `json:"attributes,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall struct {
	Arguments string `json:"arguments,required"`
	CallID    string `json:"call_id,required"`
	Name      string `json:"name,required"`
	ID        string `json:"id,nullable"`
	Status    string `json:"status,nullable"`
	// Any of "function_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageMcpCall struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	Error       string `json:"error,nullable"`
	Output      string `json:"output,nullable"`
	// Any of "mcp_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageMcpCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseListResponseInputOpenAIResponseOutputMessageMcpListTools struct {
	ID          string                                                                 `json:"id,required"`
	ServerLabel string                                                                 `json:"server_label,required"`
	Tools       []ResponseListResponseInputOpenAIResponseOutputMessageMcpListToolsTool `json:"tools,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseListResponseInputOpenAIResponseOutputMessageMcpListToolsTool struct {
	InputSchema map[string]any `json:"input_schema,required"`
	Name        string         `json:"name,required"`
	Description string         `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseOutputMessageMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseListResponseInputOpenAIResponseMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput struct {
	CallID string `json:"call_id,required"`
	Output string `json:"output,required"`
	ID     string `json:"id,nullable"`
	Status string `json:"status,nullable"`
	// Any of "function_call_output".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Output      respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ResponseListResponseInputOpenAIResponseMcpApprovalResponse struct {
	ApprovalRequestID string `json:"approval_request_id,required"`
	Approve           bool   `json:"approve,required"`
	ID                string `json:"id,nullable"`
	Reason            string `json:"reason,nullable"`
	// Any of "mcp_approval_response".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		ID                respjson.Field
		Reason            respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseInputOpenAIResponseMcpApprovalResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputUnion contains all possible properties and values from
// [ResponseListResponseOutputMessage], [ResponseListResponseOutputWebSearchCall],
// [ResponseListResponseOutputFileSearchCall],
// [ResponseListResponseOutputFunctionCall], [ResponseListResponseOutputMcpCall],
// [ResponseListResponseOutputMcpListTools],
// [ResponseListResponseOutputMcpApprovalRequest].
//
// Use the [ResponseListResponseOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputUnion struct {
	// This field is from variant [ResponseListResponseOutputMessage].
	Content ResponseListResponseOutputMessageContentUnion `json:"content"`
	// This field is from variant [ResponseListResponseOutputMessage].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type string `json:"type"`
	// This field is from variant [ResponseListResponseOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseListResponseOutputFileSearchCall].
	Results   []ResponseListResponseOutputFileSearchCallResult `json:"results"`
	Arguments string                                           `json:"arguments"`
	// This field is from variant [ResponseListResponseOutputFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseListResponseOutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [ResponseListResponseOutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [ResponseListResponseOutputMcpListTools].
	Tools []ResponseListResponseOutputMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseListResponseOutput is implemented by each variant of
// [ResponseListResponseOutputUnion] to add type safety for the return type of
// [ResponseListResponseOutputUnion.AsAny]
type anyResponseListResponseOutput interface {
	implResponseListResponseOutputUnion()
}

func (ResponseListResponseOutputMessage) implResponseListResponseOutputUnion()            {}
func (ResponseListResponseOutputWebSearchCall) implResponseListResponseOutputUnion()      {}
func (ResponseListResponseOutputFileSearchCall) implResponseListResponseOutputUnion()     {}
func (ResponseListResponseOutputFunctionCall) implResponseListResponseOutputUnion()       {}
func (ResponseListResponseOutputMcpCall) implResponseListResponseOutputUnion()            {}
func (ResponseListResponseOutputMcpListTools) implResponseListResponseOutputUnion()       {}
func (ResponseListResponseOutputMcpApprovalRequest) implResponseListResponseOutputUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessage:
//	case llamastackclient.ResponseListResponseOutputWebSearchCall:
//	case llamastackclient.ResponseListResponseOutputFileSearchCall:
//	case llamastackclient.ResponseListResponseOutputFunctionCall:
//	case llamastackclient.ResponseListResponseOutputMcpCall:
//	case llamastackclient.ResponseListResponseOutputMcpListTools:
//	case llamastackclient.ResponseListResponseOutputMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputUnion) AsAny() anyResponseListResponseOutput {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseListResponseOutputUnion) AsMessage() (v ResponseListResponseOutputMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsWebSearchCall() (v ResponseListResponseOutputWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsFileSearchCall() (v ResponseListResponseOutputFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsFunctionCall() (v ResponseListResponseOutputFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsMcpCall() (v ResponseListResponseOutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsMcpListTools() (v ResponseListResponseOutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsMcpApprovalRequest() (v ResponseListResponseOutputMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseListResponseOutputMessage struct {
	Content ResponseListResponseOutputMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string `json:"role,required"`
	ID     string `json:"id,nullable"`
	Status string `json:"status,nullable"`
	// Any of "message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentUnion contains all possible properties
// and values from [string],
// [[]ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ResponseListResponseOutputMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseListResponseOutputMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseOutputMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                        `json:"text,required"`
	Annotations []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Logprobs    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                          `json:"token,required"`
	Logprob     float64                                                                                                                                                         `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                         `json:"bytes,nullable"`
	TopLogprobs []ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseListResponseOutputWebSearchCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseListResponseOutputFileSearchCall struct {
	ID      string                                           `json:"id,required"`
	Queries []string                                         `json:"queries,required"`
	Status  string                                           `json:"status,required"`
	Results []ResponseListResponseOutputFileSearchCallResult `json:"results,nullable"`
	// Any of "file_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Results     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseListResponseOutputFileSearchCallResult struct {
	Attributes map[string]any `json:"attributes,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseListResponseOutputFunctionCall struct {
	Arguments string `json:"arguments,required"`
	CallID    string `json:"call_id,required"`
	Name      string `json:"name,required"`
	ID        string `json:"id,nullable"`
	Status    string `json:"status,nullable"`
	// Any of "function_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseListResponseOutputMcpCall struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	Error       string `json:"error,nullable"`
	Output      string `json:"output,nullable"`
	// Any of "mcp_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseListResponseOutputMcpListTools struct {
	ID          string                                       `json:"id,required"`
	ServerLabel string                                       `json:"server_label,required"`
	Tools       []ResponseListResponseOutputMcpListToolsTool `json:"tools,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseListResponseOutputMcpListToolsTool struct {
	InputSchema map[string]any `json:"input_schema,required"`
	Name        string         `json:"name,required"`
	Description string         `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseListResponseOutputMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseOutputMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for failed OpenAI response requests.
type ResponseListResponseError struct {
	Code    string `json:"code,required"`
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseObject string

const (
	ResponseListResponseObjectResponse ResponseListResponseObject = "response"
)

// OpenAI compatible Prompt object that is used in OpenAI responses.
type ResponseListResponsePrompt struct {
	ID        string                                             `json:"id,required"`
	Variables map[string]ResponseListResponsePromptVariableUnion `json:"variables,nullable"`
	Version   string                                             `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Variables   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponsePrompt) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponsePrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponsePromptVariableUnion contains all possible properties and
// values from [ResponseListResponsePromptVariableInputText],
// [ResponseListResponsePromptVariableInputImage],
// [ResponseListResponsePromptVariableInputFile].
//
// Use the [ResponseListResponsePromptVariableUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponsePromptVariableUnion struct {
	// This field is from variant [ResponseListResponsePromptVariableInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant [ResponseListResponsePromptVariableInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant [ResponseListResponsePromptVariableInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant [ResponseListResponsePromptVariableInputFile].
	FileData string `json:"file_data"`
	// This field is from variant [ResponseListResponsePromptVariableInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant [ResponseListResponsePromptVariableInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponsePromptVariable is implemented by each variant of
// [ResponseListResponsePromptVariableUnion] to add type safety for the return type
// of [ResponseListResponsePromptVariableUnion.AsAny]
type anyResponseListResponsePromptVariable interface {
	implResponseListResponsePromptVariableUnion()
}

func (ResponseListResponsePromptVariableInputText) implResponseListResponsePromptVariableUnion()  {}
func (ResponseListResponsePromptVariableInputImage) implResponseListResponsePromptVariableUnion() {}
func (ResponseListResponsePromptVariableInputFile) implResponseListResponsePromptVariableUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponsePromptVariableUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponsePromptVariableInputText:
//	case llamastackclient.ResponseListResponsePromptVariableInputImage:
//	case llamastackclient.ResponseListResponsePromptVariableInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponsePromptVariableUnion) AsAny() anyResponseListResponsePromptVariable {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
	}
	return nil
}

func (u ResponseListResponsePromptVariableUnion) AsInputText() (v ResponseListResponsePromptVariableInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponsePromptVariableUnion) AsInputImage() (v ResponseListResponsePromptVariableInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponsePromptVariableUnion) AsInputFile() (v ResponseListResponsePromptVariableInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponsePromptVariableUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponsePromptVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseListResponsePromptVariableInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
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
func (r ResponseListResponsePromptVariableInputText) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponsePromptVariableInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseListResponsePromptVariableInputImage struct {
	// Any of "low", "high", "auto".
	Detail   string `json:"detail"`
	FileID   string `json:"file_id,nullable"`
	ImageURL string `json:"image_url,nullable"`
	// Any of "input_image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponsePromptVariableInputImage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponsePromptVariableInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseListResponsePromptVariableInputFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	FileURL  string `json:"file_url,nullable"`
	Filename string `json:"filename,nullable"`
	// Any of "input_file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponsePromptVariableInputFile) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponsePromptVariableInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text response configuration for OpenAI responses.
type ResponseListResponseText struct {
	// Configuration for Responses API text format.
	Format ResponseListResponseTextFormat `json:"format,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseText) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Responses API text format.
type ResponseListResponseTextFormat struct {
	Description string         `json:"description,nullable"`
	Name        string         `json:"name,nullable"`
	Schema      map[string]any `json:"schema,nullable"`
	Strict      bool           `json:"strict,nullable"`
	// Any of "text", "json_schema", "json_object".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Schema      respjson.Field
		Strict      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolChoiceUnion contains all possible properties and values
// from [string],
// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools],
// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFileSearch],
// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceWebSearch],
// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFunctionTool],
// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMcpTool],
// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceCustomTool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOpenAIResponseInputToolChoiceMode]
type ResponseListResponseToolChoiceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfOpenAIResponseInputToolChoiceMode string `json:",inline"`
	// This field is from variant
	// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools].
	Tools []map[string]string `json:"tools"`
	// This field is from variant
	// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools].
	Mode string `json:"mode"`
	Type string `json:"type"`
	Name string `json:"name"`
	// This field is from variant
	// [ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMcpTool].
	ServerLabel string `json:"server_label"`
	JSON        struct {
		OfOpenAIResponseInputToolChoiceMode respjson.Field
		Tools                               respjson.Field
		Mode                                respjson.Field
		Type                                respjson.Field
		Name                                respjson.Field
		ServerLabel                         respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceMode() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceAllowedTools() (v ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceFileSearch() (v ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFileSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceWebSearch() (v ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceWebSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceFunctionTool() (v ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFunctionTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceMcpTool() (v ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMcpTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolChoiceUnion) AsOpenAIResponseInputToolChoiceCustomTool() (v ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceCustomTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolChoiceUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMode string

const (
	ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceModeAuto     ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMode = "auto"
	ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceModeRequired ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMode = "required"
	ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceModeNone     ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMode = "none"
)

// Constrains the tools available to the model to a pre-defined set.
type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools struct {
	Tools []map[string]string `json:"tools,required"`
	// Any of "auto", "required".
	Mode string `json:"mode"`
	// Any of "allowed_tools".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tools       respjson.Field
		Mode        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceAllowedTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates that the model should use file search to generate a response.
type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFileSearch struct {
	// Any of "file_search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFileSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates that the model should use web search to generate a response
type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceWebSearch struct {
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11",
	// "web_search_2025_08_26".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceWebSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forces the model to call a specific function.
type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFunctionTool struct {
	Name string `json:"name,required"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFunctionTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceFunctionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forces the model to call a specific tool on a remote MCP server
type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMcpTool struct {
	ServerLabel string `json:"server_label,required"`
	Name        string `json:"name,nullable"`
	// Any of "mcp".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServerLabel respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMcpTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceMcpTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forces the model to call a custom tool.
type ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceCustomTool struct {
	Name string `json:"name,required"`
	// Any of "custom".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceCustomTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolChoiceOpenAIResponseInputToolChoiceCustomTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolUnion contains all possible properties and values from
// [ResponseListResponseToolOpenAIResponseInputToolWebSearch],
// [ResponseListResponseToolFileSearch], [ResponseListResponseToolFunction],
// [ResponseListResponseToolMcp].
//
// Use the [ResponseListResponseToolUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseToolUnion struct {
	// This field is from variant
	// [ResponseListResponseToolOpenAIResponseInputToolWebSearch].
	SearchContextSize string `json:"search_context_size"`
	// Any of nil, "file_search", "function", "mcp".
	Type string `json:"type"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	VectorStoreIDs []string `json:"vector_store_ids"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	Filters map[string]any `json:"filters"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	MaxNumResults int64 `json:"max_num_results"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	RankingOptions ResponseListResponseToolFileSearchRankingOptions `json:"ranking_options"`
	// This field is from variant [ResponseListResponseToolFunction].
	Name string `json:"name"`
	// This field is from variant [ResponseListResponseToolFunction].
	Parameters map[string]any `json:"parameters"`
	// This field is from variant [ResponseListResponseToolFunction].
	Description string `json:"description"`
	// This field is from variant [ResponseListResponseToolFunction].
	Strict bool `json:"strict"`
	// This field is from variant [ResponseListResponseToolMcp].
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseListResponseToolMcp].
	AllowedTools ResponseListResponseToolMcpAllowedToolsUnion `json:"allowed_tools"`
	JSON         struct {
		SearchContextSize respjson.Field
		Type              respjson.Field
		VectorStoreIDs    respjson.Field
		Filters           respjson.Field
		MaxNumResults     respjson.Field
		RankingOptions    respjson.Field
		Name              respjson.Field
		Parameters        respjson.Field
		Description       respjson.Field
		Strict            respjson.Field
		ServerLabel       respjson.Field
		AllowedTools      respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseListResponseToolUnion) AsOpenAIResponseInputToolWebSearch() (v ResponseListResponseToolOpenAIResponseInputToolWebSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolUnion) AsFileSearch() (v ResponseListResponseToolFileSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolUnion) AsFunction() (v ResponseListResponseToolFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolUnion) AsMcp() (v ResponseListResponseToolMcp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool configuration for OpenAI response inputs.
type ResponseListResponseToolOpenAIResponseInputToolWebSearch struct {
	SearchContextSize string `json:"search_context_size,nullable"`
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11",
	// "web_search_2025_08_26".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SearchContextSize respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolOpenAIResponseInputToolWebSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool configuration for OpenAI response inputs.
type ResponseListResponseToolFileSearch struct {
	VectorStoreIDs []string       `json:"vector_store_ids,required"`
	Filters        map[string]any `json:"filters,nullable"`
	MaxNumResults  int64          `json:"max_num_results,nullable"`
	// Options for ranking and filtering search results.
	RankingOptions ResponseListResponseToolFileSearchRankingOptions `json:"ranking_options,nullable"`
	// Any of "file_search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VectorStoreIDs respjson.Field
		Filters        respjson.Field
		MaxNumResults  respjson.Field
		RankingOptions respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolFileSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for ranking and filtering search results.
type ResponseListResponseToolFileSearchRankingOptions struct {
	Ranker         string  `json:"ranker,nullable"`
	ScoreThreshold float64 `json:"score_threshold,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ranker         respjson.Field
		ScoreThreshold respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolFileSearchRankingOptions) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool configuration for OpenAI response inputs.
type ResponseListResponseToolFunction struct {
	Name        string         `json:"name,required"`
	Parameters  map[string]any `json:"parameters,required"`
	Description string         `json:"description,nullable"`
	Strict      bool           `json:"strict,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Parameters  respjson.Field
		Description respjson.Field
		Strict      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolFunction) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) tool configuration for OpenAI response object.
type ResponseListResponseToolMcp struct {
	ServerLabel string `json:"server_label,required"`
	// Filter configuration for restricting which MCP tools can be used.
	AllowedTools ResponseListResponseToolMcpAllowedToolsUnion `json:"allowed_tools,nullable"`
	// Any of "mcp".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServerLabel  respjson.Field
		AllowedTools respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolMcp) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolMcpAllowedToolsUnion contains all possible properties
// and values from [[]string],
// [ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfListString]
type ResponseListResponseToolMcpAllowedToolsUnion struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfListString []string `json:",inline"`
	// This field is from variant
	// [ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter].
	ToolNames []string `json:"tool_names"`
	JSON      struct {
		OfListString respjson.Field
		ToolNames    respjson.Field
		raw          string
	} `json:"-"`
}

func (u ResponseListResponseToolMcpAllowedToolsUnion) AsListString() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolMcpAllowedToolsUnion) AsAllowedToolsFilter() (v ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolMcpAllowedToolsUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter configuration for restricting which MCP tools can be used.
type ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter struct {
	ToolNames []string `json:"tool_names,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolNames   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage information for OpenAI response.
type ResponseListResponseUsage struct {
	InputTokens  int64 `json:"input_tokens,required"`
	OutputTokens int64 `json:"output_tokens,required"`
	TotalTokens  int64 `json:"total_tokens,required"`
	// Token details for input tokens in OpenAI response usage.
	InputTokensDetails ResponseListResponseUsageInputTokensDetails `json:"input_tokens_details,nullable"`
	// Token details for output tokens in OpenAI response usage.
	OutputTokensDetails ResponseListResponseUsageOutputTokensDetails `json:"output_tokens_details,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		OutputTokens        respjson.Field
		TotalTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokensDetails respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for input tokens in OpenAI response usage.
type ResponseListResponseUsageInputTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for output tokens in OpenAI response usage.
type ResponseListResponseUsageOutputTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object confirming deletion of an OpenAI response.
type ResponseDeleteResponse struct {
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted"`
	// Any of "response".
	Object ResponseDeleteResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseDeleteResponseObject string

const (
	ResponseDeleteResponseObjectResponse ResponseDeleteResponseObject = "response"
)

type ResponseNewParams struct {
	Input              ResponseNewParamsInputUnion `json:"input,omitzero,required"`
	Model              string                      `json:"model,required"`
	Conversation       param.Opt[string]           `json:"conversation,omitzero"`
	Instructions       param.Opt[string]           `json:"instructions,omitzero"`
	MaxInferIters      param.Opt[int64]            `json:"max_infer_iters,omitzero"`
	MaxToolCalls       param.Opt[int64]            `json:"max_tool_calls,omitzero"`
	ParallelToolCalls  param.Opt[bool]             `json:"parallel_tool_calls,omitzero"`
	PreviousResponseID param.Opt[string]           `json:"previous_response_id,omitzero"`
	Store              param.Opt[bool]             `json:"store,omitzero"`
	Temperature        param.Opt[float64]          `json:"temperature,omitzero"`
	// Any of "web_search_call.action.sources", "code_interpreter_call.outputs",
	// "computer_call_output.output.image_url", "file_search_call.results",
	// "message.input_image.image_url", "message.output_text.logprobs",
	// "reasoning.encrypted_content".
	Include  []string          `json:"include,omitzero"`
	Metadata map[string]string `json:"metadata,omitzero"`
	// OpenAI compatible Prompt object that is used in OpenAI responses.
	Prompt ResponseNewParamsPrompt `json:"prompt,omitzero"`
	// Text response configuration for OpenAI responses.
	Text ResponseNewParamsText `json:"text,omitzero"`
	// Constrains the tools available to the model to a pre-defined set.
	ToolChoice ResponseNewParamsToolChoiceUnion `json:"tool_choice,omitzero"`
	Tools      []ResponseNewParamsToolUnion     `json:"tools,omitzero"`
	paramObj
}

func (r ResponseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputUnion struct {
	OfString                                                                  param.Opt[string]                                                                                        `json:",omitzero,inline"`
	OfListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutput []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutput)
}
func (u *ResponseNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutput) {
		return &u.OfListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutput
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion struct {
	OfOpenAIResponseMessageInput                                                                                             *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput                    `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageWebSearchToolCall                                                                           *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageWebSearchToolCall  `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageFileSearchToolCall                                                                          *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCall `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageFunctionToolCall                                                                            *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFunctionToolCall   `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageMcpCall                                                                                     *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpCall            `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageMcpListTools                                                                                *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListTools       `json:",omitzero,inline"`
	OfOpenAIResponseMcpApprovalRequest                                                                                       *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalRequest              `json:",omitzero,inline"`
	OfOpenAIResponseInputFunctionToolCallOutput                                                                              *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseInputFunctionToolCallOutput     `json:",omitzero,inline"`
	OfOpenAIResponseMcpApprovalResponse                                                                                      *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalResponse             `json:",omitzero,inline"`
	OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput                    `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseMessageInput,
		u.OfOpenAIResponseOutputMessageWebSearchToolCall,
		u.OfOpenAIResponseOutputMessageFileSearchToolCall,
		u.OfOpenAIResponseOutputMessageFunctionToolCall,
		u.OfOpenAIResponseOutputMessageMcpCall,
		u.OfOpenAIResponseOutputMessageMcpListTools,
		u.OfOpenAIResponseMcpApprovalRequest,
		u.OfOpenAIResponseInputFunctionToolCallOutput,
		u.OfOpenAIResponseMcpApprovalResponse,
		u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput)
}
func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseMessageInput) {
		return u.OfOpenAIResponseMessageInput
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageWebSearchToolCall) {
		return u.OfOpenAIResponseOutputMessageWebSearchToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageFileSearchToolCall) {
		return u.OfOpenAIResponseOutputMessageFileSearchToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageFunctionToolCall) {
		return u.OfOpenAIResponseOutputMessageFunctionToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageMcpCall) {
		return u.OfOpenAIResponseOutputMessageMcpCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageMcpListTools) {
		return u.OfOpenAIResponseOutputMessageMcpListTools
	} else if !param.IsOmitted(u.OfOpenAIResponseMcpApprovalRequest) {
		return u.OfOpenAIResponseMcpApprovalRequest
	} else if !param.IsOmitted(u.OfOpenAIResponseInputFunctionToolCallOutput) {
		return u.OfOpenAIResponseInputFunctionToolCallOutput
	} else if !param.IsOmitted(u.OfOpenAIResponseMcpApprovalResponse) {
		return u.OfOpenAIResponseMcpApprovalResponse
	} else if !param.IsOmitted(u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput) {
		return u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetQueries() []string {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetResults() []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCallResult {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetError() *string {
	if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil && vt.Error.Valid() {
		return &vt.Error.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetTools() []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListToolsTool {
	if vt := u.OfOpenAIResponseOutputMessageMcpListTools; vt != nil {
		return vt.Tools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetApprovalRequestID() *string {
	if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil {
		return &vt.ApprovalRequestID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetApprove() *bool {
	if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil {
		return &vt.Approve
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetReason() *string {
	if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetRole() *string {
	if vt := u.OfOpenAIResponseMessageInput; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetID() *string {
	if vt := u.OfOpenAIResponseMessageInput; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageMcpListTools; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetStatus() *string {
	if vt := u.OfOpenAIResponseMessageInput; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetType() *string {
	if vt := u.OfOpenAIResponseMessageInput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageMcpListTools; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetArguments() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Arguments)
	} else if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil {
		return (*string)(&vt.Arguments)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Arguments)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetCallID() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.CallID)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.CallID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetName() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetServerLabel() *string {
	if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil {
		return (*string)(&vt.ServerLabel)
	} else if vt := u.OfOpenAIResponseOutputMessageMcpListTools; vt != nil {
		return (*string)(&vt.ServerLabel)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.ServerLabel)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetOutput() *string {
	if vt := u.OfOpenAIResponseOutputMessageMcpCall; vt != nil && vt.Output.Valid() {
		return &vt.Output.Value
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.Output)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemUnion) GetContent() *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentUnion {
	if vt := u.OfOpenAIResponseMessageInput; vt != nil {
		return &vt.Content
	} else if vt := u.OfResponseNewsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput; vt != nil {
		return &vt.Content
	}
	return nil
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput struct {
	Content ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentUnion `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string            `json:"role,omitzero,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// Any of "message".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput](
		"role", "system", "developer", "user", "assistant",
	)
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInput](
		"type", "message",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentUnion struct {
	OfString                                                                                                               param.Opt[string]                                                                                                                                                                                                                                                 `json:",omitzero,inline"`
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",omitzero,inline"`
	OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal                                []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion                                `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile, u.OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal)
}
func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile) {
		return &u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
	} else if !param.IsOmitted(u.OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal) {
		return &u.OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	OfInputText  *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText  `json:",omitzero,inline"`
	OfInputImage *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage `json:",omitzero,inline"`
	OfInputFile  *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile  `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage, u.OfInputFile)
}
func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) asAny() any {
	if !param.IsOmitted(u.OfInputText) {
		return u.OfInputText
	} else if !param.IsOmitted(u.OfInputImage) {
		return u.OfInputImage
	} else if !param.IsOmitted(u.OfInputFile) {
		return u.OfInputFile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return &vt.Detail
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileData() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileData.Valid() {
		return &vt.FileData.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileURL() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileURL.Valid() {
		return &vt.FileURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFilename() *string {
	if vt := u.OfInputFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetType() *string {
	if vt := u.OfInputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputFile; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileID() *string {
	if vt := u.OfInputImage; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	} else if vt := u.OfInputFile; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText]("input_text"),
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage]("input_image"),
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile]("input_file"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The property Text is required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText](
		"type", "input_text",
	)
}

// Image content for input messages in OpenAI response format.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	// Any of "input_image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"detail", "low", "high", "auto",
	)
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"type", "input_image",
	)
}

// File content for input messages in OpenAI response format.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	FileURL  param.Opt[string] `json:"file_url,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Any of "input_file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile](
		"type", "input_file",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion struct {
	OfOutputText *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText `json:",omitzero,inline"`
	OfRefusal    *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal    `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOutputText, u.OfRefusal)
}
func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOutputText) {
		return u.OfOutputText
	} else if !param.IsOmitted(u.OfRefusal) {
		return u.OfRefusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetText() *string {
	if vt := u.OfOutputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetAnnotations() []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion {
	if vt := u.OfOutputText; vt != nil {
		return vt.Annotations
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetLogprobs() []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob {
	if vt := u.OfOutputText; vt != nil {
		return vt.Logprobs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetRefusal() *string {
	if vt := u.OfRefusal; vt != nil {
		return &vt.Refusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetType() *string {
	if vt := u.OfOutputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRefusal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText]("output_text"),
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal]("refusal"),
	)
}

// The property Text is required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                                                                                                                 `json:"text,required"`
	Logprobs    []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,omitzero"`
	Annotations []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations,omitzero"`
	// Any of "output_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText](
		"type", "output_text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	OfFileCitation          *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation          `json:",omitzero,inline"`
	OfURLCitation           *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation           `json:",omitzero,inline"`
	OfContainerFileCitation *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation `json:",omitzero,inline"`
	OfFilePath              *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) asAny() any {
	if !param.IsOmitted(u.OfFileCitation) {
		return u.OfFileCitation
	} else if !param.IsOmitted(u.OfURLCitation) {
		return u.OfURLCitation
	} else if !param.IsOmitted(u.OfContainerFileCitation) {
		return u.OfContainerFileCitation
	} else if !param.IsOmitted(u.OfFilePath) {
		return u.OfFilePath
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFileID() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.FileID)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.FileID)
	} else if vt := u.OfFilePath; vt != nil {
		return (*string)(&vt.FileID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetType() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfURLCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFilePath; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation]("file_citation"),
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation]("url_citation"),
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation]("container_file_citation"),
		apijson.Discriminator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath]("file_path"),
	)
}

// File citation annotation for referencing specific files in response content.
//
// The properties FileID, Filename, Index are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation](
		"type", "file_citation",
	)
}

// URL citation annotation for referencing external web resources.
//
// The properties EndIndex, StartIndex, Title, URL are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation](
		"type", "url_citation",
	)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation](
		"type", "container_file_citation",
	)
}

// The properties FileID, Index are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath](
		"type", "file_path",
	)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
//
// The properties Token, Logprob are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                                                                                                                   `json:"token,required"`
	Logprob     float64                                                                                                                                                                                                                                                  `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                                                                                                                  `json:"bytes,omitzero"`
	TopLogprobs []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
//
// The properties Token, Logprob are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
//
// The property Refusal is required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMessageInputContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal](
		"type", "refusal",
	)
}

// Web search tool call output message for OpenAI responses.
//
// The properties ID, Status are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageWebSearchToolCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageWebSearchToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageWebSearchToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageWebSearchToolCall](
		"type", "web_search_call",
	)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCall struct {
	ID      string                                                                                                                                                 `json:"id,required"`
	Queries []string                                                                                                                                               `json:"queries,omitzero,required"`
	Status  string                                                                                                                                                 `json:"status,required"`
	Results []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results,omitzero"`
	// Any of "file_search_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCall](
		"type", "file_search_call",
	)
}

// Search results returned by the file search operation.
//
// The properties Attributes, FileID, Filename, Score, Text are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCallResult struct {
	Attributes map[string]any `json:"attributes,omitzero,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCallResult) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCallResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
//
// The properties Arguments, CallID, Name are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFunctionToolCall struct {
	Arguments string            `json:"arguments,required"`
	CallID    string            `json:"call_id,required"`
	Name      string            `json:"name,required"`
	ID        param.Opt[string] `json:"id,omitzero"`
	Status    param.Opt[string] `json:"status,omitzero"`
	// Any of "function_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFunctionToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFunctionToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageFunctionToolCall](
		"type", "function_call",
	)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpCall struct {
	ID          string            `json:"id,required"`
	Arguments   string            `json:"arguments,required"`
	Name        string            `json:"name,required"`
	ServerLabel string            `json:"server_label,required"`
	Error       param.Opt[string] `json:"error,omitzero"`
	Output      param.Opt[string] `json:"output,omitzero"`
	// Any of "mcp_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpCall](
		"type", "mcp_call",
	)
}

// MCP list tools output message containing available tools from an MCP server.
//
// The properties ID, ServerLabel, Tools are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListTools struct {
	ID          string                                                                                                                                         `json:"id,required"`
	ServerLabel string                                                                                                                                         `json:"server_label,required"`
	Tools       []ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListToolsTool `json:"tools,omitzero,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListTools) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListTools
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListTools](
		"type", "mcp_list_tools",
	)
}

// Tool definition returned by MCP list tools operation.
//
// The properties InputSchema, Name are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListToolsTool struct {
	InputSchema map[string]any    `json:"input_schema,omitzero,required"`
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListToolsTool) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListToolsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseOutputMessageMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalRequest) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalRequest](
		"type", "mcp_approval_request",
	)
}

// This represents the output of a function call that gets passed back to the
// model.
//
// The properties CallID, Output are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseInputFunctionToolCallOutput struct {
	CallID string            `json:"call_id,required"`
	Output string            `json:"output,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// Any of "function_call_output".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseInputFunctionToolCallOutput) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseInputFunctionToolCallOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseInputFunctionToolCallOutput](
		"type", "function_call_output",
	)
}

// A response to an MCP approval request.
//
// The properties ApprovalRequestID, Approve are required.
type ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalResponse struct {
	ApprovalRequestID string            `json:"approval_request_id,required"`
	Approve           bool              `json:"approve,required"`
	ID                param.Opt[string] `json:"id,omitzero"`
	Reason            param.Opt[string] `json:"reason,omitzero"`
	// Any of "mcp_approval_response".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalResponse) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsInputListOpenAIResponseMessageUnionOpenAIResponseInputFunctionToolCallOutputItemOpenAIResponseMcpApprovalResponse](
		"type", "mcp_approval_response",
	)
}

// OpenAI compatible Prompt object that is used in OpenAI responses.
//
// The property ID is required.
type ResponseNewParamsPrompt struct {
	ID        string                                          `json:"id,required"`
	Version   param.Opt[string]                               `json:"version,omitzero"`
	Variables map[string]ResponseNewParamsPromptVariableUnion `json:"variables,omitzero"`
	paramObj
}

func (r ResponseNewParamsPrompt) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsPrompt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsPromptVariableUnion struct {
	OfInputText  *ResponseNewParamsPromptVariableInputText  `json:",omitzero,inline"`
	OfInputImage *ResponseNewParamsPromptVariableInputImage `json:",omitzero,inline"`
	OfInputFile  *ResponseNewParamsPromptVariableInputFile  `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsPromptVariableUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage, u.OfInputFile)
}
func (u *ResponseNewParamsPromptVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsPromptVariableUnion) asAny() any {
	if !param.IsOmitted(u.OfInputText) {
		return u.OfInputText
	} else if !param.IsOmitted(u.OfInputImage) {
		return u.OfInputImage
	} else if !param.IsOmitted(u.OfInputFile) {
		return u.OfInputFile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return &vt.Detail
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetFileData() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileData.Valid() {
		return &vt.FileData.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetFileURL() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileURL.Valid() {
		return &vt.FileURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetFilename() *string {
	if vt := u.OfInputFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetType() *string {
	if vt := u.OfInputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputFile; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsPromptVariableUnion) GetFileID() *string {
	if vt := u.OfInputImage; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	} else if vt := u.OfInputFile; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsPromptVariableUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsPromptVariableInputText]("input_text"),
		apijson.Discriminator[ResponseNewParamsPromptVariableInputImage]("input_image"),
		apijson.Discriminator[ResponseNewParamsPromptVariableInputFile]("input_file"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The property Text is required.
type ResponseNewParamsPromptVariableInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsPromptVariableInputText) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsPromptVariableInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsPromptVariableInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsPromptVariableInputText](
		"type", "input_text",
	)
}

// Image content for input messages in OpenAI response format.
type ResponseNewParamsPromptVariableInputImage struct {
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	// Any of "input_image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsPromptVariableInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsPromptVariableInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsPromptVariableInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsPromptVariableInputImage](
		"detail", "low", "high", "auto",
	)
	apijson.RegisterFieldValidator[ResponseNewParamsPromptVariableInputImage](
		"type", "input_image",
	)
}

// File content for input messages in OpenAI response format.
type ResponseNewParamsPromptVariableInputFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	FileURL  param.Opt[string] `json:"file_url,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Any of "input_file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsPromptVariableInputFile) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsPromptVariableInputFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsPromptVariableInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsPromptVariableInputFile](
		"type", "input_file",
	)
}

// Text response configuration for OpenAI responses.
type ResponseNewParamsText struct {
	// Configuration for Responses API text format.
	Format ResponseNewParamsTextFormat `json:"format,omitzero"`
	paramObj
}

func (r ResponseNewParamsText) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Responses API text format.
type ResponseNewParamsTextFormat struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Strict      param.Opt[bool]   `json:"strict,omitzero"`
	Schema      map[string]any    `json:"schema,omitzero"`
	// Any of "text", "json_schema", "json_object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsTextFormat) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsTextFormat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsTextFormat](
		"type", "text", "json_schema", "json_object",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolChoiceUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfOpenAIResponseInputToolChoiceMode)
	OfOpenAIResponseInputToolChoiceMode         param.Opt[string]                                                     `json:",omitzero,inline"`
	OfOpenAIResponseInputToolChoiceAllowedTools *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools `json:",omitzero,inline"`
	OfOpenAIResponseInputToolChoiceFileSearch   *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFileSearch   `json:",omitzero,inline"`
	OfOpenAIResponseInputToolChoiceWebSearch    *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceWebSearch    `json:",omitzero,inline"`
	OfOpenAIResponseInputToolChoiceFunctionTool *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFunctionTool `json:",omitzero,inline"`
	OfOpenAIResponseInputToolChoiceMcpTool      *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMcpTool      `json:",omitzero,inline"`
	OfOpenAIResponseInputToolChoiceCustomTool   *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceCustomTool   `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseInputToolChoiceMode,
		u.OfOpenAIResponseInputToolChoiceAllowedTools,
		u.OfOpenAIResponseInputToolChoiceFileSearch,
		u.OfOpenAIResponseInputToolChoiceWebSearch,
		u.OfOpenAIResponseInputToolChoiceFunctionTool,
		u.OfOpenAIResponseInputToolChoiceMcpTool,
		u.OfOpenAIResponseInputToolChoiceCustomTool)
}
func (u *ResponseNewParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceMode) {
		return &u.OfOpenAIResponseInputToolChoiceMode
	} else if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceAllowedTools) {
		return u.OfOpenAIResponseInputToolChoiceAllowedTools
	} else if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceFileSearch) {
		return u.OfOpenAIResponseInputToolChoiceFileSearch
	} else if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceWebSearch) {
		return u.OfOpenAIResponseInputToolChoiceWebSearch
	} else if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceFunctionTool) {
		return u.OfOpenAIResponseInputToolChoiceFunctionTool
	} else if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceMcpTool) {
		return u.OfOpenAIResponseInputToolChoiceMcpTool
	} else if !param.IsOmitted(u.OfOpenAIResponseInputToolChoiceCustomTool) {
		return u.OfOpenAIResponseInputToolChoiceCustomTool
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolChoiceUnion) GetTools() []map[string]string {
	if vt := u.OfOpenAIResponseInputToolChoiceAllowedTools; vt != nil {
		return vt.Tools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolChoiceUnion) GetMode() *string {
	if vt := u.OfOpenAIResponseInputToolChoiceAllowedTools; vt != nil {
		return &vt.Mode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolChoiceUnion) GetServerLabel() *string {
	if vt := u.OfOpenAIResponseInputToolChoiceMcpTool; vt != nil {
		return &vt.ServerLabel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolChoiceUnion) GetType() *string {
	if vt := u.OfOpenAIResponseInputToolChoiceAllowedTools; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputToolChoiceFileSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputToolChoiceWebSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputToolChoiceFunctionTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputToolChoiceMcpTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputToolChoiceCustomTool; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolChoiceUnion) GetName() *string {
	if vt := u.OfOpenAIResponseInputToolChoiceFunctionTool; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfOpenAIResponseInputToolChoiceMcpTool; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfOpenAIResponseInputToolChoiceCustomTool; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMode string

const (
	ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceModeAuto     ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMode = "auto"
	ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceModeRequired ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMode = "required"
	ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceModeNone     ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMode = "none"
)

// Constrains the tools available to the model to a pre-defined set.
//
// The property Tools is required.
type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools struct {
	Tools []map[string]string `json:"tools,omitzero,required"`
	// Any of "auto", "required".
	Mode string `json:"mode,omitzero"`
	// Any of "allowed_tools".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools](
		"mode", "auto", "required",
	)
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceAllowedTools](
		"type", "allowed_tools",
	)
}

// Indicates that the model should use file search to generate a response.
type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFileSearch struct {
	// Any of "file_search".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFileSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFileSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFileSearch](
		"type", "file_search",
	)
}

// Indicates that the model should use web search to generate a response
type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceWebSearch struct {
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11",
	// "web_search_2025_08_26".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceWebSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceWebSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceWebSearch](
		"type", "web_search", "web_search_preview", "web_search_preview_2025_03_11", "web_search_2025_08_26",
	)
}

// Forces the model to call a specific function.
//
// The property Name is required.
type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFunctionTool struct {
	Name string `json:"name,required"`
	// Any of "function".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFunctionTool) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFunctionTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFunctionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceFunctionTool](
		"type", "function",
	)
}

// Forces the model to call a specific tool on a remote MCP server
//
// The property ServerLabel is required.
type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMcpTool struct {
	ServerLabel string            `json:"server_label,required"`
	Name        param.Opt[string] `json:"name,omitzero"`
	// Any of "mcp".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMcpTool) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMcpTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMcpTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceMcpTool](
		"type", "mcp",
	)
}

// Forces the model to call a custom tool.
//
// The property Name is required.
type ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceCustomTool struct {
	Name string `json:"name,required"`
	// Any of "custom".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceCustomTool) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceCustomTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceCustomTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolChoiceOpenAIResponseInputToolChoiceCustomTool](
		"type", "custom",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolUnion struct {
	OfOpenAIResponseInputToolWebSearch *ResponseNewParamsToolOpenAIResponseInputToolWebSearch `json:",omitzero,inline"`
	OfFileSearch                       *ResponseNewParamsToolFileSearch                       `json:",omitzero,inline"`
	OfFunction                         *ResponseNewParamsToolFunction                         `json:",omitzero,inline"`
	OfMcp                              *ResponseNewParamsToolMcp                              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseInputToolWebSearch, u.OfFileSearch, u.OfFunction, u.OfMcp)
}
func (u *ResponseNewParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseInputToolWebSearch) {
		return u.OfOpenAIResponseInputToolWebSearch
	} else if !param.IsOmitted(u.OfFileSearch) {
		return u.OfFileSearch
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfMcp) {
		return u.OfMcp
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetSearchContextSize() *string {
	if vt := u.OfOpenAIResponseInputToolWebSearch; vt != nil && vt.SearchContextSize.Valid() {
		return &vt.SearchContextSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetVectorStoreIDs() []string {
	if vt := u.OfFileSearch; vt != nil {
		return vt.VectorStoreIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetFilters() map[string]any {
	if vt := u.OfFileSearch; vt != nil {
		return vt.Filters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetMaxNumResults() *int64 {
	if vt := u.OfFileSearch; vt != nil && vt.MaxNumResults.Valid() {
		return &vt.MaxNumResults.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetRankingOptions() *ResponseNewParamsToolFileSearchRankingOptions {
	if vt := u.OfFileSearch; vt != nil {
		return &vt.RankingOptions
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetName() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetParameters() map[string]any {
	if vt := u.OfFunction; vt != nil {
		return vt.Parameters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetDescription() *string {
	if vt := u.OfFunction; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetStrict() *bool {
	if vt := u.OfFunction; vt != nil && vt.Strict.Valid() {
		return &vt.Strict.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetServerLabel() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ServerLabel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetServerURL() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ServerURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetAllowedTools() *ResponseNewParamsToolMcpAllowedToolsUnion {
	if vt := u.OfMcp; vt != nil {
		return &vt.AllowedTools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetAuthorization() *string {
	if vt := u.OfMcp; vt != nil && vt.Authorization.Valid() {
		return &vt.Authorization.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetHeaders() map[string]any {
	if vt := u.OfMcp; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetRequireApproval() *ResponseNewParamsToolMcpRequireApprovalUnion {
	if vt := u.OfMcp; vt != nil {
		return &vt.RequireApproval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetType() *string {
	if vt := u.OfOpenAIResponseInputToolWebSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcp; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsToolUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search"),
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_preview"),
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_preview_2025_03_11"),
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_2025_08_26"),
		apijson.Discriminator[ResponseNewParamsToolFileSearch]("file_search"),
		apijson.Discriminator[ResponseNewParamsToolFunction]("function"),
		apijson.Discriminator[ResponseNewParamsToolMcp]("mcp"),
	)
}

// Web search tool configuration for OpenAI response inputs.
type ResponseNewParamsToolOpenAIResponseInputToolWebSearch struct {
	SearchContextSize param.Opt[string] `json:"search_context_size,omitzero"`
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11",
	// "web_search_2025_08_26".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolOpenAIResponseInputToolWebSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolOpenAIResponseInputToolWebSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch](
		"type", "web_search", "web_search_preview", "web_search_preview_2025_03_11", "web_search_2025_08_26",
	)
}

// File search tool configuration for OpenAI response inputs.
//
// The property VectorStoreIDs is required.
type ResponseNewParamsToolFileSearch struct {
	VectorStoreIDs []string         `json:"vector_store_ids,omitzero,required"`
	MaxNumResults  param.Opt[int64] `json:"max_num_results,omitzero"`
	Filters        map[string]any   `json:"filters,omitzero"`
	// Options for ranking and filtering search results.
	RankingOptions ResponseNewParamsToolFileSearchRankingOptions `json:"ranking_options,omitzero"`
	// Any of "file_search".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolFileSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolFileSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolFileSearch](
		"type", "file_search",
	)
}

// Options for ranking and filtering search results.
type ResponseNewParamsToolFileSearchRankingOptions struct {
	Ranker         param.Opt[string]  `json:"ranker,omitzero"`
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolFileSearchRankingOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolFileSearchRankingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool configuration for OpenAI response inputs.
//
// The properties Name, Parameters are required.
type ResponseNewParamsToolFunction struct {
	Parameters  map[string]any    `json:"parameters,omitzero,required"`
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Strict      param.Opt[bool]   `json:"strict,omitzero"`
	// Any of "function".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolFunction) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolFunction](
		"type", "function",
	)
}

// Model Context Protocol (MCP) tool configuration for OpenAI response inputs.
//
// The properties ServerLabel, ServerURL are required.
type ResponseNewParamsToolMcp struct {
	ServerLabel   string            `json:"server_label,required"`
	ServerURL     string            `json:"server_url,required"`
	Authorization param.Opt[string] `json:"authorization,omitzero"`
	// Filter configuration for restricting which MCP tools can be used.
	AllowedTools ResponseNewParamsToolMcpAllowedToolsUnion `json:"allowed_tools,omitzero"`
	Headers      map[string]any                            `json:"headers,omitzero"`
	// Filter configuration for MCP tool approval requirements.
	RequireApproval ResponseNewParamsToolMcpRequireApprovalUnion `json:"require_approval,omitzero"`
	// Any of "mcp".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolMcp) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolMcp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ResponseNewParamsToolMcp](
		"type", "mcp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolMcpAllowedToolsUnion struct {
	OfListString         []string                                                `json:",omitzero,inline"`
	OfAllowedToolsFilter *ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolMcpAllowedToolsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfListString, u.OfAllowedToolsFilter)
}
func (u *ResponseNewParamsToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolMcpAllowedToolsUnion) asAny() any {
	if !param.IsOmitted(u.OfListString) {
		return &u.OfListString
	} else if !param.IsOmitted(u.OfAllowedToolsFilter) {
		return u.OfAllowedToolsFilter
	}
	return nil
}

// Filter configuration for restricting which MCP tools can be used.
type ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter struct {
	ToolNames []string `json:"tool_names,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolMcpRequireApprovalUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfResponseNewsToolMcpRequireApprovalString)
	OfResponseNewsToolMcpRequireApprovalString param.Opt[ResponseNewParamsToolMcpRequireApprovalString] `json:",omitzero,inline"`
	OfApprovalFilter                           *ResponseNewParamsToolMcpRequireApprovalApprovalFilter   `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolMcpRequireApprovalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfResponseNewsToolMcpRequireApprovalString, u.OfApprovalFilter)
}
func (u *ResponseNewParamsToolMcpRequireApprovalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolMcpRequireApprovalUnion) asAny() any {
	if !param.IsOmitted(u.OfResponseNewsToolMcpRequireApprovalString) {
		return &u.OfResponseNewsToolMcpRequireApprovalString
	} else if !param.IsOmitted(u.OfApprovalFilter) {
		return u.OfApprovalFilter
	}
	return nil
}

type ResponseNewParamsToolMcpRequireApprovalString string

const (
	ResponseNewParamsToolMcpRequireApprovalStringAlways ResponseNewParamsToolMcpRequireApprovalString = "always"
	ResponseNewParamsToolMcpRequireApprovalStringNever  ResponseNewParamsToolMcpRequireApprovalString = "never"
)

// Filter configuration for MCP tool approval requirements.
type ResponseNewParamsToolMcpRequireApprovalApprovalFilter struct {
	Always []string `json:"always,omitzero"`
	Never  []string `json:"never,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolMcpRequireApprovalApprovalFilter) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolMcpRequireApprovalApprovalFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolMcpRequireApprovalApprovalFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListParams struct {
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	// Sort order for paginated responses.
	//
	// Any of "asc", "desc".
	Order ResponseListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResponseListParams]'s query parameters as `url.Values`.
func (r ResponseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for paginated responses.
type ResponseListParamsOrder string

const (
	ResponseListParamsOrderAsc  ResponseListParamsOrder = "asc"
	ResponseListParamsOrderDesc ResponseListParamsOrder = "desc"
)
