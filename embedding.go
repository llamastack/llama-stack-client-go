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
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
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
// EmbeddingService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r EmbeddingService) {
	r = EmbeddingService{}
	r.Options = opts
	return
}

// Generate OpenAI-compatible embeddings for the given input using the specified
// model.
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *CreateEmbeddingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response from an OpenAI-compatible embeddings request.
type CreateEmbeddingsResponse struct {
	// List of embedding data objects.
	Data []CreateEmbeddingsResponseData `json:"data" api:"required"`
	// The model that was used to generate the embeddings.
	Model string `json:"model" api:"required"`
	// Usage information.
	Usage CreateEmbeddingsResponseUsage `json:"usage" api:"required"`
	// The object type.
	//
	// Any of "list".
	Object CreateEmbeddingsResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Model       respjson.Field
		Usage       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateEmbeddingsResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single embedding data object from an OpenAI-compatible embeddings response.
type CreateEmbeddingsResponseData struct {
	// The embedding vector as a list of floats (when encoding_format='float') or as a
	// base64-encoded string.
	Embedding CreateEmbeddingsResponseDataEmbeddingUnion `json:"embedding" api:"required"`
	// The index of the embedding in the input list.
	Index int64 `json:"index" api:"required"`
	// The object type.
	//
	// Any of "embedding".
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding   respjson.Field
		Index       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateEmbeddingsResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateEmbeddingsResponseDataEmbeddingUnion contains all possible properties and
// values from [[]float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfListNumber OfString]
type CreateEmbeddingsResponseDataEmbeddingUnion struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	OfListNumber []float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfListNumber respjson.Field
		OfString     respjson.Field
		raw          string
	} `json:"-"`
}

func (u CreateEmbeddingsResponseDataEmbeddingUnion) AsListNumber() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreateEmbeddingsResponseDataEmbeddingUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CreateEmbeddingsResponseDataEmbeddingUnion) RawJSON() string { return u.JSON.raw }

func (r *CreateEmbeddingsResponseDataEmbeddingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage information.
type CreateEmbeddingsResponseUsage struct {
	// The number of tokens in the input.
	PromptTokens int64 `json:"prompt_tokens" api:"required"`
	// The total number of tokens used.
	TotalTokens int64 `json:"total_tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptTokens respjson.Field
		TotalTokens  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateEmbeddingsResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingsResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type CreateEmbeddingsResponseObject string

const (
	CreateEmbeddingsResponseObjectList CreateEmbeddingsResponseObject = "list"
)

type EmbeddingNewParams struct {
	// Input text to embed, encoded as a string or array of tokens.
	Input EmbeddingNewParamsInputUnion `json:"input,omitzero" api:"required"`
	// The identifier of the model to use.
	Model string `json:"model" api:"required"`
	// The number of dimensions for output embeddings.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// A unique identifier representing your end-user.
	User param.Opt[string] `json:"user,omitzero"`
	// The format to return the embeddings in.
	//
	// Any of "float", "base64".
	EncodingFormat EmbeddingNewParamsEncodingFormat `json:"encoding_format,omitzero"`
	paramObj
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbeddingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbeddingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmbeddingNewParamsInputUnion struct {
	OfString             param.Opt[string] `json:",omitzero,inline"`
	OfArrayOfStrings     []string          `json:",omitzero,inline"`
	OfArrayOfTokens      []int64           `json:",omitzero,inline"`
	OfArrayOfTokenArrays [][]int64         `json:",omitzero,inline"`
	paramUnion
}

func (u EmbeddingNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfArrayOfStrings, u.OfArrayOfTokens, u.OfArrayOfTokenArrays)
}
func (u *EmbeddingNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmbeddingNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfArrayOfStrings) {
		return &u.OfArrayOfStrings
	} else if !param.IsOmitted(u.OfArrayOfTokens) {
		return &u.OfArrayOfTokens
	} else if !param.IsOmitted(u.OfArrayOfTokenArrays) {
		return &u.OfArrayOfTokenArrays
	}
	return nil
}

// The format to return the embeddings in.
type EmbeddingNewParamsEncodingFormat string

const (
	EmbeddingNewParamsEncodingFormatFloat  EmbeddingNewParamsEncodingFormat = "float"
	EmbeddingNewParamsEncodingFormatBase64 EmbeddingNewParamsEncodingFormat = "base64"
)
