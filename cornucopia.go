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
	starlark.Universe["NewFile"] = starlark.NewBuiltin("NewFile", wrappedNewFile)
	starlark.Universe["Bool"] = starlark.NewBuiltin("Bool", wrappedBool)
	starlark.Universe["Case"] = starlark.NewBuiltin("Case", wrappedCase)
	starlark.Universe["Complex64"] = starlark.NewBuiltin("Complex64", wrappedComplex64)
	starlark.Universe["Complex128"] = starlark.NewBuiltin("Complex64", wrappedComplex128)
	starlark.Universe["Empty"] = starlark.NewBuiltin("Empty", wrappedEmpty)
	starlark.Universe["Fallthrough"] = starlark.NewBuiltin("Id", wrappedFallthrough)
	starlark.Universe["Float32"] = starlark.NewBuiltin("Float32", wrappedFloat32)
	starlark.Universe["Float64"] = starlark.NewBuiltin("Float64", wrappedFloat64)
	starlark.Universe["Id"] = starlark.NewBuiltin("Id", wrappedId)
	starlark.Universe["Int"] = starlark.NewBuiltin("Int", wrappedInt)
	starlark.Universe["Int8"] = starlark.NewBuiltin("Int8", wrappedInt8)
	starlark.Universe["Int16"] = starlark.NewBuiltin("Int16", wrappedInt16)
	starlark.Universe["Int32"] = starlark.NewBuiltin("Int32", wrappedInt32)
	starlark.Universe["Int64"] = starlark.NewBuiltin("Int64", wrappedInt64)
	starlark.Universe["Qual"] = starlark.NewBuiltin("Qual", wrappedQual)
	starlark.Universe["Lit"] = starlark.NewBuiltin("Lit", wrappedLit)
	starlark.Universe["Nil"] = starlark.NewBuiltin("Nil", wrappedNil)
	starlark.Universe["Return"] = starlark.NewBuiltin("Return", wrappedReturn)
	starlark.Universe["String"] = starlark.NewBuiltin("Uint", wrappedString)
	starlark.Universe["Uint"] = starlark.NewBuiltin("Uint", wrappedUint)
	starlark.Universe["Uint8"] = starlark.NewBuiltin("Uint8", wrappedUint8)
	starlark.Universe["Uint16"] = starlark.NewBuiltin("Uint16", wrappedUint16)
	starlark.Universe["Uint32"] = starlark.NewBuiltin("Uint32", wrappedUint32)
	starlark.Universe["Uint64"] = starlark.NewBuiltin("Uint64", wrappedUint64)
	// TODO a lot of adapter
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

func wrappedComplex64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Complex64(), wType: &jenStatementWrappedType}, nil
}

func wrappedComplex128(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Complex128(), wType: &jenStatementWrappedType}, nil
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

func wrappedNil(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Nil(), wType: &jenStatementWrappedType}, nil
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
