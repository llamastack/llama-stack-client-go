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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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

// Generate an OpenAI-compatible chat completion for the given messages using the
// specified model.
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletionNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate an OpenAI-compatible chat completion for the given messages using the
// specified model.
func (r *ChatCompletionService) NewStreaming(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

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

// List all chat completions.
func (r *ChatCompletionService) List(ctx context.Context, query ChatCompletionListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[ChatCompletionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/chat/completions"
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

// List all chat completions.
func (r *ChatCompletionService) ListAutoPaging(ctx context.Context, query ChatCompletionListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ChatCompletionListResponse] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// ChatCompletionNewResponseUnion contains all possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletion], [ChatCompletionChunk].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoice],
	// [[]ChatCompletionChunkChoice]
	Choices ChatCompletionNewResponseUnionChoices `json:"choices"`
	Created int64                                 `json:"created"`
	Model   string                                `json:"model"`
	Object  string                                `json:"object"`
	JSON    struct {
		ID      respjson.Field
		Choices respjson.Field
		Created respjson.Field
		Model   respjson.Field
		Object  respjson.Field
		raw     string
	} `json:"-"`
}

func (u ChatCompletionNewResponseUnion) AsOpenAIChatCompletion() (v ChatCompletionNewResponseOpenAIChatCompletion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseUnion) AsOpenAIChatCompletionChunk() (v ChatCompletionChunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseUnionChoices is an implicit subunion of
// [ChatCompletionNewResponseUnion]. ChatCompletionNewResponseUnionChoices provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionNewResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChatCompletionNewResponseOpenAIChatCompletionChoices
// OfChatCompletionChunkChoices]
type ChatCompletionNewResponseUnionChoices struct {
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoice] instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoices []ChatCompletionNewResponseOpenAIChatCompletionChoice `json:",inline"`
	// This field will be present if the value is a [[]ChatCompletionChunkChoice]
	// instead of an object.
	OfChatCompletionChunkChoices []ChatCompletionChunkChoice `json:",inline"`
	JSON                         struct {
		OfChatCompletionNewResponseOpenAIChatCompletionChoices respjson.Field
		OfChatCompletionChunkChoices                           respjson.Field
		raw                                                    string
	} `json:"-"`
}

func (r *ChatCompletionNewResponseUnionChoices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from an OpenAI-compatible chat completion request.
type ChatCompletionNewResponseOpenAIChatCompletion struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []ChatCompletionNewResponseOpenAIChatCompletionChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created int64 `json:"created,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion"
	Object constant.ChatCompletion `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletion) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseOpenAIChatCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type ChatCompletionNewResponseOpenAIChatCompletionChoice struct {
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// The message from the model
	Message ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion `json:"message,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobs `json:"logprobs"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion contains all
// possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper].
//
// Use the [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion struct {
	// This field is a union of
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion],
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion],
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion],
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion],
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion]
	Content ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnionContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant].
	ToolCalls []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessage is implemented by
// each variant of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion] to add type
// safety for the return type of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion.AsAny]
type anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessage interface {
	implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion()
}

func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessage {
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

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) AsUser() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) AsSystem() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) AsAssistant() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) AsTool() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) AsDeveloper() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnionContent is an
// implicit subunion of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion].
// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnionContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray]
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                                              struct {
		OfString                                                                          respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray      respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray    respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray      respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion
// contains all possible properties and values from [string],
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray]
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile].
//
// Use the
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile].
	File ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItem
// is implemented by each variant of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItem interface {
	implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion()
}

