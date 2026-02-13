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

// Generate an OpenAI-compatible chat completion for the given messages using the
// specified model.
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletionNewResponse, err error) {
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
	opts = append(opts, option.WithJSONSet("stream", true))
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

// List chat completions.
func (r *ChatCompletionService) List(ctx context.Context, query ChatCompletionListParams, opts ...option.RequestOption) (res *ChatCompletionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response from an OpenAI-compatible chat completion request.
type ChatCompletionNewResponse struct {
	// The ID of the chat completion.
	ID string `json:"id,required"`
	// List of choices.
	Choices []ChatCompletionNewResponseChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created.
	Created int64 `json:"created,required"`
	// The model that was used to generate the chat completion.
	Model string `json:"model,required"`
	// The object type.
	//
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
	// The reason the model stopped generating.
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice.
	Index int64 `json:"index,required"`
	// The message from the model.
	Message ChatCompletionNewResponseChoiceMessage `json:"message,required"`
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

// The message from the model.
type ChatCompletionNewResponseChoiceMessage struct {
	// Annotations for the message, when applicable.
	Annotations []map[string]any `json:"annotations"`
	// Audio response data when using audio output modality.
	Audio map[string]any `json:"audio,nullable"`
	// The content of the message.
	Content string `json:"content,nullable"`
	// Deprecated: the name and arguments of a function that should be called.
	FunctionCall ChatCompletionNewResponseChoiceMessageFunctionCall `json:"function_call"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal,nullable"`
	// The role of the message author, always 'assistant' in responses.
	//
	// Any of "assistant".
	Role string `json:"role"`
	// The tool calls generated by the model.
	ToolCalls []ChatCompletionNewResponseChoiceMessageToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations  respjson.Field
		Audio        respjson.Field
		Content      respjson.Field
		FunctionCall respjson.Field
		Refusal      respjson.Field
		Role         respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionNewResponseChoiceMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: the name and arguments of a function that should be called.
type ChatCompletionNewResponseChoiceMessageFunctionCall struct {
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
func (r ChatCompletionNewResponseChoiceMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionNewResponseChoiceMessageToolCall struct {
	// Unique identifier for the tool call.
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionNewResponseChoiceMessageToolCallFunction `json:"function,nullable"`
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
func (r ChatCompletionNewResponseChoiceMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionNewResponseChoiceMessageToolCallFunction struct {
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
func (r ChatCompletionNewResponseChoiceMessageToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionNewResponseChoiceLogprobs struct {
	// The log probabilities for the tokens in the message.
	Content []ChatCompletionNewResponseChoiceLogprobsContent `json:"content,nullable"`
	// The log probabilities for the refusal tokens.
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
type ChatCompletionNewResponseChoiceLogprobsContent struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
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
type ChatCompletionNewResponseChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionNewResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionNewResponseChoiceLogprobsRefusal struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
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
type ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob struct {
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
func (r ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type ChatCompletionNewResponseObject string

const (
	ChatCompletionNewResponseObjectChatCompletion ChatCompletionNewResponseObject = "chat.completion"
)

// Usage information for OpenAI chat completion.
type ChatCompletionNewResponseUsage struct {
	// Number of tokens in the completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total tokens used (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
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
func (r ChatCompletionNewResponseUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionNewResponseUsagePromptTokensDetails struct {
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
func (r ChatCompletionNewResponseUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionNewResponseUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionGetResponse struct {
	// The ID of the chat completion.
	ID string `json:"id,required"`
	// List of choices.
	Choices []ChatCompletionGetResponseChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created.
	Created int64 `json:"created,required"`
	// The input messages used to generate this completion.
	InputMessages []ChatCompletionGetResponseInputMessageUnion `json:"input_messages,required"`
	// The model that was used to generate the chat completion.
	Model string `json:"model,required"`
	// The object type.
	//
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
	// The reason the model stopped generating.
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice.
	Index int64 `json:"index,required"`
	// The message from the model.
	Message ChatCompletionGetResponseChoiceMessage `json:"message,required"`
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

// The message from the model.
type ChatCompletionGetResponseChoiceMessage struct {
	// Annotations for the message, when applicable.
	Annotations []map[string]any `json:"annotations"`
	// Audio response data when using audio output modality.
	Audio map[string]any `json:"audio,nullable"`
	// The content of the message.
	Content string `json:"content,nullable"`
	// Deprecated: the name and arguments of a function that should be called.
	FunctionCall ChatCompletionGetResponseChoiceMessageFunctionCall `json:"function_call"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal,nullable"`
	// The role of the message author, always 'assistant' in responses.
	//
	// Any of "assistant".
	Role string `json:"role"`
	// The tool calls generated by the model.
	ToolCalls []ChatCompletionGetResponseChoiceMessageToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations  respjson.Field
		Audio        respjson.Field
		Content      respjson.Field
		FunctionCall respjson.Field
		Refusal      respjson.Field
		Role         respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionGetResponseChoiceMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: the name and arguments of a function that should be called.
type ChatCompletionGetResponseChoiceMessageFunctionCall struct {
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
func (r ChatCompletionGetResponseChoiceMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionGetResponseChoiceMessageToolCall struct {
	// Unique identifier for the tool call.
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionGetResponseChoiceMessageToolCallFunction `json:"function,nullable"`
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
func (r ChatCompletionGetResponseChoiceMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionGetResponseChoiceMessageToolCallFunction struct {
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
func (r ChatCompletionGetResponseChoiceMessageToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionGetResponseChoiceLogprobs struct {
	// The log probabilities for the tokens in the message.
	Content []ChatCompletionGetResponseChoiceLogprobsContent `json:"content,nullable"`
	// The log probabilities for the refusal tokens.
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
type ChatCompletionGetResponseChoiceLogprobsContent struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
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
type ChatCompletionGetResponseChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionGetResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionGetResponseChoiceLogprobsRefusal struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
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
type ChatCompletionGetResponseChoiceLogprobsRefusalTopLogprob struct {
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
	// The content of the message, which can include text and other media.
	Content ChatCompletionGetResponseInputMessageUserContentUnion `json:"content,required"`
	// The name of the user message participant.
	Name string `json:"name,nullable"`
	// Must be 'user' to identify this as a user message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Image URL specification and processing details.
	ImageURL ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Must be 'image_url' to identify this as image content.
	//
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

// Image URL specification and processing details.
type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	// URL of the image to include in the message.
	URL string `json:"url,required"`
	// Level of detail for image processing. Can be 'low', 'high', or 'auto'.
	//
	// Any of "low", "high", "auto".
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
	// File specification.
	File ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Must be 'file' to identify this as file content.
	//
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

// File specification.
type ChatCompletionGetResponseInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	// Base64-encoded file data.
	FileData string `json:"file_data,nullable"`
	// ID of an uploaded file.
	FileID string `json:"file_id,nullable"`
	// Name of the file.
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
	// The content of the 'system prompt'. If multiple system messages are provided,
	// they are concatenated.
	Content ChatCompletionGetResponseInputMessageSystemContentUnion `json:"content,required"`
	// The name of the system message participant.
	Name string `json:"name,nullable"`
	// Must be 'system' to identify this as a system message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// The content of the model's response.
	Content ChatCompletionGetResponseInputMessageAssistantContentUnion `json:"content,nullable"`
	// The name of the assistant message participant.
	Name string `json:"name,nullable"`
	// Must be 'assistant' to identify this as the model's response.
	//
	// Any of "assistant".
	Role string `json:"role"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Unique identifier for the tool call.
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionGetResponseInputMessageAssistantToolCallFunction `json:"function,nullable"`
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
func (r ChatCompletionGetResponseInputMessageAssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionGetResponseInputMessageAssistantToolCallFunction struct {
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
func (r ChatCompletionGetResponseInputMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionGetResponseInputMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionGetResponseInputMessageTool struct {
	// The response content from the tool.
	Content ChatCompletionGetResponseInputMessageToolContentUnion `json:"content,required"`
	// Unique identifier for the tool call this response is for.
	ToolCallID string `json:"tool_call_id,required"`
	// Must be 'tool' to identify this as a tool response.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// The content of the developer message.
	Content ChatCompletionGetResponseInputMessageDeveloperContentUnion `json:"content,required"`
	// The name of the developer message participant.
	Name string `json:"name,nullable"`
	// Must be 'developer' to identify this as a developer message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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

// The object type.
type ChatCompletionGetResponseObject string

const (
	ChatCompletionGetResponseObjectChatCompletion ChatCompletionGetResponseObject = "chat.completion"
)

// Usage information for OpenAI chat completion.
type ChatCompletionGetResponseUsage struct {
	// Number of tokens in the completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total tokens used (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
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
func (r ChatCompletionGetResponseUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionGetResponseUsagePromptTokensDetails struct {
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
func (r ChatCompletionGetResponseUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionGetResponseUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from listing OpenAI-compatible chat completions.
type ChatCompletionListResponse struct {
	// List of chat completion objects with their input messages.
	Data []ChatCompletionListResponseData `json:"data,required"`
	// ID of the first completion in this list.
	FirstID string `json:"first_id,required"`
	// Whether there are more completions available beyond this list.
	HasMore bool `json:"has_more,required"`
	// ID of the last completion in this list.
	LastID string `json:"last_id,required"`
	// Must be 'list' to identify this as a list response.
	//
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
	// The ID of the chat completion.
	ID string `json:"id,required"`
	// List of choices.
	Choices []ChatCompletionListResponseDataChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created.
	Created int64 `json:"created,required"`
	// The input messages used to generate this completion.
	InputMessages []ChatCompletionListResponseDataInputMessageUnion `json:"input_messages,required"`
	// The model that was used to generate the chat completion.
	Model string `json:"model,required"`
	// The object type.
	//
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
	// The reason the model stopped generating.
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice.
	Index int64 `json:"index,required"`
	// The message from the model.
	Message ChatCompletionListResponseDataChoiceMessage `json:"message,required"`
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

// The message from the model.
type ChatCompletionListResponseDataChoiceMessage struct {
	// Annotations for the message, when applicable.
	Annotations []map[string]any `json:"annotations"`
	// Audio response data when using audio output modality.
	Audio map[string]any `json:"audio,nullable"`
	// The content of the message.
	Content string `json:"content,nullable"`
	// Deprecated: the name and arguments of a function that should be called.
	FunctionCall ChatCompletionListResponseDataChoiceMessageFunctionCall `json:"function_call"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal,nullable"`
	// The role of the message author, always 'assistant' in responses.
	//
	// Any of "assistant".
	Role string `json:"role"`
	// The tool calls generated by the model.
	ToolCalls []ChatCompletionListResponseDataChoiceMessageToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations  respjson.Field
		Audio        respjson.Field
		Content      respjson.Field
		FunctionCall respjson.Field
		Refusal      respjson.Field
		Role         respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionListResponseDataChoiceMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: the name and arguments of a function that should be called.
type ChatCompletionListResponseDataChoiceMessageFunctionCall struct {
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
func (r ChatCompletionListResponseDataChoiceMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionListResponseDataChoiceMessageToolCall struct {
	// Unique identifier for the tool call.
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionListResponseDataChoiceMessageToolCallFunction `json:"function,nullable"`
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
func (r ChatCompletionListResponseDataChoiceMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataChoiceMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionListResponseDataChoiceMessageToolCallFunction struct {
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
func (r ChatCompletionListResponseDataChoiceMessageToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChatCompletionListResponseDataChoiceLogprobs struct {
	// The log probabilities for the tokens in the message.
	Content []ChatCompletionListResponseDataChoiceLogprobsContent `json:"content,nullable"`
	// The log probabilities for the refusal tokens.
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
type ChatCompletionListResponseDataChoiceLogprobsContent struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
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
type ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob struct {
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
func (r ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type ChatCompletionListResponseDataChoiceLogprobsRefusal struct {
	// The token.
	Token string `json:"token,required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob,required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes,nullable"`
	// The top log probabilities for the token.
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
type ChatCompletionListResponseDataChoiceLogprobsRefusalTopLogprob struct {
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
	// The content of the message, which can include text and other media.
	Content ChatCompletionListResponseDataInputMessageUserContentUnion `json:"content,required"`
	// The name of the user message participant.
	Name string `json:"name,nullable"`
	// Must be 'user' to identify this as a user message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Image URL specification and processing details.
	ImageURL ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,required"`
	// Must be 'image_url' to identify this as image content.
	//
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

// Image URL specification and processing details.
type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	// URL of the image to include in the message.
	URL string `json:"url,required"`
	// Level of detail for image processing. Can be 'low', 'high', or 'auto'.
	//
	// Any of "low", "high", "auto".
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
	// File specification.
	File ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,required"`
	// Must be 'file' to identify this as file content.
	//
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

// File specification.
type ChatCompletionListResponseDataInputMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	// Base64-encoded file data.
	FileData string `json:"file_data,nullable"`
	// ID of an uploaded file.
	FileID string `json:"file_id,nullable"`
	// Name of the file.
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
	// The content of the 'system prompt'. If multiple system messages are provided,
	// they are concatenated.
	Content ChatCompletionListResponseDataInputMessageSystemContentUnion `json:"content,required"`
	// The name of the system message participant.
	Name string `json:"name,nullable"`
	// Must be 'system' to identify this as a system message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// The content of the model's response.
	Content ChatCompletionListResponseDataInputMessageAssistantContentUnion `json:"content,nullable"`
	// The name of the assistant message participant.
	Name string `json:"name,nullable"`
	// Must be 'assistant' to identify this as the model's response.
	//
	// Any of "assistant".
	Role string `json:"role"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Unique identifier for the tool call.
	ID string `json:"id,nullable"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionListResponseDataInputMessageAssistantToolCallFunction `json:"function,nullable"`
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
func (r ChatCompletionListResponseDataInputMessageAssistantToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function call details for OpenAI-compatible tool calls.
type ChatCompletionListResponseDataInputMessageAssistantToolCallFunction struct {
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
func (r ChatCompletionListResponseDataInputMessageAssistantToolCallFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataInputMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type ChatCompletionListResponseDataInputMessageTool struct {
	// The response content from the tool.
	Content ChatCompletionListResponseDataInputMessageToolContentUnion `json:"content,required"`
	// Unique identifier for the tool call this response is for.
	ToolCallID string `json:"tool_call_id,required"`
	// Must be 'tool' to identify this as a tool response.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// The content of the developer message.
	Content ChatCompletionListResponseDataInputMessageDeveloperContentUnion `json:"content,required"`
	// The name of the developer message participant.
	Name string `json:"name,nullable"`
	// Must be 'developer' to identify this as a developer message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Number of tokens in the completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total tokens used (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
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
func (r ChatCompletionListResponseDataUsageCompletionTokensDetails) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatCompletionListResponseDataUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token details for prompt tokens in OpenAI chat completion usage.
type ChatCompletionListResponseDataUsagePromptTokensDetails struct {
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
func (r ChatCompletionListResponseDataUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionListResponseDataUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be 'list' to identify this as a list response.
type ChatCompletionListResponseObject string

const (
	ChatCompletionListResponseObjectList ChatCompletionListResponseObject = "list"
)

type ChatCompletionNewParams struct {
	// List of messages in the conversation.
	Messages []ChatCompletionNewParamsMessageUnion `json:"messages,omitzero,required"`
	// The identifier of the model to use.
	Model string `json:"model,required"`
	// The penalty for repeated tokens.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// The log probabilities to use.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// The maximum number of tokens to generate.
	MaxCompletionTokens param.Opt[int64] `json:"max_completion_tokens,omitzero"`
	// The maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// The number of completions to generate.
	N param.Opt[int64] `json:"n,omitzero"`
	// Whether to parallelize tool calls.
	ParallelToolCalls param.Opt[bool] `json:"parallel_tool_calls,omitzero"`
	// The penalty for repeated tokens.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// A key to use when reading from or writing to the prompt cache.
	PromptCacheKey param.Opt[string] `json:"prompt_cache_key,omitzero"`
	// A stable identifier used for safety monitoring and abuse detection.
	SafetyIdentifier param.Opt[string] `json:"safety_identifier,omitzero"`
	// The seed to use.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The temperature to use.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// The top log probabilities to use.
	TopLogprobs param.Opt[int64] `json:"top_logprobs,omitzero"`
	// The top p to use.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// The user to use.
	User param.Opt[string] `json:"user,omitzero"`
	// The function call to use.
	FunctionCall ChatCompletionNewParamsFunctionCallUnion `json:"function_call,omitzero"`
	// List of functions to use.
	Functions []map[string]any `json:"functions,omitzero"`
	// The logit bias to use.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// The effort level for reasoning models.
	//
	// Any of "none", "minimal", "low", "medium", "high", "xhigh".
	ReasoningEffort ChatCompletionNewParamsReasoningEffort `json:"reasoning_effort,omitzero"`
	// The response format to use.
	ResponseFormat ChatCompletionNewParamsResponseFormatUnion `json:"response_format,omitzero"`
	// The stop tokens to use.
	Stop ChatCompletionNewParamsStopUnion `json:"stop,omitzero"`
	// The stream options to use.
	StreamOptions map[string]any `json:"stream_options,omitzero"`
	// The tool choice to use.
	ToolChoice ChatCompletionNewParamsToolChoiceUnion `json:"tool_choice,omitzero"`
	// The tools to use.
	Tools []map[string]any `json:"tools,omitzero"`
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
	// The content of the message, which can include text and other media.
	Content ChatCompletionNewParamsMessageUserContentUnion `json:"content,omitzero,required"`
	// The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be 'user' to identify this as a user message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Image URL specification and processing details.
	ImageURL ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Must be 'image_url' to identify this as image content.
	//
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

// Image URL specification and processing details.
//
// The property URL is required.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	// URL of the image to include in the message.
	URL string `json:"url,required"`
	// Level of detail for image processing. Can be 'low', 'high', or 'auto'.
	//
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL](
		"detail", "low", "high", "auto",
	)
}

// The property File is required.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	// File specification.
	File ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,omitzero,required"`
	// Must be 'file' to identify this as file content.
	//
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

// File specification.
type ChatCompletionNewParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	// Base64-encoded file data.
	FileData param.Opt[string] `json:"file_data,omitzero"`
	// ID of an uploaded file.
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// Name of the file.
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
	// The content of the 'system prompt'. If multiple system messages are provided,
	// they are concatenated.
	Content ChatCompletionNewParamsMessageSystemContentUnion `json:"content,omitzero,required"`
	// The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be 'system' to identify this as a system message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response.
	Content ChatCompletionNewParamsMessageAssistantContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionNewParamsMessageAssistantToolCall `json:"tool_calls,omitzero"`
	// Must be 'assistant' to identify this as the model's response.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// Unique identifier for the tool call.
	ID param.Opt[string] `json:"id,omitzero"`
	// Index of the tool call in the list.
	Index param.Opt[int64] `json:"index,omitzero"`
	// Function call details for OpenAI-compatible tool calls.
	Function ChatCompletionNewParamsMessageAssistantToolCallFunction `json:"function,omitzero"`
	// Must be 'function' to identify this as a function call.
	//
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
	// Arguments to pass to the function as a JSON string.
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	// Name of the function to call.
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
// The properties Content, ToolCallID are required.
type ChatCompletionNewParamsMessageTool struct {
	// The response content from the tool.
	Content ChatCompletionNewParamsMessageToolContentUnion `json:"content,omitzero,required"`
	// Unique identifier for the tool call this response is for.
	ToolCallID string `json:"tool_call_id,required"`
	// Must be 'tool' to identify this as a tool response.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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
	// The content of the developer message.
	Content ChatCompletionNewParamsMessageDeveloperContentUnion `json:"content,omitzero,required"`
	// The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be 'developer' to identify this as a developer message.
	//
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
	// The text content of the message.
	Text string `json:"text,required"`
	// Must be 'text' to identify this as text content.
	//
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

// The effort level for reasoning models.
type ChatCompletionNewParamsReasoningEffort string

const (
	ChatCompletionNewParamsReasoningEffortNone    ChatCompletionNewParamsReasoningEffort = "none"
	ChatCompletionNewParamsReasoningEffortMinimal ChatCompletionNewParamsReasoningEffort = "minimal"
	ChatCompletionNewParamsReasoningEffortLow     ChatCompletionNewParamsReasoningEffort = "low"
	ChatCompletionNewParamsReasoningEffortMedium  ChatCompletionNewParamsReasoningEffort = "medium"
	ChatCompletionNewParamsReasoningEffortHigh    ChatCompletionNewParamsReasoningEffort = "high"
	ChatCompletionNewParamsReasoningEffortXhigh   ChatCompletionNewParamsReasoningEffort = "xhigh"
)

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
	// Must be 'text' to indicate plain text response format.
	//
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
	// The JSON schema specification for the response.
	JsonSchema ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero,required"`
	// Must be 'json_schema' to indicate structured JSON response format.
	//
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

// The JSON schema specification for the response.
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
	// Must be 'json_object' to indicate generic JSON object response format.
	//
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
