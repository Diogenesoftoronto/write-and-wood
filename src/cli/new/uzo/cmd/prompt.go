/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// ideaCmd represents the idea command
var promptCmd = &cobra.Command{
	Use:   "prompt [args]",
	Short: "This command will respond with a prompt of your choice",
	Long: `This command will help you manage getting and answering prompts.
			To get a prompt simply use the prompt command without any arguments.
			In order for this command to work you should be signed in with you account first. 
			For additional information visit: 
			https://www.github.com/Diogenesoftoronto/write-and-wood/README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("here is a prompt")
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	promptCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "A File name to unzip and open in IDE")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ideaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
