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

package render

import (
	"encoding/json"
	"fmt"
	"github.com/chatgpt-accesstoken/errors"
	"net/http"
)

// ErrorCode writes the json-encoded error message to the response.
func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, &errors.Error{Message: err.Error()}, status)
}

// InternalError writes the json-encoded error message to the response
// with a 500 internal server error.
func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 500)
}

// InternalErrorf writes the json-encoded error message to the response
// with a 500 internal server error.
func InternalErrorf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 500)
}

// NotImplemented writes the json-encoded error message to the
// response with a 501 not found status code.
func NotImplemented(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 501)
}

// NotFound writes the json-encoded error message to the response
// with a 404 not found status code.
func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 404)
}

// NotFoundf writes the json-encoded error message to the response
// with a 404 not found status code.
func NotFoundf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 404)
}

// Unauthorized writes the json-encoded error message to the response
// with a 401 unauthorized status code.
func Unauthorized(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 401)
}

// Forbidden writes the json-encoded error message to the response
// with a 403 forbidden status code.
func Forbidden(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 403)
}

// BadRequest writes the json-encoded error message to the response
// with a 400 bad request status code.
func BadRequest(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 400)
}

// BadRequestf writes the json-encoded error message to the response
// with a 400 bad request status code.
func BadRequestf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 400)
}

// JSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
