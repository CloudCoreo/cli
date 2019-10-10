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
	CmdTeamListExample = "vss team list"

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

	ErrorProviderNotSupported = "Provider not supported, either input AWS or Azure\n"
)
