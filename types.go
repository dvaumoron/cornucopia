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
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"go.starlark.net/starlark"
)

var jenFileWrappedType wrappedType
var jenStatementWrappedType wrappedType

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
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "filename", &filename); err != nil {
		return nil, err
	}

	if index := strings.LastIndexByte(filename, '/'); index != -1 {
		if err := os.MkdirAll(filename[:index], 0755); err != nil {
			return nil, err
		}
	}

	recv := b.Receiver().(wrapper[*jen.File])
	if err := recv.inner.Save(filename); err != nil {
		return nil, err
	}
	return starlark.None, nil
}

func jenStatement_Add(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Add(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Append(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Append(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Assert(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var type_ starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "type", &type_); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Assert(convertToCode(type_))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Block(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Block(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Call(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Call(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Cap(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Cap(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
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

func jenStatement_Defs(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Defs(convertToCodeSlice(args)...)
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

func jenStatement_Imag(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Imag(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Index(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Index(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Interface(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Interface(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Len(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Len(convertToCode(value))
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

func jenStatement_Make(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Make(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Map(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var type_ starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "type", &type_); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Map(convertToCode(type_))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
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

func jenStatement_Params(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Params(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Parens(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var item starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "item", &item); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Parens(convertToCode(item))
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

func jenStatement_Range(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Range()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Real(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Real(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Recover(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Recover()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Rune(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Rune()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_String(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.String()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Struct(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Struct(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Tag(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Tag(convertToMapString(kwargs))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Types(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Types(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Uint(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Uint8(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint8()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Uint16(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint16()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Uint32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint32()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Uint64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint64()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Uintptr(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uintptr()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Values(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Values(convertToDictOrCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}
