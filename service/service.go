package service

import (
	"encoding/json"
	"log"
	"simplebot/db"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func DoService(c *gin.Context) {

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(400, "")
		} else {
			c.JSON(500, "")
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			userID := event.Source.UserID
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				str := ""
				if message.Text == "#ls" {
					str = list(userID)
				} else {
					str = save(userID, message.Text)
					str = message.Text + " " + str
				}
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do()
				if err != nil {
					log.Print(err)
				}
			}
		}
	}

}

func list(userID string) string {
	str := ""
	msgs, ok := db.Get(userID)
	if !ok {
		str = "list msg failed"
	} else {
		bytes, _ := json.Marshal(msgs)
		str = string(bytes)
	}
	return str
}

func save(userID string, msg string) string {
	str := ""
	ok := db.Insert(db.Message{ID: userID, Msg: msg})
	if !ok {
		str = "save msg failed"
	}
	return str
}
