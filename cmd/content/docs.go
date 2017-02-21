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
	//CmdFlagTypeLong cmd
	CmdFlagTypeLong = "type"

	//CmdFlagTypeShort cmd
	CmdFlagTypeShort = "t"

	//CmdFlagTypeDescription cmd
	CmdFlagTypeDescription = "the type of documentation to generate (markdown, man, bash)"

	//CmdDocsUse docs command
	CmdDocsUse = "docs"

	//CmdDocsShort short description
	CmdDocsShort = "Generate documentation as markdown or man pages"

	//CmdDocsLong long description
	CmdDocsLong = `
Generate documentation files for Coreo.

This command can generate documentation for Coreo in the following formats:

- Markdown
- Man pages

It can also generate bash autocompletions.

	$ coreo docs markdown --dir mydocs/
`
	//ErrorDocGeneration description
	ErrorDocGeneration = "unknown doc type %q. Try 'markdown' or 'man'"
)
