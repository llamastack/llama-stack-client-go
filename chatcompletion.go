// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/llama-stack-client-go/internal/apijson"
	"github.com/stainless-sdks/llama-stack-client-go/internal/apiquery"
	"github.com/stainless-sdks/llama-stack-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llama-stack-client-go/option"
	"github.com/stainless-sdks/llama-stack-client-go/packages/param"
	"github.com/stainless-sdks/llama-stack-client-go/packages/respjson"
	"github.com/stainless-sdks/llama-stack-client-go/packages/ssestream"
	"github.com/stainless-sdks/llama-stack-client-go/shared/constant"
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
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/chat/completions"
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/openai/v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

// Describe a chat completion by its ID.
func (r *ChatCompletionService) Get(ctx context.Context, completionID string, opts ...option.RequestOption) (res *ChatCompletionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if completionID == "" {
		err = errors.New("missing required completion_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/chat/completions/%s", completionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all chat completions.
func (r *ChatCompletionService) List(ctx context.Context, query ChatCompletionListParams, opts ...option.RequestOption) (res *ChatCompletionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion `json:",inline"`
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
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL].
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
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
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

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemText struct {
	Text string        `json:"text,required"`
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

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                              `json:"type,required"`
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

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageUserContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion].
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
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion) {
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

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL].
//
// Use the
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL].
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem
// is implemented by each variant of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion.AsAny]
type anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem interface {
	implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion()
}

func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion) AsText() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion) AsImageURL() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                                `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageSystemContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion].
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
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion) {
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

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL].
//
// Use the
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL].
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem
// is implemented by each variant of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion.AsAny]
type anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem interface {
	implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion()
}

func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion) AsText() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion) AsImageURL() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCall struct {
	Type     constant.Function                                                                   `json:"type,required"`
	ID       string                                                                              `json:"id"`
	Function ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCallFunction `json:"function"`
	Index    int64                                                                               `json:"index"`
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

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
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
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion].
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
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion) {
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

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL].
//
// Use the
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL].
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem
// is implemented by each variant of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion.AsAny]
type anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem interface {
	implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion()
}

func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion) AsText() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion) AsImageURL() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageToolContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion].
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
	// [[]ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentUnion) AsChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArray() (v []ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion) {
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

// ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText],
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL].
//
// Use the
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL].
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem
// is implemented by each variant of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion.AsAny]
type anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem interface {
	implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion()
}

func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion() {
}
func (ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL) implChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText:
//	case llamastackclient.ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion) AsAny() anyChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion) AsText() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion) AsImageURL() (v ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionNewResponseOpenAIChatCompletionChoiceMessageDeveloperContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
	// [[]ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageSystemContentArray []ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseChoiceMessageAssistantContentArray []ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageToolContentArray []ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseChoiceMessageDeveloperContentArray []ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion `json:",inline"`
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
// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL].
//
// Use the [ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageUserContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
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

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageUserContentArrayItemText struct {
	Text string        `json:"text,required"`
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

type ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                          `json:"type,required"`
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

type ChatCompletionGetResponseChoiceMessageUserContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
// [[]ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageSystemContentArray []ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) AsChatCompletionGetResponseChoiceMessageSystemContentArray() (v []ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText],
// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL].
//
// Use the
// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessageSystemContentArrayItem is implemented
// by each variant of
// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion] to add type
// safety for the return type of
// [ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessageSystemContentArrayItem interface {
	implChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion()
}

func (ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText) implChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL) implChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageSystemContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion) AsText() (v ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageSystemContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageSystemContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                            `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageSystemContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [[]ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseChoiceMessageAssistantContentArray []ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseChoiceMessageAssistantContentUnion) AsChatCompletionGetResponseChoiceMessageAssistantContentArray() (v []ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion) {
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

// ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion contains
// all possible properties and values from
// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText],
// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL].
//
// Use the
// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessageAssistantContentArrayItem is
// implemented by each variant of
// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessageAssistantContentArrayItem interface {
	implChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion()
}

