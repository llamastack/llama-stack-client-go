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
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// VectorIoService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorIoService] method instead.
type VectorIoService struct {
	Options []option.RequestOption
}

// NewVectorIoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorIoService(opts ...option.RequestOption) (r VectorIoService) {
	r = VectorIoService{}
	r.Options = opts
	return
}

// Insert chunks into a vector database.
func (r *VectorIoService) Insert(ctx context.Context, body VectorIoInsertParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/vector-io/insert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Query chunks from a vector database.
func (r *VectorIoService) Query(ctx context.Context, body VectorIoQueryParams, opts ...option.RequestOption) (res *QueryChunksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vector-io/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from querying chunks in a vector database.
type QueryChunksResponse struct {
	Chunks []QueryChunksResponseChunk `json:"chunks,required"`
	Scores []float64                  `json:"scores,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chunks      respjson.Field
		Scores      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryChunksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk of content that can be inserted into a vector database.
type QueryChunksResponseChunk struct {
	ChunkID string `json:"chunk_id,required"`
	// A image content item
	Content QueryChunksResponseChunkContentUnion `json:"content,required"`
	// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
	// additional information about the chunk that will not be used in the context
	// during inference, but is required for backend functionality. The `ChunkMetadata`
	// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
	// expected to change after. Use `Chunk.metadata` for metadata that will be used in
	// the context during inference.
	ChunkMetadata QueryChunksResponseChunkChunkMetadata `json:"chunk_metadata,nullable"`
	Embedding     []float64                             `json:"embedding,nullable"`
	Metadata      map[string]any                        `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID       respjson.Field
		Content       respjson.Field
		ChunkMetadata respjson.Field
		Embedding     respjson.Field
		Metadata      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunk) RawJSON() string { return r.JSON.raw }
func (r *QueryChunksResponseChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QueryChunksResponseChunkContentUnion contains all possible properties and values
// from [string], [QueryChunksResponseChunkContentImageContentItemOutput],
// [QueryChunksResponseChunkContentTextContentItem],
// [[]QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfListImageContentItemOutputTextContentItem]
type QueryChunksResponseChunkContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion]
	// instead of an object.
	OfListImageContentItemOutputTextContentItem []QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion `json:",inline"`
	// This field is from variant
	// [QueryChunksResponseChunkContentImageContentItemOutput].
	Image QueryChunksResponseChunkContentImageContentItemOutputImage `json:"image"`
	Type  string                                                     `json:"type"`
	// This field is from variant [QueryChunksResponseChunkContentTextContentItem].
	Text string `json:"text"`
	JSON struct {
		OfString                                    respjson.Field
		OfListImageContentItemOutputTextContentItem respjson.Field
		Image                                       respjson.Field
		Type                                        respjson.Field
		Text                                        respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u QueryChunksResponseChunkContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryChunksResponseChunkContentUnion) AsImageContentItemOutput() (v QueryChunksResponseChunkContentImageContentItemOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryChunksResponseChunkContentUnion) AsTextContentItem() (v QueryChunksResponseChunkContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryChunksResponseChunkContentUnion) AsListImageContentItemOutputTextContentItem() (v []QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QueryChunksResponseChunkContentUnion) RawJSON() string { return u.JSON.raw }

func (r *QueryChunksResponseChunkContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A image content item
type QueryChunksResponseChunkContentImageContentItemOutput struct {
	// A URL or a base64 encoded string
	Image QueryChunksResponseChunkContentImageContentItemOutputImage `json:"image,required"`
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentImageContentItemOutput) RawJSON() string { return r.JSON.raw }
func (r *QueryChunksResponseChunkContentImageContentItemOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL or a base64 encoded string
type QueryChunksResponseChunkContentImageContentItemOutputImage struct {
	Data string `json:"data,nullable"`
	// A URL reference to external content.
	URL QueryChunksResponseChunkContentImageContentItemOutputImageURL `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentImageContentItemOutputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *QueryChunksResponseChunkContentImageContentItemOutputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
type QueryChunksResponseChunkContentImageContentItemOutputImageURL struct {
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentImageContentItemOutputImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *QueryChunksResponseChunkContentImageContentItemOutputImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type QueryChunksResponseChunkContentTextContentItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *QueryChunksResponseChunkContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion
// contains all possible properties and values from
// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage],
// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText].
//
// Use the
// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion struct {
	// This field is from variant
	// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage].
	Image QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImage `json:"image"`
	// Any of "image", "text".
	Type string `json:"type"`
	// This field is from variant
	// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText].
	Text string `json:"text"`
	JSON struct {
		Image respjson.Field
		Type  respjson.Field
		Text  respjson.Field
		raw   string
	} `json:"-"`
}

// anyQueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItem
// is implemented by each variant of
// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion]
// to add type safety for the return type of
// [QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion.AsAny]
type anyQueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItem interface {
	implQueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion()
}

func (QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage) implQueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion() {
}
func (QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText) implQueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion.AsAny().(type) {
//	case llamastackclient.QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage:
//	case llamastackclient.QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion) AsAny() anyQueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItem {
	switch u.Type {
	case "image":
		return u.AsImage()
	case "text":
		return u.AsText()
	}
	return nil
}

func (u QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion) AsImage() (v QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion) AsText() (v QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A image content item
type QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage struct {
	// A URL or a base64 encoded string
	Image QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImage `json:"image,required"`
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage) RawJSON() string {
	return r.JSON.raw
}
func (r *QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL or a base64 encoded string
type QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImage struct {
	Data string `json:"data,nullable"`
	// A URL reference to external content.
	URL QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImageURL `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImage) RawJSON() string {
	return r.JSON.raw
}
func (r *QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
type QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImageURL struct {
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *QueryChunksResponseChunkContentListImageContentItemOutputTextContentItemItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
// additional information about the chunk that will not be used in the context
// during inference, but is required for backend functionality. The `ChunkMetadata`
// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
// expected to change after. Use `Chunk.metadata` for metadata that will be used in
// the context during inference.
type QueryChunksResponseChunkChunkMetadata struct {
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
func (r QueryChunksResponseChunkChunkMetadata) RawJSON() string { return r.JSON.raw }
func (r *QueryChunksResponseChunkChunkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorIoInsertParams struct {
	Chunks        []VectorIoInsertParamsChunk `json:"chunks,omitzero,required"`
	VectorStoreID string                      `json:"vector_store_id,required"`
	TtlSeconds    param.Opt[int64]            `json:"ttl_seconds,omitzero"`
	paramObj
}

func (r VectorIoInsertParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk of content that can be inserted into a vector database.
//
// The properties ChunkID, Content are required.
type VectorIoInsertParamsChunk struct {
	ChunkID string `json:"chunk_id,required"`
	// A image content item
	Content VectorIoInsertParamsChunkContentUnion `json:"content,omitzero,required"`
	// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
	// additional information about the chunk that will not be used in the context
	// during inference, but is required for backend functionality. The `ChunkMetadata`
	// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
	// expected to change after. Use `Chunk.metadata` for metadata that will be used in
	// the context during inference.
	ChunkMetadata VectorIoInsertParamsChunkChunkMetadata `json:"chunk_metadata,omitzero"`
	Embedding     []float64                              `json:"embedding,omitzero"`
	Metadata      map[string]any                         `json:"metadata,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunk) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunk
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorIoInsertParamsChunkContentUnion struct {
	OfString                                   param.Opt[string]                                                                   `json:",omitzero,inline"`
	OfImageContentItemInput                    *VectorIoInsertParamsChunkContentImageContentItemInput                              `json:",omitzero,inline"`
	OfTextContentItem                          *VectorIoInsertParamsChunkContentTextContentItem                                    `json:",omitzero,inline"`
	OfListImageContentItemInputTextContentItem []VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VectorIoInsertParamsChunkContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfImageContentItemInput, u.OfTextContentItem, u.OfListImageContentItemInputTextContentItem)
}
func (u *VectorIoInsertParamsChunkContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorIoInsertParamsChunkContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItemInput) {
		return u.OfImageContentItemInput
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfListImageContentItemInputTextContentItem) {
		return &u.OfListImageContentItemInputTextContentItem
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoInsertParamsChunkContentUnion) GetImage() *VectorIoInsertParamsChunkContentImageContentItemInputImage {
	if vt := u.OfImageContentItemInput; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoInsertParamsChunkContentUnion) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoInsertParamsChunkContentUnion) GetType() *string {
	if vt := u.OfImageContentItemInput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The property Image is required.
type VectorIoInsertParamsChunkContentImageContentItemInput struct {
	// A URL or a base64 encoded string
	Image VectorIoInsertParamsChunkContentImageContentItemInputImage `json:"image,omitzero,required"`
	// Any of "image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentImageContentItemInput) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentImageContentItemInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentImageContentItemInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoInsertParamsChunkContentImageContentItemInput](
		"type", "image",
	)
}

// A URL or a base64 encoded string
type VectorIoInsertParamsChunkContentImageContentItemInputImage struct {
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL reference to external content.
	URL VectorIoInsertParamsChunkContentImageContentItemInputImageURL `json:"url,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentImageContentItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentImageContentItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentImageContentItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type VectorIoInsertParamsChunkContentImageContentItemInputImageURL struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentImageContentItemInputImageURL) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentImageContentItemInputImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentImageContentItemInputImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The property Text is required.
type VectorIoInsertParamsChunkContentTextContentItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentTextContentItem) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentTextContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoInsertParamsChunkContentTextContentItem](
		"type", "text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion struct {
	OfImage *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage `json:",omitzero,inline"`
	OfText  *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText  `json:",omitzero,inline"`
	paramUnion
}

func (u VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfText)
}
func (u *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion) asAny() any {
	if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion) GetImage() *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImage {
	if vt := u.OfImage; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion) GetType() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemUnion](
		"type",
		apijson.Discriminator[VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage]("image"),
		apijson.Discriminator[VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText]("text"),
	)
}

// A image content item
//
// The property Image is required.
type VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage struct {
	// A URL or a base64 encoded string
	Image VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImage `json:"image,omitzero,required"`
	// Any of "image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImage](
		"type", "image",
	)
}

// A URL or a base64 encoded string
type VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImage struct {
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL reference to external content.
	URL VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImageURL `json:"url,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImage) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImageURL struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImageURL) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The property Text is required.
type VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoInsertParamsChunkContentListImageContentItemInputTextContentItemItemText](
		"type", "text",
	)
}

// `ChunkMetadata` is backend metadata for a `Chunk` that is used to store
// additional information about the chunk that will not be used in the context
// during inference, but is required for backend functionality. The `ChunkMetadata`
// is set during chunk creation in `MemoryToolRuntimeImpl().insert()`and is not
// expected to change after. Use `Chunk.metadata` for metadata that will be used in
// the context during inference.
type VectorIoInsertParamsChunkChunkMetadata struct {
	ChunkEmbeddingDimension param.Opt[int64]  `json:"chunk_embedding_dimension,omitzero"`
	ChunkEmbeddingModel     param.Opt[string] `json:"chunk_embedding_model,omitzero"`
	ChunkID                 param.Opt[string] `json:"chunk_id,omitzero"`
	ChunkTokenizer          param.Opt[string] `json:"chunk_tokenizer,omitzero"`
	ChunkWindow             param.Opt[string] `json:"chunk_window,omitzero"`
	ContentTokenCount       param.Opt[int64]  `json:"content_token_count,omitzero"`
	CreatedTimestamp        param.Opt[int64]  `json:"created_timestamp,omitzero"`
	DocumentID              param.Opt[string] `json:"document_id,omitzero"`
	MetadataTokenCount      param.Opt[int64]  `json:"metadata_token_count,omitzero"`
	Source                  param.Opt[string] `json:"source,omitzero"`
	UpdatedTimestamp        param.Opt[int64]  `json:"updated_timestamp,omitzero"`
	paramObj
}

func (r VectorIoInsertParamsChunkChunkMetadata) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParamsChunkChunkMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParamsChunkChunkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorIoQueryParams struct {
	// A image content item
	Query         VectorIoQueryParamsQueryUnion `json:"query,omitzero,required"`
	VectorStoreID string                        `json:"vector_store_id,required"`
	Params        map[string]any                `json:"params,omitzero"`
	paramObj
}

func (r VectorIoQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorIoQueryParamsQueryUnion struct {
	OfString                                   param.Opt[string]                                                           `json:",omitzero,inline"`
	OfImageContentItemInput                    *VectorIoQueryParamsQueryImageContentItemInput                              `json:",omitzero,inline"`
	OfTextContentItem                          *VectorIoQueryParamsQueryTextContentItem                                    `json:",omitzero,inline"`
	OfListImageContentItemInputTextContentItem []VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u VectorIoQueryParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfImageContentItemInput, u.OfTextContentItem, u.OfListImageContentItemInputTextContentItem)
}
func (u *VectorIoQueryParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorIoQueryParamsQueryUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItemInput) {
		return u.OfImageContentItemInput
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfListImageContentItemInputTextContentItem) {
		return &u.OfListImageContentItemInputTextContentItem
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoQueryParamsQueryUnion) GetImage() *VectorIoQueryParamsQueryImageContentItemInputImage {
	if vt := u.OfImageContentItemInput; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoQueryParamsQueryUnion) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoQueryParamsQueryUnion) GetType() *string {
	if vt := u.OfImageContentItemInput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The property Image is required.
type VectorIoQueryParamsQueryImageContentItemInput struct {
	// A URL or a base64 encoded string
	Image VectorIoQueryParamsQueryImageContentItemInputImage `json:"image,omitzero,required"`
	// Any of "image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoQueryParamsQueryImageContentItemInput) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryImageContentItemInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryImageContentItemInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoQueryParamsQueryImageContentItemInput](
		"type", "image",
	)
}

// A URL or a base64 encoded string
type VectorIoQueryParamsQueryImageContentItemInputImage struct {
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL reference to external content.
	URL VectorIoQueryParamsQueryImageContentItemInputImageURL `json:"url,omitzero"`
	paramObj
}

func (r VectorIoQueryParamsQueryImageContentItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryImageContentItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryImageContentItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type VectorIoQueryParamsQueryImageContentItemInputImageURL struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r VectorIoQueryParamsQueryImageContentItemInputImageURL) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryImageContentItemInputImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryImageContentItemInputImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The property Text is required.
type VectorIoQueryParamsQueryTextContentItem struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoQueryParamsQueryTextContentItem) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryTextContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoQueryParamsQueryTextContentItem](
		"type", "text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion struct {
	OfImage *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage `json:",omitzero,inline"`
	OfText  *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText  `json:",omitzero,inline"`
	paramUnion
}

func (u VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfText)
}
func (u *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion) asAny() any {
	if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion) GetImage() *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImage {
	if vt := u.OfImage; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion) GetType() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemUnion](
		"type",
		apijson.Discriminator[VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage]("image"),
		apijson.Discriminator[VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText]("text"),
	)
}

// A image content item
//
// The property Image is required.
type VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage struct {
	// A URL or a base64 encoded string
	Image VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImage `json:"image,omitzero,required"`
	// Any of "image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImage](
		"type", "image",
	)
}

// A URL or a base64 encoded string
type VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImage struct {
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL reference to external content.
	URL VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImageURL `json:"url,omitzero"`
	paramObj
}

func (r VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImage) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImageURL struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImageURL) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The property Text is required.
type VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VectorIoQueryParamsQueryListImageContentItemInputTextContentItemItemText](
		"type", "text",
	)
}
