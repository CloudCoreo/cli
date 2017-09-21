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
	//CmdDevTimeKeyUse gitKey command
	CmdDevTimeKeyUse = "devtime"

	//CmdDevTimeShot short description
	CmdDevTimeShot = "Manage devtime"

	//CmdDevTimeLong long description
	CmdDevTimeLong = `Manage or create devtime.`

	//CmdDevTimeCreateShort short description
	CmdDevTimeCreateShort = "Create a devtime url"

	//CmdDevTimeCreateLong long description
	CmdDevTimeCreateLong = `Create a devtime url for a given context and task.`

	//CmdDevTimeStartShort short description
	CmdDevTimeStartShort = "Start a devtime"

	//CmdDevTimeStartLong long description
	CmdDevTimeStartLong = `Start a devtime for a given devtime ID.`

	//CmdDevTimeStopShort short description
	CmdDevTimeStopShort = "Stop a devtime"

	//CmdDevTimeStopLong long description
	CmdDevTimeStopLong = `Stop a devtime for a given devtime ID.`

	//CmdDevTimeResultsShort short description
	CmdDevTimeResultsShort = "Get devtime results"

	//CmdDevTimeResultsLong long description
	CmdDevTimeResultsLong = `Get devtime results for a given devtime ID.`

	//CmdDevTimeJobsShort short description
	CmdDevTimeJobsShort = "Get running devtime jobs"

	//CmdDevTimeJobsLong long description
	CmdDevTimeJobsLong = `Get running devtime jobs for a given devtime ID.`

	// CmdFlagDevTimeContextLong flag
	CmdFlagDevTimeContextLong = "context"

	// CmdFlagDevTimeTaskLong flag
	CmdFlagDevTimeTaskLong = "task"

	// CmdFlagDevTimeIDLong flag
	CmdFlagDevTimeIDLong = "devtime-id"

	//ErrorDevTimeIDMissing error
	ErrorDevTimeIDMissing = "DevTimeID is required for this command. Use flag --devtime-id."

	//ErrorContextMissing error
	ErrorContextMissing = "Context is required for this command. Use flag --context."

	//ErrorTaskMissing error
	ErrorTaskMissing = "Task is required for this command. Use flag --task."

	//CmdFlagDevTimeContextDescription Devtime context
	CmdFlagDevTimeContextDescription = "Devtime context"

	//CmdFlagDevTimeTaskDescription Devtime task
	CmdFlagDevTimeTaskDescription = "Devtime task"

	//CmdFlagDevTimeIDDescription Devtime ID
	CmdFlagDevTimeIDDescription = "Devtime ID"

	//InfoDevTimeStarted informational using DevTime id
	InfoDevTimeStarted = "Started DevTime for ID %s.\n"

	//InfoDevTimeStopped informational using DevTime id
	InfoDevTimeStopped = "Stopped DevTime for ID %s.\n"
)
