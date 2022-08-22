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
	
	This will show the amount of views you have, and other engagment stats about you and your content.
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
		 You can also use the --prompt/-p or --challenge/-c flags to show the view metrics for a specific prompt or challenge. 
		- "--words": This shows you just the words metrics. Like the engagment metrics, this can be set with the --time/-t flag.
			You can observe how many words you have written overtime to see if you are increasing or not. 
		- "commits": This shows you how many changes or additions you have made.
			Those familiar with programming the system is very similar to git. You can set the time as usual with the --time/-t flag.
	
	You can combine status flags to create a more statistics.
	
	wraw status -e -v -w
	
	this example status will show you how many words you have written, your engagment metrics, and the views.
		
	If you are not logged in this command will not work.
	
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
