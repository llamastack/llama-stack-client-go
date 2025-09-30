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

// InferenceService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceService] method instead.
type InferenceService struct {
	Options []option.RequestOption
}

// NewInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceService(opts ...option.RequestOption) (r InferenceService) {
	r = InferenceService{}
	r.Options = opts
	return
}

// Rerank a list of documents based on their relevance to a query.
func (r *InferenceService) Rerank(ctx context.Context, body InferenceRerankParams, opts ...option.RequestOption) (res *[]InferenceRerankResponse, err error) {
	var env InferenceRerankResponseEnvelope
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
type InferenceRerankResponse struct {
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
func (r InferenceRerankResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceRerankResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceRerankParams struct {
	// List of items to rerank. Each item can be a string, text content part, or image
	// content part. Each input must not exceed the model's max input token length.
	Items []InferenceRerankParamsItemUnion `json:"items,omitzero,required"`
	// The identifier of the reranking model to use.
	Model string `json:"model,required"`
	// The search query to rank items against. Can be a string, text content part, or
	// image content part. The input must not exceed the model's max input token
	// length.
	Query InferenceRerankParamsQueryUnion `json:"query,omitzero,required"`
	// (Optional) Maximum number of results to return. Default: returns all.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	paramObj
}

func (r InferenceRerankParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InferenceRerankParamsItemUnion struct {
	OfString                               param.Opt[string]                                                   `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartText  *InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceRerankParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *InferenceRerankParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceRerankParamsItemUnion) asAny() any {
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
func (u InferenceRerankParamsItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsItemUnion) GetImageURL() *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsItemUnion) GetType() *string {
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
type InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam struct {
	// Image URL specification and processing details
	ImageURL InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InferenceRerankParamsQueryUnion struct {
	OfString                               param.Opt[string]                                                    `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartText  *InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceRerankParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *InferenceRerankParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceRerankParamsQueryUnion) asAny() any {
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
func (u InferenceRerankParamsQueryUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsQueryUnion) GetImageURL() *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsQueryUnion) GetType() *string {
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
type InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam struct {
	// Image URL specification and processing details
	ImageURL InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from a reranking request.
type InferenceRerankResponseEnvelope struct {
	// List of rerank result objects, sorted by relevance score (descending)
	Data []InferenceRerankResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRerankResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *InferenceRerankResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
