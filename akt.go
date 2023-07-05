/*
Copyright 2022 The Workpieces LLC.

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

package chatgpt_accesstoken

import (
	"context"
	"time"

	"github.com/acheong08/OpenAIAuth/auth"
)

type OpenaiAuthRequest struct {
	Email       string `json:"email"`                  // Email Openai chatgpt email.
	Password    string `json:"password"`               // Password Openai chatgpt password.
	Proxy       string `json:"proxy,omitempty"`        // Proxy global proxy default: http://username:password@ip:port
	MFA         string `json:"mfa,omitempty"`          // MFA don't need input.
	AccessToken string `json:"access_token,omitempty"` //  AccessToken used to get puid
}

type OpenaiAuthService interface {
	// All Get access_token and puid, if one of them reports an error, return an exception.
	All(ctx context.Context, req *OpenaiAuthRequest) (*auth.AuthResult, error)
	// AccessToken just getting access_token will not trigger getting puid.
	AccessToken(ctx context.Context, req *OpenaiAuthRequest) (*auth.AuthResult, error)
	// PUID just getting the PUID will not trigger getting the access token.
	PUID(ctx context.Context, req *OpenaiAuthRequest) (*auth.AuthResult, error)
}

type ProxyService interface {
	// List get proxy list.
	List(ctx context.Context) ([]string, error)
	// Add insert proxy ip into redis db.
	Add(ctx context.Context, ip string) error
	// Delete proxy ip from redis db.
	Delete(ctx context.Context, ip string) error
}

type AuthExpireResult struct {
	*auth.AuthResult
	Expires time.Time `json:"expires"`
}

type AccessTokenStore interface {
	Add(ctx context.Context, email string, ak *AuthExpireResult) error
	Delete(ctx context.Context, email string) error
	Get(ctx context.Context, email string) (*AuthExpireResult, error)
}
