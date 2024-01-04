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

package go_glu

import (
	"github.com/dave/jennifer/jen"
	"github.com/dvaumoron/cornucopia/glu"
	"go.starlark.net/starlark"
)

//go:generate cornucopia self.crn

func wrappedNewFile(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var packageName string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "packageName", &packageName); err != nil {
		return nil, err
	}

	file := jen.NewFile(packageName)
	return glu.Wrapper{Inner: file, WType: &jenFileWrappedType}, nil
}

func wrappedComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &value); err != nil {
		return nil, err
	}
	jen.Comment(value)
	return starlark.None, nil
}

func wrappedComplex(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}
	stmt := jen.Complex(convertToCode(m), convertToCode(key))
	return glu.Wrapper{Inner: stmt, WType: &jenStatementWrappedType}, nil
}

func wrappedCopy(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var dst starlark.Value
	var src starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "dst", &dst, "src", &src); err != nil {
		return nil, err
	}
	stmt := jen.Copy(convertToCode(dst), convertToCode(src))
	return glu.Wrapper{Inner: stmt, WType: &jenStatementWrappedType}, nil
}

func wrappedDelete(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}
	stmt := jen.Delete(convertToCode(m), convertToCode(key))
	return glu.Wrapper{Inner: stmt, WType: &jenStatementWrappedType}, nil
}

func wrappedId(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}
	id := jen.Id(name)
	return glu.Wrapper{Inner: id, WType: &jenStatementWrappedType}, nil
}

func wrappedLit(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.Lit(glu.ConvertToGoBaseType(value))
	return glu.Wrapper{Inner: lit, WType: &jenStatementWrappedType}, nil
}

func wrappedLitByte(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.LitByte(convertToGoByte(value))
	return glu.Wrapper{Inner: lit, WType: &jenStatementWrappedType}, nil
}

func wrappedLitRune(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.LitRune(convertToGoRune(value))
	return glu.Wrapper{Inner: lit, WType: &jenStatementWrappedType}, nil
}

func wrappedOp(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "op", &value); err != nil {
		return nil, err
	}
	op := jen.Op(value)
	return glu.Wrapper{Inner: op, WType: &jenStatementWrappedType}, nil
}

func wrappedQual(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}
	qual := jen.Qual(path, name)
	return glu.Wrapper{Inner: qual, WType: &jenStatementWrappedType}, nil
}

func wrappedValues(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Values(convertToDictOrCodeSlice(args)...)
	return glu.Wrapper{Inner: stmt, WType: &jenStatementWrappedType}, nil
}
