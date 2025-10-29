// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"encoding/json"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A message containing the model's (assistant) response in a chat conversation.
//
// The properties Content, Role, StopReason are required.
type CompletionMessageParam struct {
	// The content of the model's response
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Reason why the model stopped generating. Options are: -
	// `StopReason.end_of_turn`: The model finished generating the entire response. -
	// `StopReason.end_of_message`: The model finished generating but generated a
	// partial response -- usually, a tool call. The user may call the tool and
	// continue the conversation with the tool's response. -
	// `StopReason.out_of_tokens`: The model ran out of token budget.
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionMessageStopReason `json:"stop_reason,omitzero,required"`
	// List of tool calls. Each tool call is a ToolCall object.
	ToolCalls []ToolCallParam `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r CompletionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reason why the model stopped generating. Options are: -
// `StopReason.end_of_turn`: The model finished generating the entire response. -
// `StopReason.end_of_message`: The model finished generating but generated a
// partial response -- usually, a tool call. The user may call the tool and
// continue the conversation with the tool's response. -
// `StopReason.out_of_tokens`: The model ran out of token budget.
type CompletionMessageStopReason string

const (
	CompletionMessageStopReasonEndOfTurn    CompletionMessageStopReason = "end_of_turn"
	CompletionMessageStopReasonEndOfMessage CompletionMessageStopReason = "end_of_message"
	CompletionMessageStopReasonOutOfTokens  CompletionMessageStopReason = "out_of_tokens"
)

// A document to be used for document ingestion in the RAG Tool.
//
// The properties Content, DocumentID, Metadata are required.
type DocumentParam struct {
	// The content of the document.
	Content DocumentContentUnionParam `json:"content,omitzero,required"`
	// The unique identifier for the document.
	DocumentID string `json:"document_id,required"`
	// Additional metadata for the document.
	Metadata map[string]DocumentMetadataUnionParam `json:"metadata,omitzero,required"`
	// The MIME type of the document.
	MimeType param.Opt[string] `json:"mime_type,omitzero"`
	paramObj
}

func (r DocumentParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentContentUnionParam struct {
	OfString                      param.Opt[string]                     `json:",omitzero,inline"`
	OfImageContentItem            *DocumentContentImageContentItemParam `json:",omitzero,inline"`
	OfTextContentItem             *DocumentContentTextContentItemParam  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam    `json:",omitzero,inline"`
	OfURL                         *DocumentContentURLParam              `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfImageContentItem,
		u.OfTextContentItem,
		u.OfInterleavedContentItemArray,
		u.OfURL)
}
func (u *DocumentContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	} else if !param.IsOmitted(u.OfURL) {
		return u.OfURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetImage() *DocumentContentImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetUri() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The properties Image, Type are required.
type DocumentContentImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image DocumentContentImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r DocumentContentImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type DocumentContentImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL DocumentContentImageContentItemImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r DocumentContentImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type DocumentContentImageContentItemImageURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r DocumentContentImageContentItemImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentImageContentItemImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentImageContentItemImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type DocumentContentTextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r DocumentContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type DocumentContentURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r DocumentContentURLParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentMetadataUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DocumentMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentMetadataUnionParam) asAny() any {
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

// InterleavedContentUnion contains all possible properties and values from
// [string], [InterleavedContentImageContentItem],
// [InterleavedContentTextContentItem], [[]InterleavedContentItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInterleavedContentItemArray]
type InterleavedContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]InterleavedContentItemUnion]
	// instead of an object.
	OfInterleavedContentItemArray []InterleavedContentItemUnion `json:",inline"`
	// This field is from variant [InterleavedContentImageContentItem].
	Image InterleavedContentImageContentItemImage `json:"image"`
	Type  string                                  `json:"type"`
	// This field is from variant [InterleavedContentTextContentItem].
	Text string `json:"text"`
	JSON struct {
		OfString                      respjson.Field
		OfInterleavedContentItemArray respjson.Field
		Image                         respjson.Field
		Type                          respjson.Field
		Text                          respjson.Field
		raw                           string
	} `json:"-"`
}

