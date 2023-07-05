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

package launcher

import (
	"context"
	"os"
	"time"

	"github.com/chatgpt-accesstoken/store/redisdb"

	"github.com/chatgpt-accesstoken/signals"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"github.com/workpieces/log"
)

func NewAccessTokensCommand(ctx context.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "akt",
		Args:  cobra.NoArgs,
		Short: "start the access token service.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := Environ()
			if err != nil {
				return err
			}
			if err := cfg.Validate(); err != nil {
				return err
			}
			return cmdRunE(ctx, cfg)
		},
	}
	return rootCmd
}

func cmdRunE(ctx context.Context, o Config) error {
	l := NewLauncher()
	l.logger = log.New(log.ParseLevel(o.LogLevel))

	// Start the launcher and wait for it to exit on SIGINT or SIGTERM.
	if err := l.run(signals.WithStandardSignals(ctx), o); err != nil {
		return err
	}
	<-l.Done()

	// Tear down the launcher, allowing it a few seconds to finish any
	// in-progress requests.
	shutdownCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return l.Shutdown(shutdownCtx)
}

type Config struct {
	// LogLevel set environment variable log level.
	LogLevel string `envconfig:"LOGGER_LEVEL" default:"info"`
	// HttpBindAddress set the environment http address and port.
	HttpBindAddress string `envconfig:"HTTP_BIND_ADDRESS" default:":8080"`
	// UseLocalDB use local cache db.
	UseLocalDB bool `envconfig:"USE_LOCAL_DB" default:"true"`
	// ProxyFileName set the environment proxy filename.
	ProxyFileName string `envconfig:"PROXY_FILENAME"`
	// RedisDB set the environment variables of redis
	RedisDB redisdb.Config
}

// Environ returns the settings from the environment.
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}

func (c Config) Validate() error {
	if c.UseLocalDB {
		if c.ProxyFileName != "" {
			_, err := os.Stat(c.ProxyFileName)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return c.RedisDB.Validate()
}
