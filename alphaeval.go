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

// AlphaEvalService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaEvalService] method instead.
type AlphaEvalService struct {
	Options []option.RequestOption
	Jobs    AlphaEvalJobService
}

// NewAlphaEvalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaEvalService(opts ...option.RequestOption) (r AlphaEvalService) {
	r = AlphaEvalService{}
	r.Options = opts
	r.Jobs = NewAlphaEvalJobService(opts...)
	return
}

// Evaluate a list of rows on a benchmark.
func (r *AlphaEvalService) EvaluateRows(ctx context.Context, benchmarkID string, body AlphaEvalEvaluateRowsParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/evaluations", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Evaluate a list of rows on a benchmark.
func (r *AlphaEvalService) EvaluateRowsAlpha(ctx context.Context, benchmarkID string, body AlphaEvalEvaluateRowsAlphaParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/evaluations", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run an evaluation on a benchmark.
func (r *AlphaEvalService) RunEval(ctx context.Context, benchmarkID string, body AlphaEvalRunEvalParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/jobs", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run an evaluation on a benchmark.
func (r *AlphaEvalService) RunEvalAlpha(ctx context.Context, benchmarkID string, body AlphaEvalRunEvalAlphaParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/jobs", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A benchmark configuration for evaluation.
//
// The property EvalCandidate is required.
type BenchmarkConfigParam struct {
	// A model candidate for evaluation.
	EvalCandidate BenchmarkConfigEvalCandidateParam `json:"eval_candidate,omitzero,required"`
	// Number of examples to evaluate (useful for testing), if not provided, all
	// examples in the dataset will be evaluated
	NumExamples param.Opt[int64] `json:"num_examples,omitzero"`
	// Map between scoring function id and parameters for each scoring function you
	// want to run
	ScoringParams map[string]BenchmarkConfigScoringParamUnionParam `json:"scoring_params,omitzero"`
	paramObj
}

func (r BenchmarkConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model candidate for evaluation.
//
// The properties Model, SamplingParams are required.
type BenchmarkConfigEvalCandidateParam struct {
	Model string `json:"model,required"`
	// Sampling parameters.
	SamplingParams SamplingParams `json:"sampling_params,omitzero,required"`
	// A system message providing instructions or context to the model.
	SystemMessage SystemMessageParam `json:"system_message,omitzero"`
	// Any of "model".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BenchmarkConfigEvalCandidateParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigEvalCandidateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigEvalCandidateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BenchmarkConfigEvalCandidateParam](
		"type", "model",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BenchmarkConfigScoringParamUnionParam struct {
	OfLlmAsJudge  *BenchmarkConfigScoringParamLlmAsJudgeParam  `json:",omitzero,inline"`
	OfRegexParser *BenchmarkConfigScoringParamRegexParserParam `json:",omitzero,inline"`
	OfBasic       *BenchmarkConfigScoringParamBasicParam       `json:",omitzero,inline"`
	paramUnion
}

func (u BenchmarkConfigScoringParamUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlmAsJudge, u.OfRegexParser, u.OfBasic)
}
func (u *BenchmarkConfigScoringParamUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BenchmarkConfigScoringParamUnionParam) asAny() any {
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
func (u BenchmarkConfigScoringParamUnionParam) GetJudgeModel() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return &vt.JudgeModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigScoringParamUnionParam) GetJudgeScoreRegexes() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.JudgeScoreRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigScoringParamUnionParam) GetPromptTemplate() *string {
	if vt := u.OfLlmAsJudge; vt != nil && vt.PromptTemplate.Valid() {
		return &vt.PromptTemplate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigScoringParamUnionParam) GetParsingRegexes() []string {
	if vt := u.OfRegexParser; vt != nil {
		return vt.ParsingRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigScoringParamUnionParam) GetType() *string {
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
func (u BenchmarkConfigScoringParamUnionParam) GetAggregationFunctions() []string {
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
	apijson.RegisterUnion[BenchmarkConfigScoringParamUnionParam](
		"type",
		apijson.Discriminator[BenchmarkConfigScoringParamLlmAsJudgeParam]("llm_as_judge"),
		apijson.Discriminator[BenchmarkConfigScoringParamRegexParserParam]("regex_parser"),
		apijson.Discriminator[BenchmarkConfigScoringParamBasicParam]("basic"),
	)
}

// Parameters for LLM-as-judge scoring function configuration.
//
// The property JudgeModel is required.
type BenchmarkConfigScoringParamLlmAsJudgeParam struct {
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

func (r BenchmarkConfigScoringParamLlmAsJudgeParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigScoringParamLlmAsJudgeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigScoringParamLlmAsJudgeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BenchmarkConfigScoringParamLlmAsJudgeParam](
		"type", "llm_as_judge",
	)
}

// Parameters for regex parser scoring function configuration.
type BenchmarkConfigScoringParamRegexParserParam struct {
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

func (r BenchmarkConfigScoringParamRegexParserParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigScoringParamRegexParserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigScoringParamRegexParserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BenchmarkConfigScoringParamRegexParserParam](
		"type", "regex_parser",
	)
}

// Parameters for basic scoring function configuration.
type BenchmarkConfigScoringParamBasicParam struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero"`
	// Any of "basic".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BenchmarkConfigScoringParamBasicParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigScoringParamBasicParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigScoringParamBasicParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BenchmarkConfigScoringParamBasicParam](
		"type", "basic",
	)
}

// The response from an evaluation.
type EvaluateResponse struct {
	Generations []map[string]any         `json:"generations,required"`
	Scores      map[string]ScoringResult `json:"scores,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Generations respjson.Field
		Scores      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluateResponse) RawJSON() string { return r.JSON.raw }
func (r *EvaluateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A job execution instance with status tracking.
type Job struct {
	JobID string `json:"job_id,required"`
	// Status of a job execution.
	//
	// Any of "completed", "in_progress", "failed", "scheduled", "cancelled".
	Status JobStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a job execution.
type JobStatus string

const (
	JobStatusCompleted  JobStatus = "completed"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusFailed     JobStatus = "failed"
	JobStatusScheduled  JobStatus = "scheduled"
	JobStatusCancelled  JobStatus = "cancelled"
)

type AlphaEvalEvaluateRowsParams struct {
	// A benchmark configuration for evaluation.
	BenchmarkConfig  BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	InputRows        []map[string]any     `json:"input_rows,omitzero,required"`
	ScoringFunctions []string             `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r AlphaEvalEvaluateRowsParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaEvalEvaluateRowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaEvalEvaluateRowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaEvalEvaluateRowsAlphaParams struct {
	// A benchmark configuration for evaluation.
	BenchmarkConfig  BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	InputRows        []map[string]any     `json:"input_rows,omitzero,required"`
	ScoringFunctions []string             `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r AlphaEvalEvaluateRowsAlphaParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaEvalEvaluateRowsAlphaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaEvalEvaluateRowsAlphaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaEvalRunEvalParams struct {
	// A benchmark configuration for evaluation.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	paramObj
}

func (r AlphaEvalRunEvalParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaEvalRunEvalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaEvalRunEvalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaEvalRunEvalAlphaParams struct {
	// A benchmark configuration for evaluation.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	paramObj
}

func (r AlphaEvalRunEvalAlphaParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaEvalRunEvalAlphaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaEvalRunEvalAlphaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
