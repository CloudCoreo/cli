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
	"fmt"
	"os"
	"os/exec"

	"github.com/CloudCoreo/cli/cmd/content"
)

// CheckGitInstall check if user has git installed
func CheckGitInstall() error {
	cmd := exec.Command("git", "version")

	if _, err := cmd.Output(); err != nil {
		println("Unable to execute 'git': " + err.Error())
		return err
	}

	return nil
}

// CreateGitSubmodule create composite git submodule
func CreateGitSubmodule(directory, gitRepoURL string) error {

	fmt.Printf(content.InfoCreatingGitSubmodule, gitRepoURL, directory)
	if directory == "" {
		return fmt.Errorf(content.ErrorDirectoryPathNotProvided)
	}

	if gitRepoURL == "" {
		return fmt.Errorf(content.ErrorGitURLNotProvided)
	}

	err := os.Chdir(directory)
	if err != nil {
		return fmt.Errorf(content.ErrorInvalidDirectory, directory)
	}

	_, err = exec.Command("git", "rev-parse", "--is-inside-work-tree").CombinedOutput()
	if err != nil {
		return fmt.Errorf(content.ErrorGitInitNotRan)
	}

	output, err := exec.Command("git", "submodule", "add", "-f", gitRepoURL, "extends").CombinedOutput()
	if err != nil {
		return fmt.Errorf(content.ErrorGitSubmoduleFailed, output)
	}

	return nil
}
