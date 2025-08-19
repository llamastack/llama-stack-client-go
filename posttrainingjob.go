// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// PostTrainingJobService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostTrainingJobService] method instead.
type PostTrainingJobService struct {
	Options []option.RequestOption
}

// NewPostTrainingJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostTrainingJobService(opts ...option.RequestOption) (r PostTrainingJobService) {
	r = PostTrainingJobService{}
	r.Options = opts
	return
}

// Get all training jobs.
func (r *PostTrainingJobService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ListPostTrainingJobsResponseData, err error) {
	var env ListPostTrainingJobsResponse
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get the artifacts of a training job.
func (r *PostTrainingJobService) Artifacts(ctx context.Context, query PostTrainingJobArtifactsParams, opts ...option.RequestOption) (res *PostTrainingJobArtifactsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/job/artifacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel a training job.
func (r *PostTrainingJobService) Cancel(ctx context.Context, body PostTrainingJobCancelParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/post-training/job/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the status of a training job.
func (r *PostTrainingJobService) Status(ctx context.Context, query PostTrainingJobStatusParams, opts ...option.RequestOption) (res *PostTrainingJobStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/job/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Artifacts of a finetuning job.
type PostTrainingJobArtifactsResponse struct {
	Checkpoints []any  `json:"checkpoints,required"`
	JobUuid     string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checkpoints respjson.Field
		JobUuid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingJobArtifactsResponse) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingJobArtifactsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a finetuning job.
type PostTrainingJobStatusResponse struct {
	Checkpoints []any  `json:"checkpoints,required"`
	JobUuid     string `json:"job_uuid,required"`
	// Any of "completed", "in_progress", "failed", "scheduled", "cancelled".
	Status             PostTrainingJobStatusResponseStatus                             `json:"status,required"`
	CompletedAt        time.Time                                                       `json:"completed_at" format:"date-time"`
	ResourcesAllocated map[string]PostTrainingJobStatusResponseResourcesAllocatedUnion `json:"resources_allocated"`
	ScheduledAt        time.Time                                                       `json:"scheduled_at" format:"date-time"`
	StartedAt          time.Time                                                       `json:"started_at" format:"date-time"`
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
func (r PostTrainingJobStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingJobStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingJobStatusResponseStatus string

const (
	PostTrainingJobStatusResponseStatusCompleted  PostTrainingJobStatusResponseStatus = "completed"
	PostTrainingJobStatusResponseStatusInProgress PostTrainingJobStatusResponseStatus = "in_progress"
	PostTrainingJobStatusResponseStatusFailed     PostTrainingJobStatusResponseStatus = "failed"
	PostTrainingJobStatusResponseStatusScheduled  PostTrainingJobStatusResponseStatus = "scheduled"
	PostTrainingJobStatusResponseStatusCancelled  PostTrainingJobStatusResponseStatus = "cancelled"
)

// PostTrainingJobStatusResponseResourcesAllocatedUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type PostTrainingJobStatusResponseResourcesAllocatedUnion struct {
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

func (u PostTrainingJobStatusResponseResourcesAllocatedUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PostTrainingJobStatusResponseResourcesAllocatedUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PostTrainingJobStatusResponseResourcesAllocatedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PostTrainingJobStatusResponseResourcesAllocatedUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PostTrainingJobStatusResponseResourcesAllocatedUnion) RawJSON() string { return u.JSON.raw }

func (r *PostTrainingJobStatusResponseResourcesAllocatedUnion) UnmarshalJSON(data []byte) error {
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

type PostTrainingJobArtifactsParams struct {
	// The UUID of the job to get the artifacts of.
	JobUuid string `query:"job_uuid,required" json:"-"`
	paramObj
}

// URLQuery serializes [PostTrainingJobArtifactsParams]'s query parameters as
// `url.Values`.
func (r PostTrainingJobArtifactsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PostTrainingJobCancelParams struct {
	// The UUID of the job to cancel.
	JobUuid string `json:"job_uuid,required"`
	paramObj
}

func (r PostTrainingJobCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingJobCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingJobCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingJobStatusParams struct {
	// The UUID of the job to get the status of.
	JobUuid string `query:"job_uuid,required" json:"-"`
	paramObj
}

// URLQuery serializes [PostTrainingJobStatusParams]'s query parameters as
// `url.Values`.
func (r PostTrainingJobStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
