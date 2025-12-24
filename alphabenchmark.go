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

// AlphaBenchmarkService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaBenchmarkService] method instead.
type AlphaBenchmarkService struct {
	Options []option.RequestOption
}

// NewAlphaBenchmarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaBenchmarkService(opts ...option.RequestOption) (r AlphaBenchmarkService) {
	r = AlphaBenchmarkService{}
	r.Options = opts
	return
}

// Get a benchmark by its ID.
func (r *AlphaBenchmarkService) Get(ctx context.Context, benchmarkID string, opts ...option.RequestOption) (res *Benchmark, err error) {
	opts = slices.Concat(r.Options, opts)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all benchmarks.
func (r *AlphaBenchmarkService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Benchmark, err error) {
	var env ListBenchmarksResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/eval/benchmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Register a benchmark.
//
// Deprecated: deprecated
func (r *AlphaBenchmarkService) Register(ctx context.Context, body AlphaBenchmarkRegisterParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1alpha/eval/benchmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Unregister a benchmark.
//
// Deprecated: deprecated
func (r *AlphaBenchmarkService) Unregister(ctx context.Context, benchmarkID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1alpha/eval/benchmarks/%s", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A benchmark resource for evaluating model performance.
type Benchmark struct {
	// Identifier of the dataset to use for the benchmark evaluation.
	DatasetID string `json:"dataset_id,required"`
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// List of scoring function identifiers to apply during evaluation.
	ScoringFunctions []string `json:"scoring_functions,required"`
	// Metadata for this evaluation task.
	Metadata map[string]any `json:"metadata"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// The resource type, always benchmark.
	//
	// Any of "benchmark".
	Type BenchmarkType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatasetID          respjson.Field
		Identifier         respjson.Field
		ProviderID         respjson.Field
		ScoringFunctions   respjson.Field
		Metadata           respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Benchmark) RawJSON() string { return r.JSON.raw }
func (r *Benchmark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The resource type, always benchmark.
type BenchmarkType string

const (
	BenchmarkTypeBenchmark BenchmarkType = "benchmark"
)

// Response containing a list of benchmark objects.
type ListBenchmarksResponse struct {
	// List of benchmark objects.
	Data []Benchmark `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListBenchmarksResponse) RawJSON() string { return r.JSON.raw }
func (r *ListBenchmarksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaBenchmarkRegisterParams struct {
	// The ID of the benchmark to register.
	BenchmarkID string `json:"benchmark_id,required"`
	// The ID of the dataset to use for the benchmark.
	DatasetID string `json:"dataset_id,required"`
	// The scoring functions to use for the benchmark.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	// The ID of the provider benchmark to use for the benchmark.
	ProviderBenchmarkID param.Opt[string] `json:"provider_benchmark_id,omitzero"`
	// The ID of the provider to use for the benchmark.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The metadata to use for the benchmark.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r AlphaBenchmarkRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaBenchmarkRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaBenchmarkRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
