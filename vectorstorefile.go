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

// Attach a file to a vector store.
func (r *VectorStoreFileService) New(ctx context.Context, vectorStoreID string, body VectorStoreFileNewParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a vector store file.
func (r *VectorStoreFileService) Get(ctx context.Context, fileID string, query VectorStoreFileGetParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = append(r.Options[:], opts...)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s", query.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a vector store file.
func (r *VectorStoreFileService) Update(ctx context.Context, fileID string, params VectorStoreFileUpdateParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = append(r.Options[:], opts...)
	if params.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s", params.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List files in a vector store.
func (r *VectorStoreFileService) List(ctx context.Context, vectorStoreID string, query VectorStoreFileListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[VectorStoreFile], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files", vectorStoreID)
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

// List files in a vector store.
func (r *VectorStoreFileService) ListAutoPaging(ctx context.Context, vectorStoreID string, query VectorStoreFileListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[VectorStoreFile] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, vectorStoreID, query, opts...))
}

// Delete a vector store file.
func (r *VectorStoreFileService) Delete(ctx context.Context, fileID string, body VectorStoreFileDeleteParams, opts ...option.RequestOption) (res *VectorStoreFileDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s", body.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieves the contents of a vector store file.
func (r *VectorStoreFileService) Content(ctx context.Context, fileID string, query VectorStoreFileContentParams, opts ...option.RequestOption) (res *VectorStoreFileContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s/content", query.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// OpenAI Vector Store File object.
type VectorStoreFile struct {
	// Unique identifier for the file
	ID string `json:"id,required"`
	// Key-value attributes associated with the file
	Attributes map[string]VectorStoreFileAttributeUnion `json:"attributes,required"`
	// Strategy used for splitting the file into chunks
	ChunkingStrategy VectorStoreFileChunkingStrategyUnion `json:"chunking_strategy,required"`
	// Timestamp when the file was added to the vector store
	CreatedAt int64 `json:"created_at,required"`
	// Object type identifier, always "vector_store.file"
	Object string `json:"object,required"`
	// Current processing status of the file
	//
	// Any of "completed", "in_progress", "cancelled", "failed".
	Status VectorStoreFileStatus `json:"status,required"`
	// Storage space used by this file in bytes
	UsageBytes int64 `json:"usage_bytes,required"`
	// ID of the vector store containing this file
	VectorStoreID string `json:"vector_store_id,required"`
	// (Optional) Error information if file processing failed
	LastError VectorStoreFileLastError `json:"last_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Attributes       respjson.Field
		ChunkingStrategy respjson.Field
		CreatedAt        respjson.Field
		Object           respjson.Field
		Status           respjson.Field
		UsageBytes       respjson.Field
		VectorStoreID    respjson.Field
		LastError        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFile) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreFileAttributeUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreFileAttributeUnion struct {
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

func (u VectorStoreFileAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreFileAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreFileAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreFileChunkingStrategyUnion contains all possible properties and values
// from [VectorStoreFileChunkingStrategyAuto],
// [VectorStoreFileChunkingStrategyStatic].
//
// Use the [VectorStoreFileChunkingStrategyUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type VectorStoreFileChunkingStrategyUnion struct {
	// Any of "auto", "static".
	Type string `json:"type"`
	// This field is from variant [VectorStoreFileChunkingStrategyStatic].
	Static VectorStoreFileChunkingStrategyStaticStatic `json:"static"`
	JSON   struct {
		Type   respjson.Field
		Static respjson.Field
		raw    string
	} `json:"-"`
}

// anyVectorStoreFileChunkingStrategy is implemented by each variant of
// [VectorStoreFileChunkingStrategyUnion] to add type safety for the return type of
// [VectorStoreFileChunkingStrategyUnion.AsAny]
type anyVectorStoreFileChunkingStrategy interface {
	implVectorStoreFileChunkingStrategyUnion()
}

func (VectorStoreFileChunkingStrategyAuto) implVectorStoreFileChunkingStrategyUnion()   {}
func (VectorStoreFileChunkingStrategyStatic) implVectorStoreFileChunkingStrategyUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := VectorStoreFileChunkingStrategyUnion.AsAny().(type) {
//	case llamastackclient.VectorStoreFileChunkingStrategyAuto:
//	case llamastackclient.VectorStoreFileChunkingStrategyStatic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u VectorStoreFileChunkingStrategyUnion) AsAny() anyVectorStoreFileChunkingStrategy {
	switch u.Type {
	case "auto":
		return u.AsAuto()
	case "static":
		return u.AsStatic()
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

// Returns the unmodified JSON received from the API
func (u VectorStoreFileChunkingStrategyUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreFileChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic chunking strategy for vector store files.
type VectorStoreFileChunkingStrategyAuto struct {
	// Strategy type, always "auto" for automatic chunking
	Type constant.Auto `json:"type,required"`
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
	// Configuration parameters for the static chunking strategy
	Static VectorStoreFileChunkingStrategyStaticStatic `json:"static,required"`
	// Strategy type, always "static" for static chunking
	Type constant.Static `json:"type,required"`
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

// Configuration parameters for the static chunking strategy
type VectorStoreFileChunkingStrategyStaticStatic struct {
	// Number of tokens to overlap between adjacent chunks
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens,required"`
	// Maximum number of tokens per chunk, must be between 100 and 4096
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens,required"`
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

// Current processing status of the file
type VectorStoreFileStatus string

const (
	VectorStoreFileStatusCompleted  VectorStoreFileStatus = "completed"
	VectorStoreFileStatusInProgress VectorStoreFileStatus = "in_progress"
	VectorStoreFileStatusCancelled  VectorStoreFileStatus = "cancelled"
	VectorStoreFileStatusFailed     VectorStoreFileStatus = "failed"
)

// (Optional) Error information if file processing failed
type VectorStoreFileLastError struct {
	// Error code indicating the type of failure
	//
	// Any of "server_error", "rate_limit_exceeded".
	Code VectorStoreFileLastErrorCode `json:"code,required"`
	// Human-readable error message describing the failure
	Message string `json:"message,required"`
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

// Error code indicating the type of failure
type VectorStoreFileLastErrorCode string

const (
	VectorStoreFileLastErrorCodeServerError       VectorStoreFileLastErrorCode = "server_error"
	VectorStoreFileLastErrorCodeRateLimitExceeded VectorStoreFileLastErrorCode = "rate_limit_exceeded"
)

// Response from deleting a vector store file.
type VectorStoreFileDeleteResponse struct {
	// Unique identifier of the deleted file
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
func (r VectorStoreFileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from retrieving the contents of a vector store file.
type VectorStoreFileContentResponse struct {
	// Key-value attributes associated with the file
	Attributes map[string]VectorStoreFileContentResponseAttributeUnion `json:"attributes,required"`
	// List of content items from the file
	Content []VectorStoreFileContentResponseContent `json:"content,required"`
	// Unique identifier for the file
	FileID string `json:"file_id,required"`
	// Name of the file
	Filename string `json:"filename,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Content     respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileContentResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreFileContentResponseAttributeUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreFileContentResponseAttributeUnion struct {
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

func (u VectorStoreFileContentResponseAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileContentResponseAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileContentResponseAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileContentResponseAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreFileContentResponseAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreFileContentResponseAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content item from a vector store file or search result.
type VectorStoreFileContentResponseContent struct {
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
func (r VectorStoreFileContentResponseContent) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileContentResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileNewParams struct {
	// The ID of the file to attach to the vector store.
	FileID string `json:"file_id,required"`
	// The key-value attributes stored with the file, which can be used for filtering.
	Attributes map[string]VectorStoreFileNewParamsAttributeUnion `json:"attributes,omitzero"`
	// The chunking strategy to use for the file.
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
type VectorStoreFileNewParamsAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileNewParamsAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreFileNewParamsAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileNewParamsAttributeUnion) asAny() any {
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
type VectorStoreFileNewParamsChunkingStrategyUnion struct {
	OfAuto   *VectorStoreFileNewParamsChunkingStrategyAuto   `json:",omitzero,inline"`
	OfStatic *VectorStoreFileNewParamsChunkingStrategyStatic `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic)
}
func (u *VectorStoreFileNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileNewParamsChunkingStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfAuto) {
		return u.OfAuto
	} else if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
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
func (u VectorStoreFileNewParamsChunkingStrategyUnion) GetType() *string {
	if vt := u.OfAuto; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VectorStoreFileNewParamsChunkingStrategyUnion](
		"type",
		apijson.Discriminator[VectorStoreFileNewParamsChunkingStrategyAuto]("auto"),
		apijson.Discriminator[VectorStoreFileNewParamsChunkingStrategyStatic]("static"),
	)
}

func NewVectorStoreFileNewParamsChunkingStrategyAuto() VectorStoreFileNewParamsChunkingStrategyAuto {
	return VectorStoreFileNewParamsChunkingStrategyAuto{
		Type: "auto",
	}
}

// Automatic chunking strategy for vector store files.
//
// This struct has a constant value, construct it with
// [NewVectorStoreFileNewParamsChunkingStrategyAuto].
type VectorStoreFileNewParamsChunkingStrategyAuto struct {
	// Strategy type, always "auto" for automatic chunking
	Type constant.Auto `json:"type,required"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyAuto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Static chunking strategy with configurable parameters.
//
// The properties Static, Type are required.
type VectorStoreFileNewParamsChunkingStrategyStatic struct {
	// Configuration parameters for the static chunking strategy
	Static VectorStoreFileNewParamsChunkingStrategyStaticStatic `json:"static,omitzero,required"`
	// Strategy type, always "static" for static chunking
	//
	// This field can be elided, and will marshal its zero value as "static".
	Type constant.Static `json:"type,required"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration parameters for the static chunking strategy
//
// The properties ChunkOverlapTokens, MaxChunkSizeTokens are required.
type VectorStoreFileNewParamsChunkingStrategyStaticStatic struct {
	// Number of tokens to overlap between adjacent chunks
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens,required"`
	// Maximum number of tokens per chunk, must be between 100 and 4096
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens,required"`
	paramObj
}

func (r VectorStoreFileNewParamsChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileNewParamsChunkingStrategyStaticStatic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileNewParamsChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileGetParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

type VectorStoreFileUpdateParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	// The updated key-value attributes to store with the file.
	Attributes map[string]VectorStoreFileUpdateParamsAttributeUnion `json:"attributes,omitzero,required"`
	paramObj
}

func (r VectorStoreFileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorStoreFileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorStoreFileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorStoreFileUpdateParamsAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorStoreFileUpdateParamsAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorStoreFileUpdateParamsAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorStoreFileUpdateParamsAttributeUnion) asAny() any {
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

type VectorStoreFileListParams struct {
	// (Optional) A cursor for use in pagination. `after` is an object ID that defines
	// your place in the list.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// (Optional) A cursor for use in pagination. `before` is an object ID that defines
	// your place in the list.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// (Optional) A limit on the number of objects to be returned. Limit can range
	// between 1 and 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// (Optional) Sort order by the `created_at` timestamp of the objects. `asc` for
	// ascending order and `desc` for descending order.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// (Optional) Filter by file status to only return files with the specified status.
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

// (Optional) Filter by file status to only return files with the specified status.
type VectorStoreFileListParamsFilter string

const (
	VectorStoreFileListParamsFilterCompleted  VectorStoreFileListParamsFilter = "completed"
	VectorStoreFileListParamsFilterInProgress VectorStoreFileListParamsFilter = "in_progress"
	VectorStoreFileListParamsFilterCancelled  VectorStoreFileListParamsFilter = "cancelled"
	VectorStoreFileListParamsFilterFailed     VectorStoreFileListParamsFilter = "failed"
)

type VectorStoreFileDeleteParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

type VectorStoreFileContentParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}
