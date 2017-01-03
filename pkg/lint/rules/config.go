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

package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CloudCoreo/cli/pkg/lint/support"
)

// Config lints a composites's config.yaml file.
func Config(linter *support.Linter) {
	file := "services/config.rb"
	vf := filepath.Join(linter.CompositeDir, file)

	linter.RunLinterRule(support.InfoSev, file, validateConfigFileExistence(linter, vf))
}

func validateConfigFileExistence(linter *support.Linter, configPath string) error {
	_, err := os.Stat(configPath)
	if err != nil {
		return fmt.Errorf("file does not exist")
	}
	return nil
}
