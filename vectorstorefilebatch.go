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

// Create a vector store file batch (OpenAI-compatible).
func (r *VectorStoreFileBatchService) New(ctx context.Context, vectorStoreID string, body VectorStoreFileBatchNewParams, opts ...option.RequestOption) (res *VectorStoreFileBatches, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a vector store file batch (OpenAI-compatible).
func (r *VectorStoreFileBatchService) Get(ctx context.Context, batchID string, query VectorStoreFileBatchGetParams, opts ...option.RequestOption) (res *VectorStoreFileBatches, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches/%s", query.VectorStoreID, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancel a vector store file batch (OpenAI-compatible).
func (r *VectorStoreFileBatchService) Cancel(ctx context.Context, batchID string, body VectorStoreFileBatchCancelParams, opts ...option.RequestOption) (res *VectorStoreFileBatches, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s/file_batches/%s/cancel", body.VectorStoreID, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// List files in a vector store file batch (OpenAI-compatible).
func (r *VectorStoreFileBatchService) ListFiles(ctx context.Context, batchID string, params VectorStoreFileBatchListFilesParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[VectorStoreFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
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

// List files in a vector store file batch (OpenAI-compatible).
func (r *VectorStoreFileBatchService) ListFilesAutoPaging(ctx context.Context, batchID string, params VectorStoreFileBatchListFilesParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[VectorStoreFile] {
	return pagination.NewOpenAICursorPageAutoPager(r.ListFiles(ctx, batchID, params, opts...))
}

// Response from listing files in a vector store file batch.
type ListVectorStoreFilesInBatchResponse struct {
	Data    []VectorStoreFile `json:"data" api:"required"`
	FirstID string            `json:"first_id" api:"required"`
	HasMore bool              `json:"has_more" api:"required"`
	LastID  string            `json:"last_id" api:"required"`
	// Any of "list".
	Object ListVectorStoreFilesInBatchResponseObject `json:"object"`
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

type ListVectorStoreFilesInBatchResponseObject string

const (
	ListVectorStoreFilesInBatchResponseObjectList ListVectorStoreFilesInBatchResponseObject = "list"
)

// OpenAI Vector Store File Batch object.
type VectorStoreFileBatches struct {
	ID        string `json:"id" api:"required"`
	CreatedAt int64  `json:"created_at" api:"required"`
	// File processing status counts for a vector store.
	FileCounts VectorStoreFileBatchesFileCounts `json:"file_counts" api:"required"`
	// Any of "in_progress", "completed", "cancelled", "failed".
	Status        VectorStoreFileBatchesStatus `json:"status" api:"required"`
	VectorStoreID string                       `json:"vector_store_id" api:"required"`
	// Any of "vector_store.files_batch".
	Object VectorStoreFileBatchesObject `json:"object"`
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
	Cancelled  int64 `json:"cancelled" api:"required"`
	Completed  int64 `json:"completed" api:"required"`
	Failed     int64 `json:"failed" api:"required"`
	InProgress int64 `json:"in_progress" api:"required"`
	Total      int64 `json:"total" api:"required"`
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
	VectorStoreFileBatchesStatusInProgress VectorStoreFileBatchesStatus = "in_progress"
	VectorStoreFileBatchesStatusCompleted  VectorStoreFileBatchesStatus = "completed"
	VectorStoreFileBatchesStatusCancelled  VectorStoreFileBatchesStatus = "cancelled"
	VectorStoreFileBatchesStatusFailed     VectorStoreFileBatchesStatus = "failed"
)

type VectorStoreFileBatchesObject string

const (
	VectorStoreFileBatchesObjectVectorStoreFilesBatch VectorStoreFileBatchesObject = "vector_store.files_batch"
)

type VectorStoreFileBatchNewParams struct {
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard. Keys are strings with a maximum
	// length of 64 characters. Values are strings with a maximum length of 512
	// characters, booleans, or numbers.
	Attributes map[string]VectorStoreFileBatchNewParamsAttributeUnion `json:"attributes,omitzero"`
	// Automatic chunking strategy for vector store files.
	ChunkingStrategy VectorStoreFileBatchNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	Files            []VectorStoreFileBatchNewParamsFile                `json:"files,omitzero"`
	FileIDs          []string                                           `json:"file_ids,omitzero"`
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
type VectorStoreFileBatchNewParamsAttributeUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileBatchNewParamsAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *VectorStoreFileBatchNewParamsAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileBatchNewParamsAttributeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreFileBatchNewParamsChunkingStrategyUnion struct {
	OfAuto       *VectorStoreFileBatchNewParamsChunkingStrategyAuto       `json:",omitzero,inline"`
	OfStatic     *VectorStoreFileBatchNewParamsChunkingStrategyStatic     `json:",omitzero,inline"`
	OfContextual *VectorStoreFileBatchNewParamsChunkingStrategyContextual `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileBatchNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic, u.OfContextual)
}
func (u *VectorStoreFileBatchNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileBatchNewParamsChunkingStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfAuto) {
		return u.OfAuto
	} else if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
	} else if !param.IsOmitted(u.OfContextual) {
		return u.OfContextual
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
func (u VectorStoreFileBatchNewParamsChunkingStrategyUnion) GetContextual() *VectorStoreFileBatchNewParamsChunkingStrategyContextualContextual {
	if vt := u.OfContextual; vt != nil {
		return &vt.Contextual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileBatchNewParamsChunkingStrategyUnion) GetType() *string {
	if vt := u.OfAuto; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContextual; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorStoreFileBatchNewParamsChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreFileBatchNewParamsChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreFileBatchNewParamsChunkingStrategyStatic]("static"),
		apijson.Discriminator[VectorStoreFileBatchNewParamsChunkingStrategyContextual]("contextual"),
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
	Static VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic `json:"static,omitzero" api:"required"`
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

// Contextual chunking strategy that uses an LLM to situate chunks within the
// document.
//
// The property Contextual is required.
type VectorStoreFileBatchNewParamsChunkingStrategyContextual struct {
	// Configuration for contextual chunking.
	Contextual VectorStoreFileBatchNewParamsChunkingStrategyContextualContextual `json:"contextual,omitzero" api:"required"`
	// Strategy type identifier.
	//
	// Any of "contextual".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileBatchNewParamsChunkingStrategyContextual](
		"type", "contextual",
	)
}

// Configuration for contextual chunking.
type VectorStoreFileBatchNewParamsChunkingStrategyContextualContextual struct {
	// Maximum concurrent LLM calls. Falls back to config default if not provided.
	MaxConcurrency param.Opt[int64] `json:"max_concurrency,omitzero"`
	// LLM model for generating context. Falls back to
	// VectorStoresConfig.contextual_retrieval_params.model if not provided.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Timeout per LLM call in seconds. Falls back to config default if not provided.
	TimeoutSeconds param.Opt[int64] `json:"timeout_seconds,omitzero"`
	// Tokens to overlap between adjacent chunks. Must be less than
	// max_chunk_size_tokens.
	ChunkOverlapTokens param.Opt[int64] `json:"chunk_overlap_tokens,omitzero"`
	// Prompt template for contextual retrieval. Uses WHOLE_DOCUMENT and CHUNK_CONTENT
	// placeholders wrapped in double curly braces.
	ContextPrompt param.Opt[string] `json:"context_prompt,omitzero"`
	// Maximum tokens per chunk. Suggested ~700 to allow room for prepended context.
	MaxChunkSizeTokens param.Opt[int64] `json:"max_chunk_size_tokens,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyContextualContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyContextualContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyContextualContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file entry for creating a vector store file batch with per-file options.
//
// The property FileID is required.
type VectorStoreFileBatchNewParamsFile struct {
	FileID string `json:"file_id" api:"required"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard. Keys are strings with a maximum
	// length of 64 characters. Values are strings with a maximum length of 512
	// characters, booleans, or numbers.
	Attributes map[string]VectorStoreFileBatchNewParamsFileAttributeUnion `json:"attributes,omitzero"`
	// Automatic chunking strategy for vector store files.
	ChunkingStrategy VectorStoreFileBatchNewParamsFileChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreFileBatchNewParamsFileAttributeUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileBatchNewParamsFileAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *VectorStoreFileBatchNewParamsFileAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileBatchNewParamsFileAttributeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreFileBatchNewParamsFileChunkingStrategyUnion struct {
	OfAuto       *VectorStoreFileBatchNewParamsFileChunkingStrategyAuto       `json:",omitzero,inline"`
	OfStatic     *VectorStoreFileBatchNewParamsFileChunkingStrategyStatic     `json:",omitzero,inline"`
	OfContextual *VectorStoreFileBatchNewParamsFileChunkingStrategyContextual `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileBatchNewParamsFileChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic, u.OfContextual)
}
func (u *VectorStoreFileBatchNewParamsFileChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileBatchNewParamsFileChunkingStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfAuto) {
		return u.OfAuto
	} else if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
	} else if !param.IsOmitted(u.OfContextual) {
		return u.OfContextual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileBatchNewParamsFileChunkingStrategyUnion) GetStatic() *VectorStoreFileBatchNewParamsFileChunkingStrategyStaticStatic {
	if vt := u.OfStatic; vt != nil {
		return &vt.Static
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileBatchNewParamsFileChunkingStrategyUnion) GetContextual() *VectorStoreFileBatchNewParamsFileChunkingStrategyContextualContextual {
	if vt := u.OfContextual; vt != nil {
		return &vt.Contextual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileBatchNewParamsFileChunkingStrategyUnion) GetType() *string {
	if vt := u.OfAuto; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContextual; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorStoreFileBatchNewParamsFileChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreFileBatchNewParamsFileChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreFileBatchNewParamsFileChunkingStrategyStatic]("static"),
		apijson.Discriminator[VectorStoreFileBatchNewParamsFileChunkingStrategyContextual]("contextual"),
	)
}

// Automatic chunking strategy for vector store files.
type VectorStoreFileBatchNewParamsFileChunkingStrategyAuto struct {
	// Any of "auto".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsFileChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsFileChunkingStrategyAuto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsFileChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileBatchNewParamsFileChunkingStrategyAuto](
		"type", "auto",
	)
}

// Static chunking strategy with configurable parameters.
//
// The property Static is required.
type VectorStoreFileBatchNewParamsFileChunkingStrategyStatic struct {
	// Configuration for static chunking strategy.
	Static VectorStoreFileBatchNewParamsFileChunkingStrategyStaticStatic `json:"static,omitzero" api:"required"`
	// Any of "static".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsFileChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsFileChunkingStrategyStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsFileChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileBatchNewParamsFileChunkingStrategyStatic](
		"type", "static",
	)
}

// Configuration for static chunking strategy.
type VectorStoreFileBatchNewParamsFileChunkingStrategyStaticStatic struct {
	ChunkOverlapTokens param.Opt[int64] `json:"chunk_overlap_tokens,omitzero"`
	MaxChunkSizeTokens param.Opt[int64] `json:"max_chunk_size_tokens,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsFileChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsFileChunkingStrategyStaticStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsFileChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contextual chunking strategy that uses an LLM to situate chunks within the
// document.
//
// The property Contextual is required.
type VectorStoreFileBatchNewParamsFileChunkingStrategyContextual struct {
	// Configuration for contextual chunking.
	Contextual VectorStoreFileBatchNewParamsFileChunkingStrategyContextualContextual `json:"contextual,omitzero" api:"required"`
	// Strategy type identifier.
	//
	// Any of "contextual".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsFileChunkingStrategyContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsFileChunkingStrategyContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsFileChunkingStrategyContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileBatchNewParamsFileChunkingStrategyContextual](
		"type", "contextual",
	)
}

// Configuration for contextual chunking.
type VectorStoreFileBatchNewParamsFileChunkingStrategyContextualContextual struct {
	// Maximum concurrent LLM calls. Falls back to config default if not provided.
	MaxConcurrency param.Opt[int64] `json:"max_concurrency,omitzero"`
	// LLM model for generating context. Falls back to
	// VectorStoresConfig.contextual_retrieval_params.model if not provided.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Timeout per LLM call in seconds. Falls back to config default if not provided.
	TimeoutSeconds param.Opt[int64] `json:"timeout_seconds,omitzero"`
	// Tokens to overlap between adjacent chunks. Must be less than
	// max_chunk_size_tokens.
	ChunkOverlapTokens param.Opt[int64] `json:"chunk_overlap_tokens,omitzero"`
	// Prompt template for contextual retrieval. Uses WHOLE_DOCUMENT and CHUNK_CONTENT
	// placeholders wrapped in double curly braces.
	ContextPrompt param.Opt[string] `json:"context_prompt,omitzero"`
	// Maximum tokens per chunk. Suggested ~700 to allow room for prepended context.
	MaxChunkSizeTokens param.Opt[int64] `json:"max_chunk_size_tokens,omitzero"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsFileChunkingStrategyContextualContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsFileChunkingStrategyContextualContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsFileChunkingStrategyContextualContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileBatchGetParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	paramObj
}

type VectorStoreFileBatchCancelParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	paramObj
}

type VectorStoreFileBatchListFilesParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	// Pagination cursor (after).
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor (before).
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by file status.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Maximum number of files to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by created_at: asc or desc.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
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
