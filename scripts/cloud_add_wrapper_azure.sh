#!/bin/bash
# account_name, application_id, key_value, subscription_id, environment
# To use this script, execute `sh cloud_add.sh < input_file` in Terminal
# This script assumes account name is unique, so you can run it multiple times if adding any account fails.
IFS=", "
team_id="YOUR_TEAM_ID"
directory_id="YOUR_DIRECTORY_ID"
accounts=$(vss cloud list --json | jq -r .[].name)
while read account_name application_id key_value subscription_id environment; do
	echo "Adding account: $subscription_id"
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
		cloud_id=$(vss cloud list --json | jq  -r '.[] | select(.name=="'$account_name'")|._id')
		vss event setup --cloud-id $cloud_id
		continue
	fi
	cloud_id=$(vss cloud add --name $account_name --application-id $application_id --key-value $key_value --subscription-id $subscription_id --directory-id $directory_id  --environment $environment --provider Azure --json | jq -r ._id)
	vss event setup --cloud-id $cloud_id
done
unset IFS
