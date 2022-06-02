package app

import (
	"fmt"
	"github.com/gin-gonic/gin"

	. "aliyun/alert_notification_forwarding/pkg/alert"
	"aliyun/alert_notification_forwarding/pkg/chatbot"
)

var (
	webhook = map[string]string{
		"cem":  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=86fc18c6-18b7-413a-b0ac-d9cb206837a8",
		"kbt":  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=c25cb739-4ae3-4269-bf79-1b1b2d9e0b42",
		"prw":  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=b7eab72c-5486-4ed6-89ea-cb564c9622f5",
		"wjyb": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=f8775071-65ca-44ca-b066-5679258cd2f7",
	}
)

func asynNotifaction(alertMessage AliyunAlertMessage) {
	// Old Method
	// var markdown common.Markdown
	// markdown.Content = alertMessage.ToMarkdown()
	//
	// var alertBot common.WXWorkBot
	// alertBot.Message = common.MarkdownMessage{
	// 	MessageType: "markdown",
	// 	Markdown:    &markdown,
	// }
	markdown, err := alertMessage.ToMarkdown()
	if err != nil {
		return
	}

	bot := new(chatbot.WXWorkBot)
	if value, ok := webhook[alertMessage.Project]; ok {
		bot.Webhook = value
		bot.MarkdownMessage(markdown).Send()
	} else {
		fmt.Printf("webhook not found\n")
		return
	}
}

func Run() {
	router := gin.Default()
	router.POST("/alert/:project", func(context *gin.Context) {
		if err := context.Request.ParseForm(); err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		var alerMessage AliyunAlertMessage
		alerMessage.Form = context.Request.PostForm
		alerMessage.Project = context.Param("project")
		go asynNotifaction(alerMessage)

		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})

	})
	// gin.SetMode(gin.ReleaseMode)
	err := router.Run("0.0.0.0:19099")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
