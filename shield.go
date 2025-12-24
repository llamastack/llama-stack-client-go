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
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// ShieldService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShieldService] method instead.
type ShieldService struct {
	Options []option.RequestOption
}

// NewShieldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewShieldService(opts ...option.RequestOption) (r ShieldService) {
	r = ShieldService{}
	r.Options = opts
	return
}

// Get a shield by its identifier.
func (r *ShieldService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *Shield, err error) {
	opts = slices.Concat(r.Options, opts)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("v1/shields/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all shields.
func (r *ShieldService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Shield, err error) {
	var env ListShieldsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/shields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Unregister a shield.
//
// Deprecated: deprecated
func (r *ShieldService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("v1/shields/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Register a shield.
//
// Deprecated: deprecated
func (r *ShieldService) Register(ctx context.Context, body ShieldRegisterParams, opts ...option.RequestOption) (res *Shield, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/shields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ListShieldsResponse struct {
	Data []Shield `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListShieldsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListShieldsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A safety shield resource that can be used to check content.
type Shield struct {
	// Unique identifier for this resource in llama stack
	Identifier string `json:"identifier,required"`
	// ID of the provider that owns this resource
	ProviderID string         `json:"provider_id,required"`
	Params     map[string]any `json:"params,nullable"`
	// Unique identifier for this resource in the provider
	ProviderResourceID string `json:"provider_resource_id,nullable"`
	// Any of "shield".
	Type ShieldType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Params             respjson.Field
		ProviderResourceID respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Shield) RawJSON() string { return r.JSON.raw }
func (r *Shield) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldType string

const (
	ShieldTypeShield ShieldType = "shield"
)

type ShieldRegisterParams struct {
	ShieldID         string            `json:"shield_id,required"`
	ProviderID       param.Opt[string] `json:"provider_id,omitzero"`
	ProviderShieldID param.Opt[string] `json:"provider_shield_id,omitzero"`
	Params           map[string]any    `json:"params,omitzero"`
	paramObj
}

func (r ShieldRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow ShieldRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShieldRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
