// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/llamastack/llama-stack-client-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type String24h string // Always "24h"
type Batch string     // Always "batch"
type CreatedAt string // Always "created_at"
type Text string      // Always "text"

func (c String24h) Default() String24h { return "24h" }
func (c Batch) Default() Batch         { return "batch" }
func (c CreatedAt) Default() CreatedAt { return "created_at" }
func (c Text) Default() Text           { return "text" }

func (c String24h) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Batch) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c CreatedAt) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
