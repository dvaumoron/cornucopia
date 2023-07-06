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
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/dvaumoron/cornucopia/common"
	"go.starlark.net/starlark"
)

var errLoadCycle = errors.New("cycle in load graph")

type entry struct {
	globals starlark.StringDict
	err     error
}

type ModuleLoader struct {
	threadPrefix string
	path         string
	url          string
	cache        map[string]*entry
}

func MakeLoader(threadPrefix string, repositoryPath string, repositoryUrl string) ModuleLoader {
	return ModuleLoader{threadPrefix: threadPrefix, path: repositoryPath, url: repositoryUrl, cache: map[string]*entry{}}
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

		e = &entry{}
		var src []byte
		src, e.err = ml.read(modulename)
		if e.err == nil {
			// Load it.
			thread := &starlark.Thread{Name: ml.threadPrefix + modulename, Load: thread.Load}
			e.globals, e.err = starlark.ExecFile(thread, modulename, src, nil)
		}

		// Update the cache.
		ml.cache[modulename] = e
	}
	return e.globals, e.err
}

// Read data, trying the following resolution for relative path :
//   - read from current directory as base path
//   - read from local repository path as base path
//   - download from repository url and write the content in local repository path
func (ml ModuleLoader) read(modulename string) ([]byte, error) {
	data, err := os.ReadFile(modulename)
	if err == nil || path.IsAbs(modulename) {
		return data, err
	}

	repoPath := path.Join(ml.path, modulename)
	if data, err = os.ReadFile(repoPath); err == nil {
		return data, nil
	}

	dUrl, err := url.JoinPath(ml.url, modulename)
	if err != nil {
		return nil, err
	}

	if data, err = unsecureDownload(dUrl); err != nil {
		return nil, err
	}

	if err = common.WriteFile(repoPath, data); err != nil {
		return nil, err
	}
	return data, nil
}

func unsecureDownload(dUrl string) ([]byte, error) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(dUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// supposing module will not be "too big"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
