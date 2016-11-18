// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package content

const (
	//CmdPlanUse plan command
	CmdPlanUse = "plan"

	//CmdPlanShort short descripton
	CmdPlanShort = "Manage Coreo Plans"

	//CmdPlanLong long description
	CmdPlanLong = `Manage Coreo Plans.`

	//CmdPlanListShort short descripton
	CmdPlanListShort = "List all plans"

	//CmdPlanListLong long description
	CmdPlanListLong = `List all plans.`

	//CmdPlanInitShort short descripton
	CmdPlanInitShort = "Init a plan"

	//CmdPlanInitLong long description
	CmdPlanInitLong = `Init a plan.`

	//CmdPlanCreateShort short descripton
	CmdPlanCreateShort = "Create a plan"

	//CmdPlanCreateLong long description
	CmdPlanCreateLong = `Create a plan.`

	//CmdPlanDeleteShort short descripton
	CmdPlanDeleteShort = "Delete a plan"

	//CmdPlanDeleteLong long description
	CmdPlanDeleteLong = `Delete a plan.`

	//CmdPlanShowShort short descripton
	CmdPlanShowShort = "Show a plan"

	//CmdPlanShowLong long description
	CmdPlanShowLong = `Show a plan.`

	//CmdPlanDisableShort short descripton
	CmdPlanDisableShort = "Disable a plan"

	//CmdPlanDisableLong long description
	CmdPlanDisableLong = `Disable a plan.`

	//CmdPlanEnableShort short descripton
	CmdPlanEnableShort = "Enable a plan"

	//CmdPlanEnableLong long description
	CmdPlanEnableLong = `Enable a plan.`

	//CmdPlanRunShort short descripton
	CmdPlanRunShort = "Run a plan"

	//CmdPlanRunLong long description
	CmdPlanRunLong = `Run a plan.`

	//CmdFlagPlanIDLong flag
	CmdFlagPlanIDLong = "plan-id"

	//CmdFlagPlanIDDescription flag description
	CmdFlagPlanIDDescription = "Coreo plan id"

	//CmdFlagBranchLong flag
	CmdFlagBranchLong = "branch"

	//CmdFlagBranchDescription flag description
	CmdFlagBranchDescription = "Git branch for plan"

	//CmdFlagGitCommitIDLong commit id flag
	CmdFlagGitCommitIDLong = "gitcommit-id"

	//CmdFlagGitCommitIDDescription flag description
	CmdFlagGitCommitIDDescription = "Git commit id for branch"

	//CmdFlagCloudRegionLong cloud region flag
	CmdFlagCloudRegionLong = "region"

	//CmdFlagCloudRegionDescription flag description
	CmdFlagCloudRegionDescription = "Cloud region, e.g. AWS 'us-east-1'"

	//CmdFlagIntervalLong interval flag
	CmdFlagIntervalLong = "interval"

	//CmdFlagIntervalDescription flag description
	CmdFlagIntervalDescription = "Refresh rate"

	//InfoPlanDeleted information
	InfoPlanDeleted = "[Done] Plan was deleted.\n"

	//InfoUsingPlanID informational using Plan id
	InfoUsingPlanID = "[ OK ] Using Plan ID %s\n"

	//ErrorPlanIDRequired error message
	ErrorPlanIDRequired = "[ ERROR ] Plan ID is required for this command. Use flag '--plan-id'\n"
)
