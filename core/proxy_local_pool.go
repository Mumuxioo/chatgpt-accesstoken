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
	akt "github.com/chatgpt-accesstoken"
	"sync"
)

type proxyLocalService struct {
	db   map[string]bool
	lock sync.RWMutex
}

func NewProxyLocalService() akt.ProxyService {
	return &proxyLocalService{
		db: make(map[string]bool),
	}
}

func (s *proxyLocalService) List(ctx context.Context) ([]string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if len(s.db) == 0 {
		return nil, fmt.Errorf("local: cannot find proxy")
	}

	list := make([]string, 0, len(s.db))
	for k, _ := range s.db {
		list = append(list, k)
	}
	return list, nil
}

func (s *proxyLocalService) Add(ctx context.Context, ip string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.db[ip] = true
	return nil
}

func (s *proxyLocalService) Delete(ctx context.Context, ip string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.db[ip]
	if ok {
		delete(s.db, ip)
		return nil
	}
	return fmt.Errorf("local: %s don't exist", ip)
}
