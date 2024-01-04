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
	"fmt"

	"go.starlark.net/starlark"
)

func convertToString(value starlark.Value) string {
	switch casted := value.(type) {
	case starlark.Bool:
		if casted {
			return "true"
		}
		return "false"
	case starlark.Int:
		if res, ok := casted.Int64(); ok {
			return fmt.Sprint(res)
		}
	case starlark.Float:
		return fmt.Sprint(float64(casted))
	case starlark.String:
		return string(casted)
	}
	return ""
}

func convertToStringSlice(args starlark.Tuple) []string {
	res := make([]string, 0, len(args))
	for _, arg := range args {
		res = append(res, convertToString(arg))
	}
	return res
}
