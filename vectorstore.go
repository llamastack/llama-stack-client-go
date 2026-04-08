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

// Create a vector store (OpenAI-compatible).
func (r *VectorStoreService) New(ctx context.Context, body VectorStoreNewParams, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vector_stores"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a vector store (OpenAI-compatible).
func (r *VectorStoreService) Get(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a vector store (OpenAI-compatible).
func (r *VectorStoreService) Update(ctx context.Context, vectorStoreID string, body VectorStoreUpdateParams, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List vector stores (OpenAI-compatible).
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

// List vector stores (OpenAI-compatible).
func (r *VectorStoreService) ListAutoPaging(ctx context.Context, query VectorStoreListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[VectorStore] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a vector store (OpenAI-compatible).
func (r *VectorStoreService) Delete(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStoreDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Search a vector store (OpenAI-compatible).
func (r *VectorStoreService) Search(ctx context.Context, vectorStoreID string, body VectorStoreSearchParams, opts ...option.RequestOption) (res *VectorStoreSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/vector_stores/%s/search", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response from listing vector stores.
type ListVectorStoresResponse struct {
	Data    []VectorStore `json:"data" api:"required"`
	FirstID string        `json:"first_id" api:"required"`
	HasMore bool          `json:"has_more" api:"required"`
	LastID  string        `json:"last_id" api:"required"`
	// Any of "list".
	Object ListVectorStoresResponseObject `json:"object"`
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

type ListVectorStoresResponseObject string

const (
	ListVectorStoresResponseObjectList ListVectorStoresResponseObject = "list"
)

// OpenAI Vector Store object.
type VectorStore struct {
	ID        string `json:"id" api:"required"`
	CreatedAt int64  `json:"created_at" api:"required"`
	// File processing status counts for a vector store.
	FileCounts VectorStoreFileCounts `json:"file_counts" api:"required"`
	// Any of "expired", "in_progress", "completed".
	Status VectorStoreStatus `json:"status" api:"required"`
	// Expiration policy for a vector store.
	ExpiresAfter VectorStoreExpiresAfter `json:"expires_after" api:"nullable"`
	ExpiresAt    int64                   `json:"expires_at" api:"nullable"`
	LastActiveAt int64                   `json:"last_active_at" api:"nullable"`
	Metadata     map[string]any          `json:"metadata" api:"nullable"`
	Name         string                  `json:"name"`
	// Any of "vector_store".
	Object     VectorStoreObject `json:"object"`
	UsageBytes int64             `json:"usage_bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		FileCounts   respjson.Field
		Status       respjson.Field
		ExpiresAfter respjson.Field
		ExpiresAt    respjson.Field
		LastActiveAt respjson.Field
		Metadata     respjson.Field
		Name         respjson.Field
		Object       respjson.Field
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
func (r VectorStoreFileCounts) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreStatus string

const (
	VectorStoreStatusExpired    VectorStoreStatus = "expired"
	VectorStoreStatusInProgress VectorStoreStatus = "in_progress"
	VectorStoreStatusCompleted  VectorStoreStatus = "completed"
)

// Expiration policy for a vector store.
type VectorStoreExpiresAfter struct {
	// Anchor timestamp after which the expiration policy applies.
	//
	// Any of "last_active_at".
	Anchor string `json:"anchor" api:"required"`
	// The number of days after the anchor time that the vector store will expire.
	Days int64 `json:"days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Anchor      respjson.Field
		Days        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreExpiresAfter) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreExpiresAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreObject string

const (
	VectorStoreObjectVectorStore VectorStoreObject = "vector_store"
)

// Response from deleting a vector store.
type VectorStoreDeleteResponse struct {
	ID      string `json:"id" api:"required"`
	Deleted bool   `json:"deleted" api:"required"`
	// Any of "vector_store.deleted".
	Object VectorStoreDeleteResponseObject `json:"object"`
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

type VectorStoreDeleteResponseObject string

const (
	VectorStoreDeleteResponseObjectVectorStoreDeleted VectorStoreDeleteResponseObject = "vector_store.deleted"
)

// Paginated response from searching a vector store.
type VectorStoreSearchResponse struct {
	Data        []VectorStoreSearchResponseData `json:"data" api:"required"`
	HasMore     bool                            `json:"has_more" api:"required"`
	SearchQuery []string                        `json:"search_query" api:"required"`
	NextPage    string                          `json:"next_page" api:"nullable"`
	// Any of "vector_store.search_results.page".
	Object VectorStoreSearchResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		SearchQuery respjson.Field
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
	Content    []VectorStoreSearchResponseDataContent                 `json:"content" api:"required"`
	FileID     string                                                 `json:"file_id" api:"required"`
	Filename   string                                                 `json:"filename" api:"required"`
	Score      float64                                                `json:"score" api:"required"`
	Attributes map[string]VectorStoreSearchResponseDataAttributeUnion `json:"attributes" api:"nullable"`
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
	Text string `json:"text" api:"required"`
	// Any of "text".
	Type string `json:"type" api:"required"`
	// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
	// additional information about the chunk that will not be used in the context
	// during inference, but is required for backend functionality. The `ChunkMetadata`
	// is set during chunk creation in `FileSearchToolRuntimeImpl().insert()`and is not
	// expected to change after. Use `Chunk.metadata` for metadata that will be used in
	// the context during inference.
	ChunkMetadata VectorStoreSearchResponseDataContentChunkMetadata `json:"chunk_metadata" api:"nullable"`
	Embedding     []float64                                         `json:"embedding" api:"nullable"`
	Metadata      map[string]any                                    `json:"metadata" api:"nullable"`
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
// is set during chunk creation in `FileSearchToolRuntimeImpl().insert()`and is not
// expected to change after. Use `Chunk.metadata` for metadata that will be used in
// the context during inference.
type VectorStoreSearchResponseDataContentChunkMetadata struct {
	ChunkID            string `json:"chunk_id" api:"nullable"`
	ChunkTokenizer     string `json:"chunk_tokenizer" api:"nullable"`
	ChunkWindow        string `json:"chunk_window" api:"nullable"`
	ContentTokenCount  int64  `json:"content_token_count" api:"nullable"`
	CreatedTimestamp   int64  `json:"created_timestamp" api:"nullable"`
	DocumentID         string `json:"document_id" api:"nullable"`
	MetadataTokenCount int64  `json:"metadata_token_count" api:"nullable"`
	Source             string `json:"source" api:"nullable"`
	UpdatedTimestamp   int64  `json:"updated_timestamp" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID            respjson.Field
		ChunkTokenizer     respjson.Field
		ChunkWindow        respjson.Field
		ContentTokenCount  respjson.Field
		CreatedTimestamp   respjson.Field
		DocumentID         respjson.Field
		MetadataTokenCount respjson.Field
		Source             respjson.Field
		UpdatedTimestamp   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
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

type VectorStoreSearchResponseObject string

const (
	VectorStoreSearchResponseObjectVectorStoreSearchResultsPage VectorStoreSearchResponseObject = "vector_store.search_results.page"
)

type VectorStoreNewParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	// Automatic chunking strategy for vector store files.
	ChunkingStrategy VectorStoreNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	// Expiration policy for a vector store.
	ExpiresAfter VectorStoreNewParamsExpiresAfter `json:"expires_after,omitzero"`
	FileIDs      []string                         `json:"file_ids,omitzero"`
	Metadata     map[string]any                   `json:"metadata,omitzero"`
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
	OfAuto       *VectorStoreNewParamsChunkingStrategyAuto       `json:",omitzero,inline"`
	OfStatic     *VectorStoreNewParamsChunkingStrategyStatic     `json:",omitzero,inline"`
	OfContextual *VectorStoreNewParamsChunkingStrategyContextual `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic, u.OfContextual)
}
func (u *VectorStoreNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreNewParamsChunkingStrategyUnion) asAny() any {
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
func (u VectorStoreNewParamsChunkingStrategyUnion) GetStatic() *VectorStoreNewParamsChunkingStrategyStaticStatic {
	if vt := u.OfStatic; vt != nil {
		return &vt.Static
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreNewParamsChunkingStrategyUnion) GetContextual() *VectorStoreNewParamsChunkingStrategyContextualContextual {
	if vt := u.OfContextual; vt != nil {
		return &vt.Contextual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreNewParamsChunkingStrategyUnion) GetType() *string {
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
	apijson.RegisterUnion[VectorStoreNewParamsChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreNewParamsChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreNewParamsChunkingStrategyStatic]("static"),
		apijson.Discriminator[VectorStoreNewParamsChunkingStrategyContextual]("contextual"),
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
	Static VectorStoreNewParamsChunkingStrategyStaticStatic `json:"static,omitzero" api:"required"`
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

// Contextual chunking strategy that uses an LLM to situate chunks within the
// document.
//
// The property Contextual is required.
type VectorStoreNewParamsChunkingStrategyContextual struct {
	// Configuration for contextual chunking.
	Contextual VectorStoreNewParamsChunkingStrategyContextualContextual `json:"contextual,omitzero" api:"required"`
	// Strategy type identifier.
	//
	// Any of "contextual".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreNewParamsChunkingStrategyContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParamsChunkingStrategyContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParamsChunkingStrategyContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreNewParamsChunkingStrategyContextual](
		"type", "contextual",
	)
}

// Configuration for contextual chunking.
type VectorStoreNewParamsChunkingStrategyContextualContextual struct {
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

func (r VectorStoreNewParamsChunkingStrategyContextualContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParamsChunkingStrategyContextualContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParamsChunkingStrategyContextualContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Expiration policy for a vector store.
//
// The properties Anchor, Days are required.
type VectorStoreNewParamsExpiresAfter struct {
	// Anchor timestamp after which the expiration policy applies.
	//
	// Any of "last_active_at".
	Anchor string `json:"anchor,omitzero" api:"required"`
	// The number of days after the anchor time that the vector store will expire.
	Days int64 `json:"days" api:"required"`
	paramObj
}

func (r VectorStoreNewParamsExpiresAfter) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreNewParamsExpiresAfter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreNewParamsExpiresAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreNewParamsExpiresAfter](
		"anchor", "last_active_at",
	)
}

type VectorStoreUpdateParams struct {
	// The new name for the vector store.
	Name param.Opt[string] `json:"name,omitzero"`
	// Expiration policy for a vector store.
	ExpiresAfter VectorStoreUpdateParamsExpiresAfter `json:"expires_after,omitzero"`
	// Metadata to associate with the vector store.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r VectorStoreUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Expiration policy for a vector store.
//
// The properties Anchor, Days are required.
type VectorStoreUpdateParamsExpiresAfter struct {
	// Anchor timestamp after which the expiration policy applies.
	//
	// Any of "last_active_at".
	Anchor string `json:"anchor,omitzero" api:"required"`
	// The number of days after the anchor time that the vector store will expire.
	Days int64 `json:"days" api:"required"`
	paramObj
}

func (r VectorStoreUpdateParamsExpiresAfter) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreUpdateParamsExpiresAfter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreUpdateParamsExpiresAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreUpdateParamsExpiresAfter](
		"anchor", "last_active_at",
	)
}

type VectorStoreListParams struct {
	// Pagination cursor (after).
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor (before).
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of vector stores to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by created_at: asc or desc.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
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
	// The search query string or list of query strings.
	Query VectorStoreSearchParamsQueryUnion `json:"query,omitzero" api:"required"`
	// The search mode to use (e.g., 'vector', 'keyword').
	SearchMode param.Opt[string] `json:"search_mode,omitzero"`
	// Maximum number of results to return.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	// Whether to rewrite the query for better results.
	RewriteQuery param.Opt[bool] `json:"rewrite_query,omitzero"`
	// Filters to apply to the search.
	Filters map[string]any `json:"filters,omitzero"`
	// Options for ranking and filtering search results.
	//
	// This class configures how search results are ranked and filtered. You can use
	// algorithm-based rerankers (weighted, RRF) or neural rerankers. Defaults from
	// VectorStoresConfig are used when parameters are not provided.
	//
	// Examples: # Weighted ranker with custom alpha
	// SearchRankingOptions(ranker="weighted", alpha=0.7)
	//
	//	# RRF ranker with custom impact factor
	//	SearchRankingOptions(ranker="rrf", impact_factor=50.0)
	//
	//	# Use config defaults (just specify ranker type)
	//	SearchRankingOptions(ranker="weighted")  # Uses alpha from VectorStoresConfig
	//
	//	# Score threshold filtering
	//	SearchRankingOptions(ranker="weighted", score_threshold=0.5)
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
//
// This class configures how search results are ranked and filtered. You can use
// algorithm-based rerankers (weighted, RRF) or neural rerankers. Defaults from
// VectorStoresConfig are used when parameters are not provided.
//
// Examples: # Weighted ranker with custom alpha
// SearchRankingOptions(ranker="weighted", alpha=0.7)
//
//	# RRF ranker with custom impact factor
//	SearchRankingOptions(ranker="rrf", impact_factor=50.0)
//
//	# Use config defaults (just specify ranker type)
//	SearchRankingOptions(ranker="weighted")  # Uses alpha from VectorStoresConfig
//
//	# Score threshold filtering
//	SearchRankingOptions(ranker="weighted", score_threshold=0.5)
type VectorStoreSearchParamsRankingOptions struct {
	// Weight factor for weighted ranker
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Impact factor for RRF algorithm
	ImpactFactor param.Opt[float64] `json:"impact_factor,omitzero"`
	// Model identifier for neural reranker
	Model          param.Opt[string]  `json:"model,omitzero"`
	Ranker         param.Opt[string]  `json:"ranker,omitzero"`
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	// Weights for combining vector, keyword, and neural scores. Keys: 'vector',
	// 'keyword', 'neural'
	Weights map[string]float64 `json:"weights,omitzero"`
	paramObj
}

func (r VectorStoreSearchParamsRankingOptions) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreSearchParamsRankingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreSearchParamsRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
