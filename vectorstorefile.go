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

// VectorStoreFileService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorStoreFileService] method instead.
type VectorStoreFileService struct {
	Options []option.RequestOption
}

// NewVectorStoreFileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorStoreFileService(opts ...option.RequestOption) (r VectorStoreFileService) {
	r = VectorStoreFileService{}
	r.Options = opts
	return
}

// Attach a file to a vector store (OpenAI-compatible).
func (r *VectorStoreFileService) New(ctx context.Context, vectorStoreID string, body VectorStoreFileNewParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = slices.Concat(r.Options, opts)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/files", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a vector store file (OpenAI-compatible).
func (r *VectorStoreFileService) Get(ctx context.Context, fileID string, query VectorStoreFileGetParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/files/%s", query.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a vector store file (OpenAI-compatible).
func (r *VectorStoreFileService) Update(ctx context.Context, fileID string, params VectorStoreFileUpdateParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/files/%s", params.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List files in a vector store (OpenAI-compatible).
func (r *VectorStoreFileService) List(ctx context.Context, vectorStoreID string, query VectorStoreFileListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[VectorStoreFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/files", vectorStoreID)
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

// List files in a vector store (OpenAI-compatible).
func (r *VectorStoreFileService) ListAutoPaging(ctx context.Context, vectorStoreID string, query VectorStoreFileListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[VectorStoreFile] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, vectorStoreID, query, opts...))
}

// Delete a vector store file (OpenAI-compatible).
func (r *VectorStoreFileService) Delete(ctx context.Context, fileID string, body VectorStoreFileDeleteParams, opts ...option.RequestOption) (res *VectorStoreFileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/files/%s", body.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve vector store file contents (OpenAI-compatible).
func (r *VectorStoreFileService) Content(ctx context.Context, fileID string, params VectorStoreFileContentParams, opts ...option.RequestOption) (res *VectorStoreFileContentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector_stores/%s/files/%s/content", params.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// OpenAI Vector Store File object.
type VectorStoreFile struct {
	ID string `json:"id" api:"required"`
	// Automatic chunking strategy for vector store files.
	ChunkingStrategy VectorStoreFileChunkingStrategyUnion `json:"chunking_strategy" api:"required"`
	CreatedAt        int64                                `json:"created_at" api:"required"`
	// Any of "completed", "in_progress", "cancelled", "failed".
	Status        VectorStoreFileStatus `json:"status" api:"required"`
	VectorStoreID string                `json:"vector_store_id" api:"required"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard. Keys are strings with a maximum
	// length of 64 characters. Values are strings with a maximum length of 512
	// characters, booleans, or numbers.
	Attributes map[string]VectorStoreFileAttributeUnion `json:"attributes"`
	// Error information for failed vector store file processing.
	LastError  VectorStoreFileLastError `json:"last_error" api:"nullable"`
	Object     string                   `json:"object"`
	UsageBytes int64                    `json:"usage_bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ChunkingStrategy respjson.Field
		CreatedAt        respjson.Field
		Status           respjson.Field
		VectorStoreID    respjson.Field
		Attributes       respjson.Field
		LastError        respjson.Field
		Object           respjson.Field
		UsageBytes       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFile) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreFileChunkingStrategyUnion contains all possible properties and values
// from [VectorStoreFileChunkingStrategyAuto],
// [VectorStoreFileChunkingStrategyStatic],
// [VectorStoreFileChunkingStrategyContextual].
//
// Use the [VectorStoreFileChunkingStrategyUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type VectorStoreFileChunkingStrategyUnion struct {
	// Any of "auto", "static", "contextual".
	Type string `json:"type"`
	// This field is from variant [VectorStoreFileChunkingStrategyStatic].
	Static VectorStoreFileChunkingStrategyStaticStatic `json:"static"`
	// This field is from variant [VectorStoreFileChunkingStrategyContextual].
	Contextual VectorStoreFileChunkingStrategyContextualContextual `json:"contextual"`
	JSON       struct {
		Type       respjson.Field
		Static     respjson.Field
		Contextual respjson.Field
		raw        string
	} `json:"-"`
}

// anyVectorStoreFileChunkingStrategy is implemented by each variant of
// [VectorStoreFileChunkingStrategyUnion] to add type safety for the return type of
// [VectorStoreFileChunkingStrategyUnion.AsAny]
type anyVectorStoreFileChunkingStrategy interface {
	implVectorStoreFileChunkingStrategyUnion()
}

func (VectorStoreFileChunkingStrategyAuto) implVectorStoreFileChunkingStrategyUnion()       {}
func (VectorStoreFileChunkingStrategyStatic) implVectorStoreFileChunkingStrategyUnion()     {}
func (VectorStoreFileChunkingStrategyContextual) implVectorStoreFileChunkingStrategyUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := VectorStoreFileChunkingStrategyUnion.AsAny().(type) {
//	case llamastackclient.VectorStoreFileChunkingStrategyAuto:
//	case llamastackclient.VectorStoreFileChunkingStrategyStatic:
//	case llamastackclient.VectorStoreFileChunkingStrategyContextual:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u VectorStoreFileChunkingStrategyUnion) AsAny() anyVectorStoreFileChunkingStrategy {
	switch u.Type {
	case "auto":
		return u.AsAuto()
	case "static":
		return u.AsStatic()
	case "contextual":
		return u.AsContextual()
	}
	return nil
}

func (u VectorStoreFileChunkingStrategyUnion) AsAuto() (v VectorStoreFileChunkingStrategyAuto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileChunkingStrategyUnion) AsStatic() (v VectorStoreFileChunkingStrategyStatic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileChunkingStrategyUnion) AsContextual() (v VectorStoreFileChunkingStrategyContextual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreFileChunkingStrategyUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreFileChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic chunking strategy for vector store files.
type VectorStoreFileChunkingStrategyAuto struct {
	// Any of "auto".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileChunkingStrategyAuto) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Static chunking strategy with configurable parameters.
type VectorStoreFileChunkingStrategyStatic struct {
	// Configuration for static chunking strategy.
	Static VectorStoreFileChunkingStrategyStaticStatic `json:"static" api:"required"`
	// Any of "static".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Static      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileChunkingStrategyStatic) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for static chunking strategy.
type VectorStoreFileChunkingStrategyStaticStatic struct {
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlapTokens respjson.Field
		MaxChunkSizeTokens respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileChunkingStrategyStaticStatic) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contextual chunking strategy that uses an LLM to situate chunks within the
// document.
type VectorStoreFileChunkingStrategyContextual struct {
	// Configuration for contextual chunking.
	Contextual VectorStoreFileChunkingStrategyContextualContextual `json:"contextual" api:"required"`
	// Strategy type identifier.
	//
	// Any of "contextual".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contextual  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileChunkingStrategyContextual) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileChunkingStrategyContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for contextual chunking.
type VectorStoreFileChunkingStrategyContextualContextual struct {
	// Tokens to overlap between adjacent chunks. Must be less than
	// max_chunk_size_tokens.
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
	// Prompt template for contextual retrieval. Uses WHOLE_DOCUMENT and CHUNK_CONTENT
	// placeholders wrapped in double curly braces.
	ContextPrompt string `json:"context_prompt"`
	// Maximum tokens per chunk. Suggested ~700 to allow room for prepended context.
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// Maximum concurrent LLM calls. Falls back to config default if not provided.
	MaxConcurrency int64 `json:"max_concurrency" api:"nullable"`
	// LLM model for generating context. Falls back to
	// VectorStoresConfig.contextual_retrieval_params.model if not provided.
	ModelID string `json:"model_id" api:"nullable"`
	// Timeout per LLM call in seconds. Falls back to config default if not provided.
	TimeoutSeconds int64 `json:"timeout_seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlapTokens respjson.Field
		ContextPrompt      respjson.Field
		MaxChunkSizeTokens respjson.Field
		MaxConcurrency     respjson.Field
		ModelID            respjson.Field
		TimeoutSeconds     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileChunkingStrategyContextualContextual) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileChunkingStrategyContextualContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileStatus string

const (
	VectorStoreFileStatusCompleted  VectorStoreFileStatus = "completed"
	VectorStoreFileStatusInProgress VectorStoreFileStatus = "in_progress"
	VectorStoreFileStatusCancelled  VectorStoreFileStatus = "cancelled"
	VectorStoreFileStatusFailed     VectorStoreFileStatus = "failed"
)

// VectorStoreFileAttributeUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type VectorStoreFileAttributeUnion struct {
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

func (u VectorStoreFileAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreFileAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreFileAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information for failed vector store file processing.
type VectorStoreFileLastError struct {
	// Any of "server_error", "rate_limit_exceeded".
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileLastError) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a vector store file.
type VectorStoreFileDeleteResponse struct {
	ID      string `json:"id" api:"required"`
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
func (r VectorStoreFileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the parsed content of a vector store file.
type VectorStoreFileContentResponse struct {
	Data     []VectorStoreFileContentResponseData `json:"data" api:"required"`
	HasMore  bool                                 `json:"has_more"`
	NextPage string                               `json:"next_page" api:"nullable"`
	// Any of "vector_store.file_content.page".
	Object VectorStoreFileContentResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextPage    respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileContentResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content item from a vector store file or search result.
type VectorStoreFileContentResponseData struct {
	Text string        `json:"text" api:"required"`
	Type constant.Text `json:"type" api:"required"`
	// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
	// additional information about the chunk that will not be used in the context
	// during inference, but is required for backend functionality. The `ChunkMetadata`
	// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
	// expected to change after. Use `Chunk.metadata` for metadata that will be used in
	// the context during inference.
	ChunkMetadata VectorStoreFileContentResponseDataChunkMetadata `json:"chunk_metadata" api:"nullable"`
	Embedding     []float64                                       `json:"embedding" api:"nullable"`
	Metadata      map[string]any                                  `json:"metadata" api:"nullable"`
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
func (r VectorStoreFileContentResponseData) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileContentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
// additional information about the chunk that will not be used in the context
// during inference, but is required for backend functionality. The `ChunkMetadata`
// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
// expected to change after. Use `Chunk.metadata` for metadata that will be used in
// the context during inference.
type VectorStoreFileContentResponseDataChunkMetadata struct {
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
func (r VectorStoreFileContentResponseDataChunkMetadata) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileContentResponseDataChunkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileContentResponseObject string

const (
	VectorStoreFileContentResponseObjectVectorStoreFileContentPage VectorStoreFileContentResponseObject = "vector_store.file_content.page"
)

type VectorStoreFileNewParams struct {
	// The ID of the file to attach.
	FileID string `json:"file_id" api:"required"`
	// Attributes to associate with the file.
	Attributes map[string]any `json:"attributes,omitzero"`
	// Strategy for chunking the file content.
	ChunkingStrategy VectorStoreFileNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	paramObj
}

func (r VectorStoreFileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreFileNewParamsChunkingStrategyUnion struct {
	OfAuto       *VectorStoreFileNewParamsChunkingStrategyAuto       `json:",omitzero,inline"`
	OfStatic     *VectorStoreFileNewParamsChunkingStrategyStatic     `json:",omitzero,inline"`
	OfContextual *VectorStoreFileNewParamsChunkingStrategyContextual `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic, u.OfContextual)
}
func (u *VectorStoreFileNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileNewParamsChunkingStrategyUnion) asAny() any {
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
func (u VectorStoreFileNewParamsChunkingStrategyUnion) GetStatic() *VectorStoreFileNewParamsChunkingStrategyStaticStatic {
	if vt := u.OfStatic; vt != nil {
		return &vt.Static
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileNewParamsChunkingStrategyUnion) GetContextual() *VectorStoreFileNewParamsChunkingStrategyContextualContextual {
	if vt := u.OfContextual; vt != nil {
		return &vt.Contextual
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorStoreFileNewParamsChunkingStrategyUnion) GetType() *string {
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
	apijson.RegisterUnion[VectorStoreFileNewParamsChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreFileNewParamsChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreFileNewParamsChunkingStrategyStatic]("static"),
		apijson.Discriminator[VectorStoreFileNewParamsChunkingStrategyContextual]("contextual"),
	)
}

// Automatic chunking strategy for vector store files.
type VectorStoreFileNewParamsChunkingStrategyAuto struct {
	// Any of "auto".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyAuto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileNewParamsChunkingStrategyAuto](
		"type", "auto",
	)
}

// Static chunking strategy with configurable parameters.
//
// The property Static is required.
type VectorStoreFileNewParamsChunkingStrategyStatic struct {
	// Configuration for static chunking strategy.
	Static VectorStoreFileNewParamsChunkingStrategyStaticStatic `json:"static,omitzero" api:"required"`
	// Any of "static".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileNewParamsChunkingStrategyStatic](
		"type", "static",
	)
}

// Configuration for static chunking strategy.
type VectorStoreFileNewParamsChunkingStrategyStaticStatic struct {
	ChunkOverlapTokens param.Opt[int64] `json:"chunk_overlap_tokens,omitzero"`
	MaxChunkSizeTokens param.Opt[int64] `json:"max_chunk_size_tokens,omitzero"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyStaticStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contextual chunking strategy that uses an LLM to situate chunks within the
// document.
//
// The property Contextual is required.
type VectorStoreFileNewParamsChunkingStrategyContextual struct {
	// Configuration for contextual chunking.
	Contextual VectorStoreFileNewParamsChunkingStrategyContextualContextual `json:"contextual,omitzero" api:"required"`
	// Strategy type identifier.
	//
	// Any of "contextual".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorStoreFileNewParamsChunkingStrategyContextual](
		"type", "contextual",
	)
}

// Configuration for contextual chunking.
type VectorStoreFileNewParamsChunkingStrategyContextualContextual struct {
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

func (r VectorStoreFileNewParamsChunkingStrategyContextualContextual) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyContextualContextual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyContextualContextual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileGetParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	paramObj
}

type VectorStoreFileUpdateParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	// The new attributes for the file.
	Attributes map[string]any `json:"attributes,omitzero" api:"required"`
	paramObj
}

func (r VectorStoreFileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileListParams struct {
	// Pagination cursor (after).
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Pagination cursor (before).
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of files to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by created_at: asc or desc.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// Filter by file status.
	//
	// Any of "completed", "in_progress", "cancelled", "failed".
	Filter VectorStoreFileListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VectorStoreFileListParams]'s query parameters as
// `url.Values`.
func (r VectorStoreFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by file status.
type VectorStoreFileListParamsFilter string

const (
	VectorStoreFileListParamsFilterCompleted  VectorStoreFileListParamsFilter = "completed"
	VectorStoreFileListParamsFilterInProgress VectorStoreFileListParamsFilter = "in_progress"
	VectorStoreFileListParamsFilterCancelled  VectorStoreFileListParamsFilter = "cancelled"
	VectorStoreFileListParamsFilterFailed     VectorStoreFileListParamsFilter = "failed"
)

type VectorStoreFileDeleteParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	paramObj
}

type VectorStoreFileContentParams struct {
	// The vector store identifier.
	VectorStoreID string `path:"vector_store_id" api:"required" json:"-"`
	// Include embedding vectors.
	IncludeEmbeddings param.Opt[bool] `query:"include_embeddings,omitzero" json:"-"`
	// Include chunk metadata.
	IncludeMetadata param.Opt[bool] `query:"include_metadata,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VectorStoreFileContentParams]'s query parameters as
// `url.Values`.
func (r VectorStoreFileContentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
