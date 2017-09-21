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

	//ErrorFailedToDeleteComposite error
	ErrorFailedToDeleteComposite = "Failed to delete composite ID %s found under team ID %s."

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

	//InfoPlanCreationMessage info
	InfoPlanCreationMessage = "Initializing plan and generating planconfig json file, please wait..."

	//InfoPlanCompilingMessage info
	InfoPlanCompilingMessage = "Compiling plan, please wait..."

	//InfoPlanRunNowMessageBlock info
	InfoPlanRunNowMessageBlock = "Adding plan to queue and executing it, please wait..."

	//InfoPlanRunNowMessage info
	InfoPlanRunNowMessage = "Adding plan to queue and executing it.\n"

	//ErrorPlanConfigVaribaleMissing error
	ErrorPlanConfigVaribaleMissing = "Error in plan config, missing required values."

	//ErrorPlanConfigRequiredVariableMissing error
	ErrorPlanConfigRequiredVariableMissing = "Missing plan config value for key %s.\n"

	//ErrorPlanIntervalMintuesInvalid error
	ErrorPlanIntervalMintuesInvalid = "Interval value should be equal to or more than 2 minutes and less than or equal to 525600 minutes."

	//ErrorPlanCreation error
	ErrorPlanCreation = "Something went wrong when adding a plan. More info -> %q"

	//ErrorPlanRunNow error
	ErrorPlanRunNow = "Something went wrong when running this plan. More info -> %q"

	//ErrorPlanCompileNow error
	ErrorPlanCompileNow = "Something went wrong compiling this plan. More info -> %q"

	//ErrorFailedToCreateDevTime error
	ErrorFailedToCreateDevTime = "Failed to create devTime under team ID %s."

	//ErrorNoDevTimesFound error
	ErrorNoDevTimesFound = "No devTimes found under team ID %s."

	//ErrorNoDevTimeWithIDFound error
	ErrorNoDevTimeWithIDFound = "No devTime with ID %s found under team ID %s."
)
