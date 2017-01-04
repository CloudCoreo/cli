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
	"strings"

	"github.com/CloudCoreo/cli/cmd/content"
)

func checkFlag(flag, error string) error {
	if flag == "" {
		return fmt.Errorf(error)
	}

	return nil
}

func checkGitRepoURL(gitRepoURL string) error {
	if gitRepoURL == "" {
		return fmt.Errorf(content.ErrorGitRepoURLMissing)
	} else if !strings.Contains(gitRepoURL, "git@") {
		return fmt.Errorf(content.ErrorInvalidGitRepoURL)
	}

	return nil
}

// CheckCloudShowOrDeleteFlag flags check for cloud show or delete command
func CheckCloudShowOrDeleteFlag(cloudID string, verbose bool) error {

	if err := checkFlag(cloudID, content.ErrorCloudIDRequired); err != nil {
		return err
	}

	if verbose {
		fmt.Printf(content.InfoUsingCloudAccount, cloudID)
	}

	return nil
}

// CheckCloudAddFlags flag check for cloud add command
func CheckCloudAddFlags(resourceName, resourceKey, resourceSecret string) error {
	if err := checkFlag(resourceName, content.ErrorNameMissing); err != nil {
		return err
	}

	if err := checkFlag(resourceKey, content.ErrorKeyMissing); err != nil {
		return err
	}

	if err := checkFlag(resourceSecret, content.ErrorSecretMissing); err != nil {
		return err
	}

	return nil
}

// CheckTokenShowOrDeleteFlag flag check for token show or delete command
func CheckTokenShowOrDeleteFlag(tokenID string, verbose bool) error {
	if err := checkFlag(tokenID, content.ErrorTokenIDMissing); err != nil {
		return err
	}

	if verbose {
		fmt.Printf(content.InfoUsingTokenID, tokenID)
	}

	return nil
}

// CheckGitKeyShowOrDeleteFlag flag check for Git key show or delete command
func CheckGitKeyShowOrDeleteFlag(gitKeyID string, verbose bool) error {
	if err := checkFlag(gitKeyID, content.ErrorGitKeyIDMissing); err != nil {
		return err
	}

	if verbose {
		fmt.Printf(content.InfoUsingGitKeyID, gitKeyID)
	}

	return nil
}

// CheckGitKeyAddFlags flag check for git key add command
func CheckGitKeyAddFlags(name, secret string) error {
	if err := checkFlag(name, content.ErrorNameMissing); err != nil {
		return err
	}

	if err := checkFlag(secret, content.ErrorSecretMissing); err != nil {
		return err
	}

	return nil
}

// CheckCompositeShowOrDeleteFlag flag check for composite show or delete command
func CheckCompositeShowOrDeleteFlag(compositeID string, verbose bool) error {
	if err := checkFlag(compositeID, content.ErrorCompositeIDRequired); err != nil {
		return err
	}

	if verbose {
		fmt.Printf(content.InfoUsingCompsiteID, compositeID)
	}

	return nil

}

// CheckCompositeCreateFlags flags check for composite create command
func CheckCompositeCreateFlags(name, gitRepoURL string) error {
	if err := checkGitRepoURL(gitRepoURL); err != nil {
		return err
	}

	if err := checkFlag(name, content.ErrorNameMissing); err != nil {
		return err
	}

	return nil

}

// CheckLayersFlags flag check for composite layer command
func CheckLayersFlags(name, gitRepoURL string) error {
	if err := checkFlag(name, content.ErrorNameMissing); err != nil {
		return err
	}

	if err := checkGitRepoURL(gitRepoURL); err != nil {

		return err
	}

	return nil
}

// CheckExtendFlags flag check for composite extend command
func CheckExtendFlags(gitRepoURL string) error {
	if err := checkGitRepoURL(gitRepoURL); err != nil {
		return err
	}

	return nil
}

// CheckTeamIDFlag flag check for team id
func CheckTeamIDFlag(teamID, userProfile string, verbose bool) (string, error) {
	if teamID == content.None {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.TeamID)
		teamID = GetValueFromConfig(teamIDKey, false)

		if teamID == content.None {
			return teamID, fmt.Errorf(content.ErrorTeamIDMissing)
		}
	}

	if verbose {
		fmt.Printf(content.InfoUsingTeamID, teamID)
	}

	return teamID, nil
}

// CheckAPIKeyFlag flag check for api key
func CheckAPIKeyFlag(apiKey string, userProfile string) (string, error) {
	if apiKey == content.None {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.AccessKey)
		apiKey = GetValueFromConfig(teamIDKey, false)

		if apiKey == content.None {
			return apiKey, fmt.Errorf(content.ErrorAPIKeyMissing)
		}
	}

	return apiKey, nil
}

// CheckSecretKeyFlag flag check for secret key
func CheckSecretKeyFlag(secretKey string, userProfile string) (string, error) {
	if secretKey == content.None {
		secretIDKey := fmt.Sprintf("%s.%s", userProfile, content.SecretKey)
		secretKey = GetValueFromConfig(secretIDKey, false)

		if secretKey == content.None {
			return secretKey, fmt.Errorf(content.ErrorAPISecretMissing)
		}
	}

	return secretKey, nil
}

//CheckCompositeIDAndPlandIDFlag Check for compositeID and planID
func CheckCompositeIDAndPlandIDFlag(compositeID, planID string, verbose bool) error {
	if err := checkFlag(compositeID, content.ErrorCompositeIDRequired); err != nil {
		return err
	}

	if verbose {
		fmt.Printf(content.InfoUsingCompsiteID, compositeID)
	}

	if err := checkFlag(planID, content.ErrorPlanIDRequired); err != nil {
		return err
	}

	if verbose {
		fmt.Printf(content.InfoUsingPlanID, planID)
	}

	return nil
}

//CheckArgsCount check for args
func CheckArgsCount(args []string) {
	if len(args) > 0 {
		fmt.Print(content.ErrorAcceptsNoArgs)
		os.Exit(-1)
	}
}
