// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
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
func (r *AlphaPostTrainingJobService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ListPostTrainingJobsResponseData, err error) {
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	// List of model checkpoints created during training
	Checkpoints []AlphaPostTrainingJobArtifactsResponseCheckpoint `json:"checkpoints,required"`
	// Unique identifier for the training job
	JobUuid string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checkpoints respjson.Field
		JobUuid     respjson.Field
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
	// Timestamp when the checkpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Training epoch when the checkpoint was saved
	Epoch int64 `json:"epoch,required"`
	// Unique identifier for the checkpoint
	Identifier string `json:"identifier,required"`
	// File system path where the checkpoint is stored
	Path string `json:"path,required"`
	// Identifier of the training job that created this checkpoint
	PostTrainingJobID string `json:"post_training_job_id,required"`
	// (Optional) Training metrics associated with this checkpoint
	TrainingMetrics AlphaPostTrainingJobArtifactsResponseCheckpointTrainingMetrics `json:"training_metrics"`
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

// (Optional) Training metrics associated with this checkpoint
type AlphaPostTrainingJobArtifactsResponseCheckpointTrainingMetrics struct {
	// Training epoch number
	Epoch int64 `json:"epoch,required"`
	// Perplexity metric indicating model confidence
	Perplexity float64 `json:"perplexity,required"`
	// Loss value on the training dataset
	TrainLoss float64 `json:"train_loss,required"`
	// Loss value on the validation dataset
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
	// List of model checkpoints created during training
	Checkpoints []AlphaPostTrainingJobStatusResponseCheckpoint `json:"checkpoints,required"`
	// Unique identifier for the training job
	JobUuid string `json:"job_uuid,required"`
	// Current status of the training job
	//
	// Any of "completed", "in_progress", "failed", "scheduled", "cancelled".
	Status AlphaPostTrainingJobStatusResponseStatus `json:"status,required"`
	// (Optional) Timestamp when the job finished, if completed
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// (Optional) Information about computational resources allocated to the job
	ResourcesAllocated map[string]AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion `json:"resources_allocated"`
	// (Optional) Timestamp when the job was scheduled
	ScheduledAt time.Time `json:"scheduled_at" format:"date-time"`
	// (Optional) Timestamp when the job execution began
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checkpoints        respjson.Field
		JobUuid            respjson.Field
		Status             respjson.Field
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

// Checkpoint created during training runs.
type AlphaPostTrainingJobStatusResponseCheckpoint struct {
	// Timestamp when the checkpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Training epoch when the checkpoint was saved
	Epoch int64 `json:"epoch,required"`
	// Unique identifier for the checkpoint
	Identifier string `json:"identifier,required"`
	// File system path where the checkpoint is stored
	Path string `json:"path,required"`
	// Identifier of the training job that created this checkpoint
	PostTrainingJobID string `json:"post_training_job_id,required"`
	// (Optional) Training metrics associated with this checkpoint
	TrainingMetrics AlphaPostTrainingJobStatusResponseCheckpointTrainingMetrics `json:"training_metrics"`
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

// (Optional) Training metrics associated with this checkpoint
type AlphaPostTrainingJobStatusResponseCheckpointTrainingMetrics struct {
	// Training epoch number
	Epoch int64 `json:"epoch,required"`
	// Perplexity metric indicating model confidence
	Perplexity float64 `json:"perplexity,required"`
	// Loss value on the training dataset
	TrainLoss float64 `json:"train_loss,required"`
	// Loss value on the validation dataset
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

// Current status of the training job
type AlphaPostTrainingJobStatusResponseStatus string

const (
	AlphaPostTrainingJobStatusResponseStatusCompleted  AlphaPostTrainingJobStatusResponseStatus = "completed"
	AlphaPostTrainingJobStatusResponseStatusInProgress AlphaPostTrainingJobStatusResponseStatus = "in_progress"
	AlphaPostTrainingJobStatusResponseStatusFailed     AlphaPostTrainingJobStatusResponseStatus = "failed"
	AlphaPostTrainingJobStatusResponseStatusScheduled  AlphaPostTrainingJobStatusResponseStatus = "scheduled"
	AlphaPostTrainingJobStatusResponseStatusCancelled  AlphaPostTrainingJobStatusResponseStatus = "cancelled"
)

// AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion struct {
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

func (u AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AlphaPostTrainingJobStatusResponseResourcesAllocatedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListPostTrainingJobsResponse struct {
	Data []ListPostTrainingJobsResponseData `json:"data,required"`
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

type ListPostTrainingJobsResponseData struct {
	JobUuid string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobUuid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListPostTrainingJobsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListPostTrainingJobsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaPostTrainingJobArtifactsParams struct {
	// The UUID of the job to get the artifacts of.
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
	// The UUID of the job to cancel.
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
	// The UUID of the job to get the status of.
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
