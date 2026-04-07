// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/option"
)

// AlphaService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaService] method instead.
type AlphaService struct {
	Options []option.RequestOption
	Admin   AlphaAdminService
	// Llama Stack Inference API for generating completions, chat completions, and
	// embeddings.
	//
	// This API provides the raw interface to the underlying models. Three kinds of
	// models are supported:
	//
	//   - LLM models: these models generate "raw" and "chat" (conversational)
	//     completions.
	//   - Embedding models: these models generate embeddings to be used for semantic
	//     search.
	//   - Rerank models: these models reorder the documents based on their relevance to
	//     a query.
	Inference AlphaInferenceService
}

// NewAlphaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlphaService(opts ...option.RequestOption) (r AlphaService) {
	r = AlphaService{}
	r.Options = opts
	r.Admin = NewAlphaAdminService(opts...)
	r.Inference = NewAlphaInferenceService(opts...)
	return
}
