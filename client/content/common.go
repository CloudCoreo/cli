package content

const (
	//ErrorMissingAPIOrSecretKey error
	ErrorMissingAPIOrSecretKey = "Missing API key or/and Secret key. Please run 'coreo configure' to configure them."

	//ErrorNoCloudAccountsFound error
	ErrorNoCloudAccountsFound = "No cloud accounts found under team ID %s."

	//ErrorNoCloudAccountWithIDFound error
	ErrorNoCloudAccountWithIDFound = "No cloud account with ID %s found under team ID %s."

	//ErrorFailedToCreateCloudAccount error
	ErrorFailedToCreateCloudAccount = "Failed to create cloud account under team ID %s."

	//ErrorFailedToDeleteCloudAccount error
	ErrorFailedToDeleteCloudAccount = "Failed to delete cloud account with ID %s under team ID %s."

	//ErrorFailedToCreateComposite error
	ErrorFailedToCreateComposite = "Failed to create composite under team ID %s."

	//ErrorNoCompositesFound error
	ErrorNoCompositesFound = "No composites found under team ID %s."

	//ErrorNoCompositeWithIDFound error
	ErrorNoCompositeWithIDFound = "No composite with ID %s found under team ID %s."

	//ErrorNoGitKeysFound error
	ErrorNoGitKeysFound = "No git keys found under team ID %s."

	//ErrorNoGitKeyWithIDFound error
	ErrorNoGitKeyWithIDFound = "No git key with ID %s found under team ID %s."

	//ErrorFailedToCreateGitKey error
	ErrorFailedToCreateGitKey = "Failed to create git key under team ID %s."

	//ErrorFailedToDeleteGitKey error
	ErrorFailedToDeleteGitKey = "Failed to delete git key with ID %s under team ID %s."

	//ErrorNoPlansFound error
	ErrorNoPlansFound = "No plans found under team team ID %s and composite ID %s."

	//ErrorNoPlanWithIDFound error
	ErrorNoPlanWithIDFound = "No plan with ID %s found under team ID %s and composite ID %s."

	//ErrorFailedToDeletePlan error
	ErrorFailedToDeletePlan = "Failed to delete plan ID %s found under team ID %s and composite ID %s."

	//ErrorFailedToEnablePlan error
	ErrorFailedToEnablePlan = "Failed to enable plan ID %s found under team ID %s and composite ID %s."

	//ErrorFailedToDisblePlan error
	ErrorFailedToDisblePlan = "Failed to disable plan ID %s found under team ID %s and composite ID %s."

	//ErrorNoTokensFound error
	ErrorNoTokensFound = "No tokens found. To create a token use `coreo token add [flags]` command."

	//ErrorNoTokenWithIDFound error
	ErrorNoTokenWithIDFound = "No token with ID %s found."

	//ErrorFailedTokenCreation error
	ErrorFailedTokenCreation = "Failed to create token."

	//ErrorFailedToDeleteToken error
	ErrorFailedToDeleteToken = "Failed to delete token with ID %s."

	//ErrorNoTeamWithIDFound error
	ErrorNoTeamWithIDFound = "No team with ID %s found."
)
