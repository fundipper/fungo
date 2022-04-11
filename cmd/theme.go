/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/fundipper/fungo/internal/cli"
	"github.com/spf13/cobra"
)

// themeCmd represents the theme command
var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "create a new site with default content.",
	Long:  `create a new theme with default content.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := "create new theme success"
		err := cli.NewTheme(args[0]).Create()
		if err != nil {
			message = err.Error()
		}
		fmt.Println(message)
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(themeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// themeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// themeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
