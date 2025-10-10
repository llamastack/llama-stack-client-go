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

// ConversationItemService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationItemService] method instead.
type ConversationItemService struct {
	Options []option.RequestOption
}

// NewConversationItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConversationItemService(opts ...option.RequestOption) (r ConversationItemService) {
	r = ConversationItemService{}
	r.Options = opts
	return
}

// Create items in the conversation.
func (r *ConversationItemService) New(ctx context.Context, conversationID string, body ConversationItemNewParams, opts ...option.RequestOption) (res *ConversationItemNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List items in the conversation.
func (r *ConversationItemService) List(ctx context.Context, conversationID string, query ConversationItemListParams, opts ...option.RequestOption) (res *ConversationItemListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a conversation item.
func (r *ConversationItemService) Get(ctx context.Context, itemID string, query ConversationItemGetParams, opts ...option.RequestOption) (res *ConversationItemGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ConversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("v1/conversations/%s/items/%s", query.ConversationID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of conversation items with pagination.
type ConversationItemNewResponse struct {
	Data    []ConversationItemNewResponseDataUnion `json:"data,required"`
	HasMore bool                                   `json:"has_more,required"`
	Object  string                                 `json:"object,required"`
	FirstID string                                 `json:"first_id"`
	LastID  string                                 `json:"last_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		FirstID     respjson.Field
		LastID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataUnion contains all possible properties and values
// from [ConversationItemNewResponseDataMessage],
// [ConversationItemNewResponseDataFunctionCall],
// [ConversationItemNewResponseDataFileSearchCall],
// [ConversationItemNewResponseDataWebSearchCall],
// [ConversationItemNewResponseDataMcpCall],
// [ConversationItemNewResponseDataMcpListTools].
//
// Use the [ConversationItemNewResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataUnion struct {
	// This field is from variant [ConversationItemNewResponseDataMessage].
	Content ConversationItemNewResponseDataMessageContentUnion `json:"content"`
	// This field is from variant [ConversationItemNewResponseDataMessage].
	Role ConversationItemNewResponseDataMessageRole `json:"role"`
	// Any of "message", "function_call", "file_search_call", "web_search_call",
	// "mcp_call", "mcp_list_tools".
	Type      string `json:"type"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	Arguments string `json:"arguments"`
	// This field is from variant [ConversationItemNewResponseDataFunctionCall].
	CallID string `json:"call_id"`
	Name   string `json:"name"`
	// This field is from variant [ConversationItemNewResponseDataFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemNewResponseDataFileSearchCall].
	Results     []ConversationItemNewResponseDataFileSearchCallResult `json:"results"`
	ServerLabel string                                                `json:"server_label"`
	// This field is from variant [ConversationItemNewResponseDataMcpCall].
	Error string `json:"error"`
	// This field is from variant [ConversationItemNewResponseDataMcpCall].
	Output string `json:"output"`
	// This field is from variant [ConversationItemNewResponseDataMcpListTools].
	Tools []ConversationItemNewResponseDataMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyConversationItemNewResponseData is implemented by each variant of
// [ConversationItemNewResponseDataUnion] to add type safety for the return type of
// [ConversationItemNewResponseDataUnion.AsAny]
type anyConversationItemNewResponseData interface {
	implConversationItemNewResponseDataUnion()
}

func (ConversationItemNewResponseDataMessage) implConversationItemNewResponseDataUnion()        {}
func (ConversationItemNewResponseDataFunctionCall) implConversationItemNewResponseDataUnion()   {}
func (ConversationItemNewResponseDataFileSearchCall) implConversationItemNewResponseDataUnion() {}
func (ConversationItemNewResponseDataWebSearchCall) implConversationItemNewResponseDataUnion()  {}
func (ConversationItemNewResponseDataMcpCall) implConversationItemNewResponseDataUnion()        {}
func (ConversationItemNewResponseDataMcpListTools) implConversationItemNewResponseDataUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessage:
//	case llamastackclient.ConversationItemNewResponseDataFunctionCall:
//	case llamastackclient.ConversationItemNewResponseDataFileSearchCall:
//	case llamastackclient.ConversationItemNewResponseDataWebSearchCall:
//	case llamastackclient.ConversationItemNewResponseDataMcpCall:
//	case llamastackclient.ConversationItemNewResponseDataMcpListTools:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataUnion) AsAny() anyConversationItemNewResponseData {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "function_call":
		return u.AsFunctionCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	}
	return nil
}

func (u ConversationItemNewResponseDataUnion) AsMessage() (v ConversationItemNewResponseDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsFunctionCall() (v ConversationItemNewResponseDataFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsFileSearchCall() (v ConversationItemNewResponseDataFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsWebSearchCall() (v ConversationItemNewResponseDataWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsMcpCall() (v ConversationItemNewResponseDataMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataUnion) AsMcpListTools() (v ConversationItemNewResponseDataMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemNewResponseDataMessage struct {
	Content ConversationItemNewResponseDataMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemNewResponseDataMessageRole `json:"role,required"`
	Type   constant.Message                           `json:"type,required"`
	ID     string                                     `json:"id"`
	Status string                                     `json:"status"`
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
func (r ConversationItemNewResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentUnion contains all possible
// properties and values from [string],
// [[]ConversationItemNewResponseDataMessageContentArrayItemUnion],
// [[]ConversationItemNewResponseDataMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfConversationItemNewResponseDataMessageContentArray
// OfVariant2]
type ConversationItemNewResponseDataMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentArrayItemUnion] instead of an
	// object.
	OfConversationItemNewResponseDataMessageContentArray []ConversationItemNewResponseDataMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemNewResponseDataMessageContentArrayItem] instead of an object.
	OfVariant2 []ConversationItemNewResponseDataMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                             respjson.Field
		OfConversationItemNewResponseDataMessageContentArray respjson.Field
		OfVariant2                                           respjson.Field
		raw                                                  string
	} `json:"-"`
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsConversationItemNewResponseDataMessageContentArray() (v []ConversationItemNewResponseDataMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentUnion) AsVariant2() (v []ConversationItemNewResponseDataMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemNewResponseDataMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentArrayItemUnion contains all
// possible properties and values from
// [ConversationItemNewResponseDataMessageContentArrayItemInputText],
// [ConversationItemNewResponseDataMessageContentArrayItemInputImage].
//
// Use the [ConversationItemNewResponseDataMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputImage].
	Detail ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyConversationItemNewResponseDataMessageContentArrayItem is implemented by each
// variant of [ConversationItemNewResponseDataMessageContentArrayItemUnion] to add
// type safety for the return type of
// [ConversationItemNewResponseDataMessageContentArrayItemUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentArrayItem interface {
	implConversationItemNewResponseDataMessageContentArrayItemUnion()
}

func (ConversationItemNewResponseDataMessageContentArrayItemInputText) implConversationItemNewResponseDataMessageContentArrayItemUnion() {
}
func (ConversationItemNewResponseDataMessageContentArrayItemInputImage) implConversationItemNewResponseDataMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemInputText:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsAny() anyConversationItemNewResponseDataMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsInputText() (v ConversationItemNewResponseDataMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) AsInputImage() (v ConversationItemNewResponseDataMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentArrayItemInputText struct {
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
func (r ConversationItemNewResponseDataMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemNewResponseDataMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ConversationItemNewResponseDataMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail string

const (
	ConversationItemNewResponseDataMessageContentArrayItemInputImageDetailLow  ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail = "low"
	ConversationItemNewResponseDataMessageContentArrayItemInputImageDetailHigh ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail = "high"
	ConversationItemNewResponseDataMessageContentArrayItemInputImageDetailAuto ConversationItemNewResponseDataMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemNewResponseDataMessageContentArrayItemDetail string

const (
	ConversationItemNewResponseDataMessageContentArrayItemDetailLow  ConversationItemNewResponseDataMessageContentArrayItemDetail = "low"
	ConversationItemNewResponseDataMessageContentArrayItemDetailHigh ConversationItemNewResponseDataMessageContentArrayItemDetail = "high"
	ConversationItemNewResponseDataMessageContentArrayItemDetailAuto ConversationItemNewResponseDataMessageContentArrayItemDetail = "auto"
)

type ConversationItemNewResponseDataMessageContentArrayItem struct {
	Annotations []ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                  `json:"text,required"`
	Type        constant.OutputText                                                     `json:"type,required"`
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
func (r ConversationItemNewResponseDataMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion contains
// all possible properties and values from
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation],
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation],
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation],
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyConversationItemNewResponseDataMessageContentArrayItemAnnotation is
// implemented by each variant of
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion] to add
// type safety for the return type of
// [ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion.AsAny]
type anyConversationItemNewResponseDataMessageContentArrayItemAnnotation interface {
	implConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion()
}

func (ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation) implConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation) implConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation) implConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath) implConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) AsAny() anyConversationItemNewResponseDataMessageContentArrayItemAnnotation {
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

func (u ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) AsFilePath() (v ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath struct {
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
func (r ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemNewResponseDataMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataMessageRole string

const (
	ConversationItemNewResponseDataMessageRoleSystem    ConversationItemNewResponseDataMessageRole = "system"
	ConversationItemNewResponseDataMessageRoleDeveloper ConversationItemNewResponseDataMessageRole = "developer"
	ConversationItemNewResponseDataMessageRoleUser      ConversationItemNewResponseDataMessageRole = "user"
	ConversationItemNewResponseDataMessageRoleAssistant ConversationItemNewResponseDataMessageRole = "assistant"
)

// Function tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFunctionCall struct {
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
func (r ConversationItemNewResponseDataFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemNewResponseDataFileSearchCallResult `json:"results"`
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
func (r ConversationItemNewResponseDataFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ConversationItemNewResponseDataFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemNewResponseDataFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ConversationItemNewResponseDataFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemNewResponseDataFileSearchCallResultAttributeUnion struct {
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

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ConversationItemNewResponseDataWebSearchCall struct {
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
func (r ConversationItemNewResponseDataWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemNewResponseDataMcpCall struct {
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
func (r ConversationItemNewResponseDataMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemNewResponseDataMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemNewResponseDataMcpListToolsTool `json:"tools,required"`
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
func (r ConversationItemNewResponseDataMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ConversationItemNewResponseDataMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ConversationItemNewResponseDataMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemNewResponseDataMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion struct {
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

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemNewResponseDataMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewResponseDataRole string

const (
	ConversationItemNewResponseDataRoleSystem    ConversationItemNewResponseDataRole = "system"
	ConversationItemNewResponseDataRoleDeveloper ConversationItemNewResponseDataRole = "developer"
	ConversationItemNewResponseDataRoleUser      ConversationItemNewResponseDataRole = "user"
	ConversationItemNewResponseDataRoleAssistant ConversationItemNewResponseDataRole = "assistant"
)

// List of conversation items with pagination.
type ConversationItemListResponse struct {
	Data    []ConversationItemListResponseDataUnion `json:"data,required"`
	HasMore bool                                    `json:"has_more,required"`
	Object  string                                  `json:"object,required"`
	FirstID string                                  `json:"first_id"`
	LastID  string                                  `json:"last_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		FirstID     respjson.Field
		LastID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseDataUnion contains all possible properties and
// values from [ConversationItemListResponseDataMessage],
// [ConversationItemListResponseDataFunctionCall],
// [ConversationItemListResponseDataFileSearchCall],
// [ConversationItemListResponseDataWebSearchCall],
// [ConversationItemListResponseDataMcpCall],
// [ConversationItemListResponseDataMcpListTools].
//
// Use the [ConversationItemListResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseDataUnion struct {
	// This field is from variant [ConversationItemListResponseDataMessage].
	Content ConversationItemListResponseDataMessageContentUnion `json:"content"`
	// This field is from variant [ConversationItemListResponseDataMessage].
	Role ConversationItemListResponseDataMessageRole `json:"role"`
	// Any of "message", "function_call", "file_search_call", "web_search_call",
	// "mcp_call", "mcp_list_tools".
	Type      string `json:"type"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	Arguments string `json:"arguments"`
	// This field is from variant [ConversationItemListResponseDataFunctionCall].
	CallID string `json:"call_id"`
	Name   string `json:"name"`
	// This field is from variant [ConversationItemListResponseDataFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemListResponseDataFileSearchCall].
	Results     []ConversationItemListResponseDataFileSearchCallResult `json:"results"`
	ServerLabel string                                                 `json:"server_label"`
	// This field is from variant [ConversationItemListResponseDataMcpCall].
	Error string `json:"error"`
	// This field is from variant [ConversationItemListResponseDataMcpCall].
	Output string `json:"output"`
	// This field is from variant [ConversationItemListResponseDataMcpListTools].
	Tools []ConversationItemListResponseDataMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyConversationItemListResponseData is implemented by each variant of
// [ConversationItemListResponseDataUnion] to add type safety for the return type
// of [ConversationItemListResponseDataUnion.AsAny]
type anyConversationItemListResponseData interface {
	implConversationItemListResponseDataUnion()
}

func (ConversationItemListResponseDataMessage) implConversationItemListResponseDataUnion()        {}
func (ConversationItemListResponseDataFunctionCall) implConversationItemListResponseDataUnion()   {}
func (ConversationItemListResponseDataFileSearchCall) implConversationItemListResponseDataUnion() {}
func (ConversationItemListResponseDataWebSearchCall) implConversationItemListResponseDataUnion()  {}
func (ConversationItemListResponseDataMcpCall) implConversationItemListResponseDataUnion()        {}
func (ConversationItemListResponseDataMcpListTools) implConversationItemListResponseDataUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseDataUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseDataMessage:
//	case llamastackclient.ConversationItemListResponseDataFunctionCall:
//	case llamastackclient.ConversationItemListResponseDataFileSearchCall:
//	case llamastackclient.ConversationItemListResponseDataWebSearchCall:
//	case llamastackclient.ConversationItemListResponseDataMcpCall:
//	case llamastackclient.ConversationItemListResponseDataMcpListTools:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseDataUnion) AsAny() anyConversationItemListResponseData {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "function_call":
		return u.AsFunctionCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	}
	return nil
}

func (u ConversationItemListResponseDataUnion) AsMessage() (v ConversationItemListResponseDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataUnion) AsFunctionCall() (v ConversationItemListResponseDataFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataUnion) AsFileSearchCall() (v ConversationItemListResponseDataFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataUnion) AsWebSearchCall() (v ConversationItemListResponseDataWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataUnion) AsMcpCall() (v ConversationItemListResponseDataMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataUnion) AsMcpListTools() (v ConversationItemListResponseDataMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemListResponseDataMessage struct {
	Content ConversationItemListResponseDataMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemListResponseDataMessageRole `json:"role,required"`
	Type   constant.Message                            `json:"type,required"`
	ID     string                                      `json:"id"`
	Status string                                      `json:"status"`
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
func (r ConversationItemListResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseDataMessageContentUnion contains all possible
// properties and values from [string],
// [[]ConversationItemListResponseDataMessageContentArrayItemUnion],
// [[]ConversationItemListResponseDataMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfConversationItemListResponseDataMessageContentArray
// OfVariant2]
type ConversationItemListResponseDataMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseDataMessageContentArrayItemUnion] instead of an
	// object.
	OfConversationItemListResponseDataMessageContentArray []ConversationItemListResponseDataMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemListResponseDataMessageContentArrayItem] instead of an
	// object.
	OfVariant2 []ConversationItemListResponseDataMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                              respjson.Field
		OfConversationItemListResponseDataMessageContentArray respjson.Field
		OfVariant2                                            respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u ConversationItemListResponseDataMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMessageContentUnion) AsConversationItemListResponseDataMessageContentArray() (v []ConversationItemListResponseDataMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMessageContentUnion) AsVariant2() (v []ConversationItemListResponseDataMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseDataMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemListResponseDataMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseDataMessageContentArrayItemUnion contains all
// possible properties and values from
// [ConversationItemListResponseDataMessageContentArrayItemInputText],
// [ConversationItemListResponseDataMessageContentArrayItemInputImage].
//
// Use the [ConversationItemListResponseDataMessageContentArrayItemUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseDataMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ConversationItemListResponseDataMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemListResponseDataMessageContentArrayItemInputImage].
	Detail ConversationItemListResponseDataMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ConversationItemListResponseDataMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyConversationItemListResponseDataMessageContentArrayItem is implemented by
// each variant of [ConversationItemListResponseDataMessageContentArrayItemUnion]
// to add type safety for the return type of
// [ConversationItemListResponseDataMessageContentArrayItemUnion.AsAny]
type anyConversationItemListResponseDataMessageContentArrayItem interface {
	implConversationItemListResponseDataMessageContentArrayItemUnion()
}

func (ConversationItemListResponseDataMessageContentArrayItemInputText) implConversationItemListResponseDataMessageContentArrayItemUnion() {
}
func (ConversationItemListResponseDataMessageContentArrayItemInputImage) implConversationItemListResponseDataMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseDataMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseDataMessageContentArrayItemInputText:
//	case llamastackclient.ConversationItemListResponseDataMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseDataMessageContentArrayItemUnion) AsAny() anyConversationItemListResponseDataMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ConversationItemListResponseDataMessageContentArrayItemUnion) AsInputText() (v ConversationItemListResponseDataMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMessageContentArrayItemUnion) AsInputImage() (v ConversationItemListResponseDataMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseDataMessageContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseDataMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemListResponseDataMessageContentArrayItemInputText struct {
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
func (r ConversationItemListResponseDataMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseDataMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemListResponseDataMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemListResponseDataMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ConversationItemListResponseDataMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseDataMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemListResponseDataMessageContentArrayItemInputImageDetail string

const (
	ConversationItemListResponseDataMessageContentArrayItemInputImageDetailLow  ConversationItemListResponseDataMessageContentArrayItemInputImageDetail = "low"
	ConversationItemListResponseDataMessageContentArrayItemInputImageDetailHigh ConversationItemListResponseDataMessageContentArrayItemInputImageDetail = "high"
	ConversationItemListResponseDataMessageContentArrayItemInputImageDetailAuto ConversationItemListResponseDataMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemListResponseDataMessageContentArrayItemDetail string

const (
	ConversationItemListResponseDataMessageContentArrayItemDetailLow  ConversationItemListResponseDataMessageContentArrayItemDetail = "low"
	ConversationItemListResponseDataMessageContentArrayItemDetailHigh ConversationItemListResponseDataMessageContentArrayItemDetail = "high"
	ConversationItemListResponseDataMessageContentArrayItemDetailAuto ConversationItemListResponseDataMessageContentArrayItemDetail = "auto"
)

type ConversationItemListResponseDataMessageContentArrayItem struct {
	Annotations []ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                                   `json:"text,required"`
	Type        constant.OutputText                                                      `json:"type,required"`
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
func (r ConversationItemListResponseDataMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion contains
// all possible properties and values from
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation],
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation],
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation],
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyConversationItemListResponseDataMessageContentArrayItemAnnotation is
// implemented by each variant of
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion] to add
// type safety for the return type of
// [ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion.AsAny]
type anyConversationItemListResponseDataMessageContentArrayItemAnnotation interface {
	implConversationItemListResponseDataMessageContentArrayItemAnnotationUnion()
}

func (ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation) implConversationItemListResponseDataMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation) implConversationItemListResponseDataMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation) implConversationItemListResponseDataMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath) implConversationItemListResponseDataMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) AsAny() anyConversationItemListResponseDataMessageContentArrayItemAnnotation {
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

func (u ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) AsFilePath() (v ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseDataMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseDataMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseDataMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseDataMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath struct {
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
func (r ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemListResponseDataMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseDataMessageRole string

const (
	ConversationItemListResponseDataMessageRoleSystem    ConversationItemListResponseDataMessageRole = "system"
	ConversationItemListResponseDataMessageRoleDeveloper ConversationItemListResponseDataMessageRole = "developer"
	ConversationItemListResponseDataMessageRoleUser      ConversationItemListResponseDataMessageRole = "user"
	ConversationItemListResponseDataMessageRoleAssistant ConversationItemListResponseDataMessageRole = "assistant"
)

// Function tool call output message for OpenAI responses.
type ConversationItemListResponseDataFunctionCall struct {
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
func (r ConversationItemListResponseDataFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemListResponseDataFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemListResponseDataFileSearchCallResult `json:"results"`
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
func (r ConversationItemListResponseDataFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ConversationItemListResponseDataFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemListResponseDataFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ConversationItemListResponseDataFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseDataFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemListResponseDataFileSearchCallResultAttributeUnion struct {
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

func (u ConversationItemListResponseDataFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseDataFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseDataFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ConversationItemListResponseDataWebSearchCall struct {
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
func (r ConversationItemListResponseDataWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemListResponseDataMcpCall struct {
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
func (r ConversationItemListResponseDataMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemListResponseDataMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemListResponseDataMcpListToolsTool `json:"tools,required"`
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
func (r ConversationItemListResponseDataMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ConversationItemListResponseDataMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ConversationItemListResponseDataMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemListResponseDataMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion struct {
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

func (u ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemListResponseDataMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemListResponseDataRole string

const (
	ConversationItemListResponseDataRoleSystem    ConversationItemListResponseDataRole = "system"
	ConversationItemListResponseDataRoleDeveloper ConversationItemListResponseDataRole = "developer"
	ConversationItemListResponseDataRoleUser      ConversationItemListResponseDataRole = "user"
	ConversationItemListResponseDataRoleAssistant ConversationItemListResponseDataRole = "assistant"
)

// ConversationItemGetResponseUnion contains all possible properties and values
// from [ConversationItemGetResponseMessage],
// [ConversationItemGetResponseFunctionCall],
// [ConversationItemGetResponseFileSearchCall],
// [ConversationItemGetResponseWebSearchCall],
// [ConversationItemGetResponseMcpCall], [ConversationItemGetResponseMcpListTools].
//
// Use the [ConversationItemGetResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseUnion struct {
	// This field is from variant [ConversationItemGetResponseMessage].
	Content ConversationItemGetResponseMessageContentUnion `json:"content"`
	// This field is from variant [ConversationItemGetResponseMessage].
	Role ConversationItemGetResponseMessageRole `json:"role"`
	// Any of "message", "function_call", "file_search_call", "web_search_call",
	// "mcp_call", "mcp_list_tools".
	Type      string `json:"type"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	Arguments string `json:"arguments"`
	// This field is from variant [ConversationItemGetResponseFunctionCall].
	CallID string `json:"call_id"`
	Name   string `json:"name"`
	// This field is from variant [ConversationItemGetResponseFileSearchCall].
	Queries []string `json:"queries"`
	// This field is from variant [ConversationItemGetResponseFileSearchCall].
	Results     []ConversationItemGetResponseFileSearchCallResult `json:"results"`
	ServerLabel string                                            `json:"server_label"`
	// This field is from variant [ConversationItemGetResponseMcpCall].
	Error string `json:"error"`
	// This field is from variant [ConversationItemGetResponseMcpCall].
	Output string `json:"output"`
	// This field is from variant [ConversationItemGetResponseMcpListTools].
	Tools []ConversationItemGetResponseMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

// anyConversationItemGetResponse is implemented by each variant of
// [ConversationItemGetResponseUnion] to add type safety for the return type of
// [ConversationItemGetResponseUnion.AsAny]
type anyConversationItemGetResponse interface {
	implConversationItemGetResponseUnion()
}

func (ConversationItemGetResponseMessage) implConversationItemGetResponseUnion()        {}
func (ConversationItemGetResponseFunctionCall) implConversationItemGetResponseUnion()   {}
func (ConversationItemGetResponseFileSearchCall) implConversationItemGetResponseUnion() {}
func (ConversationItemGetResponseWebSearchCall) implConversationItemGetResponseUnion()  {}
func (ConversationItemGetResponseMcpCall) implConversationItemGetResponseUnion()        {}
func (ConversationItemGetResponseMcpListTools) implConversationItemGetResponseUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseMessage:
//	case llamastackclient.ConversationItemGetResponseFunctionCall:
//	case llamastackclient.ConversationItemGetResponseFileSearchCall:
//	case llamastackclient.ConversationItemGetResponseWebSearchCall:
//	case llamastackclient.ConversationItemGetResponseMcpCall:
//	case llamastackclient.ConversationItemGetResponseMcpListTools:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseUnion) AsAny() anyConversationItemGetResponse {
	switch u.Type {
	case "message":
		return u.AsMessage()
	case "function_call":
		return u.AsFunctionCall()
	case "file_search_call":
		return u.AsFileSearchCall()
	case "web_search_call":
		return u.AsWebSearchCall()
	case "mcp_call":
		return u.AsMcpCall()
	case "mcp_list_tools":
		return u.AsMcpListTools()
	}
	return nil
}

func (u ConversationItemGetResponseUnion) AsMessage() (v ConversationItemGetResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsFunctionCall() (v ConversationItemGetResponseFunctionCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsFileSearchCall() (v ConversationItemGetResponseFileSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsWebSearchCall() (v ConversationItemGetResponseWebSearchCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsMcpCall() (v ConversationItemGetResponseMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseUnion) AsMcpListTools() (v ConversationItemGetResponseMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ConversationItemGetResponseMessage struct {
	Content ConversationItemGetResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemGetResponseMessageRole `json:"role,required"`
	Type   constant.Message                       `json:"type,required"`
	ID     string                                 `json:"id"`
	Status string                                 `json:"status"`
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
func (r ConversationItemGetResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentUnion contains all possible properties
// and values from [string],
// [[]ConversationItemGetResponseMessageContentArrayItemUnion],
// [[]ConversationItemGetResponseMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfConversationItemGetResponseMessageContentArray
// OfVariant2]
type ConversationItemGetResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentArrayItemUnion] instead of an
	// object.
	OfConversationItemGetResponseMessageContentArray []ConversationItemGetResponseMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ConversationItemGetResponseMessageContentArrayItem] instead of an object.
	OfVariant2 []ConversationItemGetResponseMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                                         respjson.Field
		OfConversationItemGetResponseMessageContentArray respjson.Field
		OfVariant2                                       respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u ConversationItemGetResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentUnion) AsConversationItemGetResponseMessageContentArray() (v []ConversationItemGetResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentUnion) AsVariant2() (v []ConversationItemGetResponseMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentArrayItemUnion contains all possible
// properties and values from
// [ConversationItemGetResponseMessageContentArrayItemInputText],
// [ConversationItemGetResponseMessageContentArrayItemInputImage].
//
// Use the [ConversationItemGetResponseMessageContentArrayItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseMessageContentArrayItemUnion struct {
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputImage].
	Detail ConversationItemGetResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyConversationItemGetResponseMessageContentArrayItem is implemented by each
// variant of [ConversationItemGetResponseMessageContentArrayItemUnion] to add type
// safety for the return type of
// [ConversationItemGetResponseMessageContentArrayItemUnion.AsAny]
type anyConversationItemGetResponseMessageContentArrayItem interface {
	implConversationItemGetResponseMessageContentArrayItemUnion()
}

func (ConversationItemGetResponseMessageContentArrayItemInputText) implConversationItemGetResponseMessageContentArrayItemUnion() {
}
func (ConversationItemGetResponseMessageContentArrayItemInputImage) implConversationItemGetResponseMessageContentArrayItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemInputText:
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsAny() anyConversationItemGetResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsInputText() (v ConversationItemGetResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentArrayItemUnion) AsInputImage() (v ConversationItemGetResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationItemGetResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentArrayItemInputText struct {
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
func (r ConversationItemGetResponseMessageContentArrayItemInputText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
type ConversationItemGetResponseMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemGetResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
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
func (r ConversationItemGetResponseMessageContentArrayItemInputImage) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemGetResponseMessageContentArrayItemInputImageDetail string

const (
	ConversationItemGetResponseMessageContentArrayItemInputImageDetailLow  ConversationItemGetResponseMessageContentArrayItemInputImageDetail = "low"
	ConversationItemGetResponseMessageContentArrayItemInputImageDetailHigh ConversationItemGetResponseMessageContentArrayItemInputImageDetail = "high"
	ConversationItemGetResponseMessageContentArrayItemInputImageDetailAuto ConversationItemGetResponseMessageContentArrayItemInputImageDetail = "auto"
)

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemGetResponseMessageContentArrayItemDetail string

const (
	ConversationItemGetResponseMessageContentArrayItemDetailLow  ConversationItemGetResponseMessageContentArrayItemDetail = "low"
	ConversationItemGetResponseMessageContentArrayItemDetailHigh ConversationItemGetResponseMessageContentArrayItemDetail = "high"
	ConversationItemGetResponseMessageContentArrayItemDetailAuto ConversationItemGetResponseMessageContentArrayItemDetail = "auto"
)

type ConversationItemGetResponseMessageContentArrayItem struct {
	Annotations []ConversationItemGetResponseMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                                              `json:"text,required"`
	Type        constant.OutputText                                                 `json:"type,required"`
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
func (r ConversationItemGetResponseMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMessageContentArrayItemAnnotationUnion contains all
// possible properties and values from
// [ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation],
// [ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation],
// [ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation],
// [ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath].
//
// Use the
// [ConversationItemGetResponseMessageContentArrayItemAnnotationUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationItemGetResponseMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation].
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

// anyConversationItemGetResponseMessageContentArrayItemAnnotation is implemented
// by each variant of
// [ConversationItemGetResponseMessageContentArrayItemAnnotationUnion] to add type
// safety for the return type of
// [ConversationItemGetResponseMessageContentArrayItemAnnotationUnion.AsAny]
type anyConversationItemGetResponseMessageContentArrayItemAnnotation interface {
	implConversationItemGetResponseMessageContentArrayItemAnnotationUnion()
}

func (ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation) implConversationItemGetResponseMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation) implConversationItemGetResponseMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation) implConversationItemGetResponseMessageContentArrayItemAnnotationUnion() {
}
func (ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath) implConversationItemGetResponseMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationItemGetResponseMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) AsAny() anyConversationItemGetResponseMessageContentArrayItemAnnotation {
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

func (u ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) AsFilePath() (v ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File citation annotation for referencing specific files in response content.
type ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation struct {
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
func (r ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
type ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation struct {
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
func (r ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
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
func (r ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath struct {
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
func (r ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationItemGetResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseMessageRole string

const (
	ConversationItemGetResponseMessageRoleSystem    ConversationItemGetResponseMessageRole = "system"
	ConversationItemGetResponseMessageRoleDeveloper ConversationItemGetResponseMessageRole = "developer"
	ConversationItemGetResponseMessageRoleUser      ConversationItemGetResponseMessageRole = "user"
	ConversationItemGetResponseMessageRoleAssistant ConversationItemGetResponseMessageRole = "assistant"
)

// Function tool call output message for OpenAI responses.
type ConversationItemGetResponseFunctionCall struct {
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
func (r ConversationItemGetResponseFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
type ConversationItemGetResponseFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// Tool call type identifier, always "file_search_call"
	Type constant.FileSearchCall `json:"type,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemGetResponseFileSearchCallResult `json:"results"`
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
func (r ConversationItemGetResponseFileSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
type ConversationItemGetResponseFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemGetResponseFileSearchCallResultAttributeUnion `json:"attributes,required"`
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
func (r ConversationItemGetResponseFileSearchCallResult) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseFileSearchCallResultAttributeUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemGetResponseFileSearchCallResultAttributeUnion struct {
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

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseFileSearchCallResultAttributeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web search tool call output message for OpenAI responses.
type ConversationItemGetResponseWebSearchCall struct {
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
func (r ConversationItemGetResponseWebSearchCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
type ConversationItemGetResponseMcpCall struct {
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
func (r ConversationItemGetResponseMcpCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
type ConversationItemGetResponseMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemGetResponseMcpListToolsTool `json:"tools,required"`
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
func (r ConversationItemGetResponseMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
type ConversationItemGetResponseMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemGetResponseMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
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
func (r ConversationItemGetResponseMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemGetResponseMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationItemGetResponseMcpListToolsToolInputSchemaUnion contains all
// possible properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ConversationItemGetResponseMcpListToolsToolInputSchemaUnion struct {
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

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationItemGetResponseMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemGetResponseRole string

const (
	ConversationItemGetResponseRoleSystem    ConversationItemGetResponseRole = "system"
	ConversationItemGetResponseRoleDeveloper ConversationItemGetResponseRole = "developer"
	ConversationItemGetResponseRoleUser      ConversationItemGetResponseRole = "user"
	ConversationItemGetResponseRoleAssistant ConversationItemGetResponseRole = "assistant"
)

type ConversationItemNewParams struct {
	// Items to include in the conversation context.
	Items []ConversationItemNewParamsItemUnion `json:"items,omitzero,required"`
	paramObj
}

func (r ConversationItemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemUnion struct {
	OfMessage        *ConversationItemNewParamsItemMessage        `json:",omitzero,inline"`
	OfFunctionCall   *ConversationItemNewParamsItemFunctionCall   `json:",omitzero,inline"`
	OfFileSearchCall *ConversationItemNewParamsItemFileSearchCall `json:",omitzero,inline"`
	OfWebSearchCall  *ConversationItemNewParamsItemWebSearchCall  `json:",omitzero,inline"`
	OfMcpCall        *ConversationItemNewParamsItemMcpCall        `json:",omitzero,inline"`
	OfMcpListTools   *ConversationItemNewParamsItemMcpListTools   `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessage,
		u.OfFunctionCall,
		u.OfFileSearchCall,
		u.OfWebSearchCall,
		u.OfMcpCall,
		u.OfMcpListTools)
}
func (u *ConversationItemNewParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemUnion) asAny() any {
	if !param.IsOmitted(u.OfMessage) {
		return u.OfMessage
	} else if !param.IsOmitted(u.OfFunctionCall) {
		return u.OfFunctionCall
	} else if !param.IsOmitted(u.OfFileSearchCall) {
		return u.OfFileSearchCall
	} else if !param.IsOmitted(u.OfWebSearchCall) {
		return u.OfWebSearchCall
	} else if !param.IsOmitted(u.OfMcpCall) {
		return u.OfMcpCall
	} else if !param.IsOmitted(u.OfMcpListTools) {
		return u.OfMcpListTools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetContent() *ConversationItemNewParamsItemMessageContentUnion {
	if vt := u.OfMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetRole() *string {
	if vt := u.OfMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetCallID() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetQueries() []string {
	if vt := u.OfFileSearchCall; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetResults() []ConversationItemNewParamsItemFileSearchCallResult {
	if vt := u.OfFileSearchCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetError() *string {
	if vt := u.OfMcpCall; vt != nil && vt.Error.Valid() {
		return &vt.Error.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetOutput() *string {
	if vt := u.OfMcpCall; vt != nil && vt.Output.Valid() {
		return &vt.Output.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetTools() []ConversationItemNewParamsItemMcpListToolsTool {
	if vt := u.OfMcpListTools; vt != nil {
		return vt.Tools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetType() *string {
	if vt := u.OfMessage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileSearchCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebSearchCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcpListTools; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetID() *string {
	if vt := u.OfMessage; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfFunctionCall; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfFileSearchCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfWebSearchCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMcpListTools; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetStatus() *string {
	if vt := u.OfMessage; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfFunctionCall; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfFileSearchCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfWebSearchCall; vt != nil {
		return (*string)(&vt.Status)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetArguments() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.Arguments)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.Arguments)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetName() *string {
	if vt := u.OfFunctionCall; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemUnion) GetServerLabel() *string {
	if vt := u.OfMcpCall; vt != nil {
		return (*string)(&vt.ServerLabel)
	} else if vt := u.OfMcpListTools; vt != nil {
		return (*string)(&vt.ServerLabel)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessage]("message"),
		apijson.Discriminator[ConversationItemNewParamsItemFunctionCall]("function_call"),
		apijson.Discriminator[ConversationItemNewParamsItemFileSearchCall]("file_search_call"),
		apijson.Discriminator[ConversationItemNewParamsItemWebSearchCall]("web_search_call"),
		apijson.Discriminator[ConversationItemNewParamsItemMcpCall]("mcp_call"),
		apijson.Discriminator[ConversationItemNewParamsItemMcpListTools]("mcp_list_tools"),
	)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role, Type are required.
type ConversationItemNewParamsItemMessage struct {
	Content ConversationItemNewParamsItemMessageContentUnion `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ConversationItemNewParamsItemMessageRole `json:"role,omitzero,required"`
	ID     param.Opt[string]                        `json:"id,omitzero"`
	Status param.Opt[string]                        `json:"status,omitzero"`
	// This field can be elided, and will marshal its zero value as "message".
	Type constant.Message `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentUnion struct {
	OfString                                      param.Opt[string]                                           `json:",omitzero,inline"`
	OfConversationItemNewsItemMessageContentArray []ConversationItemNewParamsItemMessageContentArrayItemUnion `json:",omitzero,inline"`
	OfVariant2                                    []ConversationItemNewParamsItemMessageContentArrayItem      `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConversationItemNewsItemMessageContentArray, u.OfVariant2)
}
func (u *ConversationItemNewParamsItemMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfConversationItemNewsItemMessageContentArray) {
		return &u.OfConversationItemNewsItemMessageContentArray
	} else if !param.IsOmitted(u.OfVariant2) {
		return &u.OfVariant2
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentArrayItemUnion struct {
	OfInputText  *ConversationItemNewParamsItemMessageContentArrayItemInputText  `json:",omitzero,inline"`
	OfInputImage *ConversationItemNewParamsItemMessageContentArrayItemInputImage `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage)
}
func (u *ConversationItemNewParamsItemMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfInputText) {
		return u.OfInputText
	} else if !param.IsOmitted(u.OfInputImage) {
		return u.OfInputImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Detail)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemUnion) GetType() *string {
	if vt := u.OfInputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentArrayItemUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemInputText]("input_text"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemInputImage]("input_image"),
	)
}

// Text content for input messages in OpenAI response format.
//
// The properties Text, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItemInputText struct {
	// The text content of the input message
	Text string `json:"text,required"`
	// Content type identifier, always "input_text"
	//
	// This field can be elided, and will marshal its zero value as "input_text".
	Type constant.InputText `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItemInputText) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemInputText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content for input messages in OpenAI response format.
//
// The properties Detail, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItemInputImage struct {
	// Level of detail for image processing, can be "low", "high", or "auto"
	//
	// Any of "low", "high", "auto".
	Detail ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail `json:"detail,omitzero,required"`
	// (Optional) URL of the image content
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Content type identifier, always "input_image"
	//
	// This field can be elided, and will marshal its zero value as "input_image".
	Type constant.InputImage `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItemInputImage) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemInputImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Level of detail for image processing, can be "low", "high", or "auto"
type ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail string

const (
	ConversationItemNewParamsItemMessageContentArrayItemInputImageDetailLow  ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail = "low"
	ConversationItemNewParamsItemMessageContentArrayItemInputImageDetailHigh ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail = "high"
	ConversationItemNewParamsItemMessageContentArrayItemInputImageDetailAuto ConversationItemNewParamsItemMessageContentArrayItemInputImageDetail = "auto"
)

// The properties Annotations, Text, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItem struct {
	Annotations []ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion `json:"annotations,omitzero,required"`
	Text        string                                                                `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "output_text".
	Type constant.OutputText `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion struct {
	OfFileCitation          *ConversationItemNewParamsItemMessageContentArrayItemAnnotationFileCitation          `json:",omitzero,inline"`
	OfURLCitation           *ConversationItemNewParamsItemMessageContentArrayItemAnnotationURLCitation           `json:",omitzero,inline"`
	OfContainerFileCitation *ConversationItemNewParamsItemMessageContentArrayItemAnnotationContainerFileCitation `json:",omitzero,inline"`
	OfFilePath              *ConversationItemNewParamsItemMessageContentArrayItemAnnotationFilePath              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) asAny() any {
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
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetFileID() *string {
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
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetType() *string {
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
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationItemNewParamsItemMessageContentArrayItemAnnotationUnion](
		"type",
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemAnnotationFileCitation]("file_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemAnnotationURLCitation]("url_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemAnnotationContainerFileCitation]("container_file_citation"),
		apijson.Discriminator[ConversationItemNewParamsItemMessageContentArrayItemAnnotationFilePath]("file_path"),
	)
}

// File citation annotation for referencing specific files in response content.
//
// The properties FileID, Filename, Index, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItemAnnotationFileCitation struct {
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

func (r ConversationItemNewParamsItemMessageContentArrayItemAnnotationFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemAnnotationFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL citation annotation for referencing external web resources.
//
// The properties EndIndex, StartIndex, Title, Type, URL are required.
type ConversationItemNewParamsItemMessageContentArrayItemAnnotationURLCitation struct {
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

func (r ConversationItemNewParamsItemMessageContentArrayItemAnnotationURLCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemAnnotationURLCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex, Type are
// required.
type ConversationItemNewParamsItemMessageContentArrayItemAnnotationContainerFileCitation struct {
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

func (r ConversationItemNewParamsItemMessageContentArrayItemAnnotationContainerFileCitation) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemAnnotationContainerFileCitation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FileID, Index, Type are required.
type ConversationItemNewParamsItemMessageContentArrayItemAnnotationFilePath struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// This field can be elided, and will marshal its zero value as "file_path".
	Type constant.FilePath `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMessageContentArrayItemAnnotationFilePath) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMessageContentArrayItemAnnotationFilePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemNewParamsItemMessageRole string

const (
	ConversationItemNewParamsItemMessageRoleSystem    ConversationItemNewParamsItemMessageRole = "system"
	ConversationItemNewParamsItemMessageRoleDeveloper ConversationItemNewParamsItemMessageRole = "developer"
	ConversationItemNewParamsItemMessageRoleUser      ConversationItemNewParamsItemMessageRole = "user"
	ConversationItemNewParamsItemMessageRoleAssistant ConversationItemNewParamsItemMessageRole = "assistant"
)

// Function tool call output message for OpenAI responses.
//
// The properties Arguments, CallID, Name, Type are required.
type ConversationItemNewParamsItemFunctionCall struct {
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

func (r ConversationItemNewParamsItemFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File search tool call output message for OpenAI responses.
//
// The properties ID, Queries, Status, Type are required.
type ConversationItemNewParamsItemFileSearchCall struct {
	// Unique identifier for this tool call
	ID string `json:"id,required"`
	// List of search queries executed
	Queries []string `json:"queries,omitzero,required"`
	// Current status of the file search operation
	Status string `json:"status,required"`
	// (Optional) Search results returned by the file search operation
	Results []ConversationItemNewParamsItemFileSearchCallResult `json:"results,omitzero"`
	// Tool call type identifier, always "file_search_call"
	//
	// This field can be elided, and will marshal its zero value as "file_search_call".
	Type constant.FileSearchCall `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemFileSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFileSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFileSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results returned by the file search operation.
//
// The properties Attributes, FileID, Filename, Score, Text are required.
type ConversationItemNewParamsItemFileSearchCallResult struct {
	// (Optional) Key-value attributes associated with the file
	Attributes map[string]ConversationItemNewParamsItemFileSearchCallResultAttributeUnion `json:"attributes,omitzero,required"`
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

func (r ConversationItemNewParamsItemFileSearchCallResult) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemFileSearchCallResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemFileSearchCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemFileSearchCallResultAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemFileSearchCallResultAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ConversationItemNewParamsItemFileSearchCallResultAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemFileSearchCallResultAttributeUnion) asAny() any {
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

// Web search tool call output message for OpenAI responses.
//
// The properties ID, Status, Type are required.
type ConversationItemNewParamsItemWebSearchCall struct {
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

func (r ConversationItemNewParamsItemWebSearchCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemWebSearchCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemWebSearchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model Context Protocol (MCP) call output message for OpenAI responses.
//
// The properties ID, Arguments, Name, ServerLabel, Type are required.
type ConversationItemNewParamsItemMcpCall struct {
	// Unique identifier for this MCP call
	ID string `json:"id,required"`
	// JSON string containing the MCP call arguments
	Arguments string `json:"arguments,required"`
	// Name of the MCP method being called
	Name string `json:"name,required"`
	// Label identifying the MCP server handling the call
	ServerLabel string `json:"server_label,required"`
	// (Optional) Error message if the MCP call failed
	Error param.Opt[string] `json:"error,omitzero"`
	// (Optional) Output result from the successful MCP call
	Output param.Opt[string] `json:"output,omitzero"`
	// Tool call type identifier, always "mcp_call"
	//
	// This field can be elided, and will marshal its zero value as "mcp_call".
	Type constant.McpCall `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpCall) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MCP list tools output message containing available tools from an MCP server.
//
// The properties ID, ServerLabel, Tools, Type are required.
type ConversationItemNewParamsItemMcpListTools struct {
	// Unique identifier for this MCP list tools operation
	ID string `json:"id,required"`
	// Label identifying the MCP server providing the tools
	ServerLabel string `json:"server_label,required"`
	// List of available tools provided by the MCP server
	Tools []ConversationItemNewParamsItemMcpListToolsTool `json:"tools,omitzero,required"`
	// Tool call type identifier, always "mcp_list_tools"
	//
	// This field can be elided, and will marshal its zero value as "mcp_list_tools".
	Type constant.McpListTools `json:"type,required"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpListTools) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpListTools
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition returned by MCP list tools operation.
//
// The properties InputSchema, Name are required.
type ConversationItemNewParamsItemMcpListToolsTool struct {
	// JSON schema defining the tool's input parameters
	InputSchema map[string]ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion `json:"input_schema,omitzero,required"`
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Description of what the tool does
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ConversationItemNewParamsItemMcpListToolsTool) MarshalJSON() (data []byte, err error) {
	type shadow ConversationItemNewParamsItemMcpListToolsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationItemNewParamsItemMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationItemNewParamsItemMcpListToolsToolInputSchemaUnion) asAny() any {
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

type ConversationItemListParams struct {
	// An item ID to list items after, used in pagination.
	After any `query:"after,omitzero,required" json:"-"`
	// Specify additional output data to include in the response.
	Include []string `query:"include,omitzero,required" json:"-"`
	// A limit on the number of objects to be returned (1-100, default 20).
	Limit any `query:"limit,omitzero,required" json:"-"`
	// The order to return items in (asc or desc, default desc).
	Order ConversationItemListParamsOrder `query:"order,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationItemListParams]'s query parameters as
// `url.Values`.
func (r ConversationItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order to return items in (asc or desc, default desc).
type ConversationItemListParamsOrder string

const (
	ConversationItemListParamsOrderAsc  ConversationItemListParamsOrder = "asc"
	ConversationItemListParamsOrderDesc ConversationItemListParamsOrder = "desc"
)

type ConversationItemGetParams struct {
	ConversationID string `path:"conversation_id,required" json:"-"`
	paramObj
}
