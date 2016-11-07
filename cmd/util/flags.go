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
		return fmt.Errorf(content.ERROR_GIT_REPO_URL_MISSING)
	} else if !strings.Contains(gitRepoUrl, "git@") {
		return fmt.Errorf(content.ERROR_INVALID_GIT_REPO_URL)
	}

	return nil
}

func CheckCloudShowOrDeleteFlag(cloudID string) error{
	return checkFlag(cloudID, "Cloud id is required for this command")
}

func CheckCloudAddFlags(resourceName, resourceKey, resourceSecret string) error{
	if err := checkFlag(resourceName, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	if err := checkFlag(resourceKey, content.ERROR_KEY_MISSING); err != nil {
		return err
	}

	if err := checkFlag(resourceSecret, content.ERROR_SECRET_MISSING); err != nil {
		return err
	}

	return nil
}

func CheckTokenShowOrDeleteFlag(tokenID string) error{
	return checkFlag(tokenID, content.ERROR_TOKEN_ID_MISSING)
}

func CheckTokenAddFlags(name, description string) error{
	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	if err := checkFlag(description, content.ERROR_DESCRIPTION_MISSING); err != nil {
		return err
	}

	return nil
}

func CheckGitKeyShowOrDeleteFlag(gitKeyID string) error{
	return checkFlag(gitKeyID, content.ERROR_ID_MISSING)
}

func CheckGitKeyAddFlags(name, secret string) error{
	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	if err := checkFlag(secret, content.ERROR_SECRET_MISSING); err != nil {
		return err
	}

	return nil
}


func CheckCompositeShowOrDeleteFlag(compositeID string) error {
	return checkFlag(compositeID, content.ERROR_ID_MISSING)
}

func CheckCompositeCreateFlags(name, gitRepoUrl string ) error{
	if err := checkGitRepoUrl(gitRepoUrl); err != nil  {
		return err
	}

	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	return nil

}

func CheckLayersFlags(name, gitRepoUrl string) error{
	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
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
			return teamID, fmt.Errorf(content.ERROR_ID_MISSING)
		}
	}

	return teamID, nil
}

func CheckAPIKeyFlag(apiKey string, userProfile string) (string, error) {
	if apiKey == content.NONE {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.ACCESS_KEY)
		apiKey = GetValueFromConfig(teamIDKey, false)

		if apiKey == content.NONE {
			return apiKey, fmt.Errorf(content.ERROR_KEY_MISSING)
		}
	}

	return apiKey, nil
}


func CheckSecretKeyFlag(secretKey string, userProfile string) (string, error) {
	if secretKey == content.NONE {
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.SECRET_KEY)
		secretKey = GetValueFromConfig(teamIDKey, false)

		if secretKey == content.NONE {
			return secretKey, fmt.Errorf(content.ERROR_SECRET_MISSING)
		}
	}

	return secretKey, nil
}