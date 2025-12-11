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
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ConversationService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationService] method instead.
type ConversationService struct {
	Options []option.RequestOption
	Items   ConversationItemService
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.Options = opts
	r.Items = NewConversationItemService(opts...)
	return
}

// Create a conversation.
//
// Create a conversation.
func (r *ConversationService) New(ctx context.Context, body ConversationNewParams, opts ...option.RequestOption) (res *ConversationObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a conversation.
//
// Get a conversation with the given ID.
func (r *ConversationService) Get(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *ConversationObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a conversation.
//
// Update a conversation's metadata with the given ID.
func (r *ConversationService) Update(ctx context.Context, conversationID string, body ConversationUpdateParams, opts ...option.RequestOption) (res *ConversationObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a conversation.
//
// Delete a conversation with the given ID.
func (r *ConversationService) Delete(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *ConversationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// OpenAI-compatible conversation object.
type ConversationObject struct {
	// The unique ID of the conversation.
	ID string `json:"id,required"`
	// The time at which the conversation was created, measured in seconds since the
	// Unix epoch.
	CreatedAt int64 `json:"created_at,required"`
	// Initial items to include in the conversation context. You may add up to 20 items
	// at a time.
	Items []map[string]any `json:"items,nullable"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard.
	Metadata map[string]string `json:"metadata,nullable"`
	// The object type, which is always conversation.
	//
	// Any of "conversation".
	Object ConversationObjectObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Items       respjson.Field
		Metadata    respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationObject) RawJSON() string { return r.JSON.raw }
func (r *ConversationObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type, which is always conversation.
type ConversationObjectObject string

const (
	ConversationObjectObjectConversation ConversationObjectObject = "conversation"
)

// Response for deleted conversation.
type ConversationDeleteResponse struct {
	// The deleted conversation identifier
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
func (r ConversationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationNewParams struct {
	Items    []ConversationNewParamsItemUnion `json:"items,omitzero"`
	Metadata map[string]string                `json:"metadata,omitzero"`
	paramObj
}

func (r ConversationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationNewParamsItemUnion struct {
	OfMessage             *ConversationNewParamsItemMessage             `json:",omitzero,inline"`
	OfWebSearchCall       *ConversationNewParamsItemWebSearchCall       `json:",omitzero,inline"`
	OfFileSearchCall      *ConversationNewParamsItemFileSearchCall      `json:",omitzero,inline"`
	OfFunctionCall        *ConversationNewParamsItemFunctionCall        `json:",omitzero,inline"`
	OfFunctionCallOutput  *ConversationNewParamsItemFunctionCallOutput  `json:",omitzero,inline"`
	OfMcpApprovalRequest  *ConversationNewParamsItemMcpApprovalRequest  `json:",omitzero,inline"`
	OfMcpApprovalResponse *ConversationNewParamsItemMcpApprovalResponse `json:",omitzero,inline"`
	OfMcpCall             *ConversationNewParamsItemMcpCall             `json:",omitzero,inline"`
	OfMcpListTools        *ConversationNewParamsItemMcpListTools        `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationNewParamsItemUnion) MarshalJSON() ([]byte, error) {
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
func (u *ConversationNewParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationNewParamsItemUnion) asAny() any {
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
func (u ConversationNewParamsItemUnion) GetContent() *ConversationNewParamsItemMessageContentUnion {
	if vt := u.OfMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetRole() *string {
	if vt := u.OfMessage; vt != nil {
		return &vt.Role
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetQueries() []string {
	if vt := u.OfFileSearchCall; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetResults() []ConversationNewParamsItemFileSearchCallResult {
	if vt := u.OfFileSearchCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetApprovalRequestID() *string {
	if vt := u.OfMcpApprovalResponse; vt != nil {
		return &vt.ApprovalRequestID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetApprove() *bool {
	if vt := u.OfMcpApprovalResponse; vt != nil {
		return &vt.Approve
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetReason() *string {
	if vt := u.OfMcpApprovalResponse; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetError() *string {
	if vt := u.OfMcpCall; vt != nil && vt.Error.Valid() {
		return &vt.Error.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetTools() []ConversationNewParamsItemMcpListToolsTool {
	if vt := u.OfMcpListTools; vt != nil {
		return vt.Tools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetID() *string {
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
func (u ConversationNewParamsItemUnion) GetStatus() *string {
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
func (u ConversationNewParamsItemUnion) GetType() *string {
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
func (u ConversationNewParamsItemUnion) GetArguments() *string {
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
func (u ConversationNewParamsItemUnion) GetCallID() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.CallID)
	} else if vt := u.OfFunctionCallOutput; vt != nil {
		return (*string)(&vt.CallID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetName() *string {
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
func (u ConversationNewParamsItemUnion) GetOutput() *string {
	if vt := u.OfFunctionCallOutput; vt != nil {
		return (*string)(&vt.Output)
	} else if vt := u.OfMcpCall; vt != nil && vt.Output.Valid() {
		return &vt.Output.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemUnion) GetServerLabel() *string {
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
	apijson.RegisterUnion[ConversationNewParamsItemUnion](
		"type",
		apijson.Discriminator[ConversationNewParamsItemMessage]("message"),
		apijson.Discriminator[ConversationNewParamsItemWebSearchCall]("web_search_call"),
		apijson.Discriminator[ConversationNewParamsItemFileSearchCall]("file_search_call"),
		apijson.Discriminator[ConversationNewParamsItemFunctionCall]("function_call"),
		apijson.Discriminator[ConversationNewParamsItemFunctionCallOutput]("function_call_output"),
		apijson.Discriminator[ConversationNewParamsItemMcpApprovalRequest]("mcp_approval_request"),
		apijson.Discriminator[ConversationNewParamsItemMcpApprovalResponse]("mcp_approval_response"),
		apijson.Discriminator[ConversationNewParamsItemMcpCall]("mcp_call"),
		apijson.Discriminator[ConversationNewParamsItemMcpListTools]("mcp_list_tools"),
	)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role are required.
type ConversationNewParamsItemMessage struct {
	Content ConversationNewParamsItemMessageContentUnion `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   string            `json:"role,omitzero,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// Any of "message".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessage](
		"role", "system", "developer", "user", "assistant",
	)
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessage](
		"type", "message",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationNewParamsItemMessageContentUnion struct {
	OfString                                                                                                               param.Opt[string]                                                                                                                                                      `json:",omitzero,inline"`
	OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile []ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion `json:",omitzero,inline"`
	OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal                                []ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion                                `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationNewParamsItemMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFile, u.OfListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusal)
}
func (u *ConversationNewParamsItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationNewParamsItemMessageContentUnion) asAny() any {
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
type ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion struct {
	OfInputText  *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText  `json:",omitzero,inline"`
	OfInputImage *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage `json:",omitzero,inline"`
	OfInputFile  *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile  `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage, u.OfInputFile)
}
func (u *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) asAny() any {
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
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return &vt.Detail
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileData() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileData.Valid() {
		return &vt.FileData.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileURL() *string {
	if vt := u.OfInputFile; vt != nil && vt.FileURL.Valid() {
		return &vt.FileURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFilename() *string {
	if vt := u.OfInputFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetType() *string {
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
func (u ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion) GetFileID() *string {
	if vt := u.OfInputImage; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	} else if vt := u.OfInputFile; vt != nil && vt.FileID.Valid() {
		return &vt.FileID.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemUnion](
		"type",
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText]("input_text"),
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage]("input_image"),
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile]("input_file"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The property Text is required.
type ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText struct {
	Text string `json:"text,required"`
	// Any of "input_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputText](
		"type", "input_text",
	)
}

// Image content for input messages in OpenAI response format.
type ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage struct {
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	// Any of "input_image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"detail", "low", "high", "auto",
	)
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputImage](
		"type", "input_image",
	)
}

// File content for input messages in OpenAI response format.
type ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	FileURL  param.Opt[string] `json:"file_url,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Any of "input_file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseInputMessageContentTextOpenAIResponseInputMessageContentImageOpenAIResponseInputMessageContentFileItemInputFile](
		"type", "input_file",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion struct {
	OfOutputText *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText `json:",omitzero,inline"`
	OfRefusal    *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal    `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOutputText, u.OfRefusal)
}
func (u *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOutputText) {
		return u.OfOutputText
	} else if !param.IsOmitted(u.OfRefusal) {
		return u.OfRefusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetText() *string {
	if vt := u.OfOutputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetAnnotations() []ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion {
	if vt := u.OfOutputText; vt != nil {
		return vt.Annotations
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetLogprobs() []ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob {
	if vt := u.OfOutputText; vt != nil {
		return vt.Logprobs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetRefusal() *string {
	if vt := u.OfRefusal; vt != nil {
		return &vt.Refusal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion) GetType() *string {
	if vt := u.OfOutputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRefusal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemUnion](
		"type",
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText]("output_text"),
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal]("refusal"),
	)
}

// The property Text is required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText struct {
	Text        string                                                                                                                                                      `json:"text,required"`
	Logprobs    []ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob         `json:"logprobs,omitzero"`
	Annotations []ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion `json:"annotations,omitzero"`
	// Any of "output_text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputText](
		"type", "output_text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion struct {
	OfFileCitation          *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation          `json:",omitzero,inline"`
	OfURLCitation           *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation           `json:",omitzero,inline"`
	OfContainerFileCitation *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation `json:",omitzero,inline"`
	OfFilePath              *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) asAny() any {
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
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFileID() *string {
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
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetType() *string {
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
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationUnion](
		"type",
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation]("file_citation"),
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation]("url_citation"),
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation]("container_file_citation"),
		apijson.Discriminator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath]("file_path"),
	)
}

// File citation annotation for referencing specific files in response content.
//
// The properties FileID, Filename, Index are required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// Any of "file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFileCitation](
		"type", "file_citation",
	)
}

// URL citation annotation for referencing external web resources.
//
// The properties EndIndex, StartIndex, Title, URL are required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// Any of "url_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationURLCitation](
		"type", "url_citation",
	)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex are required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// Any of "container_file_citation".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationContainerFileCitation](
		"type", "container_file_citation",
	)
}

// The properties FileID, Index are required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// Any of "file_path".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextAnnotationFilePath](
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
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob struct {
	Token       string                                                                                                                                                        `json:"token,required"`
	Logprob     float64                                                                                                                                                       `json:"logprob,required"`
	Bytes       []int64                                                                                                                                                       `json:"bytes,omitzero"`
	TopLogprobs []ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob `json:"top_logprobs,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
//
// The properties Token, Logprob are required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
//
// The property Refusal is required.
type ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal struct {
	Refusal string `json:"refusal,required"`
	// Any of "refusal".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMessageContentListOpenAIResponseOutputMessageContentOutputTextInputOpenAIResponseContentPartRefusalItemRefusal](
		"type", "refusal",
	)
}

// Web search tool call output message for OpenAI responses.
//
// The properties ID, Status are required.
type ConversationNewParamsItemWebSearchCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// Any of "web_search_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemWebSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemWebSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemWebSearchCall](
		"type", "web_search_call",
	)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status are required.
type ConversationNewParamsItemFileSearchCall struct {
	ID      string                                          `json:"id,required"`
	Queries []string                                        `json:"queries,omitzero,required"`
	Status  string                                          `json:"status,required"`
	Results []ConversationNewParamsItemFileSearchCallResult `json:"results,omitzero"`
	// Any of "file_search_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemFileSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemFileSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemFileSearchCall](
		"type", "file_search_call",
	)
}

// Search results returned by the file search operation.
//
// The properties Attributes, FileID, Filename, Score, Text are required.
type ConversationNewParamsItemFileSearchCallResult struct {
	Attributes map[string]any `json:"attributes,omitzero,required"`
	FileID     string         `json:"file_id,required"`
	Filename   string         `json:"filename,required"`
	Score      float64        `json:"score,required"`
	Text       string         `json:"text,required"`
	paramObj
}

func (r ConversationNewParamsItemFileSearchCallResult) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemFileSearchCallResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
//
// The properties Arguments, CallID, Name are required.
type ConversationNewParamsItemFunctionCall struct {
	Arguments string            `json:"arguments,required"`
	CallID    string            `json:"call_id,required"`
	Name      string            `json:"name,required"`
	ID        param.Opt[string] `json:"id,omitzero"`
	Status    param.Opt[string] `json:"status,omitzero"`
	// Any of "function_call".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemFunctionCall](
		"type", "function_call",
	)
}

// This represents the output of a function call that gets passed back to the
// model.
//
// The properties CallID, Output are required.
type ConversationNewParamsItemFunctionCallOutput struct {
	CallID string            `json:"call_id,required"`
	Output string            `json:"output,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// Any of "function_call_output".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemFunctionCallOutput) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemFunctionCallOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemFunctionCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemFunctionCallOutput](
		"type", "function_call_output",
	)
}

// A request for human approval of a tool invocation.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ConversationNewParamsItemMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// Any of "mcp_approval_request".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMcpApprovalRequest) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMcpApprovalRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMcpApprovalRequest](
		"type", "mcp_approval_request",
	)
}

// A response to an MCP approval request.
//
// The properties ApprovalRequestID, Approve are required.
type ConversationNewParamsItemMcpApprovalResponse struct {
	ApprovalRequestID string            `json:"approval_request_id,required"`
	Approve           bool              `json:"approve,required"`
	ID                param.Opt[string] `json:"id,omitzero"`
	Reason            param.Opt[string] `json:"reason,omitzero"`
	// Any of "mcp_approval_response".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMcpApprovalResponse) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMcpApprovalResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMcpApprovalResponse](
		"type", "mcp_approval_response",
	)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
//
// The properties ID, Arguments, Name, ServerLabel are required.
type ConversationNewParamsItemMcpCall struct {
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

func (r ConversationNewParamsItemMcpCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMcpCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMcpCall](
		"type", "mcp_call",
	)
}

// MCP list tools output message containing available tools from an MCP server.
//
// The properties ID, ServerLabel, Tools are required.
type ConversationNewParamsItemMcpListTools struct {
	ID          string                                      `json:"id,required"`
	ServerLabel string                                      `json:"server_label,required"`
	Tools       []ConversationNewParamsItemMcpListToolsTool `json:"tools,omitzero,required"`
	// Any of "mcp_list_tools".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMcpListTools) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMcpListTools
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationNewParamsItemMcpListTools](
		"type", "mcp_list_tools",
	)
}

// Tool definition returned by MCP list tools operation.
//
// The properties InputSchema, Name are required.
type ConversationNewParamsItemMcpListToolsTool struct {
	InputSchema map[string]any    `json:"input_schema,omitzero,required"`
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ConversationNewParamsItemMcpListToolsTool) MarshalJSON() (data []byte, err error) {
	type shadow ConversationNewParamsItemMcpListToolsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationNewParamsItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationUpdateParams struct {
	Metadata map[string]string `json:"metadata,omitzero,required"`
	paramObj
}

func (r ConversationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConversationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
