/*
 *
 * Copyright 2023 cornucopia authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"errors"
	"io"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"go.starlark.net/starlark"
)

// check interface compliance
var _ starlark.HasAttrs = wrapper[jen.Code]{}
var _ coder = wrapper[jen.Code]{}

var errUnhashable = errors.New("unhashable builtin type")

type coder interface {
	code() jen.Code
}

type renderer interface {
	Render(writer io.Writer) error
}

type wrappedType struct {
	name        string
	methods     map[string]*starlark.Builtin
	methodNames []string
}

func makeWrappedType(name string, methods ...*starlark.Builtin) wrappedType {
	size := len(methods)
	methodsMap := make(map[string]*starlark.Builtin, size)
	methodNames := make([]string, 0, size)

	for _, m := range methods {
		mName := m.Name()
		methodsMap[mName] = m
		methodNames = append(methodNames, mName)
	}
	sort.Strings(methodNames)

	return wrappedType{name: name, methods: methodsMap, methodNames: methodNames}
}

type wrapper[T jen.Code] struct {
	inner T
	wType *wrappedType
}

func (w wrapper[T]) String() string {
	casted, ok := w.code().(renderer)
	if ok {
		var buffer strings.Builder
		if err := casted.Render(&buffer); err == nil {
			return buffer.String()
		}
	}
	return ""
}

func (w wrapper[T]) Type() string {
	return w.wType.name
}

func (w wrapper[T]) Freeze() {
	// No op
}

func (w wrapper[T]) Truth() starlark.Bool {
	return starlark.True
}

func (w wrapper[T]) Hash() (uint32, error) {
	return 0, errUnhashable
}

func (w wrapper[T]) Attr(name string) (starlark.Value, error) {
	m, ok := w.wType.methods[name]
	if !ok {
		return nil, nil
	}
	return m.BindReceiver(w), nil
}

func (w wrapper[T]) AttrNames() []string {
	return w.wType.methodNames
}

func (w wrapper[T]) code() jen.Code {
	return w.inner
}

func convertToGoBuiltin(value starlark.Value) any {
	switch casted := value.(type) {
	case starlark.Bool:
		return bool(casted)
	case starlark.Int:
		res, ok := casted.Int64()
		if ok {
			return res
		}
	case starlark.Float:
		return float64(casted)
	case starlark.String:
		return string(casted)
	}
	return nil
}
func convertToGoByte(value starlark.Value) byte {
	casted, ok := value.(starlark.Int)
	if ok {
		res, ok2 := casted.Int64()
		if ok2 {
			return byte(res)
		}
	}
	return 0
}

func convertToGoRune(value starlark.Value) rune {
	switch casted := value.(type) {
	case starlark.Int:
		res, ok := casted.Int64()
		if ok {
			return rune(res)
		}
	case starlark.String:
		for _, r := range string(casted) {
			return r
		}
	}
	return 0
}

func convertToCode(value starlark.Value) jen.Code {
	casted, ok := value.(coder)
	if !ok {
		return jen.Lit(convertToGoBuiltin(value))
	}
	return casted.code()
}

func convertToCodeSlice(args starlark.Tuple) []jen.Code {
	res := make([]jen.Code, 0, len(args))
	for _, arg := range args {
		res = append(res, convertToCode(arg))
	}
	return res
}

func convertToStringSlice(args starlark.Tuple) []string {
	res := make([]string, 0, len(args))
	for _, arg := range args {
		s, ok := starlark.AsString(arg)
		if ok {
			res = append(res, s)
		}
	}
	return res
}

func convertToMapString(items []starlark.Tuple) map[string]string {
	mapString := make(map[string]string, len(items))
	for _, item := range items {
		key, ok := starlark.AsString(item[0])
		value, ok2 := starlark.AsString(item[1])
		if ok && ok2 {
			mapString[key] = value
		}
	}
	return mapString
}
