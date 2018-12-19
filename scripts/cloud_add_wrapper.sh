#!/bin/bash
# account name, account id, environment, profile
# To use this script, execute `sh cloud_add.sh < input_file` in Terminal
IFS=", "
team_id="YOUR_TEAM_ID"
role_name="securestate_role"
while read account_name account_id environment profile; do
	echo "creating role for account: $account_id"
	accounts=$(coreo cloud list --json --team-id $team_id | jq -r .[].name)
	is_exist=false
	while read account
	do
		if [ "$account" = "$account_name" ]; then
			is_exist=true
			break
		fi
	done <<< "$accounts"
	if [ $is_exist = true ]; then
		echo "Cloud account with name $account_name already exist, skip for this account"
		continue
	fi
	cloud_id=$(coreo cloud add --team-id $team_id --name $account_name --role $role_name --aws-profile $profile --environment $environment --json | jq -r .id)
	coreo event setup --team-id $team_id --cloud-id $cloud_id --aws-profile $profile
done
unset IFS