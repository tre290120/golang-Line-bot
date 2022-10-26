package service

import (
	"log"
	"os"
	"simplebot/config"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func InitBot(configPath string) {
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
