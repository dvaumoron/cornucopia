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

package common

import (
	"os"
	"strings"
)

const Prefix = "Cornucopia: "

// Create the parents directories if needed and write the file
func WriteFile(path string, data []byte) error {
	if err := EnsureWrite(path); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// Create the parents directories if needed
func EnsureWrite(path string) error {
	if index := strings.LastIndexByte(path, '/'); index != -1 {
		return os.MkdirAll(path[:index], 0755)
	}
	return nil
}
