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
	"net/url"
	"slices"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// AlphaPostTrainingJobService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaPostTrainingJobService] method instead.
type AlphaPostTrainingJobService struct {
	Options []option.RequestOption
}

// NewAlphaPostTrainingJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlphaPostTrainingJobService(opts ...option.RequestOption) (r AlphaPostTrainingJobService) {
	r = AlphaPostTrainingJobService{}
	r.Options = opts
	return
}

// Get all training jobs.
func (r *AlphaPostTrainingJobService) List(ctx context.Context, opts ...option.RequestOption) (res *[]PostTrainingJob, err error) {
	var env ListPostTrainingJobsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/post-training/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get the artifacts of a training job.
func (r *AlphaPostTrainingJobService) Artifacts(ctx context.Context, query AlphaPostTrainingJobArtifactsParams, opts ...option.RequestOption) (res *AlphaPostTrainingJobArtifactsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/post-training/job/artifacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel a training job.
func (r *AlphaPostTrainingJobService) Cancel(ctx context.Context, body AlphaPostTrainingJobCancelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1alpha/post-training/job/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the status of a training job.
func (r *AlphaPostTrainingJobService) Status(ctx context.Context, query AlphaPostTrainingJobStatusParams, opts ...option.RequestOption) (res *AlphaPostTrainingJobStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/post-training/job/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Artifacts of a finetuning job.
type AlphaPostTrainingJobArtifactsResponse struct {
	JobUuid     string                                            `json:"job_uuid,required"`
	Checkpoints []AlphaPostTrainingJobArtifactsResponseCheckpoint `json:"checkpoints"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobUuid     respjson.Field
		Checkpoints respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaPostTrainingJobArtifactsResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaPostTrainingJobArtifactsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Checkpoint created during training runs.
type AlphaPostTrainingJobArtifactsResponseCheckpoint struct {
	CreatedAt         time.Time `json:"created_at,required" format:"date-time"`
	Epoch             int64     `json:"epoch,required"`
	Identifier        string    `json:"identifier,required"`
	Path              string    `json:"path,required"`
	PostTrainingJobID string    `json:"post_training_job_id,required"`
	// Training metrics captured during post-training jobs.
	TrainingMetrics AlphaPostTrainingJobArtifactsResponseCheckpointTrainingMetrics `json:"training_metrics,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		Epoch             respjson.Field
		Identifier        respjson.Field
		Path              respjson.Field
		PostTrainingJobID respjson.Field
		TrainingMetrics   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaPostTrainingJobArtifactsResponseCheckpoint) RawJSON() string { return r.JSON.raw }
func (r *AlphaPostTrainingJobArtifactsResponseCheckpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Training metrics captured during post-training jobs.
type AlphaPostTrainingJobArtifactsResponseCheckpointTrainingMetrics struct {
	Epoch          int64   `json:"epoch,required"`
	Perplexity     float64 `json:"perplexity,required"`
	TrainLoss      float64 `json:"train_loss,required"`
	ValidationLoss float64 `json:"validation_loss,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Epoch          respjson.Field
		Perplexity     respjson.Field
		TrainLoss      respjson.Field
		ValidationLoss respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaPostTrainingJobArtifactsResponseCheckpointTrainingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaPostTrainingJobArtifactsResponseCheckpointTrainingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a finetuning job.
type AlphaPostTrainingJobStatusResponse struct {
	JobUuid string `json:"job_uuid,required"`
	// Status of a job execution.
	//
	// Any of "completed", "in_progress", "failed", "scheduled", "cancelled".
	Status             AlphaPostTrainingJobStatusResponseStatus       `json:"status,required"`
	Checkpoints        []AlphaPostTrainingJobStatusResponseCheckpoint `json:"checkpoints"`
	CompletedAt        time.Time                                      `json:"completed_at,nullable" format:"date-time"`
	ResourcesAllocated map[string]any                                 `json:"resources_allocated,nullable"`
	ScheduledAt        time.Time                                      `json:"scheduled_at,nullable" format:"date-time"`
	StartedAt          time.Time                                      `json:"started_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobUuid            respjson.Field
		Status             respjson.Field
		Checkpoints        respjson.Field
		CompletedAt        respjson.Field
		ResourcesAllocated respjson.Field
		ScheduledAt        respjson.Field
		StartedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaPostTrainingJobStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaPostTrainingJobStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a job execution.
type AlphaPostTrainingJobStatusResponseStatus string

const (
	AlphaPostTrainingJobStatusResponseStatusCompleted  AlphaPostTrainingJobStatusResponseStatus = "completed"
	AlphaPostTrainingJobStatusResponseStatusInProgress AlphaPostTrainingJobStatusResponseStatus = "in_progress"
	AlphaPostTrainingJobStatusResponseStatusFailed     AlphaPostTrainingJobStatusResponseStatus = "failed"
	AlphaPostTrainingJobStatusResponseStatusScheduled  AlphaPostTrainingJobStatusResponseStatus = "scheduled"
	AlphaPostTrainingJobStatusResponseStatusCancelled  AlphaPostTrainingJobStatusResponseStatus = "cancelled"
)

// Checkpoint created during training runs.
type AlphaPostTrainingJobStatusResponseCheckpoint struct {
	CreatedAt         time.Time `json:"created_at,required" format:"date-time"`
	Epoch             int64     `json:"epoch,required"`
	Identifier        string    `json:"identifier,required"`
	Path              string    `json:"path,required"`
	PostTrainingJobID string    `json:"post_training_job_id,required"`
	// Training metrics captured during post-training jobs.
	TrainingMetrics AlphaPostTrainingJobStatusResponseCheckpointTrainingMetrics `json:"training_metrics,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		Epoch             respjson.Field
		Identifier        respjson.Field
		Path              respjson.Field
		PostTrainingJobID respjson.Field
		TrainingMetrics   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaPostTrainingJobStatusResponseCheckpoint) RawJSON() string { return r.JSON.raw }
func (r *AlphaPostTrainingJobStatusResponseCheckpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Training metrics captured during post-training jobs.
type AlphaPostTrainingJobStatusResponseCheckpointTrainingMetrics struct {
	Epoch          int64   `json:"epoch,required"`
	Perplexity     float64 `json:"perplexity,required"`
	TrainLoss      float64 `json:"train_loss,required"`
	ValidationLoss float64 `json:"validation_loss,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Epoch          respjson.Field
		Perplexity     respjson.Field
		TrainLoss      respjson.Field
		ValidationLoss respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaPostTrainingJobStatusResponseCheckpointTrainingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaPostTrainingJobStatusResponseCheckpointTrainingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListPostTrainingJobsResponse struct {
	Data []PostTrainingJob `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListPostTrainingJobsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListPostTrainingJobsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaPostTrainingJobArtifactsParams struct {
	JobUuid string `query:"job_uuid,required" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaPostTrainingJobArtifactsParams]'s query parameters as
// `url.Values`.
func (r AlphaPostTrainingJobArtifactsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlphaPostTrainingJobCancelParams struct {
	JobUuid string `json:"job_uuid,required"`
	paramObj
}

func (r AlphaPostTrainingJobCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingJobCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingJobCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaPostTrainingJobStatusParams struct {
	JobUuid string `query:"job_uuid,required" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaPostTrainingJobStatusParams]'s query parameters as
// `url.Values`.
func (r AlphaPostTrainingJobStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
