// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// ModelOpenAIService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelOpenAIService] method instead.
type ModelOpenAIService struct {
	Options []option.RequestOption
}

// NewModelOpenAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelOpenAIService(opts ...option.RequestOption) (r ModelOpenAIService) {
	r = ModelOpenAIService{}
	r.Options = opts
	return
}

// List models using the OpenAI API.
func (r *ModelOpenAIService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ModelOpenAIListResponse, err error) {
	var env ModelOpenAIListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v1/openai/v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// A model from OpenAI.
type ModelOpenAIListResponse struct {
	ID      string         `json:"id,required"`
	Created int64          `json:"created,required"`
	Object  constant.Model `json:"object,required"`
	OwnedBy string         `json:"owned_by,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelOpenAIListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelOpenAIListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelOpenAIListResponseEnvelope struct {
	Data []ModelOpenAIListResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelOpenAIListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ModelOpenAIListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
