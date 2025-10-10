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
	"github.com/llamastack/llama-stack-client-go/packages/pagination"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// ResponseService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseService] method instead.
type ResponseService struct {
	Options    []option.RequestOption
	InputItems ResponseInputItemService
}

// NewResponseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResponseService(opts ...option.RequestOption) (r ResponseService) {
	r = ResponseService{}
	r.Options = opts
	r.InputItems = NewResponseInputItemService(opts...)
	return
}

// Create a model response.
func (r *ResponseService) New(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a model response.
func (r *ResponseService) NewStreaming(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ResponseObjectStreamUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ResponseObjectStreamUnion](ssestream.NewDecoder(raw), err)
}

// Get a model response.
func (r *ResponseService) Get(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all responses.
func (r *ResponseService) List(ctx context.Context, query ResponseListParams, opts ...option.RequestOption) (res *pagination.OpenAICursorPage[ResponseListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/responses"
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

// List all responses.
func (r *ResponseService) ListAutoPaging(ctx context.Context, query ResponseListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ResponseListResponse] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a response.
func (r *ResponseService) Delete(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Complete OpenAI response object containing generation results and metadata.
type ResponseObject struct {
	// Unique identifier for this response
	ID string `json:"id,required"`
	// Unix timestamp when the response was created
	CreatedAt int64 `json:"created_at,required"`
	// Model identifier used for generation
	Model string `json:"model,required"`
	// Object type identifier, always "response"
	Object constant.Response `json:"object,required"`
	// List of generated output items (messages, tool calls, etc.)
	Output []ResponseObjectOutputUnion `json:"output,required"`
	// Whether tool calls can be executed in parallel
	ParallelToolCalls bool `json:"parallel_tool_calls,required"`
	// Current status of the response generation
	Status string `json:"status,required"`
	// Text formatting configuration for the response
	Text ResponseObjectText `json:"text,required"`
	// (Optional) Error details if the response generation failed
	Error ResponseObjectError `json:"error"`
	// (Optional) ID of the previous response in a conversation
	PreviousResponseID string `json:"previous_response_id"`
	// (Optional) Sampling temperature used for generation
	Temperature float64 `json:"temperature"`
	// (Optional) An array of tools the model may call while generating a response.
	Tools []ResponseObjectToolUnion `json:"tools"`
	// (Optional) Nucleus sampling parameter used for generation
	TopP float64 `json:"top_p"`
	// (Optional) Truncation strategy applied to the response
	Truncation string `json:"truncation"`
	// (Optional) Token usage information for the response
	Usage ResponseObjectUsage `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Model              respjson.Field
		Object             respjson.Field
		Output             respjson.Field
		ParallelToolCalls  respjson.Field
		Status             respjson.Field
		Text               respjson.Field
		Error              respjson.Field
		PreviousResponseID respjson.Field
		Temperature        respjson.Field
		Tools              respjson.Field
		TopP               respjson.Field
		Truncation         respjson.Field
		Usage              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputUnion contains all possible properties and values from
// [ResponseObjectOutputMessage], [ResponseObjectOutputWebSearchCall],
// [ResponseObjectOutputFileSearchCall], [ResponseObjectOutputFunctionCall],
// [ResponseObjectOutputMcpCall], [ResponseObjectOutputMcpListTools],
// [ResponseObjectOutputMcpApprovalRequest].
//
// Use the [ResponseObjectOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputUnion struct {
	// This field is from variant [ResponseObjectOutputMessage].
	Content ResponseObjectOutputMessageContentUnion `json:"content"`
	// This field is from variant [ResponseObjectOutputMessage].
	Role ResponseObjectOutputMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [ResponseObjectOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseObjectOutputFileSearchCall].
	Results   []ResponseObjectOutputFileSearchCallResult `json:"results"`
	Arguments string                                     `json:"arguments"`
	// This field is from variant [ResponseObjectOutputFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseObjectOutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [ResponseObjectOutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [ResponseObjectOutputMcpListTools].
	Tools []ResponseObjectOutputMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectOutput is implemented by each variant of
// [ResponseObjectOutputUnion] to add type safety for the return type of
// [ResponseObjectOutputUnion.AsAny]
type anyResponseObjectOutput interface {
	implResponseObjectOutputUnion()
}

func (ResponseObjectOutputMessage) implResponseObjectOutputUnion()            {}
func (ResponseObjectOutputWebSearchCall) implResponseObjectOutputUnion()      {}
func (ResponseObjectOutputFileSearchCall) implResponseObjectOutputUnion()     {}
func (ResponseObjectOutputFunctionCall) implResponseObjectOutputUnion()       {}
func (ResponseObjectOutputMcpCall) implResponseObjectOutputUnion()            {}
func (ResponseObjectOutputMcpListTools) implResponseObjectOutputUnion()       {}
func (ResponseObjectOutputMcpApprovalRequest) implResponseObjectOutputUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessage:
//	case llamastackclient.ResponseObjectOutputWebSearchCall:
//	case llamastackclient.ResponseObjectOutputFileSearchCall:
//	case llamastackclient.ResponseObjectOutputFunctionCall:
//	case llamastackclient.ResponseObjectOutputMcpCall:
//	case llamastackclient.ResponseObjectOutputMcpListTools:
//	case llamastackclient.ResponseObjectOutputMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputUnion) AsAny() anyResponseObjectOutput {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseObjectOutputUnion) AsMessage() (v ResponseObjectOutputMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsWebSearchCall() (v ResponseObjectOutputWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsFileSearchCall() (v ResponseObjectOutputFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsFunctionCall() (v ResponseObjectOutputFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsMcpCall() (v ResponseObjectOutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsMcpListTools() (v ResponseObjectOutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputUnion) AsMcpApprovalRequest() (v ResponseObjectOutputMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseObjectOutputMessage struct {
	Content ResponseObjectOutputMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseObjectOutputMessageRole `json:"role,required"`
	Type   constant.Message                `json:"type,required"`
	ID     string                          `json:"id"`
	Status string                          `json:"status"`
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
func (r ResponseObjectOutputMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentUnion contains all possible properties and
// values from [string], [[]ResponseObjectOutputMessageContentArrayItemUnion],
// [[]ResponseObjectOutputMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfResponseObjectOutputMessageContentArray OfVariant2]
type ResponseObjectOutputMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectOutputMessageContentArrayItemUnion] instead of an object.
	OfResponseObjectOutputMessageContentArray []ResponseObjectOutputMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectOutputMessageContentArrayItem] instead of an object.
	OfVariant2 []ResponseObjectOutputMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                  respjson.Field
		OfResponseObjectOutputMessageContentArray respjson.Field
		OfVariant2                                respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ResponseObjectOutputMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentUnion) AsResponseObjectOutputMessageContentArray() (v []ResponseObjectOutputMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentUnion) AsVariant2() (v []ResponseObjectOutputMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentArrayItemUnion contains all possible
// properties and values from
// [ResponseObjectOutputMessageContentArrayItemInputText],
// [ResponseObjectOutputMessageContentArrayItemInputImage].
//
// Use the [ResponseObjectOutputMessageContentArrayItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseObjectOutputMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentArrayItemInputImage].
	Detail ResponseObjectOutputMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectOutputMessageContentArrayItem is implemented by each variant of
// [ResponseObjectOutputMessageContentArrayItemUnion] to add type safety for the
// return type of [ResponseObjectOutputMessageContentArrayItemUnion.AsAny]
type anyResponseObjectOutputMessageContentArrayItem interface {
	implResponseObjectOutputMessageContentArrayItemUnion()
}

func (ResponseObjectOutputMessageContentArrayItemInputText) implResponseObjectOutputMessageContentArrayItemUnion() {
}
func (ResponseObjectOutputMessageContentArrayItemInputImage) implResponseObjectOutputMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessageContentArrayItemInputText:
//	case llamastackclient.ResponseObjectOutputMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputMessageContentArrayItemUnion) AsAny() anyResponseObjectOutputMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseObjectOutputMessageContentArrayItemUnion) AsInputText() (v ResponseObjectOutputMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentArrayItemUnion) AsInputImage() (v ResponseObjectOutputMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectOutputMessageContentArrayItemInputText struct {
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
func (r ResponseObjectOutputMessageContentArrayItemInputText) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectOutputMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseObjectOutputMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseObjectOutputMessageContentArrayItemInputImage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseObjectOutputMessageContentArrayItemInputImageDetail string

const (
	ResponseObjectOutputMessageContentArrayItemInputImageDetailLow  ResponseObjectOutputMessageContentArrayItemInputImageDetail = "low"
	ResponseObjectOutputMessageContentArrayItemInputImageDetailHigh ResponseObjectOutputMessageContentArrayItemInputImageDetail = "high"
	ResponseObjectOutputMessageContentArrayItemInputImageDetailAuto ResponseObjectOutputMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseObjectOutputMessageContentArrayItemDetail string

const (
	ResponseObjectOutputMessageContentArrayItemDetailLow  ResponseObjectOutputMessageContentArrayItemDetail = "low"
	ResponseObjectOutputMessageContentArrayItemDetailHigh ResponseObjectOutputMessageContentArrayItemDetail = "high"
	ResponseObjectOutputMessageContentArrayItemDetailAuto ResponseObjectOutputMessageContentArrayItemDetail = "auto"
)

type ResponseObjectOutputMessageContentArrayItem struct {
	Annotations []ResponseObjectOutputMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                       `json:"text,required"`
	Type        constant.OutputText                                          `json:"type,required"`
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
func (r ResponseObjectOutputMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMessageContentArrayItemAnnotationUnion contains all possible
// properties and values from
// [ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation],
// [ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation],
// [ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseObjectOutputMessageContentArrayItemAnnotationFilePath].
//
// Use the [ResponseObjectOutputMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectOutputMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseObjectOutputMessageContentArrayItemAnnotation is implemented by each
// variant of [ResponseObjectOutputMessageContentArrayItemAnnotationUnion] to add
// type safety for the return type of
// [ResponseObjectOutputMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseObjectOutputMessageContentArrayItemAnnotation interface {
	implResponseObjectOutputMessageContentArrayItemAnnotationUnion()
}

func (ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation) implResponseObjectOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation) implResponseObjectOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation) implResponseObjectOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectOutputMessageContentArrayItemAnnotationFilePath) implResponseObjectOutputMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectOutputMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectOutputMessageContentArrayItemAnnotationUnion) AsAny() anyResponseObjectOutputMessageContentArrayItemAnnotation {
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

func (u ResponseObjectOutputMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseObjectOutputMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectOutputMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseObjectOutputMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectOutputMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputMessageRole string

const (
	ResponseObjectOutputMessageRoleSystem    ResponseObjectOutputMessageRole = "system"
	ResponseObjectOutputMessageRoleDeveloper ResponseObjectOutputMessageRole = "developer"
	ResponseObjectOutputMessageRoleUser      ResponseObjectOutputMessageRole = "user"
	ResponseObjectOutputMessageRoleAssistant ResponseObjectOutputMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ResponseObjectOutputWebSearchCall struct {
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
func (r ResponseObjectOutputWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseObjectOutputFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseObjectOutputFileSearchCallResult `json:"results"`
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
func (r ResponseObjectOutputFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseObjectOutputFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseObjectOutputFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseObjectOutputFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputFileSearchCallResultAttributeUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectOutputFileSearchCallResultAttributeUnion struct {
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

func (u ResponseObjectOutputFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputFileSearchCallResultAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseObjectOutputFunctionCall struct {
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
func (r ResponseObjectOutputFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseObjectOutputMcpCall struct {
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
func (r ResponseObjectOutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseObjectOutputMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ResponseObjectOutputMcpListToolsTool `json:"tools,required"`
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
func (r ResponseObjectOutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseObjectOutputMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ResponseObjectOutputMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ResponseObjectOutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectOutputMcpListToolsToolInputSchemaUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectOutputMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseObjectOutputMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputMcpListToolsToolInputSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseObjectOutputMcpApprovalRequest struct {
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
func (r ResponseObjectOutputMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectOutputMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputRole string

const (
	ResponseObjectOutputRoleSystem    ResponseObjectOutputRole = "system"
	ResponseObjectOutputRoleDeveloper ResponseObjectOutputRole = "developer"
	ResponseObjectOutputRoleUser      ResponseObjectOutputRole = "user"
	ResponseObjectOutputRoleAssistant ResponseObjectOutputRole = "assistant"
)

// Text formatting configuration for the response
type ResponseObjectText struct {
	// (Optional) Text format configuration specifying output format requirements
	Format ResponseObjectTextFormat `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectText) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Text format configuration specifying output format requirements
type ResponseObjectTextFormat struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseObjectTextFormatType `json:"type,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description string `json:"description"`
	// The name of the response format. Only used for json_schema.
	Name string `json:"name"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseObjectTextFormatSchemaUnion `json:"schema"`
	// (Optional) Whether to strictly enforce the JSON schema. If true, the response
	// must match the schema exactly. Only used for json_schema.
	Strict bool `json:"strict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Schema      respjson.Field
		Strict      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "text", "json_schema", or "json_object" to identify the format type
type ResponseObjectTextFormatType string

const (
	ResponseObjectTextFormatTypeText       ResponseObjectTextFormatType = "text"
	ResponseObjectTextFormatTypeJsonSchema ResponseObjectTextFormatType = "json_schema"
	ResponseObjectTextFormatTypeJsonObject ResponseObjectTextFormatType = "json_object"
)

// ResponseObjectTextFormatSchemaUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectTextFormatSchemaUnion struct {
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

func (u ResponseObjectTextFormatSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectTextFormatSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectTextFormatSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectTextFormatSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectTextFormatSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectTextFormatSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Error details if the response generation failed
type ResponseObjectError struct {
	// Error code identifying the type of failure
	Code string `json:"code,required"`
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
func (r ResponseObjectError) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolUnion contains all possible properties and values from
// [ResponseObjectToolOpenAIResponseInputToolWebSearch],
// [ResponseObjectToolFileSearch], [ResponseObjectToolFunction],
// [ResponseObjectToolMcp].
//
// Use the [ResponseObjectToolUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectToolUnion struct {
	// Any of nil, "file_search", "function", "mcp".
	Type string `json:"type"`
	// This field is from variant [ResponseObjectToolOpenAIResponseInputToolWebSearch].
	SearchContextSize string `json:"search_context_size"`
	// This field is from variant [ResponseObjectToolFileSearch].
	VectorStoreIDs []string `json:"vector_store_ids"`
	// This field is from variant [ResponseObjectToolFileSearch].
	Filters map[string]ResponseObjectToolFileSearchFilterUnion `json:"filters"`
	// This field is from variant [ResponseObjectToolFileSearch].
	MaxNumResults int64 `json:"max_num_results"`
	// This field is from variant [ResponseObjectToolFileSearch].
	RankingOptions ResponseObjectToolFileSearchRankingOptions `json:"ranking_options"`
	// This field is from variant [ResponseObjectToolFunction].
	Name string `json:"name"`
	// This field is from variant [ResponseObjectToolFunction].
	Description string `json:"description"`
	// This field is from variant [ResponseObjectToolFunction].
	Parameters map[string]ResponseObjectToolFunctionParameterUnion `json:"parameters"`
	// This field is from variant [ResponseObjectToolFunction].
	Strict bool `json:"strict"`
	// This field is from variant [ResponseObjectToolMcp].
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseObjectToolMcp].
	AllowedTools ResponseObjectToolMcpAllowedToolsUnion `json:"allowed_tools"`
	JSON         struct {
		Type              respjson.Field
		SearchContextSize respjson.Field
		VectorStoreIDs    respjson.Field
		Filters           respjson.Field
		MaxNumResults     respjson.Field
		RankingOptions    respjson.Field
		Name              respjson.Field
		Description       respjson.Field
		Parameters        respjson.Field
		Strict            respjson.Field
		ServerLabel       respjson.Field
		AllowedTools      respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseObjectToolUnion) AsOpenAIResponseInputToolWebSearch() (v ResponseObjectToolOpenAIResponseInputToolWebSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolUnion) AsFileSearch() (v ResponseObjectToolFileSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolUnion) AsFunction() (v ResponseObjectToolFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolUnion) AsMcp() (v ResponseObjectToolMcp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool configuration for OpenAI response inputs.
type ResponseObjectToolOpenAIResponseInputToolWebSearch struct {
	// Web search tool type variant to use
	//
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11".
	Type ResponseObjectToolOpenAIResponseInputToolWebSearchType `json:"type,required"`
	// (Optional) Size of search context, must be "low", "medium", or "high"
	SearchContextSize string `json:"search_context_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		SearchContextSize respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolOpenAIResponseInputToolWebSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool type variant to use
type ResponseObjectToolOpenAIResponseInputToolWebSearchType string

const (
	ResponseObjectToolOpenAIResponseInputToolWebSearchTypeWebSearch                  ResponseObjectToolOpenAIResponseInputToolWebSearchType = "web_search"
	ResponseObjectToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview           ResponseObjectToolOpenAIResponseInputToolWebSearchType = "web_search_preview"
	ResponseObjectToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview2025_03_11 ResponseObjectToolOpenAIResponseInputToolWebSearchType = "web_search_preview_2025_03_11"
)

// File search tool configuration for OpenAI response inputs.
type ResponseObjectToolFileSearch struct {
	// Tool type identifier, always "file_search"
	Type constant.FileSearch `json:"type,required"`
	// List of vector store identifiers to search within
	VectorStoreIDs []string `json:"vector_store_ids,required"`
	// (Optional) Additional filters to apply to the search
	Filters map[string]ResponseObjectToolFileSearchFilterUnion `json:"filters"`
	// (Optional) Maximum number of search results to return (1-50)
	MaxNumResults int64 `json:"max_num_results"`
	// (Optional) Options for ranking and scoring search results
	RankingOptions ResponseObjectToolFileSearchRankingOptions `json:"ranking_options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type           respjson.Field
		VectorStoreIDs respjson.Field
		Filters        respjson.Field
		MaxNumResults  respjson.Field
		RankingOptions respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolFileSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolFileSearchFilterUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectToolFileSearchFilterUnion struct {
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

func (u ResponseObjectToolFileSearchFilterUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolFileSearchFilterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolFileSearchFilterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolFileSearchFilterUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolFileSearchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolFileSearchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Options for ranking and scoring search results
type ResponseObjectToolFileSearchRankingOptions struct {
	// (Optional) Name of the ranking algorithm to use
	Ranker string `json:"ranker"`
	// (Optional) Minimum relevance score threshold for results
	ScoreThreshold float64 `json:"score_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ranker         respjson.Field
		ScoreThreshold respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolFileSearchRankingOptions) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool configuration for OpenAI response inputs.
type ResponseObjectToolFunction struct {
	// Name of the function that can be called
	Name string `json:"name,required"`
	// Tool type identifier, always "function"
	Type constant.Function `json:"type,required"`
	// (Optional) Description of what the function does
	Description string `json:"description"`
	// (Optional) JSON schema defining the function's parameters
	Parameters map[string]ResponseObjectToolFunctionParameterUnion `json:"parameters"`
	// (Optional) Whether to enforce strict parameter validation
	Strict bool `json:"strict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Parameters  respjson.Field
		Strict      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolFunction) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolFunctionParameterUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectToolFunctionParameterUnion struct {
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

func (u ResponseObjectToolFunctionParameterUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolFunctionParameterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolFunctionParameterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolFunctionParameterUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolFunctionParameterUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolFunctionParameterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) tool configuration for OpenAI response object.
type ResponseObjectToolMcp struct {
	// Label to identify this MCP server
	ServerLabel string `json:"server_label,required"`
	// Tool type identifier, always "mcp"
	Type constant.Mcp `json:"type,required"`
	// (Optional) Restriction on which tools can be used from this server
	AllowedTools ResponseObjectToolMcpAllowedToolsUnion `json:"allowed_tools"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServerLabel  respjson.Field
		Type         respjson.Field
		AllowedTools respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolMcp) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectToolMcpAllowedToolsUnion contains all possible properties and
// values from [[]string], [ResponseObjectToolMcpAllowedToolsAllowedToolsFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfStringArray]
type ResponseObjectToolMcpAllowedToolsUnion struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant
	// [ResponseObjectToolMcpAllowedToolsAllowedToolsFilter].
	ToolNames []string `json:"tool_names"`
	JSON      struct {
		OfStringArray respjson.Field
		ToolNames     respjson.Field
		raw           string
	} `json:"-"`
}

func (u ResponseObjectToolMcpAllowedToolsUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectToolMcpAllowedToolsUnion) AsAllowedToolsFilter() (v ResponseObjectToolMcpAllowedToolsAllowedToolsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectToolMcpAllowedToolsUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter configuration for restricting which MCP tools can be used.
type ResponseObjectToolMcpAllowedToolsAllowedToolsFilter struct {
	// (Optional) List of specific tool names that are allowed
	ToolNames []string `json:"tool_names"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolNames   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectToolMcpAllowedToolsAllowedToolsFilter) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool type variant to use
type ResponseObjectToolType string

const (
	ResponseObjectToolTypeWebSearch                  ResponseObjectToolType = "web_search"
	ResponseObjectToolTypeWebSearchPreview           ResponseObjectToolType = "web_search_preview"
	ResponseObjectToolTypeWebSearchPreview2025_03_11 ResponseObjectToolType = "web_search_preview_2025_03_11"
	ResponseObjectToolTypeFileSearch                 ResponseObjectToolType = "file_search"
	ResponseObjectToolTypeFunction                   ResponseObjectToolType = "function"
	ResponseObjectToolTypeMcp                        ResponseObjectToolType = "mcp"
)

// (Optional) Token usage information for the response
type ResponseObjectUsage struct {
	// Number of tokens in the input
	InputTokens int64 `json:"input_tokens,required"`
	// Number of tokens in the output
	OutputTokens int64 `json:"output_tokens,required"`
	// Total tokens used (input + output)
	TotalTokens int64 `json:"total_tokens,required"`
	// Detailed breakdown of input token usage
	InputTokensDetails ResponseObjectUsageInputTokensDetails `json:"input_tokens_details"`
	// Detailed breakdown of output token usage
	OutputTokensDetails ResponseObjectUsageOutputTokensDetails `json:"output_tokens_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		OutputTokens        respjson.Field
		TotalTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokensDetails respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectUsage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of input token usage
type ResponseObjectUsageInputTokensDetails struct {
	// Number of tokens retrieved from cache
	CachedTokens int64 `json:"cached_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of output token usage
type ResponseObjectUsageOutputTokensDetails struct {
	// Number of tokens used for reasoning (o1/o3 models)
	ReasoningTokens int64 `json:"reasoning_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnion contains all possible properties and values from
// [ResponseObjectStreamResponseCreated], [ResponseObjectStreamResponseInProgress],
// [ResponseObjectStreamResponseOutputItemAdded],
// [ResponseObjectStreamResponseOutputItemDone],
// [ResponseObjectStreamResponseOutputTextDelta],
// [ResponseObjectStreamResponseOutputTextDone],
// [ResponseObjectStreamResponseFunctionCallArgumentsDelta],
// [ResponseObjectStreamResponseFunctionCallArgumentsDone],
// [ResponseObjectStreamResponseWebSearchCallInProgress],
// [ResponseObjectStreamResponseWebSearchCallSearching],
// [ResponseObjectStreamResponseWebSearchCallCompleted],
// [ResponseObjectStreamResponseMcpListToolsInProgress],
// [ResponseObjectStreamResponseMcpListToolsFailed],
// [ResponseObjectStreamResponseMcpListToolsCompleted],
// [ResponseObjectStreamResponseMcpCallArgumentsDelta],
// [ResponseObjectStreamResponseMcpCallArgumentsDone],
// [ResponseObjectStreamResponseMcpCallInProgress],
// [ResponseObjectStreamResponseMcpCallFailed],
// [ResponseObjectStreamResponseMcpCallCompleted],
// [ResponseObjectStreamResponseContentPartAdded],
// [ResponseObjectStreamResponseContentPartDone],
// [ResponseObjectStreamResponseIncomplete], [ResponseObjectStreamResponseFailed],
// [ResponseObjectStreamResponseCompleted].
//
// Use the [ResponseObjectStreamUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamUnion struct {
	// This field is from variant [ResponseObjectStreamResponseCreated].
	Response ResponseObject `json:"response"`
	// Any of "response.created", "response.in_progress", "response.output_item.added",
	// "response.output_item.done", "response.output_text.delta",
	// "response.output_text.done", "response.function_call_arguments.delta",
	// "response.function_call_arguments.done", "response.web_search_call.in_progress",
	// "response.web_search_call.searching", "response.web_search_call.completed",
	// "response.mcp_list_tools.in_progress", "response.mcp_list_tools.failed",
	// "response.mcp_list_tools.completed", "response.mcp_call.arguments.delta",
	// "response.mcp_call.arguments.done", "response.mcp_call.in_progress",
	// "response.mcp_call.failed", "response.mcp_call.completed",
	// "response.content_part.added", "response.content_part.done",
	// "response.incomplete", "response.failed", "response.completed".
	Type           string `json:"type"`
	SequenceNumber int64  `json:"sequence_number"`
	// This field is a union of [ResponseObjectStreamResponseOutputItemAddedItemUnion],
	// [ResponseObjectStreamResponseOutputItemDoneItemUnion]
	Item         ResponseObjectStreamUnionItem `json:"item"`
	OutputIndex  int64                         `json:"output_index"`
	ResponseID   string                        `json:"response_id"`
	ContentIndex int64                         `json:"content_index"`
	Delta        string                        `json:"delta"`
	ItemID       string                        `json:"item_id"`
	// This field is from variant [ResponseObjectStreamResponseOutputTextDone].
	Text      string `json:"text"`
	Arguments string `json:"arguments"`
	// This field is a union of
	// [ResponseObjectStreamResponseContentPartAddedPartUnion],
	// [ResponseObjectStreamResponseContentPartDonePartUnion]
	Part ResponseObjectStreamUnionPart `json:"part"`
	JSON struct {
		Response       respjson.Field
		Type           respjson.Field
		SequenceNumber respjson.Field
		Item           respjson.Field
		OutputIndex    respjson.Field
		ResponseID     respjson.Field
		ContentIndex   respjson.Field
		Delta          respjson.Field
		ItemID         respjson.Field
		Text           respjson.Field
		Arguments      respjson.Field
		Part           respjson.Field
		raw            string
	} `json:"-"`
}

// anyResponseObjectStream is implemented by each variant of
// [ResponseObjectStreamUnion] to add type safety for the return type of
// [ResponseObjectStreamUnion.AsAny]
type anyResponseObjectStream interface {
	implResponseObjectStreamUnion()
}

func (ResponseObjectStreamResponseCreated) implResponseObjectStreamUnion()                    {}
func (ResponseObjectStreamResponseInProgress) implResponseObjectStreamUnion()                 {}
func (ResponseObjectStreamResponseOutputItemAdded) implResponseObjectStreamUnion()            {}
func (ResponseObjectStreamResponseOutputItemDone) implResponseObjectStreamUnion()             {}
func (ResponseObjectStreamResponseOutputTextDelta) implResponseObjectStreamUnion()            {}
func (ResponseObjectStreamResponseOutputTextDone) implResponseObjectStreamUnion()             {}
func (ResponseObjectStreamResponseFunctionCallArgumentsDelta) implResponseObjectStreamUnion() {}
func (ResponseObjectStreamResponseFunctionCallArgumentsDone) implResponseObjectStreamUnion()  {}
func (ResponseObjectStreamResponseWebSearchCallInProgress) implResponseObjectStreamUnion()    {}
func (ResponseObjectStreamResponseWebSearchCallSearching) implResponseObjectStreamUnion()     {}
func (ResponseObjectStreamResponseWebSearchCallCompleted) implResponseObjectStreamUnion()     {}
func (ResponseObjectStreamResponseMcpListToolsInProgress) implResponseObjectStreamUnion()     {}
func (ResponseObjectStreamResponseMcpListToolsFailed) implResponseObjectStreamUnion()         {}
func (ResponseObjectStreamResponseMcpListToolsCompleted) implResponseObjectStreamUnion()      {}
func (ResponseObjectStreamResponseMcpCallArgumentsDelta) implResponseObjectStreamUnion()      {}
func (ResponseObjectStreamResponseMcpCallArgumentsDone) implResponseObjectStreamUnion()       {}
func (ResponseObjectStreamResponseMcpCallInProgress) implResponseObjectStreamUnion()          {}
func (ResponseObjectStreamResponseMcpCallFailed) implResponseObjectStreamUnion()              {}
func (ResponseObjectStreamResponseMcpCallCompleted) implResponseObjectStreamUnion()           {}
func (ResponseObjectStreamResponseContentPartAdded) implResponseObjectStreamUnion()           {}
func (ResponseObjectStreamResponseContentPartDone) implResponseObjectStreamUnion()            {}
func (ResponseObjectStreamResponseIncomplete) implResponseObjectStreamUnion()                 {}
func (ResponseObjectStreamResponseFailed) implResponseObjectStreamUnion()                     {}
func (ResponseObjectStreamResponseCompleted) implResponseObjectStreamUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseCreated:
//	case llamastackclient.ResponseObjectStreamResponseInProgress:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAdded:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDone:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextDelta:
//	case llamastackclient.ResponseObjectStreamResponseOutputTextDone:
//	case llamastackclient.ResponseObjectStreamResponseFunctionCallArgumentsDelta:
//	case llamastackclient.ResponseObjectStreamResponseFunctionCallArgumentsDone:
//	case llamastackclient.ResponseObjectStreamResponseWebSearchCallInProgress:
//	case llamastackclient.ResponseObjectStreamResponseWebSearchCallSearching:
//	case llamastackclient.ResponseObjectStreamResponseWebSearchCallCompleted:
//	case llamastackclient.ResponseObjectStreamResponseMcpListToolsInProgress:
//	case llamastackclient.ResponseObjectStreamResponseMcpListToolsFailed:
//	case llamastackclient.ResponseObjectStreamResponseMcpListToolsCompleted:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallArgumentsDelta:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallArgumentsDone:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallInProgress:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallFailed:
//	case llamastackclient.ResponseObjectStreamResponseMcpCallCompleted:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAdded:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDone:
//	case llamastackclient.ResponseObjectStreamResponseIncomplete:
//	case llamastackclient.ResponseObjectStreamResponseFailed:
//	case llamastackclient.ResponseObjectStreamResponseCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamUnion) AsAny() anyResponseObjectStream {
	switch u.Type {
	case "response.created":
		return u.AsResponseCreated()
	case "response.in_progress":
		return u.AsResponseInProgress()
	case "response.output_item.added":
		return u.AsResponseOutputItemAdded()
	case "response.output_item.done":
		return u.AsResponseOutputItemDone()
	case "response.output_text.delta":
		return u.AsResponseOutputTextDelta()
	case "response.output_text.done":
		return u.AsResponseOutputTextDone()
	case "response.function_call_arguments.delta":
		return u.AsResponseFunctionCallArgumentsDelta()
	case "response.function_call_arguments.done":
		return u.AsResponseFunctionCallArgumentsDone()
	case "response.web_search_call.in_progress":
		return u.AsResponseWebSearchCallInProgress()
	case "response.web_search_call.searching":
		return u.AsResponseWebSearchCallSearching()
	case "response.web_search_call.completed":
		return u.AsResponseWebSearchCallCompleted()
	case "response.mcp_list_tools.in_progress":
		return u.AsResponseMcpListToolsInProgress()
	case "response.mcp_list_tools.failed":
		return u.AsResponseMcpListToolsFailed()
	case "response.mcp_list_tools.completed":
		return u.AsResponseMcpListToolsCompleted()
	case "response.mcp_call.arguments.delta":
		return u.AsResponseMcpCallArgumentsDelta()
	case "response.mcp_call.arguments.done":
		return u.AsResponseMcpCallArgumentsDone()
	case "response.mcp_call.in_progress":
		return u.AsResponseMcpCallInProgress()
	case "response.mcp_call.failed":
		return u.AsResponseMcpCallFailed()
	case "response.mcp_call.completed":
		return u.AsResponseMcpCallCompleted()
	case "response.content_part.added":
		return u.AsResponseContentPartAdded()
	case "response.content_part.done":
		return u.AsResponseContentPartDone()
	case "response.incomplete":
		return u.AsResponseIncomplete()
	case "response.failed":
		return u.AsResponseFailed()
	case "response.completed":
		return u.AsResponseCompleted()
	}
	return nil
}

func (u ResponseObjectStreamUnion) AsResponseCreated() (v ResponseObjectStreamResponseCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseInProgress() (v ResponseObjectStreamResponseInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputItemAdded() (v ResponseObjectStreamResponseOutputItemAdded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputItemDone() (v ResponseObjectStreamResponseOutputItemDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputTextDelta() (v ResponseObjectStreamResponseOutputTextDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseOutputTextDone() (v ResponseObjectStreamResponseOutputTextDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFunctionCallArgumentsDelta() (v ResponseObjectStreamResponseFunctionCallArgumentsDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFunctionCallArgumentsDone() (v ResponseObjectStreamResponseFunctionCallArgumentsDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseWebSearchCallInProgress() (v ResponseObjectStreamResponseWebSearchCallInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseWebSearchCallSearching() (v ResponseObjectStreamResponseWebSearchCallSearching) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseWebSearchCallCompleted() (v ResponseObjectStreamResponseWebSearchCallCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpListToolsInProgress() (v ResponseObjectStreamResponseMcpListToolsInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpListToolsFailed() (v ResponseObjectStreamResponseMcpListToolsFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpListToolsCompleted() (v ResponseObjectStreamResponseMcpListToolsCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallArgumentsDelta() (v ResponseObjectStreamResponseMcpCallArgumentsDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallArgumentsDone() (v ResponseObjectStreamResponseMcpCallArgumentsDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallInProgress() (v ResponseObjectStreamResponseMcpCallInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallFailed() (v ResponseObjectStreamResponseMcpCallFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseMcpCallCompleted() (v ResponseObjectStreamResponseMcpCallCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseContentPartAdded() (v ResponseObjectStreamResponseContentPartAdded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseContentPartDone() (v ResponseObjectStreamResponseContentPartDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseIncomplete() (v ResponseObjectStreamResponseIncomplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseFailed() (v ResponseObjectStreamResponseFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamUnion) AsResponseCompleted() (v ResponseObjectStreamResponseCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItem is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItem provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
type ResponseObjectStreamUnionItem struct {
	// This field is a union of
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion],
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion]
	Content ResponseObjectStreamUnionItemContent `json:"content"`
	Role    string                               `json:"role"`
	Type    string                               `json:"type"`
	ID      string                               `json:"id"`
	Status  string                               `json:"status"`
	Queries []string                             `json:"queries"`
	// This field is a union of
	// [[]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult],
	// [[]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult]
	Results     ResponseObjectStreamUnionItemResults `json:"results"`
	Arguments   string                               `json:"arguments"`
	CallID      string                               `json:"call_id"`
	Name        string                               `json:"name"`
	ServerLabel string                               `json:"server_label"`
	Error       string                               `json:"error"`
	Output      string                               `json:"output"`
	// This field is a union of
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool],
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool]
	Tools ResponseObjectStreamUnionItemTools `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItemContent is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItemContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseObjectStreamResponseOutputItemAddedItemMessageContentArray OfVariant2
// OfResponseObjectStreamResponseOutputItemDoneItemMessageContentArray]
type ResponseObjectStreamUnionItemContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseObjectStreamResponseOutputItemAddedItemMessageContentArray []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem]
	// instead of an object.
	OfVariant2 []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseObjectStreamResponseOutputItemDoneItemMessageContentArray []ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion `json:",inline"`
	JSON                                                                struct {
		OfString                                                             respjson.Field
		OfResponseObjectStreamResponseOutputItemAddedItemMessageContentArray respjson.Field
		OfVariant2                                                           respjson.Field
		OfResponseObjectStreamResponseOutputItemDoneItemMessageContentArray  respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItemContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItemResults is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItemResults provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResults
// OfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResults]
type ResponseObjectStreamUnionItemResults struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult] instead
	// of an object.
	OfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResults []ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult] instead
	// of an object.
	OfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResults []ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult `json:",inline"`
	JSON                                                                  struct {
		OfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResults respjson.Field
		OfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResults  respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItemResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionItemTools is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionItemTools provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTools
// OfResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTools]
type ResponseObjectStreamUnionItemTools struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool] instead of
	// an object.
	OfResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTools []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool] instead of an
	// object.
	OfResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTools []ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool `json:",inline"`
	JSON                                                              struct {
		OfResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTools respjson.Field
		OfResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTools  respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionItemTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionPart is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionPart provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
type ResponseObjectStreamUnionPart struct {
	// This field is a union of
	// [[]ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion],
	// [[]ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion]
	Annotations ResponseObjectStreamUnionPartAnnotations `json:"annotations"`
	Text        string                                   `json:"text"`
	Type        string                                   `json:"type"`
	// This field is a union of
	// [[]map[string]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion],
	// [[]map[string]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion]
	Logprobs ResponseObjectStreamUnionPartLogprobs `json:"logprobs"`
	Refusal  string                                `json:"refusal"`
	JSON     struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Logprobs    respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionPartAnnotations is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionPartAnnotations provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotations
// OfResponseObjectStreamResponseContentPartDonePartOutputTextAnnotations]
type ResponseObjectStreamUnionPartAnnotations struct {
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion]
	// instead of an object.
	OfResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotations []ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion]
	// instead of an object.
	OfResponseObjectStreamResponseContentPartDonePartOutputTextAnnotations []ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion `json:",inline"`
	JSON                                                                   struct {
		OfResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotations respjson.Field
		OfResponseObjectStreamResponseContentPartDonePartOutputTextAnnotations  respjson.Field
		raw                                                                     string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPartAnnotations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamUnionPartLogprobs is an implicit subunion of
// [ResponseObjectStreamUnion]. ResponseObjectStreamUnionPartLogprobs provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseObjectStreamUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfMapOfResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobMap
// OfMapOfResponseObjectStreamResponseContentPartDonePartOutputTextLogprobMap]
type ResponseObjectStreamUnionPartLogprobs struct {
	// This field will be present if the value is a
	// [[]map[string]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion]
	// instead of an object.
	OfMapOfResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobMap []map[string]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]map[string]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion]
	// instead of an object.
	OfMapOfResponseObjectStreamResponseContentPartDonePartOutputTextLogprobMap []map[string]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion `json:",inline"`
	JSON                                                                       struct {
		OfMapOfResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobMap respjson.Field
		OfMapOfResponseObjectStreamResponseContentPartDonePartOutputTextLogprobMap  respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPartLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating a new response has been created.
type ResponseObjectStreamResponseCreated struct {
	// The response object that was created
	Response ResponseObject `json:"response,required"`
	// Event type identifier, always "response.created"
	Type constant.ResponseCreated `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseCreated) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating the response remains in progress.
type ResponseObjectStreamResponseInProgress struct {
	// Current response state while in progress
	Response ResponseObject `json:"response,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.in_progress"
	Type constant.ResponseInProgress `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response       respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a new output item is added to the response.
type ResponseObjectStreamResponseOutputItemAdded struct {
	// The output item that was added (message, tool call, etc.)
	Item ResponseObjectStreamResponseOutputItemAddedItemUnion `json:"item,required"`
	// Index position of this item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Unique identifier of the response containing this output
	ResponseID string `json:"response_id,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.output_item.added"
	Type constant.ResponseOutputItemAdded `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item           respjson.Field
		OutputIndex    respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemAdded) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessage],
// [ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall],
// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall],
// [ResponseObjectStreamResponseOutputItemAddedItemFunctionCall],
// [ResponseObjectStreamResponseOutputItemAddedItemMcpCall],
// [ResponseObjectStreamResponseOutputItemAddedItemMcpListTools],
// [ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest].
//
// Use the [ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessage].
	Content ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion `json:"content"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessage].
	Role ResponseObjectStreamResponseOutputItemAddedItemMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall].
	Results   []ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult `json:"results"`
	Arguments string                                                                `json:"arguments"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMcpCall].
	Error string `json:"error"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMcpCall].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMcpListTools].
	Tools []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemAddedItem is implemented by each
// variant of [ResponseObjectStreamResponseOutputItemAddedItemUnion] to add type
// safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItem interface {
	implResponseObjectStreamResponseOutputItemAddedItemUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessage) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMcpCall) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) implResponseObjectStreamResponseOutputItemAddedItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemFunctionCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpListTools:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItem {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMessage() (v ResponseObjectStreamResponseOutputItemAddedItemMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsWebSearchCall() (v ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsFileSearchCall() (v ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsFunctionCall() (v ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMcpCall() (v ResponseObjectStreamResponseOutputItemAddedItemMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMcpListTools() (v ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) AsMcpApprovalRequest() (v ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseOutputItemAddedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseObjectStreamResponseOutputItemAddedItemMessage struct {
	Content ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseObjectStreamResponseOutputItemAddedItemMessageRole `json:"role,required"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion],
// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseObjectStreamResponseOutputItemAddedItemMessageContentArray OfVariant2]
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseObjectStreamResponseOutputItemAddedItemMessageContentArray []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem]
	// instead of an object.
	OfVariant2 []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                                             respjson.Field
		OfResponseObjectStreamResponseOutputItemAddedItemMessageContentArray respjson.Field
		OfVariant2                                                           respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) AsResponseObjectStreamResponseOutputItemAddedItemMessageContentArray() (v []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) AsVariant2() (v []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage].
//
// Use the
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage].
	Detail ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem is
// implemented by each variant of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion] to
// add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem interface {
	implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText) implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage) implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion) AsInputText() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion) AsInputImage() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail string

const (
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetailLow  ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail = "low"
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetailHigh ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail = "high"
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetailAuto ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetail string

const (
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetailLow  ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetail = "low"
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetailHigh ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetail = "high"
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetailAuto ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemDetail = "auto"
)

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem struct {
	Annotations []ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotation
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotation interface {
	implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion()
}

func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation) implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation) implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation) implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath) implResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) AsAny() anyResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotation {
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

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemMessageRole string

const (
	ResponseObjectStreamResponseOutputItemAddedItemMessageRoleSystem    ResponseObjectStreamResponseOutputItemAddedItemMessageRole = "system"
	ResponseObjectStreamResponseOutputItemAddedItemMessageRoleDeveloper ResponseObjectStreamResponseOutputItemAddedItemMessageRole = "developer"
	ResponseObjectStreamResponseOutputItemAddedItemMessageRoleUser      ResponseObjectStreamResponseOutputItemAddedItemMessageRole = "user"
	ResponseObjectStreamResponseOutputItemAddedItemMessageRoleAssistant ResponseObjectStreamResponseOutputItemAddedItemMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult `json:"results"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion struct {
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

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemFunctionCall struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemAddedItemMcpCall struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseObjectStreamResponseOutputItemAddedItemMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:"tools,required"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest struct {
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
func (r ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemAddedItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemRole string

const (
	ResponseObjectStreamResponseOutputItemAddedItemRoleSystem    ResponseObjectStreamResponseOutputItemAddedItemRole = "system"
	ResponseObjectStreamResponseOutputItemAddedItemRoleDeveloper ResponseObjectStreamResponseOutputItemAddedItemRole = "developer"
	ResponseObjectStreamResponseOutputItemAddedItemRoleUser      ResponseObjectStreamResponseOutputItemAddedItemRole = "user"
	ResponseObjectStreamResponseOutputItemAddedItemRoleAssistant ResponseObjectStreamResponseOutputItemAddedItemRole = "assistant"
)

// Streaming event for when an output item is completed.
type ResponseObjectStreamResponseOutputItemDone struct {
	// The completed output item (message, tool call, etc.)
	Item ResponseObjectStreamResponseOutputItemDoneItemUnion `json:"item,required"`
	// Index position of this item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Unique identifier of the response containing this output
	ResponseID string `json:"response_id,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.output_item.done"
	Type constant.ResponseOutputItemDone `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item           respjson.Field
		OutputIndex    respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputItemDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessage],
// [ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall],
// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall],
// [ResponseObjectStreamResponseOutputItemDoneItemFunctionCall],
// [ResponseObjectStreamResponseOutputItemDoneItemMcpCall],
// [ResponseObjectStreamResponseOutputItemDoneItemMcpListTools],
// [ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest].
//
// Use the [ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessage].
	Content ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion `json:"content"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessage].
	Role ResponseObjectStreamResponseOutputItemDoneItemMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall].
	Results   []ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult `json:"results"`
	Arguments string                                                               `json:"arguments"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMcpCall].
	Error string `json:"error"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMcpCall].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMcpListTools].
	Tools []ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemDoneItem is implemented by each variant
// of [ResponseObjectStreamResponseOutputItemDoneItemUnion] to add type safety for
// the return type of [ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItem interface {
	implResponseObjectStreamResponseOutputItemDoneItemUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessage) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMcpCall) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) implResponseObjectStreamResponseOutputItemDoneItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemFunctionCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpListTools:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItem {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMessage() (v ResponseObjectStreamResponseOutputItemDoneItemMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsWebSearchCall() (v ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsFileSearchCall() (v ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsFunctionCall() (v ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMcpCall() (v ResponseObjectStreamResponseOutputItemDoneItemMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMcpListTools() (v ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) AsMcpApprovalRequest() (v ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseOutputItemDoneItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseObjectStreamResponseOutputItemDoneItemMessage struct {
	Content ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseObjectStreamResponseOutputItemDoneItemMessageRole `json:"role,required"`
	Type   constant.Message                                          `json:"type,required"`
	ID     string                                                    `json:"id"`
	Status string                                                    `json:"status"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion],
// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseObjectStreamResponseOutputItemDoneItemMessageContentArray OfVariant2]
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseObjectStreamResponseOutputItemDoneItemMessageContentArray []ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem]
	// instead of an object.
	OfVariant2 []ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                                            respjson.Field
		OfResponseObjectStreamResponseOutputItemDoneItemMessageContentArray respjson.Field
		OfVariant2                                                          respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) AsResponseObjectStreamResponseOutputItemDoneItemMessageContentArray() (v []ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) AsVariant2() (v []ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage].
//
// Use the
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage].
	Detail ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem is
// implemented by each variant of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion] to
// add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem interface {
	implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText) implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage) implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion) AsInputText() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion) AsInputImage() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail string

const (
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetailLow  ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail = "low"
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetailHigh ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail = "high"
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetailAuto ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetail string

const (
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetailLow  ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetail = "low"
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetailHigh ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetail = "high"
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetailAuto ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemDetail = "auto"
)

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem struct {
	Annotations []ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                                 `json:"text,required"`
	Type        constant.OutputText                                                                    `json:"type,required"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotation
// is implemented by each variant of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion]
// to add type safety for the return type of
// [ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotation interface {
	implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion()
}

func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation) implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation) implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation) implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion() {
}
func (ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath) implResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) AsAny() anyResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotation {
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

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemMessageRole string

const (
	ResponseObjectStreamResponseOutputItemDoneItemMessageRoleSystem    ResponseObjectStreamResponseOutputItemDoneItemMessageRole = "system"
	ResponseObjectStreamResponseOutputItemDoneItemMessageRoleDeveloper ResponseObjectStreamResponseOutputItemDoneItemMessageRole = "developer"
	ResponseObjectStreamResponseOutputItemDoneItemMessageRoleUser      ResponseObjectStreamResponseOutputItemDoneItemMessageRole = "user"
	ResponseObjectStreamResponseOutputItemDoneItemMessageRoleAssistant ResponseObjectStreamResponseOutputItemDoneItemMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult `json:"results"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion struct {
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

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemFunctionCall struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseObjectStreamResponseOutputItemDoneItemMcpCall struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseObjectStreamResponseOutputItemDoneItemMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool `json:"tools,required"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest struct {
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
func (r ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseOutputItemDoneItemMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemRole string

const (
	ResponseObjectStreamResponseOutputItemDoneItemRoleSystem    ResponseObjectStreamResponseOutputItemDoneItemRole = "system"
	ResponseObjectStreamResponseOutputItemDoneItemRoleDeveloper ResponseObjectStreamResponseOutputItemDoneItemRole = "developer"
	ResponseObjectStreamResponseOutputItemDoneItemRoleUser      ResponseObjectStreamResponseOutputItemDoneItemRole = "user"
	ResponseObjectStreamResponseOutputItemDoneItemRoleAssistant ResponseObjectStreamResponseOutputItemDoneItemRole = "assistant"
)

// Streaming event for incremental text content updates.
type ResponseObjectStreamResponseOutputTextDelta struct {
	// Index position within the text content
	ContentIndex int64 `json:"content_index,required"`
	// Incremental text content being added
	Delta string `json:"delta,required"`
	// Unique identifier of the output item being updated
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.output_text.delta"
	Type constant.ResponseOutputTextDelta `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputTextDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when text output is completed.
type ResponseObjectStreamResponseOutputTextDone struct {
	// Index position within the text content
	ContentIndex int64 `json:"content_index,required"`
	// Unique identifier of the completed output item
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Final complete text content of the output item
	Text string `json:"text,required"`
	// Event type identifier, always "response.output_text.done"
	Type constant.ResponseOutputTextDone `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Text           respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseOutputTextDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseOutputTextDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for incremental function call argument updates.
type ResponseObjectStreamResponseFunctionCallArgumentsDelta struct {
	// Incremental function call arguments being added
	Delta string `json:"delta,required"`
	// Unique identifier of the function call being updated
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.function_call_arguments.delta"
	Type constant.ResponseFunctionCallArgumentsDelta `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFunctionCallArgumentsDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFunctionCallArgumentsDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when function call arguments are completed.
type ResponseObjectStreamResponseFunctionCallArgumentsDone struct {
	// Final complete arguments JSON string for the function call
	Arguments string `json:"arguments,required"`
	// Unique identifier of the completed function call
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.function_call_arguments.done"
	Type constant.ResponseFunctionCallArgumentsDone `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments      respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFunctionCallArgumentsDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFunctionCallArgumentsDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for web search calls in progress.
type ResponseObjectStreamResponseWebSearchCallInProgress struct {
	// Unique identifier of the web search call
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.web_search_call.in_progress"
	Type constant.ResponseWebSearchCallInProgress `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseWebSearchCallInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseWebSearchCallInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseWebSearchCallSearching struct {
	ItemID         string                                  `json:"item_id,required"`
	OutputIndex    int64                                   `json:"output_index,required"`
	SequenceNumber int64                                   `json:"sequence_number,required"`
	Type           constant.ResponseWebSearchCallSearching `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseWebSearchCallSearching) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseWebSearchCallSearching) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for completed web search calls.
type ResponseObjectStreamResponseWebSearchCallCompleted struct {
	// Unique identifier of the completed web search call
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.web_search_call.completed"
	Type constant.ResponseWebSearchCallCompleted `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseWebSearchCallCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseWebSearchCallCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpListToolsInProgress struct {
	SequenceNumber int64                                   `json:"sequence_number,required"`
	Type           constant.ResponseMcpListToolsInProgress `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpListToolsInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpListToolsInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpListToolsFailed struct {
	SequenceNumber int64                               `json:"sequence_number,required"`
	Type           constant.ResponseMcpListToolsFailed `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpListToolsFailed) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpListToolsFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpListToolsCompleted struct {
	SequenceNumber int64                                  `json:"sequence_number,required"`
	Type           constant.ResponseMcpListToolsCompleted `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpListToolsCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpListToolsCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpCallArgumentsDelta struct {
	Delta          string                                 `json:"delta,required"`
	ItemID         string                                 `json:"item_id,required"`
	OutputIndex    int64                                  `json:"output_index,required"`
	SequenceNumber int64                                  `json:"sequence_number,required"`
	Type           constant.ResponseMcpCallArgumentsDelta `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta          respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallArgumentsDelta) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallArgumentsDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseMcpCallArgumentsDone struct {
	Arguments      string                                `json:"arguments,required"`
	ItemID         string                                `json:"item_id,required"`
	OutputIndex    int64                                 `json:"output_index,required"`
	SequenceNumber int64                                 `json:"sequence_number,required"`
	Type           constant.ResponseMcpCallArgumentsDone `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments      respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallArgumentsDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallArgumentsDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for MCP calls in progress.
type ResponseObjectStreamResponseMcpCallInProgress struct {
	// Unique identifier of the MCP call
	ItemID string `json:"item_id,required"`
	// Index position of the item in the output list
	OutputIndex int64 `json:"output_index,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.mcp_call.in_progress"
	Type constant.ResponseMcpCallInProgress `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallInProgress) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for failed MCP calls.
type ResponseObjectStreamResponseMcpCallFailed struct {
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.mcp_call.failed"
	Type constant.ResponseMcpCallFailed `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallFailed) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for completed MCP calls.
type ResponseObjectStreamResponseMcpCallCompleted struct {
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.mcp_call.completed"
	Type constant.ResponseMcpCallCompleted `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseMcpCallCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseMcpCallCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a new content part is added to a response item.
type ResponseObjectStreamResponseContentPartAdded struct {
	// Index position of the part within the content array
	ContentIndex int64 `json:"content_index,required"`
	// Unique identifier of the output item containing this content part
	ItemID string `json:"item_id,required"`
	// Index position of the output item in the response
	OutputIndex int64 `json:"output_index,required"`
	// The content part that was added
	Part ResponseObjectStreamResponseContentPartAddedPartUnion `json:"part,required"`
	// Unique identifier of the response containing this content
	ResponseID string `json:"response_id,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.content_part.added"
	Type constant.ResponseContentPartAdded `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Part           respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAdded) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartAddedPartUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseContentPartAddedPartOutputText],
// [ResponseObjectStreamResponseContentPartAddedPartRefusal],
// [ResponseObjectStreamResponseContentPartAddedPartReasoningText].
//
// Use the [ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartAddedPartUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputText].
	Annotations []ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion `json:"annotations"`
	Text        string                                                                      `json:"text"`
	// Any of "output_text", "refusal", "reasoning_text".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputText].
	Logprobs []map[string]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion `json:"logprobs"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Logprobs    respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseContentPartAddedPart is implemented by each
// variant of [ResponseObjectStreamResponseContentPartAddedPartUnion] to add type
// safety for the return type of
// [ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny]
type anyResponseObjectStreamResponseContentPartAddedPart interface {
	implResponseObjectStreamResponseContentPartAddedPartUnion()
}

func (ResponseObjectStreamResponseContentPartAddedPartOutputText) implResponseObjectStreamResponseContentPartAddedPartUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartRefusal) implResponseObjectStreamResponseContentPartAddedPartUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartReasoningText) implResponseObjectStreamResponseContentPartAddedPartUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputText:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartRefusal:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartReasoningText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsAny() anyResponseObjectStreamResponseContentPartAddedPart {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	case "reasoning_text":
		return u.AsReasoningText()
	}
	return nil
}

func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsOutputText() (v ResponseObjectStreamResponseContentPartAddedPartOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsRefusal() (v ResponseObjectStreamResponseContentPartAddedPartRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsReasoningText() (v ResponseObjectStreamResponseContentPartAddedPartReasoningText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartAddedPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseContentPartAddedPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content within a streamed response part.
type ResponseObjectStreamResponseContentPartAddedPartOutputText struct {
	// Structured annotations associated with the text
	Annotations []ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion `json:"annotations,required"`
	// Text emitted for this content part
	Text string `json:"text,required"`
	// Content part type identifier, always "output_text"
	Type constant.OutputText `json:"type,required"`
	// (Optional) Token log probability details
	Logprobs []map[string]ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Logprobs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation],
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation],
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation].
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

// anyResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotation is
// implemented by each variant of
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion] to
// add type safety for the return type of
// [ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotation interface {
	implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion()
}

func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) implResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsAny() anyResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotation {
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

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation struct {
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
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation struct {
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
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation struct {
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
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath struct {
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
func (r ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion contains
// all possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion struct {
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

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseContentPartAddedPartOutputTextLogprobUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectStreamResponseContentPartAddedPartRefusal struct {
	// Refusal text supplied by the model
	Refusal string `json:"refusal,required"`
	// Content part type identifier, always "refusal"
	Type constant.Refusal `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartRefusal) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartAddedPartRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text emitted as part of a streamed response.
type ResponseObjectStreamResponseContentPartAddedPartReasoningText struct {
	// Reasoning text supplied by the model
	Text string `json:"text,required"`
	// Content part type identifier, always "reasoning_text"
	Type constant.ReasoningText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartAddedPartReasoningText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartAddedPartReasoningText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event for when a content part is completed.
type ResponseObjectStreamResponseContentPartDone struct {
	// Index position of the part within the content array
	ContentIndex int64 `json:"content_index,required"`
	// Unique identifier of the output item containing this content part
	ItemID string `json:"item_id,required"`
	// Index position of the output item in the response
	OutputIndex int64 `json:"output_index,required"`
	// The completed content part
	Part ResponseObjectStreamResponseContentPartDonePartUnion `json:"part,required"`
	// Unique identifier of the response containing this content
	ResponseID string `json:"response_id,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.content_part.done"
	Type constant.ResponseContentPartDone `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentIndex   respjson.Field
		ItemID         respjson.Field
		OutputIndex    respjson.Field
		Part           respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDone) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartDonePartUnion contains all possible
// properties and values from
// [ResponseObjectStreamResponseContentPartDonePartOutputText],
// [ResponseObjectStreamResponseContentPartDonePartRefusal],
// [ResponseObjectStreamResponseContentPartDonePartReasoningText].
//
// Use the [ResponseObjectStreamResponseContentPartDonePartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartDonePartUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputText].
	Annotations []ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion `json:"annotations"`
	Text        string                                                                     `json:"text"`
	// Any of "output_text", "refusal", "reasoning_text".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputText].
	Logprobs []map[string]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion `json:"logprobs"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Logprobs    respjson.Field
		Refusal     respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseObjectStreamResponseContentPartDonePart is implemented by each
// variant of [ResponseObjectStreamResponseContentPartDonePartUnion] to add type
// safety for the return type of
// [ResponseObjectStreamResponseContentPartDonePartUnion.AsAny]
type anyResponseObjectStreamResponseContentPartDonePart interface {
	implResponseObjectStreamResponseContentPartDonePartUnion()
}

func (ResponseObjectStreamResponseContentPartDonePartOutputText) implResponseObjectStreamResponseContentPartDonePartUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartRefusal) implResponseObjectStreamResponseContentPartDonePartUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartReasoningText) implResponseObjectStreamResponseContentPartDonePartUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartDonePartUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputText:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartRefusal:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartReasoningText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsAny() anyResponseObjectStreamResponseContentPartDonePart {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
	case "reasoning_text":
		return u.AsReasoningText()
	}
	return nil
}

func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsOutputText() (v ResponseObjectStreamResponseContentPartDonePartOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsRefusal() (v ResponseObjectStreamResponseContentPartDonePartRefusal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsReasoningText() (v ResponseObjectStreamResponseContentPartDonePartReasoningText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartDonePartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseContentPartDonePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content within a streamed response part.
type ResponseObjectStreamResponseContentPartDonePartOutputText struct {
	// Structured annotations associated with the text
	Annotations []ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion `json:"annotations,required"`
	// Text emitted for this content part
	Text string `json:"text,required"`
	// Content part type identifier, always "output_text"
	Type constant.OutputText `json:"type,required"`
	// (Optional) Token log probability details
	Logprobs []map[string]ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Logprobs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion
// contains all possible properties and values from
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation],
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation],
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation],
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath].
//
// Use the
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation].
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

// anyResponseObjectStreamResponseContentPartDonePartOutputTextAnnotation is
// implemented by each variant of
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion] to
// add type safety for the return type of
// [ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion.AsAny]
type anyResponseObjectStreamResponseContentPartDonePartOutputTextAnnotation interface {
	implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion()
}

func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}
func (ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) implResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsAny() anyResponseObjectStreamResponseContentPartDonePartOutputTextAnnotation {
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

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsFileCitation() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsURLCitation() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsContainerFileCitation() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) AsFilePath() (v ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation struct {
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
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation struct {
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
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation struct {
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
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath struct {
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
func (r ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion contains
// all possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion struct {
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

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseContentPartDonePartOutputTextLogprobUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Refusal content within a streamed response part.
type ResponseObjectStreamResponseContentPartDonePartRefusal struct {
	// Refusal text supplied by the model
	Refusal string `json:"refusal,required"`
	// Content part type identifier, always "refusal"
	Type constant.Refusal `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refusal     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartRefusal) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseContentPartDonePartRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning text emitted as part of a streamed response.
type ResponseObjectStreamResponseContentPartDonePartReasoningText struct {
	// Reasoning text supplied by the model
	Text string `json:"text,required"`
	// Content part type identifier, always "reasoning_text"
	Type constant.ReasoningText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseContentPartDonePartReasoningText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseObjectStreamResponseContentPartDonePartReasoningText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event emitted when a response ends in an incomplete state.
type ResponseObjectStreamResponseIncomplete struct {
	// Response object describing the incomplete state
	Response ResponseObject `json:"response,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.incomplete"
	Type constant.ResponseIncomplete `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response       respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseIncomplete) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseIncomplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event emitted when a response fails.
type ResponseObjectStreamResponseFailed struct {
	// Response object describing the failure
	Response ResponseObject `json:"response,required"`
	// Sequential number for ordering streaming events
	SequenceNumber int64 `json:"sequence_number,required"`
	// Event type identifier, always "response.failed"
	Type constant.ResponseFailed `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response       respjson.Field
		SequenceNumber respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseFailed) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating a response has been completed.
type ResponseObjectStreamResponseCompleted struct {
	// Completed response object
	Response ResponseObject `json:"response,required"`
	// Event type identifier, always "response.completed"
	Type constant.ResponseCompleted `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObjectStreamResponseCompleted) RawJSON() string { return r.JSON.raw }
func (r *ResponseObjectStreamResponseCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI response object extended with input context information.
type ResponseListResponse struct {
	// Unique identifier for this response
	ID string `json:"id,required"`
	// Unix timestamp when the response was created
	CreatedAt int64 `json:"created_at,required"`
	// List of input items that led to this response
	Input []ResponseListResponseInputUnion `json:"input,required"`
	// Model identifier used for generation
	Model string `json:"model,required"`
	// Object type identifier, always "response"
	Object constant.Response `json:"object,required"`
	// List of generated output items (messages, tool calls, etc.)
	Output []ResponseListResponseOutputUnion `json:"output,required"`
	// Whether tool calls can be executed in parallel
	ParallelToolCalls bool `json:"parallel_tool_calls,required"`
	// Current status of the response generation
	Status string `json:"status,required"`
	// Text formatting configuration for the response
	Text ResponseListResponseText `json:"text,required"`
	// (Optional) Error details if the response generation failed
	Error ResponseListResponseError `json:"error"`
	// (Optional) ID of the previous response in a conversation
	PreviousResponseID string `json:"previous_response_id"`
	// (Optional) Sampling temperature used for generation
	Temperature float64 `json:"temperature"`
	// (Optional) An array of tools the model may call while generating a response.
	Tools []ResponseListResponseToolUnion `json:"tools"`
	// (Optional) Nucleus sampling parameter used for generation
	TopP float64 `json:"top_p"`
	// (Optional) Truncation strategy applied to the response
	Truncation string `json:"truncation"`
	// (Optional) Token usage information for the response
	Usage ResponseListResponseUsage `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Input              respjson.Field
		Model              respjson.Field
		Object             respjson.Field
		Output             respjson.Field
		ParallelToolCalls  respjson.Field
		Status             respjson.Field
		Text               respjson.Field
		Error              respjson.Field
		PreviousResponseID respjson.Field
		Temperature        respjson.Field
		Tools              respjson.Field
		TopP               respjson.Field
		Truncation         respjson.Field
		Usage              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputUnion contains all possible properties and values from
// [ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput],
// [ResponseListResponseInputOpenAIResponseMcpApprovalRequest],
// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse],
// [ResponseListResponseInputOpenAIResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputUnion struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall].
	Results   []ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	Arguments string                                                                         `json:"arguments"`
	CallID    string                                                                         `json:"call_id"`
	Name      string                                                                         `json:"name"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalRequest].
	ServerLabel string `json:"server_label"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	// This field is from variant [ResponseListResponseInputOpenAIResponseMessage].
	Content ResponseListResponseInputOpenAIResponseMessageContentUnion `json:"content"`
	// This field is from variant [ResponseListResponseInputOpenAIResponseMessage].
	Role ResponseListResponseInputOpenAIResponseMessageRole `json:"role"`
	JSON struct {
		ID                respjson.Field
		Status            respjson.Field
		Type              respjson.Field
		Queries           respjson.Field
		Results           respjson.Field
		Arguments         respjson.Field
		CallID            respjson.Field
		Name              respjson.Field
		Output            respjson.Field
		ServerLabel       respjson.Field
		ApprovalRequestID respjson.Field
		Approve           respjson.Field
		Reason            respjson.Field
		Content           respjson.Field
		Role              respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseMcpApprovalRequest() (v ResponseListResponseInputOpenAIResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseMcpApprovalResponse() (v ResponseListResponseInputOpenAIResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputUnion) AsOpenAIResponseMessage() (v ResponseListResponseInputOpenAIResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall struct {
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
func (r ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
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
func (r ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion struct {
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

func (u ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall struct {
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
func (r ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput struct {
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
func (r ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseListResponseInputOpenAIResponseMcpApprovalRequest struct {
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
func (r ResponseListResponseInputOpenAIResponseMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ResponseListResponseInputOpenAIResponseMcpApprovalResponse struct {
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
func (r ResponseListResponseInputOpenAIResponseMcpApprovalResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseListResponseInputOpenAIResponseMessage struct {
	Content ResponseListResponseInputOpenAIResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseListResponseInputOpenAIResponseMessageRole `json:"role,required"`
	Type   constant.Message                                   `json:"type,required"`
	ID     string                                             `json:"id"`
	Status string                                             `json:"status"`
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
func (r ResponseListResponseInputOpenAIResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseInputOpenAIResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageContentUnion contains all possible
// properties and values from [string],
// [[]ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion],
// [[]ResponseListResponseInputOpenAIResponseMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseListResponseInputOpenAIResponseMessageContentArray OfVariant2]
type ResponseListResponseInputOpenAIResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion] instead
	// of an object.
	OfResponseListResponseInputOpenAIResponseMessageContentArray []ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseInputOpenAIResponseMessageContentArrayItem] instead of an
	// object.
	OfVariant2 []ResponseListResponseInputOpenAIResponseMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                                     respjson.Field
		OfResponseListResponseInputOpenAIResponseMessageContentArray respjson.Field
		OfVariant2                                                   respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (u ResponseListResponseInputOpenAIResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageContentUnion) AsResponseListResponseInputOpenAIResponseMessageContentArray() (v []ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageContentUnion) AsVariant2() (v []ResponseListResponseInputOpenAIResponseMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion contains all
// possible properties and values from
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText],
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage].
//
// Use the
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage].
	Detail ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponseInputOpenAIResponseMessageContentArrayItem is implemented
// by each variant of
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion.AsAny]
type anyResponseListResponseInputOpenAIResponseMessageContentArrayItem interface {
	implResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion()
}

func (ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText) implResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage) implResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion) AsAny() anyResponseListResponseInputOpenAIResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion) AsInputText() (v ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion) AsInputImage() (v ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText struct {
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetail string

const (
	ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetailLow  ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetail = "low"
	ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetailHigh ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetail = "high"
	ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetailAuto ResponseListResponseInputOpenAIResponseMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetail string

const (
	ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetailLow  ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetail = "low"
	ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetailHigh ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetail = "high"
	ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetailAuto ResponseListResponseInputOpenAIResponseMessageContentArrayItemDetail = "auto"
)

type ResponseListResponseInputOpenAIResponseMessageContentArrayItem struct {
	Annotations []ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                          `json:"text,required"`
	Type        constant.OutputText                                                             `json:"type,required"`
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion
// contains all possible properties and values from
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation],
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation],
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotation is
// implemented by each variant of
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion]
// to add type safety for the return type of
// [ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotation interface {
	implResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion()
}

func (ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) implResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) implResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) implResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) implResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsAny() anyResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotation {
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

func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseInputOpenAIResponseMessageRole string

const (
	ResponseListResponseInputOpenAIResponseMessageRoleSystem    ResponseListResponseInputOpenAIResponseMessageRole = "system"
	ResponseListResponseInputOpenAIResponseMessageRoleDeveloper ResponseListResponseInputOpenAIResponseMessageRole = "developer"
	ResponseListResponseInputOpenAIResponseMessageRoleUser      ResponseListResponseInputOpenAIResponseMessageRole = "user"
	ResponseListResponseInputOpenAIResponseMessageRoleAssistant ResponseListResponseInputOpenAIResponseMessageRole = "assistant"
)

type ResponseListResponseInputRole string

const (
	ResponseListResponseInputRoleSystem    ResponseListResponseInputRole = "system"
	ResponseListResponseInputRoleDeveloper ResponseListResponseInputRole = "developer"
	ResponseListResponseInputRoleUser      ResponseListResponseInputRole = "user"
	ResponseListResponseInputRoleAssistant ResponseListResponseInputRole = "assistant"
)

// ResponseListResponseOutputUnion contains all possible properties and values from
// [ResponseListResponseOutputMessage], [ResponseListResponseOutputWebSearchCall],
// [ResponseListResponseOutputFileSearchCall],
// [ResponseListResponseOutputFunctionCall], [ResponseListResponseOutputMcpCall],
// [ResponseListResponseOutputMcpListTools],
// [ResponseListResponseOutputMcpApprovalRequest].
//
// Use the [ResponseListResponseOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputUnion struct {
	// This field is from variant [ResponseListResponseOutputMessage].
	Content ResponseListResponseOutputMessageContentUnion `json:"content"`
	// This field is from variant [ResponseListResponseOutputMessage].
	Role ResponseListResponseOutputMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [ResponseListResponseOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseListResponseOutputFileSearchCall].
	Results   []ResponseListResponseOutputFileSearchCallResult `json:"results"`
	Arguments string                                           `json:"arguments"`
	// This field is from variant [ResponseListResponseOutputFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseListResponseOutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [ResponseListResponseOutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [ResponseListResponseOutputMcpListTools].
	Tools []ResponseListResponseOutputMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseListResponseOutput is implemented by each variant of
// [ResponseListResponseOutputUnion] to add type safety for the return type of
// [ResponseListResponseOutputUnion.AsAny]
type anyResponseListResponseOutput interface {
	implResponseListResponseOutputUnion()
}

func (ResponseListResponseOutputMessage) implResponseListResponseOutputUnion()            {}
func (ResponseListResponseOutputWebSearchCall) implResponseListResponseOutputUnion()      {}
func (ResponseListResponseOutputFileSearchCall) implResponseListResponseOutputUnion()     {}
func (ResponseListResponseOutputFunctionCall) implResponseListResponseOutputUnion()       {}
func (ResponseListResponseOutputMcpCall) implResponseListResponseOutputUnion()            {}
func (ResponseListResponseOutputMcpListTools) implResponseListResponseOutputUnion()       {}
func (ResponseListResponseOutputMcpApprovalRequest) implResponseListResponseOutputUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessage:
//	case llamastackclient.ResponseListResponseOutputWebSearchCall:
//	case llamastackclient.ResponseListResponseOutputFileSearchCall:
//	case llamastackclient.ResponseListResponseOutputFunctionCall:
//	case llamastackclient.ResponseListResponseOutputMcpCall:
//	case llamastackclient.ResponseListResponseOutputMcpListTools:
//	case llamastackclient.ResponseListResponseOutputMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputUnion) AsAny() anyResponseListResponseOutput {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "function_call":
		return u.AsFunctionCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	case "mcp_approval_request":
		return u.AsMcpApprovalRequest()
	}
	return nil
}

func (u ResponseListResponseOutputUnion) AsMessage() (v ResponseListResponseOutputMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsWebSearchCall() (v ResponseListResponseOutputWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsFileSearchCall() (v ResponseListResponseOutputFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsFunctionCall() (v ResponseListResponseOutputFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsMcpCall() (v ResponseListResponseOutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsMcpListTools() (v ResponseListResponseOutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputUnion) AsMcpApprovalRequest() (v ResponseListResponseOutputMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseListResponseOutputMessage struct {
	Content ResponseListResponseOutputMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseListResponseOutputMessageRole `json:"role,required"`
	Type   constant.Message                      `json:"type,required"`
	ID     string                                `json:"id"`
	Status string                                `json:"status"`
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
func (r ResponseListResponseOutputMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentUnion contains all possible properties
// and values from [string],
// [[]ResponseListResponseOutputMessageContentArrayItemUnion],
// [[]ResponseListResponseOutputMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfResponseListResponseOutputMessageContentArray
// OfVariant2]
type ResponseListResponseOutputMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseOutputMessageContentArrayItemUnion] instead of an object.
	OfResponseListResponseOutputMessageContentArray []ResponseListResponseOutputMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseOutputMessageContentArrayItem] instead of an object.
	OfVariant2 []ResponseListResponseOutputMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                        respjson.Field
		OfResponseListResponseOutputMessageContentArray respjson.Field
		OfVariant2                                      respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u ResponseListResponseOutputMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentUnion) AsResponseListResponseOutputMessageContentArray() (v []ResponseListResponseOutputMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentUnion) AsVariant2() (v []ResponseListResponseOutputMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseOutputMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentArrayItemUnion contains all possible
// properties and values from
// [ResponseListResponseOutputMessageContentArrayItemInputText],
// [ResponseListResponseOutputMessageContentArrayItemInputImage].
//
// Use the [ResponseListResponseOutputMessageContentArrayItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseListResponseOutputMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentArrayItemInputImage].
	Detail ResponseListResponseOutputMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponseOutputMessageContentArrayItem is implemented by each
// variant of [ResponseListResponseOutputMessageContentArrayItemUnion] to add type
// safety for the return type of
// [ResponseListResponseOutputMessageContentArrayItemUnion.AsAny]
type anyResponseListResponseOutputMessageContentArrayItem interface {
	implResponseListResponseOutputMessageContentArrayItemUnion()
}

func (ResponseListResponseOutputMessageContentArrayItemInputText) implResponseListResponseOutputMessageContentArrayItemUnion() {
}
func (ResponseListResponseOutputMessageContentArrayItemInputImage) implResponseListResponseOutputMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessageContentArrayItemInputText:
//	case llamastackclient.ResponseListResponseOutputMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputMessageContentArrayItemUnion) AsAny() anyResponseListResponseOutputMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseListResponseOutputMessageContentArrayItemUnion) AsInputText() (v ResponseListResponseOutputMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentArrayItemUnion) AsInputImage() (v ResponseListResponseOutputMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseOutputMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseListResponseOutputMessageContentArrayItemInputText struct {
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
func (r ResponseListResponseOutputMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseListResponseOutputMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseListResponseOutputMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseListResponseOutputMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseListResponseOutputMessageContentArrayItemInputImageDetail string

const (
	ResponseListResponseOutputMessageContentArrayItemInputImageDetailLow  ResponseListResponseOutputMessageContentArrayItemInputImageDetail = "low"
	ResponseListResponseOutputMessageContentArrayItemInputImageDetailHigh ResponseListResponseOutputMessageContentArrayItemInputImageDetail = "high"
	ResponseListResponseOutputMessageContentArrayItemInputImageDetailAuto ResponseListResponseOutputMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseListResponseOutputMessageContentArrayItemDetail string

const (
	ResponseListResponseOutputMessageContentArrayItemDetailLow  ResponseListResponseOutputMessageContentArrayItemDetail = "low"
	ResponseListResponseOutputMessageContentArrayItemDetailHigh ResponseListResponseOutputMessageContentArrayItemDetail = "high"
	ResponseListResponseOutputMessageContentArrayItemDetailAuto ResponseListResponseOutputMessageContentArrayItemDetail = "auto"
)

type ResponseListResponseOutputMessageContentArrayItem struct {
	Annotations []ResponseListResponseOutputMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                             `json:"text,required"`
	Type        constant.OutputText                                                `json:"type,required"`
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
func (r ResponseListResponseOutputMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMessageContentArrayItemAnnotationUnion contains all
// possible properties and values from
// [ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation],
// [ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation],
// [ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath].
//
// Use the [ResponseListResponseOutputMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseOutputMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseListResponseOutputMessageContentArrayItemAnnotation is implemented by
// each variant of
// [ResponseListResponseOutputMessageContentArrayItemAnnotationUnion] to add type
// safety for the return type of
// [ResponseListResponseOutputMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseListResponseOutputMessageContentArrayItemAnnotation interface {
	implResponseListResponseOutputMessageContentArrayItemAnnotationUnion()
}

func (ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation) implResponseListResponseOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation) implResponseListResponseOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation) implResponseListResponseOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath) implResponseListResponseOutputMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseOutputMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) AsAny() anyResponseListResponseOutputMessageContentArrayItemAnnotation {
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

func (u ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseOutputMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseOutputMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputMessageRole string

const (
	ResponseListResponseOutputMessageRoleSystem    ResponseListResponseOutputMessageRole = "system"
	ResponseListResponseOutputMessageRoleDeveloper ResponseListResponseOutputMessageRole = "developer"
	ResponseListResponseOutputMessageRoleUser      ResponseListResponseOutputMessageRole = "user"
	ResponseListResponseOutputMessageRoleAssistant ResponseListResponseOutputMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ResponseListResponseOutputWebSearchCall struct {
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
func (r ResponseListResponseOutputWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseListResponseOutputFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseListResponseOutputFileSearchCallResult `json:"results"`
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
func (r ResponseListResponseOutputFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseListResponseOutputFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseListResponseOutputFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseListResponseOutputFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseOutputFileSearchCallResultAttributeUnion struct {
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

func (u ResponseListResponseOutputFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseOutputFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseListResponseOutputFunctionCall struct {
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
func (r ResponseListResponseOutputFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseListResponseOutputMcpCall struct {
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
func (r ResponseListResponseOutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseListResponseOutputMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ResponseListResponseOutputMcpListToolsTool `json:"tools,required"`
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
func (r ResponseListResponseOutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseListResponseOutputMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ResponseListResponseOutputMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ResponseListResponseOutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseOutputMcpListToolsToolInputSchemaUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseOutputMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseListResponseOutputMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseOutputMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseOutputMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseOutputMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseListResponseOutputMcpApprovalRequest struct {
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
func (r ResponseListResponseOutputMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseOutputMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseOutputRole string

const (
	ResponseListResponseOutputRoleSystem    ResponseListResponseOutputRole = "system"
	ResponseListResponseOutputRoleDeveloper ResponseListResponseOutputRole = "developer"
	ResponseListResponseOutputRoleUser      ResponseListResponseOutputRole = "user"
	ResponseListResponseOutputRoleAssistant ResponseListResponseOutputRole = "assistant"
)

// Text formatting configuration for the response
type ResponseListResponseText struct {
	// (Optional) Text format configuration specifying output format requirements
	Format ResponseListResponseTextFormat `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseText) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Text format configuration specifying output format requirements
type ResponseListResponseTextFormat struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseListResponseTextFormatType `json:"type,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description string `json:"description"`
	// The name of the response format. Only used for json_schema.
	Name string `json:"name"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseListResponseTextFormatSchemaUnion `json:"schema"`
	// (Optional) Whether to strictly enforce the JSON schema. If true, the response
	// must match the schema exactly. Only used for json_schema.
	Strict bool `json:"strict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Schema      respjson.Field
		Strict      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "text", "json_schema", or "json_object" to identify the format type
type ResponseListResponseTextFormatType string

const (
	ResponseListResponseTextFormatTypeText       ResponseListResponseTextFormatType = "text"
	ResponseListResponseTextFormatTypeJsonSchema ResponseListResponseTextFormatType = "json_schema"
	ResponseListResponseTextFormatTypeJsonObject ResponseListResponseTextFormatType = "json_object"
)

// ResponseListResponseTextFormatSchemaUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseTextFormatSchemaUnion struct {
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

func (u ResponseListResponseTextFormatSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseTextFormatSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseTextFormatSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseTextFormatSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseTextFormatSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseTextFormatSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Error details if the response generation failed
type ResponseListResponseError struct {
	// Error code identifying the type of failure
	Code string `json:"code,required"`
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
func (r ResponseListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolUnion contains all possible properties and values from
// [ResponseListResponseToolOpenAIResponseInputToolWebSearch],
// [ResponseListResponseToolFileSearch], [ResponseListResponseToolFunction],
// [ResponseListResponseToolMcp].
//
// Use the [ResponseListResponseToolUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseToolUnion struct {
	// Any of nil, "file_search", "function", "mcp".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseToolOpenAIResponseInputToolWebSearch].
	SearchContextSize string `json:"search_context_size"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	VectorStoreIDs []string `json:"vector_store_ids"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	Filters map[string]ResponseListResponseToolFileSearchFilterUnion `json:"filters"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	MaxNumResults int64 `json:"max_num_results"`
	// This field is from variant [ResponseListResponseToolFileSearch].
	RankingOptions ResponseListResponseToolFileSearchRankingOptions `json:"ranking_options"`
	// This field is from variant [ResponseListResponseToolFunction].
	Name string `json:"name"`
	// This field is from variant [ResponseListResponseToolFunction].
	Description string `json:"description"`
	// This field is from variant [ResponseListResponseToolFunction].
	Parameters map[string]ResponseListResponseToolFunctionParameterUnion `json:"parameters"`
	// This field is from variant [ResponseListResponseToolFunction].
	Strict bool `json:"strict"`
	// This field is from variant [ResponseListResponseToolMcp].
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseListResponseToolMcp].
	AllowedTools ResponseListResponseToolMcpAllowedToolsUnion `json:"allowed_tools"`
	JSON         struct {
		Type              respjson.Field
		SearchContextSize respjson.Field
		VectorStoreIDs    respjson.Field
		Filters           respjson.Field
		MaxNumResults     respjson.Field
		RankingOptions    respjson.Field
		Name              respjson.Field
		Description       respjson.Field
		Parameters        respjson.Field
		Strict            respjson.Field
		ServerLabel       respjson.Field
		AllowedTools      respjson.Field
		raw               string
	} `json:"-"`
}

func (u ResponseListResponseToolUnion) AsOpenAIResponseInputToolWebSearch() (v ResponseListResponseToolOpenAIResponseInputToolWebSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolUnion) AsFileSearch() (v ResponseListResponseToolFileSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolUnion) AsFunction() (v ResponseListResponseToolFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolUnion) AsMcp() (v ResponseListResponseToolMcp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool configuration for OpenAI response inputs.
type ResponseListResponseToolOpenAIResponseInputToolWebSearch struct {
	// Web search tool type variant to use
	//
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11".
	Type ResponseListResponseToolOpenAIResponseInputToolWebSearchType `json:"type,required"`
	// (Optional) Size of search context, must be "low", "medium", or "high"
	SearchContextSize string `json:"search_context_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		SearchContextSize respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolOpenAIResponseInputToolWebSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool type variant to use
type ResponseListResponseToolOpenAIResponseInputToolWebSearchType string

const (
	ResponseListResponseToolOpenAIResponseInputToolWebSearchTypeWebSearch                  ResponseListResponseToolOpenAIResponseInputToolWebSearchType = "web_search"
	ResponseListResponseToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview           ResponseListResponseToolOpenAIResponseInputToolWebSearchType = "web_search_preview"
	ResponseListResponseToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview2025_03_11 ResponseListResponseToolOpenAIResponseInputToolWebSearchType = "web_search_preview_2025_03_11"
)

// File search tool configuration for OpenAI response inputs.
type ResponseListResponseToolFileSearch struct {
	// Tool type identifier, always "file_search"
	Type constant.FileSearch `json:"type,required"`
	// List of vector store identifiers to search within
	VectorStoreIDs []string `json:"vector_store_ids,required"`
	// (Optional) Additional filters to apply to the search
	Filters map[string]ResponseListResponseToolFileSearchFilterUnion `json:"filters"`
	// (Optional) Maximum number of search results to return (1-50)
	MaxNumResults int64 `json:"max_num_results"`
	// (Optional) Options for ranking and scoring search results
	RankingOptions ResponseListResponseToolFileSearchRankingOptions `json:"ranking_options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type           respjson.Field
		VectorStoreIDs respjson.Field
		Filters        respjson.Field
		MaxNumResults  respjson.Field
		RankingOptions respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolFileSearch) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolFileSearchFilterUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseToolFileSearchFilterUnion struct {
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

func (u ResponseListResponseToolFileSearchFilterUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolFileSearchFilterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolFileSearchFilterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolFileSearchFilterUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolFileSearchFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolFileSearchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Options for ranking and scoring search results
type ResponseListResponseToolFileSearchRankingOptions struct {
	// (Optional) Name of the ranking algorithm to use
	Ranker string `json:"ranker"`
	// (Optional) Minimum relevance score threshold for results
	ScoreThreshold float64 `json:"score_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ranker         respjson.Field
		ScoreThreshold respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolFileSearchRankingOptions) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool configuration for OpenAI response inputs.
type ResponseListResponseToolFunction struct {
	// Name of the function that can be called
	Name string `json:"name,required"`
	// Tool type identifier, always "function"
	Type constant.Function `json:"type,required"`
	// (Optional) Description of what the function does
	Description string `json:"description"`
	// (Optional) JSON schema defining the function's parameters
	Parameters map[string]ResponseListResponseToolFunctionParameterUnion `json:"parameters"`
	// (Optional) Whether to enforce strict parameter validation
	Strict bool `json:"strict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Parameters  respjson.Field
		Strict      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolFunction) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolFunctionParameterUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseToolFunctionParameterUnion struct {
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

func (u ResponseListResponseToolFunctionParameterUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolFunctionParameterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolFunctionParameterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolFunctionParameterUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolFunctionParameterUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolFunctionParameterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) tool configuration for OpenAI response object.
type ResponseListResponseToolMcp struct {
	// Label to identify this MCP server
	ServerLabel string `json:"server_label,required"`
	// Tool type identifier, always "mcp"
	Type constant.Mcp `json:"type,required"`
	// (Optional) Restriction on which tools can be used from this server
	AllowedTools ResponseListResponseToolMcpAllowedToolsUnion `json:"allowed_tools"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServerLabel  respjson.Field
		Type         respjson.Field
		AllowedTools respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolMcp) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseToolMcpAllowedToolsUnion contains all possible properties
// and values from [[]string],
// [ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfStringArray]
type ResponseListResponseToolMcpAllowedToolsUnion struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant
	// [ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter].
	ToolNames []string `json:"tool_names"`
	JSON      struct {
		OfStringArray respjson.Field
		ToolNames     respjson.Field
		raw           string
	} `json:"-"`
}

func (u ResponseListResponseToolMcpAllowedToolsUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseToolMcpAllowedToolsUnion) AsAllowedToolsFilter() (v ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseToolMcpAllowedToolsUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter configuration for restricting which MCP tools can be used.
type ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter struct {
	// (Optional) List of specific tool names that are allowed
	ToolNames []string `json:"tool_names"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolNames   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool type variant to use
type ResponseListResponseToolType string

const (
	ResponseListResponseToolTypeWebSearch                  ResponseListResponseToolType = "web_search"
	ResponseListResponseToolTypeWebSearchPreview           ResponseListResponseToolType = "web_search_preview"
	ResponseListResponseToolTypeWebSearchPreview2025_03_11 ResponseListResponseToolType = "web_search_preview_2025_03_11"
	ResponseListResponseToolTypeFileSearch                 ResponseListResponseToolType = "file_search"
	ResponseListResponseToolTypeFunction                   ResponseListResponseToolType = "function"
	ResponseListResponseToolTypeMcp                        ResponseListResponseToolType = "mcp"
)

// (Optional) Token usage information for the response
type ResponseListResponseUsage struct {
	// Number of tokens in the input
	InputTokens int64 `json:"input_tokens,required"`
	// Number of tokens in the output
	OutputTokens int64 `json:"output_tokens,required"`
	// Total tokens used (input + output)
	TotalTokens int64 `json:"total_tokens,required"`
	// Detailed breakdown of input token usage
	InputTokensDetails ResponseListResponseUsageInputTokensDetails `json:"input_tokens_details"`
	// Detailed breakdown of output token usage
	OutputTokensDetails ResponseListResponseUsageOutputTokensDetails `json:"output_tokens_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens         respjson.Field
		OutputTokens        respjson.Field
		TotalTokens         respjson.Field
		InputTokensDetails  respjson.Field
		OutputTokensDetails respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of input token usage
type ResponseListResponseUsageInputTokensDetails struct {
	// Number of tokens retrieved from cache
	CachedTokens int64 `json:"cached_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseUsageInputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseUsageInputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of output token usage
type ResponseListResponseUsageOutputTokensDetails struct {
	// Number of tokens used for reasoning (o1/o3 models)
	ReasoningTokens int64 `json:"reasoning_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasoningTokens respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseUsageOutputTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseUsageOutputTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object confirming deletion of an OpenAI response.
type ResponseDeleteResponse struct {
	// Unique identifier of the deleted response
	ID string `json:"id,required"`
	// Deletion confirmation flag, always True
	Deleted bool `json:"deleted,required"`
	// Object type identifier, always "response"
	Object constant.Response `json:"object,required"`
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
func (r ResponseDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewParams struct {
	// Input message(s) to create the response.
	Input ResponseNewParamsInputUnion `json:"input,omitzero,required"`
	// The underlying LLM used for completions.
	Model string `json:"model,required"`
	// (Optional) The ID of a conversation to add the response to. Must begin with
	// 'conv\_'. Input and output messages will be automatically added to the
	// conversation.
	Conversation  param.Opt[string] `json:"conversation,omitzero"`
	Instructions  param.Opt[string] `json:"instructions,omitzero"`
	MaxInferIters param.Opt[int64]  `json:"max_infer_iters,omitzero"`
	// (Optional) if specified, the new response will be a continuation of the previous
	// response. This can be used to easily fork-off new responses from existing
	// responses.
	PreviousResponseID param.Opt[string]  `json:"previous_response_id,omitzero"`
	Store              param.Opt[bool]    `json:"store,omitzero"`
	Temperature        param.Opt[float64] `json:"temperature,omitzero"`
	// (Optional) Additional fields to include in the response.
	Include []string `json:"include,omitzero"`
	// Text response configuration for OpenAI responses.
	Text  ResponseNewParamsText        `json:"text,omitzero"`
	Tools []ResponseNewParamsToolUnion `json:"tools,omitzero"`
	paramObj
}

func (r ResponseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputUnion struct {
	OfString                 param.Opt[string]                      `json:",omitzero,inline"`
	OfResponseNewsInputArray []ResponseNewParamsInputArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfResponseNewsInputArray)
}
func (u *ResponseNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfResponseNewsInputArray) {
		return &u.OfResponseNewsInputArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputArrayItemUnion struct {
	OfOpenAIResponseOutputMessageWebSearchToolCall  *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageWebSearchToolCall  `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageFileSearchToolCall *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCall `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageFunctionToolCall   *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFunctionToolCall   `json:",omitzero,inline"`
	OfOpenAIResponseInputFunctionToolCallOutput     *ResponseNewParamsInputArrayItemOpenAIResponseInputFunctionToolCallOutput     `json:",omitzero,inline"`
	OfOpenAIResponseMcpApprovalRequest              *ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalRequest              `json:",omitzero,inline"`
	OfOpenAIResponseMcpApprovalResponse             *ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalResponse             `json:",omitzero,inline"`
	OfOpenAIResponseMessage                         *ResponseNewParamsInputArrayItemOpenAIResponseMessage                         `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseOutputMessageWebSearchToolCall,
		u.OfOpenAIResponseOutputMessageFileSearchToolCall,
		u.OfOpenAIResponseOutputMessageFunctionToolCall,
		u.OfOpenAIResponseInputFunctionToolCallOutput,
		u.OfOpenAIResponseMcpApprovalRequest,
		u.OfOpenAIResponseMcpApprovalResponse,
		u.OfOpenAIResponseMessage)
}
func (u *ResponseNewParamsInputArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseOutputMessageWebSearchToolCall) {
		return u.OfOpenAIResponseOutputMessageWebSearchToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageFileSearchToolCall) {
		return u.OfOpenAIResponseOutputMessageFileSearchToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageFunctionToolCall) {
		return u.OfOpenAIResponseOutputMessageFunctionToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseInputFunctionToolCallOutput) {
		return u.OfOpenAIResponseInputFunctionToolCallOutput
	} else if !param.IsOmitted(u.OfOpenAIResponseMcpApprovalRequest) {
		return u.OfOpenAIResponseMcpApprovalRequest
	} else if !param.IsOmitted(u.OfOpenAIResponseMcpApprovalResponse) {
		return u.OfOpenAIResponseMcpApprovalResponse
	} else if !param.IsOmitted(u.OfOpenAIResponseMessage) {
		return u.OfOpenAIResponseMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetQueries() []string {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetResults() []ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResult {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetOutput() *string {
	if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return &vt.Output
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetServerLabel() *string {
	if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return &vt.ServerLabel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetApprovalRequestID() *string {
	if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil {
		return &vt.ApprovalRequestID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetApprove() *bool {
	if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil {
		return &vt.Approve
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetReason() *string {
	if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetContent() *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentUnion {
	if vt := u.OfOpenAIResponseMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetRole() *string {
	if vt := u.OfOpenAIResponseMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetID() *string {
	if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseMessage; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetStatus() *string {
	if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfOpenAIResponseMessage; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetType() *string {
	if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseMcpApprovalResponse; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseMessage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetArguments() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Arguments)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Arguments)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetCallID() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.CallID)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.CallID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetName() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfOpenAIResponseMcpApprovalRequest; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Web search tool call output message for OpenAI responses.
//
// The properties ID, Status, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageWebSearchToolCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// Current status of the web search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "web_search_call"
	//
	// This field can be elided, and will marshal its zero value as "web_search_call".
	Type constant.WebSearchCall `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageWebSearchToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageWebSearchToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,omitzero,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results,omitzero"`
	// Tool call type identifier, always "file_search_call"
	//
	// This field can be elided, and will marshal its zero value as "file_search_call".
	Type constant.FileSearchCall `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
//
// The properties Attributes, FileID, Filename, Score, Text are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion `json:"attributes,omitzero,required"`
	// Unique identifier of the file containing the result
	FileID string `json:"file_id,required"`
	// Name of the file containing the result
	Filename string `json:"filename,required"`
	// Relevance score for this search result (between 0 and 1)
	Score float64 `json:"score,required"`
	// Text content of the search result
	Text string `json:"text,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResult) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) asAny() any {
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

// Function tool call output message for OpenAI responses.
//
// The properties Arguments, CallID, Name, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFunctionToolCall struct {
	// JSON string containing the function arguments
	Arguments string `json:"arguments,required"`
	// Unique identifier for the function call
	CallID string `json:"call_id,required"`
	// Name of the function being called
	Name string `json:"name,required"`
	// (Optional) Additional identifier for the tool call
	ID param.Opt[string] `json:"id,omitzero"`
	// (Optional) Current status of the function call execution
	Status param.Opt[string] `json:"status,omitzero"`
	// Tool call type identifier, always "function_call"
	//
	// This field can be elided, and will marshal its zero value as "function_call".
	Type constant.FunctionCall `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFunctionToolCall) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFunctionToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
//
// The properties CallID, Output, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseInputFunctionToolCallOutput struct {
	CallID string            `json:"call_id,required"`
	Output string            `json:"output,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "function_call_output".
	Type constant.FunctionCallOutput `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseInputFunctionToolCallOutput) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseInputFunctionToolCallOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
//
// The properties ID, Arguments, Name, ServerLabel, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalRequest struct {
	ID          string `json:"id,required"`
	Arguments   string `json:"arguments,required"`
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	// This field can be elided, and will marshal its zero value as
	// "mcp_approval_request".
	Type constant.McpApprovalRequest `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalRequest) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
//
// The properties ApprovalRequestID, Approve, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalResponse struct {
	ApprovalRequestID string            `json:"approval_request_id,required"`
	Approve           bool              `json:"approve,required"`
	ID                param.Opt[string] `json:"id,omitzero"`
	Reason            param.Opt[string] `json:"reason,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "mcp_approval_response".
	Type constant.McpApprovalResponse `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalResponse) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessage struct {
	Content ResponseNewParamsInputArrayItemOpenAIResponseMessageContentUnion `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseNewParamsInputArrayItemOpenAIResponseMessageRole `json:"role,omitzero,required"`
	ID     param.Opt[string]                                        `json:"id,omitzero"`
	Status param.Opt[string]                                        `json:"status,omitzero"`
	// This field can be elided, and will marshal its zero value as "message".
	Type constant.Message `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessage) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentUnion struct {
	OfString                                                      param.Opt[string]                                                           `json:",omitzero,inline"`
	OfResponseNewsInputArrayItemOpenAIResponseMessageContentArray []ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion `json:",omitzero,inline"`
	OfVariant2                                                    []ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItem      `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfResponseNewsInputArrayItemOpenAIResponseMessageContentArray, u.OfVariant2)
}
func (u *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfResponseNewsInputArrayItemOpenAIResponseMessageContentArray) {
		return &u.OfResponseNewsInputArrayItemOpenAIResponseMessageContentArray
	} else if !param.IsOmitted(u.OfVariant2) {
		return &u.OfVariant2
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion struct {
	OfInputText  *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText  `json:",omitzero,inline"`
	OfInputImage *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage)
}
func (u *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfInputText) {
		return u.OfInputText
	} else if !param.IsOmitted(u.OfInputImage) {
		return u.OfInputImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Detail)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion) GetType() *string {
	if vt := u.OfInputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText]("input_text"),
		apijson.Discriminator[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage]("input_image"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The properties Text, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	//
	// This field can be elided, and will marshal its zero value as "input_text".
	Type constant.InputText `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
//
// The properties Detail, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail,omitzero,required"`
	// (Optional) URL of the image content
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Content type identifier, always "input_image"
	//
	// This field can be elided, and will marshal its zero value as "input_image".
	Type constant.InputImage `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetail string

const (
	ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetailLow  ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetail = "low"
	ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetailHigh ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetail = "high"
	ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetailAuto ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetail = "auto"
)

// The properties Annotations, Text, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItem struct {
	Annotations []ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion `json:"annotations,omitzero,required"`
	Text        string                                                                                `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "output_text".
	Type constant.OutputText `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion struct {
	OfFileCitation          *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation          `json:",omitzero,inline"`
	OfURLCitation           *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation           `json:",omitzero,inline"`
	OfContainerFileCitation *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation `json:",omitzero,inline"`
	OfFilePath              *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFilePath              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) asAny() any {
	if !param.IsOmitted(u.OfFileCitation) {
		return u.OfFileCitation
	} else if !param.IsOmitted(u.OfURLCitation) {
		return u.OfURLCitation
	} else if !param.IsOmitted(u.OfContainerFileCitation) {
		return u.OfContainerFileCitation
	} else if !param.IsOmitted(u.OfFilePath) {
		return u.OfFilePath
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetFileID() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.FileID)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.FileID)
	} else if vt := u.OfFilePath; vt != nil {
		return (*string)(&vt.FileID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetType() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfURLCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFilePath; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation]("file_citation"),
		apijson.Discriminator[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation]("url_citation"),
		apijson.Discriminator[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation]("container_file_citation"),
		apijson.Discriminator[ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFilePath]("file_path"),
	)
}

// File citation annotation for referencing specific files in response content.
//
// The properties FileID, Filename, Index, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation struct {
	// Unique identifier of the referenced file
	FileID string `json:"file_id,required"`
	// Name of the referenced file
	Filename string `json:"filename,required"`
	// Position index of the citation within the content
	Index int64 `json:"index,required"`
	// Annotation type identifier, always "file_citation"
	//
	// This field can be elided, and will marshal its zero value as "file_citation".
	Type constant.FileCitation `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
//
// The properties EndIndex, StartIndex, Title, Type, URL are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation struct {
	// End position of the citation span in the content
	EndIndex int64 `json:"end_index,required"`
	// Start position of the citation span in the content
	StartIndex int64 `json:"start_index,required"`
	// Title of the referenced web resource
	Title string `json:"title,required"`
	// URL of the referenced web resource
	URL string `json:"url,required"`
	// Annotation type identifier, always "url_citation"
	//
	// This field can be elided, and will marshal its zero value as "url_citation".
	Type constant.URLCitation `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex, Type are
// required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// This field can be elided, and will marshal its zero value as
	// "container_file_citation".
	Type constant.ContainerFileCitation `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FileID, Index, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// This field can be elided, and will marshal its zero value as "file_path".
	Type constant.FilePath `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFilePath) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFilePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewParamsInputArrayItemOpenAIResponseMessageRole string

const (
	ResponseNewParamsInputArrayItemOpenAIResponseMessageRoleSystem    ResponseNewParamsInputArrayItemOpenAIResponseMessageRole = "system"
	ResponseNewParamsInputArrayItemOpenAIResponseMessageRoleDeveloper ResponseNewParamsInputArrayItemOpenAIResponseMessageRole = "developer"
	ResponseNewParamsInputArrayItemOpenAIResponseMessageRoleUser      ResponseNewParamsInputArrayItemOpenAIResponseMessageRole = "user"
	ResponseNewParamsInputArrayItemOpenAIResponseMessageRoleAssistant ResponseNewParamsInputArrayItemOpenAIResponseMessageRole = "assistant"
)

// Text response configuration for OpenAI responses.
type ResponseNewParamsText struct {
	// (Optional) Text format configuration specifying output format requirements
	Format ResponseNewParamsTextFormat `json:"format,omitzero"`
	paramObj
}

func (r ResponseNewParamsText) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Text format configuration specifying output format requirements
//
// The property Type is required.
type ResponseNewParamsTextFormat struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseNewParamsTextFormatType `json:"type,omitzero,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the response format. Only used for json_schema.
	Name param.Opt[string] `json:"name,omitzero"`
	// (Optional) Whether to strictly enforce the JSON schema. If true, the response
	// must match the schema exactly. Only used for json_schema.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseNewParamsTextFormatSchemaUnion `json:"schema,omitzero"`
	paramObj
}

func (r ResponseNewParamsTextFormat) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsTextFormat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "text", "json_schema", or "json_object" to identify the format type
type ResponseNewParamsTextFormatType string

const (
	ResponseNewParamsTextFormatTypeText       ResponseNewParamsTextFormatType = "text"
	ResponseNewParamsTextFormatTypeJsonSchema ResponseNewParamsTextFormatType = "json_schema"
	ResponseNewParamsTextFormatTypeJsonObject ResponseNewParamsTextFormatType = "json_object"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsTextFormatSchemaUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsTextFormatSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseNewParamsTextFormatSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsTextFormatSchemaUnion) asAny() any {
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
type ResponseNewParamsToolUnion struct {
	OfOpenAIResponseInputToolWebSearch *ResponseNewParamsToolOpenAIResponseInputToolWebSearch `json:",omitzero,inline"`
	OfFileSearch                       *ResponseNewParamsToolFileSearch                       `json:",omitzero,inline"`
	OfFunction                         *ResponseNewParamsToolFunction                         `json:",omitzero,inline"`
	OfMcp                              *ResponseNewParamsToolMcp                              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseInputToolWebSearch, u.OfFileSearch, u.OfFunction, u.OfMcp)
}
func (u *ResponseNewParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseInputToolWebSearch) {
		return u.OfOpenAIResponseInputToolWebSearch
	} else if !param.IsOmitted(u.OfFileSearch) {
		return u.OfFileSearch
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfMcp) {
		return u.OfMcp
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetSearchContextSize() *string {
	if vt := u.OfOpenAIResponseInputToolWebSearch; vt != nil && vt.SearchContextSize.Valid() {
		return &vt.SearchContextSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetVectorStoreIDs() []string {
	if vt := u.OfFileSearch; vt != nil {
		return vt.VectorStoreIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetFilters() map[string]ResponseNewParamsToolFileSearchFilterUnion {
	if vt := u.OfFileSearch; vt != nil {
		return vt.Filters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetMaxNumResults() *int64 {
	if vt := u.OfFileSearch; vt != nil && vt.MaxNumResults.Valid() {
		return &vt.MaxNumResults.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetRankingOptions() *ResponseNewParamsToolFileSearchRankingOptions {
	if vt := u.OfFileSearch; vt != nil {
		return &vt.RankingOptions
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetName() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetDescription() *string {
	if vt := u.OfFunction; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetParameters() map[string]ResponseNewParamsToolFunctionParameterUnion {
	if vt := u.OfFunction; vt != nil {
		return vt.Parameters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetStrict() *bool {
	if vt := u.OfFunction; vt != nil && vt.Strict.Valid() {
		return &vt.Strict.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetRequireApproval() *ResponseNewParamsToolMcpRequireApprovalUnion {
	if vt := u.OfMcp; vt != nil {
		return &vt.RequireApproval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetServerLabel() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ServerLabel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetServerURL() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ServerURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetAllowedTools() *ResponseNewParamsToolMcpAllowedToolsUnion {
	if vt := u.OfMcp; vt != nil {
		return &vt.AllowedTools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetHeaders() map[string]ResponseNewParamsToolMcpHeaderUnion {
	if vt := u.OfMcp; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsToolUnion) GetType() *string {
	if vt := u.OfOpenAIResponseInputToolWebSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcp; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseNewParamsToolUnion](
		"type",
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search"),
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_preview"),
		apijson.Discriminator[ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_preview_2025_03_11"),
		apijson.Discriminator[ResponseNewParamsToolFileSearch]("file_search"),
		apijson.Discriminator[ResponseNewParamsToolFunction]("function"),
		apijson.Discriminator[ResponseNewParamsToolMcp]("mcp"),
	)
}

// Web search tool configuration for OpenAI response inputs.
//
// The property Type is required.
type ResponseNewParamsToolOpenAIResponseInputToolWebSearch struct {
	// Web search tool type variant to use
	//
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11".
	Type ResponseNewParamsToolOpenAIResponseInputToolWebSearchType `json:"type,omitzero,required"`
	// (Optional) Size of search context, must be "low", "medium", or "high"
	SearchContextSize param.Opt[string] `json:"search_context_size,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolOpenAIResponseInputToolWebSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolOpenAIResponseInputToolWebSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool type variant to use
type ResponseNewParamsToolOpenAIResponseInputToolWebSearchType string

const (
	ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch                  ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search"
	ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview           ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search_preview"
	ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview2025_03_11 ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search_preview_2025_03_11"
)

// File search tool configuration for OpenAI response inputs.
//
// The properties Type, VectorStoreIDs are required.
type ResponseNewParamsToolFileSearch struct {
	// List of vector store identifiers to search within
	VectorStoreIDs []string `json:"vector_store_ids,omitzero,required"`
	// (Optional) Maximum number of search results to return (1-50)
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	// (Optional) Additional filters to apply to the search
	Filters map[string]ResponseNewParamsToolFileSearchFilterUnion `json:"filters,omitzero"`
	// (Optional) Options for ranking and scoring search results
	RankingOptions ResponseNewParamsToolFileSearchRankingOptions `json:"ranking_options,omitzero"`
	// Tool type identifier, always "file_search"
	//
	// This field can be elided, and will marshal its zero value as "file_search".
	Type constant.FileSearch `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsToolFileSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolFileSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolFileSearchFilterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolFileSearchFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseNewParamsToolFileSearchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolFileSearchFilterUnion) asAny() any {
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

// (Optional) Options for ranking and scoring search results
type ResponseNewParamsToolFileSearchRankingOptions struct {
	// (Optional) Name of the ranking algorithm to use
	Ranker param.Opt[string] `json:"ranker,omitzero"`
	// (Optional) Minimum relevance score threshold for results
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolFileSearchRankingOptions) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolFileSearchRankingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool configuration for OpenAI response inputs.
//
// The properties Name, Type are required.
type ResponseNewParamsToolFunction struct {
	// Name of the function that can be called
	Name string `json:"name,required"`
	// (Optional) Description of what the function does
	Description param.Opt[string] `json:"description,omitzero"`
	// (Optional) Whether to enforce strict parameter validation
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// (Optional) JSON schema defining the function's parameters
	Parameters map[string]ResponseNewParamsToolFunctionParameterUnion `json:"parameters,omitzero"`
	// Tool type identifier, always "function"
	//
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsToolFunction) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolFunctionParameterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolFunctionParameterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseNewParamsToolFunctionParameterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolFunctionParameterUnion) asAny() any {
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

// Model Context Protocol (MCP) tool configuration for OpenAI response inputs.
//
// The properties RequireApproval, ServerLabel, ServerURL, Type are required.
type ResponseNewParamsToolMcp struct {
	// Approval requirement for tool calls ("always", "never", or filter)
	RequireApproval ResponseNewParamsToolMcpRequireApprovalUnion `json:"require_approval,omitzero,required"`
	// Label to identify this MCP server
	ServerLabel string `json:"server_label,required"`
	// URL endpoint of the MCP server
	ServerURL string `json:"server_url,required"`
	// (Optional) Restriction on which tools can be used from this server
	AllowedTools ResponseNewParamsToolMcpAllowedToolsUnion `json:"allowed_tools,omitzero"`
	// (Optional) HTTP headers to include when connecting to the server
	Headers map[string]ResponseNewParamsToolMcpHeaderUnion `json:"headers,omitzero"`
	// Tool type identifier, always "mcp"
	//
	// This field can be elided, and will marshal its zero value as "mcp".
	Type constant.Mcp `json:"type,required"`
	paramObj
}

func (r ResponseNewParamsToolMcp) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolMcp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolMcpRequireApprovalUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfResponseNewsToolMcpRequireApprovalString)
	OfResponseNewsToolMcpRequireApprovalString param.Opt[ResponseNewParamsToolMcpRequireApprovalString] `json:",omitzero,inline"`
	OfApprovalFilter                           *ResponseNewParamsToolMcpRequireApprovalApprovalFilter   `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolMcpRequireApprovalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfResponseNewsToolMcpRequireApprovalString, u.OfApprovalFilter)
}
func (u *ResponseNewParamsToolMcpRequireApprovalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolMcpRequireApprovalUnion) asAny() any {
	if !param.IsOmitted(u.OfResponseNewsToolMcpRequireApprovalString) {
		return &u.OfResponseNewsToolMcpRequireApprovalString
	} else if !param.IsOmitted(u.OfApprovalFilter) {
		return u.OfApprovalFilter
	}
	return nil
}

type ResponseNewParamsToolMcpRequireApprovalString string

const (
	ResponseNewParamsToolMcpRequireApprovalStringAlways ResponseNewParamsToolMcpRequireApprovalString = "always"
	ResponseNewParamsToolMcpRequireApprovalStringNever  ResponseNewParamsToolMcpRequireApprovalString = "never"
)

// Filter configuration for MCP tool approval requirements.
type ResponseNewParamsToolMcpRequireApprovalApprovalFilter struct {
	// (Optional) List of tool names that always require approval
	Always []string `json:"always,omitzero"`
	// (Optional) List of tool names that never require approval
	Never []string `json:"never,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolMcpRequireApprovalApprovalFilter) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolMcpRequireApprovalApprovalFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolMcpRequireApprovalApprovalFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolMcpAllowedToolsUnion struct {
	OfStringArray        []string                                                `json:",omitzero,inline"`
	OfAllowedToolsFilter *ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolMcpAllowedToolsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfAllowedToolsFilter)
}
func (u *ResponseNewParamsToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolMcpAllowedToolsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfAllowedToolsFilter) {
		return u.OfAllowedToolsFilter
	}
	return nil
}

// Filter configuration for restricting which MCP tools can be used.
type ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter struct {
	// (Optional) List of specific tool names that are allowed
	ToolNames []string `json:"tool_names,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsToolMcpHeaderUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsToolMcpHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseNewParamsToolMcpHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsToolMcpHeaderUnion) asAny() any {
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

type ResponseListParams struct {
	// The ID of the last response to return.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The number of responses to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The model to filter responses by.
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	// The order to sort responses by when sorted by created_at ('asc' or 'desc').
	//
	// Any of "asc", "desc".
	Order ResponseListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResponseListParams]'s query parameters as `url.Values`.
func (r ResponseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order to sort responses by when sorted by created_at ('asc' or 'desc').
type ResponseListParamsOrder string

const (
	ResponseListParamsOrderAsc  ResponseListParamsOrder = "asc"
	ResponseListParamsOrderDesc ResponseListParamsOrder = "desc"
)
