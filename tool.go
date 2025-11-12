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
)

// ToolService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r ToolService) {
	r = ToolService{}
	r.Options = opts
	return
}

// List tools with optional tool group.
func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *[]ToolDef, err error) {
	var env ToolListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a tool by its name.
func (r *ToolService) Get(ctx context.Context, toolName string, opts ...option.RequestOption) (res *ToolDef, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolName == "" {
		err = errors.New("missing required tool_name parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s", toolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ToolListParams struct {
	// The ID of the tool group to list tools for.
	ToolgroupID param.Opt[string] `query:"toolgroup_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolListParams]'s query parameters as `url.Values`.
func (r ToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response containing a list of tool definitions.
type ToolListResponseEnvelope struct {
	// List of tool definitions
	Data []ToolDef `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ToolListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
