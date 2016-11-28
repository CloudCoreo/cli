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

package main

import (
	"github.com/CloudCoreo/cli/cmd"
)

var version = "No Version Provided"
var buildstamp = "Unknown buildstamp"
var githash = "Unknown githash"
var buildID = "Unknown buildID"

func main() {
	cmd.Version = version
	cmd.Buildstamp = buildstamp
	cmd.Githash = githash
	cmd.BuildID = buildID

	cmd.Execute()
}
