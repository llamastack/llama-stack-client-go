// Copyright (c) The OGX Contributors.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ogxclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ogx-ai/ogx-client-go/internal/apijson"
	"github.com/ogx-ai/ogx-client-go/internal/apiquery"
	"github.com/ogx-ai/ogx-client-go/internal/requestconfig"
	"github.com/ogx-ai/ogx-client-go/option"
	"github.com/ogx-ai/ogx-client-go/packages/param"
	"github.com/ogx-ai/ogx-client-go/packages/respjson"
)

// ModelOpenAIService contains methods and other services that help with
// interacting with the ogx-client API.
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

// List models. Returns OpenAI, Anthropic, or Google response format based on SDK
// detection headers.
func (r *ModelOpenAIService) List(ctx context.Context, params ModelOpenAIListParams, opts ...option.RequestOption) (res *ModelOpenAIListResponseUnion, err error) {
	if !param.IsOmitted(params.AnthropicVersion) {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%v", params.AnthropicVersion.Value)))
	}
	if !param.IsOmitted(params.XGoogAPIClient) {
		opts = append(opts, option.WithHeader("x-goog-api-client", fmt.Sprintf("%v", params.XGoogAPIClient.Value)))
	}
	if !param.IsOmitted(params.XGoogAPIKey) {
		opts = append(opts, option.WithHeader("x-goog-api-key", fmt.Sprintf("%v", params.XGoogAPIKey.Value)))
	}
	if !param.IsOmitted(params.XGoogUserProject) {
		opts = append(opts, option.WithHeader("x-goog-user-project", fmt.Sprintf("%v", params.XGoogUserProject.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// ModelOpenAIListResponseUnion contains all possible properties and values from
// [ListModelsResponse], [ModelOpenAIListResponseAnthropicListModelsResponse],
// [ModelOpenAIListResponseGoogleListModelsResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ModelOpenAIListResponseUnion struct {
	// This field is a union of [[]Model],
	// [[]ModelOpenAIListResponseAnthropicListModelsResponseData]
	Data ModelOpenAIListResponseUnionData `json:"data"`
	// This field is from variant [ListModelsResponse].
	Object ListModelsResponseObject `json:"object"`
	// This field is from variant [ModelOpenAIListResponseAnthropicListModelsResponse].
	FirstID string `json:"first_id"`
	// This field is from variant [ModelOpenAIListResponseAnthropicListModelsResponse].
	HasMore bool `json:"has_more"`
	// This field is from variant [ModelOpenAIListResponseAnthropicListModelsResponse].
	LastID string `json:"last_id"`
	// This field is from variant [ModelOpenAIListResponseGoogleListModelsResponse].
	Models []ModelOpenAIListResponseGoogleListModelsResponseModel `json:"models"`
	JSON   struct {
		Data    respjson.Field
		Object  respjson.Field
		FirstID respjson.Field
		HasMore respjson.Field
		LastID  respjson.Field
		Models  respjson.Field
		raw     string
	} `json:"-"`
}

func (u ModelOpenAIListResponseUnion) AsOpenAIListModelsResponse() (v ListModelsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelOpenAIListResponseUnion) AsAnthropicListModelsResponse() (v ModelOpenAIListResponseAnthropicListModelsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelOpenAIListResponseUnion) AsGoogleListModelsResponse() (v ModelOpenAIListResponseGoogleListModelsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelOpenAIListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelOpenAIListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelOpenAIListResponseUnionData is an implicit subunion of
// [ModelOpenAIListResponseUnion]. ModelOpenAIListResponseUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ModelOpenAIListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfData]
type ModelOpenAIListResponseUnionData struct {
	// This field will be present if the value is a [[]Model] instead of an object.
	OfData []Model `json:",inline"`
	JSON   struct {
		OfData respjson.Field
		raw    string
	} `json:"-"`
}

func (r *ModelOpenAIListResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of Anthropic model objects.
type ModelOpenAIListResponseAnthropicListModelsResponse struct {
	// List of Anthropic model objects.
	Data []ModelOpenAIListResponseAnthropicListModelsResponseData `json:"data" api:"required"`
	// First ID in the data list, usable as before_id for the previous page.
	FirstID string `json:"first_id" api:"nullable"`
	// Whether there are more results in the requested page direction.
	HasMore bool `json:"has_more"`
	// Last ID in the data list, usable as after_id for the next page.
	LastID string `json:"last_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelOpenAIListResponseAnthropicListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelOpenAIListResponseAnthropicListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Anthropic model info response object.
//
// :id: Unique model identifier
type ModelOpenAIListResponseAnthropicListModelsResponseData struct {
	// Unique model identifier.
	ID string `json:"id" api:"required"`
	// RFC 3339 datetime string representing when the model was released.
	CreatedAt string `json:"created_at" api:"required"`
	// A human-readable name for the model.
	DisplayName string `json:"display_name" api:"required"`
	// Maximum input context window size in tokens.
	MaxInputTokens int64 `json:"max_input_tokens" api:"nullable"`
	// Maximum value for the max_tokens parameter when using this model.
	MaxTokens int64 `json:"max_tokens" api:"nullable"`
	// Object type, always 'model'.
	//
	// Any of "model".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DisplayName    respjson.Field
		MaxInputTokens respjson.Field
		MaxTokens      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelOpenAIListResponseAnthropicListModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ModelOpenAIListResponseAnthropicListModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of Google model objects.
type ModelOpenAIListResponseGoogleListModelsResponse struct {
	// List of Google model objects.
	Models []ModelOpenAIListResponseGoogleListModelsResponseModel `json:"models" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Models      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelOpenAIListResponseGoogleListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelOpenAIListResponseGoogleListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Google model info response object.
//
// :name: Model resource name, e.g. 'models/gemini-pro' :display_name: A
// human-readable name for the model :description: A description of the model
type ModelOpenAIListResponseGoogleListModelsResponseModel struct {
	// A human-readable name for the model.
	DisplayName string `json:"display_name" api:"required"`
	// Model resource name, e.g. 'models/gemini-pro'.
	Name string `json:"name" api:"required"`
	// A description of the model.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelOpenAIListResponseGoogleListModelsResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ModelOpenAIListResponseGoogleListModelsResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelOpenAIListParams struct {
	// Return models after this model ID (Anthropic SDK format only).
	AfterID param.Opt[string] `query:"after_id,omitzero" json:"-"`
	// Return models before this model ID (Anthropic SDK format only).
	BeforeID param.Opt[string] `query:"before_id,omitzero" json:"-"`
	// Maximum number of models to return (Anthropic SDK format only).
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	AnthropicVersion param.Opt[string] `header:"anthropic-version,omitzero" json:"-"`
	XGoogAPIClient   param.Opt[string] `header:"x-goog-api-client,omitzero" json:"-"`
	XGoogAPIKey      param.Opt[string] `header:"x-goog-api-key,omitzero" json:"-"`
	XGoogUserProject param.Opt[string] `header:"x-goog-user-project,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ModelOpenAIListParams]'s query parameters as `url.Values`.
func (r ModelOpenAIListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
