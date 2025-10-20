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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1beta/datasetio/append-rows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a paginated list of rows from a dataset. Uses offset-based pagination where:
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
func (r *BetaDatasetService) Register(ctx context.Context, body BetaDatasetRegisterParams, opts ...option.RequestOption) (res *BetaDatasetRegisterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1beta/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unregister a dataset by its ID.
func (r *BetaDatasetService) Unregister(ctx context.Context, datasetID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	Identifier string `json:"identifier,required"`
	// Additional metadata for the dataset
	Metadata   map[string]ListDatasetsResponseDataMetadataUnion `json:"metadata,required"`
	ProviderID string                                           `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose string `json:"purpose,required"`
	// Data source configuration for the dataset
	Source ListDatasetsResponseDataSourceUnion `json:"source,required"`
	// Type of resource, always 'dataset' for datasets
	Type               constant.Dataset `json:"type,required"`
	ProviderResourceID string           `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDatasetsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ListDatasetsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListDatasetsResponseDataMetadataUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ListDatasetsResponseDataMetadataUnion struct {
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

func (u ListDatasetsResponseDataMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListDatasetsResponseDataMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ListDatasetsResponseDataMetadataUnion) UnmarshalJSON(data []byte) error {
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
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [ListDatasetsResponseDataSourceUri].
	Uri string `json:"uri"`
	// This field is from variant [ListDatasetsResponseDataSourceRows].
	Rows []map[string]ListDatasetsResponseDataSourceRowsRowUnion `json:"rows"`
	JSON struct {
		Type respjson.Field
		Uri  respjson.Field
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
	Type constant.Uri `json:"type,required"`
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Uri         respjson.Field
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
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]ListDatasetsResponseDataSourceRowsRowUnion `json:"rows,required"`
	Type constant.Rows                                           `json:"type,required"`
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

// ListDatasetsResponseDataSourceRowsRowUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ListDatasetsResponseDataSourceRowsRowUnion struct {
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

func (u ListDatasetsResponseDataSourceRowsRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataSourceRowsRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataSourceRowsRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListDatasetsResponseDataSourceRowsRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListDatasetsResponseDataSourceRowsRowUnion) RawJSON() string { return u.JSON.raw }

func (r *ListDatasetsResponseDataSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dataset resource for storing and accessing training or evaluation data.
type BetaDatasetGetResponse struct {
	Identifier string `json:"identifier,required"`
	// Additional metadata for the dataset
	Metadata   map[string]BetaDatasetGetResponseMetadataUnion `json:"metadata,required"`
	ProviderID string                                         `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose BetaDatasetGetResponsePurpose `json:"purpose,required"`
	// Data source configuration for the dataset
	Source BetaDatasetGetResponseSourceUnion `json:"source,required"`
	// Type of resource, always 'dataset' for datasets
	Type               constant.Dataset `json:"type,required"`
	ProviderResourceID string           `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDatasetGetResponseMetadataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BetaDatasetGetResponseMetadataUnion struct {
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

func (u BetaDatasetGetResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetGetResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetGetResponseMetadataUnion) UnmarshalJSON(data []byte) error {
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
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [BetaDatasetGetResponseSourceUri].
	Uri string `json:"uri"`
	// This field is from variant [BetaDatasetGetResponseSourceRows].
	Rows []map[string]BetaDatasetGetResponseSourceRowsRowUnion `json:"rows"`
	JSON struct {
		Type respjson.Field
		Uri  respjson.Field
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
	Type constant.Uri `json:"type,required"`
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Uri         respjson.Field
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
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]BetaDatasetGetResponseSourceRowsRowUnion `json:"rows,required"`
	Type constant.Rows                                         `json:"type,required"`
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

// BetaDatasetGetResponseSourceRowsRowUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BetaDatasetGetResponseSourceRowsRowUnion struct {
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

func (u BetaDatasetGetResponseSourceRowsRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseSourceRowsRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseSourceRowsRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetGetResponseSourceRowsRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetGetResponseSourceRowsRowUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetGetResponseSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A generic paginated response that follows a simple format.
type BetaDatasetIterrowsResponse struct {
	// The list of items for the current page
	Data []map[string]BetaDatasetIterrowsResponseDataUnion `json:"data,required"`
	// Whether there are more items available after this set
	HasMore bool `json:"has_more,required"`
	// The URL for accessing this list
	URL string `json:"url"`
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

// BetaDatasetIterrowsResponseDataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BetaDatasetIterrowsResponseDataUnion struct {
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

func (u BetaDatasetIterrowsResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetIterrowsResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetIterrowsResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetIterrowsResponseDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetIterrowsResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetIterrowsResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dataset resource for storing and accessing training or evaluation data.
type BetaDatasetRegisterResponse struct {
	Identifier string `json:"identifier,required"`
	// Additional metadata for the dataset
	Metadata   map[string]BetaDatasetRegisterResponseMetadataUnion `json:"metadata,required"`
	ProviderID string                                              `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose BetaDatasetRegisterResponsePurpose `json:"purpose,required"`
	// Data source configuration for the dataset
	Source BetaDatasetRegisterResponseSourceUnion `json:"source,required"`
	// Type of resource, always 'dataset' for datasets
	Type               constant.Dataset `json:"type,required"`
	ProviderResourceID string           `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDatasetRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDatasetRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDatasetRegisterResponseMetadataUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BetaDatasetRegisterResponseMetadataUnion struct {
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

func (u BetaDatasetRegisterResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetRegisterResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetRegisterResponseMetadataUnion) UnmarshalJSON(data []byte) error {
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
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [BetaDatasetRegisterResponseSourceUri].
	Uri string `json:"uri"`
	// This field is from variant [BetaDatasetRegisterResponseSourceRows].
	Rows []map[string]BetaDatasetRegisterResponseSourceRowsRowUnion `json:"rows"`
	JSON struct {
		Type respjson.Field
		Uri  respjson.Field
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
	Type constant.Uri `json:"type,required"`
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Uri         respjson.Field
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
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]BetaDatasetRegisterResponseSourceRowsRowUnion `json:"rows,required"`
	Type constant.Rows                                              `json:"type,required"`
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

// BetaDatasetRegisterResponseSourceRowsRowUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BetaDatasetRegisterResponseSourceRowsRowUnion struct {
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

func (u BetaDatasetRegisterResponseSourceRowsRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseSourceRowsRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseSourceRowsRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDatasetRegisterResponseSourceRowsRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDatasetRegisterResponseSourceRowsRowUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDatasetRegisterResponseSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaDatasetAppendrowsParams struct {
	// The rows to append to the dataset.
	Rows []map[string]BetaDatasetAppendrowsParamsRowUnion `json:"rows,omitzero,required"`
	paramObj
}

func (r BetaDatasetAppendrowsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetAppendrowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetAppendrowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaDatasetAppendrowsParamsRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u BetaDatasetAppendrowsParamsRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *BetaDatasetAppendrowsParamsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaDatasetAppendrowsParamsRowUnion) asAny() any {
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

type BetaDatasetIterrowsParams struct {
	// The number of rows to get.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Index into dataset for the first row to get. Get all rows if None.
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
	// The purpose of the dataset. One of: - "post-training/messages": The dataset
	// contains a messages column with list of messages for post-training. {
	// "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
	// "assistant", "content": "Hello, world!"}, ] } - "eval/question-answer": The
	// dataset contains a question column and an answer column for evaluation. {
	// "question": "What is the capital of France?", "answer": "Paris" } -
	// "eval/messages-answer": The dataset contains a messages column with list of
	// messages and an answer column for evaluation. { "messages": [ {"role": "user",
	// "content": "Hello, my name is John Doe."}, {"role": "assistant", "content":
	// "Hello, John Doe. How can I help you today?"}, {"role": "user", "content":
	// "What's my name?"}, ], "answer": "John Doe" }
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose BetaDatasetRegisterParamsPurpose `json:"purpose,omitzero,required"`
	// The data source of the dataset. Ensure that the data source schema is compatible
	// with the purpose of the dataset. Examples: - { "type": "uri", "uri":
	// "https://mywebsite.com/mydata.jsonl" } - { "type": "uri", "uri":
	// "lsfs://mydata.jsonl" } - { "type": "uri", "uri":
	// "data:csv;base64,{base64_content}" } - { "type": "uri", "uri":
	// "huggingface://llamastack/simpleqa?split=train" } - { "type": "rows", "rows": [
	// { "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
	// "assistant", "content": "Hello, world!"}, ] } ] }
	Source BetaDatasetRegisterParamsSourceUnion `json:"source,omitzero,required"`
	// The ID of the dataset. If not provided, an ID will be generated.
	DatasetID param.Opt[string] `json:"dataset_id,omitzero"`
	// The metadata for the dataset. - E.g. {"description": "My dataset"}.
	Metadata map[string]BetaDatasetRegisterParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r BetaDatasetRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The purpose of the dataset. One of: - "post-training/messages": The dataset
// contains a messages column with list of messages for post-training. {
// "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
// "assistant", "content": "Hello, world!"}, ] } - "eval/question-answer": The
// dataset contains a question column and an answer column for evaluation. {
// "question": "What is the capital of France?", "answer": "Paris" } -
// "eval/messages-answer": The dataset contains a messages column with list of
// messages and an answer column for evaluation. { "messages": [ {"role": "user",
// "content": "Hello, my name is John Doe."}, {"role": "assistant", "content":
// "Hello, John Doe. How can I help you today?"}, {"role": "user", "content":
// "What's my name?"}, ], "answer": "John Doe" }
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
func (u BetaDatasetRegisterParamsSourceUnion) GetRows() []map[string]BetaDatasetRegisterParamsSourceRowsRowUnion {
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
// The properties Type, Uri are required.
type BetaDatasetRegisterParamsSourceUri struct {
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// This field can be elided, and will marshal its zero value as "uri".
	Type constant.Uri `json:"type,required"`
	paramObj
}

func (r BetaDatasetRegisterParamsSourceUri) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetRegisterParamsSourceUri
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetRegisterParamsSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
//
// The properties Rows, Type are required.
type BetaDatasetRegisterParamsSourceRows struct {
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]BetaDatasetRegisterParamsSourceRowsRowUnion `json:"rows,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "rows".
	Type constant.Rows `json:"type,required"`
	paramObj
}

func (r BetaDatasetRegisterParamsSourceRows) MarshalJSON() (data []byte, err error) {
	type shadow BetaDatasetRegisterParamsSourceRows
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDatasetRegisterParamsSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaDatasetRegisterParamsSourceRowsRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u BetaDatasetRegisterParamsSourceRowsRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *BetaDatasetRegisterParamsSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaDatasetRegisterParamsSourceRowsRowUnion) asAny() any {
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
type BetaDatasetRegisterParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u BetaDatasetRegisterParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *BetaDatasetRegisterParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaDatasetRegisterParamsMetadataUnion) asAny() any {
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
