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

package marshal_glu

import (
	"encoding/json"
	"path"

	"github.com/dvaumoron/cornucopia/common"
	"github.com/dvaumoron/cornucopia/glu"
	"go.starlark.net/starlark"
	"gopkg.in/yaml.v3"
)

func InitCornucopiaMarshalGlobals() {
	starlark.Universe["SaveJsonFile"] = starlark.NewBuiltin("SaveJsonFile", wrappedSaveJsonFile)
	starlark.Universe["SaveYamlFile"] = starlark.NewBuiltin("SaveYamlFile", wrappedSaveYamlFile)
}

func wrappedSaveJsonFile(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	return saveMarshallFile(json.Marshal, b, args, kwargs)
}

func wrappedSaveYamlFile(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	return saveMarshallFile(yaml.Marshal, b, args, kwargs)
}

func saveMarshallFile(marshalFunc func(any) ([]byte, error), b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var fileName string
	var object starlark.Value
	err := starlark.UnpackArgs(b.Name(), args, kwargs, "fileName", &fileName, "object", &object)
	if err != nil {
		return nil, err
	}

	if path.IsAbs(fileName) {
		return nil, glu.ErrForbidAbsolute
	}

	data, err := marshalFunc(convertToMarshal(object))
	if err != nil {
		return nil, err
	}

	if err = common.WriteFile(fileName, data); err != nil {
		return nil, err
	}

	glu.GeneratedFileNames = append(glu.GeneratedFileNames, fileName)
	return starlark.None, nil
}

func convertToMarshal(value starlark.Value) any {
	switch casted := value.(type) {
	case starlark.Tuple:
		res := make([]any, len(casted))
		for _, elem := range casted {
			res = append(res, convertToMarshal(elem))
		}
		return res
	case *starlark.Dict:
		res := make(map[string]any, casted.Len())
		for _, item := range casted.Items() {
			res[glu.ConvertToString(item[0])] = convertToMarshal(item[1])
		}
		return res
	}
	return glu.ConvertToGoBaseType(value)
}
