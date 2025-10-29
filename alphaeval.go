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
// The properties EvalCandidate, ScoringParams are required.
type BenchmarkConfigParam struct {
	// The candidate to evaluate.
	EvalCandidate BenchmarkConfigEvalCandidateUnionParam `json:"eval_candidate,omitzero,required"`
	// Map between scoring function id and parameters for each scoring function you
	// want to run
	ScoringParams map[string]ScoringFnParamsUnion `json:"scoring_params,omitzero,required"`
	// (Optional) The number of examples to evaluate. If not provided, all examples in
	// the dataset will be evaluated
	NumExamples param.Opt[int64] `json:"num_examples,omitzero"`
	paramObj
}

func (r BenchmarkConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BenchmarkConfigEvalCandidateUnionParam struct {
	OfModel *BenchmarkConfigEvalCandidateModelParam `json:",omitzero,inline"`
	OfAgent *BenchmarkConfigEvalCandidateAgentParam `json:",omitzero,inline"`
	paramUnion
}

func (u BenchmarkConfigEvalCandidateUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModel, u.OfAgent)
}
func (u *BenchmarkConfigEvalCandidateUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BenchmarkConfigEvalCandidateUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModel) {
		return u.OfModel
	} else if !param.IsOmitted(u.OfAgent) {
		return u.OfAgent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetModel() *string {
	if vt := u.OfModel; vt != nil {
		return &vt.Model
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetSamplingParams() *SamplingParams {
	if vt := u.OfModel; vt != nil {
		return &vt.SamplingParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetSystemMessage() *SystemMessageParam {
	if vt := u.OfModel; vt != nil {
		return &vt.SystemMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetConfig() *AgentConfigParam {
	if vt := u.OfAgent; vt != nil {
		return &vt.Config
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetType() *string {
	if vt := u.OfModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAgent; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BenchmarkConfigEvalCandidateUnionParam](
		"type",
		apijson.Discriminator[BenchmarkConfigEvalCandidateModelParam]("model"),
		apijson.Discriminator[BenchmarkConfigEvalCandidateAgentParam]("agent"),
	)
}

// A model candidate for evaluation.
//
// The properties Model, SamplingParams, Type are required.
type BenchmarkConfigEvalCandidateModelParam struct {
	// The model ID to evaluate.
	Model string `json:"model,required"`
	// The sampling parameters for the model.
	SamplingParams SamplingParams `json:"sampling_params,omitzero,required"`
	// (Optional) The system message providing instructions or context to the model.
	SystemMessage SystemMessageParam `json:"system_message,omitzero"`
	// This field can be elided, and will marshal its zero value as "model".
	Type constant.Model `json:"type,required"`
	paramObj
}

func (r BenchmarkConfigEvalCandidateModelParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigEvalCandidateModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigEvalCandidateModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An agent candidate for evaluation.
//
// The properties Config, Type are required.
type BenchmarkConfigEvalCandidateAgentParam struct {
	// The configuration for the agent candidate.
	Config AgentConfigParam `json:"config,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "agent".
	Type constant.Agent `json:"type,required"`
	paramObj
}

func (r BenchmarkConfigEvalCandidateAgentParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigEvalCandidateAgentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigEvalCandidateAgentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from an evaluation.
type EvaluateResponse struct {
	// The generations from the evaluation.
	Generations []map[string]EvaluateResponseGenerationUnion `json:"generations,required"`
	// The scores from the evaluation.
	Scores map[string]ScoringResult `json:"scores,required"`
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

// EvaluateResponseGenerationUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type EvaluateResponseGenerationUnion struct {
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

func (u EvaluateResponseGenerationUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseGenerationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseGenerationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseGenerationUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvaluateResponseGenerationUnion) RawJSON() string { return u.JSON.raw }

func (r *EvaluateResponseGenerationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A job execution instance with status tracking.
type Job struct {
	// Unique identifier for the job
	JobID string `json:"job_id,required"`
	// Current execution status of the job
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

// Current execution status of the job
type JobStatus string

const (
	JobStatusCompleted  JobStatus = "completed"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusFailed     JobStatus = "failed"
	JobStatusScheduled  JobStatus = "scheduled"
	JobStatusCancelled  JobStatus = "cancelled"
)

type AlphaEvalEvaluateRowsParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	// The rows to evaluate.
	InputRows []map[string]AlphaEvalEvaluateRowsParamsInputRowUnion `json:"input_rows,omitzero,required"`
	// The scoring functions to use for the evaluation.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r AlphaEvalEvaluateRowsParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaEvalEvaluateRowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaEvalEvaluateRowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaEvalEvaluateRowsParamsInputRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaEvalEvaluateRowsParamsInputRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AlphaEvalEvaluateRowsParamsInputRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaEvalEvaluateRowsParamsInputRowUnion) asAny() any {
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

type AlphaEvalEvaluateRowsAlphaParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	// The rows to evaluate.
	InputRows []map[string]AlphaEvalEvaluateRowsAlphaParamsInputRowUnion `json:"input_rows,omitzero,required"`
	// The scoring functions to use for the evaluation.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r AlphaEvalEvaluateRowsAlphaParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaEvalEvaluateRowsAlphaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaEvalEvaluateRowsAlphaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaEvalEvaluateRowsAlphaParamsInputRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaEvalEvaluateRowsAlphaParamsInputRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AlphaEvalEvaluateRowsAlphaParamsInputRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaEvalEvaluateRowsAlphaParamsInputRowUnion) asAny() any {
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

type AlphaEvalRunEvalParams struct {
	// The configuration for the benchmark.
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
	// The configuration for the benchmark.
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
