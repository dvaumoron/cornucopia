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

package glu

import (
	"errors"
	"fmt"
	"math"

	"go.starlark.net/starlark"
)

var (
	ErrForbidAbsolute = errors.New("writing on an absolute path is forbidden")
	ErrCast           = errors.New("not expected type")
)

// mutex not needed
var GeneratedFileNames []string

func ConvertToGoBaseType(value starlark.Value) any {
	switch casted := value.(type) {
	case starlark.Bool:
		return bool(casted)
	case starlark.Int:
		if res, ok := casted.Int64(); ok {
			if math.MinInt <= res && res <= math.MaxInt {
				return int(res)
			}
			return res
		}
	case starlark.Float:
		return float64(casted)
	case starlark.String:
		return string(casted)
	}
	return nil
}

func ConvertToString(value starlark.Value) string {
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
