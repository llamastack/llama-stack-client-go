// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// ChatService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options     []option.RequestOption
	Completions ChatCompletionService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	r.Completions = NewChatCompletionService(opts...)
	return
}

// Chunk from a streaming response to an OpenAI-compatible chat completion request.
type ChatCompletionChunk struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []ChatCompletionChunkChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created int64 `json:"created,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion.chunk"
	Object constant.ChatCompletionChunk `json:"object,required"`
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
func (r ChatCompletionChunk) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk choice from an OpenAI-compatible chat completion streaming response.
type ChatCompletionChunkChoice struct {
	// The delta from the chunk
	Delta ChatCompletionChunkChoiceDelta `json:"delta,required"`
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChatCompletionChunkChoiceLogprobs `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta        respjson.Field
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The delta from the chunk
type ChatCompletionChunkChoiceDelta struct {
	// (Optional) The content of the delta
	Content string `json:"content"`
	// (Optional) The refusal of the delta
	Refusal string `json:"refusal"`
	// (Optional) The role of the delta
	Role string `json:"role"`
	// (Optional) The tool calls of the delta
	ToolCalls []ChatCompletionChunkChoiceDeltaToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkChoiceDelta) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionChunkChoiceDeltaToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionChunkChoiceDeltaToolCallFunction `json:"function"`
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
func (r ChatCompletionChunkChoiceDeltaToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDeltaToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionChunkChoiceDeltaToolCallFunction struct {
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
func (r ChatCompletionChunkChoiceDeltaToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDeltaToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The log probabilities for the tokens in the message
type ChatCompletionChunkChoiceLogprobs struct {
	// (Optional) The log probabilities for the tokens in the message
	Content []ChatCompletionChunkChoiceLogprobsContent `json:"content"`
	// (Optional) The log probabilities for the tokens in the message
	Refusal []ChatCompletionChunkChoiceLogprobsRefusal `json:"refusal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionChunkChoiceLogprobsContent struct {
	Token       string                                               `json:"token,required"`
	Logprob     float64                                              `json:"logprob,required"`
	TopLogprobs []ChatCompletionChunkChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                              `json:"bytes"`
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
func (r ChatCompletionChunkChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionChunkChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionChunkChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionChunkChoiceLogprobsRefusal struct {
	Token       string                                               `json:"token,required"`
	Logprob     float64                                              `json:"logprob,required"`
	TopLogprobs []ChatCompletionChunkChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                              `json:"bytes"`
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
func (r ChatCompletionChunkChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionChunkChoiceLogprobsRefusalTopLogprob struct {
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
func (r ChatCompletionChunkChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