func (u InterleavedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsImageContentItem() (v InterleavedContentImageContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsTextContentItem() (v InterleavedContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsInterleavedContentItemArray() (v []InterleavedContentItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InterleavedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *InterleavedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InterleavedContentUnion to a InterleavedContentUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InterleavedContentUnionParam.Overrides()
func (r InterleavedContentUnion) ToParam() InterleavedContentUnionParam {
	return param.Override[InterleavedContentUnionParam](json.RawMessage(r.RawJSON()))
}

// A image content item
type InterleavedContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentImageContentItemImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	Type constant.Image `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentImageContentItemImageURL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
type InterleavedContentImageContentItemImageURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItemImageURL) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InterleavedContentParamOfImageContentItem(image InterleavedContentImageContentItemImageParam) InterleavedContentUnionParam {
	var variant InterleavedContentImageContentItemParam
	variant.Image = image
	return InterleavedContentUnionParam{OfImageContentItem: &variant}
}

func InterleavedContentParamOfTextContentItem(text string) InterleavedContentUnionParam {
	var variant InterleavedContentTextContentItemParam
	variant.Text = text
	return InterleavedContentUnionParam{OfTextContentItem: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentUnionParam struct {
	OfString                      param.Opt[string]                        `json:",omitzero,inline"`
	OfImageContentItem            *InterleavedContentImageContentItemParam `json:",omitzero,inline"`
	OfTextContentItem             *InterleavedContentTextContentItemParam  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam       `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfImageContentItem, u.OfTextContentItem, u.OfInterleavedContentItemArray)
}
func (u *InterleavedContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetImage() *InterleavedContentImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r InterleavedContentImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentImageContentItemImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type InterleavedContentImageContentItemImageURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r InterleavedContentImageContentItemImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentTextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InterleavedContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InterleavedContentItemUnion contains all possible properties and values from
// [InterleavedContentItemImage], [InterleavedContentItemText].
//
// Use the [InterleavedContentItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InterleavedContentItemUnion struct {
	// This field is from variant [InterleavedContentItemImage].
	Image InterleavedContentItemImageImage `json:"image"`
	// Any of "image", "text".
	Type string `json:"type"`
	// This field is from variant [InterleavedContentItemText].
	Text string `json:"text"`
	JSON struct {
		Image respjson.Field
		Type  respjson.Field
		Text  respjson.Field
		raw   string
	} `json:"-"`
}

// anyInterleavedContentItem is implemented by each variant of
// [InterleavedContentItemUnion] to add type safety for the return type of
// [InterleavedContentItemUnion.AsAny]
type anyInterleavedContentItem interface {
	implInterleavedContentItemUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := InterleavedContentItemUnion.AsAny().(type) {
//	case llamastackclient.InterleavedContentItemImage:
//	case llamastackclient.InterleavedContentItemText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InterleavedContentItemUnion) AsAny() anyInterleavedContentItem {
	switch u.Type {
	case "image":
		return u.AsImage()
	case "text":
		return u.AsText()
	}
	return nil
}

func (u InterleavedContentItemUnion) AsImage() (v InterleavedContentItemImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentItemUnion) AsText() (v InterleavedContentItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InterleavedContentItemUnion) RawJSON() string { return u.JSON.raw }

func (r *InterleavedContentItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InterleavedContentItemUnion to a
// InterleavedContentItemUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InterleavedContentItemUnionParam.Overrides()
func (r InterleavedContentItemUnion) ToParam() InterleavedContentItemUnionParam {
	return param.Override[InterleavedContentItemUnionParam](json.RawMessage(r.RawJSON()))
}

// A image content item
type InterleavedContentItemImage struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	Type constant.Image `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (InterleavedContentItemImage) implInterleavedContentItemUnion() {}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentItemImageImageURL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImageImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
type InterleavedContentItemImageImageURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImageImageURL) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentItemText struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemText) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (InterleavedContentItemText) implInterleavedContentItemUnion() {}

func InterleavedContentItemParamOfImage(image InterleavedContentItemImageImageParam) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemImageParam
	variant.Image = image
	return InterleavedContentItemUnionParam{OfImage: &variant}
}

func InterleavedContentItemParamOfText(text string) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemTextParam
	variant.Text = text
	return InterleavedContentItemUnionParam{OfText: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentItemUnionParam struct {
	OfImage *InterleavedContentItemImageParam `json:",omitzero,inline"`
	OfText  *InterleavedContentItemTextParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfText)
}
func (u *InterleavedContentItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetImage() *InterleavedContentItemImageImageParam {
	if vt := u.OfImage; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetType() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InterleavedContentItemUnionParam](
		"type",
		apijson.Discriminator[InterleavedContentItemImageParam]("image"),
		apijson.Discriminator[InterleavedContentItemTextParam]("text"),
	)
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentItemImageParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r InterleavedContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentItemImageImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentItemImageImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type InterleavedContentItemImageImageURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r InterleavedContentItemImageImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentItemTextParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InterleavedContentItemTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func MessageParamOfUser[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T) MessageUnionParam {
	var user UserMessageParam
	switch v := any(content).(type) {
	case string:
		user.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		user.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		user.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		user.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfUser: &user}
}

func MessageParamOfSystem[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T) MessageUnionParam {
	var system SystemMessageParam
	switch v := any(content).(type) {
	case string:
		system.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		system.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		system.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		system.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfSystem: &system}
}

func MessageParamOfTool[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](callID string, content T) MessageUnionParam {
	var tool ToolResponseMessageParam
	tool.CallID = callID
	switch v := any(content).(type) {
	case string:
		tool.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		tool.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		tool.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		tool.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfTool: &tool}
}

func MessageParamOfAssistant[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T, stopReason CompletionMessageStopReason) MessageUnionParam {
	var assistant CompletionMessageParam
	switch v := any(content).(type) {
	case string:
		assistant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		assistant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		assistant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		assistant.Content.OfInterleavedContentItemArray = v
	}
	assistant.StopReason = stopReason
	return MessageUnionParam{OfAssistant: &assistant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageUnionParam struct {
	OfUser      *UserMessageParam         `json:",omitzero,inline"`
	OfSystem    *SystemMessageParam       `json:",omitzero,inline"`
	OfTool      *ToolResponseMessageParam `json:",omitzero,inline"`
	OfAssistant *CompletionMessageParam   `json:",omitzero,inline"`
	paramUnion
}

func (u MessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser, u.OfSystem, u.OfTool, u.OfAssistant)
}
func (u *MessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetContext() *InterleavedContentUnionParam {
	if vt := u.OfUser; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetStopReason() *string {
	if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.StopReason)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetToolCalls() []ToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u MessageUnionParam) GetContent() *InterleavedContentUnionParam {
	if vt := u.OfUser; vt != nil {
		return &vt.Content
	} else if vt := u.OfSystem; vt != nil {
		return &vt.Content
	} else if vt := u.OfTool; vt != nil {
		return &vt.Content
	} else if vt := u.OfAssistant; vt != nil {
		return &vt.Content
	}
	return nil
}

func init() {
	apijson.RegisterUnion[MessageUnionParam](
		"role",
		apijson.Discriminator[UserMessageParam]("user"),
		apijson.Discriminator[SystemMessageParam]("system"),
		apijson.Discriminator[ToolResponseMessageParam]("tool"),
		apijson.Discriminator[CompletionMessageParam]("assistant"),
	)
}

// Configuration for the RAG query generation.
//
// The properties ChunkTemplate, MaxChunks, MaxTokensInContext,
// QueryGeneratorConfig are required.
type QueryConfigParam struct {
	// Template for formatting each retrieved chunk in the context. Available
	// placeholders: {index} (1-based chunk ordinal), {chunk.content} (chunk content
	// string), {metadata} (chunk metadata dict). Default: "Result {index}\nContent:
	// {chunk.content}\nMetadata: {metadata}\n"
	ChunkTemplate string `json:"chunk_template,required"`
	// Maximum number of chunks to retrieve.
	MaxChunks int64 `json:"max_chunks,required"`
	// Maximum number of tokens in the context.
	MaxTokensInContext int64 `json:"max_tokens_in_context,required"`
	// Configuration for the query generator.
	QueryGeneratorConfig QueryConfigQueryGeneratorConfigUnionParam `json:"query_generator_config,omitzero,required"`
	// Search mode for retrieval—either "vector", "keyword", or "hybrid". Default
	// "vector".
	//
	// Any of "vector", "keyword", "hybrid".
	Mode QueryConfigMode `json:"mode,omitzero"`
	// Configuration for the ranker to use in hybrid search. Defaults to RRF ranker.
	Ranker QueryConfigRankerUnionParam `json:"ranker,omitzero"`
	paramObj
}

func (r QueryConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryConfigQueryGeneratorConfigUnionParam struct {
	OfDefault *QueryConfigQueryGeneratorConfigDefaultParam `json:",omitzero,inline"`
	OfLlm     *QueryConfigQueryGeneratorConfigLlmParam     `json:",omitzero,inline"`
	paramUnion
}

func (u QueryConfigQueryGeneratorConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDefault, u.OfLlm)
}
func (u *QueryConfigQueryGeneratorConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryConfigQueryGeneratorConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDefault) {
		return u.OfDefault
	} else if !param.IsOmitted(u.OfLlm) {
		return u.OfLlm
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigQueryGeneratorConfigUnionParam) GetSeparator() *string {
	if vt := u.OfDefault; vt != nil {
		return &vt.Separator
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigQueryGeneratorConfigUnionParam) GetModel() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Model
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigQueryGeneratorConfigUnionParam) GetTemplate() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigQueryGeneratorConfigUnionParam) GetType() *string {
	if vt := u.OfDefault; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLlm; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[QueryConfigQueryGeneratorConfigUnionParam](
		"type",
		apijson.Discriminator[QueryConfigQueryGeneratorConfigDefaultParam]("default"),
		apijson.Discriminator[QueryConfigQueryGeneratorConfigLlmParam]("llm"),
	)
}

// Configuration for the default RAG query generator.
//
// The properties Separator, Type are required.
type QueryConfigQueryGeneratorConfigDefaultParam struct {
	// String separator used to join query terms
	Separator string `json:"separator,required"`
	// Type of query generator, always 'default'
	//
	// This field can be elided, and will marshal its zero value as "default".
	Type constant.Default `json:"type,required"`
	paramObj
}

func (r QueryConfigQueryGeneratorConfigDefaultParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigQueryGeneratorConfigDefaultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigQueryGeneratorConfigDefaultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the LLM-based RAG query generator.
//
// The properties Model, Template, Type are required.
type QueryConfigQueryGeneratorConfigLlmParam struct {
	// Name of the language model to use for query generation
	Model string `json:"model,required"`
	// Template string for formatting the query generation prompt
	Template string `json:"template,required"`
	// Type of query generator, always 'llm'
	//
	// This field can be elided, and will marshal its zero value as "llm".
	Type constant.Llm `json:"type,required"`
	paramObj
}

func (r QueryConfigQueryGeneratorConfigLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigQueryGeneratorConfigLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigQueryGeneratorConfigLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search mode for retrieval—either "vector", "keyword", or "hybrid". Default
// "vector".
type QueryConfigMode string

const (
	QueryConfigModeVector  QueryConfigMode = "vector"
	QueryConfigModeKeyword QueryConfigMode = "keyword"
	QueryConfigModeHybrid  QueryConfigMode = "hybrid"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryConfigRankerUnionParam struct {
	OfRrf      *QueryConfigRankerRrfParam      `json:",omitzero,inline"`
	OfWeighted *QueryConfigRankerWeightedParam `json:",omitzero,inline"`
	paramUnion
}

func (u QueryConfigRankerUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRrf, u.OfWeighted)
}
func (u *QueryConfigRankerUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryConfigRankerUnionParam) asAny() any {
	if !param.IsOmitted(u.OfRrf) {
		return u.OfRrf
	} else if !param.IsOmitted(u.OfWeighted) {
		return u.OfWeighted
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigRankerUnionParam) GetImpactFactor() *float64 {
	if vt := u.OfRrf; vt != nil {
		return &vt.ImpactFactor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigRankerUnionParam) GetAlpha() *float64 {
	if vt := u.OfWeighted; vt != nil {
		return &vt.Alpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigRankerUnionParam) GetType() *string {
	if vt := u.OfRrf; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeighted; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[QueryConfigRankerUnionParam](
		"type",
		apijson.Discriminator[QueryConfigRankerRrfParam]("rrf"),
		apijson.Discriminator[QueryConfigRankerWeightedParam]("weighted"),
	)
}

// Reciprocal Rank Fusion (RRF) ranker configuration.
//
// The properties ImpactFactor, Type are required.
type QueryConfigRankerRrfParam struct {
	// The impact factor for RRF scoring. Higher values give more weight to
	// higher-ranked results. Must be greater than 0
	ImpactFactor float64 `json:"impact_factor,required"`
	// The type of ranker, always "rrf"
	//
	// This field can be elided, and will marshal its zero value as "rrf".
	Type constant.Rrf `json:"type,required"`
	paramObj
}

func (r QueryConfigRankerRrfParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigRankerRrfParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigRankerRrfParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weighted ranker configuration that combines vector and keyword scores.
//
// The properties Alpha, Type are required.
type QueryConfigRankerWeightedParam struct {
	// Weight factor between 0 and 1. 0 means only use keyword scores, 1 means only use
	// vector scores, values in between blend both scores.
	Alpha float64 `json:"alpha,required"`
	// The type of ranker, always "weighted"
	//
	// This field can be elided, and will marshal its zero value as "weighted".
	Type constant.Weighted `json:"type,required"`
	paramObj
}

func (r QueryConfigRankerWeightedParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigRankerWeightedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigRankerWeightedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a RAG query containing retrieved content and metadata.
type QueryResult struct {
	// Additional metadata about the query result
	Metadata map[string]QueryResultMetadataUnion `json:"metadata,required"`
	// (Optional) The retrieved content from the query
	Content InterleavedContentUnion `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryResult) RawJSON() string { return r.JSON.raw }
func (r *QueryResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QueryResultMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type QueryResultMetadataUnion struct {
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

func (u QueryResultMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryResultMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryResultMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryResultMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QueryResultMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *QueryResultMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of a safety violation detected by content moderation.
type SafetyViolation struct {
	// Additional metadata including specific violation codes for debugging and
	// telemetry
	Metadata map[string]SafetyViolationMetadataUnion `json:"metadata,required"`
	// Severity level of the violation
	//
	// Any of "info", "warn", "error".
	ViolationLevel SafetyViolationViolationLevel `json:"violation_level,required"`
	// (Optional) Message to convey to the user about the violation
	UserMessage string `json:"user_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata       respjson.Field
		ViolationLevel respjson.Field
		UserMessage    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafetyViolation) RawJSON() string { return r.JSON.raw }
func (r *SafetyViolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SafetyViolationMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SafetyViolationMetadataUnion struct {
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

func (u SafetyViolationMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SafetyViolationMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *SafetyViolationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Severity level of the violation
type SafetyViolationViolationLevel string

const (
	SafetyViolationViolationLevelInfo  SafetyViolationViolationLevel = "info"
	SafetyViolationViolationLevelWarn  SafetyViolationViolationLevel = "warn"
	SafetyViolationViolationLevelError SafetyViolationViolationLevel = "error"
)

// A scoring result for a single row.
type ScoringResult struct {
	// Map of metric name to aggregated value
	AggregatedResults map[string]ScoringResultAggregatedResultUnion `json:"aggregated_results,required"`
	// The scoring result for each row. Each row is a map of column name to value.
	ScoreRows []map[string]ScoringResultScoreRowUnion `json:"score_rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregatedResults respjson.Field
		ScoreRows         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringResult) RawJSON() string { return r.JSON.raw }
func (r *ScoringResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringResultAggregatedResultUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringResultAggregatedResultUnion struct {
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

func (u ScoringResultAggregatedResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringResultAggregatedResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringResultAggregatedResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringResultScoreRowUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringResultScoreRowUnion struct {
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

func (u ScoringResultScoreRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringResultScoreRowUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringResultScoreRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type SystemMessageParam struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r SystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Arguments, CallID, ToolName are required.
type ToolCallParam struct {
	Arguments string           `json:"arguments,required"`
	CallID    string           `json:"call_id,required"`
	ToolName  ToolCallToolName `json:"tool_name,omitzero,required"`
	paramObj
}

func (r ToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolCallToolName string

const (
	ToolCallToolNameBraveSearch     ToolCallToolName = "brave_search"
	ToolCallToolNameWolframAlpha    ToolCallToolName = "wolfram_alpha"
	ToolCallToolNamePhotogen        ToolCallToolName = "photogen"
	ToolCallToolNameCodeInterpreter ToolCallToolName = "code_interpreter"
)

// A message representing the result of a tool invocation.
//
// The properties CallID, Content, Role are required.
type ToolResponseMessageParam struct {
	// Unique identifier for the tool call this response is for
	CallID string `json:"call_id,required"`
	// The response content from the tool
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r ToolResponseMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolResponseMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolResponseMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in a chat conversation.
//
// The properties Content, Role are required.
type UserMessageParam struct {
	// The content of the message, which can include text and other media
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// (Optional) This field is used internally by Llama Stack to pass RAG context.
	// This field may be removed in the API in the future.
	Context InterleavedContentUnionParam `json:"context,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r UserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow UserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
