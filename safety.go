// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
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
func (r *SafetyService) OpenAIModerations(ctx context.Context, body SafetyOpenAIModerationsParams, opts ...option.RequestOption) (res *OpenAIModerationsResponse, err error) {
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
type OpenAIModerationsResponse struct {
	// The unique identifier for the moderation request.
	ID string `json:"id,required"`
	// The model used to generate the moderation results.
	Model string `json:"model,required"`
	// A list of moderation objects
	Results []OpenAIModerationsResponseResult `json:"results,required"`
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
func (r OpenAIModerationsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAIModerationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A moderation object.
type OpenAIModerationsResponseResult struct {
	// A list of the categories, and whether they are flagged or not.
	Categories                map[string]bool     `json:"categories,required"`
	CategoryAppliedInputTypes map[string][]string `json:"category_applied_input_types,required"`
	CategoryMessages          map[string]string   `json:"category_messages,required"`
	// A list of the categories along with their scores as predicted by model.
	CategoryScores map[string]float64 `json:"category_scores,required"`
	// Whether any of the below categories are flagged.
	Flagged bool `json:"flagged,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories                respjson.Field
		CategoryAppliedInputTypes respjson.Field
		CategoryMessages          respjson.Field
		CategoryScores            respjson.Field
		Flagged                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIModerationsResponseResult) RawJSON() string { return r.JSON.raw }
func (r *OpenAIModerationsResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunShieldResponse struct {
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

type SafetyOpenAIModerationsParams struct {
	// Input (or inputs) to classify. Can be a single string, an array of strings, or
	// an array of multi-modal input objects similar to other models.
	Input SafetyOpenAIModerationsParamsInputUnion `json:"input,omitzero,required"`
	// The content moderation model you would like to use.
	Model param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r SafetyOpenAIModerationsParams) MarshalJSON() (data []byte, err error) {
	type shadow SafetyOpenAIModerationsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyOpenAIModerationsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyOpenAIModerationsParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyOpenAIModerationsParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *SafetyOpenAIModerationsParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyOpenAIModerationsParamsInputUnion) asAny() any {
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
