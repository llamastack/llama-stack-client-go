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
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ScoringService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoringService] method instead.
type ScoringService struct {
	Options []option.RequestOption
}

// NewScoringService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScoringService(opts ...option.RequestOption) (r ScoringService) {
	r = ScoringService{}
	r.Options = opts
	return
}

// Score a list of rows.
func (r *ScoringService) Score(ctx context.Context, body ScoringScoreParams, opts ...option.RequestOption) (res *ScoringScoreResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/scoring/score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Score a batch of rows.
func (r *ScoringService) ScoreBatch(ctx context.Context, body ScoringScoreBatchParams, opts ...option.RequestOption) (res *ScoringScoreBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/scoring/score-batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The response from scoring.
type ScoringScoreResponse struct {
	Results map[string]ScoringResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringScoreResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoringScoreResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from batch scoring operations on datasets.
type ScoringScoreBatchResponse struct {
	Results   map[string]ScoringResult `json:"results,required"`
	DatasetID string                   `json:"dataset_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		DatasetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringScoreBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoringScoreBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringScoreParams struct {
	InputRows        []map[string]any                                  `json:"input_rows,omitzero,required"`
	ScoringFunctions map[string]ScoringScoreParamsScoringFunctionUnion `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r ScoringScoreParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScoringScoreParamsScoringFunctionUnion struct {
	OfLlmAsJudge  *ScoringScoreParamsScoringFunctionLlmAsJudge  `json:",omitzero,inline"`
	OfRegexParser *ScoringScoreParamsScoringFunctionRegexParser `json:",omitzero,inline"`
	OfBasic       *ScoringScoreParamsScoringFunctionBasic       `json:",omitzero,inline"`
	paramUnion
}

func (u ScoringScoreParamsScoringFunctionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlmAsJudge, u.OfRegexParser, u.OfBasic)
}
func (u *ScoringScoreParamsScoringFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScoringScoreParamsScoringFunctionUnion) asAny() any {
	if !param.IsOmitted(u.OfLlmAsJudge) {
		return u.OfLlmAsJudge
	} else if !param.IsOmitted(u.OfRegexParser) {
		return u.OfRegexParser
	} else if !param.IsOmitted(u.OfBasic) {
		return u.OfBasic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreParamsScoringFunctionUnion) GetJudgeModel() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return &vt.JudgeModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreParamsScoringFunctionUnion) GetJudgeScoreRegexes() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.JudgeScoreRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreParamsScoringFunctionUnion) GetPromptTemplate() *string {
	if vt := u.OfLlmAsJudge; vt != nil && vt.PromptTemplate.Valid() {
		return &vt.PromptTemplate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreParamsScoringFunctionUnion) GetParsingRegexes() []string {
	if vt := u.OfRegexParser; vt != nil {
		return vt.ParsingRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreParamsScoringFunctionUnion) GetType() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRegexParser; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBasic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's AggregationFunctions property, if
// present.
func (u ScoringScoreParamsScoringFunctionUnion) GetAggregationFunctions() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfRegexParser; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfBasic; vt != nil {
		return vt.AggregationFunctions
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ScoringScoreParamsScoringFunctionUnion](
		"type",
		apijson.Discriminator[ScoringScoreParamsScoringFunctionLlmAsJudge]("llm_as_judge"),
		apijson.Discriminator[ScoringScoreParamsScoringFunctionRegexParser]("regex_parser"),
		apijson.Discriminator[ScoringScoreParamsScoringFunctionBasic]("basic"),
	)
}

// Parameters for LLM-as-judge scoring function configuration.
//
// The property JudgeModel is required.
type ScoringScoreParamsScoringFunctionLlmAsJudge struct {
	JudgeModel     string            `json:"judge_model,required"`
	PromptTemplate param.Opt[string] `json:"prompt_template,omitzero"`
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes,omitzero"`
	// Any of "llm_as_judge".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ScoringScoreParamsScoringFunctionLlmAsJudge) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreParamsScoringFunctionLlmAsJudge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreParamsScoringFunctionLlmAsJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringScoreParamsScoringFunctionLlmAsJudge](
		"type", "llm_as_judge",
	)
}

// Parameters for regex parser scoring function configuration.
type ScoringScoreParamsScoringFunctionRegexParser struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes,omitzero"`
	// Any of "regex_parser".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ScoringScoreParamsScoringFunctionRegexParser) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreParamsScoringFunctionRegexParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreParamsScoringFunctionRegexParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringScoreParamsScoringFunctionRegexParser](
		"type", "regex_parser",
	)
}

