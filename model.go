// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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

// Get model. Get a model by its identifier.
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

// Register model. Register a model.
func (r *ModelService) Register(ctx context.Context, body ModelRegisterParams, opts ...option.RequestOption) (res *ModelRegisterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unregister model. Unregister a model.
func (r *ModelService) Unregister(ctx context.Context, modelID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
type Model struct {
	ID             string                              `json:"id,required"`
	Created        int64                               `json:"created,required"`
	Object         constant.Model                      `json:"object,required"`
	OwnedBy        string                              `json:"owned_by,required"`
	CustomMetadata map[string]ModelCustomMetadataUnion `json:"custom_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Created        respjson.Field
		Object         respjson.Field
		OwnedBy        respjson.Field
		CustomMetadata respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelCustomMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ModelCustomMetadataUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ModelCustomMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelCustomMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelCustomMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelCustomMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelCustomMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model resource representing an AI model registered in Llama Stack.
type ModelGetResponse struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// Any additional metadata for this model
	Metadata map[string]ModelGetResponseMetadataUnion `json:"metadata,required"`
	// The type of model (LLM or embedding model)
	//
	// Any of "llm", "embedding", "rerank".
	ModelType ModelGetResponseModelType `json:"model_type,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// The resource type, always 'model' for model resources
	Type constant.Model `json:"type,required"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
		ProviderID         respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelGetResponseMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ModelGetResponseMetadataUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ModelGetResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelGetResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelGetResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelGetResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelGetResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelGetResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of model (LLM or embedding model)
type ModelGetResponseModelType string

const (
	ModelGetResponseModelTypeLlm       ModelGetResponseModelType = "llm"
	ModelGetResponseModelTypeEmbedding ModelGetResponseModelType = "embedding"
	ModelGetResponseModelTypeRerank    ModelGetResponseModelType = "rerank"
)

// A model resource representing an AI model registered in Llama Stack.
type ModelRegisterResponse struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// Any additional metadata for this model
	Metadata map[string]ModelRegisterResponseMetadataUnion `json:"metadata,required"`
	// The type of model (LLM or embedding model)
	//
	// Any of "llm", "embedding", "rerank".
	ModelType ModelRegisterResponseModelType `json:"model_type,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// The resource type, always 'model' for model resources
	Type constant.Model `json:"type,required"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
		ProviderID         respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelRegisterResponseMetadataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ModelRegisterResponseMetadataUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ModelRegisterResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelRegisterResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelRegisterResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelRegisterResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelRegisterResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelRegisterResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of model (LLM or embedding model)
type ModelRegisterResponseModelType string

const (
	ModelRegisterResponseModelTypeLlm       ModelRegisterResponseModelType = "llm"
	ModelRegisterResponseModelTypeEmbedding ModelRegisterResponseModelType = "embedding"
	ModelRegisterResponseModelTypeRerank    ModelRegisterResponseModelType = "rerank"
)

type ModelRegisterParams struct {
	// The identifier of the model to register.
	ModelID string `json:"model_id,required"`
	// The identifier of the provider.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The identifier of the model in the provider.
	ProviderModelID param.Opt[string] `json:"provider_model_id,omitzero"`
	// Any additional metadata for this model.
	Metadata map[string]ModelRegisterParamsMetadataUnion `json:"metadata,omitzero"`
	// The type of model to register.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelRegisterParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ModelRegisterParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ModelRegisterParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelRegisterParamsMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// The type of model to register.
type ModelRegisterParamsModelType string

const (
	ModelRegisterParamsModelTypeLlm       ModelRegisterParamsModelType = "llm"
	ModelRegisterParamsModelTypeEmbedding ModelRegisterParamsModelType = "embedding"
	ModelRegisterParamsModelTypeRerank    ModelRegisterParamsModelType = "rerank"
)
