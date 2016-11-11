package util

import (
	"fmt"
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
		return fmt.Errorf(content.ERROR_GIT_REPO_URL_MISSING)
	} else if !strings.Contains(gitRepoURL, "git@") {
		return fmt.Errorf(content.ERROR_INVALID_GIT_REPO_URL)
	}

	return nil
}

// CheckCloudShowOrDeleteFlag flags check for cloud show or delete command
func CheckCloudShowOrDeleteFlag(cloudID string) error {
	return checkFlag(cloudID, "Cloud id is required for this command")
}

// CheckCloudAddFlags flag check for cloud add command
func CheckCloudAddFlags(resourceName, resourceKey, resourceSecret string) error {
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

// CheckTokenShowOrDeleteFlag flag check for token show or delete command
func CheckTokenShowOrDeleteFlag(tokenID string) error {
	return checkFlag(tokenID, content.ERROR_ID_MISSING)
}

// CheckTokenAddFlags flag check for token add command
func CheckTokenAddFlags(name, description string) error {
	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	if err := checkFlag(description, content.ERROR_DESCRIPTION_MISSING); err != nil {
		return err
	}

	return nil
}

// CheckGitKeyShowOrDeleteFlag flag check for Git key show or delete command
func CheckGitKeyShowOrDeleteFlag(gitKeyID string) error {
	return checkFlag(gitKeyID, content.ERROR_ID_MISSING)
}

// CheckGitKeyAddFlags flag check for git key add command
func CheckGitKeyAddFlags(name, secret string) error {
	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	if err := checkFlag(secret, content.ERROR_SECRET_MISSING); err != nil {
		return err
	}

	return nil
}

// CheckCompositeShowOrDeleteFlag flag check for composite show or delete command
func CheckCompositeShowOrDeleteFlag(compositeID string) error {
	return checkFlag(compositeID, content.ERROR_ID_MISSING)
}

// CheckCompositeCreateFlags flags check for composite create command
func CheckCompositeCreateFlags(name, gitRepoURL string) error {
	if err := checkGitRepoURL(gitRepoURL); err != nil {
		return err
	}

	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
		return err
	}

	return nil

}

// CheckLayersFlags flag check for composite layer command
func CheckLayersFlags(name, gitRepoURL string) error {
	if err := checkFlag(name, content.ERROR_NAME_MISSING); err != nil {
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

// CheckAPIKeyFlag flag check for api key
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

// CheckSecretKeyFlag flag check for secret key
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

// Check for compositeID and planID
func CheckCompositeIdAndPlandIdFlag(compositeID, planID string) error {
	if err := checkFlag(compositeID, content.ERROR_ID_MISSING); err != nil {
		return err
	}

	if err := checkFlag(planID, content.ERROR_ID_MISSING); err != nil {
		return err
	}

	return nil
}
