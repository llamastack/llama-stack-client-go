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

// VectorStoreService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorStoreService] method instead.
type VectorStoreService struct {
	Options     []option.RequestOption
	Files       VectorStoreFileService
	FileBatches VectorStoreFileBatchService
}

// NewVectorStoreService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorStoreService(opts ...option.RequestOption) (r VectorStoreService) {
	r = VectorStoreService{}
	r.Options = opts
	r.Files = NewVectorStoreFileService(opts...)
	r.FileBatches = NewVectorStoreFileBatchService(opts...)
	return
}

// Creates a vector store.
//
// Generate an OpenAI-compatible vector store with the given parameters.
func (r *VectorStoreService) New(ctx context.Context, body VectorStoreNewParams, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vector_stores"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a vector store.
func (r *VectorStoreService) Get(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a vector store.
func (r *VectorStoreService) Update(ctx context.Context, vectorStoreID string, body VectorStoreUpdateParams, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of vector stores.
func (r *VectorStoreService) List(ctx context.Context, query VectorStoreListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[VectorStore], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/vector_stores"
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

// Returns a list of vector stores.
func (r *VectorStoreService) ListAutoPaging(ctx context.Context, query VectorStoreListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[VectorStore] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a vector store.
func (r *VectorStoreService) Delete(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStoreDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Search for chunks in a vector store.
//
// Searches a vector store for relevant chunks based on a query and optional file
// attribute filters.
func (r *VectorStoreService) Search(ctx context.Context, vectorStoreID string, body VectorStoreSearchParams, opts ...option.RequestOption) (res *VectorStoreSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/search", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from listing vector stores.
type ListVectorStoresResponse struct {
	Data    []VectorStore `json:"data,required"`
	FirstID string        `json:"first_id,nullable"`
	HasMore bool          `json:"has_more"`
	LastID  string        `json:"last_id,nullable"`
	Object  string        `json:"object"`
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
func (r ListVectorStoresResponse) RawJSON() string { return r.JSON.raw }
func (r *ListVectorStoresResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI Vector Store object.
type VectorStore struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	// File processing status counts for a vector store.
	FileCounts   VectorStoreFileCounts `json:"file_counts,required"`
	ExpiresAfter map[string]any        `json:"expires_after,nullable"`
	ExpiresAt    int64                 `json:"expires_at,nullable"`
	LastActiveAt int64                 `json:"last_active_at,nullable"`
	Metadata     map[string]any        `json:"metadata"`
	Name         string                `json:"name,nullable"`
	Object       string                `json:"object"`
	Status       string                `json:"status"`
	UsageBytes   int64                 `json:"usage_bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		FileCounts   respjson.Field
		ExpiresAfter respjson.Field
		ExpiresAt    respjson.Field
		LastActiveAt respjson.Field
		Metadata     respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		Status       respjson.Field
		UsageBytes   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStore) RawJSON() string { return r.JSON.raw }
func (r *VectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File processing status counts for a vector store.
type VectorStoreFileCounts struct {
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
func (r VectorStoreFileCounts) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a vector store.
type VectorStoreDeleteResponse struct {
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted"`
	Object  string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated response from searching a vector store.
type VectorStoreSearchResponse struct {
	Data        []VectorStoreSearchResponseData `json:"data,required"`
	SearchQuery []string                        `json:"search_query,required"`
	HasMore     bool                            `json:"has_more"`
	NextPage    string                          `json:"next_page,nullable"`
	Object      string                          `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		SearchQuery respjson.Field
		HasMore     respjson.Field
		NextPage    respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from searching a vector store.
type VectorStoreSearchResponseData struct {
	Content    []VectorStoreSearchResponseDataContent                 `json:"content,required"`
	FileID     string                                                 `json:"file_id,required"`
	Filename   string                                                 `json:"filename,required"`
	Score      float64                                                `json:"score,required"`
	Attributes map[string]VectorStoreSearchResponseDataAttributeUnion `json:"attributes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Attributes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content item from a vector store file or search result.
type VectorStoreSearchResponseDataContent struct {
	Text string        `json:"text,required"`
	Type constant.Text `json:"type,required"`
	// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
	// additional information about the chunk that will not be used in the context
	// during inference, but is required for backend functionality. The `ChunkMetadata`
	// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
	// expected to change after. Use `Chunk.metadata` for metadata that will be used in
	// the context during inference.
	ChunkMetadata VectorStoreSearchResponseDataContentChunkMetadata `json:"chunk_metadata,nullable"`
	Embedding     []float64                                         `json:"embedding,nullable"`
	Metadata      map[string]any                                    `json:"metadata,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text          respjson.Field
		Type          respjson.Field
		ChunkMetadata respjson.Field
		Embedding     respjson.Field
		Metadata      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreSearchResponseDataContent) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreSearchResponseDataContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
// additional information about the chunk that will not be used in the context
// during inference, but is required for backend functionality. The `ChunkMetadata`
// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
// expected to change after. Use `Chunk.metadata` for metadata that will be used in
// the context during inference.
type VectorStoreSearchResponseDataContentChunkMetadata struct {
	ChunkEmbeddingDimension int64  `json:"chunk_embedding_dimension,nullable"`
	ChunkEmbeddingModel     string `json:"chunk_embedding_model,nullable"`
	ChunkID                 string `json:"chunk_id,nullable"`
	ChunkTokenizer          string `json:"chunk_tokenizer,nullable"`
	ChunkWindow             string `json:"chunk_window,nullable"`
	ContentTokenCount       int64  `json:"content_token_count,nullable"`
	CreatedTimestamp        int64  `json:"created_timestamp,nullable"`
	DocumentID              string `json:"document_id,nullable"`
	MetadataTokenCount      int64  `json:"metadata_token_count,nullable"`
	Source                  string `json:"source,nullable"`
	UpdatedTimestamp        int64  `json:"updated_timestamp,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkEmbeddingDimension respjson.Field
		ChunkEmbeddingModel     respjson.Field
		ChunkID                 respjson.Field
		ChunkTokenizer          respjson.Field
		ChunkWindow             respjson.Field
		ContentTokenCount       respjson.Field
		CreatedTimestamp        respjson.Field
		DocumentID              respjson.Field
		MetadataTokenCount      respjson.Field
		Source                  respjson.Field
		UpdatedTimestamp        respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreSearchResponseDataContentChunkMetadata) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreSearchResponseDataContentChunkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreSearchResponseDataAttributeUnion contains all possible properties and
// values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type VectorStoreSearchResponseDataAttributeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u VectorStoreSearchResponseDataAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreSearchResponseDataAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreSearchResponseDataAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreSearchResponseDataAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreSearchResponseDataAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreNewParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Automatic chunking strategy for vector store files.
	ChunkingStrategy VectorStoreNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	ExpiresAfter     map[string]any                            `json:"expires_after,omitzero"`
	FileIDs          []string                                  `json:"file_ids,omitzero"`
	Metadata         map[string]any                            `json:"metadata,omitzero"`
	paramObj
}

func (r VectorStoreNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreNewParamsChunkingStrategyUnion struct {
	OfAuto   *VectorStoreNewParamsChunkingStrategyAuto   `json:",omitzero,inline"`
	OfStatic *VectorStoreNewParamsChunkingStrategyStatic `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic)
}
func (u *VectorStoreNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreNewParamsChunkingStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfAuto) {
		return u.OfAuto
	} else if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreNewParamsChunkingStrategyUnion) GetStatic() *VectorStoreNewParamsChunkingStrategyStaticStatic {
	if vt := u.OfStatic; vt != nil {
		return &vt.Static
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreNewParamsChunkingStrategyUnion) GetType() *string {
	if vt := u.OfAuto; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorStoreNewParamsChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreNewParamsChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreNewParamsChunkingStrategyStatic]("static"),
	)
}

// Automatic chunking strategy for vector store files.
type VectorStoreNewParamsChunkingStrategyAuto struct {
	// Any of "auto".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreNewParamsChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParamsChunkingStrategyAuto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParamsChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreNewParamsChunkingStrategyAuto](
		"type", "auto",
	)
}

// Static chunking strategy with configurable parameters.
//
// The property Static is required.
type VectorStoreNewParamsChunkingStrategyStatic struct {
	// Configuration for static chunking strategy.
	Static VectorStoreNewParamsChunkingStrategyStaticStatic `json:"static,omitzero,required"`
	// Any of "static".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreNewParamsChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParamsChunkingStrategyStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParamsChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreNewParamsChunkingStrategyStatic](
		"type", "static",
	)
}

// Configuration for static chunking strategy.
type VectorStoreNewParamsChunkingStrategyStaticStatic struct {
	ChunkOverlapTokens param.Opt[int64] `json:"chunk_overlap_tokens,omitzero"`
	MaxChunkSizeTokens param.Opt[int64] `json:"max_chunk_size_tokens,omitzero"`
	paramObj
}

func (r VectorStoreNewParamsChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParamsChunkingStrategyStaticStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParamsChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreUpdateParams struct {
	Name         param.Opt[string] `json:"name,omitzero"`
	ExpiresAfter map[string]any    `json:"expires_after,omitzero"`
	Metadata     map[string]any    `json:"metadata,omitzero"`
	paramObj
}

func (r VectorStoreUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreListParams struct {
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Order  param.Opt[string] `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VectorStoreListParams]'s query parameters as `url.Values`.
func (r VectorStoreListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VectorStoreSearchParams struct {
	Query         VectorStoreSearchParamsQueryUnion `json:"query,omitzero,required"`
	MaxNumResults param.Opt[int64]                  `json:"max_num_results,omitzero"`
	RewriteQuery  param.Opt[bool]                   `json:"rewrite_query,omitzero"`
	SearchMode    param.Opt[string]                 `json:"search_mode,omitzero"`
	Filters       map[string]any                    `json:"filters,omitzero"`
	// Options for ranking and filtering search results.
	RankingOptions VectorStoreSearchParamsRankingOptions `json:"ranking_options,omitzero"`
	paramObj
}

func (r VectorStoreSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreSearchParamsQueryUnion struct {
	OfString     param.Opt[string] `json:",omitzero,inline"`
	OfListString []string          `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreSearchParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListString)
}
func (u *VectorStoreSearchParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreSearchParamsQueryUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListString) {
		return &u.OfListString
	}
	return nil
}

// Options for ranking and filtering search results.
type VectorStoreSearchParamsRankingOptions struct {
	Ranker         param.Opt[string]  `json:"ranker,omitzero"`
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	paramObj
}

func (r VectorStoreSearchParamsRankingOptions) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreSearchParamsRankingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreSearchParamsRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
