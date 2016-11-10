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
func CreateGitSubmodule(directory, gitRepoUrl string) error {

	fmt.Printf("INFO_CREATING_GITSUBMODULE", gitRepoUrl, directory)
	if directory == "" {
		return fmt.Errorf(content.ERROR_DIRECTORY_PATH_NOT_PROVIDED)
	}

	if gitRepoUrl == "" {
		return fmt.Errorf(content.ERROR_GIT_URL_NOT_PROVIDED)
	}

	err := os.Chdir(directory)
	if err != nil {
		return fmt.Errorf(content.ERROR_INVALID_DIRECTORY, directory)
	}

	_, err = exec.Command("git", "remote",  "show", "origin").CombinedOutput()
	if  err != nil {
		return fmt.Errorf(content.ERROR_INVALID_DIRECTORY)
	}

	output, err := exec.Command("git", "submodule", "add", "-f", gitRepoUrl, "extends").CombinedOutput()
	if err != nil {
		return fmt.Errorf(content.ERROR_GIT_SUBMODULE_FAILED, output)
	}

	return nil
}