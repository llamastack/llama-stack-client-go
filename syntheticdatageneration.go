// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared"
)

// SyntheticDataGenerationService contains methods and other services that help
// with interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSyntheticDataGenerationService] method instead.
type SyntheticDataGenerationService struct {
	Options []option.RequestOption
}

// NewSyntheticDataGenerationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSyntheticDataGenerationService(opts ...option.RequestOption) (r SyntheticDataGenerationService) {
	r = SyntheticDataGenerationService{}
	r.Options = opts
	return
}

// Generate synthetic data based on input dialogs and apply filtering.
func (r *SyntheticDataGenerationService) Generate(ctx context.Context, body SyntheticDataGenerationGenerateParams, opts ...option.RequestOption) (res *SyntheticDataGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/synthetic-data-generation/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from the synthetic data generation. Batch of (prompt, response, score)
// tuples that pass the threshold.
type SyntheticDataGenerationResponse struct {
	// List of generated synthetic data samples that passed the filtering criteria
	SyntheticData []map[string]SyntheticDataGenerationResponseSyntheticDataUnion `json:"synthetic_data,required"`
	// (Optional) Statistical information about the generation process and filtering
	// results
	Statistics map[string]SyntheticDataGenerationResponseStatisticUnion `json:"statistics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SyntheticData respjson.Field
		Statistics    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyntheticDataGenerationResponse) RawJSON() string { return r.JSON.raw }
func (r *SyntheticDataGenerationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SyntheticDataGenerationResponseSyntheticDataUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SyntheticDataGenerationResponseSyntheticDataUnion struct {
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

func (u SyntheticDataGenerationResponseSyntheticDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SyntheticDataGenerationResponseSyntheticDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SyntheticDataGenerationResponseSyntheticDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SyntheticDataGenerationResponseSyntheticDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SyntheticDataGenerationResponseSyntheticDataUnion) RawJSON() string { return u.JSON.raw }

func (r *SyntheticDataGenerationResponseSyntheticDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SyntheticDataGenerationResponseStatisticUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SyntheticDataGenerationResponseStatisticUnion struct {
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

func (u SyntheticDataGenerationResponseStatisticUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SyntheticDataGenerationResponseStatisticUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SyntheticDataGenerationResponseStatisticUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SyntheticDataGenerationResponseStatisticUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SyntheticDataGenerationResponseStatisticUnion) RawJSON() string { return u.JSON.raw }

func (r *SyntheticDataGenerationResponseStatisticUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SyntheticDataGenerationGenerateParams struct {
	// List of conversation messages to use as input for synthetic data generation
	Dialogs []shared.MessageUnionParam `json:"dialogs,omitzero,required"`
	// Type of filtering to apply to generated synthetic data samples
	//
	// Any of "none", "random", "top_k", "top_p", "top_k_top_p", "sigmoid".
	FilteringFunction SyntheticDataGenerationGenerateParamsFilteringFunction `json:"filtering_function,omitzero,required"`
	// (Optional) The identifier of the model to use. The model must be registered with
	// Llama Stack and available via the /models endpoint
	Model param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r SyntheticDataGenerationGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow SyntheticDataGenerationGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SyntheticDataGenerationGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of filtering to apply to generated synthetic data samples
type SyntheticDataGenerationGenerateParamsFilteringFunction string

const (
	SyntheticDataGenerationGenerateParamsFilteringFunctionNone     SyntheticDataGenerationGenerateParamsFilteringFunction = "none"
	SyntheticDataGenerationGenerateParamsFilteringFunctionRandom   SyntheticDataGenerationGenerateParamsFilteringFunction = "random"
	SyntheticDataGenerationGenerateParamsFilteringFunctionTopK     SyntheticDataGenerationGenerateParamsFilteringFunction = "top_k"
	SyntheticDataGenerationGenerateParamsFilteringFunctionTopP     SyntheticDataGenerationGenerateParamsFilteringFunction = "top_p"
	SyntheticDataGenerationGenerateParamsFilteringFunctionTopKTopP SyntheticDataGenerationGenerateParamsFilteringFunction = "top_k_top_p"
	SyntheticDataGenerationGenerateParamsFilteringFunctionSigmoid  SyntheticDataGenerationGenerateParamsFilteringFunction = "sigmoid"
)
