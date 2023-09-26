/*
 *
 * Copyright 2023 cornucopia authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cmd

import (
	"fmt"

	"github.com/dvaumoron/cornucopia/common"
	"github.com/dvaumoron/cornucopia/config"
	"github.com/dvaumoron/cornucopia/glu"
	"github.com/dvaumoron/cornucopia/module"
	"github.com/spf13/cobra"
	"go.starlark.net/starlark"
)

var conf config.Config

func Init() *cobra.Command {
	defaultRepoPath, defaultRepoUrl, err := config.InitDefault("CORNUCOPIA_REPO_PATH", "CORNUCOPIA_REPO_URL")

	cmd := &cobra.Command{
		Use:   "cornucopia scriptname",
		Short: "cornucopia run a code generation script.",
		Long: `cornucopia run a code generation script, find more details at :
https://github.com/dvaumoron/cornucopia`,
		Args: cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			if err != nil {
				return err
			}

			scriptname := args[0]
			loader := module.MakeLoader(common.Prefix, &conf)
			thread := &starlark.Thread{Name: common.Prefix + scriptname, Load: loader.Load}
			glu.InitCornucopiaGlobals()

			if _, err = starlark.ExecFile(thread, scriptname, nil, nil); err != nil {
				if len(glu.GeneratedFilenames) != 0 {
					displayGeneratedFileNames("Before error, the following file have been generated :")
				}
				return err
			}

			if len(glu.GeneratedFilenames) == 0 {
				fmt.Println("Successful without file generation")
			} else {
				displayGeneratedFileNames("Successful, the following file have been generated :")
			}
			return nil
		},
	}

	cmdFlags := cmd.Flags()
	cmdFlags.StringVar(&conf.RepoPath, "repo-path", defaultRepoPath, "Local path of the script cache repository")
	cmdFlags.StringVar(&conf.RepoUrl, "repo-addr", defaultRepoUrl, "Address of the shared script repository")
	cmdFlags.BoolVarP(&conf.ForceDownload, "force-download", "f", false, "Force download in module resolution")

	return cmd
}

func displayGeneratedFileNames(msg string) {
	fmt.Println(msg)
	for _, fileName := range glu.GeneratedFilenames {
		fmt.Println(fileName)
	}
}
