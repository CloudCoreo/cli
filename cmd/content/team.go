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
	//CmdTeamUse team command
	CmdTeamUse = "team"

	//CmdTeamShort short description
	CmdTeamShort = "Manage Teams"

	//CmdTeamLong long descripton
	CmdTeamLong = `Manage Teams.`

	//CmdTeamListShort short description
	CmdTeamListShort = "Show list of team"

	//CmdTeamListLong long descripton
	CmdTeamListLong = `Show list of team.`

	//CmdTeamShowShort short description
	CmdTeamShowShort = "Show a team"

	//CmdTeamShowLong long descripton
	CmdTeamShowLong = `Show a team.`

	//InfoUsingTeamID informational using team id
	InfoUsingTeamID = "[ OK ] Using Team ID %s\n"

	//ErrorTeamIDMissing error
	ErrorTeamIDMissing = "[ ERROR ] Team ID is required for this command. Use flag --team-id\n"
)
