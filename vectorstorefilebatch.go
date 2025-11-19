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
)

// VectorStoreFileBatchService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorStoreFileBatchService] method instead.
type VectorStoreFileBatchService struct {
	Options []option.RequestOption
}

// NewVectorStoreFileBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVectorStoreFileBatchService(opts ...option.RequestOption) (r VectorStoreFileBatchService) {
	r = VectorStoreFileBatchService{}
	r.Options = opts
	return
}

// Create a vector store file batch.
//
// Generate an OpenAI-compatible vector store file batch for the given vector
// store.
func (r *VectorStoreFileBatchService) New(ctx context.Context, vectorStoreID string, body VectorStoreFileBatchNewParams, opts ...option.RequestOption) (res *VectorStoreFileBatches, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a vector store file batch.
func (r *VectorStoreFileBatchService) Get(ctx context.Context, batchID string, query VectorStoreFileBatchGetParams, opts ...option.RequestOption) (res *VectorStoreFileBatches, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches/%s", query.VectorStoreID, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancels a vector store file batch.
func (r *VectorStoreFileBatchService) Cancel(ctx context.Context, batchID string, body VectorStoreFileBatchCancelParams, opts ...option.RequestOption) (res *VectorStoreFileBatches, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches/%s/cancel", body.VectorStoreID, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Returns a list of vector store files in a batch.
func (r *VectorStoreFileBatchService) ListFiles(ctx context.Context, batchID string, params VectorStoreFileBatchListFilesParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[VectorStoreFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches/%s/files", params.VectorStoreID, batchID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns a list of vector store files in a batch.
func (r *VectorStoreFileBatchService) ListFilesAutoPaging(ctx context.Context, batchID string, params VectorStoreFileBatchListFilesParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[VectorStoreFile] {
	return pagination.NewOpenAICursorPageAutoPager(r.ListFiles(ctx, batchID, params, opts...))
}

// Response from listing files in a vector store file batch.
type ListVectorStoreFilesInBatchResponse struct {
	Data    []VectorStoreFile `json:"data,required"`
	FirstID string            `json:"first_id,nullable"`
	HasMore bool              `json:"has_more"`
	LastID  string            `json:"last_id,nullable"`
	Object  string            `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListVectorStoreFilesInBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *ListVectorStoreFilesInBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI Vector Store File Batch object.
type VectorStoreFileBatches struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	// File processing status counts for a vector store.
	FileCounts VectorStoreFileBatchesFileCounts `json:"file_counts,required"`
	// Any of "completed", "in_progress", "cancelled", "failed".
	Status        VectorStoreFileBatchesStatus `json:"status,required"`
	VectorStoreID string                       `json:"vector_store_id,required"`
	Object        string                       `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		FileCounts    respjson.Field
		Status        respjson.Field
		VectorStoreID respjson.Field
		Object        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileBatches) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileBatches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File processing status counts for a vector store.
type VectorStoreFileBatchesFileCounts struct {
	Cancelled  int64 `json:"cancelled,required"`
	Completed  int64 `json:"completed,required"`
	Failed     int64 `json:"failed,required"`
	InProgress int64 `json:"in_progress,required"`
	Total      int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Completed   respjson.Field
		Failed      respjson.Field
		InProgress  respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileBatchesFileCounts) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileBatchesFileCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileBatchesStatus string

const (
	VectorStoreFileBatchesStatusCompleted  VectorStoreFileBatchesStatus = "completed"
	VectorStoreFileBatchesStatusInProgress VectorStoreFileBatchesStatus = "in_progress"
	VectorStoreFileBatchesStatusCancelled  VectorStoreFileBatchesStatus = "cancelled"
	VectorStoreFileBatchesStatusFailed     VectorStoreFileBatchesStatus = "failed"
)

type VectorStoreFileBatchNewParams struct {
	FileIDs    []string       `json:"file_ids,omitzero,required"`
	Attributes map[string]any `json:"attributes,omitzero"`
	// Automatic chunking strategy for vector store files.
	ChunkingStrategy VectorStoreFileBatchNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreFileBatchNewParamsChunkingStrategyUnion struct {
	OfAuto   *VectorStoreFileBatchNewParamsChunkingStrategyAuto   `json:",omitzero,inline"`
	OfStatic *VectorStoreFileBatchNewParamsChunkingStrategyStatic `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileBatchNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic)
}
func (u *VectorStoreFileBatchNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileBatchNewParamsChunkingStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfAuto) {
		return u.OfAuto
	} else if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileBatchNewParamsChunkingStrategyUnion) GetStatic() *VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic {
	if vt := u.OfStatic; vt != nil {
		return &vt.Static
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileBatchNewParamsChunkingStrategyUnion) GetType() *string {
	if vt := u.OfAuto; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorStoreFileBatchNewParamsChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreFileBatchNewParamsChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreFileBatchNewParamsChunkingStrategyStatic]("static"),
	)
}

// Automatic chunking strategy for vector store files.
type VectorStoreFileBatchNewParamsChunkingStrategyAuto struct {
	// Any of "auto".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyAuto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileBatchNewParamsChunkingStrategyAuto](
		"type", "auto",
	)
}

// Static chunking strategy with configurable parameters.
//
// The property Static is required.
type VectorStoreFileBatchNewParamsChunkingStrategyStatic struct {
	// Configuration for static chunking strategy.
	Static VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic `json:"static,omitzero,required"`
	// Any of "static".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileBatchNewParamsChunkingStrategyStatic](
		"type", "static",
	)
}

// Configuration for static chunking strategy.
type VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic struct {
	ChunkOverlapTokens param.Opt[int64] `json:"chunk_overlap_tokens,omitzero"`
	MaxChunkSizeTokens param.Opt[int64] `json:"max_chunk_size_tokens,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileBatchGetParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

type VectorStoreFileBatchCancelParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

type VectorStoreFileBatchListFilesParams struct {
	VectorStoreID string            `path:"vector_store_id,required" json:"-"`
	After         param.Opt[string] `query:"after,omitzero" json:"-"`
	Before        param.Opt[string] `query:"before,omitzero" json:"-"`
	Filter        param.Opt[string] `query:"filter,omitzero" json:"-"`
	Limit         param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Order         param.Opt[string] `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VectorStoreFileBatchListFilesParams]'s query parameters as
// `url.Values`.
func (r VectorStoreFileBatchListFilesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
