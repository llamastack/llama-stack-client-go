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

// List all OpenAI responses.
func (r *ResponseService) NewStreaming(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ResponseObjectStreamUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ResponseObjectStreamUnion](ssestream.NewDecoder(raw), err)
}

// Retrieve an OpenAI response by its ID.
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

// List all OpenAI responses.
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

// List all OpenAI responses.
func (r *ResponseService) ListAutoPaging(ctx context.Context, query ResponseListParams, opts ...option.RequestOption) *pagination.OpenAICursorPageAutoPager[ResponseListResponse] {
	return pagination.NewOpenAICursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an OpenAI response by its ID.
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
	// (Optional) Nucleus sampling parameter used for generation
	TopP float64 `json:"top_p"`
	// (Optional) Truncation strategy applied to the response
	Truncation string `json:"truncation"`
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
		TopP               respjson.Field
		Truncation         respjson.Field
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

// ResponseObjectStreamUnion contains all possible properties and values from
// [ResponseObjectStreamResponseCreated],
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
// [ResponseObjectStreamResponseCompleted].
//
// Use the [ResponseObjectStreamUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamUnion struct {
	// This field is from variant [ResponseObjectStreamResponseCreated].
	Response ResponseObject `json:"response"`
	// Any of "response.created", "response.output_item.added",
	// "response.output_item.done", "response.output_text.delta",
	// "response.output_text.done", "response.function_call_arguments.delta",
	// "response.function_call_arguments.done", "response.web_search_call.in_progress",
	// "response.web_search_call.searching", "response.web_search_call.completed",
	// "response.mcp_list_tools.in_progress", "response.mcp_list_tools.failed",
	// "response.mcp_list_tools.completed", "response.mcp_call.arguments.delta",
	// "response.mcp_call.arguments.done", "response.mcp_call.in_progress",
	// "response.mcp_call.failed", "response.mcp_call.completed",
	// "response.content_part.added", "response.content_part.done",
	// "response.completed".
	Type string `json:"type"`
	// This field is a union of [ResponseObjectStreamResponseOutputItemAddedItemUnion],
	// [ResponseObjectStreamResponseOutputItemDoneItemUnion]
	Item           ResponseObjectStreamUnionItem `json:"item"`
	OutputIndex    int64                         `json:"output_index"`
	ResponseID     string                        `json:"response_id"`
	SequenceNumber int64                         `json:"sequence_number"`
	ContentIndex   int64                         `json:"content_index"`
	Delta          string                        `json:"delta"`
	ItemID         string                        `json:"item_id"`
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
		Item           respjson.Field
		OutputIndex    respjson.Field
		ResponseID     respjson.Field
		SequenceNumber respjson.Field
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
func (ResponseObjectStreamResponseCompleted) implResponseObjectStreamUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseCreated:
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
//	case llamastackclient.ResponseObjectStreamResponseCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamUnion) AsAny() anyResponseObjectStream {
	switch u.Type {
	case "response.created":
		return u.AsResponseCreated()
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
	case "response.completed":
		return u.AsResponseCompleted()
	}
	return nil
}

