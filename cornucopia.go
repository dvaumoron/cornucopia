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
	starlark.Universe["Empty"] = starlark.NewBuiltin("Empty", wrappedEmpty)
	starlark.Universe["Id"] = starlark.NewBuiltin("Id", wrappedId)
	starlark.Universe["Qual"] = starlark.NewBuiltin("Qual", wrappedQual)
	starlark.Universe["Lit"] = starlark.NewBuiltin("Lit", wrappedLit)
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

func wrappedEmpty(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Empty(), wType: &jenStatementWrappedType}, nil
}

func wrappedId(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}

	id := jen.Id(name)
	return wrapper[*jen.Statement]{inner: id, wType: &jenStatementWrappedType}, nil
}

func wrappedQual(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}
	lit := jen.Qual(path, name)
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedLit(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.Lit(convertToGoBuiltin(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}
