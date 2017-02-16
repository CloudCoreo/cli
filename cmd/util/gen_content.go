// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"os"
	"path/filepath"
)

// CreateFolder create folder
func CreateFolder(name, path string) error {
	fp := filepath.Join(path, name)
	if _, err := os.Stat(fp); err != nil {
		// set default permissions
		os.Mkdir(fp, 0755)
	}

	return nil
}

// CreateFile create a file
func CreateFile(name, path, content string, override bool) error {
	fp := filepath.Join(path, name)
	if _, err := os.Stat(fp); err != nil || override {
		f, _ := os.Create(fp)
		defer f.Close()

		if _, err := f.WriteString(content); err != nil {
			return err
		}
	}

	return nil
}