func (ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText) implChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL) implChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageAssistantContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion) AsText() (v ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                               `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageAssistantContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageAssistantToolCall struct {
	Type     constant.Function                                               `json:"type,required"`
	ID       string                                                          `json:"id"`
	Function ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction `json:"function"`
	Index    int64                                                           `json:"index"`
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

type ChatCompletionGetResponseChoiceMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
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
// [[]ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseChoiceMessageToolContentArray []ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) AsChatCompletionGetResponseChoiceMessageToolContentArray() (v []ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseChoiceMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseChoiceMessageToolContentArrayItemText],
// [ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL].
//
// Use the [ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageToolContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessageToolContentArrayItem is implemented by
// each variant of
// [ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion] to add type
// safety for the return type of
// [ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessageToolContentArrayItem interface {
	implChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion()
}

func (ChatCompletionGetResponseChoiceMessageToolContentArrayItemText) implChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL) implChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageToolContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageToolContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion) AsText() (v ChatCompletionGetResponseChoiceMessageToolContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageToolContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageToolContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseChoiceMessageToolContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageToolContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                          `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageToolContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseChoiceMessageDeveloperContentArray []ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentUnion) AsChatCompletionGetResponseChoiceMessageDeveloperContentArray() (v []ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion) {
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

// ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion contains
// all possible properties and values from
// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText],
// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL].
//
// Use the
// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem is
// implemented by each variant of
// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem interface {
	implChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion()
}

func (ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText) implChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion() {
}
func (ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL) implChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion) AsAny() anyChatCompletionGetResponseChoiceMessageDeveloperContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion) AsText() (v ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                               `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseChoiceMessageDeveloperContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
	// [[]ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseInputMessageSystemContentArray []ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseInputMessageAssistantContentArray []ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageToolContentArrayItemUnion] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageToolContentArray []ChatCompletionGetResponseInputMessageToolContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseInputMessageDeveloperContentArray []ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion `json:",inline"`
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
// [ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL].
//
// Use the [ChatCompletionGetResponseInputMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) AsAny() anyChatCompletionGetResponseInputMessageUserContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
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

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageUserContentArrayItemText struct {
	Text string        `json:"text,required"`
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

type ChatCompletionGetResponseInputMessageUserContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                         `json:"type,required"`
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

type ChatCompletionGetResponseInputMessageUserContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
// [[]ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion] instead of
	// an object.
	OfChatCompletionGetResponseInputMessageSystemContentArray []ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseInputMessageSystemContentUnion) AsChatCompletionGetResponseInputMessageSystemContentArray() (v []ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageSystemContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseInputMessageSystemContentArrayItemText],
// [ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL].
//
// Use the [ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageSystemContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessageSystemContentArrayItem is implemented by
// each variant of
// [ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion] to add type
// safety for the return type of
// [ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseInputMessageSystemContentArrayItem interface {
	implChatCompletionGetResponseInputMessageSystemContentArrayItemUnion()
}

func (ChatCompletionGetResponseInputMessageSystemContentArrayItemText) implChatCompletionGetResponseInputMessageSystemContentArrayItemUnion() {
}
func (ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL) implChatCompletionGetResponseInputMessageSystemContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageSystemContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion) AsAny() anyChatCompletionGetResponseInputMessageSystemContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion) AsText() (v ChatCompletionGetResponseInputMessageSystemContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageSystemContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageSystemContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseInputMessageSystemContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageSystemContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                           `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageSystemContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [[]ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseInputMessageAssistantContentArray []ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseInputMessageAssistantContentUnion) AsChatCompletionGetResponseInputMessageAssistantContentArray() (v []ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion) {
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

// ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemText],
// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL].
//
// Use the
// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessageAssistantContentArrayItem is implemented
// by each variant of
// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseInputMessageAssistantContentArrayItem interface {
	implChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion()
}

func (ChatCompletionGetResponseInputMessageAssistantContentArrayItemText) implChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion() {
}
func (ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL) implChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageAssistantContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion) AsAny() anyChatCompletionGetResponseInputMessageAssistantContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion) AsText() (v ChatCompletionGetResponseInputMessageAssistantContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageAssistantContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseInputMessageAssistantContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageAssistantToolCall struct {
	Type     constant.Function                                              `json:"type,required"`
	ID       string                                                         `json:"id"`
	Function ChatCompletionGetResponseInputMessageAssistantToolCallFunction `json:"function"`
	Index    int64                                                          `json:"index"`
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

type ChatCompletionGetResponseInputMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
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
// [[]ChatCompletionGetResponseInputMessageToolContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfChatCompletionGetResponseInputMessageToolContentArray]
type ChatCompletionGetResponseInputMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionGetResponseInputMessageToolContentArrayItemUnion] instead of an
	// object.
	OfChatCompletionGetResponseInputMessageToolContentArray []ChatCompletionGetResponseInputMessageToolContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseInputMessageToolContentUnion) AsChatCompletionGetResponseInputMessageToolContentArray() (v []ChatCompletionGetResponseInputMessageToolContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatCompletionGetResponseInputMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionGetResponseInputMessageToolContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseInputMessageToolContentArrayItemText],
