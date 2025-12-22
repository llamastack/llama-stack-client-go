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

// InspectService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInspectService] method instead.
type InspectService struct {
	Options []option.RequestOption
}

// NewInspectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInspectService(opts ...option.RequestOption) (r InspectService) {
	r = InspectService{}
	r.Options = opts
	return
}

// Get the current health status of the service.
func (r *InspectService) Health(ctx context.Context, opts ...option.RequestOption) (res *HealthInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the version of the service.
func (r *InspectService) Version(ctx context.Context, opts ...option.RequestOption) (res *VersionInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
