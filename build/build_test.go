/*
Copyright 2022 The Gridsum Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package build

import "testing"

func TestSetBuildInfo(t *testing.T) {
	type args struct {
		version      string
		commit       string
		date         string
		organization string
		email        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "build",
			args: args{
				version:      "v2.0.0",
				commit:       "hash",
				date:         "2023-01-02",
				organization: "Workpieces LLC",
				email:        "workpieces.app@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetBuildInfo(tt.args.version, tt.args.commit, tt.args.date, tt.args.organization, tt.args.email)
		})
	}
}