// Parameters for basic scoring function configuration.
type ScoringScoreParamsScoringFunctionBasic struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Any of "basic".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ScoringScoreParamsScoringFunctionBasic) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreParamsScoringFunctionBasic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreParamsScoringFunctionBasic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringScoreParamsScoringFunctionBasic](
		"type", "basic",
	)
}

type ScoringScoreBatchParams struct {
	DatasetID          string                                                 `json:"dataset_id,required"`
	ScoringFunctions   map[string]ScoringScoreBatchParamsScoringFunctionUnion `json:"scoring_functions,omitzero,required"`
	SaveResultsDataset param.Opt[bool]                                        `json:"save_results_dataset,omitzero"`
	paramObj
}

func (r ScoringScoreBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScoringScoreBatchParamsScoringFunctionUnion struct {
	OfLlmAsJudge  *ScoringScoreBatchParamsScoringFunctionLlmAsJudge  `json:",omitzero,inline"`
	OfRegexParser *ScoringScoreBatchParamsScoringFunctionRegexParser `json:",omitzero,inline"`
	OfBasic       *ScoringScoreBatchParamsScoringFunctionBasic       `json:",omitzero,inline"`
	paramUnion
}

func (u ScoringScoreBatchParamsScoringFunctionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlmAsJudge, u.OfRegexParser, u.OfBasic)
}
func (u *ScoringScoreBatchParamsScoringFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScoringScoreBatchParamsScoringFunctionUnion) asAny() any {
	if !param.IsOmitted(u.OfLlmAsJudge) {
		return u.OfLlmAsJudge
	} else if !param.IsOmitted(u.OfRegexParser) {
		return u.OfRegexParser
	} else if !param.IsOmitted(u.OfBasic) {
		return u.OfBasic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreBatchParamsScoringFunctionUnion) GetJudgeModel() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return &vt.JudgeModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreBatchParamsScoringFunctionUnion) GetJudgeScoreRegexes() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.JudgeScoreRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreBatchParamsScoringFunctionUnion) GetPromptTemplate() *string {
	if vt := u.OfLlmAsJudge; vt != nil && vt.PromptTemplate.Valid() {
		return &vt.PromptTemplate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreBatchParamsScoringFunctionUnion) GetParsingRegexes() []string {
	if vt := u.OfRegexParser; vt != nil {
		return vt.ParsingRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringScoreBatchParamsScoringFunctionUnion) GetType() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRegexParser; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBasic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's AggregationFunctions property, if
// present.
func (u ScoringScoreBatchParamsScoringFunctionUnion) GetAggregationFunctions() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfRegexParser; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfBasic; vt != nil {
		return vt.AggregationFunctions
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ScoringScoreBatchParamsScoringFunctionUnion](
		"type",
		apijson.Discriminator[ScoringScoreBatchParamsScoringFunctionLlmAsJudge]("llm_as_judge"),
		apijson.Discriminator[ScoringScoreBatchParamsScoringFunctionRegexParser]("regex_parser"),
		apijson.Discriminator[ScoringScoreBatchParamsScoringFunctionBasic]("basic"),
	)
}

// Parameters for LLM-as-judge scoring function configuration.
//
// The property JudgeModel is required.
type ScoringScoreBatchParamsScoringFunctionLlmAsJudge struct {
	JudgeModel     string            `json:"judge_model,required"`
	PromptTemplate param.Opt[string] `json:"prompt_template,omitzero"`
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes,omitzero"`
	// Any of "llm_as_judge".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ScoringScoreBatchParamsScoringFunctionLlmAsJudge) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreBatchParamsScoringFunctionLlmAsJudge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreBatchParamsScoringFunctionLlmAsJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringScoreBatchParamsScoringFunctionLlmAsJudge](
		"type", "llm_as_judge",
	)
}

// Parameters for regex parser scoring function configuration.
type ScoringScoreBatchParamsScoringFunctionRegexParser struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes,omitzero"`
	// Any of "regex_parser".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ScoringScoreBatchParamsScoringFunctionRegexParser) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreBatchParamsScoringFunctionRegexParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreBatchParamsScoringFunctionRegexParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringScoreBatchParamsScoringFunctionRegexParser](
		"type", "regex_parser",
	)
}

// Parameters for basic scoring function configuration.
type ScoringScoreBatchParamsScoringFunctionBasic struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Any of "basic".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ScoringScoreBatchParamsScoringFunctionBasic) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreBatchParamsScoringFunctionBasic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreBatchParamsScoringFunctionBasic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringScoreBatchParamsScoringFunctionBasic](
		"type", "basic",
	)
}
