// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"encoding/json"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

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

// Sampling parameters.
//
// The property Strategy is required.
type SamplingParams struct {
	// The sampling strategy.
	Strategy SamplingParamsStrategyUnion `json:"strategy,omitzero,required"`
	// The maximum number of tokens that can be generated in the completion. The token
	// count of your prompt plus max_tokens cannot exceed the model's context length.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop []string `json:"stop,omitzero"`
	paramObj
}

func (r SamplingParams) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SamplingParamsStrategyUnion struct {
	OfGreedy *SamplingParamsStrategyGreedy `json:",omitzero,inline"`
	OfTopP   *SamplingParamsStrategyTopP   `json:",omitzero,inline"`
	OfTopK   *SamplingParamsStrategyTopK   `json:",omitzero,inline"`
	paramUnion
}

func (u SamplingParamsStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGreedy, u.OfTopP, u.OfTopK)
}
func (u *SamplingParamsStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SamplingParamsStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfGreedy) {
		return u.OfGreedy
	} else if !param.IsOmitted(u.OfTopP) {
		return u.OfTopP
	} else if !param.IsOmitted(u.OfTopK) {
		return u.OfTopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTemperature() *float64 {
	if vt := u.OfTopP; vt != nil && vt.Temperature.Valid() {
		return &vt.Temperature.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopP() *float64 {
	if vt := u.OfTopP; vt != nil && vt.TopP.Valid() {
		return &vt.TopP.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopK() *int64 {
	if vt := u.OfTopK; vt != nil {
		return &vt.TopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetType() *string {
	if vt := u.OfGreedy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopP; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopK; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SamplingParamsStrategyUnion](
		"type",
		apijson.Discriminator[SamplingParamsStrategyGreedy]("greedy"),
		apijson.Discriminator[SamplingParamsStrategyTopP]("top_p"),
		apijson.Discriminator[SamplingParamsStrategyTopK]("top_k"),
	)
}

func NewSamplingParamsStrategyGreedy() SamplingParamsStrategyGreedy {
	return SamplingParamsStrategyGreedy{
		Type: "greedy",
	}
}

// Greedy sampling strategy that selects the highest probability token at each
// step.
//
// This struct has a constant value, construct it with
// [NewSamplingParamsStrategyGreedy].
type SamplingParamsStrategyGreedy struct {
	// Must be "greedy" to identify this sampling strategy
	Type constant.Greedy `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyGreedy) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyGreedy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyGreedy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-p (nucleus) sampling strategy that samples from the smallest set of tokens
// with cumulative probability >= p.
//
// The property Type is required.
type SamplingParamsStrategyTopP struct {
	// Controls randomness in sampling. Higher values increase randomness
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Cumulative probability threshold for nucleus sampling. Defaults to 0.95
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Must be "top_p" to identify this sampling strategy
	//
	// This field can be elided, and will marshal its zero value as "top_p".
	Type constant.TopP `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyTopP) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-k sampling strategy that restricts sampling to the k most likely tokens.
//
// The properties TopK, Type are required.
type SamplingParamsStrategyTopK struct {
	// Number of top tokens to consider for sampling. Must be at least 1
	TopK int64 `json:"top_k,required"`
	// Must be "top_k" to identify this sampling strategy
	//
	// This field can be elided, and will marshal its zero value as "top_k".
	Type constant.TopK `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyTopK) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopK
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopK) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
