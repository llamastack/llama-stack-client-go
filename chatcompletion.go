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
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
)

// ChatCompletionService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatCompletionService] method instead.
type ChatCompletionService struct {
	Options []option.RequestOption
}

// NewChatCompletionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatCompletionService(opts ...option.RequestOption) (r ChatCompletionService) {
	r = ChatCompletionService{}
	r.Options = opts
	return
}

// Create chat completions.
//
// Generate an OpenAI-compatible chat completion for the given messages using the
// specified model.
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create chat completions.
//
// Generate an OpenAI-compatible chat completion for the given messages using the
// specified model.
func (r *ChatCompletionService) NewStreaming(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append(opts, option.WithJSONSet("stream", true))
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

// Get chat completion.
//
// Describe a chat completion by its ID.
func (r *ChatCompletionService) Get(ctx context.Context, completionID string, opts ...option.RequestOption) (res *ChatCompletionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if completionID == "" {
		err = errors.New("missing required completion_id parameter")
		return
	}
	path := fmt.Sprintf("v1/chat/completions/%s", completionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List chat completions.
func (r *ChatCompletionService) List(ctx context.Context, query ChatCompletionListParams, opts ...option.RequestOption) (res *ChatCompletionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response from an OpenAI-compatible chat completion request.
type ChatCompletionNewResponse struct {
	ID      string                            `json:"id,required"`
	Choices []ChatCompletionNewResponseChoice `json:"choices,required"`
	Created int64                             `json:"created,required"`
	Model   string                            `json:"model,required"`
	// Any of "chat.completion".
	Object ChatCompletionNewResponseObject `json:"object"`
	// Usage information for OpenAI chat completion.
	Usage ChatCompletionNewResponseUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type ChatCompletionNewResponseChoice struct {
	FinishReason string `json:"finish_reason,required"`
	Index        int64  `json:"index,required"`
	// A message from the user in an OpenAI-compatible chat completion request.
	Message ChatCompletionNewResponseChoiceMessageUnion `json:"message,required"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs ChatCompletionNewResponseChoiceLogprobs `json:"logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Message      respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageUnion contains all possible properties and
// values from [ChatCompletionNewResponseChoiceMessageUser],
// [ChatCompletionNewResponseChoiceMessageSystem],
// [ChatCompletionNewResponseChoiceMessageAssistant],
// [ChatCompletionNewResponseChoiceMessageTool],
// [ChatCompletionNewResponseChoiceMessageDeveloper].
//
// Use the [ChatCompletionNewResponseChoiceMessageUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseChoiceMessageUnion struct {
	// This field is a union of
	// [ChatCompletionNewResponseChoiceMessageUserContentUnion],
	// [ChatCompletionNewResponseChoiceMessageSystemContentUnion],
	// [ChatCompletionNewResponseChoiceMessageAssistantContentUnion],
	// [ChatCompletionNewResponseChoiceMessageToolContentUnion],
	// [ChatCompletionNewResponseChoiceMessageDeveloperContentUnion]
	Content ChatCompletionNewResponseChoiceMessageUnionContent `json:"content"`
	Name    string                                             `json:"name"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	// This field is from variant [ChatCompletionNewResponseChoiceMessageAssistant].
	ToolCalls []ChatCompletionNewResponseChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionNewResponseChoiceMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Name       respjson.Field
		Role       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionNewResponseChoiceMessage is implemented by each variant of
// [ChatCompletionNewResponseChoiceMessageUnion] to add type safety for the return
// type of [ChatCompletionNewResponseChoiceMessageUnion.AsAny]
type anyChatCompletionNewResponseChoiceMessage interface {
	implChatCompletionNewResponseChoiceMessageUnion()
}

func (ChatCompletionNewResponseChoiceMessageUser) implChatCompletionNewResponseChoiceMessageUnion() {}
func (ChatCompletionNewResponseChoiceMessageSystem) implChatCompletionNewResponseChoiceMessageUnion() {
}
func (ChatCompletionNewResponseChoiceMessageAssistant) implChatCompletionNewResponseChoiceMessageUnion() {
}
func (ChatCompletionNewResponseChoiceMessageTool) implChatCompletionNewResponseChoiceMessageUnion() {}
func (ChatCompletionNewResponseChoiceMessageDeveloper) implChatCompletionNewResponseChoiceMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseChoiceMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageUser:
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageSystem:
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageAssistant:
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageTool:
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseChoiceMessageUnion) AsAny() anyChatCompletionNewResponseChoiceMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u ChatCompletionNewResponseChoiceMessageUnion) AsUser() (v ChatCompletionNewResponseChoiceMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUnion) AsSystem() (v ChatCompletionNewResponseChoiceMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUnion) AsAssistant() (v ChatCompletionNewResponseChoiceMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUnion) AsTool() (v ChatCompletionNewResponseChoiceMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUnion) AsDeveloper() (v ChatCompletionNewResponseChoiceMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionNewResponseChoiceMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageUnionContent is an implicit subunion of
// [ChatCompletionNewResponseChoiceMessageUnion].
// ChatCompletionNewResponseChoiceMessageUnionContent provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionNewResponseChoiceMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
// OfListOpenAIChatCompletionContentPartText]
type ChatCompletionNewResponseChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		OfListOpenAIChatCompletionContentPartText                                                    respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (r *ChatCompletionNewResponseChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionNewResponseChoiceMessageUser struct {
	Content ChatCompletionNewResponseChoiceMessageUserContentUnion `json:"content,required"`
	Name    string                                                 `json:"name,nullable"`
	// Any of "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile]
type ChatCompletionNewResponseChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (u ChatCompletionNewResponseChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUserContentUnion) AsListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFile() (v []ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionNewResponseChoiceMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion
// contains all possible properties and values from
// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText],
// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL],
// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
//
// Use the
// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	// This field is from variant
	// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL].
	ImageURL ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
	File ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem
// is implemented by each variant of
// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
// to add type safety for the return type of
// [ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
type anyChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem interface {
	implChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion()
}

func (ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) implChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) implChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) implChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText:
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL:
//	case llamastackclient.ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsAny() anyChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsText() (v ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsImageURL() (v ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsFile() (v ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
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
func (r ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification for OpenAI-compatible chat completion messages.
	ImageURL ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Any of "image_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	URL    string `json:"url,required"`
	Detail string `json:"detail,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	File ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Any of "file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	Filename string `json:"filename,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionNewResponseChoiceMessageSystem struct {
	Content ChatCompletionNewResponseChoiceMessageSystemContentUnion `json:"content,required"`
	Name    string                                                   `json:"name,nullable"`
	// Any of "system".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageSystemContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionNewResponseChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionNewResponseChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageSystemContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionNewResponseChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionNewResponseChoiceMessageAssistant struct {
	Content ChatCompletionNewResponseChoiceMessageAssistantContentUnion `json:"content,nullable"`
	Name    string                                                      `json:"name,nullable"`
	// Any of "assistant".
	Role      string                                                    `json:"role"`
	ToolCalls []ChatCompletionNewResponseChoiceMessageAssistantToolCall `json:"tool_calls,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionNewResponseChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionNewResponseChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageAssistantContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseChoiceMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionNewResponseChoiceMessageAssistantToolCall struct {
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionNewResponseChoiceMessageAssistantToolCallFunction `json:"function,nullable"`
	Index    int64                                                           `json:"index,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionNewResponseChoiceMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments,nullable"`
	Name      string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionNewResponseChoiceMessageTool struct {
	Content    ChatCompletionNewResponseChoiceMessageToolContentUnion `json:"content,required"`
	ToolCallID string                                                 `json:"tool_call_id,required"`
	// Any of "tool".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ToolCallID  respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionNewResponseChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionNewResponseChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageToolContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionNewResponseChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionNewResponseChoiceMessageDeveloper struct {
	Content ChatCompletionNewResponseChoiceMessageDeveloperContentUnion `json:"content,required"`
	Name    string                                                      `json:"name,nullable"`
	// Any of "developer".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseChoiceMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionNewResponseChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionNewResponseChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseChoiceMessageDeveloperContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseChoiceMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseChoiceMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionNewResponseChoiceLogprobs struct {
	Content []ChatCompletionNewResponseChoiceLogprobsContent `json:"content,nullable"`
	Refusal []ChatCompletionNewResponseChoiceLogprobsRefusal `json:"refusal,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ChatCompletionNewResponseChoiceLogprobsContent struct {
	Token       string                                                     `json:"token,required"`
	Logprob     float64                                                    `json:"logprob,required"`
	Bytes       []int64                                                    `json:"bytes,nullable"`
	TopLogprobs []ChatCompletionNewResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionNewResponseChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ChatCompletionNewResponseChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionNewResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ChatCompletionNewResponseChoiceLogprobsRefusal struct {
	Token       string                                                     `json:"token,required"`
	Logprob     float64                                                    `json:"logprob,required"`
	Bytes       []int64                                                    `json:"bytes,nullable"`
	TopLogprobs []ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionNewResponseChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob struct {
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
func (r ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseObject string

const (
	ChatCompletionNewResponseObjectChatCompletion ChatCompletionNewResponseObject = "chat.completion"
)

// Usage information for OpenAI chat completion.
type ChatCompletionNewResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens,required"`
	PromptTokens     int64 `json:"prompt_tokens,required"`
	TotalTokens      int64 `json:"total_tokens,required"`
	// Token details for output tokens in OpenAI chat completion usage.
	CompletionTokensDetails ChatCompletionNewResponseUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	// Token details for prompt tokens in OpenAI chat completion usage.
	PromptTokensDetails ChatCompletionNewResponseUsagePromptTokensDetails `json:"prompt_tokens_details,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		PromptTokens            respjson.Field
		TotalTokens             respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokensDetails     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for output tokens in OpenAI chat completion usage.
type ChatCompletionNewResponseUsageCompletionTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionNewResponseUsagePromptTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponse struct {
	ID            string                                       `json:"id,required"`
	Choices       []ChatCompletionGetResponseChoice            `json:"choices,required"`
	Created       int64                                        `json:"created,required"`
	InputMessages []ChatCompletionGetResponseInputMessageUnion `json:"input_messages,required"`
	Model         string                                       `json:"model,required"`
	// Any of "chat.completion".
	Object ChatCompletionGetResponseObject `json:"object"`
	// Usage information for OpenAI chat completion.
	Usage ChatCompletionGetResponseUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Choices       respjson.Field
		Created       respjson.Field
		InputMessages respjson.Field
		Model         respjson.Field
		Object        respjson.Field
		Usage         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type ChatCompletionGetResponseChoice struct {
	FinishReason string `json:"finish_reason,required"`
	Index        int64  `json:"index,required"`
	// A message from the user in an OpenAI-compatible chat completion request.
	Message ChatCompletionGetResponseChoiceMessageUnion `json:"message,required"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs ChatCompletionGetResponseChoiceLogprobs `json:"logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Message      respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageUnion contains all possible properties and
// values from [ChatCompletionGetResponseChoiceMessageUser],
// [ChatCompletionGetResponseChoiceMessageSystem],
// [ChatCompletionGetResponseChoiceMessageAssistant],
// [ChatCompletionGetResponseChoiceMessageTool],
// [ChatCompletionGetResponseChoiceMessageDeveloper].
//
// Use the [ChatCompletionGetResponseChoiceMessageUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageUnion struct {
	// This field is a union of
	// [ChatCompletionGetResponseChoiceMessageUserContentUnion],
	// [ChatCompletionGetResponseChoiceMessageSystemContentUnion],
	// [ChatCompletionGetResponseChoiceMessageAssistantContentUnion],
	// [ChatCompletionGetResponseChoiceMessageToolContentUnion],
	// [ChatCompletionGetResponseChoiceMessageDeveloperContentUnion]
	Content ChatCompletionGetResponseChoiceMessageUnionContent `json:"content"`
	Name    string                                             `json:"name"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	// This field is from variant [ChatCompletionGetResponseChoiceMessageAssistant].
	ToolCalls []ChatCompletionGetResponseChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionGetResponseChoiceMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Name       respjson.Field
		Role       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessage is implemented by each variant of
// [ChatCompletionGetResponseChoiceMessageUnion] to add type safety for the return
// type of [ChatCompletionGetResponseChoiceMessageUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessage interface {
	implChatCompletionGetResponseChoiceMessageUnion()
}

func (ChatCompletionGetResponseChoiceMessageUser) implChatCompletionGetResponseChoiceMessageUnion() {}
func (ChatCompletionGetResponseChoiceMessageSystem) implChatCompletionGetResponseChoiceMessageUnion() {
}
func (ChatCompletionGetResponseChoiceMessageAssistant) implChatCompletionGetResponseChoiceMessageUnion() {
}
func (ChatCompletionGetResponseChoiceMessageTool) implChatCompletionGetResponseChoiceMessageUnion() {}
func (ChatCompletionGetResponseChoiceMessageDeveloper) implChatCompletionGetResponseChoiceMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUser:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageSystem:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageAssistant:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageTool:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageUnion) AsAny() anyChatCompletionGetResponseChoiceMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u ChatCompletionGetResponseChoiceMessageUnion) AsUser() (v ChatCompletionGetResponseChoiceMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUnion) AsSystem() (v ChatCompletionGetResponseChoiceMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUnion) AsAssistant() (v ChatCompletionGetResponseChoiceMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUnion) AsTool() (v ChatCompletionGetResponseChoiceMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUnion) AsDeveloper() (v ChatCompletionGetResponseChoiceMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageUnionContent is an implicit subunion of
// [ChatCompletionGetResponseChoiceMessageUnion].
// ChatCompletionGetResponseChoiceMessageUnionContent provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionGetResponseChoiceMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
// OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		OfListOpenAIChatCompletionContentPartText                                                    respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (r *ChatCompletionGetResponseChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseChoiceMessageUser struct {
	Content ChatCompletionGetResponseChoiceMessageUserContentUnion `json:"content,required"`
	Name    string                                                 `json:"name,nullable"`
	// Any of "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile]
type ChatCompletionGetResponseChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUserContentUnion) AsListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFile() (v []ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion
// contains all possible properties and values from
// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText],
// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL],
// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
//
// Use the
// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
	File ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem
// is implemented by each variant of
// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
// to add type safety for the return type of
// [ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem interface {
	implChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion()
}

func (ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) implChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) implChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) implChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsText() (v ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsImageURL() (v ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsFile() (v ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
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
func (r ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification for OpenAI-compatible chat completion messages.
	ImageURL ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Any of "image_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	URL    string `json:"url,required"`
	Detail string `json:"detail,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	File ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Any of "file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	Filename string `json:"filename,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionGetResponseChoiceMessageSystem struct {
	Content ChatCompletionGetResponseChoiceMessageSystemContentUnion `json:"content,required"`
	Name    string                                                   `json:"name,nullable"`
	// Any of "system".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageSystemContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseChoiceMessageAssistant struct {
	Content ChatCompletionGetResponseChoiceMessageAssistantContentUnion `json:"content,nullable"`
	Name    string                                                      `json:"name,nullable"`
	// Any of "assistant".
	Role      string                                                    `json:"role"`
	ToolCalls []ChatCompletionGetResponseChoiceMessageAssistantToolCall `json:"tool_calls,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageAssistantContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionGetResponseChoiceMessageAssistantToolCall struct {
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction `json:"function,nullable"`
	Index    int64                                                           `json:"index,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments,nullable"`
	Name      string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseChoiceMessageTool struct {
	Content    ChatCompletionGetResponseChoiceMessageToolContentUnion `json:"content,required"`
	ToolCallID string                                                 `json:"tool_call_id,required"`
	// Any of "tool".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ToolCallID  respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseChoiceMessageDeveloper struct {
	Content ChatCompletionGetResponseChoiceMessageDeveloperContentUnion `json:"content,required"`
	Name    string                                                      `json:"name,nullable"`
	// Any of "developer".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionGetResponseChoiceLogprobs struct {
	Content []ChatCompletionGetResponseChoiceLogprobsContent `json:"content,nullable"`
	Refusal []ChatCompletionGetResponseChoiceLogprobsRefusal `json:"refusal,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ChatCompletionGetResponseChoiceLogprobsContent struct {
	Token       string                                                     `json:"token,required"`
	Logprob     float64                                                    `json:"logprob,required"`
	Bytes       []int64                                                    `json:"bytes,nullable"`
	TopLogprobs []ChatCompletionGetResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionGetResponseChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ChatCompletionGetResponseChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionGetResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ChatCompletionGetResponseChoiceLogprobsRefusal struct {
	Token       string                                                     `json:"token,required"`
	Logprob     float64                                                    `json:"logprob,required"`
	Bytes       []int64                                                    `json:"bytes,nullable"`
	TopLogprobs []ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionGetResponseChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob struct {
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
func (r ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageUnion contains all possible properties and
// values from [ChatCompletionGetResponseInputMessageUser],
// [ChatCompletionGetResponseInputMessageSystem],
// [ChatCompletionGetResponseInputMessageAssistant],
// [ChatCompletionGetResponseInputMessageTool],
// [ChatCompletionGetResponseInputMessageDeveloper].
//
// Use the [ChatCompletionGetResponseInputMessageUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageUnion struct {
	// This field is a union of
	// [ChatCompletionGetResponseInputMessageUserContentUnion],
	// [ChatCompletionGetResponseInputMessageSystemContentUnion],
	// [ChatCompletionGetResponseInputMessageAssistantContentUnion],
	// [ChatCompletionGetResponseInputMessageToolContentUnion],
	// [ChatCompletionGetResponseInputMessageDeveloperContentUnion]
	Content ChatCompletionGetResponseInputMessageUnionContent `json:"content"`
	Name    string                                            `json:"name"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	// This field is from variant [ChatCompletionGetResponseInputMessageAssistant].
	ToolCalls []ChatCompletionGetResponseInputMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionGetResponseInputMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Name       respjson.Field
		Role       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessage is implemented by each variant of
// [ChatCompletionGetResponseInputMessageUnion] to add type safety for the return
// type of [ChatCompletionGetResponseInputMessageUnion.AsAny]
type anyChatCompletionGetResponseInputMessage interface {
	implChatCompletionGetResponseInputMessageUnion()
}

func (ChatCompletionGetResponseInputMessageUser) implChatCompletionGetResponseInputMessageUnion()   {}
func (ChatCompletionGetResponseInputMessageSystem) implChatCompletionGetResponseInputMessageUnion() {}
func (ChatCompletionGetResponseInputMessageAssistant) implChatCompletionGetResponseInputMessageUnion() {
}
func (ChatCompletionGetResponseInputMessageTool) implChatCompletionGetResponseInputMessageUnion() {}
func (ChatCompletionGetResponseInputMessageDeveloper) implChatCompletionGetResponseInputMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageUser:
//	case llamastackclient.ChatCompletionGetResponseInputMessageSystem:
//	case llamastackclient.ChatCompletionGetResponseInputMessageAssistant:
//	case llamastackclient.ChatCompletionGetResponseInputMessageTool:
//	case llamastackclient.ChatCompletionGetResponseInputMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageUnion) AsAny() anyChatCompletionGetResponseInputMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u ChatCompletionGetResponseInputMessageUnion) AsUser() (v ChatCompletionGetResponseInputMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUnion) AsSystem() (v ChatCompletionGetResponseInputMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUnion) AsAssistant() (v ChatCompletionGetResponseInputMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUnion) AsTool() (v ChatCompletionGetResponseInputMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUnion) AsDeveloper() (v ChatCompletionGetResponseInputMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageUnionContent is an implicit subunion of
// [ChatCompletionGetResponseInputMessageUnion].
// ChatCompletionGetResponseInputMessageUnionContent provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionGetResponseInputMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
// OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		OfListOpenAIChatCompletionContentPartText                                                    respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (r *ChatCompletionGetResponseInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseInputMessageUser struct {
	Content ChatCompletionGetResponseInputMessageUserContentUnion `json:"content,required"`
	Name    string                                                `json:"name,nullable"`
	// Any of "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile]
type ChatCompletionGetResponseInputMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUserContentUnion) AsListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFile() (v []ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion
// contains all possible properties and values from
// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText],
// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL],
// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
//
// Use the
// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
	File ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem
// is implemented by each variant of
// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
// to add type safety for the return type of
// [ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
type anyChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem interface {
	implChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion()
}

func (ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) implChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) implChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) implChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL:
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsAny() anyChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsText() (v ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsImageURL() (v ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsFile() (v ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
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
func (r ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification for OpenAI-compatible chat completion messages.
	ImageURL ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Any of "image_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	URL    string `json:"url,required"`
	Detail string `json:"detail,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	File ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Any of "file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	Filename string `json:"filename,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionGetResponseInputMessageSystem struct {
	Content ChatCompletionGetResponseInputMessageSystemContentUnion `json:"content,required"`
	Name    string                                                  `json:"name,nullable"`
	// Any of "system".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageSystemContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseInputMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageSystemContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseInputMessageAssistant struct {
	Content ChatCompletionGetResponseInputMessageAssistantContentUnion `json:"content,nullable"`
	Name    string                                                     `json:"name,nullable"`
	// Any of "assistant".
	Role      string                                                   `json:"role"`
	ToolCalls []ChatCompletionGetResponseInputMessageAssistantToolCall `json:"tool_calls,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageAssistantContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseInputMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageAssistantContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionGetResponseInputMessageAssistantToolCall struct {
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionGetResponseInputMessageAssistantToolCallFunction `json:"function,nullable"`
	Index    int64                                                          `json:"index,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionGetResponseInputMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments,nullable"`
	Name      string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseInputMessageTool struct {
	Content    ChatCompletionGetResponseInputMessageToolContentUnion `json:"content,required"`
	ToolCallID string                                                `json:"tool_call_id,required"`
	// Any of "tool".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ToolCallID  respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseInputMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageToolContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseInputMessageDeveloper struct {
	Content ChatCompletionGetResponseInputMessageDeveloperContentUnion `json:"content,required"`
	Name    string                                                     `json:"name,nullable"`
	// Any of "developer".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageDeveloperContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionGetResponseInputMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageDeveloperContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseObject string

const (
	ChatCompletionGetResponseObjectChatCompletion ChatCompletionGetResponseObject = "chat.completion"
)

// Usage information for OpenAI chat completion.
type ChatCompletionGetResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens,required"`
	PromptTokens     int64 `json:"prompt_tokens,required"`
	TotalTokens      int64 `json:"total_tokens,required"`
	// Token details for output tokens in OpenAI chat completion usage.
	CompletionTokensDetails ChatCompletionGetResponseUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	// Token details for prompt tokens in OpenAI chat completion usage.
	PromptTokensDetails ChatCompletionGetResponseUsagePromptTokensDetails `json:"prompt_tokens_details,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		PromptTokens            respjson.Field
		TotalTokens             respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokensDetails     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for output tokens in OpenAI chat completion usage.
type ChatCompletionGetResponseUsageCompletionTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionGetResponseUsagePromptTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from listing OpenAI-compatible chat completions.
type ChatCompletionListResponse struct {
	Data    []ChatCompletionListResponseData `json:"data,required"`
	FirstID string                           `json:"first_id,required"`
	HasMore bool                             `json:"has_more,required"`
	LastID  string                           `json:"last_id,required"`
	// Any of "list".
	Object ChatCompletionListResponseObject `json:"object"`
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
func (r ChatCompletionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseData struct {
	ID            string                                            `json:"id,required"`
	Choices       []ChatCompletionListResponseDataChoice            `json:"choices,required"`
	Created       int64                                             `json:"created,required"`
	InputMessages []ChatCompletionListResponseDataInputMessageUnion `json:"input_messages,required"`
	Model         string                                            `json:"model,required"`
	// Any of "chat.completion".
	Object string `json:"object"`
	// Usage information for OpenAI chat completion.
	Usage ChatCompletionListResponseDataUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Choices       respjson.Field
		Created       respjson.Field
		InputMessages respjson.Field
		Model         respjson.Field
		Object        respjson.Field
		Usage         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type ChatCompletionListResponseDataChoice struct {
	FinishReason string `json:"finish_reason,required"`
	Index        int64  `json:"index,required"`
	// A message from the user in an OpenAI-compatible chat completion request.
	Message ChatCompletionListResponseDataChoiceMessageUnion `json:"message,required"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs ChatCompletionListResponseDataChoiceLogprobs `json:"logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Message      respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageUnion contains all possible
// properties and values from [ChatCompletionListResponseDataChoiceMessageUser],
// [ChatCompletionListResponseDataChoiceMessageSystem],
// [ChatCompletionListResponseDataChoiceMessageAssistant],
// [ChatCompletionListResponseDataChoiceMessageTool],
// [ChatCompletionListResponseDataChoiceMessageDeveloper].
//
// Use the [ChatCompletionListResponseDataChoiceMessageUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageUnion struct {
	// This field is a union of
	// [ChatCompletionListResponseDataChoiceMessageUserContentUnion],
	// [ChatCompletionListResponseDataChoiceMessageSystemContentUnion],
	// [ChatCompletionListResponseDataChoiceMessageAssistantContentUnion],
	// [ChatCompletionListResponseDataChoiceMessageToolContentUnion],
	// [ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion]
	Content ChatCompletionListResponseDataChoiceMessageUnionContent `json:"content"`
	Name    string                                                  `json:"name"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageAssistant].
	ToolCalls []ChatCompletionListResponseDataChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionListResponseDataChoiceMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Name       respjson.Field
		Role       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessage is implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageUnion] to add type safety for the
// return type of [ChatCompletionListResponseDataChoiceMessageUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessage interface {
	implChatCompletionListResponseDataChoiceMessageUnion()
}

func (ChatCompletionListResponseDataChoiceMessageUser) implChatCompletionListResponseDataChoiceMessageUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageSystem) implChatCompletionListResponseDataChoiceMessageUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageAssistant) implChatCompletionListResponseDataChoiceMessageUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageTool) implChatCompletionListResponseDataChoiceMessageUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageDeveloper) implChatCompletionListResponseDataChoiceMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageUser:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageSystem:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageAssistant:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageTool:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageUnion) AsAny() anyChatCompletionListResponseDataChoiceMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageUnion) AsUser() (v ChatCompletionListResponseDataChoiceMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUnion) AsSystem() (v ChatCompletionListResponseDataChoiceMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUnion) AsAssistant() (v ChatCompletionListResponseDataChoiceMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUnion) AsTool() (v ChatCompletionListResponseDataChoiceMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUnion) AsDeveloper() (v ChatCompletionListResponseDataChoiceMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseDataChoiceMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageUnionContent is an implicit subunion
// of [ChatCompletionListResponseDataChoiceMessageUnion].
// ChatCompletionListResponseDataChoiceMessageUnionContent provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionListResponseDataChoiceMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
// OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		OfListOpenAIChatCompletionContentPartText                                                    respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (r *ChatCompletionListResponseDataChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataChoiceMessageUser struct {
	Content ChatCompletionListResponseDataChoiceMessageUserContentUnion `json:"content,required"`
	Name    string                                                      `json:"name,nullable"`
	// Any of "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageUserContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile]
type ChatCompletionListResponseDataChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentUnion) AsListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFile() (v []ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageUserContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion
// contains all possible properties and values from
// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText],
// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL],
// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
//
// Use the
// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL].
	ImageURL ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
	File ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem
// is implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
// to add type safety for the return type of
// [ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem interface {
	implChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion()
}

func (ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) implChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) implChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) implChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsAny() anyChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsText() (v ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsImageURL() (v ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsFile() (v ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
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
func (r ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification for OpenAI-compatible chat completion messages.
	ImageURL ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Any of "image_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	URL    string `json:"url,required"`
	Detail string `json:"detail,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	File ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Any of "file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	Filename string `json:"filename,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionListResponseDataChoiceMessageSystem struct {
	Content ChatCompletionListResponseDataChoiceMessageSystemContentUnion `json:"content,required"`
	Name    string                                                        `json:"name,nullable"`
	// Any of "system".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageSystemContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageSystemContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageSystemContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataChoiceMessageAssistant struct {
	Content ChatCompletionListResponseDataChoiceMessageAssistantContentUnion `json:"content,nullable"`
	Name    string                                                           `json:"name,nullable"`
	// Any of "assistant".
	Role      string                                                         `json:"role"`
	ToolCalls []ChatCompletionListResponseDataChoiceMessageAssistantToolCall `json:"tool_calls,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageAssistantContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionListResponseDataChoiceMessageAssistantToolCall struct {
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionListResponseDataChoiceMessageAssistantToolCallFunction `json:"function,nullable"`
	Index    int64                                                                `json:"index,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageAssistantToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionListResponseDataChoiceMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments,nullable"`
	Name      string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataChoiceMessageTool struct {
	Content    ChatCompletionListResponseDataChoiceMessageToolContentUnion `json:"content,required"`
	ToolCallID string                                                      `json:"tool_call_id,required"`
	// Any of "tool".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ToolCallID  respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageToolContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageToolContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageToolContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataChoiceMessageDeveloper struct {
	Content ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion `json:"content,required"`
	Name    string                                                           `json:"name,nullable"`
	// Any of "developer".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionListResponseDataChoiceLogprobs struct {
	Content []ChatCompletionListResponseDataChoiceLogprobsContent `json:"content,nullable"`
	Refusal []ChatCompletionListResponseDataChoiceLogprobsRefusal `json:"refusal,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ChatCompletionListResponseDataChoiceLogprobsContent struct {
	Token       string                                                          `json:"token,required"`
	Logprob     float64                                                         `json:"logprob,required"`
	Bytes       []int64                                                         `json:"bytes,nullable"`
	TopLogprobs []ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionListResponseDataChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type ChatCompletionListResponseDataChoiceLogprobsRefusal struct {
	Token       string                                                          `json:"token,required"`
	Logprob     float64                                                         `json:"logprob,required"`
	Bytes       []int64                                                         `json:"bytes,nullable"`
	TopLogprobs []ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionListResponseDataChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob struct {
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
func (r ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageUnion contains all possible properties
// and values from [ChatCompletionListResponseDataInputMessageUser],
// [ChatCompletionListResponseDataInputMessageSystem],
// [ChatCompletionListResponseDataInputMessageAssistant],
// [ChatCompletionListResponseDataInputMessageTool],
// [ChatCompletionListResponseDataInputMessageDeveloper].
//
// Use the [ChatCompletionListResponseDataInputMessageUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageUnion struct {
	// This field is a union of
	// [ChatCompletionListResponseDataInputMessageUserContentUnion],
	// [ChatCompletionListResponseDataInputMessageSystemContentUnion],
	// [ChatCompletionListResponseDataInputMessageAssistantContentUnion],
	// [ChatCompletionListResponseDataInputMessageToolContentUnion],
	// [ChatCompletionListResponseDataInputMessageDeveloperContentUnion]
	Content ChatCompletionListResponseDataInputMessageUnionContent `json:"content"`
	Name    string                                                 `json:"name"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageAssistant].
	ToolCalls []ChatCompletionListResponseDataInputMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionListResponseDataInputMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Name       respjson.Field
		Role       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessage is implemented by each variant of
// [ChatCompletionListResponseDataInputMessageUnion] to add type safety for the
// return type of [ChatCompletionListResponseDataInputMessageUnion.AsAny]
type anyChatCompletionListResponseDataInputMessage interface {
	implChatCompletionListResponseDataInputMessageUnion()
}

func (ChatCompletionListResponseDataInputMessageUser) implChatCompletionListResponseDataInputMessageUnion() {
}
func (ChatCompletionListResponseDataInputMessageSystem) implChatCompletionListResponseDataInputMessageUnion() {
}
func (ChatCompletionListResponseDataInputMessageAssistant) implChatCompletionListResponseDataInputMessageUnion() {
}
func (ChatCompletionListResponseDataInputMessageTool) implChatCompletionListResponseDataInputMessageUnion() {
}
func (ChatCompletionListResponseDataInputMessageDeveloper) implChatCompletionListResponseDataInputMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageUser:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageSystem:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageAssistant:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageTool:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageUnion) AsAny() anyChatCompletionListResponseDataInputMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageUnion) AsUser() (v ChatCompletionListResponseDataInputMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUnion) AsSystem() (v ChatCompletionListResponseDataInputMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUnion) AsAssistant() (v ChatCompletionListResponseDataInputMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUnion) AsTool() (v ChatCompletionListResponseDataInputMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUnion) AsDeveloper() (v ChatCompletionListResponseDataInputMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseDataInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageUnionContent is an implicit subunion
// of [ChatCompletionListResponseDataInputMessageUnion].
// ChatCompletionListResponseDataInputMessageUnionContent provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionListResponseDataInputMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
// OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		OfListOpenAIChatCompletionContentPartText                                                    respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (r *ChatCompletionListResponseDataInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataInputMessageUser struct {
	Content ChatCompletionListResponseDataInputMessageUserContentUnion `json:"content,required"`
	Name    string                                                     `json:"name,nullable"`
	// Any of "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile]
type ChatCompletionListResponseDataInputMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUserContentUnion) AsListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFile() (v []ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageUserContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion
// contains all possible properties and values from
// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText],
// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL],
// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
//
// Use the
// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL].
	ImageURL ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile].
	File ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem
// is implemented by each variant of
// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion]
// to add type safety for the return type of
// [ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny]
type anyChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem interface {
	implChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion()
}

func (ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) implChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) implChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) implChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsAny() anyChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsText() (v ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsImageURL() (v ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) AsFile() (v ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
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
func (r ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification for OpenAI-compatible chat completion messages.
	ImageURL ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Any of "image_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	URL    string `json:"url,required"`
	Detail string `json:"detail,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	File ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Any of "file".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	FileData string `json:"file_data,nullable"`
	FileID   string `json:"file_id,nullable"`
	Filename string `json:"filename,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionListResponseDataInputMessageSystem struct {
	Content ChatCompletionListResponseDataInputMessageSystemContentUnion `json:"content,required"`
	Name    string                                                       `json:"name,nullable"`
	// Any of "system".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageSystemContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataInputMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageSystemContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageSystemContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataInputMessageAssistant struct {
	Content ChatCompletionListResponseDataInputMessageAssistantContentUnion `json:"content,nullable"`
	Name    string                                                          `json:"name,nullable"`
	// Any of "assistant".
	Role      string                                                        `json:"role"`
	ToolCalls []ChatCompletionListResponseDataInputMessageAssistantToolCall `json:"tool_calls,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataInputMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageAssistantContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionListResponseDataInputMessageAssistantToolCall struct {
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionListResponseDataInputMessageAssistantToolCallFunction `json:"function,nullable"`
	Index    int64                                                               `json:"index,nullable"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageAssistantToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionListResponseDataInputMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments,nullable"`
	Name      string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataInputMessageTool struct {
	Content    ChatCompletionListResponseDataInputMessageToolContentUnion `json:"content,required"`
	ToolCallID string                                                     `json:"tool_call_id,required"`
	// Any of "tool".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ToolCallID  respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataInputMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageToolContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageToolContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataInputMessageDeveloper struct {
	Content ChatCompletionListResponseDataInputMessageDeveloperContentUnion `json:"content,required"`
	Name    string                                                          `json:"name,nullable"`
	// Any of "developer".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListOpenAIChatCompletionContentPartText]
type ChatCompletionListResponseDataInputMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
	// instead of an object.
	OfListOpenAIChatCompletionContentPartText []ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfListOpenAIChatCompletionContentPartText respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageDeveloperContentUnion) AsListOpenAIChatCompletionContentPartTextParam() (v []ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
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
func (r ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage information for OpenAI chat completion.
type ChatCompletionListResponseDataUsage struct {
	CompletionTokens int64 `json:"completion_tokens,required"`
	PromptTokens     int64 `json:"prompt_tokens,required"`
	TotalTokens      int64 `json:"total_tokens,required"`
	// Token details for output tokens in OpenAI chat completion usage.
	CompletionTokensDetails ChatCompletionListResponseDataUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	// Token details for prompt tokens in OpenAI chat completion usage.
	PromptTokensDetails ChatCompletionListResponseDataUsagePromptTokensDetails `json:"prompt_tokens_details,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		PromptTokens            respjson.Field
		TotalTokens             respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokensDetails     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataUsage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for output tokens in OpenAI chat completion usage.
type ChatCompletionListResponseDataUsageCompletionTokensDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataUsageCompletionTokensDetails) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionListResponseDataUsagePromptTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseObject string

const (
	ChatCompletionListResponseObjectList ChatCompletionListResponseObject = "list"
)

type ChatCompletionNewParams struct {
	Messages            []ChatCompletionNewParamsMessageUnion    `json:"messages,omitzero,required"`
	Model               string                                   `json:"model,required"`
	FrequencyPenalty    param.Opt[float64]                       `json:"frequency_penalty,omitzero"`
	Logprobs            param.Opt[bool]                          `json:"logprobs,omitzero"`
	MaxCompletionTokens param.Opt[int64]                         `json:"max_completion_tokens,omitzero"`
	MaxTokens           param.Opt[int64]                         `json:"max_tokens,omitzero"`
	N                   param.Opt[int64]                         `json:"n,omitzero"`
	ParallelToolCalls   param.Opt[bool]                          `json:"parallel_tool_calls,omitzero"`
	PresencePenalty     param.Opt[float64]                       `json:"presence_penalty,omitzero"`
	Seed                param.Opt[int64]                         `json:"seed,omitzero"`
	Temperature         param.Opt[float64]                       `json:"temperature,omitzero"`
	TopLogprobs         param.Opt[int64]                         `json:"top_logprobs,omitzero"`
	TopP                param.Opt[float64]                       `json:"top_p,omitzero"`
	User                param.Opt[string]                        `json:"user,omitzero"`
	FunctionCall        ChatCompletionNewParamsFunctionCallUnion `json:"function_call,omitzero"`
	Functions           []map[string]any                         `json:"functions,omitzero"`
	LogitBias           map[string]float64                       `json:"logit_bias,omitzero"`
	// Text response format for OpenAI-compatible chat completion requests.
	ResponseFormat ChatCompletionNewParamsResponseFormatUnion `json:"response_format,omitzero"`
	Stop           ChatCompletionNewParamsStopUnion           `json:"stop,omitzero"`
	StreamOptions  map[string]any                             `json:"stream_options,omitzero"`
	ToolChoice     ChatCompletionNewParamsToolChoiceUnion     `json:"tool_choice,omitzero"`
	Tools          []map[string]any                           `json:"tools,omitzero"`
	paramObj
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageUnion struct {
	OfUser      *ChatCompletionNewParamsMessageUser      `json:",omitzero,inline"`
	OfSystem    *ChatCompletionNewParamsMessageSystem    `json:",omitzero,inline"`
	OfAssistant *ChatCompletionNewParamsMessageAssistant `json:",omitzero,inline"`
	OfTool      *ChatCompletionNewParamsMessageTool      `json:",omitzero,inline"`
	OfDeveloper *ChatCompletionNewParamsMessageDeveloper `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser,
		u.OfSystem,
		u.OfAssistant,
		u.OfTool,
		u.OfDeveloper)
}
func (u *ChatCompletionNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfDeveloper) {
		return u.OfDeveloper
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetToolCalls() []ChatCompletionNewParamsMessageAssistantToolCall {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetName() *string {
	if vt := u.OfUser; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSystem; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAssistant; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfDeveloper; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfDeveloper; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ChatCompletionNewParamsMessageUnion) GetContent() (res chatCompletionNewParamsMessageUnionContent) {
	if vt := u.OfUser; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfAssistant; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfTool; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfDeveloper; vt != nil {
		res.any = vt.Content.asAny()
	}
	return
}

// Can have the runtime types [*string],
// [_[]ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion],
// [_[]ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem],
// [_[]ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem],
// [_[]ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem],
// [\*[]ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
type chatCompletionNewParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u chatCompletionNewParamsMessageUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageUnion](
		"role",
		apijson.Discriminator[ChatCompletionNewParamsMessageUser]("user"),
		apijson.Discriminator[ChatCompletionNewParamsMessageSystem]("system"),
		apijson.Discriminator[ChatCompletionNewParamsMessageAssistant]("assistant"),
		apijson.Discriminator[ChatCompletionNewParamsMessageTool]("tool"),
		apijson.Discriminator[ChatCompletionNewParamsMessageDeveloper]("developer"),
	)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The property Content is required.
type ChatCompletionNewParamsMessageUser struct {
	Content ChatCompletionNewParamsMessageUserContentUnion `json:"content,omitzero,required"`
	Name    param.Opt[string]                              `json:"name,omitzero"`
	// Any of "user".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUser) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageUser](
		"role", "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageUserContentUnion struct {
	OfString                                                                                     param.Opt[string]                                                                                                                                   `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile)
}
func (u *ChatCompletionNewParamsMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUserContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile) {
		return &u.OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	OfText     *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText     `json:",omitzero,inline"`
	OfImageURL *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL `json:",omitzero,inline"`
	OfFile     *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetImageURL() *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetFile() *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL]("image_url"),
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile]("file"),
	)
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText](
		"type", "text",
	)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The property ImageURL is required.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification for OpenAI-compatible chat completion messages.
	ImageURL ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Any of "image_url".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL](
		"type", "image_url",
	)
}

// Image URL specification for OpenAI-compatible chat completion messages.
//
// The property URL is required.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property File is required.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	File ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,omitzero,required"`
	// Any of "file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile](
		"type", "file",
	)
}

type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The property Content is required.
type ChatCompletionNewParamsMessageSystem struct {
	Content ChatCompletionNewParamsMessageSystemContentUnion `json:"content,omitzero,required"`
	Name    param.Opt[string]                                `json:"name,omitzero"`
	// Any of "system".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageSystem](
		"role", "system",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageSystemContentUnion struct {
	OfString                                  param.Opt[string]                                                                             `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageSystemContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *ChatCompletionNewParamsMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageSystemContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionNewParamsMessageAssistant struct {
	Name      param.Opt[string]                                   `json:"name,omitzero"`
	Content   ChatCompletionNewParamsMessageAssistantContentUnion `json:"content,omitzero"`
	ToolCalls []ChatCompletionNewParamsMessageAssistantToolCall   `json:"tool_calls,omitzero"`
	// Any of "assistant".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistant) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageAssistant](
		"role", "assistant",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageAssistantContentUnion struct {
	OfString                                  param.Opt[string]                                                                                `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageAssistantContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *ChatCompletionNewParamsMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageAssistantContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionNewParamsMessageAssistantToolCall struct {
	ID    param.Opt[string] `json:"id,omitzero"`
	Index param.Opt[int64]  `json:"index,omitzero"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionNewParamsMessageAssistantToolCallFunction `json:"function,omitzero"`
	// Any of "function".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageAssistantToolCall](
		"type", "function",
	)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionNewParamsMessageAssistantToolCallFunction struct {
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantToolCallFunction) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantToolCallFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
//
// The properties Content, ToolCallID are required.
type ChatCompletionNewParamsMessageTool struct {
	Content    ChatCompletionNewParamsMessageToolContentUnion `json:"content,omitzero,required"`
	ToolCallID string                                         `json:"tool_call_id,required"`
	// Any of "tool".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageTool) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageTool](
		"role", "tool",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageToolContentUnion struct {
	OfString                                  param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageToolContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *ChatCompletionNewParamsMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageToolContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The property Content is required.
type ChatCompletionNewParamsMessageDeveloper struct {
	Content ChatCompletionNewParamsMessageDeveloperContentUnion `json:"content,omitzero,required"`
	Name    param.Opt[string]                                   `json:"name,omitzero"`
	// Any of "developer".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageDeveloper](
		"role", "developer",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageDeveloperContentUnion struct {
	OfString                                  param.Opt[string]                                                                                `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageDeveloperContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *ChatCompletionNewParamsMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageDeveloperContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsFunctionCallUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsFunctionCallUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *ChatCompletionNewParamsFunctionCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsFunctionCallUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsResponseFormatUnion struct {
	OfText       *ChatCompletionNewParamsResponseFormatText       `json:",omitzero,inline"`
	OfJsonSchema *ChatCompletionNewParamsResponseFormatJsonSchema `json:",omitzero,inline"`
	OfJsonObject *ChatCompletionNewParamsResponseFormatJsonObject `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsResponseFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfJsonSchema, u.OfJsonObject)
}
func (u *ChatCompletionNewParamsResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsResponseFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfJsonSchema) {
		return u.OfJsonSchema
	} else if !param.IsOmitted(u.OfJsonObject) {
		return u.OfJsonObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsResponseFormatUnion) GetJsonSchema() *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema {
	if vt := u.OfJsonSchema; vt != nil {
		return &vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsResponseFormatUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonObject; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsResponseFormatUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsResponseFormatText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsResponseFormatJsonSchema]("json_schema"),
		apijson.Discriminator[ChatCompletionNewParamsResponseFormatJsonObject]("json_object"),
	)
}

// Text response format for OpenAI-compatible chat completion requests.
type ChatCompletionNewParamsResponseFormatText struct {
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsResponseFormatText](
		"type", "text",
	)
}

// JSON schema response format for OpenAI-compatible chat completion requests.
//
// The property JsonSchema is required.
type ChatCompletionNewParamsResponseFormatJsonSchema struct {
	// JSON schema specification for OpenAI-compatible structured response format.
	JsonSchema ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero,required"`
	// Any of "json_schema".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsResponseFormatJsonSchema](
		"type", "json_schema",
	)
}

// JSON schema specification for OpenAI-compatible structured response format.
type ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Strict      param.Opt[bool]   `json:"strict,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Schema      map[string]any    `json:"schema,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON object response format for OpenAI-compatible chat completion requests.
type ChatCompletionNewParamsResponseFormatJsonObject struct {
	// Any of "json_object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonObject) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsResponseFormatJsonObject](
		"type", "json_object",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsStopUnion struct {
	OfString     param.Opt[string] `json:",omitzero,inline"`
	OfListString []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListString)
}
func (u *ChatCompletionNewParamsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsStopUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListString) {
		return &u.OfListString
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsToolChoiceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *ChatCompletionNewParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

type ChatCompletionListParams struct {
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	// Sort order for paginated responses.
	//
	// Any of "asc", "desc".
	Order ChatCompletionListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatCompletionListParams]'s query parameters as
// `url.Values`.
func (r ChatCompletionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for paginated responses.
type ChatCompletionListParamsOrder string

const (
	ChatCompletionListParamsOrderAsc  ChatCompletionListParamsOrder = "asc"
	ChatCompletionListParamsOrderDesc ChatCompletionListParamsOrder = "desc"
)
