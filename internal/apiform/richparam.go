// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.

package apiform

import (
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"mime/multipart"
	"reflect"
)

func (e *encoder) newRichFieldTypeEncoder(t reflect.Type) encoderFunc {
	f, _ := t.FieldByName("Value")
	enc := e.newPrimitiveTypeEncoder(f.Type)
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		if opt, ok := value.Interface().(param.Optional); ok && opt.Valid() {
			return enc(key, value.FieldByIndex(f.Index), writer)
		} else if ok && param.IsNull(opt) {
			return writer.WriteField(key, "null")
		}
		return nil
	}
}
