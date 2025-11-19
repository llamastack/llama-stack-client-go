// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
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
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apiform"
	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/pagination"
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

// Upload file.
//
// Upload a file that can be used across various endpoints.
//
// The file upload should be a multipart form request with:
//
// - file: The File object (not file name) to be uploaded.
// - purpose: The intended purpose of the uploaded file.
// - expires_after: Optional form values describing expiration for the file.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve file.
//
// Returns information about a specific file.
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List files.
//
// Returns a list of files that belong to the user's organization.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[File], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/files"
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

// List files.
//
// Returns a list of files that belong to the user's organization.
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[File] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete file.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *DeleteFileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve file content.
//
// Returns the contents of the specified file.
func (r *FileService) Content(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileContentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/content", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response for deleting a file in OpenAI Files API.
type DeleteFileResponse struct {
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted,required"`
	// Any of "file".
	Object DeleteFileResponseObject `json:"object"`
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

type DeleteFileResponseObject string

const (
	DeleteFileResponseObjectFile DeleteFileResponseObject = "file"
)

// OpenAI File object as defined in the OpenAI Files API.
type File struct {
	ID        string `json:"id,required"`
	Bytes     int64  `json:"bytes,required"`
	CreatedAt int64  `json:"created_at,required"`
	ExpiresAt int64  `json:"expires_at,required"`
	Filename  string `json:"filename,required"`
	// Valid purpose values for OpenAI Files API.
	//
	// Any of "assistants", "batch".
	Purpose FilePurpose `json:"purpose,required"`
	// Any of "file".
	Object FileObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bytes       respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		Purpose     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Valid purpose values for OpenAI Files API.
type FilePurpose string

const (
	FilePurposeAssistants FilePurpose = "assistants"
	FilePurposeBatch      FilePurpose = "batch"
)

type FileObject string

const (
	FileObjectFile FileObject = "file"
)

// Response for listing files in OpenAI Files API.
type ListFilesResponse struct {
	Data    []File `json:"data,required"`
	FirstID string `json:"first_id,required"`
	HasMore bool   `json:"has_more,required"`
	LastID  string `json:"last_id,required"`
	// Any of "list".
	Object ListFilesResponseObject `json:"object"`
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

type ListFilesResponseObject string

const (
	ListFilesResponseObjectList ListFilesResponseObject = "list"
)

type FileContentResponse = any

type FileNewParams struct {
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Valid purpose values for OpenAI Files API.
	//
	// Any of "assistants", "batch".
	Purpose FileNewParamsPurpose `json:"purpose,omitzero,required"`
	// Control expiration of uploaded files.
	//
	// Params:
	//
	// - anchor, must be "created_at"
	// - seconds, must be int between 3600 and 2592000 (1 hour to 30 days)
	ExpiresAfter FileNewParamsExpiresAfter `json:"expires_after,omitzero"`
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
	FileNewParamsPurposeBatch      FileNewParamsPurpose = "batch"
)

// Control expiration of uploaded files.
//
// Params:
//
// - anchor, must be "created_at"
// - seconds, must be int between 3600 and 2592000 (1 hour to 30 days)
//
// The properties Anchor, Seconds are required.
type FileNewParamsExpiresAfter struct {
	Seconds int64 `json:"seconds,required"`
	// This field can be elided, and will marshal its zero value as "created_at".
	Anchor constant.CreatedAt `json:"anchor,required"`
	paramObj
}

func (r FileNewParamsExpiresAfter) MarshalJSON() (data []byte, err error) {
	type shadow FileNewParamsExpiresAfter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileNewParamsExpiresAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListParams struct {
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Sort order for paginated responses.
	//
	// Any of "asc", "desc".
	Order FileListParamsOrder `query:"order,omitzero" json:"-"`
	// Valid purpose values for OpenAI Files API.
	//
	// Any of "assistants", "batch".
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

// Sort order for paginated responses.
type FileListParamsOrder string

const (
	FileListParamsOrderAsc  FileListParamsOrder = "asc"
	FileListParamsOrderDesc FileListParamsOrder = "desc"
)

// Valid purpose values for OpenAI Files API.
type FileListParamsPurpose string

const (
	FileListParamsPurposeAssistants FileListParamsPurpose = "assistants"
	FileListParamsPurposeBatch      FileListParamsPurpose = "batch"
)
