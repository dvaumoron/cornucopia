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
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"go.starlark.net/repl"
	"go.starlark.net/starlark"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Usage:\n  cornucopia [filename]")
		return
	}

	load := repl.MakeLoad() // TODO check the utility of a better version
	thread := &starlark.Thread{Name: "cornucopia", Load: load}
	initCornucopiaGlobals()

	_, err := starlark.ExecFile(thread, args[0], nil, nil)
	if err != nil {
		fmt.Println("An error occured :", err)
		os.Exit(1)
	}

	fmt.Println("Successful")
}

func initCornucopiaGlobals() {
	addGlobals("NewFile", wrappedNewFile)
	addGlobals("Bool", wrappedBool)
	addGlobals("Case", wrappedCase)
	addGlobals("Chan", wrappedChan)
	addGlobals("Comment", wrappedComment)
	addGlobals("Complex64", wrappedComplex64)
	addGlobals("Complex128", wrappedComplex128)
	addGlobals("Default", wrappedDefault)
	addGlobals("Empty", wrappedEmpty)
	addGlobals("Fallthrough", wrappedFallthrough)
	addGlobals("Float32", wrappedFloat32)
	addGlobals("Float64", wrappedFloat64)
	addGlobals("Id", wrappedId)
	addGlobals("Int", wrappedInt)
	addGlobals("Int8", wrappedInt8)
	addGlobals("Int16", wrappedInt16)
	addGlobals("Int32", wrappedInt32)
	addGlobals("Int64", wrappedInt64)
	addGlobals("Parens", wrappedParens)
	addGlobals("Qual", wrappedQual)
	addGlobals("Lit", wrappedLit)
	addGlobals("LitByte", wrappedLitByte)
	addGlobals("LitRune", wrappedLitRune)
	addGlobals("Nil", wrappedNil)
	addGlobals("Op", wrappedOp)
	addGlobals("Return", wrappedReturn)
	addGlobals("String", wrappedString)
	addGlobals("Uint", wrappedUint)
	addGlobals("Uint8", wrappedUint8)
	addGlobals("Uint16", wrappedUint16)
	addGlobals("Uint32", wrappedUint32)
	addGlobals("Uint64", wrappedUint64)
}

func addGlobals(name string, wrapped func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)) {
	starlark.Universe[name] = starlark.NewBuiltin(name, wrapped)
}

func wrappedNewFile(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var packageName string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "packageName", &packageName); err != nil {
		return nil, err
	}

	file := jen.NewFile(packageName)
	return wrapper[*jen.File]{inner: file, wType: &jenFileWrappedType}, nil
}

func wrappedBool(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Bool(), wType: &jenStatementWrappedType}, nil
}

func wrappedCase(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	case_ := jen.Case(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: case_, wType: &jenStatementWrappedType}, nil
}

func wrappedChan(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Chan(), wType: &jenStatementWrappedType}, nil
}

func wrappedComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &value); err != nil {
		return nil, err
	}
	jen.Comment(value)
	return starlark.None, nil
}

func wrappedComplex64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Complex64(), wType: &jenStatementWrappedType}, nil
}

func wrappedComplex128(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Complex128(), wType: &jenStatementWrappedType}, nil
}

func wrappedDefault(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Default(), wType: &jenStatementWrappedType}, nil
}

func wrappedEmpty(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Empty(), wType: &jenStatementWrappedType}, nil
}

func wrappedFallthrough(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Fallthrough(), wType: &jenStatementWrappedType}, nil
}

func wrappedFloat32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Float32(), wType: &jenStatementWrappedType}, nil
}

func wrappedFloat64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Float64(), wType: &jenStatementWrappedType}, nil
}

func wrappedId(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}
	id := jen.Id(name)
	return wrapper[*jen.Statement]{inner: id, wType: &jenStatementWrappedType}, nil
}

func wrappedInt(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int8(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int16(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int32(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int64(), wType: &jenStatementWrappedType}, nil
}

func wrappedParens(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var item starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "item", &item); err != nil {
		return nil, err
	}
	parens := jen.Parens(convertToCode(item))
	return wrapper[*jen.Statement]{inner: parens, wType: &jenStatementWrappedType}, nil
}

func wrappedQual(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}
	qual := jen.Qual(path, name)
	return wrapper[*jen.Statement]{inner: qual, wType: &jenStatementWrappedType}, nil
}

func wrappedLit(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.Lit(convertToGoBuiltin(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedLitByte(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.LitByte(convertToGoByte(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedLitRune(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.LitRune(convertToGoRune(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedNil(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Nil(), wType: &jenStatementWrappedType}, nil
}

func wrappedOp(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "op", &value); err != nil {
		return nil, err
	}
	op := jen.Op(value)
	return wrapper[*jen.Statement]{inner: op, wType: &jenStatementWrappedType}, nil
}

func wrappedReturn(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return_ := jen.Return(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: return_, wType: &jenStatementWrappedType}, nil
}

func wrappedString(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.String(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint8(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint16(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint32(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint64(), wType: &jenStatementWrappedType}, nil
}
