#!/bin/bash
# input file format: account name,account id,environment,profile
# To use this script, execute `sh cloud_add_wrapper_aws.sh < input_file` in Terminal
# You may customize this script by modifying line 10, 23, 26, 27
# This script assumes account name is unique, so you can run it multiple times if adding any account fails.
IFS=","
role_name="securestate_role"
accounts=$(vss cloud list --json | jq -r .[].name)
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
		cloud_id=$(vss cloud list --json | jq  -r '.[] | select(.name=="'$account_name'")|._id')
		vss event setup --cloud-id $cloud_id --aws-profile $profile --ignore-missing-trails
		continue
	fi
	cloud_id=$(vss cloud add --name $account_name --role $role_name --aws-profile $profile --environment $environment --json | jq -r ._id)
	vss event setup --cloud-id $cloud_id --aws-profile $profile --ignore-missing-trails
done
unset IFS
