// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/llama-stack-client-go/internal/apijson"
	"github.com/stainless-sdks/llama-stack-client-go/internal/apiquery"
	"github.com/stainless-sdks/llama-stack-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llama-stack-client-go/option"
	"github.com/stainless-sdks/llama-stack-client-go/packages/param"
	"github.com/stainless-sdks/llama-stack-client-go/packages/respjson"
	"github.com/stainless-sdks/llama-stack-client-go/packages/ssestream"
	"github.com/stainless-sdks/llama-stack-client-go/shared/constant"
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

// Create a new OpenAI response.
func (r *ResponseService) New(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new OpenAI response.
func (r *ResponseService) NewStreaming(ctx context.Context, body ResponseNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ResponseObjectStreamUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/openai/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ResponseObjectStreamUnion](ssestream.NewDecoder(raw), err)
}

// Retrieve an OpenAI response by its ID.
func (r *ResponseService) Get(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = append(r.Options[:], opts...)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all OpenAI responses.
func (r *ResponseService) List(ctx context.Context, query ResponseListParams, opts ...option.RequestOption) (res *ResponseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ResponseObject struct {
	ID                 string                      `json:"id,required"`
	CreatedAt          int64                       `json:"created_at,required"`
	Model              string                      `json:"model,required"`
	Object             constant.Response           `json:"object,required"`
	Output             []ResponseObjectOutputUnion `json:"output,required"`
	ParallelToolCalls  bool                        `json:"parallel_tool_calls,required"`
	Status             string                      `json:"status,required"`
	Text               ResponseObjectText          `json:"text,required"`
	Error              ResponseObjectError         `json:"error"`
	PreviousResponseID string                      `json:"previous_response_id"`
	Temperature        float64                     `json:"temperature"`
	TopP               float64                     `json:"top_p"`
	Truncation         string                      `json:"truncation"`
	User               string                      `json:"user"`
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
		User               respjson.Field
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
// [ResponseObjectOutputMcpCall], [ResponseObjectOutputMcpListTools].
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
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [ResponseObjectOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseObjectOutputFileSearchCall].
	Results   []map[string]ResponseObjectOutputFileSearchCallResultUnion `json:"results"`
	Arguments string                                                     `json:"arguments"`
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

func (ResponseObjectOutputMessage) implResponseObjectOutputUnion()        {}
func (ResponseObjectOutputWebSearchCall) implResponseObjectOutputUnion()  {}
func (ResponseObjectOutputFileSearchCall) implResponseObjectOutputUnion() {}
func (ResponseObjectOutputFunctionCall) implResponseObjectOutputUnion()   {}
func (ResponseObjectOutputMcpCall) implResponseObjectOutputUnion()        {}
func (ResponseObjectOutputMcpListTools) implResponseObjectOutputUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectOutputMessage:
//	case llamastackclient.ResponseObjectOutputWebSearchCall:
//	case llamastackclient.ResponseObjectOutputFileSearchCall:
//	case llamastackclient.ResponseObjectOutputFunctionCall:
//	case llamastackclient.ResponseObjectOutputMcpCall:
//	case llamastackclient.ResponseObjectOutputMcpListTools:
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

type ResponseObjectOutputMessageContentArrayItemInputText struct {
	Text string             `json:"text,required"`
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

type ResponseObjectOutputMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseObjectOutputMessageContentArrayItemInputImageDetail `json:"detail,required"`
	Type     constant.InputImage                                         `json:"type,required"`
	ImageURL string                                                      `json:"image_url"`
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

type ResponseObjectOutputMessageContentArrayItemInputImageDetail string

const (
	ResponseObjectOutputMessageContentArrayItemInputImageDetailLow  ResponseObjectOutputMessageContentArrayItemInputImageDetail = "low"
	ResponseObjectOutputMessageContentArrayItemInputImageDetailHigh ResponseObjectOutputMessageContentArrayItemInputImageDetail = "high"
	ResponseObjectOutputMessageContentArrayItemInputImageDetailAuto ResponseObjectOutputMessageContentArrayItemInputImageDetail = "auto"
)

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

type ResponseObjectOutputMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string                `json:"file_id,required"`
	Filename string                `json:"filename,required"`
	Index    int64                 `json:"index,required"`
	Type     constant.FileCitation `json:"type,required"`
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

type ResponseObjectOutputMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64                `json:"end_index,required"`
	StartIndex int64                `json:"start_index,required"`
	Title      string               `json:"title,required"`
	Type       constant.URLCitation `json:"type,required"`
	URL        string               `json:"url,required"`
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

type ResponseObjectOutputWebSearchCall struct {
	ID     string                 `json:"id,required"`
	Status string                 `json:"status,required"`
	Type   constant.WebSearchCall `json:"type,required"`
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

type ResponseObjectOutputFileSearchCall struct {
	ID      string                                                     `json:"id,required"`
	Queries []string                                                   `json:"queries,required"`
	Status  string                                                     `json:"status,required"`
	Type    constant.FileSearchCall                                    `json:"type,required"`
	Results []map[string]ResponseObjectOutputFileSearchCallResultUnion `json:"results"`
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

// ResponseObjectOutputFileSearchCallResultUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectOutputFileSearchCallResultUnion struct {
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

func (u ResponseObjectOutputFileSearchCallResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputFileSearchCallResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputFileSearchCallResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectOutputFileSearchCallResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectOutputFileSearchCallResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseObjectOutputFileSearchCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectOutputFunctionCall struct {
	Arguments string                `json:"arguments,required"`
	CallID    string                `json:"call_id,required"`
	Name      string                `json:"name,required"`
	Type      constant.FunctionCall `json:"type,required"`
	ID        string                `json:"id"`
	Status    string                `json:"status"`
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

type ResponseObjectOutputMcpCall struct {
	ID          string           `json:"id,required"`
	Arguments   string           `json:"arguments,required"`
	Name        string           `json:"name,required"`
	ServerLabel string           `json:"server_label,required"`
	Type        constant.McpCall `json:"type,required"`
	Error       string           `json:"error"`
	Output      string           `json:"output"`
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

type ResponseObjectOutputMcpListTools struct {
	ID          string                                 `json:"id,required"`
	ServerLabel string                                 `json:"server_label,required"`
	Tools       []ResponseObjectOutputMcpListToolsTool `json:"tools,required"`
	Type        constant.McpListTools                  `json:"type,required"`
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

type ResponseObjectOutputMcpListToolsTool struct {
	InputSchema map[string]ResponseObjectOutputMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	Name        string                                                          `json:"name,required"`
	Description string                                                          `json:"description"`
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

type ResponseObjectOutputRole string

const (
	ResponseObjectOutputRoleSystem    ResponseObjectOutputRole = "system"
	ResponseObjectOutputRoleDeveloper ResponseObjectOutputRole = "developer"
	ResponseObjectOutputRoleUser      ResponseObjectOutputRole = "user"
	ResponseObjectOutputRoleAssistant ResponseObjectOutputRole = "assistant"
)

type ResponseObjectText struct {
	// Configuration for Responses API text format.
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

// Configuration for Responses API text format.
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

type ResponseObjectError struct {
	Code    string `json:"code,required"`
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
	// "response.mcp_call.failed", "response.mcp_call.completed", "response.completed".
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
	JSON      struct {
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
	// [[]map[string]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion],
	// [[]map[string]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion]
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
// OfMapOfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultMap
// OfMapOfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultMap]
type ResponseObjectStreamUnionItemResults struct {
	// This field will be present if the value is a
	// [[]map[string]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion]
	// instead of an object.
	OfMapOfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultMap []map[string]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]map[string]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion]
	// instead of an object.
	OfMapOfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultMap []map[string]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion `json:",inline"`
	JSON                                                                         struct {
		OfMapOfResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultMap respjson.Field
		OfMapOfResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultMap  respjson.Field
		raw                                                                           string
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

type ResponseObjectStreamResponseCreated struct {
	Response ResponseObject           `json:"response,required"`
	Type     constant.ResponseCreated `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemAdded struct {
	// Corresponds to the various Message types in the Responses API. They are all
	// under one type because the Responses API gives them all the same "type" value,
	// and there is no way to tell them apart in certain scenarios.
	Item           ResponseObjectStreamResponseOutputItemAddedItemUnion `json:"item,required"`
	OutputIndex    int64                                                `json:"output_index,required"`
	ResponseID     string                                               `json:"response_id,required"`
	SequenceNumber int64                                                `json:"sequence_number,required"`
	Type           constant.ResponseOutputItemAdded                     `json:"type,required"`
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
// [ResponseObjectStreamResponseOutputItemAddedItemMcpListTools].
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
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall].
	Results   []map[string]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion `json:"results"`
	Arguments string                                                                                `json:"arguments"`
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemAddedItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMessage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemFunctionCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemAddedItemMcpListTools:
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

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputText struct {
	Text string             `json:"text,required"`
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

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail `json:"detail,required"`
	Type     constant.InputImage                                                                    `json:"type,required"`
	ImageURL string                                                                                 `json:"image_url"`
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

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail string

const (
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetailLow  ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail = "low"
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetailHigh ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail = "high"
	ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetailAuto ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemInputImageDetail = "auto"
)

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

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string                `json:"file_id,required"`
	Filename string                `json:"filename,required"`
	Index    int64                 `json:"index,required"`
	Type     constant.FileCitation `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemAddedItemMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64                `json:"end_index,required"`
	StartIndex int64                `json:"start_index,required"`
	Title      string               `json:"title,required"`
	Type       constant.URLCitation `json:"type,required"`
	URL        string               `json:"url,required"`
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

type ResponseObjectStreamResponseOutputItemAddedItemWebSearchCall struct {
	ID     string                 `json:"id,required"`
	Status string                 `json:"status,required"`
	Type   constant.WebSearchCall `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCall struct {
	ID      string                                                                                `json:"id,required"`
	Queries []string                                                                              `json:"queries,required"`
	Status  string                                                                                `json:"status,required"`
	Type    constant.FileSearchCall                                                               `json:"type,required"`
	Results []map[string]ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion `json:"results"`
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

// ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion struct {
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

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemAddedItemFileSearchCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemAddedItemFunctionCall struct {
	Arguments string                `json:"arguments,required"`
	CallID    string                `json:"call_id,required"`
	Name      string                `json:"name,required"`
	Type      constant.FunctionCall `json:"type,required"`
	ID        string                `json:"id"`
	Status    string                `json:"status"`
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

type ResponseObjectStreamResponseOutputItemAddedItemMcpCall struct {
	ID          string           `json:"id,required"`
	Arguments   string           `json:"arguments,required"`
	Name        string           `json:"name,required"`
	ServerLabel string           `json:"server_label,required"`
	Type        constant.McpCall `json:"type,required"`
	Error       string           `json:"error"`
	Output      string           `json:"output"`
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

type ResponseObjectStreamResponseOutputItemAddedItemMcpListTools struct {
	ID          string                                                            `json:"id,required"`
	ServerLabel string                                                            `json:"server_label,required"`
	Tools       []ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool `json:"tools,required"`
	Type        constant.McpListTools                                             `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsTool struct {
	InputSchema map[string]ResponseObjectStreamResponseOutputItemAddedItemMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	Name        string                                                                                     `json:"name,required"`
	Description string                                                                                     `json:"description"`
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

type ResponseObjectStreamResponseOutputItemAddedItemRole string

const (
	ResponseObjectStreamResponseOutputItemAddedItemRoleSystem    ResponseObjectStreamResponseOutputItemAddedItemRole = "system"
	ResponseObjectStreamResponseOutputItemAddedItemRoleDeveloper ResponseObjectStreamResponseOutputItemAddedItemRole = "developer"
	ResponseObjectStreamResponseOutputItemAddedItemRoleUser      ResponseObjectStreamResponseOutputItemAddedItemRole = "user"
	ResponseObjectStreamResponseOutputItemAddedItemRoleAssistant ResponseObjectStreamResponseOutputItemAddedItemRole = "assistant"
)

type ResponseObjectStreamResponseOutputItemDone struct {
	// Corresponds to the various Message types in the Responses API. They are all
	// under one type because the Responses API gives them all the same "type" value,
	// and there is no way to tell them apart in certain scenarios.
	Item           ResponseObjectStreamResponseOutputItemDoneItemUnion `json:"item,required"`
	OutputIndex    int64                                               `json:"output_index,required"`
	ResponseID     string                                              `json:"response_id,required"`
	SequenceNumber int64                                               `json:"sequence_number,required"`
	Type           constant.ResponseOutputItemDone                     `json:"type,required"`
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
// [ResponseObjectStreamResponseOutputItemDoneItemMcpListTools].
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
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall].
	Results   []map[string]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion `json:"results"`
	Arguments string                                                                               `json:"arguments"`
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

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseObjectStreamResponseOutputItemDoneItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMessage:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemFunctionCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpCall:
//	case llamastackclient.ResponseObjectStreamResponseOutputItemDoneItemMcpListTools:
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

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputText struct {
	Text string             `json:"text,required"`
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

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail `json:"detail,required"`
	Type     constant.InputImage                                                                   `json:"type,required"`
	ImageURL string                                                                                `json:"image_url"`
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

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail string

const (
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetailLow  ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail = "low"
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetailHigh ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail = "high"
	ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetailAuto ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemInputImageDetail = "auto"
)

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

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string                `json:"file_id,required"`
	Filename string                `json:"filename,required"`
	Index    int64                 `json:"index,required"`
	Type     constant.FileCitation `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemDoneItemMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64                `json:"end_index,required"`
	StartIndex int64                `json:"start_index,required"`
	Title      string               `json:"title,required"`
	Type       constant.URLCitation `json:"type,required"`
	URL        string               `json:"url,required"`
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

type ResponseObjectStreamResponseOutputItemDoneItemWebSearchCall struct {
	ID     string                 `json:"id,required"`
	Status string                 `json:"status,required"`
	Type   constant.WebSearchCall `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCall struct {
	ID      string                                                                               `json:"id,required"`
	Queries []string                                                                             `json:"queries,required"`
	Status  string                                                                               `json:"status,required"`
	Type    constant.FileSearchCall                                                              `json:"type,required"`
	Results []map[string]ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion `json:"results"`
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

// ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion contains
// all possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion struct {
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

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseObjectStreamResponseOutputItemDoneItemFileSearchCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObjectStreamResponseOutputItemDoneItemFunctionCall struct {
	Arguments string                `json:"arguments,required"`
	CallID    string                `json:"call_id,required"`
	Name      string                `json:"name,required"`
	Type      constant.FunctionCall `json:"type,required"`
	ID        string                `json:"id"`
	Status    string                `json:"status"`
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

type ResponseObjectStreamResponseOutputItemDoneItemMcpCall struct {
	ID          string           `json:"id,required"`
	Arguments   string           `json:"arguments,required"`
	Name        string           `json:"name,required"`
	ServerLabel string           `json:"server_label,required"`
	Type        constant.McpCall `json:"type,required"`
	Error       string           `json:"error"`
	Output      string           `json:"output"`
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

type ResponseObjectStreamResponseOutputItemDoneItemMcpListTools struct {
	ID          string                                                           `json:"id,required"`
	ServerLabel string                                                           `json:"server_label,required"`
	Tools       []ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool `json:"tools,required"`
	Type        constant.McpListTools                                            `json:"type,required"`
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

type ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsTool struct {
	InputSchema map[string]ResponseObjectStreamResponseOutputItemDoneItemMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	Name        string                                                                                    `json:"name,required"`
	Description string                                                                                    `json:"description"`
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

type ResponseObjectStreamResponseOutputItemDoneItemRole string

const (
	ResponseObjectStreamResponseOutputItemDoneItemRoleSystem    ResponseObjectStreamResponseOutputItemDoneItemRole = "system"
	ResponseObjectStreamResponseOutputItemDoneItemRoleDeveloper ResponseObjectStreamResponseOutputItemDoneItemRole = "developer"
	ResponseObjectStreamResponseOutputItemDoneItemRoleUser      ResponseObjectStreamResponseOutputItemDoneItemRole = "user"
	ResponseObjectStreamResponseOutputItemDoneItemRoleAssistant ResponseObjectStreamResponseOutputItemDoneItemRole = "assistant"
)

type ResponseObjectStreamResponseOutputTextDelta struct {
	ContentIndex   int64                            `json:"content_index,required"`
	Delta          string                           `json:"delta,required"`
	ItemID         string                           `json:"item_id,required"`
	OutputIndex    int64                            `json:"output_index,required"`
	SequenceNumber int64                            `json:"sequence_number,required"`
	Type           constant.ResponseOutputTextDelta `json:"type,required"`
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

type ResponseObjectStreamResponseOutputTextDone struct {
	ContentIndex   int64                           `json:"content_index,required"`
	ItemID         string                          `json:"item_id,required"`
	OutputIndex    int64                           `json:"output_index,required"`
	SequenceNumber int64                           `json:"sequence_number,required"`
	Text           string                          `json:"text,required"`
	Type           constant.ResponseOutputTextDone `json:"type,required"`
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

type ResponseObjectStreamResponseFunctionCallArgumentsDelta struct {
	Delta          string                                      `json:"delta,required"`
	ItemID         string                                      `json:"item_id,required"`
	OutputIndex    int64                                       `json:"output_index,required"`
	SequenceNumber int64                                       `json:"sequence_number,required"`
	Type           constant.ResponseFunctionCallArgumentsDelta `json:"type,required"`
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

type ResponseObjectStreamResponseFunctionCallArgumentsDone struct {
	Arguments      string                                     `json:"arguments,required"`
	ItemID         string                                     `json:"item_id,required"`
	OutputIndex    int64                                      `json:"output_index,required"`
	SequenceNumber int64                                      `json:"sequence_number,required"`
	Type           constant.ResponseFunctionCallArgumentsDone `json:"type,required"`
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

type ResponseObjectStreamResponseWebSearchCallInProgress struct {
	ItemID         string                                   `json:"item_id,required"`
	OutputIndex    int64                                    `json:"output_index,required"`
	SequenceNumber int64                                    `json:"sequence_number,required"`
	Type           constant.ResponseWebSearchCallInProgress `json:"type,required"`
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

type ResponseObjectStreamResponseWebSearchCallCompleted struct {
	ItemID         string                                  `json:"item_id,required"`
	OutputIndex    int64                                   `json:"output_index,required"`
	SequenceNumber int64                                   `json:"sequence_number,required"`
	Type           constant.ResponseWebSearchCallCompleted `json:"type,required"`
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

type ResponseObjectStreamResponseMcpCallInProgress struct {
	ItemID         string                             `json:"item_id,required"`
	OutputIndex    int64                              `json:"output_index,required"`
	SequenceNumber int64                              `json:"sequence_number,required"`
	Type           constant.ResponseMcpCallInProgress `json:"type,required"`
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

type ResponseObjectStreamResponseMcpCallFailed struct {
	SequenceNumber int64                          `json:"sequence_number,required"`
	Type           constant.ResponseMcpCallFailed `json:"type,required"`
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

type ResponseObjectStreamResponseMcpCallCompleted struct {
	SequenceNumber int64                             `json:"sequence_number,required"`
	Type           constant.ResponseMcpCallCompleted `json:"type,required"`
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

type ResponseObjectStreamResponseCompleted struct {
	Response ResponseObject             `json:"response,required"`
	Type     constant.ResponseCompleted `json:"type,required"`
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

type ResponseListResponse struct {
	Data    []ResponseListResponseData `json:"data,required"`
	FirstID string                     `json:"first_id,required"`
	HasMore bool                       `json:"has_more,required"`
	LastID  string                     `json:"last_id,required"`
	Object  constant.List              `json:"object,required"`
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
func (r ResponseListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseData struct {
	ID                 string                                `json:"id,required"`
	CreatedAt          int64                                 `json:"created_at,required"`
	Input              []ResponseListResponseDataInputUnion  `json:"input,required"`
	Model              string                                `json:"model,required"`
	Object             constant.Response                     `json:"object,required"`
	Output             []ResponseListResponseDataOutputUnion `json:"output,required"`
	ParallelToolCalls  bool                                  `json:"parallel_tool_calls,required"`
	Status             string                                `json:"status,required"`
	Text               ResponseListResponseDataText          `json:"text,required"`
	Error              ResponseListResponseDataError         `json:"error"`
	PreviousResponseID string                                `json:"previous_response_id"`
	Temperature        float64                               `json:"temperature"`
	TopP               float64                               `json:"top_p"`
	Truncation         string                                `json:"truncation"`
	User               string                                `json:"user"`
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
		User               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataInputUnion contains all possible properties and values
// from
// [ResponseListResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall],
// [ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall],
// [ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall],
// [ResponseListResponseDataInputOpenAIResponseInputFunctionToolCallOutput],
// [ResponseListResponseDataInputOpenAIResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseDataInputUnion struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall].
	Results []map[string]ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion `json:"results"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall].
	Arguments string `json:"arguments"`
	CallID    string `json:"call_id"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall].
	Name string `json:"name"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseInputFunctionToolCallOutput].
	Output string `json:"output"`
	// This field is from variant [ResponseListResponseDataInputOpenAIResponseMessage].
	Content ResponseListResponseDataInputOpenAIResponseMessageContentUnion `json:"content"`
	// This field is from variant [ResponseListResponseDataInputOpenAIResponseMessage].
	Role ResponseListResponseDataInputOpenAIResponseMessageRole `json:"role"`
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

func (u ResponseListResponseDataInputUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v ResponseListResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseListResponseDataInputOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputUnion) AsOpenAIResponseMessage() (v ResponseListResponseDataInputOpenAIResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseDataInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall struct {
	ID     string                 `json:"id,required"`
	Status string                 `json:"status,required"`
	Type   constant.WebSearchCall `json:"type,required"`
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
func (r ResponseListResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseOutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall struct {
	ID      string                                                                                             `json:"id,required"`
	Queries []string                                                                                           `json:"queries,required"`
	Status  string                                                                                             `json:"status,required"`
	Type    constant.FileSearchCall                                                                            `json:"type,required"`
	Results []map[string]ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion `json:"results"`
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
func (r ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion
// contains all possible properties and values from [bool], [float64], [string],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion struct {
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

func (u ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataInputOpenAIResponseOutputMessageFileSearchToolCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall struct {
	Arguments string                `json:"arguments,required"`
	CallID    string                `json:"call_id,required"`
	Name      string                `json:"name,required"`
	Type      constant.FunctionCall `json:"type,required"`
	ID        string                `json:"id"`
	Status    string                `json:"status"`
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
func (r ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseOutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseListResponseDataInputOpenAIResponseInputFunctionToolCallOutput struct {
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
func (r ResponseListResponseDataInputOpenAIResponseInputFunctionToolCallOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseListResponseDataInputOpenAIResponseMessage struct {
	Content ResponseListResponseDataInputOpenAIResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseListResponseDataInputOpenAIResponseMessageRole `json:"role,required"`
	Type   constant.Message                                       `json:"type,required"`
	ID     string                                                 `json:"id"`
	Status string                                                 `json:"status"`
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
func (r ResponseListResponseDataInputOpenAIResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataInputOpenAIResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataInputOpenAIResponseMessageContentUnion contains all
// possible properties and values from [string],
// [[]ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion],
// [[]ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfResponseListResponseDataInputOpenAIResponseMessageContentArray OfVariant2]
type ResponseListResponseDataInputOpenAIResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion]
	// instead of an object.
	OfResponseListResponseDataInputOpenAIResponseMessageContentArray []ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem] instead
	// of an object.
	OfVariant2 []ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                                         respjson.Field
		OfResponseListResponseDataInputOpenAIResponseMessageContentArray respjson.Field
		OfVariant2                                                       respjson.Field
		raw                                                              string
	} `json:"-"`
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentUnion) AsResponseListResponseDataInputOpenAIResponseMessageContentArray() (v []ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentUnion) AsVariant2() (v []ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataInputOpenAIResponseMessageContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataInputOpenAIResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion contains
// all possible properties and values from
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText],
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage].
//
// Use the
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage].
	Detail ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponseDataInputOpenAIResponseMessageContentArrayItem is
// implemented by each variant of
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion.AsAny]
type anyResponseListResponseDataInputOpenAIResponseMessageContentArrayItem interface {
	implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion()
}

func (ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText) implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion() {
}
func (ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText:
//	case llamastackclient.ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion) AsAny() anyResponseListResponseDataInputOpenAIResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion) AsInputText() (v ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion) AsInputImage() (v ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText struct {
	Text string             `json:"text,required"`
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
	Type     constant.InputImage                                                                `json:"type,required"`
	ImageURL string                                                                             `json:"image_url"`
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail string

const (
	ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetailLow  ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail = "low"
	ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetailHigh ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail = "high"
	ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetailAuto ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemInputImageDetail = "auto"
)

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetail string

const (
	ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetailLow  ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetail = "low"
	ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetailHigh ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetail = "high"
	ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetailAuto ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemDetail = "auto"
)

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem struct {
	Annotations []ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                              `json:"text,required"`
	Type        constant.OutputText                                                                 `json:"type,required"`
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion
// contains all possible properties and values from
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation],
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation],
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotation
// is implemented by each variant of
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion]
// to add type safety for the return type of
// [ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotation interface {
	implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion()
}

func (ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) implResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsAny() anyResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotation {
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

func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string                `json:"file_id,required"`
	Filename string                `json:"filename,required"`
	Index    int64                 `json:"index,required"`
	Type     constant.FileCitation `json:"type,required"`
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64                `json:"end_index,required"`
	StartIndex int64                `json:"start_index,required"`
	Title      string               `json:"title,required"`
	Type       constant.URLCitation `json:"type,required"`
	URL        string               `json:"url,required"`
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataInputOpenAIResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataInputOpenAIResponseMessageRole string

const (
	ResponseListResponseDataInputOpenAIResponseMessageRoleSystem    ResponseListResponseDataInputOpenAIResponseMessageRole = "system"
	ResponseListResponseDataInputOpenAIResponseMessageRoleDeveloper ResponseListResponseDataInputOpenAIResponseMessageRole = "developer"
	ResponseListResponseDataInputOpenAIResponseMessageRoleUser      ResponseListResponseDataInputOpenAIResponseMessageRole = "user"
	ResponseListResponseDataInputOpenAIResponseMessageRoleAssistant ResponseListResponseDataInputOpenAIResponseMessageRole = "assistant"
)

type ResponseListResponseDataInputRole string

const (
	ResponseListResponseDataInputRoleSystem    ResponseListResponseDataInputRole = "system"
	ResponseListResponseDataInputRoleDeveloper ResponseListResponseDataInputRole = "developer"
	ResponseListResponseDataInputRoleUser      ResponseListResponseDataInputRole = "user"
	ResponseListResponseDataInputRoleAssistant ResponseListResponseDataInputRole = "assistant"
)

// ResponseListResponseDataOutputUnion contains all possible properties and values
// from [ResponseListResponseDataOutputMessage],
// [ResponseListResponseDataOutputWebSearchCall],
// [ResponseListResponseDataOutputFileSearchCall],
// [ResponseListResponseDataOutputFunctionCall],
// [ResponseListResponseDataOutputMcpCall],
// [ResponseListResponseDataOutputMcpListTools].
//
// Use the [ResponseListResponseDataOutputUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseDataOutputUnion struct {
	// This field is from variant [ResponseListResponseDataOutputMessage].
	Content ResponseListResponseDataOutputMessageContentUnion `json:"content"`
	// This field is from variant [ResponseListResponseDataOutputMessage].
	Role ResponseListResponseDataOutputMessageRole `json:"role"`
	// Any of "message", "web_search_call", "file_search_call", "function_call",
	// "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [ResponseListResponseDataOutputFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ResponseListResponseDataOutputFileSearchCall].
	Results   []map[string]ResponseListResponseDataOutputFileSearchCallResultUnion `json:"results"`
	Arguments string                                                               `json:"arguments"`
	// This field is from variant [ResponseListResponseDataOutputFunctionCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [ResponseListResponseDataOutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [ResponseListResponseDataOutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [ResponseListResponseDataOutputMcpListTools].
	Tools []ResponseListResponseDataOutputMcpListToolsTool `json:"tools"`
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

// anyResponseListResponseDataOutput is implemented by each variant of
// [ResponseListResponseDataOutputUnion] to add type safety for the return type of
// [ResponseListResponseDataOutputUnion.AsAny]
type anyResponseListResponseDataOutput interface {
	implResponseListResponseDataOutputUnion()
}

func (ResponseListResponseDataOutputMessage) implResponseListResponseDataOutputUnion()        {}
func (ResponseListResponseDataOutputWebSearchCall) implResponseListResponseDataOutputUnion()  {}
func (ResponseListResponseDataOutputFileSearchCall) implResponseListResponseDataOutputUnion() {}
func (ResponseListResponseDataOutputFunctionCall) implResponseListResponseDataOutputUnion()   {}
func (ResponseListResponseDataOutputMcpCall) implResponseListResponseDataOutputUnion()        {}
func (ResponseListResponseDataOutputMcpListTools) implResponseListResponseDataOutputUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseDataOutputUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseDataOutputMessage:
//	case llamastackclient.ResponseListResponseDataOutputWebSearchCall:
//	case llamastackclient.ResponseListResponseDataOutputFileSearchCall:
//	case llamastackclient.ResponseListResponseDataOutputFunctionCall:
//	case llamastackclient.ResponseListResponseDataOutputMcpCall:
//	case llamastackclient.ResponseListResponseDataOutputMcpListTools:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseDataOutputUnion) AsAny() anyResponseListResponseDataOutput {
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
	}
	return nil
}

func (u ResponseListResponseDataOutputUnion) AsMessage() (v ResponseListResponseDataOutputMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputUnion) AsWebSearchCall() (v ResponseListResponseDataOutputWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputUnion) AsFileSearchCall() (v ResponseListResponseDataOutputFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputUnion) AsFunctionCall() (v ResponseListResponseDataOutputFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputUnion) AsMcpCall() (v ResponseListResponseDataOutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputUnion) AsMcpListTools() (v ResponseListResponseDataOutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseDataOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseListResponseDataOutputMessage struct {
	Content ResponseListResponseDataOutputMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseListResponseDataOutputMessageRole `json:"role,required"`
	Type   constant.Message                          `json:"type,required"`
	ID     string                                    `json:"id"`
	Status string                                    `json:"status"`
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
func (r ResponseListResponseDataOutputMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataOutputMessageContentUnion contains all possible
// properties and values from [string],
// [[]ResponseListResponseDataOutputMessageContentArrayItemUnion],
// [[]ResponseListResponseDataOutputMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfResponseListResponseDataOutputMessageContentArray
// OfVariant2]
type ResponseListResponseDataOutputMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseDataOutputMessageContentArrayItemUnion] instead of an
	// object.
	OfResponseListResponseDataOutputMessageContentArray []ResponseListResponseDataOutputMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseListResponseDataOutputMessageContentArrayItem] instead of an object.
	OfVariant2 []ResponseListResponseDataOutputMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                            respjson.Field
		OfResponseListResponseDataOutputMessageContentArray respjson.Field
		OfVariant2                                          respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u ResponseListResponseDataOutputMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMessageContentUnion) AsResponseListResponseDataOutputMessageContentArray() (v []ResponseListResponseDataOutputMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMessageContentUnion) AsVariant2() (v []ResponseListResponseDataOutputMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataOutputMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseDataOutputMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataOutputMessageContentArrayItemUnion contains all possible
// properties and values from
// [ResponseListResponseDataOutputMessageContentArrayItemInputText],
// [ResponseListResponseDataOutputMessageContentArrayItemInputImage].
//
// Use the [ResponseListResponseDataOutputMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseDataOutputMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ResponseListResponseDataOutputMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ResponseListResponseDataOutputMessageContentArrayItemInputImage].
	Detail ResponseListResponseDataOutputMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ResponseListResponseDataOutputMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseListResponseDataOutputMessageContentArrayItem is implemented by each
// variant of [ResponseListResponseDataOutputMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ResponseListResponseDataOutputMessageContentArrayItemUnion.AsAny]
type anyResponseListResponseDataOutputMessageContentArrayItem interface {
	implResponseListResponseDataOutputMessageContentArrayItemUnion()
}

func (ResponseListResponseDataOutputMessageContentArrayItemInputText) implResponseListResponseDataOutputMessageContentArrayItemUnion() {
}
func (ResponseListResponseDataOutputMessageContentArrayItemInputImage) implResponseListResponseDataOutputMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseDataOutputMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseDataOutputMessageContentArrayItemInputText:
//	case llamastackclient.ResponseListResponseDataOutputMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseDataOutputMessageContentArrayItemUnion) AsAny() anyResponseListResponseDataOutputMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseListResponseDataOutputMessageContentArrayItemUnion) AsInputText() (v ResponseListResponseDataOutputMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMessageContentArrayItemUnion) AsInputImage() (v ResponseListResponseDataOutputMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataOutputMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataOutputMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemInputText struct {
	Text string             `json:"text,required"`
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
func (r ResponseListResponseDataOutputMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataOutputMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseListResponseDataOutputMessageContentArrayItemInputImageDetail `json:"detail,required"`
	Type     constant.InputImage                                                   `json:"type,required"`
	ImageURL string                                                                `json:"image_url"`
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
func (r ResponseListResponseDataOutputMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataOutputMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemInputImageDetail string

const (
	ResponseListResponseDataOutputMessageContentArrayItemInputImageDetailLow  ResponseListResponseDataOutputMessageContentArrayItemInputImageDetail = "low"
	ResponseListResponseDataOutputMessageContentArrayItemInputImageDetailHigh ResponseListResponseDataOutputMessageContentArrayItemInputImageDetail = "high"
	ResponseListResponseDataOutputMessageContentArrayItemInputImageDetailAuto ResponseListResponseDataOutputMessageContentArrayItemInputImageDetail = "auto"
)

type ResponseListResponseDataOutputMessageContentArrayItemDetail string

const (
	ResponseListResponseDataOutputMessageContentArrayItemDetailLow  ResponseListResponseDataOutputMessageContentArrayItemDetail = "low"
	ResponseListResponseDataOutputMessageContentArrayItemDetailHigh ResponseListResponseDataOutputMessageContentArrayItemDetail = "high"
	ResponseListResponseDataOutputMessageContentArrayItemDetailAuto ResponseListResponseDataOutputMessageContentArrayItemDetail = "auto"
)

type ResponseListResponseDataOutputMessageContentArrayItem struct {
	Annotations []ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                 `json:"text,required"`
	Type        constant.OutputText                                                    `json:"type,required"`
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
func (r ResponseListResponseDataOutputMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion contains
// all possible properties and values from
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation],
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation],
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyResponseListResponseDataOutputMessageContentArrayItemAnnotation is
// implemented by each variant of
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion] to add
// type safety for the return type of
// [ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseListResponseDataOutputMessageContentArrayItemAnnotation interface {
	implResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion()
}

func (ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation) implResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation) implResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) implResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}
func (ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath) implResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) AsAny() anyResponseListResponseDataOutputMessageContentArrayItemAnnotation {
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

func (u ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataOutputMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string                `json:"file_id,required"`
	Filename string                `json:"filename,required"`
	Index    int64                 `json:"index,required"`
	Type     constant.FileCitation `json:"type,required"`
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
func (r ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataOutputMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64                `json:"end_index,required"`
	StartIndex int64                `json:"start_index,required"`
	Title      string               `json:"title,required"`
	Type       constant.URLCitation `json:"type,required"`
	URL        string               `json:"url,required"`
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
func (r ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataOutputMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataOutputMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath struct {
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
func (r ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseListResponseDataOutputMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMessageRole string

const (
	ResponseListResponseDataOutputMessageRoleSystem    ResponseListResponseDataOutputMessageRole = "system"
	ResponseListResponseDataOutputMessageRoleDeveloper ResponseListResponseDataOutputMessageRole = "developer"
	ResponseListResponseDataOutputMessageRoleUser      ResponseListResponseDataOutputMessageRole = "user"
	ResponseListResponseDataOutputMessageRoleAssistant ResponseListResponseDataOutputMessageRole = "assistant"
)

type ResponseListResponseDataOutputWebSearchCall struct {
	ID     string                 `json:"id,required"`
	Status string                 `json:"status,required"`
	Type   constant.WebSearchCall `json:"type,required"`
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
func (r ResponseListResponseDataOutputWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputFileSearchCall struct {
	ID      string                                                               `json:"id,required"`
	Queries []string                                                             `json:"queries,required"`
	Status  string                                                               `json:"status,required"`
	Type    constant.FileSearchCall                                              `json:"type,required"`
	Results []map[string]ResponseListResponseDataOutputFileSearchCallResultUnion `json:"results"`
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
func (r ResponseListResponseDataOutputFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataOutputFileSearchCallResultUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseDataOutputFileSearchCallResultUnion struct {
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

func (u ResponseListResponseDataOutputFileSearchCallResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputFileSearchCallResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputFileSearchCallResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputFileSearchCallResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataOutputFileSearchCallResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseDataOutputFileSearchCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputFunctionCall struct {
	Arguments string                `json:"arguments,required"`
	CallID    string                `json:"call_id,required"`
	Name      string                `json:"name,required"`
	Type      constant.FunctionCall `json:"type,required"`
	ID        string                `json:"id"`
	Status    string                `json:"status"`
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
func (r ResponseListResponseDataOutputFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMcpCall struct {
	ID          string           `json:"id,required"`
	Arguments   string           `json:"arguments,required"`
	Name        string           `json:"name,required"`
	ServerLabel string           `json:"server_label,required"`
	Type        constant.McpCall `json:"type,required"`
	Error       string           `json:"error"`
	Output      string           `json:"output"`
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
func (r ResponseListResponseDataOutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMcpListTools struct {
	ID          string                                           `json:"id,required"`
	ServerLabel string                                           `json:"server_label,required"`
	Tools       []ResponseListResponseDataOutputMcpListToolsTool `json:"tools,required"`
	Type        constant.McpListTools                            `json:"type,required"`
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
func (r ResponseListResponseDataOutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputMcpListToolsTool struct {
	InputSchema map[string]ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	Name        string                                                                    `json:"name,required"`
	Description string                                                                    `json:"description"`
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
func (r ResponseListResponseDataOutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataOutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion struct {
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

func (u ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ResponseListResponseDataOutputMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataOutputRole string

const (
	ResponseListResponseDataOutputRoleSystem    ResponseListResponseDataOutputRole = "system"
	ResponseListResponseDataOutputRoleDeveloper ResponseListResponseDataOutputRole = "developer"
	ResponseListResponseDataOutputRoleUser      ResponseListResponseDataOutputRole = "user"
	ResponseListResponseDataOutputRoleAssistant ResponseListResponseDataOutputRole = "assistant"
)

type ResponseListResponseDataText struct {
	// Configuration for Responses API text format.
	Format ResponseListResponseDataTextFormat `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseListResponseDataText) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Responses API text format.
type ResponseListResponseDataTextFormat struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseListResponseDataTextFormatType `json:"type,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description string `json:"description"`
	// The name of the response format. Only used for json_schema.
	Name string `json:"name"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseListResponseDataTextFormatSchemaUnion `json:"schema"`
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
func (r ResponseListResponseDataTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "text", "json_schema", or "json_object" to identify the format type
type ResponseListResponseDataTextFormatType string

const (
	ResponseListResponseDataTextFormatTypeText       ResponseListResponseDataTextFormatType = "text"
	ResponseListResponseDataTextFormatTypeJsonSchema ResponseListResponseDataTextFormatType = "json_schema"
	ResponseListResponseDataTextFormatTypeJsonObject ResponseListResponseDataTextFormatType = "json_object"
)

// ResponseListResponseDataTextFormatSchemaUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseListResponseDataTextFormatSchemaUnion struct {
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

func (u ResponseListResponseDataTextFormatSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataTextFormatSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataTextFormatSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseListResponseDataTextFormatSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseListResponseDataTextFormatSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseListResponseDataTextFormatSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseListResponseDataError struct {
	Code    string `json:"code,required"`
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
func (r ResponseListResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *ResponseListResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewParams struct {
	// Input message(s) to create the response.
	Input ResponseNewParamsInputUnion `json:"input,omitzero,required"`
	// The underlying LLM used for completions.
	Model         string            `json:"model,required"`
	Instructions  param.Opt[string] `json:"instructions,omitzero"`
	MaxInferIters param.Opt[int64]  `json:"max_infer_iters,omitzero"`
	// (Optional) if specified, the new response will be a continuation of the previous
	// response. This can be used to easily fork-off new responses from existing
	// responses.
	PreviousResponseID param.Opt[string]            `json:"previous_response_id,omitzero"`
	Store              param.Opt[bool]              `json:"store,omitzero"`
	Temperature        param.Opt[float64]           `json:"temperature,omitzero"`
	Text               ResponseNewParamsText        `json:"text,omitzero"`
	Tools              []ResponseNewParamsToolUnion `json:"tools,omitzero"`
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
	OfOpenAIResponseMessage                         *ResponseNewParamsInputArrayItemOpenAIResponseMessage                         `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseOutputMessageWebSearchToolCall,
		u.OfOpenAIResponseOutputMessageFileSearchToolCall,
		u.OfOpenAIResponseOutputMessageFunctionToolCall,
		u.OfOpenAIResponseInputFunctionToolCallOutput,
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
func (u ResponseNewParamsInputArrayItemUnion) GetResults() []map[string]ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultUnion {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetArguments() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return &vt.Arguments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseNewParamsInputArrayItemUnion) GetName() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return &vt.Name
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
	} else if vt := u.OfOpenAIResponseMessage; vt != nil {
		return (*string)(&vt.Type)
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

// The properties ID, Status, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageWebSearchToolCall struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
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

// The properties ID, Queries, Status, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCall struct {
	ID      string                                                                                               `json:"id,required"`
	Queries []string                                                                                             `json:"queries,omitzero,required"`
	Status  string                                                                                               `json:"status,required"`
	Results []map[string]ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultUnion `json:"results,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFileSearchToolCallResultUnion) asAny() any {
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

// The properties Arguments, CallID, Name, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseOutputMessageFunctionToolCall struct {
	Arguments string            `json:"arguments,required"`
	CallID    string            `json:"call_id,required"`
	Name      string            `json:"name,required"`
	ID        param.Opt[string] `json:"id,omitzero"`
	Status    param.Opt[string] `json:"status,omitzero"`
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

// The properties Text, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputText struct {
	Text string `json:"text,required"`
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

// The properties Detail, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemInputImageDetail `json:"detail,omitzero,required"`
	ImageURL param.Opt[string]                                                                    `json:"image_url,omitzero"`
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

// The properties FileID, Filename, Index, Type are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
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

// The properties EndIndex, StartIndex, Title, Type, URL are required.
type ResponseNewParamsInputArrayItemOpenAIResponseMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
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

type ResponseNewParamsText struct {
	// Configuration for Responses API text format.
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

// Configuration for Responses API text format.
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

// The property Type is required.
type ResponseNewParamsToolOpenAIResponseInputToolWebSearch struct {
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11".
	Type              ResponseNewParamsToolOpenAIResponseInputToolWebSearchType `json:"type,omitzero,required"`
	SearchContextSize param.Opt[string]                                         `json:"search_context_size,omitzero"`
	paramObj
}

func (r ResponseNewParamsToolOpenAIResponseInputToolWebSearch) MarshalJSON() (data []byte, err error) {
	type shadow ResponseNewParamsToolOpenAIResponseInputToolWebSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseNewParamsToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseNewParamsToolOpenAIResponseInputToolWebSearchType string

const (
	ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch                  ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search"
	ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview           ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search_preview"
	ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview2025_03_11 ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search_preview_2025_03_11"
)

// The properties Type, VectorStoreIDs are required.
type ResponseNewParamsToolFileSearch struct {
	VectorStoreIDs []string                                              `json:"vector_store_ids,omitzero,required"`
	MaxNumResults  param.Opt[int64]                                      `json:"max_num_results,omitzero"`
	Filters        map[string]ResponseNewParamsToolFileSearchFilterUnion `json:"filters,omitzero"`
	RankingOptions ResponseNewParamsToolFileSearchRankingOptions         `json:"ranking_options,omitzero"`
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

type ResponseNewParamsToolFileSearchRankingOptions struct {
	Ranker         param.Opt[string]  `json:"ranker,omitzero"`
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

// The properties Name, Type are required.
type ResponseNewParamsToolFunction struct {
	Name        string                                                 `json:"name,required"`
	Description param.Opt[string]                                      `json:"description,omitzero"`
	Strict      param.Opt[bool]                                        `json:"strict,omitzero"`
	Parameters  map[string]ResponseNewParamsToolFunctionParameterUnion `json:"parameters,omitzero"`
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

// The properties RequireApproval, ServerLabel, ServerURL, Type are required.
type ResponseNewParamsToolMcp struct {
	RequireApproval ResponseNewParamsToolMcpRequireApprovalUnion   `json:"require_approval,omitzero,required"`
	ServerLabel     string                                         `json:"server_label,required"`
	ServerURL       string                                         `json:"server_url,required"`
	AllowedTools    ResponseNewParamsToolMcpAllowedToolsUnion      `json:"allowed_tools,omitzero"`
	Headers         map[string]ResponseNewParamsToolMcpHeaderUnion `json:"headers,omitzero"`
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

type ResponseNewParamsToolMcpRequireApprovalApprovalFilter struct {
	Always []string `json:"always,omitzero"`
	Never  []string `json:"never,omitzero"`
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

type ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter struct {
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
