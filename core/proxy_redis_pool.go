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

	akt "github.com/chatgpt-accesstoken"
	"github.com/chatgpt-accesstoken/store/redisdb"
)

type proxyService struct {
	db *redisdb.Redis
}

func NewProxyService(db *redisdb.Redis) akt.ProxyService {
	return &proxyService{db: db}
}

func (s *proxyService) List(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (s *proxyService) Add(ctx context.Context, ip string) error {
	return nil
}

func (s *proxyService) Delete(ctx context.Context, ip string) error {
	return nil
}
