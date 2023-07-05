/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this files except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package build

// BuildInfo represents the information about InfluxDB build.
type BuildInfo struct {
	Version string // Version is the current git tag with v prefix stripped
	Commit  string // Commit is the current git commit SHA
	Date    string // Date is the build date in RFC3339
}

var buildInfo BuildInfo

// SetBuildInfo sets the build information for the binary.
func SetBuildInfo(version, commit, date string) {
	buildInfo.Version = version
	buildInfo.Commit = commit
	buildInfo.Date = date
}

// GetBuildInfo returns the current build information for the binary.
func GetBuildInfo() BuildInfo {
	return buildInfo
}
