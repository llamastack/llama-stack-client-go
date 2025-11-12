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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// AlphaInferenceService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaInferenceService] method instead.
type AlphaInferenceService struct {
	Options []option.RequestOption
}

// NewAlphaInferenceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaInferenceService(opts ...option.RequestOption) (r AlphaInferenceService) {
	r = AlphaInferenceService{}
	r.Options = opts
	return
}

// Rerank a list of documents based on their relevance to a query.
func (r *AlphaInferenceService) Rerank(ctx context.Context, body AlphaInferenceRerankParams, opts ...option.RequestOption) (res *[]AlphaInferenceRerankResponse, err error) {
	var env AlphaInferenceRerankResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/inference/rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// A single rerank result from a reranking response.
type AlphaInferenceRerankResponse struct {
	// The original index of the document in the input list
	Index int64 `json:"index,required"`
	// The relevance score from the model output. Values are inverted when applicable
	// so that higher scores indicate greater relevance.
	RelevanceScore float64 `json:"relevance_score,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index          respjson.Field
		RelevanceScore respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaInferenceRerankResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaInferenceRerankResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaInferenceRerankParams struct {
	// List of items to rerank. Each item can be a string, text content part, or image
	// content part. Each input must not exceed the model's max input token length.
	Items []AlphaInferenceRerankParamsItemUnion `json:"items,omitzero,required"`
	// The identifier of the reranking model to use.
	Model string `json:"model,required"`
	// The search query to rank items against. Can be a string, text content part, or
	// image content part. The input must not exceed the model's max input token
	// length.
	Query AlphaInferenceRerankParamsQueryUnion `json:"query,omitzero,required"`
	// (Optional) Maximum number of results to return. Default: returns all.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	paramObj
}

func (r AlphaInferenceRerankParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaInferenceRerankParamsItemUnion struct {
	OfString                               param.Opt[string]                                                        `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartText  *AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaInferenceRerankParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *AlphaInferenceRerankParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaInferenceRerankParamsItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaInferenceRerankParamsItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaInferenceRerankParamsItemUnion) GetImageURL() *AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaInferenceRerankParamsItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam struct {
	// Image URL specification and processing details
	ImageURL AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaInferenceRerankParamsQueryUnion struct {
	OfString                               param.Opt[string]                                                         `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartText  *AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaInferenceRerankParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *AlphaInferenceRerankParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaInferenceRerankParamsQueryUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaInferenceRerankParamsQueryUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaInferenceRerankParamsQueryUnion) GetImageURL() *AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaInferenceRerankParamsQueryUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam struct {
	// Image URL specification and processing details
	ImageURL AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaInferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from a reranking request.
type AlphaInferenceRerankResponseEnvelope struct {
	// List of rerank result objects, sorted by relevance score (descending)
	Data []AlphaInferenceRerankResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaInferenceRerankResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AlphaInferenceRerankResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