// [ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL].
//
// Use the [ChatCompletionGetResponseInputMessageToolContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageToolContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageToolContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageToolContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessageToolContentArrayItem is implemented by
// each variant of [ChatCompletionGetResponseInputMessageToolContentArrayItemUnion]
// to add type safety for the return type of
// [ChatCompletionGetResponseInputMessageToolContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseInputMessageToolContentArrayItem interface {
	implChatCompletionGetResponseInputMessageToolContentArrayItemUnion()
}

func (ChatCompletionGetResponseInputMessageToolContentArrayItemText) implChatCompletionGetResponseInputMessageToolContentArrayItemUnion() {
}
func (ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL) implChatCompletionGetResponseInputMessageToolContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageToolContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageToolContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageToolContentArrayItemUnion) AsAny() anyChatCompletionGetResponseInputMessageToolContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseInputMessageToolContentArrayItemUnion) AsText() (v ChatCompletionGetResponseInputMessageToolContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageToolContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageToolContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageToolContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageToolContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseInputMessageToolContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageToolContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseInputMessageToolContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                         `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageToolContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageToolContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseInputMessageToolContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageToolContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [[]ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion].
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
	// [[]ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionGetResponseInputMessageDeveloperContentArray []ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion `json:",inline"`
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

func (u ChatCompletionGetResponseInputMessageDeveloperContentUnion) AsChatCompletionGetResponseInputMessageDeveloperContentArray() (v []ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion) {
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

// ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText],
// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL].
//
// Use the
// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL].
	ImageURL ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionGetResponseInputMessageDeveloperContentArrayItem is implemented
// by each variant of
// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion.AsAny]
type anyChatCompletionGetResponseInputMessageDeveloperContentArrayItem interface {
	implChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion()
}

func (ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText) implChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion() {
}
func (ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL) implChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText:
//	case llamastackclient.ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion) AsAny() anyChatCompletionGetResponseInputMessageDeveloperContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion) AsText() (v ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion) AsImageURL() (v ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionGetResponseInputMessageDeveloperContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageDeveloperContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL struct {
	ImageURL ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageDeveloperContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponse struct {
	Data    []ChatCompletionListResponseData `json:"data,required"`
	FirstID string                           `json:"first_id,required"`
	HasMore bool                             `json:"has_more,required"`
	LastID  string                           `json:"last_id,required"`
	Object  constant.List                    `json:"object,required"`
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
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []ChatCompletionListResponseDataChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created       int64                                             `json:"created,required"`
	InputMessages []ChatCompletionListResponseDataInputMessageUnion `json:"input_messages,required"`
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
func (r ChatCompletionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type ChatCompletionListResponseDataChoice struct {
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// The message from the model
	Message ChatCompletionListResponseDataChoiceMessageUnion `json:"message,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChatCompletionListResponseDataChoiceLogprobs `json:"logprobs"`
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
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageAssistant].
	ToolCalls []ChatCompletionListResponseDataChoiceMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionListResponseDataChoiceMessageTool].
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
// OfChatCompletionListResponseDataChoiceMessageUserContentArray
// OfChatCompletionListResponseDataChoiceMessageSystemContentArray
// OfChatCompletionListResponseDataChoiceMessageAssistantContentArray
// OfChatCompletionListResponseDataChoiceMessageToolContentArray
// OfChatCompletionListResponseDataChoiceMessageDeveloperContentArray]
type ChatCompletionListResponseDataChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataChoiceMessageUserContentArray []ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataChoiceMessageSystemContentArray []ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataChoiceMessageAssistantContentArray []ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataChoiceMessageToolContentArray []ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataChoiceMessageDeveloperContentArray []ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfChatCompletionListResponseDataChoiceMessageUserContentArray      respjson.Field
		OfChatCompletionListResponseDataChoiceMessageSystemContentArray    respjson.Field
		OfChatCompletionListResponseDataChoiceMessageAssistantContentArray respjson.Field
		OfChatCompletionListResponseDataChoiceMessageToolContentArray      respjson.Field
		OfChatCompletionListResponseDataChoiceMessageDeveloperContentArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (r *ChatCompletionListResponseDataChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataChoiceMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionListResponseDataChoiceMessageUserContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageUserContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataChoiceMessageUserContentArray]
type ChatCompletionListResponseDataChoiceMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataChoiceMessageUserContentArray []ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionListResponseDataChoiceMessageUserContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentUnion) AsChatCompletionListResponseDataChoiceMessageUserContentArray() (v []ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion) {
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

// ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion contains
// all possible properties and values from
// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText],
// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessageUserContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessageUserContentArrayItem interface {
	implChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion()
}

func (ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText) implChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL) implChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataChoiceMessageUserContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                               `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionListResponseDataChoiceMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionListResponseDataChoiceMessageSystemContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageSystemContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataChoiceMessageSystemContentArray]
type ChatCompletionListResponseDataChoiceMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataChoiceMessageSystemContentArray []ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion `json:",inline"`
	JSON                                                            struct {
		OfString                                                        respjson.Field
		OfChatCompletionListResponseDataChoiceMessageSystemContentArray respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageSystemContentUnion) AsChatCompletionListResponseDataChoiceMessageSystemContentArray() (v []ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion) {
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

// ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion contains
// all possible properties and values from
// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText],
// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessageSystemContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessageSystemContentArrayItem interface {
	implChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion()
}

func (ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText) implChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL) implChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataChoiceMessageSystemContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                 `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageSystemContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataChoiceMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionListResponseDataChoiceMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionListResponseDataChoiceMessageAssistantToolCall `json:"tool_calls"`
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
func (r ChatCompletionListResponseDataChoiceMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataChoiceMessageAssistantContentArray]
type ChatCompletionListResponseDataChoiceMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataChoiceMessageAssistantContentArray []ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfChatCompletionListResponseDataChoiceMessageAssistantContentArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageAssistantContentUnion) AsChatCompletionListResponseDataChoiceMessageAssistantContentArray() (v []ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion) {
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

// ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText],
// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessageAssistantContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion] to
// add type safety for the return type of
// [ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessageAssistantContentArrayItem interface {
	implChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion()
}

func (ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText) implChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL) implChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataChoiceMessageAssistantContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageAssistantToolCall struct {
	Type     constant.Function                                                    `json:"type,required"`
	ID       string                                                               `json:"id"`
	Function ChatCompletionListResponseDataChoiceMessageAssistantToolCallFunction `json:"function"`
	Index    int64                                                                `json:"index"`
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
func (r ChatCompletionListResponseDataChoiceMessageAssistantToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
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
	// The response content from the tool
	Content ChatCompletionListResponseDataChoiceMessageToolContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageToolContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataChoiceMessageToolContentArray]
type ChatCompletionListResponseDataChoiceMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataChoiceMessageToolContentArray []ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfChatCompletionListResponseDataChoiceMessageToolContentArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageToolContentUnion) AsChatCompletionListResponseDataChoiceMessageToolContentArray() (v []ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion) {
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

// ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion contains
// all possible properties and values from
// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText],
// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessageToolContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessageToolContentArrayItem interface {
	implChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion()
}

