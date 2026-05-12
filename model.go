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
	"errors"
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

// ModelService contains methods and other services that help with interacting with
// the ogx-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
	OpenAI  ModelOpenAIService
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	r.OpenAI = NewModelOpenAIService(opts...)
	return
}

// Get a model by its identifier. Returns OpenAI, Anthropic, or Google response
// format based on SDK detection headers.
func (r *ModelService) Get(ctx context.Context, modelID string, query ModelGetParams, opts ...option.RequestOption) (res *ModelGetResponseUnion, err error) {
	if !param.IsOmitted(query.AnthropicVersion) {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%v", query.AnthropicVersion.Value)))
	}
	if !param.IsOmitted(query.XGoogAPIClient) {
		opts = append(opts, option.WithHeader("x-goog-api-client", fmt.Sprintf("%v", query.XGoogAPIClient.Value)))
	}
	if !param.IsOmitted(query.XGoogAPIKey) {
		opts = append(opts, option.WithHeader("x-goog-api-key", fmt.Sprintf("%v", query.XGoogAPIKey.Value)))
	}
	if !param.IsOmitted(query.XGoogUserProject) {
		opts = append(opts, option.WithHeader("x-goog-user-project", fmt.Sprintf("%v", query.XGoogUserProject.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List models. Returns OpenAI, Anthropic, or Google response format based on SDK
// detection headers.
func (r *ModelService) List(ctx context.Context, params ModelListParams, opts ...option.RequestOption) (res *ModelListResponseUnion, err error) {
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

// Response containing a list of OpenAI model objects.
type ListModelsResponse struct {
	// List of OpenAI model objects.
	Data []Model `json:"data" api:"required"`
	// Any of "list".
	Object ListModelsResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListModelsResponseObject string

const (
	ListModelsResponseObjectList ListModelsResponseObject = "list"
)

// A model from OpenAI.
//
// :id: The ID of the model :object: The object type, which will be "model"
// :created: The Unix timestamp in seconds when the model was created :owned_by:
// The owner of the model :custom_metadata: OGX-specific metadata including
// model_type, provider info, and additional metadata
type Model struct {
	ID             string         `json:"id" api:"required"`
	Created        int64          `json:"created" api:"required"`
	OwnedBy        string         `json:"owned_by" api:"required"`
	CustomMetadata map[string]any `json:"custom_metadata" api:"nullable"`
	// Any of "model".
	Object ModelObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Created        respjson.Field
		OwnedBy        respjson.Field
		CustomMetadata respjson.Field
		Object         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelObject string

const (
	ModelObjectModel ModelObject = "model"
)

// ModelGetResponseUnion contains all possible properties and values from
// [ModelGetResponseModel], [ModelGetResponseAnthropicModelInfo],
// [ModelGetResponseGoogleModelInfo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ModelGetResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [ModelGetResponseModel].
	Identifier string `json:"identifier"`
	// This field is from variant [ModelGetResponseModel].
	Object string `json:"object"`
	// This field is from variant [ModelGetResponseModel].
	ProviderID string `json:"provider_id"`
	// This field is from variant [ModelGetResponseModel].
	Created int64 `json:"created"`
	// This field is from variant [ModelGetResponseModel].
	Metadata map[string]any `json:"metadata"`
	// This field is from variant [ModelGetResponseModel].
	ModelType string `json:"model_type"`
	// This field is from variant [ModelGetResponseModel].
	ModelValidation bool `json:"model_validation"`
	// This field is from variant [ModelGetResponseModel].
	OwnedBy string `json:"owned_by"`
	// This field is from variant [ModelGetResponseModel].
	ProviderResourceID string `json:"provider_resource_id"`
	Type               string `json:"type"`
	// This field is from variant [ModelGetResponseAnthropicModelInfo].
	CreatedAt   string `json:"created_at"`
	DisplayName string `json:"display_name"`
	// This field is from variant [ModelGetResponseAnthropicModelInfo].
	MaxInputTokens int64 `json:"max_input_tokens"`
	// This field is from variant [ModelGetResponseAnthropicModelInfo].
	MaxTokens int64 `json:"max_tokens"`
	// This field is from variant [ModelGetResponseGoogleModelInfo].
	Name string `json:"name"`
	// This field is from variant [ModelGetResponseGoogleModelInfo].
	Description string `json:"description"`
	JSON        struct {
		ID                 respjson.Field
		Identifier         respjson.Field
		Object             respjson.Field
		ProviderID         respjson.Field
		Created            respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
		ModelValidation    respjson.Field
		OwnedBy            respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		CreatedAt          respjson.Field
		DisplayName        respjson.Field
		MaxInputTokens     respjson.Field
		MaxTokens          respjson.Field
		Name               respjson.Field
		Description        respjson.Field
		raw                string
	} `json:"-"`
}

func (u ModelGetResponseUnion) AsModel() (v ModelGetResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelGetResponseUnion) AsAnthropicModelInfo() (v ModelGetResponseAnthropicModelInfo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelGetResponseUnion) AsGoogleModelInfo() (v ModelGetResponseGoogleModelInfo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model resource representing an AI model registered in OGX.
type ModelGetResponseModel struct {
	// The model identifier (OpenAI-compatible alias for identifier).
	ID string `json:"id" api:"required"`
	// Unique identifier for this resource in ogx
	Identifier string `json:"identifier" api:"required"`
	// The object type, always 'model'.
	//
	// Any of "model".
	Object string `json:"object" api:"required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id" api:"required"`
	// The Unix timestamp in seconds when the model was created.
	Created int64 `json:"created"`
	// Any additional metadata for this model
	Metadata map[string]any `json:"metadata"`
	// Enumeration of supported model types in OGX.
	//
	// Any of "llm", "embedding", "rerank".
	ModelType string `json:"model_type"`
	// Enable model availability check during registration. When false (default),
	// validation is deferred to runtime and model is preserved during provider
	// refresh.
	ModelValidation bool `json:"model_validation" api:"nullable"`
	// The owner of the model.
	OwnedBy string `json:"owned_by"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id" api:"nullable"`
	// Any of "model".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Identifier         respjson.Field
		Object             respjson.Field
		ProviderID         respjson.Field
		Created            respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
		ModelValidation    respjson.Field
		OwnedBy            respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelGetResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ModelGetResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Anthropic model info response object.
//
// :id: Unique model identifier
type ModelGetResponseAnthropicModelInfo struct {
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
func (r ModelGetResponseAnthropicModelInfo) RawJSON() string { return r.JSON.raw }
func (r *ModelGetResponseAnthropicModelInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Google model info response object.
//
// :name: Model resource name, e.g. 'models/gemini-pro' :display_name: A
// human-readable name for the model :description: A description of the model
type ModelGetResponseGoogleModelInfo struct {
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
func (r ModelGetResponseGoogleModelInfo) RawJSON() string { return r.JSON.raw }
func (r *ModelGetResponseGoogleModelInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelListResponseUnion contains all possible properties and values from
// [ListModelsResponse], [ModelListResponseAnthropicListModelsResponse],
// [ModelListResponseGoogleListModelsResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ModelListResponseUnion struct {
	// This field is a union of [[]Model],
	// [[]ModelListResponseAnthropicListModelsResponseData]
	Data ModelListResponseUnionData `json:"data"`
	// This field is from variant [ListModelsResponse].
	Object ListModelsResponseObject `json:"object"`
	// This field is from variant [ModelListResponseAnthropicListModelsResponse].
	FirstID string `json:"first_id"`
	// This field is from variant [ModelListResponseAnthropicListModelsResponse].
	HasMore bool `json:"has_more"`
	// This field is from variant [ModelListResponseAnthropicListModelsResponse].
	LastID string `json:"last_id"`
	// This field is from variant [ModelListResponseGoogleListModelsResponse].
	Models []ModelListResponseGoogleListModelsResponseModel `json:"models"`
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

func (u ModelListResponseUnion) AsOpenAIListModelsResponse() (v ListModelsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelListResponseUnion) AsAnthropicListModelsResponse() (v ModelListResponseAnthropicListModelsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelListResponseUnion) AsGoogleListModelsResponse() (v ModelListResponseGoogleListModelsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelListResponseUnionData is an implicit subunion of [ModelListResponseUnion].
// ModelListResponseUnionData provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ModelListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfData]
type ModelListResponseUnionData struct {
	// This field will be present if the value is a [[]Model] instead of an object.
	OfData []Model `json:",inline"`
	JSON   struct {
		OfData respjson.Field
		raw    string
	} `json:"-"`
}

func (r *ModelListResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of Anthropic model objects.
type ModelListResponseAnthropicListModelsResponse struct {
	// List of Anthropic model objects.
	Data []ModelListResponseAnthropicListModelsResponseData `json:"data" api:"required"`
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
func (r ModelListResponseAnthropicListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseAnthropicListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Anthropic model info response object.
//
// :id: Unique model identifier
type ModelListResponseAnthropicListModelsResponseData struct {
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
func (r ModelListResponseAnthropicListModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseAnthropicListModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of Google model objects.
type ModelListResponseGoogleListModelsResponse struct {
	// List of Google model objects.
	Models []ModelListResponseGoogleListModelsResponseModel `json:"models" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Models      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseGoogleListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseGoogleListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Google model info response object.
//
// :name: Model resource name, e.g. 'models/gemini-pro' :display_name: A
// human-readable name for the model :description: A description of the model
type ModelListResponseGoogleListModelsResponseModel struct {
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
func (r ModelListResponseGoogleListModelsResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseGoogleListModelsResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelGetParams struct {
	AnthropicVersion param.Opt[string] `header:"anthropic-version,omitzero" json:"-"`
	XGoogAPIClient   param.Opt[string] `header:"x-goog-api-client,omitzero" json:"-"`
	XGoogAPIKey      param.Opt[string] `header:"x-goog-api-key,omitzero" json:"-"`
	XGoogUserProject param.Opt[string] `header:"x-goog-user-project,omitzero" json:"-"`
	paramObj
}

type ModelListParams struct {
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

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
