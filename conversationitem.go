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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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

// Create items. Create items in the conversation.
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

// List items. List items in the conversation.
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

// List items. List items in the conversation.
func (r *ConversationItemService) ListAutoPaging(ctx context.Context, conversationID string, query ConversationItemListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ConversationItemListResponseUnion] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, conversationID, query, opts...))
}

// Retrieve an item. Retrieve a conversation item.
func (r *ConversationItemService) Get(ctx context.Context, itemID string, query ConversationItemGetParams, opts ...option.RequestOption) (res *ConversationItemGetResponseUnion, err error) {
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
	Data    []ConversationItemNewResponseDataUnion `json:"data,required"`
	HasMore bool                                   `json:"has_more,required"`
	Object  string                                 `json:"object,required"`
	FirstID string                                 `json:"first_id"`
	LastID  string                                 `json:"last_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		FirstID     respjson.Field
		LastID      respjson.Field
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
	Role ConversationItemNewResponseDataMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
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
		Type              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
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
	Role   ConversationItemNewResponseDataMessageRole `json:"role,required"`
	Type   constant.Message                           `json:"type,required"`
	ID     string                                     `json:"id"`
	Status string                                     `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
// [[]ConversationItemNewResponseDataMessageContentArrayItemUnion],
// [[]ConversationItemNewResponseDataMessageContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfVariant2]
type ConversationItemNewResponseDataMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentArrayItemUnion] instead of an
	// object.
	OfVariant2 []ConversationItemNewResponseDataMessageContentArrayItemUnion `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfVariant2 respjson.Field
		raw        string
	} `json:"-"`
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsConversationItemNewResponseDataMessageContentArray() (v []ConversationItemNewResponseDataMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsVariant2() (v []ConversationItemNewResponseDataMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentArrayItemUnion contains all
// possible properties and values from
// [ConversationItemNewResponseDataMessageContentArrayItemInputText],
// [ConversationItemNewResponseDataMessageContentArrayItemInputImage],
// [ConversationItemNewResponseDataMessageContentArrayItemInputFile].
//
// Use the [ConversationItemNewResponseDataMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputImage].
	Detail ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail `json:"detail"`
	FileID string                                                                 `json:"file_id"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputFile].
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

// anyConversationItemNewResponseDataMessageContentArrayItem is implemented by each
// variant of [ConversationItemNewResponseDataMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ConversationItemNewResponseDataMessageContentArrayItemUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentArrayItem interface {
	implConversationItemNewResponseDataMessageContentArrayItemUnion()
}

