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
	"sort"

	"github.com/dave/jennifer/jen"
	"go.starlark.net/starlark"
)

// check interface compliance
var _ starlark.HasAttrs = wrapper[any]{}

var errUnhashable = errors.New("unhashable builtin type")

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

type wrapper[T any] struct {
	inner T
	wType *wrappedType
}

func (w wrapper[T]) String() string {
	return "todo"
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

func convertToCode(value starlark.Value) jen.Code {
	// TODO
	return nil
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
		if !ok {
			res = append(res, s)
		}
	}
	return res
}
