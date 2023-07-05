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

package redisdb

import (
	"errors"
	"time"
)

type Config struct {
	Address      []string      `envconfig:"REDIS_ADDRESS"        default:"127.0.0.1:6379"`
	Password     string        `envconfig:"REDIS_PASSWORD"       default:"root"`
	DB           int           `envconfig:"REDIS_DB"             default:"0"`
	PoolSize     int           `envconfig:"REDIS_POOL_SIZE"      default:"20"`
	DailTimeOut  time.Duration `envconfig:"REDIS_DAIL_TIMEOUT"   default:"3s"`
	ReadTimeOut  time.Duration `envconfig:"REDIS_READ_TIMEOUT"   default:"5s"`
	WriteTimeout time.Duration `envconfig:"REDIS_WRITE_TIMEOUT"  default:"3s"`
}

func (c Config) Validate() error {
	if len(c.Address) == 0 {
		return errors.New("redis: cannot find address")
	}

	if c.Password == "" {
		return errors.New("redis: cannot find password")
	}
	return nil
}