func (ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText) implChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL) implChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataChoiceMessageToolContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageToolContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageToolContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                               `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageToolContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataChoiceMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataChoiceMessageDeveloperContentArray]
type ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataChoiceMessageDeveloperContentArray []ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfChatCompletionListResponseDataChoiceMessageDeveloperContentArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentUnion) AsChatCompletionListResponseDataChoiceMessageDeveloperContentArray() (v []ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion) {
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

// ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText],
// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion] to
// add type safety for the return type of
// [ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItem interface {
	implChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion()
}

func (ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText) implChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL) implChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageDeveloperContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The log probabilities for the tokens in the message
type ChatCompletionListResponseDataChoiceLogprobs struct {
	// (Optional) The log probabilities for the tokens in the message
	Content []ChatCompletionListResponseDataChoiceLogprobsContent `json:"content"`
	// (Optional) The log probabilities for the tokens in the message
	Refusal []ChatCompletionListResponseDataChoiceLogprobsRefusal `json:"refusal"`
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
type ChatCompletionListResponseDataChoiceLogprobsContent struct {
	Token       string                                                          `json:"token,required"`
	Logprob     float64                                                         `json:"logprob,required"`
	TopLogprobs []ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                         `json:"bytes"`
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
func (r ChatCompletionListResponseDataChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseDataChoiceLogprobsRefusal struct {
	Token       string                                                          `json:"token,required"`
	Logprob     float64                                                         `json:"logprob,required"`
	TopLogprobs []ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                         `json:"bytes"`
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
func (r ChatCompletionListResponseDataChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob struct {
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
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageAssistant].
	ToolCalls []ChatCompletionListResponseDataInputMessageAssistantToolCall `json:"tool_calls"`
	// This field is from variant [ChatCompletionListResponseDataInputMessageTool].
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
// OfChatCompletionListResponseDataInputMessageUserContentArray
// OfChatCompletionListResponseDataInputMessageSystemContentArray
// OfChatCompletionListResponseDataInputMessageAssistantContentArray
// OfChatCompletionListResponseDataInputMessageToolContentArray
// OfChatCompletionListResponseDataInputMessageDeveloperContentArray]
type ChatCompletionListResponseDataInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataInputMessageUserContentArray []ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataInputMessageSystemContentArray []ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataInputMessageAssistantContentArray []ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataInputMessageToolContentArray []ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataInputMessageDeveloperContentArray []ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion `json:",inline"`
	JSON                                                              struct {
		OfString                                                          respjson.Field
		OfChatCompletionListResponseDataInputMessageUserContentArray      respjson.Field
		OfChatCompletionListResponseDataInputMessageSystemContentArray    respjson.Field
		OfChatCompletionListResponseDataInputMessageAssistantContentArray respjson.Field
		OfChatCompletionListResponseDataInputMessageToolContentArray      respjson.Field
		OfChatCompletionListResponseDataInputMessageDeveloperContentArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (r *ChatCompletionListResponseDataInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataInputMessageUser struct {
	// The content of the message, which can include text and other media
	Content ChatCompletionListResponseDataInputMessageUserContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataInputMessageUser) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataInputMessageUserContentArray]
type ChatCompletionListResponseDataInputMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataInputMessageUserContentArray []ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfChatCompletionListResponseDataInputMessageUserContentArray respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUserContentUnion) AsChatCompletionListResponseDataInputMessageUserContentArray() (v []ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion) {
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

// ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionListResponseDataInputMessageUserContentArrayItemText],
// [ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageUserContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessageUserContentArrayItem is implemented
// by each variant of
// [ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataInputMessageUserContentArrayItem interface {
	implChatCompletionListResponseDataInputMessageUserContentArrayItemUnion()
}

func (ChatCompletionListResponseDataInputMessageUserContentArrayItemText) implChatCompletionListResponseDataInputMessageUserContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL) implChatCompletionListResponseDataInputMessageUserContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageUserContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataInputMessageUserContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataInputMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageUserContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataInputMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type ChatCompletionListResponseDataInputMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content ChatCompletionListResponseDataInputMessageSystemContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataInputMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageSystemContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataInputMessageSystemContentArray]
type ChatCompletionListResponseDataInputMessageSystemContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataInputMessageSystemContentArray []ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion `json:",inline"`
	JSON                                                           struct {
		OfString                                                       respjson.Field
		OfChatCompletionListResponseDataInputMessageSystemContentArray respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageSystemContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageSystemContentUnion) AsChatCompletionListResponseDataInputMessageSystemContentArray() (v []ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion) {
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

// ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion contains
// all possible properties and values from
// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemText],
// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessageSystemContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataInputMessageSystemContentArrayItem interface {
	implChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion()
}

func (ChatCompletionListResponseDataInputMessageSystemContentArrayItemText) implChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL) implChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageSystemContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataInputMessageSystemContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataInputMessageSystemContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageSystemContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageSystemContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataInputMessageSystemContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageSystemContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageSystemContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataInputMessageAssistant struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content ChatCompletionListResponseDataInputMessageAssistantContentUnion `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionListResponseDataInputMessageAssistantToolCall `json:"tool_calls"`
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
func (r ChatCompletionListResponseDataInputMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageAssistantContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataInputMessageAssistantContentArray]
type ChatCompletionListResponseDataInputMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataInputMessageAssistantContentArray []ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion `json:",inline"`
	JSON                                                              struct {
		OfString                                                          respjson.Field
		OfChatCompletionListResponseDataInputMessageAssistantContentArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageAssistantContentUnion) AsChatCompletionListResponseDataInputMessageAssistantContentArray() (v []ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion) {
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

// ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText],
// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessageAssistantContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion] to
// add type safety for the return type of
// [ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataInputMessageAssistantContentArrayItem interface {
	implChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion()
}

