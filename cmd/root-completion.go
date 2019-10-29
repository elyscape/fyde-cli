// Package cmd implements fyde-cli commands
package cmd

/*
Copyright © 2019 Fyde, Inc. <hello@fyde.com>

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

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash, zsh or PowerShell completion scripts",
}

// bashCompletionCmd represents the completion command for bash
var bashCompletionCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generates bash completion script",
	Long: `To load completion run

. <(fyde-cli completion bash)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(fyde-cli completion bash)`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

// zshCompletionCmd represents the completion command for zsh
var zshCompletionCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generates zsh completion script",
	Long: `The completion script can be generated by running

fyde-cli completion zsh
	
The generated completion script should be put somewhere in your $fpath named _fyde-cli.
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}

// psCompletionCmd represents the completion command for zsh
var psCompletionCmd = &cobra.Command{
	Use:     "powershell",
	Aliases: []string{"ps"},
	Short:   "Generates PowerShell completion script",
	Long: `The completion script can be generated by running

fyde-cli completion ps
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenPowerShellCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(bashCompletionCmd)
	completionCmd.AddCommand(zshCompletionCmd)
	completionCmd.AddCommand(psCompletionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
