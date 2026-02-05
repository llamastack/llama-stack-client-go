// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
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
	// The ID of the chat completion.
	ID string `json:"id,required"`
	// List of choices.
	Choices []ChatCompletionChunkChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created.
	Created int64 `json:"created,required"`
	// The model that was used to generate the chat completion.
	Model string `json:"model,required"`
	// The object type.
	//
	// Any of "chat.completion.chunk".
	Object ChatCompletionChunkObject `json:"object"`
	// Usage information for OpenAI chat completion.
	Usage ChatCompletionChunkUsage `json:"usage,nullable"`
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
func (r ChatCompletionChunk) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk choice from an OpenAI-compatible chat completion streaming response.
type ChatCompletionChunkChoice struct {
	// The delta from the chunk.
	Delta ChatCompletionChunkChoiceDelta `json:"delta,required"`
	// The index of the choice.
	Index int64 `json:"index,required"`
	// The reason the model stopped generating.
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason,nullable"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs ChatCompletionChunkChoiceLogprobs `json:"logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta        respjson.Field
		Index        respjson.Field
		FinishReason respjson.Field
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

// The delta from the chunk.
type ChatCompletionChunkChoiceDelta struct {
	// The content of the delta.
	Content string `json:"content,nullable"`
	// The reasoning content from the model (for o1/o3 models).
	ReasoningContent string `json:"reasoning_content,nullable"`
	// The refusal of the delta.
	Refusal string `json:"refusal,nullable"`
	// The role of the delta.
	Role string `json:"role,nullable"`
	// The tool calls of the delta.
	ToolCalls []ChatCompletionChunkChoiceDeltaToolCall `json:"tool_calls,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content          respjson.Field
		ReasoningContent respjson.Field
		Refusal          respjson.Field
		Role             respjson.Field
		ToolCalls        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkChoiceDelta) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionChunkChoiceDeltaToolCall struct {
	// Unique identifier for the tool call.
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionChunkChoiceDeltaToolCallFunction `json:"function,nullable"`
	// Index of the tool call in the list.
	Index int64 `json:"index,nullable"`
	// Must be 'function' to identify this as a function call.
	//
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
func (r ChatCompletionChunkChoiceDeltaToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDeltaToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionChunkChoiceDeltaToolCallFunction struct {
	// Arguments to pass to the function as a JSON string.
	Arguments string `json:"arguments,nullable"`
	// Name of the function to call.
	Name string `json:"name,nullable"`
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

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionChunkChoiceLogprobs struct {
	// The log probabilities for the tokens in the message.
	Content []ChatCompletionChunkChoiceLogprobsContent `json:"content,nullable"`
	// The log probabilities for the refusal tokens.
	Refusal []ChatCompletionChunkChoiceLogprobsRefusal `json:"refusal,nullable"`
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
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
	TopLogprobs []ChatCompletionChunkChoiceLogprobsContentTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionChunkChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionChunkChoiceLogprobsContentTopLogprob struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
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
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
	TopLogprobs []ChatCompletionChunkChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,nullable"`
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
func (r ChatCompletionChunkChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionChunkChoiceLogprobsRefusalTopLogprob struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
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

// The object type.
type ChatCompletionChunkObject string

const (
	ChatCompletionChunkObjectChatCompletionChunk ChatCompletionChunkObject = "chat.completion.chunk"
)

// Usage information for OpenAI chat completion.
type ChatCompletionChunkUsage struct {
	// Number of tokens in the completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total tokens used (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
	// Token details for output tokens in OpenAI chat completion usage.
	CompletionTokensDetails ChatCompletionChunkUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	// Token details for prompt tokens in OpenAI chat completion usage.
	PromptTokensDetails ChatCompletionChunkUsagePromptTokensDetails `json:"prompt_tokens_details,nullable"`
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
func (r ChatCompletionChunkUsage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for output tokens in OpenAI chat completion usage.
type ChatCompletionChunkUsageCompletionTokensDetails struct {
	// Number of tokens used for reasoning (o1/o3 models).
	ReasoningTokens int64 `json:"reasoning_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionChunkUsagePromptTokensDetails struct {
	// Number of tokens retrieved from cache.
	CachedTokens int64 `json:"cached_tokens,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
