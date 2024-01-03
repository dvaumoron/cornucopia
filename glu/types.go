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
	"errors"
	"path"

	"github.com/dave/jennifer/jen"
	"github.com/dvaumoron/cornucopia/common"
	"go.starlark.net/starlark"
)

//go:generate cornucopia self.crn

var errForbidAbsolute = errors.New("writing on an absolute path is forbidden")

var (
	jenFileWrappedType      wrappedType
	jenStatementWrappedType wrappedType
)

// mutex not needed
var GeneratedFilenames []string

func jenFile_HeaderComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.HeaderComment(comment)
	return starlark.None, nil
}

func jenFile_PackageComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.PackageComment(comment)
	return starlark.None, nil
}

func jenFile_Comment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.Comment(comment)
	return starlark.None, nil
}

func jenFile_Anon(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.Anon(convertToStringSlice(args)...)
	return starlark.None, nil
}

func jenFile_ImportAlias(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var alias string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &alias); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.ImportAlias(path, alias)
	return starlark.None, nil
}

func jenFile_ImportName(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.ImportName(path, name)
	return starlark.None, nil
}

func jenFile_ImportNames(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var names starlark.Dict
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "names", &names); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.ImportNames(convertToMapString(names.Items()))
	return starlark.None, nil
}

func jenFile_Line(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.Line()
	return starlark.None, nil
}

func jenFile_Save(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var filename string
	err := starlark.UnpackArgs(b.Name(), args, kwargs, "filename", &filename)
	if err != nil {
		return nil, err
	}

	if path.IsAbs(filename) {
		return nil, errForbidAbsolute
	}

	if err = common.EnsureWrite(filename); err != nil {
		return nil, err
	}

	recv := b.Receiver().(wrapper[*jen.File])
	if err = recv.inner.Save(filename); err != nil {
		return nil, err
	}

	GeneratedFilenames = append(GeneratedFilenames, filename)
	return starlark.None, nil
}

func jenStatement_Comment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	recv.inner.Comment(comment)
	return starlark.None, nil
}

func jenStatement_Complex(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Complex(convertToCode(m), convertToCode(key))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Dot(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Dot(name)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Id(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Id(name)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Lit(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	lit := recv.inner.Lit(convertToGoBuiltin(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func jenStatement_LitByte(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	lit := recv.inner.LitByte(convertToGoByte(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func jenStatement_LitRune(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	lit := recv.inner.LitRune(convertToGoRune(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Op(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var op string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "op", &op); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Op(op)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Qual(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Qual(path, name)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Tag(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Tag(convertToMapString(kwargs))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Values(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Values(convertToDictOrCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}
