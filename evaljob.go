// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
)

// EvalJobService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalJobService] method instead.
type EvalJobService struct {
	Options []option.RequestOption
}

// NewEvalJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalJobService(opts ...option.RequestOption) (r EvalJobService) {
	r = EvalJobService{}
	r.Options = opts
	return
}

// Get the result of a job.
func (r *EvalJobService) Get(ctx context.Context, jobID string, query EvalJobGetParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs/%s/result", query.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a job.
func (r *EvalJobService) Cancel(ctx context.Context, jobID string, body EvalJobCancelParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs/%s", body.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the status of a job.
func (r *EvalJobService) Status(ctx context.Context, jobID string, query EvalJobStatusParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = append(r.Options[:], opts...)
	if query.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs/%s", query.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EvalJobGetParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type EvalJobCancelParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type EvalJobStatusParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}
