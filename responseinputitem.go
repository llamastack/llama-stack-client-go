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

// List input items.
func (r *ResponseInputItemService) List(ctx context.Context, responseID string, query ResponseInputItemListParams, opts ...option.RequestOption) (res *ResponseInputItemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s/input_items", responseID)
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
// values from [ResponseInputItemListResponseDataOpenAIResponseMessage],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall],
// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools],
// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest],
// [ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput],
// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse],
// [ResponseInputItemListResponseDataOpenAIResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputItemListResponseDataUnion struct {
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessage].
	Content ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion `json:"content"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessage].
	Role   ResponseInputItemListResponseDataOpenAIResponseMessageRole `json:"role"`
	Type   string                                                     `json:"type"`
	ID     string                                                     `json:"id"`
	Status string                                                     `json:"status"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCall].
	Results     []ResponseInputItemListResponseDataOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	Arguments   string                                                                                 `json:"arguments"`
	CallID      string                                                                                 `json:"call_id"`
	Name        string                                                                                 `json:"name"`
	ServerLabel string                                                                                 `json:"server_label"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall].
	Error  string `json:"error"`
	Output string `json:"output"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools].
	Tools []ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool `json:"tools"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	JSON   struct {
		Content           respjson.Field
		Role              respjson.Field
		Type              respjson.Field
		ID                respjson.Field
		Status            respjson.Field
		Queries           respjson.Field
		Results           respjson.Field
		Arguments         respjson.Field
		CallID            respjson.Field
		Name              respjson.Field
		ServerLabel       respjson.Field
		Error             respjson.Field
		Output            respjson.Field
		Tools             respjson.Field
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Reason            respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMessage() (v ResponseInputItemListResponseDataOpenAIResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
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

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageMcpCall() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseOutputMessageMcpListTools() (v ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMcpApprovalRequest() (v ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseInputItemListResponseDataOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsOpenAIResponseMcpApprovalResponse() (v ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataUnion) AsResponseInputItemListResponseDataOpenAIResponseMessage() (v ResponseInputItemListResponseDataOpenAIResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseInputItemListResponseDataUnion) UnmarshalJSON(data []byte) error {
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
// [[]ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfVariant2]
type ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion]
	// instead of an object.
	OfVariant2 []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfVariant2 respjson.Field
		raw        string
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

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentUnion) AsVariant2() (v []ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) {
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
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage],
// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile].
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
	// Any of "input_text", "input_image", "input_file".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage].
	Detail ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	FileID string                                                                                 `json:"file_id"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile].
	FileData string `json:"file_data"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile].
	FileURL string `json:"file_url"`
	// This field is from variant
	// [ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		FileID   respjson.Field
		ImageURL respjson.Field
		FileData respjson.Field
		FileURL  respjson.Field
		Filename respjson.Field
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
func (ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile) implResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputText:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputImage:
//	case llamastackclient.ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) AsAny() anyResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	case "input_file":
		return u.AsInputFile()
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

func (u ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemUnion) AsInputFile() (v ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile) {
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
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// (Optional) URL of the image content
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Type        respjson.Field
		FileID      respjson.Field
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

// File content for input messages in OpenAI response format.
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile struct {
	// The type of the input item. Always `input_file`.
	Type constant.InputFile `json:"type,required"`
	// The data of the file to be sent to the model.
	FileData string `json:"file_data"`
	// (Optional) The ID of the file to be sent to the model.
	FileID string `json:"file_id"`
	// The URL of the file to be sent to the model.
	FileURL string `json:"file_url"`
	// The name of the file to be sent to the model.
	Filename string `json:"filename"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		FileData    respjson.Field
		FileID      respjson.Field
		FileURL     respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemInputFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetailLow  ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail = "low"
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetailHigh ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail = "high"
	ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetailAuto ResponseInputItemListResponseDataOpenAIResponseMessageContentArrayItemDetail = "auto"
)

type ResponseInputItemListResponseDataOpenAIResponseMessageRole string

const (
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleSystem    ResponseInputItemListResponseDataOpenAIResponseMessageRole = "system"
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleDeveloper ResponseInputItemListResponseDataOpenAIResponseMessageRole = "developer"
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleUser      ResponseInputItemListResponseDataOpenAIResponseMessageRole = "user"
	ResponseInputItemListResponseDataOpenAIResponseMessageRoleAssistant ResponseInputItemListResponseDataOpenAIResponseMessageRole = "assistant"
)

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

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall struct {
	// Unique identifier for this MCP call
	ID string `json:"id,required"`
	// JSON string containing the MCP call arguments
	Arguments string `json:"arguments,required"`
	// Name of the MCP method being called
	Name string `json:"name,required"`
	// Label identifying the MCP server handling the call
	ServerLabel string `json:"server_label,required"`
	// Tool call type identifier, always "mcp_call"
	Type constant.McpCall `json:"type,required"`
	// (Optional) Error message if the MCP call failed
	Error string `json:"error"`
	// (Optional) Output result from the successful MCP call
	Output string `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool `json:"tools,required"`
	// Tool call type identifier, always "mcp_list_tools"
	Type constant.McpListTools `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Description of what the tool does
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseInputItemListResponseDataOpenAIResponseOutputMessageMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest struct {
	ID          string                      `json:"id,required"`
	Arguments   string                      `json:"arguments,required"`
	Name        string                      `json:"name,required"`
	ServerLabel string                      `json:"server_label,required"`
	Type        constant.McpApprovalRequest `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
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

// A response to an MCP approval request.
type ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse struct {
	ApprovalRequestID string                       `json:"approval_request_id,required"`
	Approve           bool                         `json:"approve,required"`
	Type              constant.McpApprovalResponse `json:"type,required"`
	ID                string                       `json:"id"`
	Reason            string                       `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Type              respjson.Field
		ID                respjson.Field
		Reason            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseInputItemListResponseDataOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
