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

// EvalService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalService] method instead.
type EvalService struct {
	Options []option.RequestOption
	Jobs    EvalJobService
}

// NewEvalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalService(opts ...option.RequestOption) (r EvalService) {
	r = EvalService{}
	r.Options = opts
	r.Jobs = NewEvalJobService(opts...)
	return
}

// Evaluate a list of rows on a benchmark.
func (r *EvalService) EvaluateRows(ctx context.Context, benchmarkID string, body EvalEvaluateRowsParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/evaluations", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Evaluate a list of rows on a benchmark.
func (r *EvalService) EvaluateRowsAlpha(ctx context.Context, benchmarkID string, body EvalEvaluateRowsAlphaParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/evaluations", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run an evaluation on a benchmark.
func (r *EvalService) RunEval(ctx context.Context, benchmarkID string, body EvalRunEvalParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run an evaluation on a benchmark.
func (r *EvalService) RunEvalAlpha(ctx context.Context, benchmarkID string, body EvalRunEvalAlphaParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs", benchmarkID)
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

type EvalEvaluateRowsParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	// The rows to evaluate.
	InputRows []map[string]EvalEvaluateRowsParamsInputRowUnion `json:"input_rows,omitzero,required"`
	// The scoring functions to use for the evaluation.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r EvalEvaluateRowsParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalEvaluateRowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalEvaluateRowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalEvaluateRowsParamsInputRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u EvalEvaluateRowsParamsInputRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *EvalEvaluateRowsParamsInputRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalEvaluateRowsParamsInputRowUnion) asAny() any {
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

type EvalEvaluateRowsAlphaParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	// The rows to evaluate.
	InputRows []map[string]EvalEvaluateRowsAlphaParamsInputRowUnion `json:"input_rows,omitzero,required"`
	// The scoring functions to use for the evaluation.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r EvalEvaluateRowsAlphaParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalEvaluateRowsAlphaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalEvaluateRowsAlphaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalEvaluateRowsAlphaParamsInputRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u EvalEvaluateRowsAlphaParamsInputRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *EvalEvaluateRowsAlphaParamsInputRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalEvaluateRowsAlphaParamsInputRowUnion) asAny() any {
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

type EvalRunEvalParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	paramObj
}

func (r EvalRunEvalParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalRunEvalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalRunEvalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalRunEvalAlphaParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	paramObj
}

func (r EvalRunEvalAlphaParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalRunEvalAlphaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalRunEvalAlphaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
