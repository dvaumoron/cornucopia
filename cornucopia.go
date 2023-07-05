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

	"go.starlark.net/repl"
	"go.starlark.net/starlark"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage:\n  cornucopia [filename]")
		return
	}

	load := repl.MakeLoad() // TODO check the utility of a better version
	thread := &starlark.Thread{Name: "cornucopia", Load: load}
	initCornucopiaGlobals()

	_, err := starlark.ExecFile(thread, args[1], nil, nil)
	if err != nil {
		fmt.Println("An error occured :", err)
		os.Exit(1)
	}

	if len(generatedFilenames) == 0 {
		fmt.Println("Successful without file generation")
	} else {
		fmt.Println("Successful, the following file have been generated :")
		for _, filename := range generatedFilenames {
			fmt.Println(filename)
		}
	}
}
