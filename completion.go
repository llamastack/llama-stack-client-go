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
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
)

// Llama Stack Inference API for generating completions, chat completions, and
// embeddings.
//
// This API provides the raw interface to the underlying models. Three kinds of
// models are supported:
//
//   - LLM models: these models generate "raw" and "chat" (conversational)
//     completions.
//   - Embedding models: these models generate embeddings to be used for semantic
//     search.
//   - Rerank models: these models reorder the documents based on their relevance to
//     a query.
//
// CompletionService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompletionService] method instead.
type CompletionService struct {
	Options []option.RequestOption
}

// NewCompletionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompletionService(opts ...option.RequestOption) (r CompletionService) {
	r = CompletionService{}
	r.Options = opts
	return
}

// Generate an OpenAI-compatible completion for the given prompt using the
// specified model.
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *CompletionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generate an OpenAI-compatible completion for the given prompt using the
// specified model.
func (r *CompletionService) NewStreaming(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[CompletionNewResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append(opts, option.WithJSONSet("stream", true))
	path := "v1/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[CompletionNewResponse](ssestream.NewDecoder(raw), err)
}

// Response from an OpenAI-compatible completion request.
type CompletionNewResponse struct {
	// The ID of the completion.
	ID string `json:"id" api:"required"`
	// List of choices.
	Choices []CompletionNewResponseChoice `json:"choices" api:"required"`
	// The Unix timestamp in seconds when the completion was created.
	Created int64 `json:"created" api:"required"`
	// The model that was used to generate the completion.
	Model string `json:"model" api:"required"`
	// The object type.
	//
	// Any of "text_completion".
	Object CompletionNewResponseObject `json:"object"`
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
func (r CompletionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible completion response.
type CompletionNewResponseChoice struct {
	// The reason the model stopped generating.
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason" api:"required"`
	// The index of the choice.
	Index int64 `json:"index" api:"required"`
	// The text of the choice.
	Text string `json:"text" api:"required"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs CompletionNewResponseChoiceLogprobs `json:"logprobs" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Text         respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionNewResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type CompletionNewResponseChoiceLogprobs struct {
	// The log probabilities for the tokens in the message.
	Content []CompletionNewResponseChoiceLogprobsContent `json:"content" api:"nullable"`
	// The log probabilities for the refusal tokens.
	Refusal []CompletionNewResponseChoiceLogprobsRefusal `json:"refusal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionNewResponseChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type CompletionNewResponseChoiceLogprobsContent struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
	// The top log probabilities for the token.
	TopLogprobs []CompletionNewResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs" api:"nullable"`
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
func (r CompletionNewResponseChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type CompletionNewResponseChoiceLogprobsContentTopLogprob struct {
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
func (r CompletionNewResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type CompletionNewResponseChoiceLogprobsRefusal struct {
	// The token.
	Token string `json:"token" api:"required"`
	// The log probability of the token.
	Logprob float64 `json:"logprob" api:"required"`
	// The bytes for the token.
	Bytes []int64 `json:"bytes" api:"nullable"`
	// The top log probabilities for the token.
	TopLogprobs []CompletionNewResponseChoiceLogprobsRefusalTopLogprob `json:"top_logprobs" api:"nullable"`
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
func (r CompletionNewResponseChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type CompletionNewResponseChoiceLogprobsRefusalTopLogprob struct {
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
func (r CompletionNewResponseChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type CompletionNewResponseObject string

const (
	CompletionNewResponseObjectTextCompletion CompletionNewResponseObject = "text_completion"
)

type CompletionNewParams struct {
	// The identifier of the model to use.
	Model string `json:"model" api:"required"`
	// The prompt to generate a completion for.
	Prompt CompletionNewParamsPromptUnion `json:"prompt,omitzero" api:"required"`
	// The number of completions to generate.
	BestOf param.Opt[int64] `json:"best_of,omitzero"`
	// Whether to echo the prompt.
	Echo param.Opt[bool] `json:"echo,omitzero"`
	// The penalty for repeated tokens.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// The log probabilities to use.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// The maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// The number of completions to generate.
	N param.Opt[int64] `json:"n,omitzero"`
	// The penalty for repeated tokens.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// The seed to use.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The suffix that should be appended to the completion.
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// The temperature to use.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// The top p to use.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// The user to use.
	User param.Opt[string] `json:"user,omitzero"`
	// The logit bias to use.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// The stop tokens to use.
	Stop CompletionNewParamsStopUnion `json:"stop,omitzero"`
	// The stream options to use.
	StreamOptions map[string]any `json:"stream_options,omitzero"`
	paramObj
}

func (r CompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CompletionNewParamsPromptUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfListString  []string          `json:",omitzero,inline"`
	OfListInteger []int64           `json:",omitzero,inline"`
	OfListArray   [][]int64         `json:",omitzero,inline"`
	paramUnion
}

func (u CompletionNewParamsPromptUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListString, u.OfListInteger, u.OfListArray)
}
func (u *CompletionNewParamsPromptUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CompletionNewParamsPromptUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListString) {
		return &u.OfListString
	} else if !param.IsOmitted(u.OfListInteger) {
		return &u.OfListInteger
	} else if !param.IsOmitted(u.OfListArray) {
		return &u.OfListArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CompletionNewParamsStopUnion struct {
	OfString     param.Opt[string] `json:",omitzero,inline"`
	OfListString []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CompletionNewParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListString)
}
func (u *CompletionNewParamsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CompletionNewParamsStopUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListString) {
		return &u.OfListString
	}
	return nil
}
