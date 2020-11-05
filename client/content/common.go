package content

const (
	//ErrorMissingAPIOrSecretKey error
	ErrorMissingAPIOrSecretKey = "Missing API key or/and Secret key. Please run 'coreo configure' to configure them."

	//ErrorNoCloudAccountsFound error
	ErrorNoCloudAccountsFound = "No cloud accounts found."

	//ErrorNoCloudAccountWithIDFound error
	ErrorNoCloudAccountWithIDFound = "No cloud account with ID %s found."

	//ErrorFailedToCreateCloudAccount error
	ErrorFailedToCreateCloudAccount = "Failed to create cloud account."

	//ErrorFailedToDeleteCloudAccount error
	ErrorFailedToDeleteCloudAccount = "Failed to delete cloud account with ID %s under team ID %s."

	//ErrorMissingRoleInformation error
	ErrorMissingRoleInformation = "Adding cloud account failed, you need to provide either rolearn and external id or new role name"

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
