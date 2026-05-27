// Copyright (c) The OGX Contributors.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ogxclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ogx-ai/ogx-client-go/internal/apijson"
	"github.com/ogx-ai/ogx-client-go/internal/apiquery"
	"github.com/ogx-ai/ogx-client-go/internal/requestconfig"
	"github.com/ogx-ai/ogx-client-go/option"
	"github.com/ogx-ai/ogx-client-go/packages/pagination"
	"github.com/ogx-ai/ogx-client-go/packages/param"
	"github.com/ogx-ai/ogx-client-go/packages/respjson"
)

// Protocol for conversation management operations.
//
// ConversationItemService contains methods and other services that help with
// interacting with the ogx-client API.
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

// Create items in the conversation.
func (r *ConversationItemService) New(ctx context.Context, conversationID string, body ConversationItemNewParams, opts ...option.RequestOption) (res *ConversationItemNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s/items", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List items in the conversation.
func (r *ConversationItemService) List(ctx context.Context, conversationID string, query ConversationItemListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[ConversationItemListResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
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

// List items in the conversation.
func (r *ConversationItemService) ListAutoPaging(ctx context.Context, conversationID string, query ConversationItemListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ConversationItemListResponseUnion] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, conversationID, query, opts...))
}

// Delete a conversation item.
func (r *ConversationItemService) Delete(ctx context.Context, itemID string, body ConversationItemDeleteParams, opts ...option.RequestOption) (res *ConversationObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ConversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s/items/%s", body.ConversationID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve a conversation item.
func (r *ConversationItemService) Get(ctx context.Context, itemID string, params ConversationItemGetParams, opts ...option.RequestOption) (res *ConversationItemGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ConversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s/items/%s", params.ConversationID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List of conversation items with pagination.
type ConversationItemNewResponse struct {
	// List of conversation items
	Data []ConversationItemNewResponseDataUnion `json:"data" api:"required"`
	// The ID of the first item in the list.
	FirstID string `json:"first_id" api:"required"`
	// Whether there are more items available.
	HasMore bool `json:"has_more" api:"required"`
	// The ID of the last item in the list.
	LastID string `json:"last_id" api:"required"`
	// The type of object returned, must be list.
	//
	// Any of "list".
	Object ConversationItemNewResponseObject `json:"object"`
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
// [ConversationItemNewResponseDataMcpListTools],
// [ConversationItemNewResponseDataReasoning],
// [ConversationItemNewResponseDataCompaction].
//
// Use the [ConversationItemNewResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataUnion struct {
	// This field is a union of [ConversationItemNewResponseDataMessageContentUnion],
	// [[]ConversationItemNewResponseDataReasoningContent]
	Content ConversationItemNewResponseDataUnionContent `json:"content"`
	// This field is from variant [ConversationItemNewResponseDataMessage].
	Role   ConversationItemNewResponseDataMessageRole `json:"role"`
	ID     string                                     `json:"id"`
	Status string                                     `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools", "reasoning", "compaction".
	Type string `json:"type"`
	// This field is from variant [ConversationItemNewResponseDataWebSearchCall].
	Action ConversationItemNewResponseDataWebSearchCallActionUnion `json:"action"`
	// This field is from variant [ConversationItemNewResponseDataFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemNewResponseDataFileSearchCall].
	Results   []ConversationItemNewResponseDataFileSearchCallResult `json:"results"`
	Arguments string                                                `json:"arguments"`
	CallID    string                                                `json:"call_id"`
	Name      string                                                `json:"name"`
	// This field is a union of
	// [ConversationItemNewResponseDataFunctionCallOutputOutputUnion], [string]
	Output      ConversationItemNewResponseDataUnionOutput `json:"output"`
	ServerLabel string                                     `json:"server_label"`
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
	// This field is from variant [ConversationItemNewResponseDataReasoning].
	Summary []ConversationItemNewResponseDataReasoningSummary `json:"summary"`
	// This field is from variant [ConversationItemNewResponseDataCompaction].
	EncryptedContent string `json:"encrypted_content"`
	JSON             struct {
		Content           respjson.Field
		Role              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		Action            respjson.Field
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
		Summary           respjson.Field
		EncryptedContent  respjson.Field
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
func (ConversationItemNewResponseDataReasoning) implConversationItemNewResponseDataUnion()    {}
func (ConversationItemNewResponseDataCompaction) implConversationItemNewResponseDataUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataUnion.AsAny().(type) {
//	case ogxclient.ConversationItemNewResponseDataMessage:
//	case ogxclient.ConversationItemNewResponseDataWebSearchCall:
//	case ogxclient.ConversationItemNewResponseDataFileSearchCall:
//	case ogxclient.ConversationItemNewResponseDataFunctionCall:
//	case ogxclient.ConversationItemNewResponseDataFunctionCallOutput:
//	case ogxclient.ConversationItemNewResponseDataMcpApprovalRequest:
//	case ogxclient.ConversationItemNewResponseDataMcpApprovalResponse:
//	case ogxclient.ConversationItemNewResponseDataMcpCall:
//	case ogxclient.ConversationItemNewResponseDataMcpListTools:
//	case ogxclient.ConversationItemNewResponseDataReasoning:
//	case ogxclient.ConversationItemNewResponseDataCompaction:
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
	case "reasoning":
		return u.AsReasoning()
	case "compaction":
		return u.AsCompaction()
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

func (u ConversationItemNewResponseDataUnion) AsReasoning() (v ConversationItemNewResponseDataReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsCompaction() (v ConversationItemNewResponseDataCompaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataUnionContent is an implicit subunion of
// [ConversationItemNewResponseDataUnion].
// ConversationItemNewResponseDataUnionContent provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationItemNewResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal
// OfConversationItemNewResponseDataReasoningContentArray]
type ConversationItemNewResponseDataUnionContent struct {
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
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataReasoningContent] instead of an object.
	OfConversationItemNewResponseDataReasoningContentArray []ConversationItemNewResponseDataReasoningContent `json:",inline"`
	JSON                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		OfConversationItemNewResponseDataReasoningContentArray                                                                 respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ConversationItemNewResponseDataUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataUnionOutput is an implicit subunion of
// [ConversationItemNewResponseDataUnion].
// ConversationItemNewResponseDataUnionOutput provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationItemNewResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ConversationItemNewResponseDataUnionOutput struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ConversationItemNewResponseDataUnionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemNewResponseDataMessage struct {
	Content ConversationItemNewResponseDataMessageContentUnion `json:"content" api:"required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemNewResponseDataMessageRole `json:"role" api:"required"`
	ID     string                                     `json:"id" api:"nullable"`
	Status string                                     `json:"status" api:"nullable"`
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
	Detail ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                                `json:"file_id"`
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
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                                `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                                `json:"image_url" api:"nullable"`
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

type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ConversationItemNewResponseDataMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

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
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
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

// Text content within an output message of an OpenAI response.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                             `json:"text" api:"required"`
	Annotations []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs" api:"nullable"`
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
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case ogxclient.ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Container file citation annotation referencing a file within a container.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File path annotation referencing a generated file in response content.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
	// The top log probabilities for the token.
	TopLogprobs []ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs" api:"nullable"`
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
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageRole string

const (
	ConversationItemNewResponseDataMessageRoleSystem    ConversationItemNewResponseDataMessageRole = "system"
	ConversationItemNewResponseDataMessageRoleDeveloper ConversationItemNewResponseDataMessageRole = "developer"
	ConversationItemNewResponseDataMessageRoleUser      ConversationItemNewResponseDataMessageRole = "user"
	ConversationItemNewResponseDataMessageRoleAssistant ConversationItemNewResponseDataMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataWebSearchCall struct {
	ID     string `json:"id" api:"required"`
	Status string `json:"status" api:"required"`
	// Web search action: performs a search query.
	Action ConversationItemNewResponseDataWebSearchCallActionUnion `json:"action" api:"nullable"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Action      respjson.Field
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

// ConversationItemNewResponseDataWebSearchCallActionUnion contains all possible
// properties and values from
// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch],
// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionOpenPage],
// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionFind].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataWebSearchCallActionUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch].
	Query string `json:"query"`
	// This field is from variant
	// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch].
	Sources []ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearchSource `json:"sources"`
	Type    string                                                                          `json:"type"`
	URL     string                                                                          `json:"url"`
	// This field is from variant
	// [ConversationItemNewResponseDataWebSearchCallActionWebSearchActionFind].
	Pattern string `json:"pattern"`
	JSON    struct {
		Query   respjson.Field
		Queries respjson.Field
		Sources respjson.Field
		Type    respjson.Field
		URL     respjson.Field
		Pattern respjson.Field
		raw     string
	} `json:"-"`
}

func (u ConversationItemNewResponseDataWebSearchCallActionUnion) AsWebSearchActionSearch() (v ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataWebSearchCallActionUnion) AsWebSearchActionOpenPage() (v ConversationItemNewResponseDataWebSearchCallActionWebSearchActionOpenPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataWebSearchCallActionUnion) AsWebSearchActionFind() (v ConversationItemNewResponseDataWebSearchCallActionWebSearchActionFind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataWebSearchCallActionUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataWebSearchCallActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: performs a search query.
type ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch struct {
	Query   string                                                                          `json:"query" api:"required"`
	Queries []string                                                                        `json:"queries" api:"nullable"`
	Sources []ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearchSource `json:"sources" api:"nullable"`
	// Any of "search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		Queries     respjson.Field
		Sources     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source URL returned by a web search action.
type ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearchSource struct {
	URL string `json:"url" api:"required"`
	// Any of "url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearchSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataWebSearchCallActionWebSearchActionSearchSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: opens a specific URL from search results.
type ConversationItemNewResponseDataWebSearchCallActionWebSearchActionOpenPage struct {
	// Any of "open_page".
	Type string `json:"type"`
	URL  string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataWebSearchCallActionWebSearchActionOpenPage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataWebSearchCallActionWebSearchActionOpenPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: searches for a pattern within a loaded page.
type ConversationItemNewResponseDataWebSearchCallActionWebSearchActionFind struct {
	Pattern string `json:"pattern" api:"required"`
	URL     string `json:"url" api:"required"`
	// Any of "find_in_page".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pattern     respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataWebSearchCallActionWebSearchActionFind) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataWebSearchCallActionWebSearchActionFind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFileSearchCall struct {
	ID      string                                                `json:"id" api:"required"`
	Queries []string                                              `json:"queries" api:"required"`
	Status  string                                                `json:"status" api:"required"`
	Results []ConversationItemNewResponseDataFileSearchCallResult `json:"results" api:"nullable"`
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
func (r ConversationItemNewResponseDataFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFunctionCall struct {
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
func (r ConversationItemNewResponseDataFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ConversationItemNewResponseDataFunctionCallOutput struct {
	CallID string                                                       `json:"call_id" api:"required"`
	Output ConversationItemNewResponseDataFunctionCallOutputOutputUnion `json:"output" api:"required"`
	ID     string                                                       `json:"id" api:"nullable"`
	Status string                                                       `json:"status" api:"nullable"`
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

// ConversationItemNewResponseDataFunctionCallOutputOutputUnion contains all
// possible properties and values from [string],
// [[]ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ConversationItemNewResponseDataFunctionCallOutputOutputUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemNewResponseDataFunctionCallOutputOutputUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFunctionCallOutputOutputUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataFunctionCallOutputOutputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataFunctionCallOutputOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                                          `json:"file_id"`
	// This field is from variant
	// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case ogxclient.ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case ogxclient.ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case ogxclient.ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                                          `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                                          `json:"image_url" api:"nullable"`
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
func (r ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ConversationItemNewResponseDataFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

// A request for human approval of a tool invocation.
type ConversationItemNewResponseDataMcpApprovalRequest struct {
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
func (r ConversationItemNewResponseDataMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ConversationItemNewResponseDataMcpApprovalResponse struct {
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
func (r ConversationItemNewResponseDataMcpApprovalResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemNewResponseDataMcpCall struct {
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
func (r ConversationItemNewResponseDataMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemNewResponseDataMcpListTools struct {
	ID          string                                            `json:"id" api:"required"`
	ServerLabel string                                            `json:"server_label" api:"required"`
	Tools       []ConversationItemNewResponseDataMcpListToolsTool `json:"tools" api:"required"`
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
func (r ConversationItemNewResponseDataMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning output from the model, representing the model's thinking process.
type ConversationItemNewResponseDataReasoning struct {
	// Unique identifier for the reasoning output item.
	ID string `json:"id" api:"required"`
	// Summary of the reasoning output.
	Summary []ConversationItemNewResponseDataReasoningSummary `json:"summary" api:"required"`
	// The reasoning content from the model.
	Content []ConversationItemNewResponseDataReasoningContent `json:"content" api:"nullable"`
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
func (r ConversationItemNewResponseDataReasoning) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summary of reasoning output from the model.
type ConversationItemNewResponseDataReasoningSummary struct {
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
func (r ConversationItemNewResponseDataReasoningSummary) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataReasoningSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text from the model.
type ConversationItemNewResponseDataReasoningContent struct {
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
func (r ConversationItemNewResponseDataReasoningContent) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataReasoningContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A compaction item that summarizes prior conversation context.
type ConversationItemNewResponseDataCompaction struct {
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
func (r ConversationItemNewResponseDataCompaction) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataCompaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataRole string

const (
	ConversationItemNewResponseDataRoleSystem    ConversationItemNewResponseDataRole = "system"
	ConversationItemNewResponseDataRoleDeveloper ConversationItemNewResponseDataRole = "developer"
	ConversationItemNewResponseDataRoleUser      ConversationItemNewResponseDataRole = "user"
	ConversationItemNewResponseDataRoleAssistant ConversationItemNewResponseDataRole = "assistant"
)

// The type of object returned, must be list.
type ConversationItemNewResponseObject string

const (
	ConversationItemNewResponseObjectList ConversationItemNewResponseObject = "list"
)

// ConversationItemListResponseUnion contains all possible properties and values
// from [ConversationItemListResponseMessage],
// [ConversationItemListResponseWebSearchCall],
// [ConversationItemListResponseFileSearchCall],
// [ConversationItemListResponseFunctionCall],
// [ConversationItemListResponseFunctionCallOutput],
// [ConversationItemListResponseMcpApprovalRequest],
// [ConversationItemListResponseMcpApprovalResponse],
// [ConversationItemListResponseMcpCall],
// [ConversationItemListResponseMcpListTools],
// [ConversationItemListResponseReasoning],
// [ConversationItemListResponseCompaction].
//
// Use the [ConversationItemListResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseUnion struct {
	// This field is a union of [ConversationItemListResponseMessageContentUnion],
	// [[]ConversationItemListResponseReasoningContent]
	Content ConversationItemListResponseUnionContent `json:"content"`
	// This field is from variant [ConversationItemListResponseMessage].
	Role   ConversationItemListResponseMessageRole `json:"role"`
	ID     string                                  `json:"id"`
	Status string                                  `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools", "reasoning", "compaction".
	Type string `json:"type"`
	// This field is from variant [ConversationItemListResponseWebSearchCall].
	Action ConversationItemListResponseWebSearchCallActionUnion `json:"action"`
	// This field is from variant [ConversationItemListResponseFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemListResponseFileSearchCall].
	Results   []ConversationItemListResponseFileSearchCallResult `json:"results"`
	Arguments string                                             `json:"arguments"`
	CallID    string                                             `json:"call_id"`
	Name      string                                             `json:"name"`
	// This field is a union of
	// [ConversationItemListResponseFunctionCallOutputOutputUnion], [string]
	Output      ConversationItemListResponseUnionOutput `json:"output"`
	ServerLabel string                                  `json:"server_label"`
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
	// This field is from variant [ConversationItemListResponseReasoning].
	Summary []ConversationItemListResponseReasoningSummary `json:"summary"`
	// This field is from variant [ConversationItemListResponseCompaction].
	EncryptedContent string `json:"encrypted_content"`
	JSON             struct {
		Content           respjson.Field
		Role              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		Action            respjson.Field
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
		Summary           respjson.Field
		EncryptedContent  respjson.Field
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
func (ConversationItemListResponseReasoning) implConversationItemListResponseUnion()           {}
func (ConversationItemListResponseCompaction) implConversationItemListResponseUnion()          {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseUnion.AsAny().(type) {
//	case ogxclient.ConversationItemListResponseMessage:
//	case ogxclient.ConversationItemListResponseWebSearchCall:
//	case ogxclient.ConversationItemListResponseFileSearchCall:
//	case ogxclient.ConversationItemListResponseFunctionCall:
//	case ogxclient.ConversationItemListResponseFunctionCallOutput:
//	case ogxclient.ConversationItemListResponseMcpApprovalRequest:
//	case ogxclient.ConversationItemListResponseMcpApprovalResponse:
//	case ogxclient.ConversationItemListResponseMcpCall:
//	case ogxclient.ConversationItemListResponseMcpListTools:
//	case ogxclient.ConversationItemListResponseReasoning:
//	case ogxclient.ConversationItemListResponseCompaction:
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
	case "reasoning":
		return u.AsReasoning()
	case "compaction":
		return u.AsCompaction()
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

func (u ConversationItemListResponseUnion) AsReasoning() (v ConversationItemListResponseReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseUnion) AsCompaction() (v ConversationItemListResponseCompaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseUnionContent is an implicit subunion of
// [ConversationItemListResponseUnion]. ConversationItemListResponseUnionContent
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationItemListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal
// OfConversationItemListResponseReasoningContentArray]
type ConversationItemListResponseUnionContent struct {
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
	// This field will be present if the value is a
	// [[]ConversationItemListResponseReasoningContent] instead of an object.
	OfConversationItemListResponseReasoningContentArray []ConversationItemListResponseReasoningContent `json:",inline"`
	JSON                                                struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		OfConversationItemListResponseReasoningContentArray                                                                    respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ConversationItemListResponseUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseUnionOutput is an implicit subunion of
// [ConversationItemListResponseUnion]. ConversationItemListResponseUnionOutput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationItemListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ConversationItemListResponseUnionOutput struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ConversationItemListResponseUnionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemListResponseMessage struct {
	Content ConversationItemListResponseMessageContentUnion `json:"content" api:"required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemListResponseMessageRole `json:"role" api:"required"`
	ID     string                                  `json:"id" api:"nullable"`
	Status string                                  `json:"status" api:"nullable"`
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
	Detail ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                             `json:"file_id"`
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
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                             `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                             `json:"image_url" api:"nullable"`
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

type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ConversationItemListResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

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
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
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

// Text content within an output message of an OpenAI response.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                          `json:"text" api:"required"`
	Annotations []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs" api:"nullable"`
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
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case ogxclient.ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Container file citation annotation referencing a file within a container.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File path annotation referencing a generated file in response content.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
	// The top log probabilities for the token.
	TopLogprobs []ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs" api:"nullable"`
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
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseMessageRole string

const (
	ConversationItemListResponseMessageRoleSystem    ConversationItemListResponseMessageRole = "system"
	ConversationItemListResponseMessageRoleDeveloper ConversationItemListResponseMessageRole = "developer"
	ConversationItemListResponseMessageRoleUser      ConversationItemListResponseMessageRole = "user"
	ConversationItemListResponseMessageRoleAssistant ConversationItemListResponseMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ConversationItemListResponseWebSearchCall struct {
	ID     string `json:"id" api:"required"`
	Status string `json:"status" api:"required"`
	// Web search action: performs a search query.
	Action ConversationItemListResponseWebSearchCallActionUnion `json:"action" api:"nullable"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Action      respjson.Field
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

// ConversationItemListResponseWebSearchCallActionUnion contains all possible
// properties and values from
// [ConversationItemListResponseWebSearchCallActionWebSearchActionSearch],
// [ConversationItemListResponseWebSearchCallActionWebSearchActionOpenPage],
// [ConversationItemListResponseWebSearchCallActionWebSearchActionFind].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseWebSearchCallActionUnion struct {
	// This field is from variant
	// [ConversationItemListResponseWebSearchCallActionWebSearchActionSearch].
	Query string `json:"query"`
	// This field is from variant
	// [ConversationItemListResponseWebSearchCallActionWebSearchActionSearch].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ConversationItemListResponseWebSearchCallActionWebSearchActionSearch].
	Sources []ConversationItemListResponseWebSearchCallActionWebSearchActionSearchSource `json:"sources"`
	Type    string                                                                       `json:"type"`
	URL     string                                                                       `json:"url"`
	// This field is from variant
	// [ConversationItemListResponseWebSearchCallActionWebSearchActionFind].
	Pattern string `json:"pattern"`
	JSON    struct {
		Query   respjson.Field
		Queries respjson.Field
		Sources respjson.Field
		Type    respjson.Field
		URL     respjson.Field
		Pattern respjson.Field
		raw     string
	} `json:"-"`
}

func (u ConversationItemListResponseWebSearchCallActionUnion) AsWebSearchActionSearch() (v ConversationItemListResponseWebSearchCallActionWebSearchActionSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseWebSearchCallActionUnion) AsWebSearchActionOpenPage() (v ConversationItemListResponseWebSearchCallActionWebSearchActionOpenPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseWebSearchCallActionUnion) AsWebSearchActionFind() (v ConversationItemListResponseWebSearchCallActionWebSearchActionFind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseWebSearchCallActionUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseWebSearchCallActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: performs a search query.
type ConversationItemListResponseWebSearchCallActionWebSearchActionSearch struct {
	Query   string                                                                       `json:"query" api:"required"`
	Queries []string                                                                     `json:"queries" api:"nullable"`
	Sources []ConversationItemListResponseWebSearchCallActionWebSearchActionSearchSource `json:"sources" api:"nullable"`
	// Any of "search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		Queries     respjson.Field
		Sources     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseWebSearchCallActionWebSearchActionSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseWebSearchCallActionWebSearchActionSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source URL returned by a web search action.
type ConversationItemListResponseWebSearchCallActionWebSearchActionSearchSource struct {
	URL string `json:"url" api:"required"`
	// Any of "url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseWebSearchCallActionWebSearchActionSearchSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseWebSearchCallActionWebSearchActionSearchSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: opens a specific URL from search results.
type ConversationItemListResponseWebSearchCallActionWebSearchActionOpenPage struct {
	// Any of "open_page".
	Type string `json:"type"`
	URL  string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseWebSearchCallActionWebSearchActionOpenPage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseWebSearchCallActionWebSearchActionOpenPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: searches for a pattern within a loaded page.
type ConversationItemListResponseWebSearchCallActionWebSearchActionFind struct {
	Pattern string `json:"pattern" api:"required"`
	URL     string `json:"url" api:"required"`
	// Any of "find_in_page".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pattern     respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseWebSearchCallActionWebSearchActionFind) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseWebSearchCallActionWebSearchActionFind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemListResponseFileSearchCall struct {
	ID      string                                             `json:"id" api:"required"`
	Queries []string                                           `json:"queries" api:"required"`
	Status  string                                             `json:"status" api:"required"`
	Results []ConversationItemListResponseFileSearchCallResult `json:"results" api:"nullable"`
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
func (r ConversationItemListResponseFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemListResponseFunctionCall struct {
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
func (r ConversationItemListResponseFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ConversationItemListResponseFunctionCallOutput struct {
	CallID string                                                    `json:"call_id" api:"required"`
	Output ConversationItemListResponseFunctionCallOutputOutputUnion `json:"output" api:"required"`
	ID     string                                                    `json:"id" api:"nullable"`
	Status string                                                    `json:"status" api:"nullable"`
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

// ConversationItemListResponseFunctionCallOutputOutputUnion contains all possible
// properties and values from [string],
// [[]ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ConversationItemListResponseFunctionCallOutputOutputUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemListResponseFunctionCallOutputOutputUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseFunctionCallOutputOutputUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseFunctionCallOutputOutputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseFunctionCallOutputOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                                       `json:"file_id"`
	// This field is from variant
	// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case ogxclient.ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case ogxclient.ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case ogxclient.ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                                       `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                                       `json:"image_url" api:"nullable"`
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
func (r ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ConversationItemListResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

// A request for human approval of a tool invocation.
type ConversationItemListResponseMcpApprovalRequest struct {
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
func (r ConversationItemListResponseMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ConversationItemListResponseMcpApprovalResponse struct {
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
func (r ConversationItemListResponseMcpApprovalResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemListResponseMcpCall struct {
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
func (r ConversationItemListResponseMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemListResponseMcpListTools struct {
	ID          string                                         `json:"id" api:"required"`
	ServerLabel string                                         `json:"server_label" api:"required"`
	Tools       []ConversationItemListResponseMcpListToolsTool `json:"tools" api:"required"`
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
func (r ConversationItemListResponseMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning output from the model, representing the model's thinking process.
type ConversationItemListResponseReasoning struct {
	// Unique identifier for the reasoning output item.
	ID string `json:"id" api:"required"`
	// Summary of the reasoning output.
	Summary []ConversationItemListResponseReasoningSummary `json:"summary" api:"required"`
	// The reasoning content from the model.
	Content []ConversationItemListResponseReasoningContent `json:"content" api:"nullable"`
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
func (r ConversationItemListResponseReasoning) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summary of reasoning output from the model.
type ConversationItemListResponseReasoningSummary struct {
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
func (r ConversationItemListResponseReasoningSummary) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseReasoningSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text from the model.
type ConversationItemListResponseReasoningContent struct {
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
func (r ConversationItemListResponseReasoningContent) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseReasoningContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A compaction item that summarizes prior conversation context.
type ConversationItemListResponseCompaction struct {
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
func (r ConversationItemListResponseCompaction) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseCompaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseRole string

const (
	ConversationItemListResponseRoleSystem    ConversationItemListResponseRole = "system"
	ConversationItemListResponseRoleDeveloper ConversationItemListResponseRole = "developer"
	ConversationItemListResponseRoleUser      ConversationItemListResponseRole = "user"
	ConversationItemListResponseRoleAssistant ConversationItemListResponseRole = "assistant"
)

// ConversationItemGetResponseUnion contains all possible properties and values
// from [ConversationItemGetResponseMessage],
// [ConversationItemGetResponseWebSearchCall],
// [ConversationItemGetResponseFileSearchCall],
// [ConversationItemGetResponseFunctionCall],
// [ConversationItemGetResponseFunctionCallOutput],
// [ConversationItemGetResponseMcpApprovalRequest],
// [ConversationItemGetResponseMcpApprovalResponse],
// [ConversationItemGetResponseMcpCall], [ConversationItemGetResponseMcpListTools],
// [ConversationItemGetResponseReasoning], [ConversationItemGetResponseCompaction].
//
// Use the [ConversationItemGetResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseUnion struct {
	// This field is a union of [ConversationItemGetResponseMessageContentUnion],
	// [[]ConversationItemGetResponseReasoningContent]
	Content ConversationItemGetResponseUnionContent `json:"content"`
	// This field is from variant [ConversationItemGetResponseMessage].
	Role   ConversationItemGetResponseMessageRole `json:"role"`
	ID     string                                 `json:"id"`
	Status string                                 `json:"status"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools", "reasoning", "compaction".
	Type string `json:"type"`
	// This field is from variant [ConversationItemGetResponseWebSearchCall].
	Action ConversationItemGetResponseWebSearchCallActionUnion `json:"action"`
	// This field is from variant [ConversationItemGetResponseFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemGetResponseFileSearchCall].
	Results   []ConversationItemGetResponseFileSearchCallResult `json:"results"`
	Arguments string                                            `json:"arguments"`
	CallID    string                                            `json:"call_id"`
	Name      string                                            `json:"name"`
	// This field is a union of
	// [ConversationItemGetResponseFunctionCallOutputOutputUnion], [string]
	Output      ConversationItemGetResponseUnionOutput `json:"output"`
	ServerLabel string                                 `json:"server_label"`
	// This field is from variant [ConversationItemGetResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant [ConversationItemGetResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant [ConversationItemGetResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	// This field is from variant [ConversationItemGetResponseMcpCall].
	Error string `json:"error"`
	// This field is from variant [ConversationItemGetResponseMcpListTools].
	Tools []ConversationItemGetResponseMcpListToolsTool `json:"tools"`
	// This field is from variant [ConversationItemGetResponseReasoning].
	Summary []ConversationItemGetResponseReasoningSummary `json:"summary"`
	// This field is from variant [ConversationItemGetResponseCompaction].
	EncryptedContent string `json:"encrypted_content"`
	JSON             struct {
		Content           respjson.Field
		Role              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		Action            respjson.Field
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
		Summary           respjson.Field
		EncryptedContent  respjson.Field
		raw               string
	} `json:"-"`
}

// anyConversationItemGetResponse is implemented by each variant of
// [ConversationItemGetResponseUnion] to add type safety for the return type of
// [ConversationItemGetResponseUnion.AsAny]
type anyConversationItemGetResponse interface {
	implConversationItemGetResponseUnion()
}

func (ConversationItemGetResponseMessage) implConversationItemGetResponseUnion()             {}
func (ConversationItemGetResponseWebSearchCall) implConversationItemGetResponseUnion()       {}
func (ConversationItemGetResponseFileSearchCall) implConversationItemGetResponseUnion()      {}
func (ConversationItemGetResponseFunctionCall) implConversationItemGetResponseUnion()        {}
func (ConversationItemGetResponseFunctionCallOutput) implConversationItemGetResponseUnion()  {}
func (ConversationItemGetResponseMcpApprovalRequest) implConversationItemGetResponseUnion()  {}
func (ConversationItemGetResponseMcpApprovalResponse) implConversationItemGetResponseUnion() {}
func (ConversationItemGetResponseMcpCall) implConversationItemGetResponseUnion()             {}
func (ConversationItemGetResponseMcpListTools) implConversationItemGetResponseUnion()        {}
func (ConversationItemGetResponseReasoning) implConversationItemGetResponseUnion()           {}
func (ConversationItemGetResponseCompaction) implConversationItemGetResponseUnion()          {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseUnion.AsAny().(type) {
//	case ogxclient.ConversationItemGetResponseMessage:
//	case ogxclient.ConversationItemGetResponseWebSearchCall:
//	case ogxclient.ConversationItemGetResponseFileSearchCall:
//	case ogxclient.ConversationItemGetResponseFunctionCall:
//	case ogxclient.ConversationItemGetResponseFunctionCallOutput:
//	case ogxclient.ConversationItemGetResponseMcpApprovalRequest:
//	case ogxclient.ConversationItemGetResponseMcpApprovalResponse:
//	case ogxclient.ConversationItemGetResponseMcpCall:
//	case ogxclient.ConversationItemGetResponseMcpListTools:
//	case ogxclient.ConversationItemGetResponseReasoning:
//	case ogxclient.ConversationItemGetResponseCompaction:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseUnion) AsAny() anyConversationItemGetResponse {
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
	case "reasoning":
		return u.AsReasoning()
	case "compaction":
		return u.AsCompaction()
	}
	return nil
}

func (u ConversationItemGetResponseUnion) AsMessage() (v ConversationItemGetResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsWebSearchCall() (v ConversationItemGetResponseWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsFileSearchCall() (v ConversationItemGetResponseFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsFunctionCall() (v ConversationItemGetResponseFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsFunctionCallOutput() (v ConversationItemGetResponseFunctionCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsMcpApprovalRequest() (v ConversationItemGetResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsMcpApprovalResponse() (v ConversationItemGetResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsMcpCall() (v ConversationItemGetResponseMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsMcpListTools() (v ConversationItemGetResponseMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsReasoning() (v ConversationItemGetResponseReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsCompaction() (v ConversationItemGetResponseCompaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseUnionContent is an implicit subunion of
// [ConversationItemGetResponseUnion]. ConversationItemGetResponseUnionContent
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationItemGetResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal
// OfConversationItemGetResponseReasoningContentArray]
type ConversationItemGetResponseUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseReasoningContent] instead of an object.
	OfConversationItemGetResponseReasoningContentArray []ConversationItemGetResponseReasoningContent `json:",inline"`
	JSON                                               struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		OfConversationItemGetResponseReasoningContentArray                                                                     respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ConversationItemGetResponseUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseUnionOutput is an implicit subunion of
// [ConversationItemGetResponseUnion]. ConversationItemGetResponseUnionOutput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationItemGetResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ConversationItemGetResponseUnionOutput struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (r *ConversationItemGetResponseUnionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemGetResponseMessage struct {
	Content ConversationItemGetResponseMessageContentUnion `json:"content" api:"required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemGetResponseMessageRole `json:"role" api:"required"`
	ID     string                                 `json:"id" api:"nullable"`
	Status string                                 `json:"status" api:"nullable"`
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
func (r ConversationItemGetResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentUnion contains all possible properties
// and values from [string],
// [[]ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [[]ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
// OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal]
type ConversationItemGetResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
	// instead of an object.
	OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		OfListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal                               respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemGetResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentUnion) AsListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusal() (v []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                            `json:"file_id"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                            `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                            `json:"image_url" api:"nullable"`
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ConversationItemGetResponseMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

// ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion
// contains all possible properties and values from
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText],
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
//
// Use the
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Text string `json:"text"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Annotations []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText].
	Logprobs []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob `json:"logprobs"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal].
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

// anyConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem
// is implemented by each variant of
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny]
type anyConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem interface {
	implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion()
}

func (ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}
func (ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion.AsAny().(type) {
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText:
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsAny() anyConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItem {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	}
	return nil
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsOutputText() (v ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) AsRefusal() (v ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content within an output message of an OpenAI response.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                         `json:"text" api:"required"`
	Annotations []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations"`
	Logprobs    []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs" api:"nullable"`
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion
// contains all possible properties and values from
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation],
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation],
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation],
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath].
//
// Use the
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation].
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

// anyConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation
// is implemented by each variant of
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny]
type anyConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation interface {
	implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion()
}

func (ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}
func (ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) implConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion.AsAny().(type) {
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation:
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation:
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation:
//	case ogxclient.ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsAny() anyConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotation {
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

func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFileCitation() (v ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsURLCitation() (v ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsContainerFileCitation() (v ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) AsFilePath() (v ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Container file citation annotation referencing a file within a container.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File path annotation referencing a generated file in response content.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
	// The top log probabilities for the token.
	TopLogprobs []ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs" api:"nullable"`
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal struct {
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
func (r ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentListOpenAIResponseOutputMessageContentOutputTextOutputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseMessageRole string

const (
	ConversationItemGetResponseMessageRoleSystem    ConversationItemGetResponseMessageRole = "system"
	ConversationItemGetResponseMessageRoleDeveloper ConversationItemGetResponseMessageRole = "developer"
	ConversationItemGetResponseMessageRoleUser      ConversationItemGetResponseMessageRole = "user"
	ConversationItemGetResponseMessageRoleAssistant ConversationItemGetResponseMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ConversationItemGetResponseWebSearchCall struct {
	ID     string `json:"id" api:"required"`
	Status string `json:"status" api:"required"`
	// Web search action: performs a search query.
	Action ConversationItemGetResponseWebSearchCallActionUnion `json:"action" api:"nullable"`
	// Any of "web_search_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Action      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseWebSearchCallActionUnion contains all possible
// properties and values from
// [ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch],
// [ConversationItemGetResponseWebSearchCallActionWebSearchActionOpenPage],
// [ConversationItemGetResponseWebSearchCallActionWebSearchActionFind].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseWebSearchCallActionUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch].
	Query string `json:"query"`
	// This field is from variant
	// [ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch].
	Sources []ConversationItemGetResponseWebSearchCallActionWebSearchActionSearchSource `json:"sources"`
	Type    string                                                                      `json:"type"`
	URL     string                                                                      `json:"url"`
	// This field is from variant
	// [ConversationItemGetResponseWebSearchCallActionWebSearchActionFind].
	Pattern string `json:"pattern"`
	JSON    struct {
		Query   respjson.Field
		Queries respjson.Field
		Sources respjson.Field
		Type    respjson.Field
		URL     respjson.Field
		Pattern respjson.Field
		raw     string
	} `json:"-"`
}

func (u ConversationItemGetResponseWebSearchCallActionUnion) AsWebSearchActionSearch() (v ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseWebSearchCallActionUnion) AsWebSearchActionOpenPage() (v ConversationItemGetResponseWebSearchCallActionWebSearchActionOpenPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseWebSearchCallActionUnion) AsWebSearchActionFind() (v ConversationItemGetResponseWebSearchCallActionWebSearchActionFind) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseWebSearchCallActionUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseWebSearchCallActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: performs a search query.
type ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch struct {
	Query   string                                                                      `json:"query" api:"required"`
	Queries []string                                                                    `json:"queries" api:"nullable"`
	Sources []ConversationItemGetResponseWebSearchCallActionWebSearchActionSearchSource `json:"sources" api:"nullable"`
	// Any of "search".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		Queries     respjson.Field
		Sources     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseWebSearchCallActionWebSearchActionSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source URL returned by a web search action.
type ConversationItemGetResponseWebSearchCallActionWebSearchActionSearchSource struct {
	URL string `json:"url" api:"required"`
	// Any of "url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseWebSearchCallActionWebSearchActionSearchSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseWebSearchCallActionWebSearchActionSearchSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: opens a specific URL from search results.
type ConversationItemGetResponseWebSearchCallActionWebSearchActionOpenPage struct {
	// Any of "open_page".
	Type string `json:"type"`
	URL  string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseWebSearchCallActionWebSearchActionOpenPage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseWebSearchCallActionWebSearchActionOpenPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search action: searches for a pattern within a loaded page.
type ConversationItemGetResponseWebSearchCallActionWebSearchActionFind struct {
	Pattern string `json:"pattern" api:"required"`
	URL     string `json:"url" api:"required"`
	// Any of "find_in_page".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pattern     respjson.Field
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseWebSearchCallActionWebSearchActionFind) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseWebSearchCallActionWebSearchActionFind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemGetResponseFileSearchCall struct {
	ID      string                                            `json:"id" api:"required"`
	Queries []string                                          `json:"queries" api:"required"`
	Status  string                                            `json:"status" api:"required"`
	Results []ConversationItemGetResponseFileSearchCallResult `json:"results" api:"nullable"`
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
func (r ConversationItemGetResponseFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ConversationItemGetResponseFileSearchCallResult struct {
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
func (r ConversationItemGetResponseFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemGetResponseFunctionCall struct {
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
func (r ConversationItemGetResponseFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ConversationItemGetResponseFunctionCallOutput struct {
	CallID string                                                   `json:"call_id" api:"required"`
	Output ConversationItemGetResponseFunctionCallOutputOutputUnion `json:"output" api:"required"`
	ID     string                                                   `json:"id" api:"nullable"`
	Status string                                                   `json:"status" api:"nullable"`
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
func (r ConversationItemGetResponseFunctionCallOutput) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseFunctionCallOutputOutputUnion contains all possible
// properties and values from [string],
// [[]ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile]
type ConversationItemGetResponseFunctionCallOutputOutputUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
	// instead of an object.
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u ConversationItemGetResponseFunctionCallOutputOutputUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFunctionCallOutputOutputUnion) AsListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile() (v []ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseFunctionCallOutputOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseFunctionCallOutputOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion
// contains all possible properties and values from
// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText],
// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage],
// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
//
// Use the
// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	Detail ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID string                                                                                                                                                                                      `json:"file_id"`
	// This field is from variant
	// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile].
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

// anyConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem
// is implemented by each variant of
// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
// to add type safety for the return type of
// [ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny]
type anyConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem interface {
	implConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion()
}

func (ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) implConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) implConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}
func (ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) implConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion.AsAny().(type) {
//	case ogxclient.ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText:
//	case ogxclient.ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage:
//	case ogxclient.ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsAny() anyConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItem {
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

func (u ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputText() (v ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputImage() (v ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) AsInputFile() (v ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
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
func (r ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail"`
	FileID   string                                                                                                                                                                                      `json:"file_id" api:"nullable"`
	ImageURL string                                                                                                                                                                                      `json:"image_url" api:"nullable"`
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
func (r ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
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
func (r ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail string

const (
	ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailLow  ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "low"
	ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailHigh ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "high"
	ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetailAuto ConversationItemGetResponseFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemDetail = "auto"
)

// A request for human approval of a tool invocation.
type ConversationItemGetResponseMcpApprovalRequest struct {
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
func (r ConversationItemGetResponseMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ConversationItemGetResponseMcpApprovalResponse struct {
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
func (r ConversationItemGetResponseMcpApprovalResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemGetResponseMcpCall struct {
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
func (r ConversationItemGetResponseMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemGetResponseMcpListTools struct {
	ID          string                                        `json:"id" api:"required"`
	ServerLabel string                                        `json:"server_label" api:"required"`
	Tools       []ConversationItemGetResponseMcpListToolsTool `json:"tools" api:"required"`
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
func (r ConversationItemGetResponseMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ConversationItemGetResponseMcpListToolsTool struct {
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
func (r ConversationItemGetResponseMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning output from the model, representing the model's thinking process.
type ConversationItemGetResponseReasoning struct {
	// Unique identifier for the reasoning output item.
	ID string `json:"id" api:"required"`
	// Summary of the reasoning output.
	Summary []ConversationItemGetResponseReasoningSummary `json:"summary" api:"required"`
	// The reasoning content from the model.
	Content []ConversationItemGetResponseReasoningContent `json:"content" api:"nullable"`
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
func (r ConversationItemGetResponseReasoning) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summary of reasoning output from the model.
type ConversationItemGetResponseReasoningSummary struct {
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
func (r ConversationItemGetResponseReasoningSummary) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseReasoningSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text from the model.
type ConversationItemGetResponseReasoningContent struct {
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
func (r ConversationItemGetResponseReasoningContent) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseReasoningContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A compaction item that summarizes prior conversation context.
type ConversationItemGetResponseCompaction struct {
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
func (r ConversationItemGetResponseCompaction) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseCompaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseRole string

const (
	ConversationItemGetResponseRoleSystem    ConversationItemGetResponseRole = "system"
	ConversationItemGetResponseRoleDeveloper ConversationItemGetResponseRole = "developer"
	ConversationItemGetResponseRoleUser      ConversationItemGetResponseRole = "user"
	ConversationItemGetResponseRoleAssistant ConversationItemGetResponseRole = "assistant"
)

type ConversationItemNewParams struct {
	// Items to include in the conversation context. You may add up to 20 items at a
	// time.
	Items []ConversationItemNewParamsItemUnion `json:"items,omitzero" api:"required"`
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
	OfReasoning           *ConversationItemNewParamsItemReasoning           `json:",omitzero,inline"`
	OfCompaction          *ConversationItemNewParamsItemCompaction          `json:",omitzero,inline"`
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
		u.OfMcpListTools,
		u.OfReasoning,
		u.OfCompaction)
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
	} else if !param.IsOmitted(u.OfReasoning) {
		return u.OfReasoning
	} else if !param.IsOmitted(u.OfCompaction) {
		return u.OfCompaction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetRole() *string {
	if vt := u.OfMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetAction() *ConversationItemNewParamsItemWebSearchCallActionUnion {
	if vt := u.OfWebSearchCall; vt != nil {
		return &vt.Action
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
func (u ConversationItemNewParamsItemUnion) GetSummary() []ConversationItemNewParamsItemReasoningSummary {
	if vt := u.OfReasoning; vt != nil {
		return vt.Summary
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetEncryptedContent() *string {
	if vt := u.OfCompaction; vt != nil {
		return &vt.EncryptedContent
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
	} else if vt := u.OfReasoning; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCompaction; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
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
	} else if vt := u.OfReasoning; vt != nil {
		return (*string)(&vt.Status)
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
	} else if vt := u.OfReasoning; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCompaction; vt != nil {
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

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ConversationItemNewParamsItemUnion) GetContent() (res conversationItemNewParamsItemUnionContent) {
	if vt := u.OfMessage; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfReasoning; vt != nil {
		res.any = &vt.Content
	}
	return
}

// Can have the runtime types [*string],
// [_[]ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion],
// [_[]ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion],
// [\*[]ConversationItemNewParamsItemReasoningContent]
type conversationItemNewParamsItemUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]ogxclient.ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion:
//	case *[]ogxclient.ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion:
//	case *[]ogxclient.ConversationItemNewParamsItemReasoningContent:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u conversationItemNewParamsItemUnionContent) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ConversationItemNewParamsItemUnion) GetOutput() (res conversationItemNewParamsItemUnionOutput) {
	if vt := u.OfFunctionCallOutput; vt != nil {
		res.any = vt.Output.asAny()
	} else if vt := u.OfMcpCall; vt != nil && vt.Output.Valid() {
		res.any = &vt.Output.Value
	}
	return
}

// Can have the runtime types [*string],
// [\*[]ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion]
type conversationItemNewParamsItemUnionOutput struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]ogxclient.ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u conversationItemNewParamsItemUnionOutput) AsAny() any { return u.any }

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
		apijson.Discriminator[ConversationItemNewParamsItemReasoning]("reasoning"),
		apijson.Discriminator[ConversationItemNewParamsItemCompaction]("compaction"),
	)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role are required.
type ConversationItemNewParamsItemMessage struct {
	Content ConversationItemNewParamsItemMessageContentUnion `json:"content,omitzero" api:"required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemNewParamsItemMessageRole `json:"role,omitzero" api:"required"`
	ID     param.Opt[string]                        `json:"id,omitzero"`
	Status param.Opt[string]                        `json:"status,omitzero"`
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
		return (*string)(&vt.Detail)
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
	Text string `json:"text" api:"required"`
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
	Detail ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail,omitzero"`
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
		"type", "input_image",
	)
}

type ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

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

// Text content within an output message of an OpenAI response.
//
// The property Text is required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                          `json:"text" api:"required"`
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
	FileID   string `json:"file_id" api:"required"`
	Filename string `json:"filename" api:"required"`
	Index    int64  `json:"index" api:"required"`
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
	EndIndex   int64  `json:"end_index" api:"required"`
	StartIndex int64  `json:"start_index" api:"required"`
	Title      string `json:"title" api:"required"`
	URL        string `json:"url" api:"required"`
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

// Container file citation annotation referencing a file within a container.
//
// The properties ContainerID, EndIndex, FileID, Filename, StartIndex are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id" api:"required"`
	EndIndex    int64  `json:"end_index" api:"required"`
	FileID      string `json:"file_id" api:"required"`
	Filename    string `json:"filename" api:"required"`
	StartIndex  int64  `json:"start_index" api:"required"`
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

// File path annotation referencing a generated file in response content.
//
// The properties FileID, Index are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id" api:"required"`
	Index  int64  `json:"index" api:"required"`
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
// The properties Token, Logprob are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,omitzero"`
	// The top log probabilities for the token.
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
// The properties Token, Logprob are required.
type ConversationItemNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,omitzero"`
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
	Refusal string `json:"refusal" api:"required"`
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

type ConversationItemNewParamsItemMessageRole string

const (
	ConversationItemNewParamsItemMessageRoleSystem    ConversationItemNewParamsItemMessageRole = "system"
	ConversationItemNewParamsItemMessageRoleDeveloper ConversationItemNewParamsItemMessageRole = "developer"
	ConversationItemNewParamsItemMessageRoleUser      ConversationItemNewParamsItemMessageRole = "user"
	ConversationItemNewParamsItemMessageRoleAssistant ConversationItemNewParamsItemMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
//
// The properties ID, Status are required.
type ConversationItemNewParamsItemWebSearchCall struct {
	ID     string `json:"id" api:"required"`
	Status string `json:"status" api:"required"`
	// Web search action: performs a search query.
	Action ConversationItemNewParamsItemWebSearchCallActionUnion `json:"action,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemWebSearchCallActionUnion struct {
	OfWebSearchActionSearch   *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearch   `json:",omitzero,inline"`
	OfWebSearchActionOpenPage *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionOpenPage `json:",omitzero,inline"`
	OfWebSearchActionFind     *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionFind     `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemWebSearchCallActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebSearchActionSearch, u.OfWebSearchActionOpenPage, u.OfWebSearchActionFind)
}
func (u *ConversationItemNewParamsItemWebSearchCallActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemWebSearchCallActionUnion) asAny() any {
	if !param.IsOmitted(u.OfWebSearchActionSearch) {
		return u.OfWebSearchActionSearch
	} else if !param.IsOmitted(u.OfWebSearchActionOpenPage) {
		return u.OfWebSearchActionOpenPage
	} else if !param.IsOmitted(u.OfWebSearchActionFind) {
		return u.OfWebSearchActionFind
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemWebSearchCallActionUnion) GetQuery() *string {
	if vt := u.OfWebSearchActionSearch; vt != nil {
		return &vt.Query
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemWebSearchCallActionUnion) GetQueries() []string {
	if vt := u.OfWebSearchActionSearch; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemWebSearchCallActionUnion) GetSources() []ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource {
	if vt := u.OfWebSearchActionSearch; vt != nil {
		return vt.Sources
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemWebSearchCallActionUnion) GetPattern() *string {
	if vt := u.OfWebSearchActionFind; vt != nil {
		return &vt.Pattern
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemWebSearchCallActionUnion) GetType() *string {
	if vt := u.OfWebSearchActionSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebSearchActionOpenPage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebSearchActionFind; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemWebSearchCallActionUnion) GetURL() *string {
	if vt := u.OfWebSearchActionOpenPage; vt != nil && vt.URL.Valid() {
		return &vt.URL.Value
	} else if vt := u.OfWebSearchActionFind; vt != nil {
		return (*string)(&vt.URL)
	}
	return nil
}

// Web search action: performs a search query.
//
// The property Query is required.
type ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearch struct {
	Query   string                                                                        `json:"query" api:"required"`
	Queries []string                                                                      `json:"queries,omitzero"`
	Sources []ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource `json:"sources,omitzero"`
	// Any of "search".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearch) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearch](
		"type", "search",
	)
}

// A source URL returned by a web search action.
//
// The property URL is required.
type ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource struct {
	URL string `json:"url" api:"required"`
	// Any of "url".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemWebSearchCallActionWebSearchActionSearchSource](
		"type", "url",
	)
}

// Web search action: opens a specific URL from search results.
type ConversationItemNewParamsItemWebSearchCallActionWebSearchActionOpenPage struct {
	URL param.Opt[string] `json:"url,omitzero"`
	// Any of "open_page".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemWebSearchCallActionWebSearchActionOpenPage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCallActionWebSearchActionOpenPage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionOpenPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemWebSearchCallActionWebSearchActionOpenPage](
		"type", "open_page",
	)
}

// Web search action: searches for a pattern within a loaded page.
//
// The properties Pattern, URL are required.
type ConversationItemNewParamsItemWebSearchCallActionWebSearchActionFind struct {
	Pattern string `json:"pattern" api:"required"`
	URL     string `json:"url" api:"required"`
	// Any of "find_in_page".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemWebSearchCallActionWebSearchActionFind) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCallActionWebSearchActionFind
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCallActionWebSearchActionFind) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemWebSearchCallActionWebSearchActionFind](
		"type", "find_in_page",
	)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status are required.
type ConversationItemNewParamsItemFileSearchCall struct {
	ID      string                                              `json:"id" api:"required"`
	Queries []string                                            `json:"queries,omitzero" api:"required"`
	Status  string                                              `json:"status" api:"required"`
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
	Attributes map[string]any `json:"attributes,omitzero" api:"required"`
	FileID     string         `json:"file_id" api:"required"`
	Filename   string         `json:"filename" api:"required"`
	Score      float64        `json:"score" api:"required"`
	Text       string         `json:"text" api:"required"`
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
	Arguments string            `json:"arguments" api:"required"`
	CallID    string            `json:"call_id" api:"required"`
	Name      string            `json:"name" api:"required"`
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
	CallID string                                                     `json:"call_id" api:"required"`
	Output ConversationItemNewParamsItemFunctionCallOutputOutputUnion `json:"output,omitzero" api:"required"`
	ID     param.Opt[string]                                          `json:"id,omitzero"`
	Status param.Opt[string]                                          `json:"status,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemFunctionCallOutputOutputUnion struct {
	OfString                                                                                                               param.Opt[string]                                                                                                                                                                    `json:",omitzero,inline"`
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemFunctionCallOutputOutputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile)
}
func (u *ConversationItemNewParamsItemFunctionCallOutputOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemFunctionCallOutputOutputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile) {
		return &u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	OfInputText  *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText  `json:",omitzero,inline"`
	OfInputImage *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage `json:",omitzero,inline"`
	OfInputFile  *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile  `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage, u.OfInputFile)
}
func (u *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) asAny() any {
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
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Detail)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileData() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileData.Valid() {
		return &vt.FileData.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileURL() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileURL.Valid() {
		return &vt.FileURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFilename() *string {
	if vt := u.OfInputFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetType() *string {
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
func (u ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileID() *string {
	if vt := u.OfInputImage; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	} else if vt := u.OfInputFile; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText]("input_text"),
		apijson.Discriminator[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage]("input_image"),
		apijson.Discriminator[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile]("input_file"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The property Text is required.
type ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text" api:"required"`
	// Any of "input_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText](
		"type", "input_text",
	)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Any of "low", "high", "auto".
	Detail ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail `json:"detail,omitzero"`
	// Any of "input_image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"type", "input_image",
	)
}

type ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail string

const (
	ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailLow  ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "low"
	ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailHigh ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "high"
	ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetailAuto ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	FileURL  param.Opt[string] `json:"file_url,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Any of "input_file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemFunctionCallOutputOutputListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile](
		"type", "input_file",
	)
}

// A request for human approval of a tool invocation.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ConversationItemNewParamsItemMcpApprovalRequest struct {
	ID          string `json:"id" api:"required"`
	Arguments   string `json:"arguments" api:"required"`
	Name        string `json:"name" api:"required"`
	ServerLabel string `json:"server_label" api:"required"`
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
	ApprovalRequestID string            `json:"approval_request_id" api:"required"`
	Approve           bool              `json:"approve" api:"required"`
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
	ID          string            `json:"id" api:"required"`
	Arguments   string            `json:"arguments" api:"required"`
	Name        string            `json:"name" api:"required"`
	ServerLabel string            `json:"server_label" api:"required"`
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
	ID          string                                          `json:"id" api:"required"`
	ServerLabel string                                          `json:"server_label" api:"required"`
	Tools       []ConversationItemNewParamsItemMcpListToolsTool `json:"tools,omitzero" api:"required"`
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
	InputSchema map[string]any    `json:"input_schema,omitzero" api:"required"`
	Name        string            `json:"name" api:"required"`
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

// Reasoning output from the model, representing the model's thinking process.
//
// The properties ID, Summary are required.
type ConversationItemNewParamsItemReasoning struct {
	// Unique identifier for the reasoning output item.
	ID string `json:"id" api:"required"`
	// Summary of the reasoning output.
	Summary []ConversationItemNewParamsItemReasoningSummary `json:"summary,omitzero" api:"required"`
	// The reasoning content from the model.
	Content []ConversationItemNewParamsItemReasoningContent `json:"content,omitzero"`
	// The status of the reasoning output.
	//
	// Any of "in_progress", "completed", "incomplete".
	Status string `json:"status,omitzero"`
	// The type identifier, always 'reasoning'.
	//
	// Any of "reasoning".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemReasoning) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemReasoning
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemReasoning](
		"status", "in_progress", "completed", "incomplete",
	)
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemReasoning](
		"type", "reasoning",
	)
}

// A summary of reasoning output from the model.
//
// The property Text is required.
type ConversationItemNewParamsItemReasoningSummary struct {
	// The summary text of the reasoning output.
	Text string `json:"text" api:"required"`
	// The type identifier, always 'summary_text'.
	//
	// Any of "summary_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemReasoningSummary) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemReasoningSummary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemReasoningSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemReasoningSummary](
		"type", "summary_text",
	)
}

// Reasoning text from the model.
//
// The property Text is required.
type ConversationItemNewParamsItemReasoningContent struct {
	// The reasoning text content from the model.
	Text string `json:"text" api:"required"`
	// The type identifier, always 'reasoning_text'.
	//
	// Any of "reasoning_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemReasoningContent) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemReasoningContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemReasoningContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemReasoningContent](
		"type", "reasoning_text",
	)
}

// A compaction item that summarizes prior conversation context.
//
// The property EncryptedContent is required.
type ConversationItemNewParamsItemCompaction struct {
	EncryptedContent string            `json:"encrypted_content" api:"required"`
	ID               param.Opt[string] `json:"id,omitzero"`
	// Any of "compaction".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemCompaction) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemCompaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemCompaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationItemNewParamsItemCompaction](
		"type", "compaction",
	)
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
	// The conversation identifier.
	ConversationID string `path:"conversation_id" api:"required" json:"-"`
	paramObj
}

type ConversationItemGetParams struct {
	// The conversation identifier.
	ConversationID string `path:"conversation_id" api:"required" json:"-"`
	// Any of "web_search_call.action.sources", "code_interpreter_call.outputs",
	// "computer_call_output.output.image_url", "file_search_call.results",
	// "message.input_image.image_url", "message.output_text.logprobs",
	// "reasoning.encrypted_content".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationItemGetParams]'s query parameters as
// `url.Values`.
func (r ConversationItemGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
