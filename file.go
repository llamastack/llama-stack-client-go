// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/llamastack/llama-stack-client-go/internal/apiform"
	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// FileService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Upload a file that can be used across various endpoints. The file upload should
// be a multipart form request with:
//
// - file: The File object (not file name) to be uploaded.
// - purpose: The intended purpose of the uploaded file.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns information about a specific file.
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of files that belong to the user's organization.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *ListFilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a file.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *DeleteFileResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns the contents of the specified file.
func (r *FileService) Content(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/files/%s/content", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response for deleting a file in OpenAI Files API.
type DeleteFileResponse struct {
	// The file identifier that was deleted
	ID string `json:"id,required"`
	// Whether the file was successfully deleted
	Deleted bool `json:"deleted,required"`
	// The object type, which is always "file"
	Object constant.File `json:"object,required"`
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
func (r DeleteFileResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI File object as defined in the OpenAI Files API.
type File struct {
	// The file identifier, which can be referenced in the API endpoints
	ID string `json:"id,required"`
	// The size of the file, in bytes
	Bytes int64 `json:"bytes,required"`
	// The Unix timestamp (in seconds) for when the file was created
	CreatedAt int64 `json:"created_at,required"`
	// The Unix timestamp (in seconds) for when the file expires
	ExpiresAt int64 `json:"expires_at,required"`
	// The name of the file
	Filename string `json:"filename,required"`
	// The object type, which is always "file"
	Object constant.File `json:"object,required"`
	// The intended purpose of the file
	//
	// Any of "assistants".
	Purpose FilePurpose `json:"purpose,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bytes       respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		Object      respjson.Field
		Purpose     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The intended purpose of the file
type FilePurpose string

const (
	FilePurposeAssistants FilePurpose = "assistants"
)

// Response for listing files in OpenAI Files API.
type ListFilesResponse struct {
	// List of file objects
	Data    []File `json:"data,required"`
	FirstID string `json:"first_id,required"`
	HasMore bool   `json:"has_more,required"`
	LastID  string `json:"last_id,required"`
	// The object type, which is always "list"
	Object constant.List `json:"object,required"`
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
func (r ListFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileContentResponse = any

type FileNewParams struct {
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Valid purpose values for OpenAI Files API.
	//
	// Any of "assistants".
	Purpose FileNewParamsPurpose `json:"purpose,omitzero,required"`
	paramObj
}

func (r FileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Valid purpose values for OpenAI Files API.
type FileNewParamsPurpose string

const (
	FileNewParamsPurposeAssistants FileNewParamsPurpose = "assistants"
)

type FileListParams struct {
	// A cursor for use in pagination. `after` is an object ID that defines your place
	// in the list. For instance, if you make a list request and receive 100 objects,
	// ending with obj_foo, your subsequent call can include after=obj_foo in order to
	// fetch the next page of the list.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 10,000, and the default is 10,000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
	// order and `desc` for descending order.
	//
	// Any of "asc", "desc".
	Order FileListParamsOrder `query:"order,omitzero" json:"-"`
	// Only return files with the given purpose.
	//
	// Any of "assistants".
	Purpose FileListParamsPurpose `query:"purpose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
// order and `desc` for descending order.
type FileListParamsOrder string

const (
	FileListParamsOrderAsc  FileListParamsOrder = "asc"
	FileListParamsOrderDesc FileListParamsOrder = "desc"
)

// Only return files with the given purpose.
type FileListParamsPurpose string

const (
	FileListParamsPurposeAssistants FileListParamsPurpose = "assistants"
)
