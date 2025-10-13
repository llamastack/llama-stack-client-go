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

// Create a vector store file batch. Generate an OpenAI-compatible vector store
// file batch for the given vector store.
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
	// List of vector store file objects in the batch
	Data []VectorStoreFile `json:"data,required"`
	// Whether there are more files available beyond this page
	HasMore bool `json:"has_more,required"`
	// Object type identifier, always "list"
	Object string `json:"object,required"`
	// (Optional) ID of the first file in the list for pagination
	FirstID string `json:"first_id"`
	// (Optional) ID of the last file in the list for pagination
	LastID string `json:"last_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		FirstID     respjson.Field
		LastID      respjson.Field
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
	// Unique identifier for the file batch
	ID string `json:"id,required"`
	// Timestamp when the file batch was created
	CreatedAt int64 `json:"created_at,required"`
	// File processing status counts for the batch
	FileCounts VectorStoreFileBatchesFileCounts `json:"file_counts,required"`
	// Object type identifier, always "vector_store.file_batch"
	Object string `json:"object,required"`
	// Current processing status of the file batch
	//
	// Any of "completed", "in_progress", "cancelled", "failed".
	Status VectorStoreFileBatchesStatus `json:"status,required"`
	// ID of the vector store containing the file batch
	VectorStoreID string `json:"vector_store_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		FileCounts    respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		VectorStoreID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileBatches) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileBatches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File processing status counts for the batch
type VectorStoreFileBatchesFileCounts struct {
	// Number of files that had their processing cancelled
	Cancelled int64 `json:"cancelled,required"`
	// Number of files that have been successfully processed
	Completed int64 `json:"completed,required"`
	// Number of files that failed to process
	Failed int64 `json:"failed,required"`
	// Number of files currently being processed
	InProgress int64 `json:"in_progress,required"`
	// Total number of files in the vector store
	Total int64 `json:"total,required"`
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

// Current processing status of the file batch
type VectorStoreFileBatchesStatus string

const (
	VectorStoreFileBatchesStatusCompleted  VectorStoreFileBatchesStatus = "completed"
	VectorStoreFileBatchesStatusInProgress VectorStoreFileBatchesStatus = "in_progress"
	VectorStoreFileBatchesStatusCancelled  VectorStoreFileBatchesStatus = "cancelled"
	VectorStoreFileBatchesStatusFailed     VectorStoreFileBatchesStatus = "failed"
)

type VectorStoreFileBatchNewParams struct {
	// A list of File IDs that the vector store should use
	FileIDs []string `json:"file_ids,omitzero,required"`
	// (Optional) Key-value attributes to store with the files
	Attributes map[string]VectorStoreFileBatchNewParamsAttributeUnion `json:"attributes,omitzero"`
	// (Optional) The chunking strategy used to chunk the file(s). Defaults to auto
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
type VectorStoreFileBatchNewParamsAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileBatchNewParamsAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreFileBatchNewParamsAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileBatchNewParamsAttributeUnion) asAny() any {
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

func NewVectorStoreFileBatchNewParamsChunkingStrategyAuto() VectorStoreFileBatchNewParamsChunkingStrategyAuto {
	return VectorStoreFileBatchNewParamsChunkingStrategyAuto{
		Type: "auto",
	}
}

// Automatic chunking strategy for vector store files.
//
// This struct has a constant value, construct it with
// [NewVectorStoreFileBatchNewParamsChunkingStrategyAuto].
type VectorStoreFileBatchNewParamsChunkingStrategyAuto struct {
	// Strategy type, always "auto" for automatic chunking
	Type constant.Auto `json:"type,required"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyAuto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Static chunking strategy with configurable parameters.
//
// The properties Static, Type are required.
type VectorStoreFileBatchNewParamsChunkingStrategyStatic struct {
	// Configuration parameters for the static chunking strategy
	Static VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic `json:"static,omitzero,required"`
	// Strategy type, always "static" for static chunking
	//
	// This field can be elided, and will marshal its zero value as "static".
	Type constant.Static `json:"type,required"`
	paramObj
}

func (r VectorStoreFileBatchNewParamsChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileBatchNewParamsChunkingStrategyStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileBatchNewParamsChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration parameters for the static chunking strategy
//
// The properties ChunkOverlapTokens, MaxChunkSizeTokens are required.
type VectorStoreFileBatchNewParamsChunkingStrategyStaticStatic struct {
	// Number of tokens to overlap between adjacent chunks
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens,required"`
	// Maximum number of tokens per chunk, must be between 100 and 4096
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens,required"`
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
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	// A cursor for use in pagination. `after` is an object ID that defines your place
	// in the list.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A cursor for use in pagination. `before` is an object ID that defines your place
	// in the list.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by file status. One of in_progress, completed, failed, cancelled.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
	// order and `desc` for descending order.
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
