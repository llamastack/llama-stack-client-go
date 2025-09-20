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

// ResponseInputItemService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseInputItemService] method instead.
type ResponseInputItemService struct {
	Options []option.RequestOption
}

// NewResponseInputItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewResponseInputItemService(opts ...option.RequestOption) (r ResponseInputItemService) {
	r = ResponseInputItemService{}
	r.Options = opts
	return
}

// List input items for a given OpenAI response.
func (r *ResponseInputItemService) List(ctx context.Context, responseID string, query ResponseInputItemListParams, opts ...option.RequestOption) (res *ResponseInputItemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/responses/%s/input_items", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List container for OpenAI response input items.
type ResponseInputItemListResponse struct {
	// List of input items
	Data []ResponseInputItemListResponseDataUnion `json:"data,required"`
	// Object type identifier, always "list"
	Object constant.List `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseInputItemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataUnion contains all possible properties and
// values from
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput],
// [ResponseInputItemListResponseDataOpenAIResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataUnion struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall].
	Results []ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall].
	Arguments string `json:"arguments"`
	CallID    string `json:"call_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall].
	Name string `json:"name"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessage].
	Content ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion `json:"content"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessage].
	Role ResponseInputItemListResponseDataOpenAIResponseMessageRole `json:"role"`
	JSON struct {
		ID        respjson.Field
		Status    respjson.Field
		Type      respjson.Field
		Queries   respjson.Field
		Results   respjson.Field
		Arguments respjson.Field
		CallID    respjson.Field
		Name      respjson.Field
		Output    respjson.Field
		Content   respjson.Field
		Role      respjson.Field
		raw       string
	} `json:"-"`
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMessage() (v ResponseInputItemListResponseDataOpenAIResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseInputItemListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// Current status of the web search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "web_search_call"
	Type constant.WebSearchCall `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion `json:"attributes,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result (between 0 and 1)
	Score float64 `json:"score,required"`
	// Text content of the search result
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion struct {
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

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall struct {
	// JSON string containing the function arguments
	Arguments string `json:"arguments,required"`
	// Unique identifier for the function call
	CallID string `json:"call_id,required"`
	// Name of the function being called
	Name string `json:"name,required"`
	// Tool call type identifier, always "function_call"
	Type constant.FunctionCall `json:"type,required"`
	// (Optional) Additional identifier for the tool call
	ID string `json:"id"`
	// (Optional) Current status of the function call execution
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput struct {
	CallID string                      `json:"call_id,required"`
	Output string                      `json:"output,required"`
	Type   constant.FunctionCallOutput `json:"type,required"`
	ID     string                      `json:"id"`
	Status string                      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseInputItemListResponseDataOpenAIResponseMessage struct {
	Content ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseInputItemListResponseDataOpenAIResponseMessageRole `json:"role,required"`
	Type   constant.Message                                           `json:"type,required"`
	ID     string                                                     `json:"id"`
	Status string                                                     `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseInputItemListResponseDataOpenAIResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion],
// [[]ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseInputItemListResponseDataOpenAIResponseMessageContentArray OfVariant2]
type ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseInputItemListResponseDataOpenAIResponseMessageContentArray []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem]
	// instead of an object.
	OfVariant2 []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                                             respjson.Field
		OfResponseInputItemListResponseDataOpenAIResponseMessageContentArray respjson.Field
		OfVariant2                                                           respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion) AsResponseInputItemListResponseDataOpenAIResponseMessageContentArray() (v []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion) AsVariant2() (v []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText],
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage].
	Detail ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem is
// implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion] to
// add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) AsInputText() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) AsInputImage() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	Type constant.InputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
	// Content type identifier, always "input_image"
	Type constant.InputImage `json:"type,required"`
	// (Optional) URL of the image content
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Type        respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetailLow  ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetailHigh ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetailAuto ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetailLow  ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetailHigh ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetailAuto ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail = "auto"
)

type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem struct {
	Annotations []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                                  `json:"text,required"`
	Type        constant.OutputText                                                                     `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion
// contains all possible properties and values from
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotation
// is implemented by each variant of
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion]
// to add type safety for the return type of
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotation interface {
	implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion()
}

func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation struct {
	// Unique identifier of the referenced file
	FileID string `json:"file_id,required"`
	// Name of the referenced file
	Filename string `json:"filename,required"`
	// Position index of the citation within the content
	Index int64 `json:"index,required"`
	// Annotation type identifier, always "file_citation"
	Type constant.FileCitation `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation struct {
	// End position of the citation span in the content
	EndIndex int64 `json:"end_index,required"`
	// Start position of the citation span in the content
	StartIndex int64 `json:"start_index,required"`
	// Title of the referenced web resource
	Title string `json:"title,required"`
	// Annotation type identifier, always "url_citation"
	Type constant.URLCitation `json:"type,required"`
	// URL of the referenced web resource
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
	ContainerID string                         `json:"container_id,required"`
	EndIndex    int64                          `json:"end_index,required"`
	FileID      string                         `json:"file_id,required"`
	Filename    string                         `json:"filename,required"`
	StartIndex  int64                          `json:"start_index,required"`
	Type        constant.ContainerFileCitation `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath struct {
	FileID string            `json:"file_id,required"`
	Index  int64             `json:"index,required"`
	Type   constant.FilePath `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputItemListResponseDataOpenAIResponseMessageRole string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleSystem    ResponseInputItemListResponseDataOpenAIResponseMessageRole = "system"
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleDeveloper ResponseInputItemListResponseDataOpenAIResponseMessageRole = "developer"
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleUser      ResponseInputItemListResponseDataOpenAIResponseMessageRole = "user"
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleAssistant ResponseInputItemListResponseDataOpenAIResponseMessageRole = "assistant"
)

type ResponseInputItemListResponseDataRole string

const (
	ResponseInputItemListResponseDataRoleSystem    ResponseInputItemListResponseDataRole = "system"
	ResponseInputItemListResponseDataRoleDeveloper ResponseInputItemListResponseDataRole = "developer"
	ResponseInputItemListResponseDataRoleUser      ResponseInputItemListResponseDataRole = "user"
	ResponseInputItemListResponseDataRoleAssistant ResponseInputItemListResponseDataRole = "assistant"
)

type ResponseInputItemListParams struct {
	// An item ID to list items after, used for pagination.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// An item ID to list items before, used for pagination.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Additional fields to include in the response.
	Include []string `query:"include,omitzero" json:"-"`
	// The order to return the input items in. Default is desc.
	//
	// Any of "asc", "desc".
	Order ResponseInputItemListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResponseInputItemListParams]'s query parameters as
// `url.Values`.
func (r ResponseInputItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order to return the input items in. Default is desc.
type ResponseInputItemListParamsOrder string

const (
	ResponseInputItemListParamsOrderAsc  ResponseInputItemListParamsOrder = "asc"
	ResponseInputItemListParamsOrderDesc ResponseInputItemListParamsOrder = "desc"
)
