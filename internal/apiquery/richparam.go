// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.

package apiquery

import (
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"reflect"
)

func (e *encoder) newRichFieldTypeEncoder(t reflect.Type) encoderFunc {
	f, _ := t.FieldByName("Value")
	enc := e.typeEncoder(f.Type)
	return func(key string, value reflect.Value) ([]Pair, error) {
		if opt, ok := value.Interface().(param.Optional); ok && opt.Valid() {
			return enc(key, value.FieldByIndex(f.Index))
		} else if ok && param.IsNull(opt) {
			return []Pair{{key, "null"}}, nil
		}
		return nil, nil
	}
}
