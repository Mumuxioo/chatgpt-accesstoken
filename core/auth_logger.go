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
	"context"
	"fmt"
	"time"

	"github.com/acheong08/OpenAIAuth/auth"
	akt "github.com/chatgpt-accesstoken"
	"github.com/workpieces/log"
)

type openaiAuthLogger struct {
	logger log.Logger
	svc    akt.OpenaiAuthService
}

func NewOpenaiAuthLogger(logger log.Logger, svc akt.OpenaiAuthService) akt.OpenaiAuthService {
	return &openaiAuthLogger{logger: logger.WithField("openai", "service"), svc: svc}
}

func (o openaiAuthLogger) All(ctx context.Context, req *akt.OpenaiAuthRequest) (resp *auth.AuthResult, err error) {
	defer func(start time.Time) {
		olog := o.logger.WithField("took", time.Since(start))
		if err != nil {
			olog.Error(fmt.Sprintf("api: cannot find all: %s", err))
			return
		}
		olog.Info("all")
	}(time.Now())
	return o.svc.All(ctx, req)
}

func (o openaiAuthLogger) AccessToken(ctx context.Context, req *akt.OpenaiAuthRequest) (resp *auth.AuthResult, err error) {
	defer func(start time.Time) {
		olog := o.logger.WithField("took", time.Since(start))
		if err != nil {
			olog.Error(fmt.Sprintf("api: cannot find access_token: %s", err))
			return
		}
		olog.Info("access_token")
	}(time.Now())

	return o.svc.AccessToken(ctx, req)
}

func (o openaiAuthLogger) PUID(ctx context.Context, req *akt.OpenaiAuthRequest) (resp *auth.AuthResult, err error) {
	defer func(start time.Time) {
		olog := o.logger.WithField("took", time.Since(start))
		if err != nil {
			olog.Error(fmt.Sprintf("api: cannot find puid: %s", err))
			return
		}
		olog.Info("puid")
	}(time.Now())

	return o.svc.PUID(ctx, req)
}
