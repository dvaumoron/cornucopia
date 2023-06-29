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
		starlark.NewBuiltin("Call", jenStatement_Call),
		starlark.NewBuiltin("Delete", jenStatement_Delete),
		starlark.NewBuiltin("Op", jenStatement_Op),
		starlark.NewBuiltin("Params", jenStatement_Params),
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

func jenFile_Anon(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var pathIt starlark.Iterable
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "paths", &pathIt); err != nil {
		return nil, err
	}
	it := pathIt.Iterate()
	defer it.Done()
	var paths []string
	var value starlark.Value
	for it.Next(&value) {
		path, ok := starlark.AsString(value)
		if ok {
			paths = append(paths, path)
		}
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.Anon(paths...)
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
	for _, t := range items {
		packageRef, ok := starlark.AsString(t.Index(0))
		packageName, ok2 := starlark.AsString(t.Index(1))
		if ok && ok2 {
			names[packageRef] = packageName
		}
	}
	recv := b.Receiver().(wrapper[*jen.File])
	recv.inner.ImportNames(names)
	return starlark.None, nil
}

func jenFile_Const(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Const()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Var(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Var()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Type(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Type()
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenFile_Func(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
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

func jenStatement_Add(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Add(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Call(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Call(convertToCodeSlice(args)...)
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

func jenStatement_Op(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var op string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "op", &op); err != nil {
		return nil, err
	}
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Op(op)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func jenStatement_Params(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Params(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}