func (ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText) implChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL) implChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataInputMessageAssistantContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageAssistantToolCall struct {
	Type     constant.Function                                                   `json:"type,required"`
	ID       string                                                              `json:"id"`
	Function ChatCompletionListResponseDataInputMessageAssistantToolCallFunction `json:"function"`
	Index    int64                                                               `json:"index"`
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
func (r ChatCompletionListResponseDataInputMessageAssistantToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageAssistantToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
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
	// The response content from the tool
	Content ChatCompletionListResponseDataInputMessageToolContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataInputMessageTool) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageToolContentUnion contains all possible
// properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataInputMessageToolContentArray]
type ChatCompletionListResponseDataInputMessageToolContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion] instead
	// of an object.
	OfChatCompletionListResponseDataInputMessageToolContentArray []ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfChatCompletionListResponseDataInputMessageToolContentArray respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageToolContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageToolContentUnion) AsChatCompletionListResponseDataInputMessageToolContentArray() (v []ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion) {
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

// ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion contains all
// possible properties and values from
// [ChatCompletionListResponseDataInputMessageToolContentArrayItemText],
// [ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageToolContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessageToolContentArrayItem is implemented
// by each variant of
// [ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion] to add
// type safety for the return type of
// [ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataInputMessageToolContentArrayItem interface {
	implChatCompletionListResponseDataInputMessageToolContentArrayItemUnion()
}

func (ChatCompletionListResponseDataInputMessageToolContentArrayItemText) implChatCompletionListResponseDataInputMessageToolContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL) implChatCompletionListResponseDataInputMessageToolContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageToolContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataInputMessageToolContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataInputMessageToolContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageToolContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageToolContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataInputMessageToolContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageToolContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageToolContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type ChatCompletionListResponseDataInputMessageDeveloper struct {
	// The content of the developer message
	Content ChatCompletionListResponseDataInputMessageDeveloperContentUnion `json:"content,required"`
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
func (r ChatCompletionListResponseDataInputMessageDeveloper) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataInputMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatCompletionListResponseDataInputMessageDeveloperContentUnion contains all
// possible properties and values from [string],
// [[]ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfChatCompletionListResponseDataInputMessageDeveloperContentArray]
type ChatCompletionListResponseDataInputMessageDeveloperContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion]
	// instead of an object.
	OfChatCompletionListResponseDataInputMessageDeveloperContentArray []ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion `json:",inline"`
	JSON                                                              struct {
		OfString                                                          respjson.Field
		OfChatCompletionListResponseDataInputMessageDeveloperContentArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (u ChatCompletionListResponseDataInputMessageDeveloperContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageDeveloperContentUnion) AsChatCompletionListResponseDataInputMessageDeveloperContentArray() (v []ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion) {
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

// ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion
// contains all possible properties and values from
// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText],
// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL].
//
// Use the
// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion struct {
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText].
	Text string `json:"text"`
	// Any of "text", "image_url".
	Type string `json:"type"`
	// This field is from variant
	// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL].
	ImageURL ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyChatCompletionListResponseDataInputMessageDeveloperContentArrayItem is
// implemented by each variant of
// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion] to
// add type safety for the return type of
// [ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion.AsAny]
type anyChatCompletionListResponseDataInputMessageDeveloperContentArrayItem interface {
	implChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion()
}

func (ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText) implChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion() {
}
func (ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL) implChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText:
//	case llamastackclient.ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion) AsAny() anyChatCompletionListResponseDataInputMessageDeveloperContentArrayItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	}
	return nil
}

