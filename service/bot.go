package service

import (
	"log"
	"os"
	"simplebot/config"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func InitBot(configPath string) {
	/*_bot, err := linebot.New(
		"9b6efdc8313e48423db82865c92a1ea7", //channel secret
		"H3VESfAiXw1cX9inFm7RaVEUpy+rJEbeMEcdh6IIrJVQy5UVjDjPeJ9JGih/MFJeXRDy+ZPLUFFcbp/rfk9F2pO2lrWc+QI747zuB9ZB7ptkLf4FtpvcGmtRuG1hY6mzPgfzpxsl+huz1Ztfrq9NmY9PbdgDzCFqoOLOYbqAITQ=", //cahnnel token
	)*/
	log.Printf("logPath:%s\n", configPath)
	token, secret := config.GetLineConfig(configPath)
	log.Printf("token:%s secret:%s\n", token, secret)
	_bot, err := linebot.New(secret, token)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	bot = _bot
}
