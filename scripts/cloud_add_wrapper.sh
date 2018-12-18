#!/bin/bash
# account name, account id, environment, profile
# To use this script, execute `sh cloud_add.sh < input_file` in Terminal
IFS=", "
team_id="YOUR_TEAM_ID"
role_name="securestate_role"
while read account_name account_id environment profile; do
	echo "creating role for account: $account_id"
	cloud_id=$(coreo cloud add --team-id $team_id --name $account_name --role $role_name --aws-profile $profile --environment $environment --json | jq -r .id)
	coreo event setup --team-id $team_id --cloud-id $cloud_id --aws-profile $profile
done