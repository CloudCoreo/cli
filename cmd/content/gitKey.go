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
	//CmdGitKeyUse gitKey command
	CmdGitKeyUse = "git-key"

	//CmdGitKeyShort short description
	CmdGitKeyShort = "Manage Git keys"

	//CmdGitKeyLong long description
	CmdGitKeyLong = `Manage or Add Git keys to connect private repos to your CloudCoreo account. This will allow you to create private Composites and Plans.`

	//CmdGitKeyListShort short description
	CmdGitKeyListShort = "Show Git keys associated with your Coreo team"

	//CmdGitKeyListLong long description
	CmdGitKeyListLong = `Show Git keys associated with your Coreo team.`

	//CmdGitKeyAddShort short description
	CmdGitKeyAddShort = "Add a Git key"

	//CmdGitKeyAddLong long description
	CmdGitKeyAddLong = `Add Git keys to connect private repos to your CloudCoreo account. This will allow you to create private Composites and Plans.`

	//CmdGitKeyShowShort short description
	CmdGitKeyShowShort = "Show a Git key"

	//CmdGitKeyShowLong long description
	CmdGitKeyShowLong = `Show a Git key.`

	//CmdGitKeyDeleteShort delete command
	CmdGitKeyDeleteShort = "Delete a Git key"

	//CmdGitKeyDeleteLong long description
	CmdGitKeyDeleteLong = `Delete a Git key.`

	// CmdFlagGitKeyIDLong flag
	CmdFlagGitKeyIDLong = "gitKey-id"

	//CmdFlagGitKeyIDDescription flag description
	CmdFlagGitKeyIDDescription = "Coreo Git key ID"

	//ErrorGitKeyIDMissing error
	ErrorGitKeyIDMissing = "[ ERROR ] Git Key ID is required for this command. Use flag --gitKey-id\n"

	//InfoUsingGitKeyID informational using git key id
	InfoUsingGitKeyID = "[ OK ] Using Git Key ID %s\n"
)
