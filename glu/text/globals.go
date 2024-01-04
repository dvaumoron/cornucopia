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

package text_glu

import (
	"strings"

	"github.com/dvaumoron/cornucopia/glu"
	"go.starlark.net/starlark"
)

func InitCornucopiaTextGlobals() {
	starlark.Universe["NewTextFile"] = starlark.NewBuiltin("NewTextFile", wrappedNewTextFile)
}

func wrappedNewTextFile(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var packageName string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "packageName", &packageName); err != nil {
		return nil, err
	}

	builder := new(strings.Builder)
	return glu.Wrapper{Inner: builder, WType: &textFileWrappedType}, nil
}
