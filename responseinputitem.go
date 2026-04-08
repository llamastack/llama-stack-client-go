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
		return nil, err
	}
	path := fmt.Sprintf("v1/responses/%s/input_items", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List container for OpenAI response input items.
type ResponseInputItemListResponse struct {
	Data []ResponseInputItemListResponseDataUnion `json:"data" api:"required"`
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
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItem],
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput],
// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse],
// [ResponseInputItemListResponseDataOpenAIResponseCompaction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataUnion struct {
	// This field is a union of
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion],
	// [[]ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent]
	Content ResponseInputItemListResponseDataUnionContent `json:"content"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutput].
	Role   ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole `json:"role"`
	ID     string                                                           `json:"id"`
	Status string                                                           `json:"status"`
	Type   string                                                           `json:"type"`
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
	Error string `json:"error"`
	// This field is a union of [string],
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion]
	Output ResponseInputItemListResponseDataUnionOutput `json:"output"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools].
	Tools []ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool `json:"tools"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItem].
	Summary []ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemSummary `json:"summary"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseCompaction].
	EncryptedContent string `json:"encrypted_content"`
	JSON             struct {
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
		Summary           respjson.Field
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Reason            respjson.Field
		EncryptedContent  respjson.Field
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

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageReasoningItem() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItem) {
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

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseCompaction() (v ResponseInputItemListResponseDataOpenAIResponseCompaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseInputItemListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataUnionContent is an implicit subunion of
// [ResponseInputItemListResponseDataUnion].
// ResponseInputItemListResponseDataUnionContent provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseInputItemListResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal
// OfResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContentArray]
type ResponseInputItemListResponseDataUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent]
	// instead of an object.
	OfResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContentArray []ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent `json:",inline"`
	JSON                                                                                    struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		OfResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContentArray                                respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ResponseInputItemListResponseDataUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataUnionOutput is an implicit subunion of
// [ResponseInputItemListResponseDataUnion].
// ResponseInputItemListResponseDataUnionOutput provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseInputItemListResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ResponseInputItemListResponseDataUnionOutput struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ResponseInputItemListResponseDataUnionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutput struct {
	Content ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion `json:"content" api:"required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole `json:"role" api:"required"`
	ID     string                                                           `json:"id" api:"nullable"`
	Status string                                                           `json:"status" api:"nullable"`
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
// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
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

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
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
	Detail ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                                                      `json:"file_id"`
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
	Text string `json:"text" api:"required"`
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
	Detail   ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                                                      `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                                                      `json:"image_url" api:"nullable"`
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

type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data" api:"nullable"`
	FileID   string `json:"file_id" api:"nullable"`
	FileURL  string `json:"file_url" api:"nullable"`
	Filename string `json:"filename" api:"nullable"`
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

type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

// ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
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

// anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content within an output message of an OpenAI response.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                                                   `json:"text" api:"required"`
	Annotations []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs" api:"nullable"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id" api:"required"`
	Filename string `json:"filename" api:"required"`
	Index    int64  `json:"index" api:"required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index" api:"required"`
	StartIndex int64  `json:"start_index" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Container file citation annotation referencing a file within a container.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id" api:"required"`
	EndIndex    int64  `json:"end_index" api:"required"`
	FileID      string `json:"file_id" api:"required"`
	Filename    string `json:"filename" api:"required"`
	StartIndex  int64  `json:"start_index" api:"required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File path annotation referencing a generated file in response content.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id" api:"required"`
	Index  int64  `json:"index" api:"required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
	// The top log probabilities for the token.
	TopLogprobs []ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs" api:"nullable"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal" api:"required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageOutputContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputRoleSystem    ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole = "system"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputRoleDeveloper ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole = "developer"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputRoleUser      ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole = "user"
	ResponseInputItemListResponseDataOpenAIResponseMessageOutputRoleAssistant ResponseInputItemListResponseDataOpenAIResponseMessageOutputRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall struct {
	ID     string `json:"id" api:"required"`
	Status string `json:"status" api:"required"`
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
	ID      string                                                                                 `json:"id" api:"required"`
	Queries []string                                                                               `json:"queries" api:"required"`
	Status  string                                                                                 `json:"status" api:"required"`
	Results []ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results" api:"nullable"`
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
	Attributes map[string]any `json:"attributes" api:"required"`
	FileID     string         `json:"file_id" api:"required"`
	Filename   string         `json:"filename" api:"required"`
	Score      float64        `json:"score" api:"required"`
	Text       string         `json:"text" api:"required"`
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
	Arguments string `json:"arguments" api:"required"`
	CallID    string `json:"call_id" api:"required"`
	Name      string `json:"name" api:"required"`
	ID        string `json:"id" api:"nullable"`
	Status    string `json:"status" api:"nullable"`
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
	ID          string `json:"id" api:"required"`
	Arguments   string `json:"arguments" api:"required"`
	Name        string `json:"name" api:"required"`
	ServerLabel string `json:"server_label" api:"required"`
	Error       string `json:"error" api:"nullable"`
	Output      string `json:"output" api:"nullable"`
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
	ID          string                                                                         `json:"id" api:"required"`
	ServerLabel string                                                                         `json:"server_label" api:"required"`
	Tools       []ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool `json:"tools" api:"required"`
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
	InputSchema map[string]any `json:"input_schema" api:"required"`
	Name        string         `json:"name" api:"required"`
	Description string         `json:"description" api:"nullable"`
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
	ID          string `json:"id" api:"required"`
	Arguments   string `json:"arguments" api:"required"`
	Name        string `json:"name" api:"required"`
	ServerLabel string `json:"server_label" api:"required"`
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

// Reasoning output from the model, representing the model's thinking process.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItem struct {
	// Unique identifier for the reasoning output item.
	ID string `json:"id" api:"required"`
	// Summary of the reasoning output.
	Summary []ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemSummary `json:"summary" api:"required"`
	// The reasoning content from the model.
	Content []ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent `json:"content" api:"nullable"`
	// The status of the reasoning output.
	//
	// Any of "in_progress", "completed", "incomplete".
	Status string `json:"status" api:"nullable"`
	// The type identifier, always 'reasoning'.
	//
	// Any of "reasoning".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Summary     respjson.Field
		Content     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summary of reasoning output from the model.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemSummary struct {
	// The summary text of the reasoning output.
	Text string `json:"text" api:"required"`
	// The type identifier, always 'summary_text'.
	//
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text from the model.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent struct {
	// The reasoning text content from the model.
	Text string `json:"text" api:"required"`
	// The type identifier, always 'reasoning_text'.
	//
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
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageReasoningItemContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput struct {
	CallID string                                                                                `json:"call_id" api:"required"`
	Output ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion `json:"output" api:"required"`
	ID     string                                                                                `json:"id" api:"nullable"`
	Status string                                                                                `json:"status" api:"nullable"`
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

// ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion
// contains all possible properties and values from [string],
// [[]ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                                                                   `json:"file_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text" api:"required"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                                                                   `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                                                                   `json:"image_url" api:"nullable"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData string `json:"file_data" api:"nullable"`
	FileID   string `json:"file_id" api:"nullable"`
	FileURL  string `json:"file_url" api:"nullable"`
	Filename string `json:"filename" api:"nullable"`
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
func (r ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

// A response to an MCP approval request.
type ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse struct {
	ApprovalRequestID string `json:"approval_request_id" api:"required"`
	Approve           bool   `json:"approve" api:"required"`
	ID                string `json:"id" api:"nullable"`
	Reason            string `json:"reason" api:"nullable"`
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

// A compaction item that summarizes prior conversation context.
type ResponseInputItemListResponseDataOpenAIResponseCompaction struct {
	EncryptedContent string `json:"encrypted_content" api:"required"`
	ID               string `json:"id" api:"nullable"`
	// Any of "compaction".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptedContent respjson.Field
		ID               respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseCompaction) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseCompaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataRole string

const (
	ResponseInputItemListResponseDataRoleSystem    ResponseInputItemListResponseDataRole = "system"
	ResponseInputItemListResponseDataRoleDeveloper ResponseInputItemListResponseDataRole = "developer"
	ResponseInputItemListResponseDataRoleUser      ResponseInputItemListResponseDataRole = "user"
	ResponseInputItemListResponseDataRoleAssistant ResponseInputItemListResponseDataRole = "assistant"
)

type ResponseInputItemListResponseObject string

const (
	ResponseInputItemListResponseObjectList ResponseInputItemListResponseObject = "list"
)

type ResponseInputItemListParams struct {
	// An item ID to list items after, used for pagination.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// An item ID to list items before, used for pagination.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Additional fields to include in the response.
	//
	// Any of "web_search_call.action.sources", "code_interpreter_call.outputs",
	// "computer_call_output.output.image_url", "file_search_call.results",
	// "message.input_image.image_url", "message.output_text.logprobs",
	// "reasoning.encrypted_content".
	Include []string `query:"include,omitzero" json:"-"`
	// The order to return the input items in.
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

// The order to return the input items in.
type ResponseInputItemListParamsOrder string

const (
	ResponseInputItemListParamsOrderAsc  ResponseInputItemListParamsOrder = "asc"
	ResponseInputItemListParamsOrderDesc ResponseInputItemListParamsOrder = "desc"
)
