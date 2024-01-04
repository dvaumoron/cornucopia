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

package go_glu

import (
	"github.com/dave/jennifer/jen"
	"github.com/dvaumoron/cornucopia/glu"
	"go.starlark.net/starlark"
)

func convertToGoByte(value starlark.Value) byte {
	if casted, ok := value.(starlark.Int); ok {
		if res, ok2 := casted.Int64(); ok2 {
			return byte(res)
		}
	}
	return 0
}

func convertToGoRune(value starlark.Value) rune {
	switch casted := value.(type) {
	case starlark.Int:
		if res, ok := casted.Int64(); ok {
			return rune(res)
		}
	case starlark.String:
		for _, r := range string(casted) {
			return r
		}
	}
	return 0
}

func convertToCode(value starlark.Value) jen.Code {
	if casted, ok := value.(glu.Coder); ok {
		return casted.Code()
	}
	return jen.Lit(glu.ConvertToGoBaseType(value))
}

func convertToCodeSlice(args starlark.Tuple) []jen.Code {
	res := make([]jen.Code, 0, len(args))
	for _, arg := range args {
		res = append(res, convertToCode(arg))
	}
	return res
}

func convertToStringSlice(args starlark.Tuple) []string {
	res := make([]string, 0, len(args))
	for _, arg := range args {
		if s, ok := starlark.AsString(arg); ok {
			res = append(res, s)
		}
	}
	return res
}

func convertToMapString(items []starlark.Tuple) map[string]string {
	mapString := make(map[string]string, len(items))
	for _, item := range items {
		key, ok := starlark.AsString(item[0])
		value, ok2 := starlark.AsString(item[1])
		if ok && ok2 {
			mapString[key] = value
		}
	}
	return mapString
}

func convertToDictOrCodeSlice(args starlark.Tuple) []jen.Code {
	if len(args) == 1 {
		if casted, ok := args[0].(*starlark.Dict); ok {
			dict := jen.Dict{}
			for _, item := range casted.Items() {
				dict[convertToCode(item[0])] = convertToCode(item[1])
			}
			return []jen.Code{dict}
		}
	}
	return convertToCodeSlice(args)
}
