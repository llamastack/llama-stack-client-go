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

	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
)

// ModelOpenAIService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelOpenAIService] method instead.
type ModelOpenAIService struct {
	Options []option.RequestOption
}

// NewModelOpenAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelOpenAIService(opts ...option.RequestOption) (r ModelOpenAIService) {
	r = ModelOpenAIService{}
	r.Options = opts
	return
}

// List models using the OpenAI API.
func (r *ModelOpenAIService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Model, err error) {
	var env ListModelsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}
