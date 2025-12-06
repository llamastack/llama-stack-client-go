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

// ModerationService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModerationService] method instead.
type ModerationService struct {
	Options []option.RequestOption
}

// NewModerationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModerationService(opts ...option.RequestOption) (r ModerationService) {
	r = ModerationService{}
	r.Options = opts
	return
}

// Create moderation.
//
// Classifies if text and/or image inputs are potentially harmful.
func (r *ModerationService) New(ctx context.Context, body ModerationNewParams, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/moderations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A moderation object.
type CreateResponse struct {
	ID      string                 `json:"id,required"`
	Model   string                 `json:"model,required"`
	Results []CreateResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Model       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A moderation object.
type CreateResponseResult struct {
	Flagged                   bool                `json:"flagged,required"`
	Categories                map[string]bool     `json:"categories,nullable"`
	CategoryAppliedInputTypes map[string][]string `json:"category_applied_input_types,nullable"`
	CategoryScores            map[string]float64  `json:"category_scores,nullable"`
	Metadata                  map[string]any      `json:"metadata"`
	UserMessage               string              `json:"user_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flagged                   respjson.Field
		Categories                respjson.Field
		CategoryAppliedInputTypes respjson.Field
		CategoryScores            respjson.Field
		Metadata                  respjson.Field
		UserMessage               respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateResponseResult) RawJSON() string { return r.JSON.raw }
func (r *CreateResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModerationNewParams struct {
	Input ModerationNewParamsInputUnion `json:"input,omitzero,required"`
	Model param.Opt[string]             `json:"model,omitzero"`
	paramObj
}

func (r ModerationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ModerationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModerationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModerationNewParamsInputUnion struct {
	OfString     param.Opt[string] `json:",omitzero,inline"`
	OfListString []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ModerationNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfListString)
}
func (u *ModerationNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModerationNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfListString) {
		return &u.OfListString
	}
	return nil
}
