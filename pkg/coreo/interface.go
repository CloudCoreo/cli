package coreo

import "github.com/CloudCoreo/cli/client"

// Interface for Coreo client for mocking in tests
type Interface interface {
	ListTeams() ([]*client.Team, error)
	ShowTeamByID(teamID string) (*client.Team, error)

	ListTokens() ([]*client.Token, error)
	ShowTokenByID(tokenID string) (*client.Token, error)
	DeleteTokenByID(tokenID string) error

	ListCloudAccounts(teamID string) ([]*client.CloudAccount, error)
	ShowCloudAccountByID(teamID, cloudID string) (*client.CloudAccount, error)
	CreateCloudAccount(teamID, resourceKey, resourceSecret, resourceName string) (*client.CloudAccount, error)
	DeleteCloudAccountByID(teamID, cloudID string) error

	ListGitKeys(teamID string) ([]*client.GitKey, error)
	ShowGitKeyByID(teamID, gitKeyID string) (*client.GitKey, error)
	CreateGitKey(teamID, resourceSecret, resourceName string) (*client.GitKey, error)
	DeleteGitKeyByID(teamID, gitKeyID string) error

	ListComposites(teamID string) ([]*client.Composite, error)
	ShowCompositeByID(teamID, compositeID string) (*client.Composite, error)
	CreateComposite(teamID, gitRepoURL, name string) (*client.Composite, error)

	ListPlans(teamID, compositeID string) ([]*client.Plan, error)
	ShowPlanByID(teamID, compositeID, planID string) (*client.Plan, error)
	EnablePlanByID(teamID, compositeID, planID string) (*client.Plan, error)
	DisablePlanByID(teamID, compositeID, planID string) (*client.Plan, error)
	DeletePlanByID(teamID, compositeID, planID string) error
}
