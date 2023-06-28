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
	// TODO a lot of adapter

}

func wrappedNewFile(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var packageName starlark.String
	if err := starlark.UnpackArgs("NewFile", args, kwargs, "packageName", &packageName); err != nil {
		return nil, err
	}

	file := jen.NewFile(packageName.GoString())

	// TODO generic wrapper able of calling methods
	return nil, nil
}
