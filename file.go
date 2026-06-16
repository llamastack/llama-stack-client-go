// Copyright (c) The OGX Contributors.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ogxclient

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

	"github.com/ogx-ai/ogx-client-go/internal/apiform"
	"github.com/ogx-ai/ogx-client-go/internal/apijson"
	"github.com/ogx-ai/ogx-client-go/internal/apiquery"
	"github.com/ogx-ai/ogx-client-go/internal/requestconfig"
	"github.com/ogx-ai/ogx-client-go/option"
	"github.com/ogx-ai/ogx-client-go/packages/pagination"
	"github.com/ogx-ai/ogx-client-go/packages/param"
	"github.com/ogx-ai/ogx-client-go/packages/respjson"
)

// This API is used to upload documents that can be used with other OGX APIs.
//
// FileService contains methods and other services that help with interacting with
// the ogx-client API.
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

// Upload a file.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get file
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List files
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

// List files
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[File] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete file
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *DeleteFileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve file content
func (r *FileService) Content(ctx context.Context, fileID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s/content", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response for deleting a file in OpenAI Files API.
type DeleteFileResponse struct {
	// The file identifier that was deleted.
	ID string `json:"id" api:"required"`
	// Whether the file was successfully deleted.
	Deleted bool `json:"deleted" api:"required"`
	// The object type, which is always 'file'.
	//
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

// The object type, which is always 'file'.
type DeleteFileResponseObject string

const (
	DeleteFileResponseObjectFile DeleteFileResponseObject = "file"
)

// OpenAI File object as defined in the OpenAI Files API.
type File struct {
	// The file identifier, which can be referenced in the API endpoints.
	ID string `json:"id" api:"required"`
	// The size of the file, in bytes.
	Bytes int64 `json:"bytes" api:"required"`
	// The Unix timestamp (in seconds) for when the file was created.
	CreatedAt int64 `json:"created_at" api:"required"`
	// The name of the file.
	Filename string `json:"filename" api:"required"`
	// The intended purpose of the file.
	//
	// Any of "assistants", "assistants_output", "batch", "batch_output", "evals",
	// "fine-tune", "fine-tune-results", "vision", "user_data".
	Purpose FilePurpose `json:"purpose" api:"required"`
	// Deprecated. The current status of the file.
	//
	// Any of "uploaded", "processed", "error".
	//
	// Deprecated: deprecated
	Status FileStatus `json:"status" api:"required"`
	// The Unix timestamp (in seconds) for when the file will expire.
	ExpiresAt int64 `json:"expires_at"`
	// The object type, which is always 'file'.
	//
	// Any of "file".
	Object FileObject `json:"object"`
	// Deprecated. For details on why a fine-tuning training file failed validation,
	// see the error field on fine_tuning.job.
	//
	// Deprecated: deprecated
	StatusDetails string `json:"status_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Bytes         respjson.Field
		CreatedAt     respjson.Field
		Filename      respjson.Field
		Purpose       respjson.Field
		Status        respjson.Field
		ExpiresAt     respjson.Field
		Object        respjson.Field
		StatusDetails respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The intended purpose of the file.
type FilePurpose string

const (
	FilePurposeAssistants       FilePurpose = "assistants"
	FilePurposeAssistantsOutput FilePurpose = "assistants_output"
	FilePurposeBatch            FilePurpose = "batch"
	FilePurposeBatchOutput      FilePurpose = "batch_output"
	FilePurposeEvals            FilePurpose = "evals"
	FilePurposeFineTune         FilePurpose = "fine-tune"
	FilePurposeFineTuneResults  FilePurpose = "fine-tune-results"
	FilePurposeVision           FilePurpose = "vision"
	FilePurposeUserData         FilePurpose = "user_data"
)

// Deprecated. The current status of the file.
type FileStatus string

const (
	FileStatusUploaded  FileStatus = "uploaded"
	FileStatusProcessed FileStatus = "processed"
	FileStatusError     FileStatus = "error"
)

// The object type, which is always 'file'.
type FileObject string

const (
	FileObjectFile FileObject = "file"
)

// Response for listing files in OpenAI Files API.
type ListFilesResponse struct {
	// The list of files.
	Data []File `json:"data" api:"required"`
	// The ID of the first file in the list for pagination.
	FirstID string `json:"first_id" api:"required"`
	// Whether there are more files available beyond this page.
	HasMore bool `json:"has_more" api:"required"`
	// The ID of the last file in the list for pagination.
	LastID string `json:"last_id" api:"required"`
	// The object type, which is always 'list'.
	//
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

// The object type, which is always 'list'.
type ListFilesResponseObject string

const (
	ListFilesResponseObjectList ListFilesResponseObject = "list"
)

type FileNewParams struct {
	// The file to upload.
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// The intended purpose of the uploaded file.
	//
	// Any of "assistants", "batch", "fine-tune", "vision", "user_data", "evals".
	Purpose FileNewParamsPurpose `json:"purpose,omitzero" api:"required"`
	// Optional expiration settings for the file.
	ExpiresAfter param.Opt[string] `json:"expires_after,omitzero"`
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

// The intended purpose of the uploaded file.
type FileNewParamsPurpose string

const (
	FileNewParamsPurposeAssistants FileNewParamsPurpose = "assistants"
	FileNewParamsPurposeBatch      FileNewParamsPurpose = "batch"
	FileNewParamsPurposeFineTune   FileNewParamsPurpose = "fine-tune"
	FileNewParamsPurposeVision     FileNewParamsPurpose = "vision"
	FileNewParamsPurposeUserData   FileNewParamsPurpose = "user_data"
	FileNewParamsPurposeEvals      FileNewParamsPurpose = "evals"
)

type FileListParams struct {
	// A cursor for pagination. Returns files after this ID.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of files to return (1-10,000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by created_at timestamp ('asc' or 'desc').
	//
	// Any of "asc", "desc".
	Order FileListParamsOrder `query:"order,omitzero" json:"-"`
	// Filter files by purpose.
	//
	// Any of "assistants", "assistants_output", "batch", "batch_output", "evals",
	// "fine-tune", "fine-tune-results", "vision", "user_data".
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

// Sort order by created_at timestamp ('asc' or 'desc').
type FileListParamsOrder string

const (
	FileListParamsOrderAsc  FileListParamsOrder = "asc"
	FileListParamsOrderDesc FileListParamsOrder = "desc"
)

// Filter files by purpose.
type FileListParamsPurpose string

const (
	FileListParamsPurposeAssistants       FileListParamsPurpose = "assistants"
	FileListParamsPurposeAssistantsOutput FileListParamsPurpose = "assistants_output"
	FileListParamsPurposeBatch            FileListParamsPurpose = "batch"
	FileListParamsPurposeBatchOutput      FileListParamsPurpose = "batch_output"
	FileListParamsPurposeEvals            FileListParamsPurpose = "evals"
	FileListParamsPurposeFineTune         FileListParamsPurpose = "fine-tune"
	FileListParamsPurposeFineTuneResults  FileListParamsPurpose = "fine-tune-results"
	FileListParamsPurposeVision           FileListParamsPurpose = "vision"
	FileListParamsPurposeUserData         FileListParamsPurpose = "user_data"
)
