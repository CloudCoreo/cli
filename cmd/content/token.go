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
	//CmdTokenUse token command
	CmdTokenUse = "token"

	//CmdTokenShort short description
	CmdTokenShort = "Manage API Tokens"

	//CmdTokenLong long description
	CmdTokenLong = `Manage API Tokens.`

	//CmdTokenListShort short description
	CmdTokenListShort = "Show list of tokens"

	//CmdTokenListLong long description
	CmdTokenListLong = `Show list of tokens.`

	//CmdTokenAddShort short description
	CmdTokenAddShort = "Add a token"

	//CmdTokenAddLong long description
	CmdTokenAddLong = `Add a token.`

	//CmdTokenShowShort short description
	CmdTokenShowShort = "Show a token"

	//CmdTokenShowLong long description
	CmdTokenShowLong = `Show a token.`

	//CmdTokenDeleteShort short description
	CmdTokenDeleteShort = "Delete a token"

	//CmdTokenDeleteLong long description
	CmdTokenDeleteLong = `Delete a token.`

	//CmdFlagTokenIDLong flag
	CmdFlagTokenIDLong = "token-id"

	//CmdFlagTokenIDDescription flag description
	CmdFlagTokenIDDescription = "VMware Secure State token id"

	//InfoUsingTokenID informational using Token id
	InfoUsingTokenID = "[ OK ] Using Token ID %s\n"

	//InfoTokenDeleted information
	InfoTokenDeleted = "[ DONE ] Token was deleted.\n"

	//ErrorTokenIDMissing error
	ErrorTokenIDMissing = "Token ID is required for this command. Use flag --token-id\n"
)