func (u ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion) AsText() (v ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion) AsImageURL() (v ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText struct {
	Text string        `json:"text,required"`
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
func (r ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL struct {
	ImageURL ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url,required"`
	Type     constant.ImageURL                                                                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURLImageURL struct {
	URL    string `json:"url,required"`
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
func (r ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURLImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageDeveloperContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
// [_[]ChatCompletionNewParamsMessageSystemContentArrayItemUnion],
// [_[]ChatCompletionNewParamsMessageAssistantContentArrayItemUnion],
// [_[]ChatCompletionNewParamsMessageToolContentArrayItemUnion],
// [\*[]ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion]
type chatCompletionNewParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageUserContentArrayItemUnion:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageSystemContentArrayItemUnion:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageAssistantContentArrayItemUnion:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageToolContentArrayItemUnion:
//	case *[]llamastackclient.ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion:
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
	paramUnion
}

func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL)
}
func (u *ChatCompletionNewParamsMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUserContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
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
func (u ChatCompletionNewParamsMessageUserContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageUserContentArrayItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentArrayItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageUserContentArrayItemImageURL]("image_url"),
	)
}

// The properties Text, Type are required.
type ChatCompletionNewParamsMessageUserContentArrayItemText struct {
	Text string `json:"text,required"`
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

// The properties ImageURL, Type are required.
type ChatCompletionNewParamsMessageUserContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
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

// The property URL is required.
type ChatCompletionNewParamsMessageUserContentArrayItemImageURLImageURL struct {
	URL    string            `json:"url,required"`
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
	OfString                                      param.Opt[string]                                           `json:",omitzero,inline"`
	OfChatCompletionNewsMessageSystemContentArray []ChatCompletionNewParamsMessageSystemContentArrayItemUnion `json:",omitzero,inline"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageSystemContentArrayItemUnion struct {
	OfText     *ChatCompletionNewParamsMessageSystemContentArrayItemText     `json:",omitzero,inline"`
	OfImageURL *ChatCompletionNewParamsMessageSystemContentArrayItemImageURL `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageSystemContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL)
}
func (u *ChatCompletionNewParamsMessageSystemContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageSystemContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageSystemContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageSystemContentArrayItemUnion) GetImageURL() *ChatCompletionNewParamsMessageSystemContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageSystemContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageSystemContentArrayItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageSystemContentArrayItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageSystemContentArrayItemImageURL]("image_url"),
	)
}

