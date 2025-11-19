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

// Create completion.
//
// Generate an OpenAI-compatible completion for the given prompt using the
// specified model.
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *CompletionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create completion.
//
// Generate an OpenAI-compatible completion for the given prompt using the
// specified model.
func (r *CompletionService) NewStreaming(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[CompletionNewResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[CompletionNewResponse](ssestream.NewDecoder(raw), err)
}

// Response from an OpenAI-compatible completion request.
//
// :id: The ID of the completion :choices: List of choices :created: The Unix
// timestamp in seconds when the completion was created :model: The model that was
// used to generate the completion :object: The object type, which will be
// "text_completion"
type CompletionNewResponse struct {
	ID      string                        `json:"id,required"`
	Choices []CompletionNewResponseChoice `json:"choices,required"`
	Created int64                         `json:"created,required"`
	Model   string                        `json:"model,required"`
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
//
// :finish_reason: The reason the model stopped generating :text: The text of the
// choice :index: The index of the choice :logprobs: (Optional) The log
// probabilities for the tokens in the choice
type CompletionNewResponseChoice struct {
	FinishReason string `json:"finish_reason,required"`
	Index        int64  `json:"index,required"`
	Text         string `json:"text,required"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs CompletionNewResponseChoiceLogprobs `json:"logprobs,nullable"`
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
	Content []CompletionNewResponseChoiceLogprobsContent `json:"content,nullable"`
	Refusal []CompletionNewResponseChoiceLogprobsRefusal `json:"refusal,nullable"`
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
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type CompletionNewResponseChoiceLogprobsContent struct {
	Token       string                                                 `json:"token,required"`
	Logprob     float64                                                `json:"logprob,required"`
	TopLogprobs []CompletionNewResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                `json:"bytes,nullable"`
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
func (r CompletionNewResponseChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type CompletionNewResponseChoiceLogprobsContentTopLogprob struct {
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
func (r CompletionNewResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token :top_logprobs: The top log probabilities for the token
type CompletionNewResponseChoiceLogprobsRefusal struct {
	Token       string                                                 `json:"token,required"`
	Logprob     float64                                                `json:"logprob,required"`
	TopLogprobs []CompletionNewResponseChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                                `json:"bytes,nullable"`
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
func (r CompletionNewResponseChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
//
// :token: The token :bytes: (Optional) The bytes for the token :logprob: The log
// probability of the token
type CompletionNewResponseChoiceLogprobsRefusalTopLogprob struct {
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
func (r CompletionNewResponseChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionNewResponseObject string

const (
	CompletionNewResponseObjectTextCompletion CompletionNewResponseObject = "text_completion"
)

type CompletionNewParams struct {
	Model            string                         `json:"model,required"`
	Prompt           CompletionNewParamsPromptUnion `json:"prompt,omitzero,required"`
	BestOf           param.Opt[int64]               `json:"best_of,omitzero"`
	Echo             param.Opt[bool]                `json:"echo,omitzero"`
	FrequencyPenalty param.Opt[float64]             `json:"frequency_penalty,omitzero"`
	Logprobs         param.Opt[bool]                `json:"logprobs,omitzero"`
	MaxTokens        param.Opt[int64]               `json:"max_tokens,omitzero"`
	N                param.Opt[int64]               `json:"n,omitzero"`
	PresencePenalty  param.Opt[float64]             `json:"presence_penalty,omitzero"`
	Seed             param.Opt[int64]               `json:"seed,omitzero"`
	Suffix           param.Opt[string]              `json:"suffix,omitzero"`
	Temperature      param.Opt[float64]             `json:"temperature,omitzero"`
	TopP             param.Opt[float64]             `json:"top_p,omitzero"`
	User             param.Opt[string]              `json:"user,omitzero"`
	LogitBias        map[string]float64             `json:"logit_bias,omitzero"`
	Stop             CompletionNewParamsStopUnion   `json:"stop,omitzero"`
	StreamOptions    map[string]any                 `json:"stream_options,omitzero"`
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
