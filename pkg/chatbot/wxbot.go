package chatbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WXWorkBot struct {
	Webhook   string
	Header    string
	Message   []byte
	GroupChat string
	ChatId    string
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type Markdown struct {
	Content string `json:"content"`
}

type Image struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

type News struct {
	Articles *[8]Article `json:"articles"`
}

type File struct {
	MediaId string `json:"media_id"`
}

type TextMessage struct {
	MessageType string `json:"msgtype"`
	Text        *Text  `json:"text"`
}

type MarkdownMessage struct {
	MessageType string    `json:"msgtype"`
	Markdown    *Markdown `json:"markdown"`
}

type ImageMessage struct {
	MessageType string `json:"msgtype"`
}

type NewsMessage struct {
	MessageType string `json:"msgtype"`
	News        *News  `json:"news"`
}

type FileMessage struct {
	MessageType string `json:"msgtype"`
	File        *File  `json:"file"`
}

func (w *WXWorkBot) MarkdownMessage(markdown string) *WXWorkBot {
	w.Message = []byte(fmt.Sprintf(`{
		"msgtype": "markdown",
		"markdown": {
			"content": "%v"
		}
	}`, markdown))
	return w
}

func (w *WXWorkBot) Send() {
	response, err := http.Post(w.Webhook, "application/json", bytes.NewBuffer(w.Message))
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}(response.Body)
}

func (w *WXWorkBot) Send2() {
	jsonData, err := json.Marshal(&w.Message)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	response, err := http.Post(w.Webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}(response.Body)
}
