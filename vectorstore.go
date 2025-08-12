// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
	Options []option.RequestOption
	Files   VectorStoreFileService
}

// NewVectorStoreService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorStoreService(opts ...option.RequestOption) (r VectorStoreService) {
	r = VectorStoreService{}
	r.Options = opts
	r.Files = NewVectorStoreFileService(opts...)
	return
}

// Creates a vector store.
func (r *VectorStoreService) New(ctx context.Context, body VectorStoreNewParams, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/vector_stores"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a vector store.
func (r *VectorStoreService) Get(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a vector store.
func (r *VectorStoreService) Update(ctx context.Context, vectorStoreID string, body VectorStoreUpdateParams, opts ...option.RequestOption) (res *VectorStore, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of vector stores.
func (r *VectorStoreService) List(ctx context.Context, query VectorStoreListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPagination[VectorStore], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/openai/v1/vector_stores"
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
func (r *VectorStoreService) ListAutoPaging(ctx context.Context, query VectorStoreListParams, opts ...option.RequestOption) *pagination.OpenAICursorPaginationAutoPager[VectorStore] {
	return pagination.NewOpenAICursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a vector store.
func (r *VectorStoreService) Delete(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStoreDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Search for chunks in a vector store. Searches a vector store for relevant chunks
// based on a query and optional file attribute filters.
func (r *VectorStoreService) Search(ctx context.Context, vectorStoreID string, body VectorStoreSearchParams, opts ...option.RequestOption) (res *VectorStoreSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/search", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from listing vector stores.
type ListVectorStoresResponse struct {
	// List of vector store objects
	Data []VectorStore `json:"data,required"`
	// Whether there are more vector stores available beyond this page
	HasMore bool `json:"has_more,required"`
	// Object type identifier, always "list"
	Object string `json:"object,required"`
	// (Optional) ID of the first vector store in the list for pagination
	FirstID string `json:"first_id"`
	// (Optional) ID of the last vector store in the list for pagination
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
func (r ListVectorStoresResponse) RawJSON() string { return r.JSON.raw }
func (r *ListVectorStoresResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI Vector Store object.
type VectorStore struct {
	// Unique identifier for the vector store
	ID string `json:"id,required"`
	// Timestamp when the vector store was created
	CreatedAt int64 `json:"created_at,required"`
	// File processing status counts for the vector store
	FileCounts VectorStoreFileCounts `json:"file_counts,required"`
	// Set of key-value pairs that can be attached to the vector store
	Metadata map[string]VectorStoreMetadataUnion `json:"metadata,required"`
	// Object type identifier, always "vector_store"
	Object string `json:"object,required"`
	// Current status of the vector store
	Status string `json:"status,required"`
	// Storage space used by the vector store in bytes
	UsageBytes int64 `json:"usage_bytes,required"`
	// (Optional) Expiration policy for the vector store
	ExpiresAfter map[string]VectorStoreExpiresAfterUnion `json:"expires_after"`
	// (Optional) Timestamp when the vector store will expire
	ExpiresAt int64 `json:"expires_at"`
	// (Optional) Timestamp of last activity on the vector store
	LastActiveAt int64 `json:"last_active_at"`
	// (Optional) Name of the vector store
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		FileCounts   respjson.Field
		Metadata     respjson.Field
		Object       respjson.Field
		Status       respjson.Field
		UsageBytes   respjson.Field
		ExpiresAfter respjson.Field
		ExpiresAt    respjson.Field
		LastActiveAt respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStore) RawJSON() string { return r.JSON.raw }
func (r *VectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File processing status counts for the vector store
type VectorStoreFileCounts struct {
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
func (r VectorStoreFileCounts) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreMetadataUnion struct {
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

func (u VectorStoreMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreExpiresAfterUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreExpiresAfterUnion struct {
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

func (u VectorStoreExpiresAfterUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreExpiresAfterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreExpiresAfterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreExpiresAfterUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreExpiresAfterUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreExpiresAfterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a vector store.
type VectorStoreDeleteResponse struct {
	// Unique identifier of the deleted vector store
	ID string `json:"id,required"`
	// Whether the deletion operation was successful
	Deleted bool `json:"deleted,required"`
	// Object type identifier for the deletion response
	Object string `json:"object,required"`
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
	// List of search result objects
	Data []VectorStoreSearchResponseData `json:"data,required"`
	// Whether there are more results available beyond this page
	HasMore bool `json:"has_more,required"`
	// Object type identifier for the search results page
	Object string `json:"object,required"`
	// The original search query that was executed
	SearchQuery string `json:"search_query,required"`
	// (Optional) Token for retrieving the next page of results
	NextPage string `json:"next_page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		SearchQuery respjson.Field
		NextPage    respjson.Field
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
	// List of content items matching the search query
	Content []VectorStoreSearchResponseDataContent `json:"content,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result
	Score float64 `json:"score,required"`
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]VectorStoreSearchResponseDataAttributeUnion `json:"attributes"`
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
	// The actual text content
	Text string `json:"text,required"`
	// Content type, currently only "text" is supported
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreSearchResponseDataContent) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreSearchResponseDataContent) UnmarshalJSON(data []byte) error {
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
	// The dimension of the embedding vectors (default: 384).
	EmbeddingDimension param.Opt[int64] `json:"embedding_dimension,omitzero"`
	// The embedding model to use for this vector store.
	EmbeddingModel param.Opt[string] `json:"embedding_model,omitzero"`
	// A name for the vector store.
	Name param.Opt[string] `json:"name,omitzero"`
	// The ID of the provider to use for this vector store.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
	// strategy.
	ChunkingStrategy map[string]VectorStoreNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	// The expiration policy for a vector store.
	ExpiresAfter map[string]VectorStoreNewParamsExpiresAfterUnion `json:"expires_after,omitzero"`
	// A list of File IDs that the vector store should use. Useful for tools like
	// `file_search` that can access files.
	FileIDs []string `json:"file_ids,omitzero"`
	// Set of 16 key-value pairs that can be attached to an object.
	Metadata map[string]VectorStoreNewParamsMetadataUnion `json:"metadata,omitzero"`
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
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreNewParamsChunkingStrategyUnion) asAny() any {
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
type VectorStoreNewParamsExpiresAfterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreNewParamsExpiresAfterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreNewParamsExpiresAfterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreNewParamsExpiresAfterUnion) asAny() any {
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
type VectorStoreNewParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreNewParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreNewParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreNewParamsMetadataUnion) asAny() any {
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

type VectorStoreUpdateParams struct {
	// The name of the vector store.
	Name param.Opt[string] `json:"name,omitzero"`
	// The expiration policy for a vector store.
	ExpiresAfter map[string]VectorStoreUpdateParamsExpiresAfterUnion `json:"expires_after,omitzero"`
	// Set of 16 key-value pairs that can be attached to an object.
	Metadata map[string]VectorStoreUpdateParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r VectorStoreUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreUpdateParamsExpiresAfterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreUpdateParamsExpiresAfterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreUpdateParamsExpiresAfterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreUpdateParamsExpiresAfterUnion) asAny() any {
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
type VectorStoreUpdateParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreUpdateParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreUpdateParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreUpdateParamsMetadataUnion) asAny() any {
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

type VectorStoreListParams struct {
	// A cursor for use in pagination. `after` is an object ID that defines your place
	// in the list.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A cursor for use in pagination. `before` is an object ID that defines your place
	// in the list.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
	// order and `desc` for descending order.
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
	// The query string or array for performing the search.
	Query VectorStoreSearchParamsQueryUnion `json:"query,omitzero,required"`
	// Maximum number of results to return (1 to 50 inclusive, default 10).
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	// Whether to rewrite the natural language query for vector search (default false)
	RewriteQuery param.Opt[bool] `json:"rewrite_query,omitzero"`
	// The search mode to use - "keyword", "vector", or "hybrid" (default "vector")
	SearchMode param.Opt[string] `json:"search_mode,omitzero"`
	// Filters based on file attributes to narrow the search results.
	Filters map[string]VectorStoreSearchParamsFilterUnion `json:"filters,omitzero"`
	// Ranking options for fine-tuning the search results.
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
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreSearchParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *VectorStoreSearchParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreSearchParamsQueryUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreSearchParamsFilterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreSearchParamsFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreSearchParamsFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreSearchParamsFilterUnion) asAny() any {
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

// Ranking options for fine-tuning the search results.
type VectorStoreSearchParamsRankingOptions struct {
	// (Optional) Name of the ranking algorithm to use
	Ranker param.Opt[string] `json:"ranker,omitzero"`
	// (Optional) Minimum relevance score threshold for results
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
