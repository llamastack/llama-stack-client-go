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
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ogx-ai/ogx-client-go/internal/apijson"
	"github.com/ogx-ai/ogx-client-go/internal/requestconfig"
	"github.com/ogx-ai/ogx-client-go/option"
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

// Get a model by its identifier.
func (r *ModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *ModelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List models using the OpenAI API.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

// A model resource representing an AI model registered in OGX.
type ModelGetResponse struct {
	// The model identifier (OpenAI-compatible alias for identifier).
	ID string `json:"id" api:"required"`
	// Unique identifier for this resource in ogx
	Identifier string `json:"identifier" api:"required"`
	// The object type, always 'model'.
	//
	// Any of "model".
	Object ModelGetResponseObject `json:"object" api:"required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id" api:"required"`
	// The Unix timestamp in seconds when the model was created.
	Created int64 `json:"created"`
	// Any additional metadata for this model
	Metadata map[string]any `json:"metadata"`
	// Enumeration of supported model types in OGX.
	//
	// Any of "llm", "embedding", "rerank".
	ModelType ModelGetResponseModelType `json:"model_type"`
	// Enable model availability check during registration. When false (default),
	// validation is deferred to runtime and model is preserved during provider
	// refresh.
	ModelValidation bool `json:"model_validation" api:"nullable"`
	// The owner of the model.
	OwnedBy string `json:"owned_by"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id" api:"nullable"`
	// Any of "model".
	Type ModelGetResponseType `json:"type"`
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
func (r ModelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type, always 'model'.
type ModelGetResponseObject string

const (
	ModelGetResponseObjectModel ModelGetResponseObject = "model"
)

// Enumeration of supported model types in OGX.
type ModelGetResponseModelType string

const (
	ModelGetResponseModelTypeLlm       ModelGetResponseModelType = "llm"
	ModelGetResponseModelTypeEmbedding ModelGetResponseModelType = "embedding"
	ModelGetResponseModelTypeRerank    ModelGetResponseModelType = "rerank"
)

type ModelGetResponseType string

const (
	ModelGetResponseTypeModel ModelGetResponseType = "model"
)
