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

	//CmdTeamLong long description
	CmdTeamLong = `Manage Teams.`

	//CmdTeamListShort short description
	CmdTeamListShort = "Show list of team"

	//CmdTeamListLong long description
	CmdTeamListLong = `Show list of team.`

	//CmdTeamListExample is the use case for command team list
	CmdTeamListExample = "coreo team list"

	//CmdTeamShowShort short description
	CmdTeamShowShort = "Show a team"

	//CmdTeamShowLong long description
	CmdTeamShowLong = `Show a team.`

	//CmdTeamAddShort add a team short description
	CmdTeamAddShort = "Add a team"

	//CmdTeamAddLong add a team long description
	CmdTeamAddLong = `Add a team.`

	//CmdTeamNameDescription provide team name
	CmdTeamNameDescription = "Provide team name"

	//CmdTeamDescriptionDescription provide team description
	CmdTeamDescriptionDescription = "Provide team description"

	//InfoUsingTeamID informational using team id
	InfoUsingTeamID = "[ OK ] Using Team ID %s\n"

	//ErrorTeamIDMissing error
	ErrorTeamIDMissing = "Team ID is required for this command. Use flag --team-id\n"

	//ErrorTeamNameRequired error
	ErrorTeamNameRequired = "Team Name is required for this command. Use flag --name\n"

	//ErrorTeamDescriptionRequired error
	ErrorTeamDescriptionRequired = "Team description required for this command. Use flag --description\n"
)
