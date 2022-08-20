/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "This command publishes your prompts and challenges for others to enjoy!",
	Long: `This command publishes your work so others can enjoy!
	In order for this command to work you should be signed in with you account first. 
	
	example: wraw publish <path/to/file>
	
	By default, the command without arguments will publish a *.wraw.txt or *.wraw.md to our site.
	The file must be in the current working directory. The file must be either a markdown or txt file.
	If you are working on a challenge the file will be used as a response to that challenge.
	Otherwise, the publish command will be a response to the current prompt.
	You can demand this behavior explicitly by passing the --challenge/-c flag 
	or --prompt/-p flag for your given prompt/challenge id.

	example: wraw publish <path/to/file> -c <challengeid>
	example: wraw publish <path/to/file> -p <promptid>


	You can publish your responses on other sites that are supported if you provide your credentials.
	For additional information, please see the online documentation @
	https://www.github.com/Diogenesoftoronto/write-and-wood/src/wraw/README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publish called")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
