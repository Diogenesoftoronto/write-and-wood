/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "This command helps configure this tool.",
	Long: `This command is best used when you first create a new account or begin using the tool for the first time.
	Running the config command with an init argument will create a new config file and start the process of 
	setting up the tool. If you have ever used tools like npm or yarn before its a lot like those.		
	
	To configure your account you will need to use the --username and --password flags.
			
			example: wraw config --username <username> --password <password>
			
			You can also use the --token flag to configure your account with a token.
			
			example: wraw config --token <token>
			
			In order for this command to work you should be signed in with you account first.
			
			You can change the default settings for different tooling, changing almost everything about the standard behavior and adding aliasing for commands.
			This speeds up your workflow and helps make sure you dont have to constantly type the same commands.
			
			example: wraw config --default-prompt <prompt> --default-challenge <challenge> --default-alias <alias>
			
			Note: You can always manually change these settings in your wraw.yaml file.
			
			You can change your publishing configuration and where you are publishing when you use that function, 
			This allows you to publish to multiple locations at once!

			example: wraw config --publish <location> --alias <alias>
			example: wraw config --default-pub <location> <location>
			
			You can offer multiple locations by continuously adding locations as seen in the above example.

			Note: The Publisher must be in our supported list for it to be guaranteeded to work.
			
			For more additional information visit:
			https://www.github.com/Diogenesoftoronto/write-and-wood/README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
