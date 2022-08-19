/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// challengeCmd represents the challenge command
var challengeCmd = &cobra.Command{
	Use:   "challenge [args] [flags]",
	Short: "This command will allow you to challenge yourself and other players",
	Long: `This command works by creating a challenge. By default, you will be challenged by a random player that is also 
	queued for a challenge, or someone who has challenged you. 
	example: 
	
	wraw challenge
	
	this will challenge a random player with a default timed challenge.
	
	wraw challenge -p <player>
	
	this will challenge a specific player. If the player is not accepting challenges or you are blocked by them
	this command will fail.
	
	wraw challenge -s <challenge>
	
	this command will challenge yourself to a specific challenge.
	There are several challenges available: 
	
		- "time":  This is a time-based challenge. You will be challenged to write a prompt in a given amount of time. 
			Default time is 15 minutes. But this can be set with the -t flag.
		- "words": This is a word-based challenge. 
		 will be challenged to write a prompt that contains a given number of words, 
		 or that must or must not contain certain words.
		- "characters": This is a character-based challenge. 
			You will be challenged to write a prompt that contains or does not contain a given list of characters.
		- "regex": This is a regular expression challenge.
			You will be challenged to write a prompt that matches a given regular expression. The expression must be set in 
			the -e flag. Otherwise, the challenge will be invalid. regex match in gosyntax.
		- "random": This is a random challenge.
			You will be challenged to write a prompt with a random set of complex challenges. 
			The complexity can be set with the -rc flag.
	Challenges can be set to be private or public. 
	Private challenges are only visible to you and the player you are challenging.
	
	You can combine challenges to create a more complex challenge.
	
	wraw challenge time words characters -ti 30 -wr "hello, world" -ch "w" -cc 10
	
	this example challenge will challenge you to write a prompt that must be written in 30 minutes,
	must contain hello, and world, and must contain  the character "w" 10 times.
	
	For more additional information visit:
	https://www.github.com/Diogenesoftoronto/write-and-wood/README.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("challenge called")
	},
}

func init() {
	rootCmd.AddCommand(challengeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	challengeCmd.PersistentFlags().String("time", "ti", "The time in minutes that the challenge must be completed in")
	challengeCmd.PersistentFlags().String("words", "wr", "The words that the challenge must contain")
	challengeCmd.PersistentFlags().String("characters", "ch", "The characters that the challenge must contain")
	challengeCmd.PersistentFlags().String("regex", "re", "The regular expression that the challenge must match")
	challengeCmd.PersistentFlags().String("random-complexity", "rc", "The number of random challenges to be included in the challenge")
	challengeCmd.PersistentFlags().String("private", "pr", "The challenge will be private")
	challengeCmd.PersistentFlags().String("charcount", "cc", "The number of characters that the challenge must contain")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// challengeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
