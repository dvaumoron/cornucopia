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

package glu

import (
	"bytes"
	"io"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"go.starlark.net/starlark"
)

type renderer interface {
	Render(writer io.Writer) error
}

type Coder interface {
	Code() jen.Code
}

type WrappedType struct {
	name        string
	methods     map[string]*starlark.Builtin
	methodNames []string
}

func MakeWrappedType(name string, methods ...*starlark.Builtin) WrappedType {
	size := len(methods)
	methodsMap := make(map[string]*starlark.Builtin, size)
	methodNames := make([]string, 0, size)

	for _, m := range methods {
		mName := m.Name()
		methodsMap[mName] = m
		methodNames = append(methodNames, mName)
	}
	sort.Strings(methodNames)

	return WrappedType{name: name, methods: methodsMap, methodNames: methodNames}
}

type Wrapper struct {
	Inner any
	WType *WrappedType
}

func (w Wrapper) String() string {
	switch casted := any(w.Inner).(type) {
	case renderer:
		var buffer strings.Builder
		if casted.Render(&buffer) == nil {
			return buffer.String()
		}
	case bytes.Buffer:
		return casted.String()
	}
	return ""
}

func (w Wrapper) Type() string {
	return w.WType.name
}

func (w Wrapper) Freeze() {
	// No op
}

func (w Wrapper) Truth() starlark.Bool {
	return starlark.True
}

func (w Wrapper) Hash() (uint32, error) {
	return starlark.String(w.String()).Hash()
}

func (w Wrapper) Attr(name string) (starlark.Value, error) {
	if m, ok := w.WType.methods[name]; ok {
		return m.BindReceiver(w), nil
	}
	return nil, nil
}

func (w Wrapper) AttrNames() []string {
	return w.WType.methodNames
}

func (w Wrapper) Code() jen.Code {
	casted, _ := w.Inner.(jen.Code)
	return casted
}
