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

package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/chatgpt-accesstoken/cmd/launcher"

	"github.com/chatgpt-accesstoken/build"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = ""
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	if len(date) == 0 {
		date = time.Now().UTC().Format(time.RFC3339)
	}

	build.SetBuildInfo(version, commit, date)
	ctx := context.Background()
	rootCmd := NewCommand()
	rootCmd.AddCommand(launcher.NewAccessTokensCommand(ctx))
	rootCmd.SilenceUsage = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "akt",
		Args:  cobra.NoArgs,
		Short: "Command line tool to provide access token service.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErrf("See '%s -h' for help\n", cmd.CommandPath())
		},
	}
	return rootCmd
}
