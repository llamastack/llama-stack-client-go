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

// SafetyService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSafetyService] method instead.
type SafetyService struct {
	Options []option.RequestOption
}

// NewSafetyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSafetyService(opts ...option.RequestOption) (r SafetyService) {
	r = SafetyService{}
	r.Options = opts
	return
}

// Classifies if text and/or image inputs are potentially harmful.
func (r *SafetyService) New(ctx context.Context, body SafetyNewParams, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/moderations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run a shield.
func (r *SafetyService) RunShield(ctx context.Context, body SafetyRunShieldParams, opts ...option.RequestOption) (res *RunShieldResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/safety/run-shield"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A moderation object.
type CreateResponse struct {
	// The unique identifier for the moderation request.
	ID string `json:"id,required"`
	// The model used to generate the moderation results.
	Model string `json:"model,required"`
	// A list of moderation objects
	Results []CreateResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Model       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A moderation object.
type CreateResponseResult struct {
	// Whether any of the below categories are flagged.
	Flagged  bool                                         `json:"flagged,required"`
	Metadata map[string]CreateResponseResultMetadataUnion `json:"metadata,required"`
	// A list of the categories, and whether they are flagged or not.
	Categories map[string]bool `json:"categories"`
	// A list of the categories along with the input type(s) that the score applies to.
	CategoryAppliedInputTypes map[string][]string `json:"category_applied_input_types"`
	// A list of the categories along with their scores as predicted by model. Required
	// set of categories that need to be in response - violence - violence/graphic -
	// harassment - harassment/threatening - hate - hate/threatening - illicit -
	// illicit/violent - sexual - sexual/minors - self-harm - self-harm/intent -
	// self-harm/instructions
	CategoryScores map[string]float64 `json:"category_scores"`
	UserMessage    string             `json:"user_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flagged                   respjson.Field
		Metadata                  respjson.Field
		Categories                respjson.Field
		CategoryAppliedInputTypes respjson.Field
		CategoryScores            respjson.Field
		UserMessage               respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateResponseResult) RawJSON() string { return r.JSON.raw }
func (r *CreateResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateResponseResultMetadataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type CreateResponseResultMetadataUnion struct {
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

func (u CreateResponseResultMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreateResponseResultMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreateResponseResultMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreateResponseResultMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CreateResponseResultMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *CreateResponseResultMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from running a safety shield.
type RunShieldResponse struct {
	// (Optional) Safety violation detected by the shield, if any
	Violation shared.SafetyViolation `json:"violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Violation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunShieldResponse) RawJSON() string { return r.JSON.raw }
func (r *RunShieldResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SafetyNewParams struct {
	// Input (or inputs) to classify. Can be a single string, an array of strings, or
	// an array of multi-modal input objects similar to other models.
	Input SafetyNewParamsInputUnion `json:"input,omitzero,required"`
	// The content moderation model you would like to use.
	Model string `json:"model,required"`
	paramObj
}

func (r SafetyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SafetyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyNewParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *SafetyNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type SafetyRunShieldParams struct {
	// The messages to run the shield on.
	Messages []shared.MessageUnionParam `json:"messages,omitzero,required"`
	// The parameters of the shield.
	Params map[string]SafetyRunShieldParamsParamUnion `json:"params,omitzero,required"`
	// The identifier of the shield to run.
	ShieldID string `json:"shield_id,required"`
	paramObj
}

func (r SafetyRunShieldParams) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsParamUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *SafetyRunShieldParamsParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsParamUnion) asAny() any {
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
