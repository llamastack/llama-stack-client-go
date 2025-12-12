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
)

// ConversationItemService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationItemService] method instead.
type ConversationItemService struct {
	Options []option.RequestOption
}

// NewConversationItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConversationItemService(opts ...option.RequestOption) (r ConversationItemService) {
	r = ConversationItemService{}
	r.Options = opts
	return
}

// Create items.
//
// Create items in the conversation.
func (r *ConversationItemService) New(ctx context.Context, conversationID string, body ConversationItemNewParams, opts ...option.RequestOption) (res *ConversationItemNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List items.
//
// List items in the conversation.
func (r *ConversationItemService) List(ctx context.Context, conversationID string, query ConversationItemListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[ConversationItemListResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items", conversationID)
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

// List items.
//
// List items in the conversation.
func (r *ConversationItemService) ListAutoPaging(ctx context.Context, conversationID string, query ConversationItemListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ConversationItemListResponseUnion] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, conversationID, query, opts...))
}

// Delete an item.
//
// Delete a conversation item.
func (r *ConversationItemService) Delete(ctx context.Context, itemID string, body ConversationItemDeleteParams, opts ...option.RequestOption) (res *ConversationItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ConversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items/%s", body.ConversationID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve an item.
//
// Retrieve a conversation item.
func (r *ConversationItemService) Get(ctx context.Context, itemID string, query ConversationItemGetParams, opts ...option.RequestOption) (res *ConversationItemGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ConversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items/%s", query.ConversationID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of conversation items with pagination.
type ConversationItemNewResponse struct {
	// List of conversation items
	Data []ConversationItemNewResponseDataUnion `json:"data,required"`
	// The ID of the first item in the list
	FirstID string `json:"first_id,nullable"`
	// Whether there are more items available
	HasMore bool `json:"has_more"`
	// The ID of the last item in the list
	LastID string `json:"last_id,nullable"`
	// Object type
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataUnion contains all possible properties and values
// from [ConversationItemNewResponseDataMessage],
// [ConversationItemNewResponseDataWebSearchCall],
// [ConversationItemNewResponseDataFileSearchCall],
// [ConversationItemNewResponseDataFunctionCall],
// [ConversationItemNewResponseDataFunctionCallOutput],
// [ConversationItemNewResponseDataMcpApprovalRequest],
// [ConversationItemNewResponseDataMcpApprovalResponse],
// [ConversationItemNewResponseDataMcpCall],
// [ConversationItemNewResponseDataMcpListTools].
//
// Use the [ConversationItemNewResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataUnion struct {
	// This field is from variant [ConversationItemNewResponseDataMessage].
	Content ConversationItemNewResponseDataMessageContentUnion `json:"content"`
	// This field is from variant [ConversationItemNewResponseDataMessage].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools".
	Type string `json:"type"`
	// This field is from variant [ConversationItemNewResponseDataFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemNewResponseDataFileSearchCall].
	Results     []ConversationItemNewResponseDataFileSearchCallResult `json:"results"`
	Arguments   string                                                `json:"arguments"`
	CallID      string                                                `json:"call_id"`
	Name        string                                                `json:"name"`
	Output      string                                                `json:"output"`
	ServerLabel string                                                `json:"server_label"`
	// This field is from variant [ConversationItemNewResponseDataMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant [ConversationItemNewResponseDataMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant [ConversationItemNewResponseDataMcpApprovalResponse].
	Reason string `json:"reason"`
	// This field is from variant [ConversationItemNewResponseDataMcpCall].
	Error string `json:"error"`
	// This field is from variant [ConversationItemNewResponseDataMcpListTools].
	Tools []ConversationItemNewResponseDataMcpListToolsTool `json:"tools"`
	JSON  struct {
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
		Output            respjson.Field
		ServerLabel       respjson.Field
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Reason            respjson.Field
		Error             respjson.Field
		Tools             respjson.Field
		raw               string
	} `json:"-"`
}

// anyConversationItemNewResponseData is implemented by each variant of
// [ConversationItemNewResponseDataUnion] to add type safety for the return type of
// [ConversationItemNewResponseDataUnion.AsAny]
type anyConversationItemNewResponseData interface {
	implConversationItemNewResponseDataUnion()
}

func (ConversationItemNewResponseDataMessage) implConversationItemNewResponseDataUnion()            {}
func (ConversationItemNewResponseDataWebSearchCall) implConversationItemNewResponseDataUnion()      {}
func (ConversationItemNewResponseDataFileSearchCall) implConversationItemNewResponseDataUnion()     {}
func (ConversationItemNewResponseDataFunctionCall) implConversationItemNewResponseDataUnion()       {}
func (ConversationItemNewResponseDataFunctionCallOutput) implConversationItemNewResponseDataUnion() {}
func (ConversationItemNewResponseDataMcpApprovalRequest) implConversationItemNewResponseDataUnion() {}
func (ConversationItemNewResponseDataMcpApprovalResponse) implConversationItemNewResponseDataUnion() {
}
func (ConversationItemNewResponseDataMcpCall) implConversationItemNewResponseDataUnion()      {}
func (ConversationItemNewResponseDataMcpListTools) implConversationItemNewResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessage:
//	case llamastackclient.ConversationItemNewResponseDataWebSearchCall:
//	case llamastackclient.ConversationItemNewResponseDataFileSearchCall:
//	case llamastackclient.ConversationItemNewResponseDataFunctionCall:
//	case llamastackclient.ConversationItemNewResponseDataFunctionCallOutput:
//	case llamastackclient.ConversationItemNewResponseDataMcpApprovalRequest:
//	case llamastackclient.ConversationItemNewResponseDataMcpApprovalResponse:
//	case llamastackclient.ConversationItemNewResponseDataMcpCall:
//	case llamastackclient.ConversationItemNewResponseDataMcpListTools:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataUnion) AsAny() anyConversationItemNewResponseData {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "function_call_output":
		return u.AsFunctionCallOutput()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	case "mcp_approval_response":
		return u.AsMcpApprovalResponse()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	}
	return nil
}

