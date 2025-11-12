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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1alpha/eval/benchmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// A benchmark resource for evaluating model performance.
type Benchmark struct {
	// Identifier of the dataset to use for the benchmark evaluation
	DatasetID  string `json:"dataset_id,required"`
	Identifier string `json:"identifier,required"`
	// Metadata for this evaluation task
	Metadata   map[string]BenchmarkMetadataUnion `json:"metadata,required"`
	ProviderID string                            `json:"provider_id,required"`
	// List of scoring function identifiers to apply during evaluation
	ScoringFunctions []string `json:"scoring_functions,required"`
	// The resource type, always benchmark
	Type               constant.Benchmark `json:"type,required"`
	ProviderResourceID string             `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatasetID          respjson.Field
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		ScoringFunctions   respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Benchmark) RawJSON() string { return r.JSON.raw }
func (r *Benchmark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BenchmarkMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BenchmarkMetadataUnion struct {
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

func (u BenchmarkMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BenchmarkMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BenchmarkMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BenchmarkMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BenchmarkMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BenchmarkMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListBenchmarksResponse struct {
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
	Metadata map[string]AlphaBenchmarkRegisterParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r AlphaBenchmarkRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaBenchmarkRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaBenchmarkRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaBenchmarkRegisterParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaBenchmarkRegisterParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AlphaBenchmarkRegisterParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaBenchmarkRegisterParamsMetadataUnion) asAny() any {
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
