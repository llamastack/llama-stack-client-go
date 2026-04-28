// Copyright (c) The OGX Contributors.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ogxclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ogx-ai/ogx-client-go/internal/apijson"
	"github.com/ogx-ai/ogx-client-go/internal/requestconfig"
	"github.com/ogx-ai/ogx-client-go/option"
	"github.com/ogx-ai/ogx-client-go/packages/respjson"
)

// Protocol for prompt management operations.
//
// PromptVersionService contains methods and other services that help with
// interacting with the ogx-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptVersionService] method instead.
type PromptVersionService struct {
	Options []option.RequestOption
}

// NewPromptVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPromptVersionService(opts ...option.RequestOption) (r PromptVersionService) {
	r = PromptVersionService{}
	r.Options = opts
	return
}

// List all versions of a specific prompt.
func (r *PromptVersionService) List(ctx context.Context, promptID string, opts ...option.RequestOption) (res *[]Prompt, err error) {
	var env ListPromptsResponse
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/prompts/%s/versions", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Response model to list prompts.
type ListPromptsResponse struct {
	Data []Prompt `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListPromptsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListPromptsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
