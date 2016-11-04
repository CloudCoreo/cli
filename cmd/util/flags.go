package util

import (
	"fmt"
	"strings"
	"github.com/cloudcoreo/cli/cmd/content"
)

func checkFlag(flag, error string) error {
	if flag == "" {
		return fmt.Errorf(error)
	}

	return nil
}

func checkGitRepoUrl(gitRepoUrl string) error {
	if gitRepoUrl == "" {
		return fmt.Errorf("A SSH git repo url is required: -g")
	} else if !strings.Contains(gitRepoUrl, "git@") {
		return fmt.Errorf("Use a SSH git repo url for example : [-g git@github.com:CloudCoreo/audit-aws.git]")
	}

	return nil
}

func CheckCloudShowOrDeleteFlag(cloudID string) error{
	return checkFlag(cloudID, "Cloud id is required for this command")
}

func CheckCloudAddFlags(resourceName, resourceKey, resourceSecret string) error{
	if err := checkFlag(resourceName, "Resource name is required for this command"); err != nil {
		return err
	}

	if err := checkFlag(resourceKey, "Resource key is required for this command"); err != nil {
		return err
	}

	if err := checkFlag(resourceSecret, "Resource secret is required for this command"); err != nil {
		return err
	}

	return nil
}

func CheckTokenShowOrDeleteFlag(tokenID string) error{
	return checkFlag(tokenID, "Token id is required for this command")
}

func CheckTokenAddFlags(name, description string) error{
	if err := checkFlag(name, "Name flag is required for this command"); err != nil {
		return err
	}

	if err := checkFlag(description, "Description flag is required for this command"); err != nil {
		return err
	}

	return nil
}

func CheckGitKeyShowOrDeleteFlag(gitKeyID string) error{
	return checkFlag(gitKeyID, "gitKeyID id is required for this command")
}

func CheckGitKeyAddFlags(name, secret string) error{
	if err := checkFlag(name, "Name flag is required for this command"); err != nil {
		return err
	}

	if err := checkFlag(secret, "secret flag is required for this command"); err != nil {
		return err
	}

	return nil
}


func CheckLayersFlags(name, gitRepoUrl string) error{
	if err := checkFlag(name, "A composite name is required: -n"); err != nil {
		return err
	}

	if err := checkGitRepoUrl(gitRepoUrl); err != nil  {
		return err
	}

	return nil
}

func CheckExtendFlags(gitRepoUrl string) error{
	if err := checkGitRepoUrl(gitRepoUrl); err != nil  {
		return err
	}

	return nil
}

func CheckTeamIDFlag(teamID string, userProfile string) (string, error) {
	if teamID == content.NONE {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.TEAM_ID)
		teamID = GetValueFromConfig(teamIDKey, false)

		if teamID == content.NONE {
			return teamID, fmt.Errorf("Team ID is required for this command")
		}
	}

	return teamID, nil
}

func CheckAPIKeyFlag(apiKey string, userProfile string) (string, error) {
	if apiKey == content.NONE {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.ACCESS_KEY)
		apiKey = GetValueFromConfig(teamIDKey, false)

		if apiKey == content.NONE {
			return apiKey, fmt.Errorf("API key is required for this command")
		}
	}

	return apiKey, nil
}


func CheckSecretKeyFlag(secretKey string, userProfile string) (string, error) {
	if secretKey == content.NONE {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.SECRET_KEY)
		secretKey = GetValueFromConfig(teamIDKey, false)

		if secretKey == content.NONE {
			return secretKey, fmt.Errorf("Secret key is required for this command")
		}
	}

	return secretKey, nil
}