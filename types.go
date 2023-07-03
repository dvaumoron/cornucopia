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
	"github.com/dave/jennifer/jen"
	"go.starlark.net/starlark"
)

var jenFileWrappedType wrappedType
var jenStatementWrappedType wrappedType

func init() {
	jenFileWrappedType = makeWrappedType("jenFile",
		starlark.NewBuiltin("HeaderComment", jenFile_HeaderComment),
		starlark.NewBuiltin("PackageComment", jenFile_PackageComment),
		starlark.NewBuiltin("Anon", jenFile_Anon),
		starlark.NewBuiltin("ImportAlias", jenFile_ImportAlias),
		starlark.NewBuiltin("ImportName", jenFile_ImportName),
		starlark.NewBuiltin("ImportNames", jenFile_ImportNames),
		starlark.NewBuiltin("Comment", jenFile_Comment),
		starlark.NewBuiltin("Const", jenFile_Const),
		starlark.NewBuiltin("Var", jenFile_Var),
		starlark.NewBuiltin("Type", jenFile_Type),
		starlark.NewBuiltin("Func", jenFile_Func),
		starlark.NewBuiltin("Save", jenFile_Save),
	)

	jenStatementWrappedType = makeWrappedType("jenStatement",
		starlark.NewBuiltin("Add", jenStatement_Add),
		starlark.NewBuiltin("Append", jenStatement_Append),
		starlark.NewBuiltin("Assert", jenStatement_Assert),
		starlark.NewBuiltin("Block", jenStatement_Block),
		starlark.NewBuiltin("Bool", jenStatement_Bool),
		starlark.NewBuiltin("Call", jenStatement_Call),
		starlark.NewBuiltin("Comment", jenStatement_Comment),
		starlark.NewBuiltin("Complex", jenStatement_Complex),
		starlark.NewBuiltin("Complex64", jenStatement_Complex64),
		starlark.NewBuiltin("Complex128", jenStatement_Complex128),
		starlark.NewBuiltin("Copy", jenStatement_Copy),
		starlark.NewBuiltin("Delete", jenStatement_Delete),
		starlark.NewBuiltin("Dot", jenStatement_Dot),
		starlark.NewBuiltin("Float32", jenStatement_Float32),
		starlark.NewBuiltin("Float64", jenStatement_Float64),
		starlark.NewBuiltin("For", jenStatement_For),
		starlark.NewBuiltin("If", jenStatement_If),
		starlark.NewBuiltin("Index", jenStatement_Index),
		starlark.NewBuiltin("Int", jenStatement_Int),
		starlark.NewBuiltin("Int8", jenStatement_Int8),
		starlark.NewBuiltin("Int16", jenStatement_Int16),
		starlark.NewBuiltin("Int32", jenStatement_Int32),
		starlark.NewBuiltin("Int64", jenStatement_Int64),
		starlark.NewBuiltin("Interface", jenStatement_Interface),
		starlark.NewBuiltin("Map", jenStatement_Map),
		starlark.NewBuiltin("Op", jenStatement_Op),
		starlark.NewBuiltin("Params", jenStatement_Params),
		starlark.NewBuiltin("Select", jenStatement_Select),
		starlark.NewBuiltin("String", jenStatement_String),
		starlark.NewBuiltin("Struct", jenStatement_Struct),
		starlark.NewBuiltin("Switch", jenStatement_Switch),
		starlark.NewBuiltin("Types", jenStatement_Types),
		starlark.NewBuiltin("Uint", jenStatement_Uint),
		starlark.NewBuiltin("Uint8", jenStatement_Uint8),
		starlark.NewBuiltin("Uint16", jenStatement_Uint16),
		starlark.NewBuiltin("Uint32", jenStatement_Uint32),
		starlark.NewBuiltin("Uint64", jenStatement_Uint64),
		starlark.NewBuiltin("Values", jenStatement_Values),
		// TODO add useful Statement methods
	)
}

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
	var nameDict starlark.Dict
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "names", &nameDict); err != nil {
		return nil, err
	}
	items := nameDict.Items()
	names := make(map[string]string, len(items))
	for _, item := range items {
		packageRef, ok := starlark.AsString(item[0])
		packageName, ok2 := starlark.AsString(item[1])
		if ok && ok2 {
			names[packageRef] = packageName
		}
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.ImportNames(names)
	return starlark.None, nil
}

func jenFile_Const(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Const()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Var(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Var()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Type(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Type()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Func(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Func()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Save(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var filename string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "filename", &filename); err != nil {
		return nil, err
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

func jenStatement_Bool(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Bool()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Call(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Call(convertToCodeSlice(args)...)
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

func jenStatement_Complex64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Complex64()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Complex128(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Complex128()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Copy(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var dst starlark.Value
	var src starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "dst", &dst, "src", &src); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Copy(convertToCode(dst), convertToCode(src))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Delete(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Delete(convertToCode(m), convertToCode(key))
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

func jenStatement_Float32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Float32()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Float64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Float64()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_For(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.For(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_If(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.If(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Index(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Index(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Int(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Int8(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int8()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Int16(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int16()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Int32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int32()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Int64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int64()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Interface(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Interface(convertToCodeSlice(args)...)
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

func jenStatement_Select(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Select()
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

func jenStatement_Switch(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Switch(convertToCodeSlice(args)...)
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

func jenStatement_Values(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	var stmt *jen.Statement
	ok := false
	if len(args) == 1 {
		var casted *starlark.Dict
		casted, ok = args[0].(*starlark.Dict)
		if ok {
			dict := jen.Dict{}
			for _, item := range casted.Items() {
				dict[convertToCode(item[0])] = convertToCode(item[1])
			}
			stmt = recv.inner.Values(dict)
		}
	}
	if !ok {
		stmt = recv.inner.Values(convertToCode(args))
	}
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}
