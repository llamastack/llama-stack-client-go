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
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// BetaDatasetService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaDatasetService] method instead.
type BetaDatasetService struct {
	Options []option.RequestOption
}

// NewBetaDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaDatasetService(opts ...option.RequestOption) (r BetaDatasetService) {
	r = BetaDatasetService{}
	r.Options = opts
	return
}

// Get a dataset by its ID.
func (r *BetaDatasetService) Get(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *BetaDatasetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1beta/datasets/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all datasets.
func (r *BetaDatasetService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ListDatasetsResponseData, err error) {
	var env ListDatasetsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1beta/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Append rows to a dataset.
func (r *BetaDatasetService) Appendrows(ctx context.Context, datasetID string, body BetaDatasetAppendrowsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1beta/datasetio/append-rows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a paginated list of rows from a dataset.
//
// Uses offset-based pagination where:
//
// - start_index: The starting index (0-based). If None, starts from beginning.
// - limit: Number of items to return. If None or -1, returns all items.
//
// The response includes:
//
// - data: List of items for the current page.
// - has_more: Whether there are more items available after this set.
func (r *BetaDatasetService) Iterrows(ctx context.Context, datasetID string, query BetaDatasetIterrowsParams, opts ...option.RequestOption) (res *BetaDatasetIterrowsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1beta/datasetio/iterrows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Register a new dataset.
//
// Deprecated: deprecated
func (r *BetaDatasetService) Register(ctx context.Context, body BetaDatasetRegisterParams, opts ...option.RequestOption) (res *BetaDatasetRegisterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1beta/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unregister a dataset by its ID.
//
// Deprecated: deprecated
func (r *BetaDatasetService) Unregister(ctx context.Context, datasetID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1beta/datasets/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response from listing datasets.
type ListDatasetsResponse struct {
	// List of datasets
	Data []ListDatasetsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDatasetsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListDatasetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dataset resource for storing and accessing training or evaluation data.
type ListDatasetsResponseData struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose string `json:"purpose,required"`
	// Data source configuration for the dataset
	Source ListDatasetsResponseDataSourceUnion `json:"source,required"`
	// Any additional metadata for this dataset
	Metadata map[string]any `json:"metadata"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Type of resource, always 'dataset' for datasets
	//
	// Any of "dataset".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Metadata           respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDatasetsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListDatasetsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListDatasetsResponseDataSourceUnion contains all possible properties and values
// from [ListDatasetsResponseDataSourceUri], [ListDatasetsResponseDataSourceRows].
//
// Use the [ListDatasetsResponseDataSourceUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ListDatasetsResponseDataSourceUnion struct {
	// This field is from variant [ListDatasetsResponseDataSourceUri].
	Uri string `json:"uri"`
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [ListDatasetsResponseDataSourceRows].
	Rows []map[string]any `json:"rows"`
	JSON struct {
		Uri  respjson.Field
		Type respjson.Field
		Rows respjson.Field
		raw  string
	} `json:"-"`
}

// anyListDatasetsResponseDataSource is implemented by each variant of
// [ListDatasetsResponseDataSourceUnion] to add type safety for the return type of
// [ListDatasetsResponseDataSourceUnion.AsAny]
type anyListDatasetsResponseDataSource interface {
	implListDatasetsResponseDataSourceUnion()
}

func (ListDatasetsResponseDataSourceUri) implListDatasetsResponseDataSourceUnion()  {}
func (ListDatasetsResponseDataSourceRows) implListDatasetsResponseDataSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ListDatasetsResponseDataSourceUnion.AsAny().(type) {
//	case llamastackclient.ListDatasetsResponseDataSourceUri:
//	case llamastackclient.ListDatasetsResponseDataSourceRows:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ListDatasetsResponseDataSourceUnion) AsAny() anyListDatasetsResponseDataSource {
	switch u.Type {
	case "uri":
		return u.AsUri()
	case "rows":
		return u.AsRows()
	}
	return nil
}

func (u ListDatasetsResponseDataSourceUnion) AsUri() (v ListDatasetsResponseDataSourceUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataSourceUnion) AsRows() (v ListDatasetsResponseDataSourceRows) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListDatasetsResponseDataSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *ListDatasetsResponseDataSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset that can be obtained from a URI.
type ListDatasetsResponseDataSourceUri struct {
	// The dataset can be obtained from a URI. E.g.
	// "https://mywebsite.com/mydata.jsonl", "lsfs://mydata.jsonl",
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// The type of data source.
	//
	// Any of "uri".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDatasetsResponseDataSourceUri) RawJSON() string { return r.JSON.raw }
func (r *ListDatasetsResponseDataSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
type ListDatasetsResponseDataSourceRows struct {
	// The dataset is stored in rows. E.g. [{"messages": [{"role": "user", "content":
	// "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}]
	Rows []map[string]any `json:"rows,required"`
	// The type of data source.
	//
	// Any of "rows".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDatasetsResponseDataSourceRows) RawJSON() string { return r.JSON.raw }
func (r *ListDatasetsResponseDataSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dataset resource for storing and accessing training or evaluation data.
type BetaDatasetGetResponse struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose BetaDatasetGetResponsePurpose `json:"purpose,required"`
	// Data source configuration for the dataset
	Source BetaDatasetGetResponseSourceUnion `json:"source,required"`
	// Any additional metadata for this dataset
	Metadata map[string]any `json:"metadata"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Type of resource, always 'dataset' for datasets
	//
	// Any of "dataset".
	Type BetaDatasetGetResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Metadata           respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose of the dataset indicating its intended use
type BetaDatasetGetResponsePurpose string

const (
	BetaDatasetGetResponsePurposePostTrainingMessages BetaDatasetGetResponsePurpose = "post-training/messages"
	BetaDatasetGetResponsePurposeEvalQuestionAnswer   BetaDatasetGetResponsePurpose = "eval/question-answer"
	BetaDatasetGetResponsePurposeEvalMessagesAnswer   BetaDatasetGetResponsePurpose = "eval/messages-answer"
)

// BetaDatasetGetResponseSourceUnion contains all possible properties and values
// from [BetaDatasetGetResponseSourceUri], [BetaDatasetGetResponseSourceRows].
//
// Use the [BetaDatasetGetResponseSourceUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BetaDatasetGetResponseSourceUnion struct {
	// This field is from variant [BetaDatasetGetResponseSourceUri].
	Uri string `json:"uri"`
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [BetaDatasetGetResponseSourceRows].
	Rows []map[string]any `json:"rows"`
	JSON struct {
		Uri  respjson.Field
		Type respjson.Field
		Rows respjson.Field
		raw  string
	} `json:"-"`
}

// anyBetaDatasetGetResponseSource is implemented by each variant of
// [BetaDatasetGetResponseSourceUnion] to add type safety for the return type of
// [BetaDatasetGetResponseSourceUnion.AsAny]
type anyBetaDatasetGetResponseSource interface {
	implBetaDatasetGetResponseSourceUnion()
}

func (BetaDatasetGetResponseSourceUri) implBetaDatasetGetResponseSourceUnion()  {}
func (BetaDatasetGetResponseSourceRows) implBetaDatasetGetResponseSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BetaDatasetGetResponseSourceUnion.AsAny().(type) {
//	case llamastackclient.BetaDatasetGetResponseSourceUri:
//	case llamastackclient.BetaDatasetGetResponseSourceRows:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BetaDatasetGetResponseSourceUnion) AsAny() anyBetaDatasetGetResponseSource {
	switch u.Type {
	case "uri":
		return u.AsUri()
	case "rows":
		return u.AsRows()
	}
	return nil
}

func (u BetaDatasetGetResponseSourceUnion) AsUri() (v BetaDatasetGetResponseSourceUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseSourceUnion) AsRows() (v BetaDatasetGetResponseSourceRows) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetGetResponseSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetGetResponseSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset that can be obtained from a URI.
type BetaDatasetGetResponseSourceUri struct {
	// The dataset can be obtained from a URI. E.g.
	// "https://mywebsite.com/mydata.jsonl", "lsfs://mydata.jsonl",
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// The type of data source.
	//
	// Any of "uri".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetGetResponseSourceUri) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetGetResponseSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
type BetaDatasetGetResponseSourceRows struct {
	// The dataset is stored in rows. E.g. [{"messages": [{"role": "user", "content":
	// "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}]
	Rows []map[string]any `json:"rows,required"`
	// The type of data source.
	//
	// Any of "rows".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetGetResponseSourceRows) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetGetResponseSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of resource, always 'dataset' for datasets
type BetaDatasetGetResponseType string

const (
	BetaDatasetGetResponseTypeDataset BetaDatasetGetResponseType = "dataset"
)

// A generic paginated response that follows a simple format.
type BetaDatasetIterrowsResponse struct {
	Data    []map[string]any `json:"data,required"`
	HasMore bool             `json:"has_more,required"`
	URL     string           `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetIterrowsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetIterrowsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dataset resource for storing and accessing training or evaluation data.
type BetaDatasetRegisterResponse struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose BetaDatasetRegisterResponsePurpose `json:"purpose,required"`
	// Data source configuration for the dataset
	Source BetaDatasetRegisterResponseSourceUnion `json:"source,required"`
	// Any additional metadata for this dataset
	Metadata map[string]any `json:"metadata"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Type of resource, always 'dataset' for datasets
	//
	// Any of "dataset".
	Type BetaDatasetRegisterResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Metadata           respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose of the dataset indicating its intended use
type BetaDatasetRegisterResponsePurpose string

const (
	BetaDatasetRegisterResponsePurposePostTrainingMessages BetaDatasetRegisterResponsePurpose = "post-training/messages"
	BetaDatasetRegisterResponsePurposeEvalQuestionAnswer   BetaDatasetRegisterResponsePurpose = "eval/question-answer"
	BetaDatasetRegisterResponsePurposeEvalMessagesAnswer   BetaDatasetRegisterResponsePurpose = "eval/messages-answer"
)

// BetaDatasetRegisterResponseSourceUnion contains all possible properties and
// values from [BetaDatasetRegisterResponseSourceUri],
// [BetaDatasetRegisterResponseSourceRows].
//
// Use the [BetaDatasetRegisterResponseSourceUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BetaDatasetRegisterResponseSourceUnion struct {
	// This field is from variant [BetaDatasetRegisterResponseSourceUri].
	Uri string `json:"uri"`
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [BetaDatasetRegisterResponseSourceRows].
	Rows []map[string]any `json:"rows"`
	JSON struct {
		Uri  respjson.Field
		Type respjson.Field
		Rows respjson.Field
		raw  string
	} `json:"-"`
}

// anyBetaDatasetRegisterResponseSource is implemented by each variant of
// [BetaDatasetRegisterResponseSourceUnion] to add type safety for the return type
// of [BetaDatasetRegisterResponseSourceUnion.AsAny]
type anyBetaDatasetRegisterResponseSource interface {
	implBetaDatasetRegisterResponseSourceUnion()
}

func (BetaDatasetRegisterResponseSourceUri) implBetaDatasetRegisterResponseSourceUnion()  {}
func (BetaDatasetRegisterResponseSourceRows) implBetaDatasetRegisterResponseSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BetaDatasetRegisterResponseSourceUnion.AsAny().(type) {
//	case llamastackclient.BetaDatasetRegisterResponseSourceUri:
//	case llamastackclient.BetaDatasetRegisterResponseSourceRows:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BetaDatasetRegisterResponseSourceUnion) AsAny() anyBetaDatasetRegisterResponseSource {
	switch u.Type {
	case "uri":
		return u.AsUri()
	case "rows":
		return u.AsRows()
	}
	return nil
}

func (u BetaDatasetRegisterResponseSourceUnion) AsUri() (v BetaDatasetRegisterResponseSourceUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseSourceUnion) AsRows() (v BetaDatasetRegisterResponseSourceRows) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetRegisterResponseSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetRegisterResponseSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset that can be obtained from a URI.
type BetaDatasetRegisterResponseSourceUri struct {
	// The dataset can be obtained from a URI. E.g.
	// "https://mywebsite.com/mydata.jsonl", "lsfs://mydata.jsonl",
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// The type of data source.
	//
	// Any of "uri".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetRegisterResponseSourceUri) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetRegisterResponseSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
type BetaDatasetRegisterResponseSourceRows struct {
	// The dataset is stored in rows. E.g. [{"messages": [{"role": "user", "content":
	// "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}]
	Rows []map[string]any `json:"rows,required"`
	// The type of data source.
	//
	// Any of "rows".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetRegisterResponseSourceRows) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetRegisterResponseSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of resource, always 'dataset' for datasets
type BetaDatasetRegisterResponseType string

const (
	BetaDatasetRegisterResponseTypeDataset BetaDatasetRegisterResponseType = "dataset"
)

type BetaDatasetAppendrowsParams struct {
	Rows []map[string]any `json:"rows,omitzero,required"`
	paramObj
}

func (r BetaDatasetAppendrowsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetAppendrowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetAppendrowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaDatasetIterrowsParams struct {
	Limit      param.Opt[int64] `query:"limit,omitzero" json:"-"`
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDatasetIterrowsParams]'s query parameters as
// `url.Values`.
func (r BetaDatasetIterrowsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDatasetRegisterParams struct {
	// The purpose of the dataset.
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose BetaDatasetRegisterParamsPurpose `json:"purpose,omitzero,required"`
	// The data source of the dataset.
	Source BetaDatasetRegisterParamsSourceUnion `json:"source,omitzero,required"`
	// The ID of the dataset. If not provided, an ID will be generated.
	DatasetID param.Opt[string] `json:"dataset_id,omitzero"`
	// The metadata for the dataset.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r BetaDatasetRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The purpose of the dataset.
type BetaDatasetRegisterParamsPurpose string

const (
	BetaDatasetRegisterParamsPurposePostTrainingMessages BetaDatasetRegisterParamsPurpose = "post-training/messages"
	BetaDatasetRegisterParamsPurposeEvalQuestionAnswer   BetaDatasetRegisterParamsPurpose = "eval/question-answer"
	BetaDatasetRegisterParamsPurposeEvalMessagesAnswer   BetaDatasetRegisterParamsPurpose = "eval/messages-answer"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaDatasetRegisterParamsSourceUnion struct {
	OfUri  *BetaDatasetRegisterParamsSourceUri  `json:",omitzero,inline"`
	OfRows *BetaDatasetRegisterParamsSourceRows `json:",omitzero,inline"`
	paramUnion
}

func (u BetaDatasetRegisterParamsSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUri, u.OfRows)
}
func (u *BetaDatasetRegisterParamsSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaDatasetRegisterParamsSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfUri) {
		return u.OfUri
	} else if !param.IsOmitted(u.OfRows) {
		return u.OfRows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaDatasetRegisterParamsSourceUnion) GetUri() *string {
	if vt := u.OfUri; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaDatasetRegisterParamsSourceUnion) GetRows() []map[string]any {
	if vt := u.OfRows; vt != nil {
		return vt.Rows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaDatasetRegisterParamsSourceUnion) GetType() *string {
	if vt := u.OfUri; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRows; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BetaDatasetRegisterParamsSourceUnion](
		"type",
		apijson.Discriminator[BetaDatasetRegisterParamsSourceUri]("uri"),
		apijson.Discriminator[BetaDatasetRegisterParamsSourceRows]("rows"),
	)
}

// A dataset that can be obtained from a URI.
//
// The property Uri is required.
type BetaDatasetRegisterParamsSourceUri struct {
	// The dataset can be obtained from a URI. E.g.
	// "https://mywebsite.com/mydata.jsonl", "lsfs://mydata.jsonl",
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// The type of data source.
	//
	// Any of "uri".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BetaDatasetRegisterParamsSourceUri) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetRegisterParamsSourceUri
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetRegisterParamsSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaDatasetRegisterParamsSourceUri](
		"type", "uri",
	)
}

// A dataset stored in rows.
//
// The property Rows is required.
type BetaDatasetRegisterParamsSourceRows struct {
	// The dataset is stored in rows. E.g. [{"messages": [{"role": "user", "content":
	// "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}]
	Rows []map[string]any `json:"rows,omitzero,required"`
	// The type of data source.
	//
	// Any of "rows".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BetaDatasetRegisterParamsSourceRows) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetRegisterParamsSourceRows
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetRegisterParamsSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaDatasetRegisterParamsSourceRows](
		"type", "rows",
	)
}