func (u ResponseObjectStreamUnion) AsResponseCreated() (v ResponseObjectStreamResponseCreated) {
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
	Text    string `json:"text"`
	Type    string `json:"type"`
	Refusal string `json:"refusal"`
	JSON    struct {
		Text    respjson.Field
		Type    respjson.Field
		Refusal respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ResponseObjectStreamUnionPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming event indicating a new response has been created.
type ResponseObjectStreamResponseCreated struct {
	// The newly created response object
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
	// Unique identifier of the output item containing this content part
	ItemID string `json:"item_id,required"`
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
		ItemID         respjson.Field
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
// [ResponseObjectStreamResponseContentPartAddedPartRefusal].
//
// Use the [ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartAddedPartUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartOutputText].
	Text string `json:"text"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartAddedPartRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text    respjson.Field
		Type    respjson.Field
		Refusal respjson.Field
		raw     string
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartAddedPartUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartOutputText:
//	case llamastackclient.ResponseObjectStreamResponseContentPartAddedPartRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartAddedPartUnion) AsAny() anyResponseObjectStreamResponseContentPartAddedPart {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
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

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartAddedPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseContentPartAddedPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartAddedPartOutputText struct {
	Text string              `json:"text,required"`
	Type constant.OutputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
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

type ResponseObjectStreamResponseContentPartAddedPartRefusal struct {
	Refusal string           `json:"refusal,required"`
	Type    constant.Refusal `json:"type,required"`
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

// Streaming event for when a content part is completed.
type ResponseObjectStreamResponseContentPartDone struct {
	// Unique identifier of the output item containing this content part
	ItemID string `json:"item_id,required"`
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
		ItemID         respjson.Field
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
// [ResponseObjectStreamResponseContentPartDonePartRefusal].
//
// Use the [ResponseObjectStreamResponseContentPartDonePartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseObjectStreamResponseContentPartDonePartUnion struct {
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartOutputText].
	Text string `json:"text"`
	// Any of "output_text", "refusal".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseObjectStreamResponseContentPartDonePartRefusal].
	Refusal string `json:"refusal"`
	JSON    struct {
		Text    respjson.Field
		Type    respjson.Field
		Refusal respjson.Field
		raw     string
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseContentPartDonePartUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartOutputText:
//	case llamastackclient.ResponseObjectStreamResponseContentPartDonePartRefusal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseObjectStreamResponseContentPartDonePartUnion) AsAny() anyResponseObjectStreamResponseContentPartDonePart {
	switch u.Type {
	case "output_text":
		return u.AsOutputText()
	case "refusal":
		return u.AsRefusal()
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

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseContentPartDonePartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectStreamResponseContentPartDonePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseContentPartDonePartOutputText struct {
	Text string              `json:"text,required"`
	Type constant.OutputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
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

type ResponseObjectStreamResponseContentPartDonePartRefusal struct {
	Refusal string           `json:"refusal,required"`
	Type    constant.Refusal `json:"type,required"`
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

// Streaming event indicating a response has been completed.
type ResponseObjectStreamResponseCompleted struct {
	// The completed response object
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

// Paginated list of OpenAI response objects with navigation metadata.
type ResponseNewResponse struct {
	// List of response objects with their input context
	Data []ResponseNewResponseData `json:"data,required"`
	// Identifier of the first item in this page
	FirstID string `json:"first_id,required"`
	// Whether there are more results available beyond this page
	HasMore bool `json:"has_more,required"`
	// Identifier of the last item in this page
	LastID string `json:"last_id,required"`
	// Object type identifier, always "list"
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
func (r ResponseNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAI response object extended with input context information.
type ResponseNewResponseData struct {
	// Unique identifier for this response
	ID string `json:"id,required"`
	// Unix timestamp when the response was created
	CreatedAt int64 `json:"created_at,required"`
	// List of input items that led to this response
	Input []ResponseNewResponseDataInputUnion `json:"input,required"`
	// Model identifier used for generation
	Model string `json:"model,required"`
	// Object type identifier, always "response"
	Object constant.Response `json:"object,required"`
	// List of generated output items (messages, tool calls, etc.)
	Output []ResponseNewResponseDataOutputUnion `json:"output,required"`
	// Whether tool calls can be executed in parallel
	ParallelToolCalls bool `json:"parallel_tool_calls,required"`
	// Current status of the response generation
	Status string `json:"status,required"`
	// Text formatting configuration for the response
	Text ResponseNewResponseDataText `json:"text,required"`
	// (Optional) Error details if the response generation failed
	Error ResponseNewResponseDataError `json:"error"`
	// (Optional) ID of the previous response in a conversation
	PreviousResponseID string `json:"previous_response_id"`
	// (Optional) Sampling temperature used for generation
	Temperature float64 `json:"temperature"`
	// (Optional) Nucleus sampling parameter used for generation
	TopP float64 `json:"top_p"`
	// (Optional) Truncation strategy applied to the response
	Truncation string `json:"truncation"`
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
		TopP               respjson.Field
		Truncation         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataInputUnion contains all possible properties and values
// from [ResponseNewResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseNewResponseDataInputOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseNewResponseDataInputOpenAIResponseInputFunctionToolCallOutput],
// [ResponseNewResponseDataInputOpenAIResponseMcpApprovalRequest],
// [ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse],
// [ResponseNewResponseDataInputOpenAIResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseNewResponseDataInputUnion struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall].
	Results   []ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
	Arguments string                                                                            `json:"arguments"`
	CallID    string                                                                            `json:"call_id"`
	Name      string                                                                            `json:"name"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseInputFunctionToolCallOutput].
	Output string `json:"output"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMcpApprovalRequest].
	ServerLabel string `json:"server_label"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse].
	ApprovalRequestID string `json:"approval_request_id"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse].
	Approve bool `json:"approve"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse].
	Reason string `json:"reason"`
	// This field is from variant [ResponseNewResponseDataInputOpenAIResponseMessage].
	Content ResponseNewResponseDataInputOpenAIResponseMessageContentUnion `json:"content"`
	// This field is from variant [ResponseNewResponseDataInputOpenAIResponseMessage].
	Role ResponseNewResponseDataInputOpenAIResponseMessageRole `json:"role"`
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

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v ResponseNewResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v ResponseNewResponseDataInputOpenAIResponseOutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseNewResponseDataInputOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseMcpApprovalRequest() (v ResponseNewResponseDataInputOpenAIResponseMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseMcpApprovalResponse() (v ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputUnion) AsOpenAIResponseMessage() (v ResponseNewResponseDataInputOpenAIResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseNewResponseDataInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ResponseNewResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResult `json:"results"`
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
func (r ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion struct {
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

func (u ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseNewResponseDataInputOpenAIResponseOutputMessageFunctionToolCall struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseOutputMessageFunctionToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseNewResponseDataInputOpenAIResponseInputFunctionToolCallOutput struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseInputFunctionToolCallOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseNewResponseDataInputOpenAIResponseMcpApprovalRequest struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMcpApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an MCP approval request.
type ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMcpApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseNewResponseDataInputOpenAIResponseMessage struct {
	Content ResponseNewResponseDataInputOpenAIResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseNewResponseDataInputOpenAIResponseMessageRole `json:"role,required"`
	Type   constant.Message                                      `json:"type,required"`
	ID     string                                                `json:"id"`
	Status string                                                `json:"status"`
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
func (r ResponseNewResponseDataInputOpenAIResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataInputOpenAIResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataInputOpenAIResponseMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion],
// [[]ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseNewResponseDataInputOpenAIResponseMessageContentArray OfVariant2]
type ResponseNewResponseDataInputOpenAIResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseNewResponseDataInputOpenAIResponseMessageContentArray []ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem] instead of
	// an object.
	OfVariant2 []ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                                        respjson.Field
		OfResponseNewResponseDataInputOpenAIResponseMessageContentArray respjson.Field
		OfVariant2                                                      respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentUnion) AsResponseNewResponseDataInputOpenAIResponseMessageContentArray() (v []ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentUnion) AsVariant2() (v []ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataInputOpenAIResponseMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion contains
// all possible properties and values from
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText],
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage].
//
// Use the
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage].
	Detail ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem is
// implemented by each variant of
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion.AsAny]
type anyResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem interface {
	implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion()
}

func (ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText) implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion() {
}
func (ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText:
//	case llamastackclient.ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion) AsAny() anyResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion) AsInputText() (v ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion) AsInputImage() (v ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail string

const (
	ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetailLow  ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail = "low"
	ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetailHigh ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail = "high"
	ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetailAuto ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetail string

const (
	ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetailLow  ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetail = "low"
	ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetailHigh ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetail = "high"
	ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetailAuto ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemDetail = "auto"
)

type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem struct {
	Annotations []ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                             `json:"text,required"`
	Type        constant.OutputText                                                                `json:"type,required"`
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion
// contains all possible properties and values from
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation],
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation],
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotation
// is implemented by each variant of
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion]
// to add type safety for the return type of
// [ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotation interface {
	implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion()
}