func (u ConversationItemNewResponseDataUnion) AsMessage() (v ConversationItemNewResponseDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsWebSearchCall() (v ConversationItemNewResponseDataWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsFileSearchCall() (v ConversationItemNewResponseDataFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsFunctionCall() (v ConversationItemNewResponseDataFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsFunctionCallOutput() (v ConversationItemNewResponseDataFunctionCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsMcpApprovalRequest() (v ConversationItemNewResponseDataMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsMcpApprovalResponse() (v ConversationItemNewResponseDataMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsMcpCall() (v ConversationItemNewResponseDataMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsMcpListTools() (v ConversationItemNewResponseDataMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemNewResponseDataMessage struct {
	Content ConversationItemNewResponseDataMessageContentUnion `json:"content,required"`
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
func (r ConversationItemNewResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentUnion contains all possible
// properties and values from [string],
// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ConversationItemNewResponseDataMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
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

// anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                             `json:"text,required"`
	Annotations []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                               `json:"token,required"`
	Logprob     float64                                                                                                                                                              `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                              `json:"bytes,nullable"`
	TopLogprobs []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataWebSearchCall struct {
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
func (r ConversationItemNewResponseDataWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFileSearchCall struct {
	ID      string                                                `json:"id,required"`
	Queries []string                                              `json:"queries,required"`
	Status  string                                                `json:"status,required"`
	Results []ConversationItemNewResponseDataFileSearchCallResult `json:"results,nullable"`
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
func (r ConversationItemNewResponseDataFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ConversationItemNewResponseDataFileSearchCallResult struct {
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
func (r ConversationItemNewResponseDataFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFunctionCall struct {
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
func (r ConversationItemNewResponseDataFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ConversationItemNewResponseDataFunctionCallOutput struct {
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
func (r ConversationItemNewResponseDataFunctionCallOutput) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ConversationItemNewResponseDataMcpApprovalRequest struct {
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
func (r ConversationItemNewResponseDataMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ConversationItemNewResponseDataMcpApprovalResponse struct {
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
func (r ConversationItemNewResponseDataMcpApprovalResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemNewResponseDataMcpCall struct {
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
func (r ConversationItemNewResponseDataMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemNewResponseDataMcpListTools struct {
	ID          string                                            `json:"id,required"`
	ServerLabel string                                            `json:"server_label,required"`
	Tools       []ConversationItemNewResponseDataMcpListToolsTool `json:"tools,required"`
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
func (r ConversationItemNewResponseDataMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ConversationItemNewResponseDataMcpListToolsTool struct {
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
func (r ConversationItemNewResponseDataMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseUnion contains all possible properties and values
// from [ConversationItemListResponseMessage],
// [ConversationItemListResponseWebSearchCall],
// [ConversationItemListResponseFileSearchCall],
// [ConversationItemListResponseFunctionCall],
// [ConversationItemListResponseFunctionCallOutput],
// [ConversationItemListResponseMcpApprovalRequest],
// [ConversationItemListResponseMcpApprovalResponse],
// [ConversationItemListResponseMcpCall],
// [ConversationItemListResponseMcpListTools].
//
// Use the [ConversationItemListResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseUnion struct {
	// This field is from variant [ConversationItemListResponseMessage].
	Content ConversationItemListResponseMessageContentUnion `json:"content"`
	// This field is from variant [ConversationItemListResponseMessage].
	Role   string `json:"role"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools".
	Type string `json:"type"`
	// This field is from variant [ConversationItemListResponseFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemListResponseFileSearchCall].
	Results     []ConversationItemListResponseFileSearchCallResult `json:"results"`
	Arguments   string                                             `json:"arguments"`
	CallID      string                                             `json:"call_id"`
	Name        string                                             `json:"name"`
	Output      string                                             `json:"output"`
	ServerLabel string                                             `json:"server_label"`
	// This field is from variant [ConversationItemListResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant [ConversationItemListResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant [ConversationItemListResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	// This field is from variant [ConversationItemListResponseMcpCall].
	Error string `json:"error"`
	// This field is from variant [ConversationItemListResponseMcpListTools].
	Tools []ConversationItemListResponseMcpListToolsTool `json:"tools"`
	JSON  struct {
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
		Output            respjson.Field
		ServerLabel       respjson.Field
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Reason            respjson.Field
		Error             respjson.Field
		Tools             respjson.Field
		raw               string
	} `json:"-"`
}

// anyConversationItemListResponse is implemented by each variant of
// [ConversationItemListResponseUnion] to add type safety for the return type of
// [ConversationItemListResponseUnion.AsAny]
type anyConversationItemListResponse interface {
	implConversationItemListResponseUnion()
}

func (ConversationItemListResponseMessage) implConversationItemListResponseUnion()             {}
func (ConversationItemListResponseWebSearchCall) implConversationItemListResponseUnion()       {}
func (ConversationItemListResponseFileSearchCall) implConversationItemListResponseUnion()      {}
func (ConversationItemListResponseFunctionCall) implConversationItemListResponseUnion()        {}
func (ConversationItemListResponseFunctionCallOutput) implConversationItemListResponseUnion()  {}
func (ConversationItemListResponseMcpApprovalRequest) implConversationItemListResponseUnion()  {}
func (ConversationItemListResponseMcpApprovalResponse) implConversationItemListResponseUnion() {}
func (ConversationItemListResponseMcpCall) implConversationItemListResponseUnion()             {}
func (ConversationItemListResponseMcpListTools) implConversationItemListResponseUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessage:
//	case llamastackclient.ConversationItemListResponseWebSearchCall:
//	case llamastackclient.ConversationItemListResponseFileSearchCall:
//	case llamastackclient.ConversationItemListResponseFunctionCall:
//	case llamastackclient.ConversationItemListResponseFunctionCallOutput:
//	case llamastackclient.ConversationItemListResponseMcpApprovalRequest:
//	case llamastackclient.ConversationItemListResponseMcpApprovalResponse:
//	case llamastackclient.ConversationItemListResponseMcpCall:
//	case llamastackclient.ConversationItemListResponseMcpListTools:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseUnion) AsAny() anyConversationItemListResponse {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "function_call_output":
		return u.AsFunctionCallOutput()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	case "mcp_approval_response":
		return u.AsMcpApprovalResponse()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	}
	return nil
}

func (u ConversationItemListResponseUnion) AsMessage() (v ConversationItemListResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsWebSearchCall() (v ConversationItemListResponseWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsFileSearchCall() (v ConversationItemListResponseFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsFunctionCall() (v ConversationItemListResponseFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsFunctionCallOutput() (v ConversationItemListResponseFunctionCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsMcpApprovalRequest() (v ConversationItemListResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsMcpApprovalResponse() (v ConversationItemListResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsMcpCall() (v ConversationItemListResponseMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsMcpListTools() (v ConversationItemListResponseMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemListResponseMessage struct {
	Content ConversationItemListResponseMessageContentUnion `json:"content,required"`
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
func (r ConversationItemListResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseMessageContentUnion contains all possible properties
// and values from [string],
// [[]ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ConversationItemListResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemListResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
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

// anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                          `json:"text,required"`
	Annotations []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                            `json:"token,required"`
	Logprob     float64                                                                                                                                                           `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                           `json:"bytes,nullable"`
	TopLogprobs []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ConversationItemListResponseWebSearchCall struct {
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
func (r ConversationItemListResponseWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemListResponseFileSearchCall struct {
	ID      string                                             `json:"id,required"`
	Queries []string                                           `json:"queries,required"`
	Status  string                                             `json:"status,required"`
	Results []ConversationItemListResponseFileSearchCallResult `json:"results,nullable"`
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
func (r ConversationItemListResponseFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ConversationItemListResponseFileSearchCallResult struct {
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
func (r ConversationItemListResponseFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemListResponseFunctionCall struct {
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
func (r ConversationItemListResponseFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ConversationItemListResponseFunctionCallOutput struct {
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
func (r ConversationItemListResponseFunctionCallOutput) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ConversationItemListResponseMcpApprovalRequest struct {
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
func (r ConversationItemListResponseMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ConversationItemListResponseMcpApprovalResponse struct {
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
func (r ConversationItemListResponseMcpApprovalResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemListResponseMcpCall struct {
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
func (r ConversationItemListResponseMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemListResponseMcpListTools struct {
	ID          string                                         `json:"id,required"`
	ServerLabel string                                         `json:"server_label,required"`
	Tools       []ConversationItemListResponseMcpListToolsTool `json:"tools,required"`
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
func (r ConversationItemListResponseMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ConversationItemListResponseMcpListToolsTool struct {
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
func (r ConversationItemListResponseMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for deleted conversation item.
type ConversationItemDeleteResponse struct {
	// The deleted item identifier
	ID string `json:"id,required"`
	// Whether the object was deleted
	Deleted bool `json:"deleted"`
	// Object type
	Object string `json:"object"`
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
func (r ConversationItemDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemGetResponse struct {
	Content ConversationItemGetResponseContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemGetResponseRole `json:"role,required"`
	ID     string                          `json:"id,nullable"`
	Status string                          `json:"status,nullable"`
	// Any of "message".
	Type ConversationItemGetResponseType `json:"type"`
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
func (r ConversationItemGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseContentUnion contains all possible properties and
// values from [string],
// [[]ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ConversationItemGetResponseContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemGetResponseContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal() (v []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail string `json:"detail"`
	FileID string `json:"file_id"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content for input messages in OpenAI response format.
type ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText],
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
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

// anyConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem interface {
	implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion()
}

func (ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}
func (ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsAny() anyConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                            `json:"text,required"`
	Annotations []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,nullable"`
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                              `json:"token,required"`
	Logprob     float64                                                                                                                                             `json:"logprob,required"`
	Bytes       []int64                                                                                                                                             `json:"bytes,nullable"`
	TopLogprobs []ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,nullable"`
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseRole string

const (
	ConversationItemGetResponseRoleSystem    ConversationItemGetResponseRole = "system"
	ConversationItemGetResponseRoleDeveloper ConversationItemGetResponseRole = "developer"
	ConversationItemGetResponseRoleUser      ConversationItemGetResponseRole = "user"
	ConversationItemGetResponseRoleAssistant ConversationItemGetResponseRole = "assistant"
)

type ConversationItemGetResponseType string

const (
	ConversationItemGetResponseTypeMessage ConversationItemGetResponseType = "message"
)

type ConversationItemNewParams struct {
	Items []ConversationItemNewParamsItemUnion `json:"items,omitzero,required"`
	paramObj
}

func (r ConversationItemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemUnion struct {
	OfMessage             *ConversationItemNewParamsItemMessage             `json:",omitzero,inline"`
	OfWebSearchCall       *ConversationItemNewParamsItemWebSearchCall       `json:",omitzero,inline"`
	OfFileSearchCall      *ConversationItemNewParamsItemFileSearchCall      `json:",omitzero,inline"`
	OfFunctionCall        *ConversationItemNewParamsItemFunctionCall        `json:",omitzero,inline"`
	OfFunctionCallOutput  *ConversationItemNewParamsItemFunctionCallOutput  `json:",omitzero,inline"`
	OfMcpApprovalRequest  *ConversationItemNewParamsItemMcpApprovalRequest  `json:",omitzero,inline"`
	OfMcpApprovalResponse *ConversationItemNewParamsItemMcpApprovalResponse `json:",omitzero,inline"`
	OfMcpCall             *ConversationItemNewParamsItemMcpCall             `json:",omitzero,inline"`
	OfMcpListTools        *ConversationItemNewParamsItemMcpListTools        `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessage,
		u.OfWebSearchCall,
		u.OfFileSearchCall,
		u.OfFunctionCall,
		u.OfFunctionCallOutput,
		u.OfMcpApprovalRequest,
		u.OfMcpApprovalResponse,
		u.OfMcpCall,
		u.OfMcpListTools)
}
func (u *ConversationItemNewParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemUnion) asAny() any {
	if !param.IsOmitted(u.OfMessage) {
		return u.OfMessage
	} else if !param.IsOmitted(u.OfWebSearchCall) {
		return u.OfWebSearchCall
	} else if !param.IsOmitted(u.OfFileSearchCall) {
		return u.OfFileSearchCall
	} else if !param.IsOmitted(u.OfFunctionCall) {
		return u.OfFunctionCall
	} else if !param.IsOmitted(u.OfFunctionCallOutput) {
		return u.OfFunctionCallOutput
	} else if !param.IsOmitted(u.OfMcpApprovalRequest) {
		return u.OfMcpApprovalRequest
	} else if !param.IsOmitted(u.OfMcpApprovalResponse) {
		return u.OfMcpApprovalResponse
	} else if !param.IsOmitted(u.OfMcpCall) {
		return u.OfMcpCall
	} else if !param.IsOmitted(u.OfMcpListTools) {
		return u.OfMcpListTools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetContent() *ConversationItemNewParamsItemMessageContentUnion {
	if vt := u.OfMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetRole() *string {
	if vt := u.OfMessage; vt != nil {
		return &vt.Role
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetQueries() []string {
	if vt := u.OfFileSearchCall; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetResults() []ConversationItemNewParamsItemFileSearchCallResult {
	if vt := u.OfFileSearchCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetApprovalRequestID() *string {
	if vt := u.OfMcpApprovalResponse; vt != nil {
		return &vt.ApprovalRequestID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetApprove() *bool {
	if vt := u.OfMcpApprovalResponse; vt != nil {
		return &vt.Approve
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetReason() *string {
	if vt := u.OfMcpApprovalResponse; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetError() *string {
	if vt := u.OfMcpCall; vt != nil && vt.Error.Valid() {
		return &vt.Error.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetTools() []ConversationItemNewParamsItemMcpListToolsTool {
	if vt := u.OfMcpListTools; vt != nil {
		return vt.Tools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetID() *string {
	if vt := u.OfMessage; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfWebSearchCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfFileSearchCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfFunctionCall; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfFunctionCallOutput; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfMcpApprovalRequest; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMcpApprovalResponse; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMcpListTools; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetStatus() *string {
	if vt := u.OfMessage; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfWebSearchCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfFileSearchCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfFunctionCall; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfFunctionCallOutput; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetType() *string {
	if vt := u.OfMessage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebSearchCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileSearchCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunctionCallOutput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcpApprovalResponse; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcpListTools; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetArguments() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.Arguments)
	} else if vt := u.OfMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Arguments)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.Arguments)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetCallID() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.CallID)
	} else if vt := u.OfFunctionCallOutput; vt != nil {
		return (*string)(&vt.CallID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetName() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetOutput() *string {
	if vt := u.OfFunctionCallOutput; vt != nil {
		return (*string)(&vt.Output)
	} else if vt := u.OfMcpCall; vt != nil && vt.Output.Valid() {
		return &vt.Output.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetServerLabel() *string {
	if vt := u.OfMcpApprovalRequest; vt != nil {
		return (*string)(&vt.ServerLabel)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.ServerLabel)
	} else if vt := u.OfMcpListTools; vt != nil {
		return (*string)(&vt.ServerLabel)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessage]("message"),
		apijson.Discriminator[ConversationItemNewParamsItemWebSearchCall]("web_search_call"),
		apijson.Discriminator[ConversationItemNewParamsItemFileSearchCall]("file_search_call"),
		apijson.Discriminator[ConversationItemNewParamsItemFunctionCall]("function_call"),
		apijson.Discriminator[ConversationItemNewParamsItemFunctionCallOutput]("function_call_output"),
		apijson.Discriminator[ConversationItemNewParamsItemMcpApprovalRequest]("mcp_approval_request"),
		apijson.Discriminator[ConversationItemNewParamsItemMcpApprovalResponse]("mcp_approval_response"),
		apijson.Discriminator[ConversationItemNewParamsItemMcpCall]("mcp_call"),
		apijson.Discriminator[ConversationItemNewParamsItemMcpListTools]("mcp_list_tools"),
	)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role are required.
type ConversationItemNewParamsItemMessage struct {
	Content ConversationItemNewParamsItemMessageContentUnion `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string            `json:"role,omitzero,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// Any of "message".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessage](
		"role", "system", "developer", "user", "assistant",
	)
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessage](
		"type", "message",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentUnion struct {
	OfString                                                                                                               param.Opt[string]                                                                                                                                                          `json:",omitzero,inline"`
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",omitzero,inline"`
	OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal                                []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion                                `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile, u.OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal)
}
func (u *ConversationItemNewParamsItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentUnion) asAny() any {
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
type ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	OfInputText  *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText  `json:",omitzero,inline"`
	OfInputImage *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage `json:",omitzero,inline"`
	OfInputFile  *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile  `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage, u.OfInputFile)
}
func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) asAny() any {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return &vt.Detail
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileData() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileData.Valid() {
		return &vt.FileData.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileURL() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileURL.Valid() {
		return &vt.FileURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFilename() *string {
	if vt := u.OfInputFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetType() *string {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileID() *string {
	if vt := u.OfInputImage; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	} else if vt := u.OfInputFile; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText]("input_text"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage]("input_image"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile]("input_file"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The property Text is required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText](
		"type", "input_text",
	)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	// Any of "input_image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"detail", "low", "high", "auto",
	)
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"type", "input_image",
	)
}

// File content for input messages in OpenAI response format.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	FileURL  param.Opt[string] `json:"file_url,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Any of "input_file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile](
		"type", "input_file",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion struct {
	OfOutputText *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText `json:",omitzero,inline"`
	OfRefusal    *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal    `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOutputText, u.OfRefusal)
}
func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOutputText) {
		return u.OfOutputText
	} else if !param.IsOmitted(u.OfRefusal) {
		return u.OfRefusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetText() *string {
	if vt := u.OfOutputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetAnnotations() []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion {
	if vt := u.OfOutputText; vt != nil {
		return vt.Annotations
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetLogprobs() []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob {
	if vt := u.OfOutputText; vt != nil {
		return vt.Logprobs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetRefusal() *string {
	if vt := u.OfRefusal; vt != nil {
		return &vt.Refusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetType() *string {
	if vt := u.OfOutputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRefusal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText]("output_text"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal]("refusal"),
	)
}

// The property Text is required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                          `json:"text,required"`
	Logprobs    []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,omitzero"`
	Annotations []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations,omitzero"`
	// Any of "output_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText](
		"type", "output_text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	OfFileCitation          *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation          `json:",omitzero,inline"`
	OfURLCitation           *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation           `json:",omitzero,inline"`
	OfContainerFileCitation *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation `json:",omitzero,inline"`
	OfFilePath              *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) asAny() any {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFileID() *string {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetType() *string {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation]("file_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation]("url_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation]("container_file_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath]("file_path"),
	)
}

// File citation annotation for referencing specific files in response content.
//
// The properties FileID, Filename, Index are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation](
		"type", "file_citation",
	)
}

// URL citation annotation for referencing external web resources.
//
// The properties EndIndex, StartIndex, Title, URL are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation](
		"type", "url_citation",
	)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation](
		"type", "container_file_citation",
	)
}

// The properties FileID, Index are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath](
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
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                            `json:"token,required"`
	Logprob     float64                                                                                                                                                           `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                           `json:"bytes,omitzero"`
	TopLogprobs []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
//
// The properties Token, Logprob are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
//
// The property Refusal is required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal](
		"type", "refusal",
	)
}

// Web search tool call output message for OpenAI responses.
//
// The properties ID, Status are required.
type ConversationItemNewParamsItemWebSearchCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemWebSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemWebSearchCall](
		"type", "web_search_call",
	)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status are required.
type ConversationItemNewParamsItemFileSearchCall struct {
	ID      string                                              `json:"id,required"`
	Queries []string                                            `json:"queries,omitzero,required"`
	Status  string                                              `json:"status,required"`
	Results []ConversationItemNewParamsItemFileSearchCallResult `json:"results,omitzero"`
	// Any of "file_search_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemFileSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFileSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemFileSearchCall](
		"type", "file_search_call",
	)
}

// Search results returned by the file search operation.
//
// The properties Attributes, FileID, Filename, Score, Text are required.
type ConversationItemNewParamsItemFileSearchCallResult struct {
	Attributes map[string]any `json:"attributes,omitzero,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	paramObj
}

func (r ConversationItemNewParamsItemFileSearchCallResult) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFileSearchCallResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
//
// The properties Arguments, CallID, Name are required.
type ConversationItemNewParamsItemFunctionCall struct {
	Arguments string            `json:"arguments,required"`
	CallID    string            `json:"call_id,required"`
	Name      string            `json:"name,required"`
	ID        param.Opt[string] `json:"id,omitzero"`
	Status    param.Opt[string] `json:"status,omitzero"`
	// Any of "function_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemFunctionCall](
		"type", "function_call",
	)
}

// This represents the output of a function call that gets passed back to the
// model.
//
// The properties CallID, Output are required.
type ConversationItemNewParamsItemFunctionCallOutput struct {
	CallID string            `json:"call_id,required"`
	Output string            `json:"output,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// Any of "function_call_output".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCallOutput) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCallOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemFunctionCallOutput](
		"type", "function_call_output",
	)
}

// A request for human approval of a tool invocation.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ConversationItemNewParamsItemMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpApprovalRequest) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpApprovalRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMcpApprovalRequest](
		"type", "mcp_approval_request",
	)
}

// A response to an MCP approval request.
//
// The properties ApprovalRequestID, Approve are required.
type ConversationItemNewParamsItemMcpApprovalResponse struct {
	ApprovalRequestID string            `json:"approval_request_id,required"`
	Approve           bool              `json:"approve,required"`
	ID                param.Opt[string] `json:"id,omitzero"`
	Reason            param.Opt[string] `json:"reason,omitzero"`
	// Any of "mcp_approval_response".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpApprovalResponse) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpApprovalResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMcpApprovalResponse](
		"type", "mcp_approval_response",
	)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ConversationItemNewParamsItemMcpCall struct {
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

func (r ConversationItemNewParamsItemMcpCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMcpCall](
		"type", "mcp_call",
	)
}

// MCP list tools output message containing available tools from an MCP server.
//
// The properties ID, ServerLabel, Tools are required.
type ConversationItemNewParamsItemMcpListTools struct {
	ID          string                                          `json:"id,required"`
	ServerLabel string                                          `json:"server_label,required"`
	Tools       []ConversationItemNewParamsItemMcpListToolsTool `json:"tools,omitzero,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpListTools) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpListTools
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMcpListTools](
		"type", "mcp_list_tools",
	)
}

// Tool definition returned by MCP list tools operation.
//
// The properties InputSchema, Name are required.
type ConversationItemNewParamsItemMcpListToolsTool struct {
	InputSchema map[string]any    `json:"input_schema,omitzero,required"`
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpListToolsTool) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpListToolsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListParams struct {
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Any of "web_search_call.action.sources", "code_interpreter_call.outputs",
	// "computer_call_output.output.image_url", "file_search_call.results",
	// "message.input_image.image_url", "message.output_text.logprobs",
	// "reasoning.encrypted_content".
	Include []string `query:"include,omitzero" json:"-"`
	// Any of "asc", "desc".
	Order ConversationItemListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationItemListParams]'s query parameters as
// `url.Values`.
func (r ConversationItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConversationItemListParamsOrder string

const (
	ConversationItemListParamsOrderAsc  ConversationItemListParamsOrder = "asc"
	ConversationItemListParamsOrderDesc ConversationItemListParamsOrder = "desc"
)

type ConversationItemDeleteParams struct {
	ConversationID string `path:"conversation_id,required" json:"-"`
	paramObj
}

type ConversationItemGetParams struct {
	ConversationID string `path:"conversation_id,required" json:"-"`
	paramObj
}
