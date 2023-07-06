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

	"github.com/dvaumoron/cornucopia/glu"
	"github.com/dvaumoron/cornucopia/module"
	"go.starlark.net/starlark"
)

const defaultRepoUrl = "https://raw.githubusercontent.com/dvaumoron/cornucopiarecipes/main/"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage:\n  cornucopia [filename]")
		return
	}
	scriptname := args[1]

	// TODO manage config
	path := ""
	url := defaultRepoUrl

	loader := module.MakeLoader(path, url)
	thread := &starlark.Thread{Name: "cornucopia: " + scriptname, Load: loader.Load}
	glu.InitCornucopiaGlobals()

	if _, err := starlark.ExecFile(thread, scriptname, nil, nil); err != nil {
		fmt.Println("An error occured :", err)
		if len(glu.GeneratedFilenames) != 0 {
			fmt.Println("Before error, the following file have been generated :")
			for _, filename := range glu.GeneratedFilenames {
				fmt.Println(filename)
			}
		}
		os.Exit(1)
	}

	if len(glu.GeneratedFilenames) == 0 {
		fmt.Println("Successful without file generation")
	} else {
		fmt.Println("Successful, the following file have been generated :")
		for _, filename := range glu.GeneratedFilenames {
			fmt.Println(filename)
		}
	}
}
