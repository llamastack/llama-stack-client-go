// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.

package apiquery

import (
	"net/url"
	"reflect"
	"time"
)

func MarshalWithSettings(value any, settings QuerySettings) (url.Values, error) {
	e := encoder{time.RFC3339, true, settings}
	kv := url.Values{}
	val := reflect.ValueOf(value)
	if !val.IsValid() {
		return nil, nil
	}
	typ := val.Type()

	pairs, err := e.typeEncoder(typ)("", val)
	if err != nil {
		return nil, err
	}
	for _, pair := range pairs {
		kv.Add(pair.key, pair.value)
	}
	return kv, nil
}

func Marshal(value any) (url.Values, error) {
	return MarshalWithSettings(value, QuerySettings{})
}

type Queryer interface {
	URLQuery() (url.Values, error)
}

type QuerySettings struct {
	NestedFormat NestedQueryFormat
	ArrayFormat  ArrayQueryFormat
}

type NestedQueryFormat int

const (
	NestedQueryFormatBrackets NestedQueryFormat = iota
	NestedQueryFormatDots
)

type ArrayQueryFormat int

const (
	ArrayQueryFormatComma ArrayQueryFormat = iota
	ArrayQueryFormatRepeat
	ArrayQueryFormatIndices
	ArrayQueryFormatBrackets
)
