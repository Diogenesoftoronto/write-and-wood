package util

import (
	"os"
	"log"
	"strings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func index() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	USERNAME := os.Getenv("USER_NAME")
	ACCESS_KEY := os.Getenv("ACCESS_KEY_ID")
	CONSOLE_LOGIN_LINK := os.Getenv("CONSOLE_LOGIN_LINK")
	SECRET := os.Getenv("SECRET")

}