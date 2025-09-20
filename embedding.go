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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

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
	path := "v1/openai/v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from an OpenAI-compatible embeddings request.
type CreateEmbeddingsResponse struct {
	// List of embedding data objects
	Data []CreateEmbeddingsResponseData `json:"data,required"`
	// The model that was used to generate the embeddings
	Model string `json:"model,required"`
	// The object type, which will be "list"
	Object constant.List `json:"object,required"`
	// Usage information
	Usage CreateEmbeddingsResponseUsage `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
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
	// The embedding vector as a list of floats (when encoding_format="float") or as a
	// base64-encoded string (when encoding_format="base64")
	Embedding CreateEmbeddingsResponseDataEmbeddingUnion `json:"embedding,required"`
	// The index of the embedding in the input list
	Index int64 `json:"index,required"`
	// The object type, which will be "embedding"
	Object constant.Embedding `json:"object,required"`
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
// will be valid: OfFloatArray OfString]
type CreateEmbeddingsResponseDataEmbeddingUnion struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloatArray respjson.Field
		OfString     respjson.Field
		raw          string
	} `json:"-"`
}

func (u CreateEmbeddingsResponseDataEmbeddingUnion) AsFloatArray() (v []float64) {
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

// Usage information
type CreateEmbeddingsResponseUsage struct {
	// The number of tokens in the input
	PromptTokens int64 `json:"prompt_tokens,required"`
	// The total number of tokens used
	TotalTokens int64 `json:"total_tokens,required"`
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

type EmbeddingNewParams struct {
	// Input text to embed, encoded as a string or array of strings. To embed multiple
	// inputs in a single request, pass an array of strings.
	Input EmbeddingNewParamsInputUnion `json:"input,omitzero,required"`
	// The identifier of the model to use. The model must be an embedding model
	// registered with Llama Stack and available via the /models endpoint.
	Model string `json:"model,required"`
	// (Optional) The number of dimensions the resulting output embeddings should have.
	// Only supported in text-embedding-3 and later models.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// (Optional) The format to return the embeddings in. Can be either "float" or
	// "base64". Defaults to "float".
	EncodingFormat param.Opt[string] `json:"encoding_format,omitzero"`
	// (Optional) A unique identifier representing your end-user, which can help OpenAI
	// to monitor and detect abuse.
	User param.Opt[string] `json:"user,omitzero"`
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
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u EmbeddingNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *EmbeddingNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmbeddingNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}
