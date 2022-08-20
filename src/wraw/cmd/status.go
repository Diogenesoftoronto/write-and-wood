/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status [flag] [arg]",
	Short: "This command shows the status of the user",
	Long: `This command shows the status of the user.
	
	example: wraw status
	
	this will show the amount of views you have, and other engagment stats about you and your content.
	Understanding how other people engage with your content is very important to writers.
	The status command gives you information about engagment but it also provides information about your writing.
	It will tell you the number of challenges you have completed, the number of prompts created, 
	your total word count and the rank your content has compared to others based on these metrics.

	Always remember that the rank or metrics you have does not mean you are a bad or good writer, merely what we can measure.
	Beauty is in the eye of the beholder.

	There are several status flags available: 
	
		- "--engagment/-e":  This flag displays your current engagment metrics. 
			Default is to display your current metrics. But this can be set with the --time/-t flag. 
			Passing a date will display the metrics overtime between then in now.
		- "--views/-v": This flag shows you just the view metrics. Like the engagment metrics, this can be set with the --time/-t flag. 
		 You can also use the --prompt/-p or --challenge/-c flags to show the view metrics for a specific prompt or challenge., 
		 or that must or must not contain certain words.
		- "characters": This is a character-based challenge. 
			You will be challenged to write a prompt that contains or does not contain a given list of characters.
		- "regex": This is a regular expression challenge.
			You will be challenged to write a prompt that matches a given regular expression. The expression must be set in 
			the -e flag. Otherwise, the challenge will be invalid. regex match in gosyntax.
		- "random": This is a random challenge.
			You will be challenged to write a prompt with a random set of complex status flags. 
			The complexity can be set with the -rc flag.
	status flags can be set to be private or public. 
	Private status flags are only visible to you and the player you are challenging.
	
	You can combine status flags to create a more complex challenge.
	
	wraw challenge time words characters -ti 30 -wr "hello, world" -ch "w" -cc 10
	
	this example challenge will challenge you to write a prompt that must be written in 30 minutes,
	must contain hello, and world, and must contain  the character "w" 10 times.
		
	If you are not logged in this command will not work.
	no arguments are required.
	
	For additional information, please see the online documentation @
	https://www.github.com/Diogenesoftoronto/write-and-wood/src/wraw/README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	challengeCmd.PersistentFlags().String("time", "ti", "The time in ")
challengeCmd.PersistentFlags().String("words", "wr", "The words that the challenge must contain")
	challengeCmd.PersistentFlags().String("characters", "ch", "The characters that the challenge must contain")
	challengeCmd.PersistentFlags().String("regex", "re", "The regular expression that the challenge must match")
	challengeCmd.PersistentFlags().String("random-complexity", "rc", "The number of random challenges to be included in the challenge")
challengeCmd.PersistentFlags().String("private", "pr", "The challenge will be private")
	challengeCmd.PersistentFlags().String("charcount", "cc", "The number of characters that the challenge must contain")


}
