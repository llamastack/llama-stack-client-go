// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Details of a safety violation detected by content moderation.
type SafetyViolation struct {
	// Severity level of a safety violation.
	//
	// Any of "info", "warn", "error".
	ViolationLevel SafetyViolationViolationLevel `json:"violation_level,required"`
	Metadata       map[string]any                `json:"metadata"`
	UserMessage    string                        `json:"user_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ViolationLevel respjson.Field
		Metadata       respjson.Field
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

// Severity level of a safety violation.
type SafetyViolationViolationLevel string

const (
	SafetyViolationViolationLevelInfo  SafetyViolationViolationLevel = "info"
	SafetyViolationViolationLevelWarn  SafetyViolationViolationLevel = "warn"
	SafetyViolationViolationLevelError SafetyViolationViolationLevel = "error"
)

// Sampling parameters.
type SamplingParams struct {
	MaxTokens         param.Opt[int64]   `json:"max_tokens,omitzero"`
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	Stop              []string           `json:"stop,omitzero"`
	// Greedy sampling strategy that selects the highest probability token at each
	// step.
	Strategy SamplingParamsStrategyUnion `json:"strategy,omitzero"`
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

// Greedy sampling strategy that selects the highest probability token at each
// step.
type SamplingParamsStrategyGreedy struct {
	// Any of "greedy".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SamplingParamsStrategyGreedy) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyGreedy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyGreedy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SamplingParamsStrategyGreedy](
		"type", "greedy",
	)
}

// Top-p (nucleus) sampling strategy that samples from the smallest set of tokens
// with cumulative probability >= p.
//
// The property Temperature is required.
type SamplingParamsStrategyTopP struct {
	Temperature param.Opt[float64] `json:"temperature,omitzero,required"`
	TopP        param.Opt[float64] `json:"top_p,omitzero"`
	// Any of "top_p".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SamplingParamsStrategyTopP) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SamplingParamsStrategyTopP](
		"type", "top_p",
	)
}

// Top-k sampling strategy that restricts sampling to the k most likely tokens.
//
// The property TopK is required.
type SamplingParamsStrategyTopK struct {
	TopK int64 `json:"top_k,required"`
	// Any of "top_k".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SamplingParamsStrategyTopK) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopK
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopK) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SamplingParamsStrategyTopK](
		"type", "top_k",
	)
}

// A scoring result for a single row.
type ScoringResult struct {
	AggregatedResults map[string]any   `json:"aggregated_results,required"`
	ScoreRows         []map[string]any `json:"score_rows,required"`
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

// A system message providing instructions or context to the model.
//
// The property Content is required.
type SystemMessageParam struct {
	// A image content item
	Content SystemMessageContentUnionParam `json:"content,omitzero,required"`
	// Any of "system".
	Role SystemMessageRole `json:"role,omitzero"`
	paramObj
}

func (r SystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SystemMessageContentUnionParam struct {
	OfString                                   param.Opt[string]                                                            `json:",omitzero,inline"`
	OfImageContentItemInput                    *SystemMessageContentImageContentItemInputParam                              `json:",omitzero,inline"`
	OfTextContentItem                          *SystemMessageContentTextContentItemParam                                    `json:",omitzero,inline"`
	OfListImageContentItemInputTextContentItem []SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u SystemMessageContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfImageContentItemInput, u.OfTextContentItem, u.OfListImageContentItemInputTextContentItem)
}
func (u *SystemMessageContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SystemMessageContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItemInput) {
		return u.OfImageContentItemInput
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfListImageContentItemInputTextContentItem) {
		return &u.OfListImageContentItemInputTextContentItem
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemMessageContentUnionParam) GetImage() *SystemMessageContentImageContentItemInputImageParam {
	if vt := u.OfImageContentItemInput; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemMessageContentUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemMessageContentUnionParam) GetType() *string {
	if vt := u.OfImageContentItemInput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The property Image is required.
type SystemMessageContentImageContentItemInputParam struct {
	// A URL or a base64 encoded string
	Image SystemMessageContentImageContentItemInputImageParam `json:"image,omitzero,required"`
	// Any of "image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SystemMessageContentImageContentItemInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentImageContentItemInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentImageContentItemInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SystemMessageContentImageContentItemInputParam](
		"type", "image",
	)
}

// A URL or a base64 encoded string
type SystemMessageContentImageContentItemInputImageParam struct {
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL reference to external content.
	URL SystemMessageContentImageContentItemInputImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r SystemMessageContentImageContentItemInputImageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentImageContentItemInputImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentImageContentItemInputImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type SystemMessageContentImageContentItemInputImageURLParam struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r SystemMessageContentImageContentItemInputImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentImageContentItemInputImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentImageContentItemInputImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The property Text is required.
type SystemMessageContentTextContentItemParam struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SystemMessageContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SystemMessageContentTextContentItemParam](
		"type", "text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam struct {
	OfImage *SystemMessageContentListImageContentItemInputTextContentItemItemImageParam `json:",omitzero,inline"`
	OfText  *SystemMessageContentListImageContentItemInputTextContentItemItemTextParam  `json:",omitzero,inline"`
	paramUnion
}

func (u SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfText)
}
func (u *SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam) GetImage() *SystemMessageContentListImageContentItemInputTextContentItemItemImageImageParam {
	if vt := u.OfImage; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam) GetType() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SystemMessageContentListImageContentItemInputTextContentItemItemUnionParam](
		"type",
		apijson.Discriminator[SystemMessageContentListImageContentItemInputTextContentItemItemImageParam]("image"),
		apijson.Discriminator[SystemMessageContentListImageContentItemInputTextContentItemItemTextParam]("text"),
	)
}

// A image content item
//
// The property Image is required.
type SystemMessageContentListImageContentItemInputTextContentItemItemImageParam struct {
	// A URL or a base64 encoded string
	Image SystemMessageContentListImageContentItemInputTextContentItemItemImageImageParam `json:"image,omitzero,required"`
	// Any of "image".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SystemMessageContentListImageContentItemInputTextContentItemItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentListImageContentItemInputTextContentItemItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentListImageContentItemInputTextContentItemItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SystemMessageContentListImageContentItemInputTextContentItemItemImageParam](
		"type", "image",
	)
}

// A URL or a base64 encoded string
type SystemMessageContentListImageContentItemInputTextContentItemItemImageImageParam struct {
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL reference to external content.
	URL SystemMessageContentListImageContentItemInputTextContentItemItemImageImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r SystemMessageContentListImageContentItemInputTextContentItemItemImageImageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentListImageContentItemInputTextContentItemItemImageImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentListImageContentItemInputTextContentItemItemImageImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type SystemMessageContentListImageContentItemInputTextContentItemItemImageImageURLParam struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r SystemMessageContentListImageContentItemInputTextContentItemItemImageImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentListImageContentItemInputTextContentItemItemImageImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentListImageContentItemInputTextContentItemItemImageImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The property Text is required.
type SystemMessageContentListImageContentItemInputTextContentItemItemTextParam struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SystemMessageContentListImageContentItemInputTextContentItemItemTextParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageContentListImageContentItemInputTextContentItemItemTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageContentListImageContentItemInputTextContentItemItemTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SystemMessageContentListImageContentItemInputTextContentItemItemTextParam](
		"type", "text",
	)
}

type SystemMessageRole string

const (
	SystemMessageRoleSystem SystemMessageRole = "system"
)