// The properties Text, Type are required.
type ChatCompletionNewParamsMessageSystemContentArrayItemText struct {
	Text string `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystemContentArrayItemText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystemContentArrayItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystemContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ImageURL, Type are required.
type ChatCompletionNewParamsMessageSystemContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewParamsMessageSystemContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystemContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystemContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystemContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ChatCompletionNewParamsMessageSystemContentArrayItemImageURLImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageSystemContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageSystemContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageSystemContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
	OfString                                         param.Opt[string]                                              `json:",omitzero,inline"`
	OfChatCompletionNewsMessageAssistantContentArray []ChatCompletionNewParamsMessageAssistantContentArrayItemUnion `json:",omitzero,inline"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageAssistantContentArrayItemUnion struct {
	OfText     *ChatCompletionNewParamsMessageAssistantContentArrayItemText     `json:",omitzero,inline"`
	OfImageURL *ChatCompletionNewParamsMessageAssistantContentArrayItemImageURL `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageAssistantContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL)
}
func (u *ChatCompletionNewParamsMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageAssistantContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageAssistantContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageAssistantContentArrayItemUnion) GetImageURL() *ChatCompletionNewParamsMessageAssistantContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageAssistantContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageAssistantContentArrayItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageAssistantContentArrayItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageAssistantContentArrayItemImageURL]("image_url"),
	)
}

// The properties Text, Type are required.
type ChatCompletionNewParamsMessageAssistantContentArrayItemText struct {
	Text string `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantContentArrayItemText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantContentArrayItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ImageURL, Type are required.
type ChatCompletionNewParamsMessageAssistantContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewParamsMessageAssistantContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ChatCompletionNewParamsMessageAssistantContentArrayItemImageURLImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageAssistantContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageAssistantContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageAssistantContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ChatCompletionNewParamsMessageAssistantToolCall struct {
	ID       param.Opt[string]                                       `json:"id,omitzero"`
	Index    param.Opt[int64]                                        `json:"index,omitzero"`
	Function ChatCompletionNewParamsMessageAssistantToolCallFunction `json:"function,omitzero"`
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
	OfString                                    param.Opt[string]                                         `json:",omitzero,inline"`
	OfChatCompletionNewsMessageToolContentArray []ChatCompletionNewParamsMessageToolContentArrayItemUnion `json:",omitzero,inline"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageToolContentArrayItemUnion struct {
	OfText     *ChatCompletionNewParamsMessageToolContentArrayItemText     `json:",omitzero,inline"`
	OfImageURL *ChatCompletionNewParamsMessageToolContentArrayItemImageURL `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageToolContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL)
}
func (u *ChatCompletionNewParamsMessageToolContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageToolContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageToolContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageToolContentArrayItemUnion) GetImageURL() *ChatCompletionNewParamsMessageToolContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageToolContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageToolContentArrayItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageToolContentArrayItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageToolContentArrayItemImageURL]("image_url"),
	)
}

// The properties Text, Type are required.
type ChatCompletionNewParamsMessageToolContentArrayItemText struct {
	Text string `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageToolContentArrayItemText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageToolContentArrayItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageToolContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ImageURL, Type are required.
type ChatCompletionNewParamsMessageToolContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewParamsMessageToolContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageToolContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageToolContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageToolContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ChatCompletionNewParamsMessageToolContentArrayItemImageURLImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageToolContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageToolContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageToolContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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
	OfString                                         param.Opt[string]                                              `json:",omitzero,inline"`
	OfChatCompletionNewsMessageDeveloperContentArray []ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion `json:",omitzero,inline"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion struct {
	OfText     *ChatCompletionNewParamsMessageDeveloperContentArrayItemText     `json:",omitzero,inline"`
	OfImageURL *ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURL `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL)
}
func (u *ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion) GetImageURL() *ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsMessageDeveloperContentArrayItemUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsMessageDeveloperContentArrayItemText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURL]("image_url"),
	)
}

// The properties Text, Type are required.
type ChatCompletionNewParamsMessageDeveloperContentArrayItemText struct {
	Text string `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloperContentArrayItemText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloperContentArrayItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloperContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ImageURL, Type are required.
type ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURL struct {
	ImageURL ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURLImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageDeveloperContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
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

// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsResponseFormatText].
type ChatCompletionNewParamsResponseFormatText struct {
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

// The properties JsonSchema, Type are required.
type ChatCompletionNewParamsResponseFormatJsonSchema struct {
	JsonSchema ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero,required"`
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

// The property Name is required.
type ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema struct {
	Name        string                                                                          `json:"name,required"`
	Description param.Opt[string]                                                               `json:"description,omitzero"`
	Strict      param.Opt[bool]                                                                 `json:"strict,omitzero"`
	Schema      map[string]ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion `json:"schema,omitzero"`
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

// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsResponseFormatJsonObject].
type ChatCompletionNewParamsResponseFormatJsonObject struct {
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
