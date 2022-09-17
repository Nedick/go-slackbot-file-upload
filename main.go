package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	SLACK_BOT_TOKEN := goDotEnvVariable("SLACK_BOT_TOKEN")
	SLACK_CHANNEL_ID := goDotEnvVariable("SLACK_CHANNEL_ID")
	api := slack.New(SLACK_BOT_TOKEN)
	channelArr := []string{SLACK_CHANNEL_ID}
	fileArr := []string{"511224.jpg"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
