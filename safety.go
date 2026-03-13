// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// OpenAI-compatible Moderations API.
//
// SafetyService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSafetyService] method instead.
type SafetyService struct {
	Options []option.RequestOption
}

// NewSafetyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSafetyService(opts ...option.RequestOption) (r SafetyService) {
	r = SafetyService{}
	r.Options = opts
	return
}

// Run a safety shield on messages to check for policy violations.
func (r *SafetyService) RunShield(ctx context.Context, body SafetyRunShieldParams, opts ...option.RequestOption) (res *RunShieldResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/safety/run-shield"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response from running a safety shield.
type RunShieldResponse struct {
	// Details of a safety violation detected by content moderation.
	Violation SafetyViolation `json:"violation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Violation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunShieldResponse) RawJSON() string { return r.JSON.raw }
func (r *RunShieldResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SafetyRunShieldParams struct {
	// The messages to run the shield on
	Messages []SafetyRunShieldParamsMessageUnion `json:"messages,omitzero" api:"required"`
	// The identifier of the shield to run
	ShieldID string `json:"shield_id" api:"required"`
	paramObj
}

func (r SafetyRunShieldParams) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageUnion struct {
	OfUser      *SafetyRunShieldParamsMessageUser      `json:",omitzero,inline"`
	OfSystem    *SafetyRunShieldParamsMessageSystem    `json:",omitzero,inline"`
	OfAssistant *SafetyRunShieldParamsMessageAssistant `json:",omitzero,inline"`
	OfTool      *SafetyRunShieldParamsMessageTool      `json:",omitzero,inline"`
	OfDeveloper *SafetyRunShieldParamsMessageDeveloper `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser,
		u.OfSystem,
		u.OfAssistant,
		u.OfTool,
		u.OfDeveloper)
}
func (u *SafetyRunShieldParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfDeveloper) {
		return u.OfDeveloper
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetToolCalls() []SafetyRunShieldParamsMessageAssistantToolCall {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetName() *string {
	if vt := u.OfUser; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSystem; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAssistant; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfDeveloper; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfDeveloper; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SafetyRunShieldParamsMessageUnion) GetContent() (res safetyRunShieldParamsMessageUnionContent) {
	if vt := u.OfUser; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfAssistant; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfTool; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfDeveloper; vt != nil {
		res.any = vt.Content.asAny()
	}
	return
}

// Can have the runtime types [*string],
// [_[]SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion],
// [_[]SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem],
// [_[]SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem],
// [_[]SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem],
// [\*[]SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem]
type safetyRunShieldParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u safetyRunShieldParamsMessageUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[SafetyRunShieldParamsMessageUnion](
		"role",
		apijson.Discriminator[SafetyRunShieldParamsMessageUser]("user"),
		apijson.Discriminator[SafetyRunShieldParamsMessageSystem]("system"),
		apijson.Discriminator[SafetyRunShieldParamsMessageAssistant]("assistant"),
		apijson.Discriminator[SafetyRunShieldParamsMessageTool]("tool"),
		apijson.Discriminator[SafetyRunShieldParamsMessageDeveloper]("developer"),
	)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The property Content is required.
type SafetyRunShieldParamsMessageUser struct {
	// The content of the message, which can include text and other media.
	Content SafetyRunShieldParamsMessageUserContentUnion `json:"content,omitzero" api:"required"`
	// The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be 'user' to identify this as a user message.
	//
	// Any of "user".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUser) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageUser](
		"role", "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageUserContentUnion struct {
	OfString                                                                                     param.Opt[string]                                                                                                                                 `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile []SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile)
}
func (u *SafetyRunShieldParamsMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageUserContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile) {
		return &u.OfListOpenAIChatCompletionContentPartTextOpenAIChatCompletionContentPartImageParamOpenAIFile
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion struct {
	OfText     *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText     `json:",omitzero,inline"`
	OfImageURL *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL `json:",omitzero,inline"`
	OfFile     *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetImageURL() *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetFile() *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemUnion](
		"type",
		apijson.Discriminator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText]("text"),
		apijson.Discriminator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL]("image_url"),
		apijson.Discriminator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile]("file"),
	)
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText struct {
	// The text content of the message.
	Text string `json:"text" api:"required"`
	// Must be 'text' to identify this as text content.
	//
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemText](
		"type", "text",
	)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The property ImageURL is required.
type SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL struct {
	// Image URL specification and processing details.
	ImageURL SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL `json:"image_url,omitzero" api:"required"`
	// Must be 'image_url' to identify this as image content.
	//
	// Any of "image_url".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURL](
		"type", "image_url",
	)
}

// Image URL specification and processing details.
//
// The property URL is required.
type SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL struct {
	// URL of the image to include in the message.
	URL string `json:"url" api:"required"`
	// Level of detail for image processing. Can be 'low', 'high', or 'auto'.
	//
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemImageURLImageURL](
		"detail", "low", "high", "auto",
	)
}

// The property File is required.
type SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile struct {
	// File specification.
	File SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile `json:"file,omitzero" api:"required"`
	// Must be 'file' to identify this as file content.
	//
	// Any of "file".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFile](
		"type", "file",
	)
}

// File specification.
type SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile struct {
	// Base64-encoded file data.
	FileData param.Opt[string] `json:"file_data,omitzero"`
	// ID of an uploaded file.
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// Name of the file.
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentListOpenAIChatCompletionContentPartTextParamOpenAIChatCompletionContentPartImageParamOpenAIFileItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The property Content is required.
type SafetyRunShieldParamsMessageSystem struct {
	// The content of the 'system prompt'. If multiple system messages are provided,
	// they are concatenated.
	Content SafetyRunShieldParamsMessageSystemContentUnion `json:"content,omitzero" api:"required"`
	// The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be 'system' to identify this as a system message.
	//
	// Any of "system".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageSystem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageSystem](
		"role", "system",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageSystemContentUnion struct {
	OfString                                  param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageSystemContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *SafetyRunShieldParamsMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageSystemContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem struct {
	// The text content of the message.
	Text string `json:"text" api:"required"`
	// Must be 'text' to identify this as text content.
	//
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageSystemContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type SafetyRunShieldParamsMessageAssistant struct {
	// The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response.
	Content SafetyRunShieldParamsMessageAssistantContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []SafetyRunShieldParamsMessageAssistantToolCall `json:"tool_calls,omitzero"`
	// Must be 'assistant' to identify this as the model's response.
	//
	// Any of "assistant".
	Role        string         `json:"role,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistant) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistant
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *SafetyRunShieldParamsMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageAssistant](
		"role", "assistant",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageAssistantContentUnion struct {
	OfString                                  param.Opt[string]                                                                              `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageAssistantContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *SafetyRunShieldParamsMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageAssistantContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem struct {
	// The text content of the message.
	Text string `json:"text" api:"required"`
	// Must be 'text' to identify this as text content.
	//
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageAssistantContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// Tool call specification for OpenAI-compatible chat completion responses.
type SafetyRunShieldParamsMessageAssistantToolCall struct {
	// Unique identifier for the tool call.
	ID param.Opt[string] `json:"id,omitzero"`
	// Index of the tool call in the list.
	Index param.Opt[int64] `json:"index,omitzero"`
	// Function call details for OpenAI-compatible tool calls.
	Function SafetyRunShieldParamsMessageAssistantToolCallFunction `json:"function,omitzero"`
	// Must be 'function' to identify this as a function call.
	//
	// Any of "function".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistantToolCall) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistantToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageAssistantToolCall](
		"type", "function",
	)
}

// Function call details for OpenAI-compatible tool calls.
//
// The properties Arguments, Name are required.
type SafetyRunShieldParamsMessageAssistantToolCallFunction struct {
	// Arguments to pass to the function as a JSON string.
	Arguments string `json:"arguments" api:"required"`
	// Name of the function to call.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistantToolCallFunction) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistantToolCallFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
//
// The properties Content, ToolCallID are required.
type SafetyRunShieldParamsMessageTool struct {
	// The response content from the tool.
	Content SafetyRunShieldParamsMessageToolContentUnion `json:"content,omitzero" api:"required"`
	// Unique identifier for the tool call this response is for.
	ToolCallID string `json:"tool_call_id" api:"required"`
	// Must be 'tool' to identify this as a tool response.
	//
	// Any of "tool".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageTool) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageTool](
		"role", "tool",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageToolContentUnion struct {
	OfString                                  param.Opt[string]                                                                         `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageToolContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *SafetyRunShieldParamsMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageToolContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem struct {
	// The text content of the message.
	Text string `json:"text" api:"required"`
	// Must be 'text' to identify this as text content.
	//
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageToolContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The property Content is required.
type SafetyRunShieldParamsMessageDeveloper struct {
	// The content of the developer message.
	Content SafetyRunShieldParamsMessageDeveloperContentUnion `json:"content,omitzero" api:"required"`
	// The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be 'developer' to identify this as a developer message.
	//
	// Any of "developer".
	Role string `json:"role,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageDeveloper](
		"role", "developer",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageDeveloperContentUnion struct {
	OfString                                  param.Opt[string]                                                                              `json:",omitzero,inline"`
	OfListOpenAIChatCompletionContentPartText []SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageDeveloperContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListOpenAIChatCompletionContentPartText)
}
func (u *SafetyRunShieldParamsMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageDeveloperContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListOpenAIChatCompletionContentPartText) {
		return &u.OfListOpenAIChatCompletionContentPartText
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The property Text is required.
type SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem struct {
	// The text content of the message.
	Text string `json:"text" api:"required"`
	// Must be 'text' to identify this as text content.
	//
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SafetyRunShieldParamsMessageDeveloperContentListOpenAIChatCompletionContentPartTextParamItem](
		"type", "text",
	)
}
