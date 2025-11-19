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

	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
)

// AlphaEvalJobService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaEvalJobService] method instead.
type AlphaEvalJobService struct {
	Options []option.RequestOption
}

// NewAlphaEvalJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaEvalJobService(opts ...option.RequestOption) (r AlphaEvalJobService) {
	r = AlphaEvalJobService{}
	r.Options = opts
	return
}

// Get the result of a job.
func (r *AlphaEvalJobService) Get(ctx context.Context, jobID string, query AlphaEvalJobGetParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/jobs/%s/result", query.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a job.
func (r *AlphaEvalJobService) Cancel(ctx context.Context, jobID string, body AlphaEvalJobCancelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/jobs/%s", body.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the status of a job.
func (r *AlphaEvalJobService) Status(ctx context.Context, jobID string, query AlphaEvalJobStatusParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s/jobs/%s", query.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AlphaEvalJobGetParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type AlphaEvalJobCancelParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type AlphaEvalJobStatusParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}
