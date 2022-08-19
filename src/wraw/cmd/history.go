/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "This command shows the history of your prompts, challenges and more.",
	Long: `This command will help you quickly get a list of your last answered prompts, challenges, 
	and more. 
	
	To get a list of your last answered prompts simply use the history command without any arguments.

	example: wraw history

	To get a list of your last answered prompts and challenges you can use the history command with the -c/--challenges flag for challenges,
	and the -p/--prompts flag for prompts. 
	
	example: wraw history -c -p
	This example returns the last five answered challenges and prompts.

	Adding -a/--amount gives you the length of the history. 
	
	To got the total history, you can use the --all flag.
	
	example: wraw history --all
	This example returns the total history.

	You can also get the total history of prompts and challenges separately by passing the --all-prompts flag and the --all-challenges flag.

	example: wraw history --all-challenges
	This example returns the total history of challenges.

	example: wraw history --all-prompts
	This example returns the total history of prompts.
	
	In order for this command to work you should be signed in with you account first.
		For additional information visit:
		https://www.github.com/Diogenesoftoronto/write-and-wood/README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("history called")
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