func (ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) implResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsAny() anyResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotation {
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

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataInputOpenAIResponseMessageRole string

const (
	ResponseNewResponseDataInputOpenAIResponseMessageRoleSystem    ResponseNewResponseDataInputOpenAIResponseMessageRole = "system"
	ResponseNewResponseDataInputOpenAIResponseMessageRoleDeveloper ResponseNewResponseDataInputOpenAIResponseMessageRole = "developer"
	ResponseNewResponseDataInputOpenAIResponseMessageRoleUser      ResponseNewResponseDataInputOpenAIResponseMessageRole = "user"
	ResponseNewResponseDataInputOpenAIResponseMessageRoleAssistant ResponseNewResponseDataInputOpenAIResponseMessageRole = "assistant"
)

type ResponseNewResponseDataInputRole string

const (
	ResponseNewResponseDataInputRoleSystem    ResponseNewResponseDataInputRole = "system"
	ResponseNewResponseDataInputRoleDeveloper ResponseNewResponseDataInputRole = "developer"
	ResponseNewResponseDataInputRoleUser      ResponseNewResponseDataInputRole = "user"
	ResponseNewResponseDataInputRoleAssistant ResponseNewResponseDataInputRole = "assistant"
)

// ResponseNewResponseDataOutputUnion contains all possible properties and values
// from [ResponseNewResponseDataOutputMessage],
// [ResponseNewResponseDataOutputWebSearchCall],
// [ResponseNewResponseDataOutputFileSearchCall],
// [ResponseNewResponseDataOutputFunctionCall],
// [ResponseNewResponseDataOutputMcpCall],
// [ResponseNewResponseDataOutputMcpListTools],
// [ResponseNewResponseDataOutputMcpApprovalRequest].
//
// Use the [ResponseNewResponseDataOutputUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseNewResponseDataOutputUnion struct {
	// This field is from variant [ResponseNewResponseDataOutputMessage].
	Content ResponseNewResponseDataOutputMessageContentUnion `json:"content"`
	// This field is from variant [ResponseNewResponseDataOutputMessage].
	Role ResponseNewResponseDataOutputMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools", "mcp_approval_request".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [ResponseNewResponseDataOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseNewResponseDataOutputFileSearchCall].
	Results   []ResponseNewResponseDataOutputFileSearchCallResult `json:"results"`
	Arguments string                                              `json:"arguments"`
	// This field is from variant [ResponseNewResponseDataOutputFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseNewResponseDataOutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [ResponseNewResponseDataOutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [ResponseNewResponseDataOutputMcpListTools].
	Tools []ResponseNewResponseDataOutputMcpListToolsTool `json:"tools"`
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

// anyResponseNewResponseDataOutput is implemented by each variant of
// [ResponseNewResponseDataOutputUnion] to add type safety for the return type of
// [ResponseNewResponseDataOutputUnion.AsAny]
type anyResponseNewResponseDataOutput interface {
	implResponseNewResponseDataOutputUnion()
}

func (ResponseNewResponseDataOutputMessage) implResponseNewResponseDataOutputUnion()            {}
func (ResponseNewResponseDataOutputWebSearchCall) implResponseNewResponseDataOutputUnion()      {}
func (ResponseNewResponseDataOutputFileSearchCall) implResponseNewResponseDataOutputUnion()     {}
func (ResponseNewResponseDataOutputFunctionCall) implResponseNewResponseDataOutputUnion()       {}
func (ResponseNewResponseDataOutputMcpCall) implResponseNewResponseDataOutputUnion()            {}
func (ResponseNewResponseDataOutputMcpListTools) implResponseNewResponseDataOutputUnion()       {}
func (ResponseNewResponseDataOutputMcpApprovalRequest) implResponseNewResponseDataOutputUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseNewResponseDataOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseNewResponseDataOutputMessage:
//	case llamastackclient.ResponseNewResponseDataOutputWebSearchCall:
//	case llamastackclient.ResponseNewResponseDataOutputFileSearchCall:
//	case llamastackclient.ResponseNewResponseDataOutputFunctionCall:
//	case llamastackclient.ResponseNewResponseDataOutputMcpCall:
//	case llamastackclient.ResponseNewResponseDataOutputMcpListTools:
//	case llamastackclient.ResponseNewResponseDataOutputMcpApprovalRequest:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseNewResponseDataOutputUnion) AsAny() anyResponseNewResponseDataOutput {
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

func (u ResponseNewResponseDataOutputUnion) AsMessage() (v ResponseNewResponseDataOutputMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputUnion) AsWebSearchCall() (v ResponseNewResponseDataOutputWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputUnion) AsFileSearchCall() (v ResponseNewResponseDataOutputFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputUnion) AsFunctionCall() (v ResponseNewResponseDataOutputFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputUnion) AsMcpCall() (v ResponseNewResponseDataOutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputUnion) AsMcpListTools() (v ResponseNewResponseDataOutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputUnion) AsMcpApprovalRequest() (v ResponseNewResponseDataOutputMcpApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseNewResponseDataOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseNewResponseDataOutputMessage struct {
	Content ResponseNewResponseDataOutputMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseNewResponseDataOutputMessageRole `json:"role,required"`
	Type   constant.Message                         `json:"type,required"`
	ID     string                                   `json:"id"`
	Status string                                   `json:"status"`
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
func (r ResponseNewResponseDataOutputMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataOutputMessageContentUnion contains all possible
// properties and values from [string],
// [[]ResponseNewResponseDataOutputMessageContentArrayItemUnion],
// [[]ResponseNewResponseDataOutputMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfResponseNewResponseDataOutputMessageContentArray
// OfVariant2]
type ResponseNewResponseDataOutputMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseNewResponseDataOutputMessageContentArrayItemUnion] instead of an
	// object.
	OfResponseNewResponseDataOutputMessageContentArray []ResponseNewResponseDataOutputMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseNewResponseDataOutputMessageContentArrayItem] instead of an object.
	OfVariant2 []ResponseNewResponseDataOutputMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                           respjson.Field
		OfResponseNewResponseDataOutputMessageContentArray respjson.Field
		OfVariant2                                         respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u ResponseNewResponseDataOutputMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMessageContentUnion) AsResponseNewResponseDataOutputMessageContentArray() (v []ResponseNewResponseDataOutputMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMessageContentUnion) AsVariant2() (v []ResponseNewResponseDataOutputMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataOutputMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseNewResponseDataOutputMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataOutputMessageContentArrayItemUnion contains all possible
// properties and values from
// [ResponseNewResponseDataOutputMessageContentArrayItemInputText],
// [ResponseNewResponseDataOutputMessageContentArrayItemInputImage].
//
// Use the [ResponseNewResponseDataOutputMessageContentArrayItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseNewResponseDataOutputMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseNewResponseDataOutputMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseNewResponseDataOutputMessageContentArrayItemInputImage].
	Detail ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseNewResponseDataOutputMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseNewResponseDataOutputMessageContentArrayItem is implemented by each
// variant of [ResponseNewResponseDataOutputMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ResponseNewResponseDataOutputMessageContentArrayItemUnion.AsAny]
type anyResponseNewResponseDataOutputMessageContentArrayItem interface {
	implResponseNewResponseDataOutputMessageContentArrayItemUnion()
}

func (ResponseNewResponseDataOutputMessageContentArrayItemInputText) implResponseNewResponseDataOutputMessageContentArrayItemUnion() {
}
func (ResponseNewResponseDataOutputMessageContentArrayItemInputImage) implResponseNewResponseDataOutputMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseNewResponseDataOutputMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseNewResponseDataOutputMessageContentArrayItemInputText:
//	case llamastackclient.ResponseNewResponseDataOutputMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseNewResponseDataOutputMessageContentArrayItemUnion) AsAny() anyResponseNewResponseDataOutputMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseNewResponseDataOutputMessageContentArrayItemUnion) AsInputText() (v ResponseNewResponseDataOutputMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMessageContentArrayItemUnion) AsInputImage() (v ResponseNewResponseDataOutputMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataOutputMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataOutputMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ResponseNewResponseDataOutputMessageContentArrayItemInputText struct {
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
func (r ResponseNewResponseDataOutputMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataOutputMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ResponseNewResponseDataOutputMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ResponseNewResponseDataOutputMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataOutputMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetail string

const (
	ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetailLow  ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetail = "low"
	ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetailHigh ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetail = "high"
	ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetailAuto ResponseNewResponseDataOutputMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ResponseNewResponseDataOutputMessageContentArrayItemDetail string

const (
	ResponseNewResponseDataOutputMessageContentArrayItemDetailLow  ResponseNewResponseDataOutputMessageContentArrayItemDetail = "low"
	ResponseNewResponseDataOutputMessageContentArrayItemDetailHigh ResponseNewResponseDataOutputMessageContentArrayItemDetail = "high"
	ResponseNewResponseDataOutputMessageContentArrayItemDetailAuto ResponseNewResponseDataOutputMessageContentArrayItemDetail = "auto"
)

type ResponseNewResponseDataOutputMessageContentArrayItem struct {
	Annotations []ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                `json:"text,required"`
	Type        constant.OutputText                                                   `json:"type,required"`
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
func (r ResponseNewResponseDataOutputMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion contains all
// possible properties and values from
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation],
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation],
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseNewResponseDataOutputMessageContentArrayItemAnnotation is implemented
// by each variant of
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion] to add
// type safety for the return type of
// [ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseNewResponseDataOutputMessageContentArrayItemAnnotation interface {
	implResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion()
}

func (ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation) implResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation) implResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) implResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath) implResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) AsAny() anyResponseNewResponseDataOutputMessageContentArrayItemAnnotation {
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

func (u ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataOutputMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataOutputMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseNewResponseDataOutputMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataOutputMessageRole string

const (
	ResponseNewResponseDataOutputMessageRoleSystem    ResponseNewResponseDataOutputMessageRole = "system"
	ResponseNewResponseDataOutputMessageRoleDeveloper ResponseNewResponseDataOutputMessageRole = "developer"
	ResponseNewResponseDataOutputMessageRoleUser      ResponseNewResponseDataOutputMessageRole = "user"
	ResponseNewResponseDataOutputMessageRoleAssistant ResponseNewResponseDataOutputMessageRole = "assistant"
)

// Web search tool call output message for OpenAI responses.
type ResponseNewResponseDataOutputWebSearchCall struct {
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
func (r ResponseNewResponseDataOutputWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ResponseNewResponseDataOutputFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ResponseNewResponseDataOutputFileSearchCallResult `json:"results"`
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
func (r ResponseNewResponseDataOutputFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ResponseNewResponseDataOutputFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ResponseNewResponseDataOutputFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion struct {
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

func (u ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataOutputFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function tool call output message for OpenAI responses.
type ResponseNewResponseDataOutputFunctionCall struct {
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
func (r ResponseNewResponseDataOutputFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ResponseNewResponseDataOutputMcpCall struct {
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
func (r ResponseNewResponseDataOutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ResponseNewResponseDataOutputMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ResponseNewResponseDataOutputMcpListToolsTool `json:"tools,required"`
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
func (r ResponseNewResponseDataOutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ResponseNewResponseDataOutputMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ResponseNewResponseDataOutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseNewResponseDataOutputMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request for human approval of a tool invocation.
type ResponseNewResponseDataOutputMcpApprovalRequest struct {
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
func (r ResponseNewResponseDataOutputMcpApprovalRequest) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataOutputMcpApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewResponseDataOutputRole string

const (
	ResponseNewResponseDataOutputRoleSystem    ResponseNewResponseDataOutputRole = "system"
	ResponseNewResponseDataOutputRoleDeveloper ResponseNewResponseDataOutputRole = "developer"
	ResponseNewResponseDataOutputRoleUser      ResponseNewResponseDataOutputRole = "user"
	ResponseNewResponseDataOutputRoleAssistant ResponseNewResponseDataOutputRole = "assistant"
)

// Text formatting configuration for the response
type ResponseNewResponseDataText struct {
	// (Optional) Text format configuration specifying output format requirements
	Format ResponseNewResponseDataTextFormat `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseNewResponseDataText) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Text format configuration specifying output format requirements
type ResponseNewResponseDataTextFormat struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseNewResponseDataTextFormatType `json:"type,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description string `json:"description"`
	// The name of the response format. Only used for json_schema.
	Name string `json:"name"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseNewResponseDataTextFormatSchemaUnion `json:"schema"`
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
func (r ResponseNewResponseDataTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "text", "json_schema", or "json_object" to identify the format type
type ResponseNewResponseDataTextFormatType string

const (
	ResponseNewResponseDataTextFormatTypeText       ResponseNewResponseDataTextFormatType = "text"
	ResponseNewResponseDataTextFormatTypeJsonSchema ResponseNewResponseDataTextFormatType = "json_schema"
	ResponseNewResponseDataTextFormatTypeJsonObject ResponseNewResponseDataTextFormatType = "json_object"
)

// ResponseNewResponseDataTextFormatSchemaUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseNewResponseDataTextFormatSchemaUnion struct {
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

func (u ResponseNewResponseDataTextFormatSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataTextFormatSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataTextFormatSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseNewResponseDataTextFormatSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseNewResponseDataTextFormatSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseNewResponseDataTextFormatSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Error details if the response generation failed
type ResponseNewResponseDataError struct {
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
func (r ResponseNewResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *ResponseNewResponseDataError) UnmarshalJSON(data []byte) error {
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
	// (Optional) Nucleus sampling parameter used for generation
	TopP float64 `json:"top_p"`
	// (Optional) Truncation strategy applied to the response
	Truncation string `json:"truncation"`
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
		TopP               respjson.Field
		Truncation         respjson.Field
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
	// The ID of the last response to return.
	After param.Opt[string] `json:"after,omitzero"`
	// The number of responses to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The model to filter responses by.
	Model param.Opt[string] `json:"model,omitzero"`
	// The order to sort responses by when sorted by created_at ('asc' or 'desc').
	//
	// Any of "asc", "desc".
	Order ResponseNewParamsOrder `json:"order,omitzero"`
	paramObj
}

func (r ResponseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The order to sort responses by when sorted by created_at ('asc' or 'desc').
type ResponseNewParamsOrder string

const (
	ResponseNewParamsOrderAsc  ResponseNewParamsOrder = "asc"
	ResponseNewParamsOrderDesc ResponseNewParamsOrder = "desc"
)

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
