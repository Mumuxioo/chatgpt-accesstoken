/*
Copyright 2022 The deepauto-io LLC.

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

package core

import (
	"fmt"
	"strings"

	"github.com/acheong08/OpenAIAuth/auth"
)

type OError struct {
	Err *auth.Error
}

func (o OError) Error() string {
	var errs []string

	if o.Err.Location != "" {
		errs = append(errs, fmt.Sprintf("localtion: %s", o.Err.Location))
	}

	if o.Err.StatusCode != 0 {
		errs = append(errs, fmt.Sprintf("status code: %d", o.Err.StatusCode))
	}

	if o.Err.Details != "" {
		errs = append(errs, fmt.Sprintf("details: %s", o.Err.Details))
	}

	if o.Err.Error != nil {
		errs = append(errs, fmt.Sprintf("err: %s", o.Err.Error.Error()))
	}

	return strings.Join(errs, ",")
}
