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
	//CmdLintUse command
	CmdLintUse = "coreo lint [flags] PATH"

	//CmdLintShort short description
	CmdLintShort = "examines a composite for possible issues"

	//CmdLintLong long description
	CmdLintLong = `This command takes a path to a composite and runs a series of tests to verify that
  the composite is well-formed.

  If the linter encounters things that will cause the composite to fail installation,
  it will emit [ERROR] messages. If it encounters issues that break with convention
  or recommendation, it will emit [WARNING] messages.
  `
)
