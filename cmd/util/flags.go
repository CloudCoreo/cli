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

	"github.com/CloudCoreo/cli/cmd/content"
)

func checkFlag(flag, error string) error {
	if flag == "" {
		return fmt.Errorf(error)
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
func CheckCloudAddFlags(externalID, roleArn, roleName, environment string) error {

	if (externalID == "" || roleArn == "") && roleName == "" {
		return fmt.Errorf("Please either provide both externalID and roleArn or the name of the new role ")
	}

	// Check for environment set
	envSet := map[string]bool{
		"Production":  true,
		"Staging":     true,
		"Development": true,
		"Test":        true,
	}

	if !envSet[environment] && environment != "" {
		return fmt.Errorf("Environment must be one of those: Production, Staging, Development, Test ")
	}
	return nil
}

// CheckCloudAddFlags flag check for cloud add command when adding AWS cloud account
func CheckCloudAddFlagsForAWS(externalID, roleArn, roleName, environment string) error {
	if (externalID == "" || roleArn == "") && roleName == "" {
		return fmt.Errorf("Please either provide both externalID and roleArn or the name of the new role ")
	}
	return checkEnvironment(environment)
}

// CheckCloudAddFlags flag check for cloud add command when adding azure cloud account
func CheckCloudAddFlagsForAzure(keyValue, applicationID, directoryID, subscriptionID, environment string) error {
	if keyValue == "" || applicationID == "" || directoryID == "" || subscriptionID == "" {
		return fmt.Errorf("Please provide all the required info: Key Value, Application ID, Directory ID and Subscription ID ")
	}
	return checkEnvironment(environment)
}

func checkEnvironment(environment string) error {
	envSet := map[string]bool{
		"Production":  true,
		"Staging":     true,
		"Development": true,
		"Test":        true,
	}

	if !envSet[environment] && environment != "" {
		return fmt.Errorf("Environment must be one of those: Production, Staging, Development, Test ")
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

//CheckTeamAddFlags Required flags check
func CheckTeamAddFlags(teamName, teamDescription string) error {
	if err := checkFlag(teamName, content.ErrorTeamNameRequired); err != nil {
		return err
	}

	return checkFlag(teamDescription, content.ErrorTeamDescriptionRequired)
}

//CheckArgsCount check for args
func CheckArgsCount(args []string) {
	if len(args) > 0 {
		fmt.Print(content.ErrorAcceptsNoArgs)
		os.Exit(-1)
	}
}
