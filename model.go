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
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the llama-stack-client API.
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

// Get model.
//
// Get a model by its identifier.
func (r *ModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *ModelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List models using the OpenAI API.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Model, err error) {
	var env ListModelsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Register model.
//
// Register a model.
//
// Deprecated: deprecated
func (r *ModelService) Register(ctx context.Context, body ModelRegisterParams, opts ...option.RequestOption) (res *ModelRegisterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unregister model.
//
// Unregister a model.
//
// Deprecated: deprecated
func (r *ModelService) Unregister(ctx context.Context, modelID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListModelsResponse struct {
	Data []Model `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model from OpenAI.
//
// :id: The ID of the model :object: The object type, which will be "model"
// :created: The Unix timestamp in seconds when the model was created :owned_by:
// The owner of the model :custom_metadata: Llama Stack-specific metadata including
// model_type, provider info, and additional metadata
type Model struct {
	ID             string         `json:"id,required"`
	Created        int64          `json:"created,required"`
	OwnedBy        string         `json:"owned_by,required"`
	CustomMetadata map[string]any `json:"custom_metadata,nullable"`
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

// A model resource representing an AI model registered in Llama Stack.
type ModelGetResponse struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// Any additional metadata for this model
	Metadata map[string]any `json:"metadata"`
	// Enumeration of supported model types in Llama Stack.
	//
	// Any of "llm", "embedding", "rerank".
	ModelType ModelGetResponseModelType `json:"model_type"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Any of "model".
	Type ModelGetResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
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

// Enumeration of supported model types in Llama Stack.
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

// A model resource representing an AI model registered in Llama Stack.
type ModelRegisterResponse struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// Any additional metadata for this model
	Metadata map[string]any `json:"metadata"`
	// Enumeration of supported model types in Llama Stack.
	//
	// Any of "llm", "embedding", "rerank".
	ModelType ModelRegisterResponseModelType `json:"model_type"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Any of "model".
	Type ModelRegisterResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enumeration of supported model types in Llama Stack.
type ModelRegisterResponseModelType string

const (
	ModelRegisterResponseModelTypeLlm       ModelRegisterResponseModelType = "llm"
	ModelRegisterResponseModelTypeEmbedding ModelRegisterResponseModelType = "embedding"
	ModelRegisterResponseModelTypeRerank    ModelRegisterResponseModelType = "rerank"
)

type ModelRegisterResponseType string

const (
	ModelRegisterResponseTypeModel ModelRegisterResponseType = "model"
)

type ModelRegisterParams struct {
	ModelID         string            `json:"model_id,required"`
	ProviderID      param.Opt[string] `json:"provider_id,omitzero"`
	ProviderModelID param.Opt[string] `json:"provider_model_id,omitzero"`
	Metadata        map[string]any    `json:"metadata,omitzero"`
	// Enumeration of supported model types in Llama Stack.
	//
	// Any of "llm", "embedding", "rerank".
	ModelType ModelRegisterParamsModelType `json:"model_type,omitzero"`
	paramObj
}

func (r ModelRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enumeration of supported model types in Llama Stack.
type ModelRegisterParamsModelType string

const (
	ModelRegisterParamsModelTypeLlm       ModelRegisterParamsModelType = "llm"
	ModelRegisterParamsModelTypeEmbedding ModelRegisterParamsModelType = "embedding"
	ModelRegisterParamsModelTypeRerank    ModelRegisterParamsModelType = "rerank"
)
