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

// ToolRuntimeRagToolService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolRuntimeRagToolService] method instead.
type ToolRuntimeRagToolService struct {
	Options []option.RequestOption
}

// NewToolRuntimeRagToolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewToolRuntimeRagToolService(opts ...option.RequestOption) (r ToolRuntimeRagToolService) {
	r = ToolRuntimeRagToolService{}
	r.Options = opts
	return
}
