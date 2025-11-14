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
)

// ScoringFunctionService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoringFunctionService] method instead.
type ScoringFunctionService struct {
	Options []option.RequestOption
}

// NewScoringFunctionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScoringFunctionService(opts ...option.RequestOption) (r ScoringFunctionService) {
	r = ScoringFunctionService{}
	r.Options = opts
	return
}

// Get a scoring function by its ID.
func (r *ScoringFunctionService) Get(ctx context.Context, scoringFnID string, opts ...option.RequestOption) (res *ScoringFn, err error) {
	opts = slices.Concat(r.Options, opts)
	if scoringFnID == "" {
		err = errors.New("missing required scoring_fn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/scoring-functions/%s", scoringFnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all scoring functions.
func (r *ScoringFunctionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ScoringFn, err error) {
	var env ListScoringFunctionsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/scoring-functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Register a scoring function.
//
// Deprecated: deprecated
func (r *ScoringFunctionService) Register(ctx context.Context, body ScoringFunctionRegisterParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/scoring-functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Unregister a scoring function.
//
// Deprecated: deprecated
func (r *ScoringFunctionService) Unregister(ctx context.Context, scoringFnID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if scoringFnID == "" {
		err = errors.New("missing required scoring_fn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/scoring-functions/%s", scoringFnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListScoringFunctionsResponse struct {
	Data []ScoringFn `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListScoringFunctionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListScoringFunctionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A scoring function resource for evaluating model outputs.
type ScoringFn struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID  string              `json:"provider_id,required"`
	ReturnType  ScoringFnReturnType `json:"return_type,required"`
	Description string              `json:"description,nullable"`
	// Any additional metadata for this definition
	Metadata map[string]any `json:"metadata"`
	// The parameters for the scoring function for benchmark eval, these can be
	// overridden for app eval
	Params ScoringFnParamsUnion `json:"params,nullable"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Any of "scoring_function".
	Type ScoringFnType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		ReturnType         respjson.Field
		Description        respjson.Field
		Metadata           respjson.Field
		Params             respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFn) RawJSON() string { return r.JSON.raw }
func (r *ScoringFn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnReturnType struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnReturnType) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnReturnType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringFnParamsUnion contains all possible properties and values from
// [ScoringFnParamsLlmAsJudge], [ScoringFnParamsRegexParser],
// [ScoringFnParamsBasic].
//
// Use the [ScoringFnParamsUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ScoringFnParamsUnion struct {
	// This field is from variant [ScoringFnParamsLlmAsJudge].
	JudgeModel           string   `json:"judge_model"`
	AggregationFunctions []string `json:"aggregation_functions"`
	// This field is from variant [ScoringFnParamsLlmAsJudge].
	JudgeScoreRegexes []string `json:"judge_score_regexes"`
	// This field is from variant [ScoringFnParamsLlmAsJudge].
	PromptTemplate string `json:"prompt_template"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type string `json:"type"`
	// This field is from variant [ScoringFnParamsRegexParser].
	ParsingRegexes []string `json:"parsing_regexes"`
	JSON           struct {
		JudgeModel           respjson.Field
		AggregationFunctions respjson.Field
		JudgeScoreRegexes    respjson.Field
		PromptTemplate       respjson.Field
		Type                 respjson.Field
		ParsingRegexes       respjson.Field
		raw                  string
	} `json:"-"`
}

// anyScoringFnParams is implemented by each variant of [ScoringFnParamsUnion] to
// add type safety for the return type of [ScoringFnParamsUnion.AsAny]
type anyScoringFnParams interface {
	implScoringFnParamsUnion()
}

func (ScoringFnParamsLlmAsJudge) implScoringFnParamsUnion()  {}
func (ScoringFnParamsRegexParser) implScoringFnParamsUnion() {}
func (ScoringFnParamsBasic) implScoringFnParamsUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ScoringFnParamsUnion.AsAny().(type) {
//	case llamastackclient.ScoringFnParamsLlmAsJudge:
//	case llamastackclient.ScoringFnParamsRegexParser:
//	case llamastackclient.ScoringFnParamsBasic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ScoringFnParamsUnion) AsAny() anyScoringFnParams {
	switch u.Type {
	case "llm_as_judge":
		return u.AsLlmAsJudge()
	case "regex_parser":
		return u.AsRegexParser()
	case "basic":
		return u.AsBasic()
	}
	return nil
}

func (u ScoringFnParamsUnion) AsLlmAsJudge() (v ScoringFnParamsLlmAsJudge) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnParamsUnion) AsRegexParser() (v ScoringFnParamsRegexParser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnParamsUnion) AsBasic() (v ScoringFnParamsBasic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringFnParamsUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringFnParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for LLM-as-judge scoring function configuration.
type ScoringFnParamsLlmAsJudge struct {
	JudgeModel string `json:"judge_model,required"`
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes"`
	PromptTemplate    string   `json:"prompt_template,nullable"`
	// Any of "llm_as_judge".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JudgeModel           respjson.Field
		AggregationFunctions respjson.Field
		JudgeScoreRegexes    respjson.Field
		PromptTemplate       respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsLlmAsJudge) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsLlmAsJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for regex parser scoring function configuration.
type ScoringFnParamsRegexParser struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes"`
	// Any of "regex_parser".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		ParsingRegexes       respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsRegexParser) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsRegexParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for basic scoring function configuration.
type ScoringFnParamsBasic struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions"`
	// Any of "basic".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsBasic) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsBasic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnType string

const (
	ScoringFnTypeScoringFunction ScoringFnType = "scoring_function"
)

type ScoringFunctionRegisterParams struct {
	Description         any `json:"description,omitzero,required"`
	ReturnType          any `json:"return_type,omitzero,required"`
	ScoringFnID         any `json:"scoring_fn_id,omitzero,required"`
	Params              any `json:"params,omitzero"`
	ProviderID          any `json:"provider_id,omitzero"`
	ProviderScoringFnID any `json:"provider_scoring_fn_id,omitzero"`
	paramObj
}

func (r ScoringFunctionRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFunctionRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFunctionRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
