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

package module

import (
	"errors"

	"go.starlark.net/starlark"
)

var errLoadCycle = errors.New("cycle in load graph")

type entry struct {
	globals starlark.StringDict
	err     error
}

type ModuleLoader struct {
	cache map[string]*entry
	path  string
	url   string
}

func MakeLoader(repositoryPath string, repositoryUrl string) ModuleLoader {
	return ModuleLoader{cache: map[string]*entry{}, path: repositoryPath, url: repositoryUrl}
}

func (ml ModuleLoader) Load(thread *starlark.Thread, modulename string) (starlark.StringDict, error) {
	e, ok := ml.cache[modulename]
	if e == nil {
		if ok {
			// request for package whose loading is in progress
			return nil, errLoadCycle
		}

		// Add a placeholder to indicate "load in progress".
		ml.cache[modulename] = nil

		src := ml.read(modulename)

		// Load it.
		thread := &starlark.Thread{Name: "cornucopia: " + modulename, Load: thread.Load}
		globals, err := starlark.ExecFile(thread, modulename, src, nil)
		e = &entry{globals, err}

		// Update the cache.
		ml.cache[modulename] = e
	}
	return e.globals, e.err
}

// Read data, trying the following resolution :
//   - read from current directory as base path
//   - read from local repository path as base path
//   - download from repository url and write the content in local repository path
func (ml ModuleLoader) read(modulename string) []byte {
	// TODO
	return nil
}
