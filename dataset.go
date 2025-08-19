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
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// DatasetService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r DatasetService) {
	r = DatasetService{}
	r.Options = opts
	return
}

// Get a dataset by its ID.
func (r *DatasetService) Get(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *DatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasets/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all datasets.
func (r *DatasetService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ListDatasetsResponseData, err error) {
	var env ListDatasetsResponse
	opts = append(r.Options[:], opts...)
	path := "v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Append rows to a dataset.
func (r *DatasetService) Appendrows(ctx context.Context, datasetID string, body DatasetAppendrowsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasetio/append-rows/%s", datasetID)
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
func (r *DatasetService) Iterrows(ctx context.Context, datasetID string, query DatasetIterrowsParams, opts ...option.RequestOption) (res *DatasetIterrowsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasetio/iterrows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Register a new dataset.
func (r *DatasetService) Register(ctx context.Context, body DatasetRegisterParams, opts ...option.RequestOption) (res *DatasetRegisterResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unregister a dataset by its ID.
func (r *DatasetService) Unregister(ctx context.Context, datasetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasets/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListDatasetsResponse struct {
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

type ListDatasetsResponseData struct {
	Identifier string                                           `json:"identifier,required"`
	Metadata   map[string]ListDatasetsResponseDataMetadataUnion `json:"metadata,required"`
	ProviderID string                                           `json:"provider_id,required"`
	// Purpose of the dataset. Each purpose has a required input data schema.
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose string `json:"purpose,required"`
	// A dataset that can be obtained from a URI.
	Source             ListDatasetsResponseDataSourceUnion `json:"source,required"`
	Type               constant.Dataset                    `json:"type,required"`
	ProviderResourceID string                              `json:"provider_resource_id"`
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

type DatasetGetResponse struct {
	Identifier string                                     `json:"identifier,required"`
	Metadata   map[string]DatasetGetResponseMetadataUnion `json:"metadata,required"`
	ProviderID string                                     `json:"provider_id,required"`
	// Purpose of the dataset. Each purpose has a required input data schema.
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose DatasetGetResponsePurpose `json:"purpose,required"`
	// A dataset that can be obtained from a URI.
	Source             DatasetGetResponseSourceUnion `json:"source,required"`
	Type               constant.Dataset              `json:"type,required"`
	ProviderResourceID string                        `json:"provider_resource_id"`
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
func (r DatasetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DatasetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetGetResponseMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetGetResponseMetadataUnion struct {
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

func (u DatasetGetResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetGetResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetGetResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose of the dataset. Each purpose has a required input data schema.
type DatasetGetResponsePurpose string

const (
	DatasetGetResponsePurposePostTrainingMessages DatasetGetResponsePurpose = "post-training/messages"
	DatasetGetResponsePurposeEvalQuestionAnswer   DatasetGetResponsePurpose = "eval/question-answer"
	DatasetGetResponsePurposeEvalMessagesAnswer   DatasetGetResponsePurpose = "eval/messages-answer"
)

// DatasetGetResponseSourceUnion contains all possible properties and values from
// [DatasetGetResponseSourceUri], [DatasetGetResponseSourceRows].
//
// Use the [DatasetGetResponseSourceUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DatasetGetResponseSourceUnion struct {
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [DatasetGetResponseSourceUri].
	Uri string `json:"uri"`
	// This field is from variant [DatasetGetResponseSourceRows].
	Rows []map[string]DatasetGetResponseSourceRowsRowUnion `json:"rows"`
	JSON struct {
		Type respjson.Field
		Uri  respjson.Field
		Rows respjson.Field
		raw  string
	} `json:"-"`
}

// anyDatasetGetResponseSource is implemented by each variant of
// [DatasetGetResponseSourceUnion] to add type safety for the return type of
// [DatasetGetResponseSourceUnion.AsAny]
type anyDatasetGetResponseSource interface {
	implDatasetGetResponseSourceUnion()
}

func (DatasetGetResponseSourceUri) implDatasetGetResponseSourceUnion()  {}
func (DatasetGetResponseSourceRows) implDatasetGetResponseSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := DatasetGetResponseSourceUnion.AsAny().(type) {
//	case llamastackclient.DatasetGetResponseSourceUri:
//	case llamastackclient.DatasetGetResponseSourceRows:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u DatasetGetResponseSourceUnion) AsAny() anyDatasetGetResponseSource {
	switch u.Type {
	case "uri":
		return u.AsUri()
	case "rows":
		return u.AsRows()
	}
	return nil
}

func (u DatasetGetResponseSourceUnion) AsUri() (v DatasetGetResponseSourceUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseSourceUnion) AsRows() (v DatasetGetResponseSourceRows) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetGetResponseSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetGetResponseSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset that can be obtained from a URI.
type DatasetGetResponseSourceUri struct {
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
func (r DatasetGetResponseSourceUri) RawJSON() string { return r.JSON.raw }
func (r *DatasetGetResponseSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
type DatasetGetResponseSourceRows struct {
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]DatasetGetResponseSourceRowsRowUnion `json:"rows,required"`
	Type constant.Rows                                     `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetGetResponseSourceRows) RawJSON() string { return r.JSON.raw }
func (r *DatasetGetResponseSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetGetResponseSourceRowsRowUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetGetResponseSourceRowsRowUnion struct {
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

func (u DatasetGetResponseSourceRowsRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseSourceRowsRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseSourceRowsRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetGetResponseSourceRowsRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetGetResponseSourceRowsRowUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetGetResponseSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A generic paginated response that follows a simple format.
type DatasetIterrowsResponse struct {
	// The list of items for the current page
	Data []map[string]DatasetIterrowsResponseDataUnion `json:"data,required"`
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
func (r DatasetIterrowsResponse) RawJSON() string { return r.JSON.raw }
func (r *DatasetIterrowsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetIterrowsResponseDataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetIterrowsResponseDataUnion struct {
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

func (u DatasetIterrowsResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetIterrowsResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetIterrowsResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetIterrowsResponseDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetIterrowsResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetIterrowsResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetRegisterResponse struct {
	Identifier string                                          `json:"identifier,required"`
	Metadata   map[string]DatasetRegisterResponseMetadataUnion `json:"metadata,required"`
	ProviderID string                                          `json:"provider_id,required"`
	// Purpose of the dataset. Each purpose has a required input data schema.
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose DatasetRegisterResponsePurpose `json:"purpose,required"`
	// A dataset that can be obtained from a URI.
	Source             DatasetRegisterResponseSourceUnion `json:"source,required"`
	Type               constant.Dataset                   `json:"type,required"`
	ProviderResourceID string                             `json:"provider_resource_id"`
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
func (r DatasetRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *DatasetRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetRegisterResponseMetadataUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetRegisterResponseMetadataUnion struct {
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

func (u DatasetRegisterResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetRegisterResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetRegisterResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose of the dataset. Each purpose has a required input data schema.
type DatasetRegisterResponsePurpose string

const (
	DatasetRegisterResponsePurposePostTrainingMessages DatasetRegisterResponsePurpose = "post-training/messages"
	DatasetRegisterResponsePurposeEvalQuestionAnswer   DatasetRegisterResponsePurpose = "eval/question-answer"
	DatasetRegisterResponsePurposeEvalMessagesAnswer   DatasetRegisterResponsePurpose = "eval/messages-answer"
)

// DatasetRegisterResponseSourceUnion contains all possible properties and values
// from [DatasetRegisterResponseSourceUri], [DatasetRegisterResponseSourceRows].
//
// Use the [DatasetRegisterResponseSourceUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DatasetRegisterResponseSourceUnion struct {
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [DatasetRegisterResponseSourceUri].
	Uri string `json:"uri"`
	// This field is from variant [DatasetRegisterResponseSourceRows].
	Rows []map[string]DatasetRegisterResponseSourceRowsRowUnion `json:"rows"`
	JSON struct {
		Type respjson.Field
		Uri  respjson.Field
		Rows respjson.Field
		raw  string
	} `json:"-"`
}

// anyDatasetRegisterResponseSource is implemented by each variant of
// [DatasetRegisterResponseSourceUnion] to add type safety for the return type of
// [DatasetRegisterResponseSourceUnion.AsAny]
type anyDatasetRegisterResponseSource interface {
	implDatasetRegisterResponseSourceUnion()
}

func (DatasetRegisterResponseSourceUri) implDatasetRegisterResponseSourceUnion()  {}
func (DatasetRegisterResponseSourceRows) implDatasetRegisterResponseSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := DatasetRegisterResponseSourceUnion.AsAny().(type) {
//	case llamastackclient.DatasetRegisterResponseSourceUri:
//	case llamastackclient.DatasetRegisterResponseSourceRows:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u DatasetRegisterResponseSourceUnion) AsAny() anyDatasetRegisterResponseSource {
	switch u.Type {
	case "uri":
		return u.AsUri()
	case "rows":
		return u.AsRows()
	}
	return nil
}

func (u DatasetRegisterResponseSourceUnion) AsUri() (v DatasetRegisterResponseSourceUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseSourceUnion) AsRows() (v DatasetRegisterResponseSourceRows) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetRegisterResponseSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetRegisterResponseSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset that can be obtained from a URI.
type DatasetRegisterResponseSourceUri struct {
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
func (r DatasetRegisterResponseSourceUri) RawJSON() string { return r.JSON.raw }
func (r *DatasetRegisterResponseSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
type DatasetRegisterResponseSourceRows struct {
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]DatasetRegisterResponseSourceRowsRowUnion `json:"rows,required"`
	Type constant.Rows                                          `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetRegisterResponseSourceRows) RawJSON() string { return r.JSON.raw }
func (r *DatasetRegisterResponseSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetRegisterResponseSourceRowsRowUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetRegisterResponseSourceRowsRowUnion struct {
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

func (u DatasetRegisterResponseSourceRowsRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseSourceRowsRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseSourceRowsRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetRegisterResponseSourceRowsRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetRegisterResponseSourceRowsRowUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetRegisterResponseSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetAppendrowsParams struct {
	// The rows to append to the dataset.
	Rows []map[string]DatasetAppendrowsParamsRowUnion `json:"rows,omitzero,required"`
	paramObj
}

func (r DatasetAppendrowsParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetAppendrowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetAppendrowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatasetAppendrowsParamsRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DatasetAppendrowsParamsRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DatasetAppendrowsParamsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatasetAppendrowsParamsRowUnion) asAny() any {
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

type DatasetIterrowsParams struct {
	// The number of rows to get.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Index into dataset for the first row to get. Get all rows if None.
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DatasetIterrowsParams]'s query parameters as `url.Values`.
func (r DatasetIterrowsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DatasetRegisterParams struct {
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
	Purpose DatasetRegisterParamsPurpose `json:"purpose,omitzero,required"`
	// The data source of the dataset. Ensure that the data source schema is compatible
	// with the purpose of the dataset. Examples: - { "type": "uri", "uri":
	// "https://mywebsite.com/mydata.jsonl" } - { "type": "uri", "uri":
	// "lsfs://mydata.jsonl" } - { "type": "uri", "uri":
	// "data:csv;base64,{base64_content}" } - { "type": "uri", "uri":
	// "huggingface://llamastack/simpleqa?split=train" } - { "type": "rows", "rows": [
	// { "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
	// "assistant", "content": "Hello, world!"}, ] } ] }
	Source DatasetRegisterParamsSourceUnion `json:"source,omitzero,required"`
	// The ID of the dataset. If not provided, an ID will be generated.
	DatasetID param.Opt[string] `json:"dataset_id,omitzero"`
	// The metadata for the dataset. - E.g. {"description": "My dataset"}.
	Metadata map[string]DatasetRegisterParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r DatasetRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetRegisterParams) UnmarshalJSON(data []byte) error {
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
type DatasetRegisterParamsPurpose string

const (
	DatasetRegisterParamsPurposePostTrainingMessages DatasetRegisterParamsPurpose = "post-training/messages"
	DatasetRegisterParamsPurposeEvalQuestionAnswer   DatasetRegisterParamsPurpose = "eval/question-answer"
	DatasetRegisterParamsPurposeEvalMessagesAnswer   DatasetRegisterParamsPurpose = "eval/messages-answer"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatasetRegisterParamsSourceUnion struct {
	OfUri  *DatasetRegisterParamsSourceUri  `json:",omitzero,inline"`
	OfRows *DatasetRegisterParamsSourceRows `json:",omitzero,inline"`
	paramUnion
}

func (u DatasetRegisterParamsSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUri, u.OfRows)
}
func (u *DatasetRegisterParamsSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatasetRegisterParamsSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfUri) {
		return u.OfUri
	} else if !param.IsOmitted(u.OfRows) {
		return u.OfRows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DatasetRegisterParamsSourceUnion) GetUri() *string {
	if vt := u.OfUri; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DatasetRegisterParamsSourceUnion) GetRows() []map[string]DatasetRegisterParamsSourceRowsRowUnion {
	if vt := u.OfRows; vt != nil {
		return vt.Rows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DatasetRegisterParamsSourceUnion) GetType() *string {
	if vt := u.OfUri; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRows; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[DatasetRegisterParamsSourceUnion](
		"type",
		apijson.Discriminator[DatasetRegisterParamsSourceUri]("uri"),
		apijson.Discriminator[DatasetRegisterParamsSourceRows]("rows"),
	)
}

// A dataset that can be obtained from a URI.
//
// The properties Type, Uri are required.
type DatasetRegisterParamsSourceUri struct {
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// This field can be elided, and will marshal its zero value as "uri".
	Type constant.Uri `json:"type,required"`
	paramObj
}

func (r DatasetRegisterParamsSourceUri) MarshalJSON() (data []byte, err error) {
	type shadow DatasetRegisterParamsSourceUri
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetRegisterParamsSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
//
// The properties Rows, Type are required.
type DatasetRegisterParamsSourceRows struct {
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]DatasetRegisterParamsSourceRowsRowUnion `json:"rows,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "rows".
	Type constant.Rows `json:"type,required"`
	paramObj
}

func (r DatasetRegisterParamsSourceRows) MarshalJSON() (data []byte, err error) {
	type shadow DatasetRegisterParamsSourceRows
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetRegisterParamsSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatasetRegisterParamsSourceRowsRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DatasetRegisterParamsSourceRowsRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DatasetRegisterParamsSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatasetRegisterParamsSourceRowsRowUnion) asAny() any {
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
type DatasetRegisterParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DatasetRegisterParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DatasetRegisterParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatasetRegisterParamsMetadataUnion) asAny() any {
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
