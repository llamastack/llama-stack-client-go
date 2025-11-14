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
// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ConversationItemNewResponseDataMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
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

func (u ConversationItemNewResponseDataMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal() (v []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) {
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

// ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem interface {
	implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion()
}

func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsAny() anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                       `json:"text,required"`
	Annotations []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
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
// [[]ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal]
type ConversationItemListResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     respjson.Field
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

func (u ConversationItemListResponseMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal() (v []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) {
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

// ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
		Type        respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem interface {
	implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion()
}

func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsAny() anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                    `json:"text,required"`
	Annotations []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
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
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text        respjson.Field
		Annotations respjson.Field
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
	OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal                                     []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile, u.OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal)
}
func (u *ConversationItemNewParamsItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile) {
		return &u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
	} else if !param.IsOmitted(u.OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal) {
		return &u.OfListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusal
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
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion struct {
	OfOutputText *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText `json:",omitzero,inline"`
	OfRefusal    *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal    `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOutputText, u.OfRefusal)
}
func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOutputText) {
		return u.OfOutputText
	} else if !param.IsOmitted(u.OfRefusal) {
		return u.OfRefusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) GetText() *string {
	if vt := u.OfOutputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) GetAnnotations() []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion {
	if vt := u.OfOutputText; vt != nil {
		return vt.Annotations
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) GetRefusal() *string {
	if vt := u.OfRefusal; vt != nil {
		return &vt.Refusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion) GetType() *string {
	if vt := u.OfOutputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRefusal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText]("output_text"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal]("refusal"),
	)
}

// The property Text is required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                     `json:"text,required"`
	Annotations []ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations,omitzero"`
	// Any of "output_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputText](
		"type", "output_text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	OfFileCitation          *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation          `json:",omitzero,inline"`
	OfURLCitation           *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation           `json:",omitzero,inline"`
	OfContainerFileCitation *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation `json:",omitzero,inline"`
	OfFilePath              *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) asAny() any {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFileID() *string {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetType() *string {
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
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation]("file_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation]("url_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation]("container_file_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath]("file_path"),
	)
}

// File citation annotation for referencing specific files in response content.
//
// The properties FileID, Filename, Index are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation](
		"type", "file_citation",
	)
}

// URL citation annotation for referencing external web resources.
//
// The properties EndIndex, StartIndex, Title, URL are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation](
		"type", "url_citation",
	)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation](
		"type", "container_file_citation",
	)
}

// The properties FileID, Index are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath](
		"type", "file_path",
	)
}

// Refusal content within a streamed response part.
//
// The property Refusal is required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextOpenAIResponseContentPartRefusalItemRefusal](
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