func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItem {
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

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) AsText() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) AsFile() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	// Must be "image_url" to identify this as image content
	Type constant.ImageURL `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail string `json:"detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile struct {
	File ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFileFile `json:"file,required"`
	Type constant.File                                                                          `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFileFile struct {
	FileData string `json:"file_data"`
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role constant.System `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion
// contains all possible properties and values from [string],
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray]
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem `json:",inline"`
	JSON                                                                           struct {
		OfString                                                                       respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion
// contains all possible properties and values from [string],
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray]
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem `json:",inline"`
	JSON                                                                              struct {
		OfString                                                                          respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCallFunction `json:"function"`
	// (Optional) Index of the tool call in the list
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments string `json:"arguments"`
	// (Optional) Name of the function to call
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool struct {
	// The response content from the tool
	Content ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion
// contains all possible properties and values from [string],
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray]
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role constant.Developer `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion
// contains all possible properties and values from [string],
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray]
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                                              struct {
		OfString                                                                          respjson.Field
		OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The log probabilities for the tokens in the message
type ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobs struct {
	// (Optional) The log probabilities for the tokens in the message
	Content []ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContent `json:"content"`
	// (Optional) The log probabilities for the tokens in the message
	Refusal []ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusal `json:"refusal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobs) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContent struct {
	Token       string                                                                         `json:"token,required"`
	Logprob     float64                                                                        `json:"logprob,required"`
	TopLogprobs []ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                                        `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContent) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContentTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContentTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusal struct {
	Token       string                                                                         `json:"token,required"`
	Logprob     float64                                                                        `json:"logprob,required"`
	TopLogprobs []ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                                        `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusal) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusalTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusalTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponse struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []ChatCompletionGetResponseChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created       int64                                        `json:"created,required"`
	InputMessages []ChatCompletionGetResponseInputMessageUnion `json:"input_messages,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion"
	Object constant.ChatCompletion `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Choices       respjson.Field
		Created       respjson.Field
		InputMessages respjson.Field
		Model         respjson.Field
		Object        respjson.Field
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
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// The message from the model
	Message ChatCompletionGetResponseChoiceMessageUnion `json:"message,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChatCompletionGetResponseChoiceLogprobs `json:"logprobs"`
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
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [ChatCompletionGetResponseChoiceMessageAssistant].
	ToolCalls []ChatCompletionGetResponseChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionGetResponseChoiceMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
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
// will be valid: OfString OfChatCompletionGetResponseChoiceMessageUserContentArray
// OfChatCompletionGetResponseChoiceMessageSystemContentArray
// OfChatCompletionGetResponseChoiceMessageAssistantContentArray
// OfChatCompletionGetResponseChoiceMessageToolContentArray
// OfChatCompletionGetResponseChoiceMessageDeveloperContentArray]
type ChatCompletionGetResponseChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageUserContentArray []ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseChoiceMessageSystemContentArray []ChatCompletionGetResponseChoiceMessageSystemContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageAssistantContentArray []ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseChoiceMessageToolContentArray []ChatCompletionGetResponseChoiceMessageToolContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageDeveloperContentArray []ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionGetResponseChoiceMessageUserContentArray      respjson.Field
		OfChatCompletionGetResponseChoiceMessageSystemContentArray    respjson.Field
		OfChatCompletionGetResponseChoiceMessageAssistantContentArray respjson.Field
		OfChatCompletionGetResponseChoiceMessageToolContentArray      respjson.Field
		OfChatCompletionGetResponseChoiceMessageDeveloperContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (r *ChatCompletionGetResponseChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseChoiceMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionGetResponseChoiceMessageUserContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseChoiceMessageUserContentArray]
type ChatCompletionGetResponseChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageUserContentArray []ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                     struct {
		OfString                                                 respjson.Field
		OfChatCompletionGetResponseChoiceMessageUserContentArray respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUserContentUnion) AsChatCompletionGetResponseChoiceMessageUserContentArray() (v []ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemText],
// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL],
// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile].
//
// Use the [ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile].
	File ChatCompletionGetResponseChoiceMessageUserContentArrayItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessageUserContentArrayItem is implemented by
// each variant of
// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion] to add type
// safety for the return type of
// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessageUserContentArrayItem interface {
	implChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion()
}

func (ChatCompletionGetResponseChoiceMessageUserContentArrayItemText) implChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL) implChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile) implChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageUserContentArrayItem {
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

func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) AsText() (v ChatCompletionGetResponseChoiceMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) AsFile() (v ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	// Must be "image_url" to identify this as image content
	Type constant.ImageURL `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail string `json:"detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile struct {
	File ChatCompletionGetResponseChoiceMessageUserContentArrayItemFileFile `json:"file,required"`
	Type constant.File                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageUserContentArrayItemFileFile struct {
	FileData string `json:"file_data"`
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
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
func (r ChatCompletionGetResponseChoiceMessageUserContentArrayItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionGetResponseChoiceMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionGetResponseChoiceMessageSystemContentUnion `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role constant.System `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseChoiceMessageSystemContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseChoiceMessageSystemContentArray]
type ChatCompletionGetResponseChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseChoiceMessageSystemContentArray []ChatCompletionGetResponseChoiceMessageSystemContentArrayItem `json:",inline"`
	JSON                                                       struct {
		OfString                                                   respjson.Field
		OfChatCompletionGetResponseChoiceMessageSystemContentArray respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) AsChatCompletionGetResponseChoiceMessageSystemContentArray() (v []ChatCompletionGetResponseChoiceMessageSystemContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageSystemContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseChoiceMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionGetResponseChoiceMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionGetResponseChoiceMessageAssistantToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Content     respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseChoiceMessageAssistantContentArray]
type ChatCompletionGetResponseChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageAssistantContentArray []ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionGetResponseChoiceMessageAssistantContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageAssistantContentUnion) AsChatCompletionGetResponseChoiceMessageAssistantContentArray() (v []ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem) {
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
type ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionGetResponseChoiceMessageAssistantToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction `json:"function"`
	// (Optional) Index of the tool call in the list
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments string `json:"arguments"`
	// (Optional) Name of the function to call
	Name string `json:"name"`
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
	// The response content from the tool
	Content ChatCompletionGetResponseChoiceMessageToolContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ToolCallID  respjson.Field
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
// [[]ChatCompletionGetResponseChoiceMessageToolContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseChoiceMessageToolContentArray]
type ChatCompletionGetResponseChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseChoiceMessageToolContentArray []ChatCompletionGetResponseChoiceMessageToolContentArrayItem `json:",inline"`
	JSON                                                     struct {
		OfString                                                 respjson.Field
		OfChatCompletionGetResponseChoiceMessageToolContentArray respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) AsChatCompletionGetResponseChoiceMessageToolContentArray() (v []ChatCompletionGetResponseChoiceMessageToolContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseChoiceMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageToolContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseChoiceMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionGetResponseChoiceMessageDeveloperContentUnion `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role constant.Developer `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseChoiceMessageDeveloperContentArray]
type ChatCompletionGetResponseChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageDeveloperContentArray []ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionGetResponseChoiceMessageDeveloperContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) AsChatCompletionGetResponseChoiceMessageDeveloperContentArray() (v []ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem) {
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
type ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The log probabilities for the tokens in the message
type ChatCompletionGetResponseChoiceLogprobs struct {
	// (Optional) The log probabilities for the tokens in the message
	Content []ChatCompletionGetResponseChoiceLogprobsContent `json:"content"`
	// (Optional) The log probabilities for the tokens in the message
	Refusal []ChatCompletionGetResponseChoiceLogprobsRefusal `json:"refusal"`
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
type ChatCompletionGetResponseChoiceLogprobsContent struct {
	Token       string                                                     `json:"token,required"`
	Logprob     float64                                                    `json:"logprob,required"`
	TopLogprobs []ChatCompletionGetResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                    `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
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
type ChatCompletionGetResponseChoiceLogprobsContentTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
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
type ChatCompletionGetResponseChoiceLogprobsRefusal struct {
	Token       string                                                     `json:"token,required"`
	Logprob     float64                                                    `json:"logprob,required"`
	TopLogprobs []ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                    `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
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
type ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
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
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [ChatCompletionGetResponseInputMessageAssistant].
	ToolCalls []ChatCompletionGetResponseInputMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionGetResponseInputMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
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
// will be valid: OfString OfChatCompletionGetResponseInputMessageUserContentArray
// OfChatCompletionGetResponseInputMessageSystemContentArray
// OfChatCompletionGetResponseInputMessageAssistantContentArray
// OfChatCompletionGetResponseInputMessageToolContentArray
// OfChatCompletionGetResponseInputMessageDeveloperContentArray]
type ChatCompletionGetResponseInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageUserContentArrayItemUnion] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageUserContentArray []ChatCompletionGetResponseInputMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageSystemContentArray []ChatCompletionGetResponseInputMessageSystemContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageAssistantContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageAssistantContentArray []ChatCompletionGetResponseInputMessageAssistantContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageToolContentArray []ChatCompletionGetResponseInputMessageToolContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageDeveloperContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageDeveloperContentArray []ChatCompletionGetResponseInputMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfChatCompletionGetResponseInputMessageUserContentArray      respjson.Field
		OfChatCompletionGetResponseInputMessageSystemContentArray    respjson.Field
		OfChatCompletionGetResponseInputMessageAssistantContentArray respjson.Field
		OfChatCompletionGetResponseInputMessageToolContentArray      respjson.Field
		OfChatCompletionGetResponseInputMessageDeveloperContentArray respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (r *ChatCompletionGetResponseInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseInputMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionGetResponseInputMessageUserContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseInputMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfChatCompletionGetResponseInputMessageUserContentArray]
type ChatCompletionGetResponseInputMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageUserContentArrayItemUnion] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageUserContentArray []ChatCompletionGetResponseInputMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                    struct {
		OfString                                                respjson.Field
		OfChatCompletionGetResponseInputMessageUserContentArray respjson.Field
		raw                                                     string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUserContentUnion) AsChatCompletionGetResponseInputMessageUserContentArray() (v []ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageUserContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseInputMessageUserContentArrayItemText],
// [ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL],
// [ChatCompletionGetResponseInputMessageUserContentArrayItemFile].
//
// Use the [ChatCompletionGetResponseInputMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentArrayItemFile].
	File ChatCompletionGetResponseInputMessageUserContentArrayItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessageUserContentArrayItem is implemented by
// each variant of [ChatCompletionGetResponseInputMessageUserContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionGetResponseInputMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseInputMessageUserContentArrayItem interface {
	implChatCompletionGetResponseInputMessageUserContentArrayItemUnion()
}

func (ChatCompletionGetResponseInputMessageUserContentArrayItemText) implChatCompletionGetResponseInputMessageUserContentArrayItemUnion() {
}
func (ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL) implChatCompletionGetResponseInputMessageUserContentArrayItemUnion() {
}
func (ChatCompletionGetResponseInputMessageUserContentArrayItemFile) implChatCompletionGetResponseInputMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL:
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentArrayItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) AsAny() anyChatCompletionGetResponseInputMessageUserContentArrayItem {
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

func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) AsText() (v ChatCompletionGetResponseInputMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) AsFile() (v ChatCompletionGetResponseInputMessageUserContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	// Must be "image_url" to identify this as image content
	Type constant.ImageURL `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail string `json:"detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageUserContentArrayItemFile struct {
	File ChatCompletionGetResponseInputMessageUserContentArrayItemFileFile `json:"file,required"`
	Type constant.File                                                     `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageUserContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageUserContentArrayItemFileFile struct {
	FileData string `json:"file_data"`
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
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
func (r ChatCompletionGetResponseInputMessageUserContentArrayItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionGetResponseInputMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionGetResponseInputMessageSystemContentUnion `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role constant.System `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseInputMessageSystemContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseInputMessageSystemContentArray]
type ChatCompletionGetResponseInputMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageSystemContentArray []ChatCompletionGetResponseInputMessageSystemContentArrayItem `json:",inline"`
	JSON                                                      struct {
		OfString                                                  respjson.Field
		OfChatCompletionGetResponseInputMessageSystemContentArray respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageSystemContentUnion) AsChatCompletionGetResponseInputMessageSystemContentArray() (v []ChatCompletionGetResponseInputMessageSystemContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageSystemContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseInputMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionGetResponseInputMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionGetResponseInputMessageAssistantToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Content     respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseInputMessageAssistantContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseInputMessageAssistantContentArray]
type ChatCompletionGetResponseInputMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageAssistantContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageAssistantContentArray []ChatCompletionGetResponseInputMessageAssistantContentArrayItem `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfChatCompletionGetResponseInputMessageAssistantContentArray respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageAssistantContentUnion) AsChatCompletionGetResponseInputMessageAssistantContentArray() (v []ChatCompletionGetResponseInputMessageAssistantContentArrayItem) {
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
type ChatCompletionGetResponseInputMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageAssistantContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionGetResponseInputMessageAssistantToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionGetResponseInputMessageAssistantToolCallFunction `json:"function"`
	// (Optional) Index of the tool call in the list
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionGetResponseInputMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments string `json:"arguments"`
	// (Optional) Name of the function to call
	Name string `json:"name"`
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
	// The response content from the tool
	Content ChatCompletionGetResponseInputMessageToolContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ToolCallID  respjson.Field
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
// [[]ChatCompletionGetResponseInputMessageToolContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfChatCompletionGetResponseInputMessageToolContentArray]
type ChatCompletionGetResponseInputMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageToolContentArray []ChatCompletionGetResponseInputMessageToolContentArrayItem `json:",inline"`
	JSON                                                    struct {
		OfString                                                respjson.Field
		OfChatCompletionGetResponseInputMessageToolContentArray respjson.Field
		raw                                                     string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageToolContentUnion) AsChatCompletionGetResponseInputMessageToolContentArray() (v []ChatCompletionGetResponseInputMessageToolContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionGetResponseInputMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageToolContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionGetResponseInputMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionGetResponseInputMessageDeveloperContentUnion `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role constant.Developer `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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
// [[]ChatCompletionGetResponseInputMessageDeveloperContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionGetResponseInputMessageDeveloperContentArray]
type ChatCompletionGetResponseInputMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageDeveloperContentArrayItem] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageDeveloperContentArray []ChatCompletionGetResponseInputMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfChatCompletionGetResponseInputMessageDeveloperContentArray respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (u ChatCompletionGetResponseInputMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageDeveloperContentUnion) AsChatCompletionGetResponseInputMessageDeveloperContentArray() (v []ChatCompletionGetResponseInputMessageDeveloperContentArrayItem) {
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
type ChatCompletionGetResponseInputMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageDeveloperContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponse struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []ChatCompletionListResponseChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created       int64                                         `json:"created,required"`
	InputMessages []ChatCompletionListResponseInputMessageUnion `json:"input_messages,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion"
	Object constant.ChatCompletion `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Choices       respjson.Field
		Created       respjson.Field
		InputMessages respjson.Field
		Model         respjson.Field
		Object        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type ChatCompletionListResponseChoice struct {
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// The message from the model
	Message ChatCompletionListResponseChoiceMessageUnion `json:"message,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChatCompletionListResponseChoiceLogprobs `json:"logprobs"`
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
func (r ChatCompletionListResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageUnion contains all possible properties
// and values from [ChatCompletionListResponseChoiceMessageUser],
// [ChatCompletionListResponseChoiceMessageSystem],
// [ChatCompletionListResponseChoiceMessageAssistant],
// [ChatCompletionListResponseChoiceMessageTool],
// [ChatCompletionListResponseChoiceMessageDeveloper].
//
// Use the [ChatCompletionListResponseChoiceMessageUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseChoiceMessageUnion struct {
	// This field is a union of
	// [ChatCompletionListResponseChoiceMessageUserContentUnion],
	// [ChatCompletionListResponseChoiceMessageSystemContentUnion],
	// [ChatCompletionListResponseChoiceMessageAssistantContentUnion],
	// [ChatCompletionListResponseChoiceMessageToolContentUnion],
	// [ChatCompletionListResponseChoiceMessageDeveloperContentUnion]
	Content ChatCompletionListResponseChoiceMessageUnionContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [ChatCompletionListResponseChoiceMessageAssistant].
	ToolCalls []ChatCompletionListResponseChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionListResponseChoiceMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionListResponseChoiceMessage is implemented by each variant of
// [ChatCompletionListResponseChoiceMessageUnion] to add type safety for the return
// type of [ChatCompletionListResponseChoiceMessageUnion.AsAny]
type anyChatCompletionListResponseChoiceMessage interface {
	implChatCompletionListResponseChoiceMessageUnion()
}

func (ChatCompletionListResponseChoiceMessageUser) implChatCompletionListResponseChoiceMessageUnion() {
}
func (ChatCompletionListResponseChoiceMessageSystem) implChatCompletionListResponseChoiceMessageUnion() {
}
func (ChatCompletionListResponseChoiceMessageAssistant) implChatCompletionListResponseChoiceMessageUnion() {
}
func (ChatCompletionListResponseChoiceMessageTool) implChatCompletionListResponseChoiceMessageUnion() {
}
func (ChatCompletionListResponseChoiceMessageDeveloper) implChatCompletionListResponseChoiceMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseChoiceMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseChoiceMessageUser:
//	case llamastackclient.ChatCompletionListResponseChoiceMessageSystem:
//	case llamastackclient.ChatCompletionListResponseChoiceMessageAssistant:
//	case llamastackclient.ChatCompletionListResponseChoiceMessageTool:
//	case llamastackclient.ChatCompletionListResponseChoiceMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseChoiceMessageUnion) AsAny() anyChatCompletionListResponseChoiceMessage {
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

func (u ChatCompletionListResponseChoiceMessageUnion) AsUser() (v ChatCompletionListResponseChoiceMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUnion) AsSystem() (v ChatCompletionListResponseChoiceMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUnion) AsAssistant() (v ChatCompletionListResponseChoiceMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUnion) AsTool() (v ChatCompletionListResponseChoiceMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUnion) AsDeveloper() (v ChatCompletionListResponseChoiceMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseChoiceMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageUnionContent is an implicit subunion of
// [ChatCompletionListResponseChoiceMessageUnion].
// ChatCompletionListResponseChoiceMessageUnionContent provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionListResponseChoiceMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseChoiceMessageUserContentArray
// OfChatCompletionListResponseChoiceMessageSystemContentArray
// OfChatCompletionListResponseChoiceMessageAssistantContentArray
// OfChatCompletionListResponseChoiceMessageToolContentArray
// OfChatCompletionListResponseChoiceMessageDeveloperContentArray]
type ChatCompletionListResponseChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionListResponseChoiceMessageUserContentArray []ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseChoiceMessageSystemContentArray []ChatCompletionListResponseChoiceMessageSystemContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageAssistantContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseChoiceMessageAssistantContentArray []ChatCompletionListResponseChoiceMessageAssistantContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseChoiceMessageToolContentArray []ChatCompletionListResponseChoiceMessageToolContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseChoiceMessageDeveloperContentArray []ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                           struct {
		OfString                                                       respjson.Field
		OfChatCompletionListResponseChoiceMessageUserContentArray      respjson.Field
		OfChatCompletionListResponseChoiceMessageSystemContentArray    respjson.Field
		OfChatCompletionListResponseChoiceMessageAssistantContentArray respjson.Field
		OfChatCompletionListResponseChoiceMessageToolContentArray      respjson.Field
		OfChatCompletionListResponseChoiceMessageDeveloperContentArray respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *ChatCompletionListResponseChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseChoiceMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionListResponseChoiceMessageUserContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseChoiceMessageUserContentArray]
type ChatCompletionListResponseChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionListResponseChoiceMessageUserContentArray []ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                      struct {
		OfString                                                  respjson.Field
		OfChatCompletionListResponseChoiceMessageUserContentArray respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUserContentUnion) AsChatCompletionListResponseChoiceMessageUserContentArray() (v []ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseChoiceMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionListResponseChoiceMessageUserContentArrayItemText],
// [ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL],
// [ChatCompletionListResponseChoiceMessageUserContentArrayItemFile].
//
// Use the [ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseChoiceMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionListResponseChoiceMessageUserContentArrayItemFile].
	File ChatCompletionListResponseChoiceMessageUserContentArrayItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseChoiceMessageUserContentArrayItem is implemented by
// each variant of
// [ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion] to add type
// safety for the return type of
// [ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseChoiceMessageUserContentArrayItem interface {
	implChatCompletionListResponseChoiceMessageUserContentArrayItemUnion()
}

func (ChatCompletionListResponseChoiceMessageUserContentArrayItemText) implChatCompletionListResponseChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL) implChatCompletionListResponseChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionListResponseChoiceMessageUserContentArrayItemFile) implChatCompletionListResponseChoiceMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseChoiceMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL:
//	case llamastackclient.ChatCompletionListResponseChoiceMessageUserContentArrayItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) AsAny() anyChatCompletionListResponseChoiceMessageUserContentArrayItem {
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

func (u ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) AsText() (v ChatCompletionListResponseChoiceMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) AsFile() (v ChatCompletionListResponseChoiceMessageUserContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseChoiceMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseChoiceMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	// Must be "image_url" to identify this as image content
	Type constant.ImageURL `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail string `json:"detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseChoiceMessageUserContentArrayItemFile struct {
	File ChatCompletionListResponseChoiceMessageUserContentArrayItemFileFile `json:"file,required"`
	Type constant.File                                                       `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageUserContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseChoiceMessageUserContentArrayItemFileFile struct {
	FileData string `json:"file_data"`
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
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
func (r ChatCompletionListResponseChoiceMessageUserContentArrayItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionListResponseChoiceMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionListResponseChoiceMessageSystemContentUnion `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role constant.System `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageSystemContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseChoiceMessageSystemContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseChoiceMessageSystemContentArray]
type ChatCompletionListResponseChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseChoiceMessageSystemContentArray []ChatCompletionListResponseChoiceMessageSystemContentArrayItem `json:",inline"`
	JSON                                                        struct {
		OfString                                                    respjson.Field
		OfChatCompletionListResponseChoiceMessageSystemContentArray respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (u ChatCompletionListResponseChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageSystemContentUnion) AsChatCompletionListResponseChoiceMessageSystemContentArray() (v []ChatCompletionListResponseChoiceMessageSystemContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageSystemContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseChoiceMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageSystemContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseChoiceMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionListResponseChoiceMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionListResponseChoiceMessageAssistantToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseChoiceMessageAssistantContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseChoiceMessageAssistantContentArray]
type ChatCompletionListResponseChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageAssistantContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseChoiceMessageAssistantContentArray []ChatCompletionListResponseChoiceMessageAssistantContentArrayItem `json:",inline"`
	JSON                                                           struct {
		OfString                                                       respjson.Field
		OfChatCompletionListResponseChoiceMessageAssistantContentArray respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (u ChatCompletionListResponseChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageAssistantContentUnion) AsChatCompletionListResponseChoiceMessageAssistantContentArray() (v []ChatCompletionListResponseChoiceMessageAssistantContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseChoiceMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseChoiceMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageAssistantContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionListResponseChoiceMessageAssistantToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionListResponseChoiceMessageAssistantToolCallFunction `json:"function"`
	// (Optional) Index of the tool call in the list
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionListResponseChoiceMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments string `json:"arguments"`
	// (Optional) Name of the function to call
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseChoiceMessageTool struct {
	// The response content from the tool
	Content ChatCompletionListResponseChoiceMessageToolContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseChoiceMessageToolContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseChoiceMessageToolContentArray]
type ChatCompletionListResponseChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseChoiceMessageToolContentArray []ChatCompletionListResponseChoiceMessageToolContentArrayItem `json:",inline"`
	JSON                                                      struct {
		OfString                                                  respjson.Field
		OfChatCompletionListResponseChoiceMessageToolContentArray respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (u ChatCompletionListResponseChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageToolContentUnion) AsChatCompletionListResponseChoiceMessageToolContentArray() (v []ChatCompletionListResponseChoiceMessageToolContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseChoiceMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageToolContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseChoiceMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionListResponseChoiceMessageDeveloperContentUnion `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role constant.Developer `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseChoiceMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseChoiceMessageDeveloperContentArray]
type ChatCompletionListResponseChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseChoiceMessageDeveloperContentArray []ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                           struct {
		OfString                                                       respjson.Field
		OfChatCompletionListResponseChoiceMessageDeveloperContentArray respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (u ChatCompletionListResponseChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseChoiceMessageDeveloperContentUnion) AsChatCompletionListResponseChoiceMessageDeveloperContentArray() (v []ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseChoiceMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseChoiceMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The log probabilities for the tokens in the message
type ChatCompletionListResponseChoiceLogprobs struct {
	// (Optional) The log probabilities for the tokens in the message
	Content []ChatCompletionListResponseChoiceLogprobsContent `json:"content"`
	// (Optional) The log probabilities for the tokens in the message
	Refusal []ChatCompletionListResponseChoiceLogprobsRefusal `json:"refusal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseChoiceLogprobsContent struct {
	Token       string                                                      `json:"token,required"`
	Logprob     float64                                                     `json:"logprob,required"`
	TopLogprobs []ChatCompletionListResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                     `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseChoiceLogprobsContentTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
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
func (r ChatCompletionListResponseChoiceLogprobsContentTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseChoiceLogprobsRefusal struct {
	Token       string                                                      `json:"token,required"`
	Logprob     float64                                                     `json:"logprob,required"`
	TopLogprobs []ChatCompletionListResponseChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                     `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseChoiceLogprobsRefusalTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
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
func (r ChatCompletionListResponseChoiceLogprobsRefusalTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageUnion contains all possible properties and
// values from [ChatCompletionListResponseInputMessageUser],
// [ChatCompletionListResponseInputMessageSystem],
// [ChatCompletionListResponseInputMessageAssistant],
// [ChatCompletionListResponseInputMessageTool],
// [ChatCompletionListResponseInputMessageDeveloper].
//
// Use the [ChatCompletionListResponseInputMessageUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseInputMessageUnion struct {
	// This field is a union of
	// [ChatCompletionListResponseInputMessageUserContentUnion],
	// [ChatCompletionListResponseInputMessageSystemContentUnion],
	// [ChatCompletionListResponseInputMessageAssistantContentUnion],
	// [ChatCompletionListResponseInputMessageToolContentUnion],
	// [ChatCompletionListResponseInputMessageDeveloperContentUnion]
	Content ChatCompletionListResponseInputMessageUnionContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [ChatCompletionListResponseInputMessageAssistant].
	ToolCalls []ChatCompletionListResponseInputMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionListResponseInputMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChatCompletionListResponseInputMessage is implemented by each variant of
// [ChatCompletionListResponseInputMessageUnion] to add type safety for the return
// type of [ChatCompletionListResponseInputMessageUnion.AsAny]
type anyChatCompletionListResponseInputMessage interface {
	implChatCompletionListResponseInputMessageUnion()
}

func (ChatCompletionListResponseInputMessageUser) implChatCompletionListResponseInputMessageUnion() {}
func (ChatCompletionListResponseInputMessageSystem) implChatCompletionListResponseInputMessageUnion() {
}
func (ChatCompletionListResponseInputMessageAssistant) implChatCompletionListResponseInputMessageUnion() {
}
func (ChatCompletionListResponseInputMessageTool) implChatCompletionListResponseInputMessageUnion() {}
func (ChatCompletionListResponseInputMessageDeveloper) implChatCompletionListResponseInputMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseInputMessageUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseInputMessageUser:
//	case llamastackclient.ChatCompletionListResponseInputMessageSystem:
//	case llamastackclient.ChatCompletionListResponseInputMessageAssistant:
//	case llamastackclient.ChatCompletionListResponseInputMessageTool:
//	case llamastackclient.ChatCompletionListResponseInputMessageDeveloper:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseInputMessageUnion) AsAny() anyChatCompletionListResponseInputMessage {
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

func (u ChatCompletionListResponseInputMessageUnion) AsUser() (v ChatCompletionListResponseInputMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUnion) AsSystem() (v ChatCompletionListResponseInputMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUnion) AsAssistant() (v ChatCompletionListResponseInputMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUnion) AsTool() (v ChatCompletionListResponseInputMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUnion) AsDeveloper() (v ChatCompletionListResponseInputMessageDeveloper) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageUnionContent is an implicit subunion of
// [ChatCompletionListResponseInputMessageUnion].
// ChatCompletionListResponseInputMessageUnionContent provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatCompletionListResponseInputMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfChatCompletionListResponseInputMessageUserContentArray
// OfChatCompletionListResponseInputMessageSystemContentArray
// OfChatCompletionListResponseInputMessageAssistantContentArray
// OfChatCompletionListResponseInputMessageToolContentArray
// OfChatCompletionListResponseInputMessageDeveloperContentArray]
type ChatCompletionListResponseInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageUserContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionListResponseInputMessageUserContentArray []ChatCompletionListResponseInputMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseInputMessageSystemContentArray []ChatCompletionListResponseInputMessageSystemContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageAssistantContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseInputMessageAssistantContentArray []ChatCompletionListResponseInputMessageAssistantContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseInputMessageToolContentArray []ChatCompletionListResponseInputMessageToolContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageDeveloperContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseInputMessageDeveloperContentArray []ChatCompletionListResponseInputMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionListResponseInputMessageUserContentArray      respjson.Field
		OfChatCompletionListResponseInputMessageSystemContentArray    respjson.Field
		OfChatCompletionListResponseInputMessageAssistantContentArray respjson.Field
		OfChatCompletionListResponseInputMessageToolContentArray      respjson.Field
		OfChatCompletionListResponseInputMessageDeveloperContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (r *ChatCompletionListResponseInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseInputMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionListResponseInputMessageUserContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseInputMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseInputMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseInputMessageUserContentArray]
type ChatCompletionListResponseInputMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageUserContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionListResponseInputMessageUserContentArray []ChatCompletionListResponseInputMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                     struct {
		OfString                                                 respjson.Field
		OfChatCompletionListResponseInputMessageUserContentArray respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (u ChatCompletionListResponseInputMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUserContentUnion) AsChatCompletionListResponseInputMessageUserContentArray() (v []ChatCompletionListResponseInputMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseInputMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageUserContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionListResponseInputMessageUserContentArrayItemText],
// [ChatCompletionListResponseInputMessageUserContentArrayItemImageURL],
// [ChatCompletionListResponseInputMessageUserContentArrayItemFile].
//
// Use the [ChatCompletionListResponseInputMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseInputMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseInputMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseInputMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseInputMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	// This field is from variant
	// [ChatCompletionListResponseInputMessageUserContentArrayItemFile].
	File ChatCompletionListResponseInputMessageUserContentArrayItemFileFile `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseInputMessageUserContentArrayItem is implemented by
// each variant of
// [ChatCompletionListResponseInputMessageUserContentArrayItemUnion] to add type
// safety for the return type of
// [ChatCompletionListResponseInputMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseInputMessageUserContentArrayItem interface {
	implChatCompletionListResponseInputMessageUserContentArrayItemUnion()
}

func (ChatCompletionListResponseInputMessageUserContentArrayItemText) implChatCompletionListResponseInputMessageUserContentArrayItemUnion() {
}
func (ChatCompletionListResponseInputMessageUserContentArrayItemImageURL) implChatCompletionListResponseInputMessageUserContentArrayItemUnion() {
}
func (ChatCompletionListResponseInputMessageUserContentArrayItemFile) implChatCompletionListResponseInputMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseInputMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseInputMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseInputMessageUserContentArrayItemImageURL:
//	case llamastackclient.ChatCompletionListResponseInputMessageUserContentArrayItemFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseInputMessageUserContentArrayItemUnion) AsAny() anyChatCompletionListResponseInputMessageUserContentArrayItem {
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

func (u ChatCompletionListResponseInputMessageUserContentArrayItemUnion) AsText() (v ChatCompletionListResponseInputMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseInputMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageUserContentArrayItemUnion) AsFile() (v ChatCompletionListResponseInputMessageUserContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseInputMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseInputMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseInputMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL ChatCompletionListResponseInputMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	// Must be "image_url" to identify this as image content
	Type constant.ImageURL `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type ChatCompletionListResponseInputMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail string `json:"detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseInputMessageUserContentArrayItemFile struct {
	File ChatCompletionListResponseInputMessageUserContentArrayItemFileFile `json:"file,required"`
	Type constant.File                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageUserContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseInputMessageUserContentArrayItemFileFile struct {
	FileData string `json:"file_data"`
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
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
func (r ChatCompletionListResponseInputMessageUserContentArrayItemFileFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionListResponseInputMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionListResponseInputMessageSystemContentUnion `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role constant.System `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseInputMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageSystemContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseInputMessageSystemContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseInputMessageSystemContentArray]
type ChatCompletionListResponseInputMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageSystemContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseInputMessageSystemContentArray []ChatCompletionListResponseInputMessageSystemContentArrayItem `json:",inline"`
	JSON                                                       struct {
		OfString                                                   respjson.Field
		OfChatCompletionListResponseInputMessageSystemContentArray respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (u ChatCompletionListResponseInputMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageSystemContentUnion) AsChatCompletionListResponseInputMessageSystemContentArray() (v []ChatCompletionListResponseInputMessageSystemContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseInputMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseInputMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageSystemContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseInputMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionListResponseInputMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionListResponseInputMessageAssistantToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseInputMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseInputMessageAssistantContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseInputMessageAssistantContentArray]
type ChatCompletionListResponseInputMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageAssistantContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseInputMessageAssistantContentArray []ChatCompletionListResponseInputMessageAssistantContentArrayItem `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionListResponseInputMessageAssistantContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u ChatCompletionListResponseInputMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageAssistantContentUnion) AsChatCompletionListResponseInputMessageAssistantContentArray() (v []ChatCompletionListResponseInputMessageAssistantContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseInputMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseInputMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageAssistantContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionListResponseInputMessageAssistantToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionListResponseInputMessageAssistantToolCallFunction `json:"function"`
	// (Optional) Index of the tool call in the list
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionListResponseInputMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments string `json:"arguments"`
	// (Optional) Name of the function to call
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseInputMessageTool struct {
	// The response content from the tool
	Content ChatCompletionListResponseInputMessageToolContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseInputMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseInputMessageToolContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseInputMessageToolContentArray]
type ChatCompletionListResponseInputMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageToolContentArrayItem] instead of an
	// object.
	OfChatCompletionListResponseInputMessageToolContentArray []ChatCompletionListResponseInputMessageToolContentArrayItem `json:",inline"`
	JSON                                                     struct {
		OfString                                                 respjson.Field
		OfChatCompletionListResponseInputMessageToolContentArray respjson.Field
		raw                                                      string
	} `json:"-"`
}

func (u ChatCompletionListResponseInputMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageToolContentUnion) AsChatCompletionListResponseInputMessageToolContentArray() (v []ChatCompletionListResponseInputMessageToolContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionListResponseInputMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseInputMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageToolContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseInputMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionListResponseInputMessageDeveloperContentUnion `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role constant.Developer `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseInputMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseInputMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseInputMessageDeveloperContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseInputMessageDeveloperContentArray]
type ChatCompletionListResponseInputMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseInputMessageDeveloperContentArrayItem] instead of
	// an object.
	OfChatCompletionListResponseInputMessageDeveloperContentArray []ChatCompletionListResponseInputMessageDeveloperContentArrayItem `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionListResponseInputMessageDeveloperContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u ChatCompletionListResponseInputMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseInputMessageDeveloperContentUnion) AsChatCompletionListResponseInputMessageDeveloperContentArray() (v []ChatCompletionListResponseInputMessageDeveloperContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseInputMessageDeveloperContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseInputMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ChatCompletionListResponseInputMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseInputMessageDeveloperContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseInputMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewParams struct {
	// List of messages in the conversation.
	Messages []ChatCompletionNewParamsMessageUnion `json:"messages,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	Model string `json:"model,required"`
	// (Optional) The penalty for repeated tokens.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// (Optional) The log probabilities to use.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// (Optional) The maximum number of tokens to generate.
	MaxCompletionTokens param.Opt[int64] `json:"max_completion_tokens,omitzero"`
	// (Optional) The maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// (Optional) The number of completions to generate.
	N param.Opt[int64] `json:"n,omitzero"`
	// (Optional) Whether to parallelize tool calls.
	ParallelToolCalls param.Opt[bool] `json:"parallel_tool_calls,omitzero"`
	// (Optional) The penalty for repeated tokens.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// (Optional) The seed to use.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// (Optional) The temperature to use.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// (Optional) The top log probabilities to use.
	TopLogprobs param.Opt[int64] `json:"top_logprobs,omitzero"`
	// (Optional) The top p to use.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// (Optional) The user to use.
	User param.Opt[string] `json:"user,omitzero"`
	// (Optional) The function call to use.
	FunctionCall ChatCompletionNewParamsFunctionCallUnion `json:"function_call,omitzero"`
	// (Optional) List of functions to use.
	Functions []map[string]ChatCompletionNewParamsFunctionUnion `json:"functions,omitzero"`
	// (Optional) The logit bias to use.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// (Optional) The response format to use.
	ResponseFormat ChatCompletionNewParamsResponseFormatUnion `json:"response_format,omitzero"`
	// (Optional) The stop tokens to use.
	Stop ChatCompletionNewParamsStopUnion `json:"stop,omitzero"`
	// (Optional) The stream options to use.
	StreamOptions map[string]ChatCompletionNewParamsStreamOptionUnion `json:"stream_options,omitzero"`
	// (Optional) The tool choice to use.
	ToolChoice ChatCompletionNewParamsToolChoiceUnion `json:"tool_choice,omitzero"`
	// (Optional) The tools to use.
	Tools []map[string]ChatCompletionNewParamsToolUnion `json:"tools,omitzero"`
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
// [_[]ChatCompletionNewParamsMessageUserContentArrayItemUnion],
// [_[]ChatCompletionNewParamsMessageSystemContentArrayItem],
// [_[]ChatCompletionNewParamsMessageAssistantContentArrayItem],
// [_[]ChatCompletionNewParamsMessageToolContentArrayItem],
// [\*[]ChatCompletionNewParamsMessageDeveloperContentArrayItem]
type chatCompletionNewParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageUserContentArrayItemUnion:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageSystemContentArrayItem:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageAssistantContentArrayItem:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageToolContentArrayItem:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageDeveloperContentArrayItem:
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
// The properties Content, Role are required.
type ChatCompletionNewParamsMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionNewParamsMessageUserContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUser) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageUserContentUnion struct {
	OfString                                    param.Opt[string]                                         `json:",omitzero,inline"`
	OfChatCompletionNewsMessageUserContentArray []ChatCompletionNewParamsMessageUserContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsMessageUserContentArray)
}
func (u *ChatCompletionNewParamsMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUserContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageUserContentArray) {
		return &u.OfChatCompletionNewsMessageUserContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageUserContentArrayItemUnion struct {
	OfText     *ChatCompletionNewParamsMessageUserContentArrayItemText     `json:",omitzero,inline"`
	OfImageURL *ChatCompletionNewParamsMessageUserContentArrayItemImageURL `json:",omitzero,inline"`
	OfFile     *ChatCompletionNewParamsMessageUserContentArrayItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *ChatCompletionNewParamsMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUserContentArrayItemUnion) asAny() any {
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
func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) GetImageURL() *ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) GetFile() *ChatCompletionNewParamsMessageUserContentArrayItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) GetType() *string {
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
	apijson.RegisterUnion[ChatCompletionNewParamsMessageUserContentArrayItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentArrayItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentArrayItemImageURL]("image_url"),
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentArrayItemFile]("file"),
	)
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type ChatCompletionNewParamsMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentArrayItemText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentArrayItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type ChatCompletionNewParamsMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties File, Type are required.
type ChatCompletionNewParamsMessageUserContentArrayItemFile struct {
	File ChatCompletionNewParamsMessageUserContentArrayItemFileFile `json:"file,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentArrayItemFile) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentArrayItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewParamsMessageUserContentArrayItemFileFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentArrayItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentArrayItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type ChatCompletionNewParamsMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionNewParamsMessageSystemContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageSystemContentUnion struct {
	OfString                                      param.Opt[string]                                      `json:",omitzero,inline"`
	OfChatCompletionNewsMessageSystemContentArray []ChatCompletionNewParamsMessageSystemContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageSystemContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsMessageSystemContentArray)
}
func (u *ChatCompletionNewParamsMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageSystemContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageSystemContentArray) {
		return &u.OfChatCompletionNewsMessageSystemContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type ChatCompletionNewParamsMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystemContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystemContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
//
// The property Role is required.
type ChatCompletionNewParamsMessageAssistant struct {
	// (Optional) The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response
	Content ChatCompletionNewParamsMessageAssistantContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionNewParamsMessageAssistantToolCall `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistant) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageAssistantContentUnion struct {
	OfString                                         param.Opt[string]                                         `json:",omitzero,inline"`
	OfChatCompletionNewsMessageAssistantContentArray []ChatCompletionNewParamsMessageAssistantContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageAssistantContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsMessageAssistantContentArray)
}
func (u *ChatCompletionNewParamsMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageAssistantContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageAssistantContentArray) {
		return &u.OfChatCompletionNewsMessageAssistantContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type ChatCompletionNewParamsMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
//
// The property Type is required.
type ChatCompletionNewParamsMessageAssistantToolCall struct {
	// (Optional) Unique identifier for the tool call
	ID param.Opt[string] `json:"id,omitzero"`
	// (Optional) Index of the tool call in the list
	Index param.Opt[int64] `json:"index,omitzero"`
	// (Optional) Function call details
	Function ChatCompletionNewParamsMessageAssistantToolCallFunction `json:"function,omitzero"`
	// Must be "function" to identify this as a function call
	//
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionNewParamsMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	// (Optional) Name of the function to call
	Name param.Opt[string] `json:"name,omitzero"`
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
// The properties Content, Role, ToolCallID are required.
type ChatCompletionNewParamsMessageTool struct {
	// The response content from the tool
	Content ChatCompletionNewParamsMessageToolContentUnion `json:"content,omitzero,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageTool) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageToolContentUnion struct {
	OfString                                    param.Opt[string]                                    `json:",omitzero,inline"`
	OfChatCompletionNewsMessageToolContentArray []ChatCompletionNewParamsMessageToolContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageToolContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsMessageToolContentArray)
}
func (u *ChatCompletionNewParamsMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageToolContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageToolContentArray) {
		return &u.OfChatCompletionNewsMessageToolContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type ChatCompletionNewParamsMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageToolContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageToolContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type ChatCompletionNewParamsMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionNewParamsMessageDeveloperContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "developer" to identify this as a developer message
	//
	// This field can be elided, and will marshal its zero value as "developer".
	Role constant.Developer `json:"role,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageDeveloperContentUnion struct {
	OfString                                         param.Opt[string]                                         `json:",omitzero,inline"`
	OfChatCompletionNewsMessageDeveloperContentArray []ChatCompletionNewParamsMessageDeveloperContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageDeveloperContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsMessageDeveloperContentArray)
}
func (u *ChatCompletionNewParamsMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageDeveloperContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageDeveloperContentArray) {
		return &u.OfChatCompletionNewsMessageDeveloperContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type ChatCompletionNewParamsMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloperContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloperContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsFunctionCallUnion struct {
	OfString                               param.Opt[string]                                          `json:",omitzero,inline"`
	OfChatCompletionNewsFunctionCallMapMap map[string]ChatCompletionNewParamsFunctionCallMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsFunctionCallUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsFunctionCallMapMap)
}
func (u *ChatCompletionNewParamsFunctionCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsFunctionCallUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsFunctionCallMapMap) {
		return &u.OfChatCompletionNewsFunctionCallMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsFunctionCallMapItemUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsFunctionCallMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChatCompletionNewParamsFunctionCallMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsFunctionCallMapItemUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsFunctionUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsFunctionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChatCompletionNewParamsFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsFunctionUnion) asAny() any {
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

func NewChatCompletionNewParamsResponseFormatText() ChatCompletionNewParamsResponseFormatText {
	return ChatCompletionNewParamsResponseFormatText{
		Type: "text",
	}
}

// Text response format for OpenAI-compatible chat completion requests.
//
// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsResponseFormatText].
type ChatCompletionNewParamsResponseFormatText struct {
	// Must be "text" to indicate plain text response format
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON schema response format for OpenAI-compatible chat completion requests.
//
// The properties JsonSchema, Type are required.
type ChatCompletionNewParamsResponseFormatJsonSchema struct {
	// The JSON schema specification for the response
	JsonSchema ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero,required"`
	// Must be "json_schema" to indicate structured JSON response format
	//
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JsonSchema `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The JSON schema specification for the response
//
// The property Name is required.
type ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema struct {
	// Name of the schema
	Name string `json:"name,required"`
	// (Optional) Description of the schema
	Description param.Opt[string] `json:"description,omitzero"`
	// (Optional) Whether to enforce strict adherence to the schema
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// (Optional) The JSON schema definition
	Schema map[string]ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion `json:"schema,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion) asAny() any {
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

func NewChatCompletionNewParamsResponseFormatJsonObject() ChatCompletionNewParamsResponseFormatJsonObject {
	return ChatCompletionNewParamsResponseFormatJsonObject{
		Type: "json_object",
	}
}

// JSON object response format for OpenAI-compatible chat completion requests.
//
// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsResponseFormatJsonObject].
type ChatCompletionNewParamsResponseFormatJsonObject struct {
	// Must be "json_object" to indicate generic JSON object response format
	Type constant.JsonObject `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonObject) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsStopUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ChatCompletionNewParamsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsStopUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsStreamOptionUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsStreamOptionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChatCompletionNewParamsStreamOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsStreamOptionUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsToolChoiceUnion struct {
	OfString                             param.Opt[string]                                        `json:",omitzero,inline"`
	OfChatCompletionNewsToolChoiceMapMap map[string]ChatCompletionNewParamsToolChoiceMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsToolChoiceMapMap)
}
func (u *ChatCompletionNewParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsToolChoiceMapMap) {
		return &u.OfChatCompletionNewsToolChoiceMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsToolChoiceMapItemUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsToolChoiceMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChatCompletionNewParamsToolChoiceMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsToolChoiceMapItemUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsToolUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChatCompletionNewParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsToolUnion) asAny() any {
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

type ChatCompletionListParams struct {
	// The ID of the last chat completion to return.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of chat completions to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The model to filter by.
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	// The order to sort the chat completions by: "asc" or "desc". Defaults to "desc".
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

// The order to sort the chat completions by: "asc" or "desc". Defaults to "desc".
type ChatCompletionListParamsOrder string

const (
	ChatCompletionListParamsOrderAsc  ChatCompletionListParamsOrder = "asc"
	ChatCompletionListParamsOrderDesc ChatCompletionListParamsOrder = "desc"
)
