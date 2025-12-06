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
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ResponseInputItemService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseInputItemService] method instead.
type ResponseInputItemService struct {
	Options []option.RequestOption
}

// NewResponseInputItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewResponseInputItemService(opts ...option.RequestOption) (r ResponseInputItemService) {
	r = ResponseInputItemService{}
	r.Options = opts
	return
}

// List input items.
func (r *ResponseInputItemService) List(ctx context.Context, responseID string, query ResponseInputItemListParams, opts ...option.RequestOption) (res *ResponseInputItemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s/input_items", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List container for OpenAI response input items.
type ResponseInputItemListResponse struct {
	Data []ResponseInputItemListResponseDataUnion `json:"data,required"`
	// Any of "list".
	Object ResponseInputItemListResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseInputItemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataUnion contains all possible properties and
// values from [ResponseInputItemListResponseDataOpenAIResponseMessageOutput],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools],
// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest],
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput],
// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutput].
	Content ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion `json:"content"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutput].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall].
	Results     []ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	Arguments   string                                                                                 `json:"arguments"`
	CallID      string                                                                                 `json:"call_id"`
	Name        string                                                                                 `json:"name"`
	ServerLabel string                                                                                 `json:"server_label"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall].
	Error  string `json:"error"`
	Output string `json:"output"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools].
	Tools []ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool `json:"tools"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
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

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMessageOutput() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageMcpCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageMcpListTools() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMcpApprovalRequest() (v ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMcpApprovalResponse() (v ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsResponseInputItemListResponseDataOpenAIResponseMessageOutput() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseInputItemListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutput struct {
	Content ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion `json:"content,required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion
// contains all possible properties and values from [string],
// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal() (v []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                                             `json:"text,required"`
	Annotations []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// Any of "output_text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Annotations respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall struct {
	ID      string                                                                                 `json:"id,required"`
	Queries []string                                                                               `json:"queries,required"`
	Status  string                                                                                 `json:"status,required"`
	Results []ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results,nullable"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools struct {
	ID          string                                                                         `json:"id,required"`
	ServerLabel string                                                                         `json:"server_label,required"`
	Tools       []ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool `json:"tools,required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse struct {
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
func (r ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseObject string

const (
	ResponseInputItemListResponseObjectList ResponseInputItemListResponseObject = "list"
)

type ResponseInputItemListParams struct {
	After   param.Opt[string] `query:"after,omitzero" json:"-"`
	Before  param.Opt[string] `query:"before,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Include []string          `query:"include,omitzero" json:"-"`
	// Sort order for paginated responses.
	//
	// Any of "asc", "desc".
	Order ResponseInputItemListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResponseInputItemListParams]'s query parameters as
// `url.Values`.
func (r ResponseInputItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for paginated responses.
type ResponseInputItemListParamsOrder string

const (
	ResponseInputItemListParamsOrderAsc  ResponseInputItemListParamsOrder = "asc"
	ResponseInputItemListParamsOrderDesc ResponseInputItemListParamsOrder = "desc"
)