func (ConversationItemNewResponseDataMessageContentArrayItemInputText) implConversationItemNewResponseDataMessageContentArrayItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentArrayItemInputImage) implConversationItemNewResponseDataMessageContentArrayItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentArrayItemInputFile) implConversationItemNewResponseDataMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemInputText:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemInputImage:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsAny() anyConversationItemNewResponseDataMessageContentArrayItem {
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

func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsInputText() (v ConversationItemNewResponseDataMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsInputImage() (v ConversationItemNewResponseDataMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsInputFile() (v ConversationItemNewResponseDataMessageContentArrayItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	Type constant.InputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail `json:"detail,required"`
	// Content type identifier, always "input_image"
	Type constant.InputImage `json:"type,required"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// (Optional) URL of the image content
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Type        respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail string

const (
	ConversationItemNewResponseDataMessageContentArrayItemInputImageDetailLow  ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail = "low"
	ConversationItemNewResponseDataMessageContentArrayItemInputImageDetailHigh ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail = "high"
	ConversationItemNewResponseDataMessageContentArrayItemInputImageDetailAuto ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentArrayItemInputFile struct {
	// The type of the input item. Always `input_file`.
	Type constant.InputFile `json:"type,required"`
	// The data of the file to be sent to the model.
	FileData string `json:"file_data"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// The URL of the file to be sent to the model.
	FileURL string `json:"file_url"`
	// The name of the file to be sent to the model.
	Filename string `json:"filename"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponseDataMessageContentArrayItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemNewResponseDataMessageContentArrayItemDetail string

const (
	ConversationItemNewResponseDataMessageContentArrayItemDetailLow  ConversationItemNewResponseDataMessageContentArrayItemDetail = "low"
	ConversationItemNewResponseDataMessageContentArrayItemDetailHigh ConversationItemNewResponseDataMessageContentArrayItemDetail = "high"
	ConversationItemNewResponseDataMessageContentArrayItemDetailAuto ConversationItemNewResponseDataMessageContentArrayItemDetail = "auto"
)

type ConversationItemNewResponseDataMessageRole string

const (
	ConversationItemNewResponseDataMessageRoleSystem    ConversationItemNewResponseDataMessageRole = "system"
	ConversationItemNewResponseDataMessageRoleDeveloper ConversationItemNewResponseDataMessageRole = "developer"
	ConversationItemNewResponseDataMessageRoleUser      ConversationItemNewResponseDataMessageRole = "user"
	ConversationItemNewResponseDataMessageRoleAssistant ConversationItemNewResponseDataMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataWebSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// Current status of the web search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "web_search_call"
	Type constant.WebSearchCall `json:"type,required"`
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
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemNewResponseDataFileSearchCallResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Results     respjson.Field
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
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemNewResponseDataFileSearchCallResultAttributeUnion `json:"attributes,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result (between 0 and 1)
	Score float64 `json:"score,required"`
	// Text content of the search result
	Text string `json:"text,required"`
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

// ConversationItemNewResponseDataFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemNewResponseDataFileSearchCallResultAttributeUnion struct {
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

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFunctionCall struct {
	// JSON string containing the function arguments
	Arguments string `json:"arguments,required"`
	// Unique identifier for the function call
	CallID string `json:"call_id,required"`
	// Name of the function being called
	Name string `json:"name,required"`
	// Tool call type identifier, always "function_call"
	Type constant.FunctionCall `json:"type,required"`
	// (Optional) Additional identifier for the tool call
	ID string `json:"id"`
	// (Optional) Current status of the function call execution
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
	CallID string                      `json:"call_id,required"`
	Output string                      `json:"output,required"`
	Type   constant.FunctionCallOutput `json:"type,required"`
	ID     string                      `json:"id"`
	Status string                      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
	ID          string                      `json:"id,required"`
	Arguments   string                      `json:"arguments,required"`
	Name        string                      `json:"name,required"`
	ServerLabel string                      `json:"server_label,required"`
	Type        constant.McpApprovalRequest `json:"type,required"`
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
	ApprovalRequestID string                       `json:"approval_request_id,required"`
	Approve           bool                         `json:"approve,required"`
	Type              constant.McpApprovalResponse `json:"type,required"`
	ID                string                       `json:"id"`
	Reason            string                       `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Type              respjson.Field
		ID                respjson.Field
		Reason            respjson.Field
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
	// Unique identifier for this MCP call
	ID string `json:"id,required"`
	// JSON string containing the MCP call arguments
	Arguments string `json:"arguments,required"`
	// Name of the MCP method being called
	Name string `json:"name,required"`
	// Label identifying the MCP server handling the call
	ServerLabel string `json:"server_label,required"`
	// Tool call type identifier, always "mcp_call"
	Type constant.McpCall `json:"type,required"`
	// (Optional) Error message if the MCP call failed
	Error string `json:"error"`
	// (Optional) Output result from the successful MCP call
	Output string `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		Error       respjson.Field
		Output      respjson.Field
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
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemNewResponseDataMcpListToolsTool `json:"tools,required"`
	// Tool call type identifier, always "mcp_list_tools"
	Type constant.McpListTools `json:"type,required"`
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
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Description of what the tool does
	Description string `json:"description"`
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

// ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion struct {
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

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataRole string

const (
	ConversationItemNewResponseDataRoleSystem    ConversationItemNewResponseDataRole = "system"
	ConversationItemNewResponseDataRoleDeveloper ConversationItemNewResponseDataRole = "developer"
	ConversationItemNewResponseDataRoleUser      ConversationItemNewResponseDataRole = "user"
	ConversationItemNewResponseDataRoleAssistant ConversationItemNewResponseDataRole = "assistant"
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
	Role ConversationItemListResponseMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
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
		Type              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
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
	Role   ConversationItemListResponseMessageRole `json:"role,required"`
	Type   constant.Message                        `json:"type,required"`
	ID     string                                  `json:"id"`
	Status string                                  `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
// [[]ConversationItemListResponseMessageContentArrayItemUnion],
// [[]ConversationItemListResponseMessageContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfVariant2]
type ConversationItemListResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseMessageContentArrayItemUnion] instead of an
	// object.
	OfVariant2 []ConversationItemListResponseMessageContentArrayItemUnion `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfVariant2 respjson.Field
		raw        string
	} `json:"-"`
}

func (u ConversationItemListResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentUnion) AsConversationItemListResponseMessageContentArray() (v []ConversationItemListResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentUnion) AsVariant2() (v []ConversationItemListResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseMessageContentArrayItemUnion contains all possible
// properties and values from
// [ConversationItemListResponseMessageContentArrayItemInputText],
// [ConversationItemListResponseMessageContentArrayItemInputImage],
// [ConversationItemListResponseMessageContentArrayItemInputFile].
//
// Use the [ConversationItemListResponseMessageContentArrayItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ConversationItemListResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentArrayItemInputImage].
	Detail ConversationItemListResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	FileID string                                                              `json:"file_id"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentArrayItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentArrayItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemListResponseMessageContentArrayItemInputFile].
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

// anyConversationItemListResponseMessageContentArrayItem is implemented by each
// variant of [ConversationItemListResponseMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ConversationItemListResponseMessageContentArrayItemUnion.AsAny]
type anyConversationItemListResponseMessageContentArrayItem interface {
	implConversationItemListResponseMessageContentArrayItemUnion()
}

func (ConversationItemListResponseMessageContentArrayItemInputText) implConversationItemListResponseMessageContentArrayItemUnion() {
}
func (ConversationItemListResponseMessageContentArrayItemInputImage) implConversationItemListResponseMessageContentArrayItemUnion() {
}
func (ConversationItemListResponseMessageContentArrayItemInputFile) implConversationItemListResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseMessageContentArrayItemInputText:
//	case llamastackclient.ConversationItemListResponseMessageContentArrayItemInputImage:
//	case llamastackclient.ConversationItemListResponseMessageContentArrayItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseMessageContentArrayItemUnion) AsAny() anyConversationItemListResponseMessageContentArrayItem {
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

func (u ConversationItemListResponseMessageContentArrayItemUnion) AsInputText() (v ConversationItemListResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentArrayItemUnion) AsInputImage() (v ConversationItemListResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMessageContentArrayItemUnion) AsInputFile() (v ConversationItemListResponseMessageContentArrayItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	Type constant.InputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemListResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
	// Content type identifier, always "input_image"
	Type constant.InputImage `json:"type,required"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// (Optional) URL of the image content
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Type        respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemListResponseMessageContentArrayItemInputImageDetail string

const (
	ConversationItemListResponseMessageContentArrayItemInputImageDetailLow  ConversationItemListResponseMessageContentArrayItemInputImageDetail = "low"
	ConversationItemListResponseMessageContentArrayItemInputImageDetailHigh ConversationItemListResponseMessageContentArrayItemInputImageDetail = "high"
	ConversationItemListResponseMessageContentArrayItemInputImageDetailAuto ConversationItemListResponseMessageContentArrayItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemListResponseMessageContentArrayItemInputFile struct {
	// The type of the input item. Always `input_file`.
	Type constant.InputFile `json:"type,required"`
	// The data of the file to be sent to the model.
	FileData string `json:"file_data"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// The URL of the file to be sent to the model.
	FileURL string `json:"file_url"`
	// The name of the file to be sent to the model.
	Filename string `json:"filename"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponseMessageContentArrayItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseMessageContentArrayItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemListResponseMessageContentArrayItemDetail string

const (
	ConversationItemListResponseMessageContentArrayItemDetailLow  ConversationItemListResponseMessageContentArrayItemDetail = "low"
	ConversationItemListResponseMessageContentArrayItemDetailHigh ConversationItemListResponseMessageContentArrayItemDetail = "high"
	ConversationItemListResponseMessageContentArrayItemDetailAuto ConversationItemListResponseMessageContentArrayItemDetail = "auto"
)

type ConversationItemListResponseMessageRole string

const (
	ConversationItemListResponseMessageRoleSystem    ConversationItemListResponseMessageRole = "system"
	ConversationItemListResponseMessageRoleDeveloper ConversationItemListResponseMessageRole = "developer"
	ConversationItemListResponseMessageRoleUser      ConversationItemListResponseMessageRole = "user"
	ConversationItemListResponseMessageRoleAssistant ConversationItemListResponseMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ConversationItemListResponseWebSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// Current status of the web search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "web_search_call"
	Type constant.WebSearchCall `json:"type,required"`
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
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemListResponseFileSearchCallResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Results     respjson.Field
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
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemListResponseFileSearchCallResultAttributeUnion `json:"attributes,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result (between 0 and 1)
	Score float64 `json:"score,required"`
	// Text content of the search result
	Text string `json:"text,required"`
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

// ConversationItemListResponseFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemListResponseFileSearchCallResultAttributeUnion struct {
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

func (u ConversationItemListResponseFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemListResponseFunctionCall struct {
	// JSON string containing the function arguments
	Arguments string `json:"arguments,required"`
	// Unique identifier for the function call
	CallID string `json:"call_id,required"`
	// Name of the function being called
	Name string `json:"name,required"`
	// Tool call type identifier, always "function_call"
	Type constant.FunctionCall `json:"type,required"`
	// (Optional) Additional identifier for the tool call
	ID string `json:"id"`
	// (Optional) Current status of the function call execution
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
	CallID string                      `json:"call_id,required"`
	Output string                      `json:"output,required"`
	Type   constant.FunctionCallOutput `json:"type,required"`
	ID     string                      `json:"id"`
	Status string                      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
	ID          string                      `json:"id,required"`
	Arguments   string                      `json:"arguments,required"`
	Name        string                      `json:"name,required"`
	ServerLabel string                      `json:"server_label,required"`
	Type        constant.McpApprovalRequest `json:"type,required"`
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
	ApprovalRequestID string                       `json:"approval_request_id,required"`
	Approve           bool                         `json:"approve,required"`
	Type              constant.McpApprovalResponse `json:"type,required"`
	ID                string                       `json:"id"`
	Reason            string                       `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Type              respjson.Field
		ID                respjson.Field
		Reason            respjson.Field
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
	// Unique identifier for this MCP call
	ID string `json:"id,required"`
	// JSON string containing the MCP call arguments
	Arguments string `json:"arguments,required"`
	// Name of the MCP method being called
	Name string `json:"name,required"`
	// Label identifying the MCP server handling the call
	ServerLabel string `json:"server_label,required"`
	// Tool call type identifier, always "mcp_call"
	Type constant.McpCall `json:"type,required"`
	// (Optional) Error message if the MCP call failed
	Error string `json:"error"`
	// (Optional) Output result from the successful MCP call
	Output string `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		Error       respjson.Field
		Output      respjson.Field
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
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemListResponseMcpListToolsTool `json:"tools,required"`
	// Tool call type identifier, always "mcp_list_tools"
	Type constant.McpListTools `json:"type,required"`
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
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemListResponseMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Description of what the tool does
	Description string `json:"description"`
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

// ConversationItemListResponseMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemListResponseMcpListToolsToolInputSchemaUnion struct {
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

func (u ConversationItemListResponseMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
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
// [ConversationItemGetResponseMcpCall], [ConversationItemGetResponseMcpListTools].
//
// Use the [ConversationItemGetResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseUnion struct {
	// This field is from variant [ConversationItemGetResponseMessage].
	Content ConversationItemGetResponseMessageContentUnion `json:"content"`
	// This field is from variant [ConversationItemGetResponseMessage].
	Role ConversationItemGetResponseMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "function_call_output", "mcp_approval_request", "mcp_approval_response",
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [ConversationItemGetResponseFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemGetResponseFileSearchCall].
	Results     []ConversationItemGetResponseFileSearchCallResult `json:"results"`
	Arguments   string                                            `json:"arguments"`
	CallID      string                                            `json:"call_id"`
	Name        string                                            `json:"name"`
	Output      string                                            `json:"output"`
	ServerLabel string                                            `json:"server_label"`
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
	JSON  struct {
		Content           respjson.Field
		Role              respjson.Field
		Type              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseMessage:
//	case llamastackclient.ConversationItemGetResponseWebSearchCall:
//	case llamastackclient.ConversationItemGetResponseFileSearchCall:
//	case llamastackclient.ConversationItemGetResponseFunctionCall:
//	case llamastackclient.ConversationItemGetResponseFunctionCallOutput:
//	case llamastackclient.ConversationItemGetResponseMcpApprovalRequest:
//	case llamastackclient.ConversationItemGetResponseMcpApprovalResponse:
//	case llamastackclient.ConversationItemGetResponseMcpCall:
//	case llamastackclient.ConversationItemGetResponseMcpListTools:
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

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemGetResponseMessage struct {
	Content ConversationItemGetResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemGetResponseMessageRole `json:"role,required"`
	Type   constant.Message                       `json:"type,required"`
	ID     string                                 `json:"id"`
	Status string                                 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
// [[]ConversationItemGetResponseMessageContentArrayItemUnion],
// [[]ConversationItemGetResponseMessageContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfVariant2]
type ConversationItemGetResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentArrayItemUnion] instead of an
	// object.
	OfVariant2 []ConversationItemGetResponseMessageContentArrayItemUnion `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfVariant2 respjson.Field
		raw        string
	} `json:"-"`
}

func (u ConversationItemGetResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentUnion) AsConversationItemGetResponseMessageContentArray() (v []ConversationItemGetResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentUnion) AsVariant2() (v []ConversationItemGetResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentArrayItemUnion contains all possible
// properties and values from
// [ConversationItemGetResponseMessageContentArrayItemInputText],
// [ConversationItemGetResponseMessageContentArrayItemInputImage],
// [ConversationItemGetResponseMessageContentArrayItemInputFile].
//
// Use the [ConversationItemGetResponseMessageContentArrayItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputImage].
	Detail ConversationItemGetResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	FileID string                                                             `json:"file_id"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputFile].
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

// anyConversationItemGetResponseMessageContentArrayItem is implemented by each
// variant of [ConversationItemGetResponseMessageContentArrayItemUnion] to add type
// safety for the return type of
// [ConversationItemGetResponseMessageContentArrayItemUnion.AsAny]
type anyConversationItemGetResponseMessageContentArrayItem interface {
	implConversationItemGetResponseMessageContentArrayItemUnion()
}

func (ConversationItemGetResponseMessageContentArrayItemInputText) implConversationItemGetResponseMessageContentArrayItemUnion() {
}
func (ConversationItemGetResponseMessageContentArrayItemInputImage) implConversationItemGetResponseMessageContentArrayItemUnion() {
}
func (ConversationItemGetResponseMessageContentArrayItemInputFile) implConversationItemGetResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemInputText:
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemInputImage:
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsAny() anyConversationItemGetResponseMessageContentArrayItem {
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

func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsInputText() (v ConversationItemGetResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsInputImage() (v ConversationItemGetResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsInputFile() (v ConversationItemGetResponseMessageContentArrayItemInputFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	Type constant.InputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemGetResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
	// Content type identifier, always "input_image"
	Type constant.InputImage `json:"type,required"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// (Optional) URL of the image content
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Type        respjson.Field
		FileID      respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemGetResponseMessageContentArrayItemInputImageDetail string

const (
	ConversationItemGetResponseMessageContentArrayItemInputImageDetailLow  ConversationItemGetResponseMessageContentArrayItemInputImageDetail = "low"
	ConversationItemGetResponseMessageContentArrayItemInputImageDetailHigh ConversationItemGetResponseMessageContentArrayItemInputImageDetail = "high"
	ConversationItemGetResponseMessageContentArrayItemInputImageDetailAuto ConversationItemGetResponseMessageContentArrayItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentArrayItemInputFile struct {
	// The type of the input item. Always `input_file`.
	Type constant.InputFile `json:"type,required"`
	// The data of the file to be sent to the model.
	FileData string `json:"file_data"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// The URL of the file to be sent to the model.
	FileURL string `json:"file_url"`
	// The name of the file to be sent to the model.
	Filename string `json:"filename"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseMessageContentArrayItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemGetResponseMessageContentArrayItemDetail string

const (
	ConversationItemGetResponseMessageContentArrayItemDetailLow  ConversationItemGetResponseMessageContentArrayItemDetail = "low"
	ConversationItemGetResponseMessageContentArrayItemDetailHigh ConversationItemGetResponseMessageContentArrayItemDetail = "high"
	ConversationItemGetResponseMessageContentArrayItemDetailAuto ConversationItemGetResponseMessageContentArrayItemDetail = "auto"
)

type ConversationItemGetResponseMessageRole string

const (
	ConversationItemGetResponseMessageRoleSystem    ConversationItemGetResponseMessageRole = "system"
	ConversationItemGetResponseMessageRoleDeveloper ConversationItemGetResponseMessageRole = "developer"
	ConversationItemGetResponseMessageRoleUser      ConversationItemGetResponseMessageRole = "user"
	ConversationItemGetResponseMessageRoleAssistant ConversationItemGetResponseMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ConversationItemGetResponseWebSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// Current status of the web search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "web_search_call"
	Type constant.WebSearchCall `json:"type,required"`
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
func (r ConversationItemGetResponseWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemGetResponseFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemGetResponseFileSearchCallResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Results     respjson.Field
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
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemGetResponseFileSearchCallResultAttributeUnion `json:"attributes,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result (between 0 and 1)
	Score float64 `json:"score,required"`
	// Text content of the search result
	Text string `json:"text,required"`
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

// ConversationItemGetResponseFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemGetResponseFileSearchCallResultAttributeUnion struct {
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

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ConversationItemGetResponseFunctionCall struct {
	// JSON string containing the function arguments
	Arguments string `json:"arguments,required"`
	// Unique identifier for the function call
	CallID string `json:"call_id,required"`
	// Name of the function being called
	Name string `json:"name,required"`
	// Tool call type identifier, always "function_call"
	Type constant.FunctionCall `json:"type,required"`
	// (Optional) Additional identifier for the tool call
	ID string `json:"id"`
	// (Optional) Current status of the function call execution
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
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
	CallID string                      `json:"call_id,required"`
	Output string                      `json:"output,required"`
	Type   constant.FunctionCallOutput `json:"type,required"`
	ID     string                      `json:"id"`
	Status string                      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemGetResponseFunctionCallOutput) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ConversationItemGetResponseMcpApprovalRequest struct {
	ID          string                      `json:"id,required"`
	Arguments   string                      `json:"arguments,required"`
	Name        string                      `json:"name,required"`
	ServerLabel string                      `json:"server_label,required"`
	Type        constant.McpApprovalRequest `json:"type,required"`
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
	ApprovalRequestID string                       `json:"approval_request_id,required"`
	Approve           bool                         `json:"approve,required"`
	Type              constant.McpApprovalResponse `json:"type,required"`
	ID                string                       `json:"id"`
	Reason            string                       `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Type              respjson.Field
		ID                respjson.Field
		Reason            respjson.Field
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
	// Unique identifier for this MCP call
	ID string `json:"id,required"`
	// JSON string containing the MCP call arguments
	Arguments string `json:"arguments,required"`
	// Name of the MCP method being called
	Name string `json:"name,required"`
	// Label identifying the MCP server handling the call
	ServerLabel string `json:"server_label,required"`
	// Tool call type identifier, always "mcp_call"
	Type constant.McpCall `json:"type,required"`
	// (Optional) Error message if the MCP call failed
	Error string `json:"error"`
	// (Optional) Output result from the successful MCP call
	Output string `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		Error       respjson.Field
		Output      respjson.Field
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
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemGetResponseMcpListToolsTool `json:"tools,required"`
	// Tool call type identifier, always "mcp_list_tools"
	Type constant.McpListTools `json:"type,required"`
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
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemGetResponseMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Description of what the tool does
	Description string `json:"description"`
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

// ConversationItemGetResponseMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemGetResponseMcpListToolsToolInputSchemaUnion struct {
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

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
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
	// Items to include in the conversation context.
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
		return (*string)(&vt.Role)
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
// The properties Content, Role, Type are required.
type ConversationItemNewParamsItemMessage struct {
	Content ConversationItemNewParamsItemMessageContentUnion `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemNewParamsItemMessageRole `json:"role,omitzero,required"`
	ID     param.Opt[string]                        `json:"id,omitzero"`
	Status param.Opt[string]                        `json:"status,omitzero"`
	// This field can be elided, and will marshal its zero value as "message".
	Type constant.Message `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentUnion struct {
	OfString                                      param.Opt[string]                                           `json:",omitzero,inline"`
	OfConversationItemNewsItemMessageContentArray []ConversationItemNewParamsItemMessageContentArrayItemUnion `json:",omitzero,inline"`
	OfVariant2                                    []ConversationItemNewParamsItemMessageContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConversationItemNewsItemMessageContentArray, u.OfVariant2)
}
func (u *ConversationItemNewParamsItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfConversationItemNewsItemMessageContentArray) {
		return &u.OfConversationItemNewsItemMessageContentArray
	} else if !param.IsOmitted(u.OfVariant2) {
		return &u.OfVariant2
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentArrayItemUnion struct {
	OfInputText  *ConversationItemNewParamsItemMessageContentArrayItemInputText  `json:",omitzero,inline"`
	OfInputImage *ConversationItemNewParamsItemMessageContentArrayItemInputImage `json:",omitzero,inline"`
	OfInputFile  *ConversationItemNewParamsItemMessageContentArrayItemInputFile  `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage, u.OfInputFile)
}
func (u *ConversationItemNewParamsItemMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentArrayItemUnion) asAny() any {
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
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Detail)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetFileData() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileData.Valid() {
		return &vt.FileData.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetFileURL() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileURL.Valid() {
		return &vt.FileURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetFilename() *string {
	if vt := u.OfInputFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetType() *string {
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
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetFileID() *string {
	if vt := u.OfInputImage; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	} else if vt := u.OfInputFile; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentArrayItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemInputText]("input_text"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemInputImage]("input_image"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemInputFile]("input_file"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The properties Text, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	//
	// This field can be elided, and will marshal its zero value as "input_text".
	Type constant.InputText `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
//
// The properties Detail, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail `json:"detail,omitzero,required"`
	// (Optional) The ID of the file to be sent to the model.
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// (Optional) URL of the image content
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Content type identifier, always "input_image"
	//
	// This field can be elided, and will marshal its zero value as "input_image".
	Type constant.InputImage `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail string

const (
	ConversationItemNewParamsItemMessageContentArrayItemInputImageDetailLow  ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail = "low"
	ConversationItemNewParamsItemMessageContentArrayItemInputImageDetailHigh ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail = "high"
	ConversationItemNewParamsItemMessageContentArrayItemInputImageDetailAuto ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail = "auto"
)

// File content for input messages in OpenAI response format.
//
// The property Type is required.
type ConversationItemNewParamsItemMessageContentArrayItemInputFile struct {
	// The data of the file to be sent to the model.
	FileData param.Opt[string] `json:"file_data,omitzero"`
	// (Optional) The ID of the file to be sent to the model.
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// The URL of the file to be sent to the model.
	FileURL param.Opt[string] `json:"file_url,omitzero"`
	// The name of the file to be sent to the model.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// The type of the input item. Always `input_file`.
	//
	// This field can be elided, and will marshal its zero value as "input_file".
	Type constant.InputFile `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItemInputFile) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemInputFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
// The properties ID, Status, Type are required.
type ConversationItemNewParamsItemWebSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// Current status of the web search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "web_search_call"
	//
	// This field can be elided, and will marshal its zero value as "web_search_call".
	Type constant.WebSearchCall `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemWebSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status, Type are required.
type ConversationItemNewParamsItemFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,omitzero,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemNewParamsItemFileSearchCallResult `json:"results,omitzero"`
	// Tool call type identifier, always "file_search_call"
	//
	// This field can be elided, and will marshal its zero value as "file_search_call".
	Type constant.FileSearchCall `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemFileSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFileSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
//
// The properties Attributes, FileID, Filename, Score, Text are required.
type ConversationItemNewParamsItemFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemNewParamsItemFileSearchCallResultAttributeUnion `json:"attributes,omitzero,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result (between 0 and 1)
	Score float64 `json:"score,required"`
	// Text content of the search result
	Text string `json:"text,required"`
	paramObj
}

func (r ConversationItemNewParamsItemFileSearchCallResult) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFileSearchCallResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemFileSearchCallResultAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemFileSearchCallResultAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ConversationItemNewParamsItemFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemFileSearchCallResultAttributeUnion) asAny() any {
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

// Function tool call output message for OpenAI responses.
//
// The properties Arguments, CallID, Name, Type are required.
type ConversationItemNewParamsItemFunctionCall struct {
	// JSON string containing the function arguments
	Arguments string `json:"arguments,required"`
	// Unique identifier for the function call
	CallID string `json:"call_id,required"`
	// Name of the function being called
	Name string `json:"name,required"`
	// (Optional) Additional identifier for the tool call
	ID param.Opt[string] `json:"id,omitzero"`
	// (Optional) Current status of the function call execution
	Status param.Opt[string] `json:"status,omitzero"`
	// Tool call type identifier, always "function_call"
	//
	// This field can be elided, and will marshal its zero value as "function_call".
	Type constant.FunctionCall `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
//
// The properties CallID, Output, Type are required.
type ConversationItemNewParamsItemFunctionCallOutput struct {
	CallID string            `json:"call_id,required"`
	Output string            `json:"output,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "function_call_output".
	Type constant.FunctionCallOutput `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemFunctionCallOutput) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCallOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
//
// The properties ID, Arguments, Name, ServerLabel, Type are required.
type ConversationItemNewParamsItemMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// This field can be elided, and will marshal its zero value as
	// "mcp_approval_request".
	Type constant.McpApprovalRequest `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpApprovalRequest) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpApprovalRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
//
// The properties ApprovalRequestID, Approve, Type are required.
type ConversationItemNewParamsItemMcpApprovalResponse struct {
	ApprovalRequestID string            `json:"approval_request_id,required"`
	Approve           bool              `json:"approve,required"`
	ID                param.Opt[string] `json:"id,omitzero"`
	Reason            param.Opt[string] `json:"reason,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "mcp_approval_response".
	Type constant.McpApprovalResponse `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpApprovalResponse) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpApprovalResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
//
// The properties ID, Arguments, Name, ServerLabel, Type are required.
type ConversationItemNewParamsItemMcpCall struct {
	// Unique identifier for this MCP call
	ID string `json:"id,required"`
	// JSON string containing the MCP call arguments
	Arguments string `json:"arguments,required"`
	// Name of the MCP method being called
	Name string `json:"name,required"`
	// Label identifying the MCP server handling the call
	ServerLabel string `json:"server_label,required"`
	// (Optional) Error message if the MCP call failed
	Error param.Opt[string] `json:"error,omitzero"`
	// (Optional) Output result from the successful MCP call
	Output param.Opt[string] `json:"output,omitzero"`
	// Tool call type identifier, always "mcp_call"
	//
	// This field can be elided, and will marshal its zero value as "mcp_call".
	Type constant.McpCall `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
//
// The properties ID, ServerLabel, Tools, Type are required.
type ConversationItemNewParamsItemMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemNewParamsItemMcpListToolsTool `json:"tools,omitzero,required"`
	// Tool call type identifier, always "mcp_list_tools"
	//
	// This field can be elided, and will marshal its zero value as "mcp_list_tools".
	Type constant.McpListTools `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpListTools) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpListTools
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
//
// The properties InputSchema, Name are required.
type ConversationItemNewParamsItemMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion `json:"input_schema,omitzero,required"`
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Description of what the tool does
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion) asAny() any {
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

type ConversationItemListParams struct {
	// An item ID to list items after, used in pagination.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A limit on the number of objects to be returned (1-100, default 20).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Specify additional output data to include in the response.
	//
	// Any of "web_search_call.action.sources", "code_interpreter_call.outputs",
	// "computer_call_output.output.image_url", "file_search_call.results",
	// "message.input_image.image_url", "message.output_text.logprobs",
	// "reasoning.encrypted_content".
	Include []string `query:"include,omitzero" json:"-"`
	// The order to return items in (asc or desc, default desc).
	//
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

// The order to return items in (asc or desc, default desc).
type ConversationItemListParamsOrder string

const (
	ConversationItemListParamsOrderAsc  ConversationItemListParamsOrder = "asc"
	ConversationItemListParamsOrderDesc ConversationItemListParamsOrder = "desc"
)

type ConversationItemGetParams struct {
	ConversationID string `path:"conversation_id,required" json:"-"`
	paramObj
}
