// Copyright © 2018 Zechen Jiang <zechen@cloudcoreo.com>
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
	//CmdResultUse ...
	CmdResultUse = "result"

	//CmdResultShort ...
	CmdResultShort = "Show the violation results"

	//CmdResultLong ...
	CmdResultLong = "Show the violation results"

	//CmdResultRuleUse ...
	CmdResultRuleUse = "rule"

	//CmdResultRuleShort ...
	CmdResultRuleShort = "Show violating rules"

	//CmdResultRuleLong ...
	CmdResultRuleLong = "Show violating rules"

	//CmdResultObjectUse ...
	CmdResultObjectUse = "object"

	//CmdResultObjectShort ...
	CmdResultObjectShort = "Show violating objects"

	//CmdResultObjectLong ...
	CmdResultObjectLong = "Show violating objects"

	//CmdResultRuleExample ...
	CmdResultRuleExample = `  vss result rule --severity "High|Medium"
  vss result rule --severity "High|Low"
  vss result rule --cloud-id YOUR_SECURITY_STATE_CLOUD_ACCOUNT_ID --severity "Low"`

	//CmdResultObjectExample ...
	CmdResultObjectExample = `vss result object --severity "High|Medium"
  vss result object --severity "High|Low"
  vss result object --cloud-id YOUR_SECURITY_STATE_CLOUD_ACCOUNT_ID --severity "Low"`
)
