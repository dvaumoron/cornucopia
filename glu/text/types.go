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
	"bytes"
	"path"

	"github.com/dvaumoron/cornucopia/common"
	glu "github.com/dvaumoron/cornucopia/glu"
	"go.starlark.net/starlark"
)

var textFileWrappedType = glu.MakeWrappedType("testFile", starlark.NewBuiltin("Line", textFile_Line), starlark.NewBuiltin("Save", textFile_Save))

func textFile_Line(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(glu.Wrapper[*bytes.Buffer])
	// TODO
	recv.Inner.WriteByte('\n')
	return starlark.None, nil
}

func textFile_Save(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var filename string
	err := starlark.UnpackArgs(b.Name(), args, kwargs, "filename", &filename)
	if err != nil {
		return nil, err
	}

	if path.IsAbs(filename) {
		return nil, glu.ErrForbidAbsolute
	}

	recv := b.Receiver().(glu.Wrapper[*bytes.Buffer])
	if err = common.WriteFile(filename, recv.Inner.Bytes()); err != nil {
		return nil, err
	}

	glu.GeneratedFilenames = append(glu.GeneratedFilenames, filename)
	return starlark.None, nil
}
