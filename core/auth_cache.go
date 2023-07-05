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
	"math/rand"
	"time"

	"github.com/workpieces/log"

	"github.com/acheong08/OpenAIAuth/auth"
	akt "github.com/chatgpt-accesstoken"
)

type openaiAuthCache struct {
	proxySvc akt.ProxyService
	svc      akt.OpenaiAuthService
	akStore  akt.AccessTokenStore
	logger   log.Logger
}

func (o openaiAuthCache) All(ctx context.Context, req *akt.OpenaiAuthRequest) (*auth.AuthResult, error) {
	res, err := o.akStore.Get(ctx, req.Email)
	if err == nil {
		if time.Now().Before(res.Expires) {
			o.logger.Info("api: token not expire")
			return res.AuthResult, nil
		} else {
			o.logger.Info("api: token has expire")
			goto LABEL
		}
	}

LABEL:
	if req.Proxy == "" {
		list, err := o.proxySvc.List(ctx)
		if err != nil {
			return nil, err
		}

		idx := rand.Intn(len(list))
		req.Proxy = list[idx]
	}
	resp, err := o.svc.All(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := o.akStore.Add(ctx, req.Email, &akt.AuthExpireResult{
		AuthResult: resp,
		Expires:    time.Now().Add(14 * 24 * time.Hour),
	}); err != nil {
		return nil, err
	}
	return resp, nil
}

func (o openaiAuthCache) AccessToken(ctx context.Context, req *akt.OpenaiAuthRequest) (*auth.AuthResult, error) {
	res, err := o.akStore.Get(ctx, req.Email)
	if err == nil {
		if time.Now().Before(res.Expires) {
			o.logger.Info("api: token not expire")
			return res.AuthResult, nil
		} else {
			o.logger.Info("api: token has expire")
			goto LABEL
		}
	}

LABEL:

	if req.Proxy == "" {
		list, err := o.proxySvc.List(ctx)
		if err != nil {
			return nil, err
		}

		idx := rand.Intn(len(list))
		req.Proxy = list[idx]
	}
	resp, err := o.svc.AccessToken(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := o.akStore.Add(ctx, req.Email, &akt.AuthExpireResult{
		AuthResult: resp,
		Expires:    time.Now().Add(14 * 24 * time.Hour),
	}); err != nil {
		return nil, err
	}
	return resp, nil
}

func (o openaiAuthCache) PUID(ctx context.Context, req *akt.OpenaiAuthRequest) (*auth.AuthResult, error) {
	res, err := o.akStore.Get(ctx, req.Email)
	if err == nil {
		if time.Now().Before(res.Expires) {
			o.logger.Info("api: token not expire")
			return res.AuthResult, nil
		} else {
			o.logger.Info("api: token has expire")
			goto LABEL
		}
	}

LABEL:

	if req.Proxy == "" {
		list, err := o.proxySvc.List(ctx)
		if err != nil {
			return nil, err
		}

		idx := rand.Intn(len(list))
		req.Proxy = list[idx]
	}
	resp, err := o.svc.PUID(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := o.akStore.Add(ctx, req.Email, &akt.AuthExpireResult{
		AuthResult: resp,
		Expires:    time.Now().Add(14 * 24 * time.Hour),
	}); err != nil {
		return nil, err
	}
	return resp, nil
}

func NewOpenaiAuthCache(proxySvc akt.ProxyService, svc akt.OpenaiAuthService, akStore akt.AccessTokenStore, logger log.Logger) akt.OpenaiAuthService {
	return &openaiAuthCache{
		proxySvc: proxySvc,
		svc:      svc,
		akStore:  akStore,
		logger:   logger.WithField("auth", "service"),
	}
}
