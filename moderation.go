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

// Classifies if text inputs are potentially harmful. OpenAI-compatible endpoint.
func (r *ModerationService) New(ctx context.Context, body ModerationNewParams, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/moderations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A moderation object containing the results of content classification.
type CreateResponse struct {
	// The unique identifier for the moderation request
	ID string `json:"id,required"`
	// The model used to generate the moderation results
	Model string `json:"model,required"`
	// A list of moderation result objects
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

// A moderation result object containing flagged status and category information.
type CreateResponseResult struct {
	// Whether any of the below categories are flagged
	Flagged bool `json:"flagged,required"`
	// A dictionary of the categories, and whether they are flagged or not
	Categories map[string]bool `json:"categories,nullable"`
	// A dictionary of the categories along with the input type(s) that the score
	// applies to
	CategoryAppliedInputTypes map[string][]string `json:"category_applied_input_types,nullable"`
	// A dictionary of the categories along with their scores as predicted by model
	CategoryScores map[string]float64 `json:"category_scores,nullable"`
	// Additional metadata about the moderation
	Metadata map[string]any `json:"metadata"`
	// A message to convey to the user about the moderation result
	UserMessage string `json:"user_message,nullable"`
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
	// Input (or inputs) to classify. Can be a single string or an array of strings.
	Input ModerationNewParamsInputUnion `json:"input,omitzero,required"`
	// The content moderation model to use. If not specified, the default shield will
	// be used.
	Model param.Opt[string] `json:"model,omitzero"`
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
