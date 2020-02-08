/*
Copyright © 2020 Luca Zecca <l.zecca78@gmail.com>

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
package cmd

import (
	"github.com/lzecca78/git-cloner/config"
	"github.com/lzecca78/git-cloner/internal"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "this will clone defined in config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		missing, _ := cmd.Flags().GetBool("missing")
		cfg := config.GetConfig(cfgFile)
		if missing {
			internal.Clone(cfg, missing)
		} else {
			internal.Clone(cfg)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolP("missing", "m", false, "clone only not existing repos")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
