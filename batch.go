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
	"net/url"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/pagination"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// BatchService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	Options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.Options = opts
	return
}

// Create a new batch for processing multiple API requests.
func (r *BatchService) New(ctx context.Context, body BatchNewParams, opts ...option.RequestOption) (res *BatchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve information about a specific batch.
func (r *BatchService) Get(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/batches/%s", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all batches for the current user.
func (r *BatchService) List(ctx context.Context, query BatchListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[BatchListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/batches"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all batches for the current user.
func (r *BatchService) ListAutoPaging(ctx context.Context, query BatchListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[BatchListResponse] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Cancel a batch that is in progress.
func (r *BatchService) Cancel(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/batches/%s/cancel", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type BatchNewResponse struct {
	ID               string         `json:"id,required"`
	CompletionWindow string         `json:"completion_window,required"`
	CreatedAt        int64          `json:"created_at,required"`
	Endpoint         string         `json:"endpoint,required"`
	InputFileID      string         `json:"input_file_id,required"`
	Object           constant.Batch `json:"object,required"`
	// Any of "validating", "failed", "in_progress", "finalizing", "completed",
	// "expired", "cancelling", "cancelled".
	Status        BatchNewResponseStatus        `json:"status,required"`
	CancelledAt   int64                         `json:"cancelled_at,nullable"`
	CancellingAt  int64                         `json:"cancelling_at,nullable"`
	CompletedAt   int64                         `json:"completed_at,nullable"`
	ErrorFileID   string                        `json:"error_file_id,nullable"`
	Errors        BatchNewResponseErrors        `json:"errors,nullable"`
	ExpiredAt     int64                         `json:"expired_at,nullable"`
	ExpiresAt     int64                         `json:"expires_at,nullable"`
	FailedAt      int64                         `json:"failed_at,nullable"`
	FinalizingAt  int64                         `json:"finalizing_at,nullable"`
	InProgressAt  int64                         `json:"in_progress_at,nullable"`
	Metadata      map[string]string             `json:"metadata,nullable"`
	Model         string                        `json:"model,nullable"`
	OutputFileID  string                        `json:"output_file_id,nullable"`
	RequestCounts BatchNewResponseRequestCounts `json:"request_counts,nullable"`
	Usage         BatchNewResponseUsage         `json:"usage,nullable"`
	ExtraFields   map[string]any                `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletionWindow respjson.Field
		CreatedAt        respjson.Field
		Endpoint         respjson.Field
		InputFileID      respjson.Field
		Object           respjson.Field
		Status           respjson.Field
		CancelledAt      respjson.Field
		CancellingAt     respjson.Field
		CompletedAt      respjson.Field
		ErrorFileID      respjson.Field
		Errors           respjson.Field
		ExpiredAt        respjson.Field
		ExpiresAt        respjson.Field
		FailedAt         respjson.Field
		FinalizingAt     respjson.Field
		InProgressAt     respjson.Field
		Metadata         respjson.Field
		Model            respjson.Field
		OutputFileID     respjson.Field
		RequestCounts    respjson.Field
		Usage            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseStatus string

const (
	BatchNewResponseStatusValidating BatchNewResponseStatus = "validating"
	BatchNewResponseStatusFailed     BatchNewResponseStatus = "failed"
	BatchNewResponseStatusInProgress BatchNewResponseStatus = "in_progress"
	BatchNewResponseStatusFinalizing BatchNewResponseStatus = "finalizing"
	BatchNewResponseStatusCompleted  BatchNewResponseStatus = "completed"
	BatchNewResponseStatusExpired    BatchNewResponseStatus = "expired"
	BatchNewResponseStatusCancelling BatchNewResponseStatus = "cancelling"
	BatchNewResponseStatusCancelled  BatchNewResponseStatus = "cancelled"
)

type BatchNewResponseErrors struct {
	Data        []BatchNewResponseErrorsData `json:"data,nullable"`
	Object      string                       `json:"object,nullable"`
	ExtraFields map[string]any               `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseErrors) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseErrorsData struct {
	Code        string         `json:"code,nullable"`
	Line        int64          `json:"line,nullable"`
	Message     string         `json:"message,nullable"`
	Param       string         `json:"param,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Line        respjson.Field
		Message     respjson.Field
		Param       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseErrorsData) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseErrorsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseRequestCounts struct {
	Completed   int64          `json:"completed,required"`
	Failed      int64          `json:"failed,required"`
	Total       int64          `json:"total,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Failed      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseRequestCounts) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseRequestCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseUsage struct {
	InputTokens         int64                                    `json:"input_tokens,required"`
	InputTokensDetails  BatchNewResponseUsageInputTokensDetails  `json:"input_tokens_details,required"`
	OutputTokens        int64                                    `json:"output_tokens,required"`
	OutputTokensDetails BatchNewResponseUsageOutputTokensDetails `json:"output_tokens_details,required"`
	TotalTokens         int64                                    `json:"total_tokens,required"`
	ExtraFields         map[string]any                           `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokens        respjson.Field
		OutputTokensDetails respjson.Field
		TotalTokens         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseUsageInputTokensDetails struct {
	CachedTokens int64          `json:"cached_tokens,required"`
	ExtraFields  map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseUsageOutputTokensDetails struct {
	ReasoningTokens int64          `json:"reasoning_tokens,required"`
	ExtraFields     map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponse struct {
	ID               string         `json:"id,required"`
	CompletionWindow string         `json:"completion_window,required"`
	CreatedAt        int64          `json:"created_at,required"`
	Endpoint         string         `json:"endpoint,required"`
	InputFileID      string         `json:"input_file_id,required"`
	Object           constant.Batch `json:"object,required"`
	// Any of "validating", "failed", "in_progress", "finalizing", "completed",
	// "expired", "cancelling", "cancelled".
	Status        BatchGetResponseStatus        `json:"status,required"`
	CancelledAt   int64                         `json:"cancelled_at,nullable"`
	CancellingAt  int64                         `json:"cancelling_at,nullable"`
	CompletedAt   int64                         `json:"completed_at,nullable"`
	ErrorFileID   string                        `json:"error_file_id,nullable"`
	Errors        BatchGetResponseErrors        `json:"errors,nullable"`
	ExpiredAt     int64                         `json:"expired_at,nullable"`
	ExpiresAt     int64                         `json:"expires_at,nullable"`
	FailedAt      int64                         `json:"failed_at,nullable"`
	FinalizingAt  int64                         `json:"finalizing_at,nullable"`
	InProgressAt  int64                         `json:"in_progress_at,nullable"`
	Metadata      map[string]string             `json:"metadata,nullable"`
	Model         string                        `json:"model,nullable"`
	OutputFileID  string                        `json:"output_file_id,nullable"`
	RequestCounts BatchGetResponseRequestCounts `json:"request_counts,nullable"`
	Usage         BatchGetResponseUsage         `json:"usage,nullable"`
	ExtraFields   map[string]any                `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletionWindow respjson.Field
		CreatedAt        respjson.Field
		Endpoint         respjson.Field
		InputFileID      respjson.Field
		Object           respjson.Field
		Status           respjson.Field
		CancelledAt      respjson.Field
		CancellingAt     respjson.Field
		CompletedAt      respjson.Field
		ErrorFileID      respjson.Field
		Errors           respjson.Field
		ExpiredAt        respjson.Field
		ExpiresAt        respjson.Field
		FailedAt         respjson.Field
		FinalizingAt     respjson.Field
		InProgressAt     respjson.Field
		Metadata         respjson.Field
		Model            respjson.Field
		OutputFileID     respjson.Field
		RequestCounts    respjson.Field
		Usage            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseStatus string

const (
	BatchGetResponseStatusValidating BatchGetResponseStatus = "validating"
	BatchGetResponseStatusFailed     BatchGetResponseStatus = "failed"
	BatchGetResponseStatusInProgress BatchGetResponseStatus = "in_progress"
	BatchGetResponseStatusFinalizing BatchGetResponseStatus = "finalizing"
	BatchGetResponseStatusCompleted  BatchGetResponseStatus = "completed"
	BatchGetResponseStatusExpired    BatchGetResponseStatus = "expired"
	BatchGetResponseStatusCancelling BatchGetResponseStatus = "cancelling"
	BatchGetResponseStatusCancelled  BatchGetResponseStatus = "cancelled"
)

type BatchGetResponseErrors struct {
	Data        []BatchGetResponseErrorsData `json:"data,nullable"`
	Object      string                       `json:"object,nullable"`
	ExtraFields map[string]any               `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseErrors) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseErrorsData struct {
	Code        string         `json:"code,nullable"`
	Line        int64          `json:"line,nullable"`
	Message     string         `json:"message,nullable"`
	Param       string         `json:"param,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Line        respjson.Field
		Message     respjson.Field
		Param       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseErrorsData) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseErrorsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseRequestCounts struct {
	Completed   int64          `json:"completed,required"`
	Failed      int64          `json:"failed,required"`
	Total       int64          `json:"total,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Failed      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseRequestCounts) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseRequestCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseUsage struct {
	InputTokens         int64                                    `json:"input_tokens,required"`
	InputTokensDetails  BatchGetResponseUsageInputTokensDetails  `json:"input_tokens_details,required"`
	OutputTokens        int64                                    `json:"output_tokens,required"`
	OutputTokensDetails BatchGetResponseUsageOutputTokensDetails `json:"output_tokens_details,required"`
	TotalTokens         int64                                    `json:"total_tokens,required"`
	ExtraFields         map[string]any                           `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokens        respjson.Field
		OutputTokensDetails respjson.Field
		TotalTokens         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseUsageInputTokensDetails struct {
	CachedTokens int64          `json:"cached_tokens,required"`
	ExtraFields  map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponseUsageOutputTokensDetails struct {
	ReasoningTokens int64          `json:"reasoning_tokens,required"`
	ExtraFields     map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponse struct {
	ID               string         `json:"id,required"`
	CompletionWindow string         `json:"completion_window,required"`
	CreatedAt        int64          `json:"created_at,required"`
	Endpoint         string         `json:"endpoint,required"`
	InputFileID      string         `json:"input_file_id,required"`
	Object           constant.Batch `json:"object,required"`
	// Any of "validating", "failed", "in_progress", "finalizing", "completed",
	// "expired", "cancelling", "cancelled".
	Status        BatchListResponseStatus        `json:"status,required"`
	CancelledAt   int64                          `json:"cancelled_at,nullable"`
	CancellingAt  int64                          `json:"cancelling_at,nullable"`
	CompletedAt   int64                          `json:"completed_at,nullable"`
	ErrorFileID   string                         `json:"error_file_id,nullable"`
	Errors        BatchListResponseErrors        `json:"errors,nullable"`
	ExpiredAt     int64                          `json:"expired_at,nullable"`
	ExpiresAt     int64                          `json:"expires_at,nullable"`
	FailedAt      int64                          `json:"failed_at,nullable"`
	FinalizingAt  int64                          `json:"finalizing_at,nullable"`
	InProgressAt  int64                          `json:"in_progress_at,nullable"`
	Metadata      map[string]string              `json:"metadata,nullable"`
	Model         string                         `json:"model,nullable"`
	OutputFileID  string                         `json:"output_file_id,nullable"`
	RequestCounts BatchListResponseRequestCounts `json:"request_counts,nullable"`
	Usage         BatchListResponseUsage         `json:"usage,nullable"`
	ExtraFields   map[string]any                 `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletionWindow respjson.Field
		CreatedAt        respjson.Field
		Endpoint         respjson.Field
		InputFileID      respjson.Field
		Object           respjson.Field
		Status           respjson.Field
		CancelledAt      respjson.Field
		CancellingAt     respjson.Field
		CompletedAt      respjson.Field
		ErrorFileID      respjson.Field
		Errors           respjson.Field
		ExpiredAt        respjson.Field
		ExpiresAt        respjson.Field
		FailedAt         respjson.Field
		FinalizingAt     respjson.Field
		InProgressAt     respjson.Field
		Metadata         respjson.Field
		Model            respjson.Field
		OutputFileID     respjson.Field
		RequestCounts    respjson.Field
		Usage            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseStatus string

const (
	BatchListResponseStatusValidating BatchListResponseStatus = "validating"
	BatchListResponseStatusFailed     BatchListResponseStatus = "failed"
	BatchListResponseStatusInProgress BatchListResponseStatus = "in_progress"
	BatchListResponseStatusFinalizing BatchListResponseStatus = "finalizing"
	BatchListResponseStatusCompleted  BatchListResponseStatus = "completed"
	BatchListResponseStatusExpired    BatchListResponseStatus = "expired"
	BatchListResponseStatusCancelling BatchListResponseStatus = "cancelling"
	BatchListResponseStatusCancelled  BatchListResponseStatus = "cancelled"
)

type BatchListResponseErrors struct {
	Data        []BatchListResponseErrorsData `json:"data,nullable"`
	Object      string                        `json:"object,nullable"`
	ExtraFields map[string]any                `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseErrors) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseErrorsData struct {
	Code        string         `json:"code,nullable"`
	Line        int64          `json:"line,nullable"`
	Message     string         `json:"message,nullable"`
	Param       string         `json:"param,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Line        respjson.Field
		Message     respjson.Field
		Param       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseErrorsData) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseErrorsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseRequestCounts struct {
	Completed   int64          `json:"completed,required"`
	Failed      int64          `json:"failed,required"`
	Total       int64          `json:"total,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Failed      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseRequestCounts) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseRequestCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseUsage struct {
	InputTokens         int64                                     `json:"input_tokens,required"`
	InputTokensDetails  BatchListResponseUsageInputTokensDetails  `json:"input_tokens_details,required"`
	OutputTokens        int64                                     `json:"output_tokens,required"`
	OutputTokensDetails BatchListResponseUsageOutputTokensDetails `json:"output_tokens_details,required"`
	TotalTokens         int64                                     `json:"total_tokens,required"`
	ExtraFields         map[string]any                            `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokens        respjson.Field
		OutputTokensDetails respjson.Field
		TotalTokens         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseUsageInputTokensDetails struct {
	CachedTokens int64          `json:"cached_tokens,required"`
	ExtraFields  map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseUsageOutputTokensDetails struct {
	ReasoningTokens int64          `json:"reasoning_tokens,required"`
	ExtraFields     map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponse struct {
	ID               string         `json:"id,required"`
	CompletionWindow string         `json:"completion_window,required"`
	CreatedAt        int64          `json:"created_at,required"`
	Endpoint         string         `json:"endpoint,required"`
	InputFileID      string         `json:"input_file_id,required"`
	Object           constant.Batch `json:"object,required"`
	// Any of "validating", "failed", "in_progress", "finalizing", "completed",
	// "expired", "cancelling", "cancelled".
	Status        BatchCancelResponseStatus        `json:"status,required"`
	CancelledAt   int64                            `json:"cancelled_at,nullable"`
	CancellingAt  int64                            `json:"cancelling_at,nullable"`
	CompletedAt   int64                            `json:"completed_at,nullable"`
	ErrorFileID   string                           `json:"error_file_id,nullable"`
	Errors        BatchCancelResponseErrors        `json:"errors,nullable"`
	ExpiredAt     int64                            `json:"expired_at,nullable"`
	ExpiresAt     int64                            `json:"expires_at,nullable"`
	FailedAt      int64                            `json:"failed_at,nullable"`
	FinalizingAt  int64                            `json:"finalizing_at,nullable"`
	InProgressAt  int64                            `json:"in_progress_at,nullable"`
	Metadata      map[string]string                `json:"metadata,nullable"`
	Model         string                           `json:"model,nullable"`
	OutputFileID  string                           `json:"output_file_id,nullable"`
	RequestCounts BatchCancelResponseRequestCounts `json:"request_counts,nullable"`
	Usage         BatchCancelResponseUsage         `json:"usage,nullable"`
	ExtraFields   map[string]any                   `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletionWindow respjson.Field
		CreatedAt        respjson.Field
		Endpoint         respjson.Field
		InputFileID      respjson.Field
		Object           respjson.Field
		Status           respjson.Field
		CancelledAt      respjson.Field
		CancellingAt     respjson.Field
		CompletedAt      respjson.Field
		ErrorFileID      respjson.Field
		Errors           respjson.Field
		ExpiredAt        respjson.Field
		ExpiresAt        respjson.Field
		FailedAt         respjson.Field
		FinalizingAt     respjson.Field
		InProgressAt     respjson.Field
		Metadata         respjson.Field
		Model            respjson.Field
		OutputFileID     respjson.Field
		RequestCounts    respjson.Field
		Usage            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseStatus string

const (
	BatchCancelResponseStatusValidating BatchCancelResponseStatus = "validating"
	BatchCancelResponseStatusFailed     BatchCancelResponseStatus = "failed"
	BatchCancelResponseStatusInProgress BatchCancelResponseStatus = "in_progress"
	BatchCancelResponseStatusFinalizing BatchCancelResponseStatus = "finalizing"
	BatchCancelResponseStatusCompleted  BatchCancelResponseStatus = "completed"
	BatchCancelResponseStatusExpired    BatchCancelResponseStatus = "expired"
	BatchCancelResponseStatusCancelling BatchCancelResponseStatus = "cancelling"
	BatchCancelResponseStatusCancelled  BatchCancelResponseStatus = "cancelled"
)

type BatchCancelResponseErrors struct {
	Data        []BatchCancelResponseErrorsData `json:"data,nullable"`
	Object      string                          `json:"object,nullable"`
	ExtraFields map[string]any                  `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseErrors) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseErrorsData struct {
	Code        string         `json:"code,nullable"`
	Line        int64          `json:"line,nullable"`
	Message     string         `json:"message,nullable"`
	Param       string         `json:"param,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Line        respjson.Field
		Message     respjson.Field
		Param       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseErrorsData) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseErrorsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseRequestCounts struct {
	Completed   int64          `json:"completed,required"`
	Failed      int64          `json:"failed,required"`
	Total       int64          `json:"total,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Failed      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseRequestCounts) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseRequestCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseUsage struct {
	InputTokens         int64                                       `json:"input_tokens,required"`
	InputTokensDetails  BatchCancelResponseUsageInputTokensDetails  `json:"input_tokens_details,required"`
	OutputTokens        int64                                       `json:"output_tokens,required"`
	OutputTokensDetails BatchCancelResponseUsageOutputTokensDetails `json:"output_tokens_details,required"`
	TotalTokens         int64                                       `json:"total_tokens,required"`
	ExtraFields         map[string]any                              `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokens        respjson.Field
		OutputTokensDetails respjson.Field
		TotalTokens         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseUsageInputTokensDetails struct {
	CachedTokens int64          `json:"cached_tokens,required"`
	ExtraFields  map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCancelResponseUsageOutputTokensDetails struct {
	ReasoningTokens int64          `json:"reasoning_tokens,required"`
	ExtraFields     map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCancelResponseUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *BatchCancelResponseUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewParams struct {
	// The endpoint to be used for all requests in the batch.
	Endpoint string `json:"endpoint,required"`
	// The ID of an uploaded file containing requests for the batch.
	InputFileID string `json:"input_file_id,required"`
	// Optional idempotency key. When provided, enables idempotent behavior.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Optional metadata for the batch.
	Metadata map[string]string `json:"metadata,omitzero"`
	// The time window within which the batch should be processed.
	//
	// This field can be elided, and will marshal its zero value as "24h".
	CompletionWindow constant.String24h `json:"completion_window,required"`
	paramObj
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListParams struct {
	// Optional cursor for pagination. Returns batches after this ID.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of batches to return. Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchListParams]'s query parameters as `url.Values`.
func (r BatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
