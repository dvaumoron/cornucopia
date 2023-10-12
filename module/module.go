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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/dvaumoron/cornucopia/common"
	"github.com/dvaumoron/cornucopia/config"
	"go.starlark.net/starlark"
)

var errLoadCycle = errors.New("cycle in load graph")

type entry struct {
	globals starlark.StringDict
	err     error
}

type ModuleLoader struct {
	threadPrefix string
	conf         *config.Config
	cache        map[string]*entry
}

func MakeLoader(threadPrefix string, conf *config.Config) ModuleLoader {
	return ModuleLoader{threadPrefix: threadPrefix, conf: conf, cache: map[string]*entry{}}
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
//   - read with current directory as base path
//   - read with local repository path as base path (skipped with force download flag)
//   - download from repository url and write the content in local repository path
func (ml ModuleLoader) read(modulename string) ([]byte, error) {
	data, err := os.ReadFile(modulename)
	if err == nil || path.IsAbs(modulename) {
		return data, err
	}

	verbose := ml.conf.Verbose
	if verbose {
		fmt.Println("Failed to resolve", modulename, "from current directory :", err)
	}

	modulePath := path.Join(ml.conf.RepoPath, modulename)
	if !ml.conf.ForceDownload {
		if data, err = os.ReadFile(modulePath); err == nil {
			return data, nil
		}

		if verbose {
			fmt.Println("Failed to resolve", modulename, "from", modulePath, ":", err)
		}
	}

	moduleUrl, err := url.JoinPath(ml.conf.RepoUrl, modulename)
	if err != nil {
		return nil, err
	}

	if data, err = unsecureDownload(moduleUrl); err != nil {
		return nil, err
	}
	return data, common.WriteFile(modulePath, data)
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
	return io.ReadAll(resp.Body)
}
