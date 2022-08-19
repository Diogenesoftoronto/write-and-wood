/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "This command logs you in to your account.",
	Long: `This command will help you log you in to your account.
	
	To log you in you need to provide your username and password.
	
	example: wraw login -u <username> -p <password>
	
	You can also log you in with your email address and password.
	
	example: wraw login -e <email> -p <password>
	
	You can also login using a secret token and userNAME
	
	example: wraw login -t <token> -u <username>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
