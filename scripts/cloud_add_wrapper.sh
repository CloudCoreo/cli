#!/bin/bash
# input file format: account name,account id,environment,profile
# To use this script, execute `sh cloud_add.sh < input_file` in Terminal
# You may customize this script by modifying line 10, 23, 26, 27
# This script assumes account name is unique, so you can run it multiple times if adding any account fails.
IFS=","
team_id="YOUR_TEAM_ID"
role_name="securestate_role"
accounts=$(vss cloud list --json --team-id $team_id | jq -r .[].name)
while read account_name account_id environment profile; do
	echo "creating role for account: $account_id"
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
		cloud_id=$(vss cloud list --json --team-id $team_id | jq  -r '.[] | select(.name=="'$account_name'")|.id')
		vss event setup --team-id $team_id --cloud-id $cloud_id --aws-profile $profile
		continue
	fi
	cloud_id=$(vss cloud add --team-id $team_id --name $account_name --role $role_name --aws-profile $profile --environment $environment --json | jq -r .id)
	vss event setup --team-id $team_id --cloud-id $cloud_id --aws-profile $profile
done
unset IFS
