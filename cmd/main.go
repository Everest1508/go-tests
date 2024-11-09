package main

import (
	"log"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func main() {
	token := "your-slack-token"
	client := slack.New(token, slack.OptionDebug(true))

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <url>")
	}

	attachment := slack.Attachment{
		Pretext: "bulk test failed",
		Text:    `Test Failed ` + os.Args[1],
		Color:   "FF0000",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().Format("Mon 2006 Jan 02 15:04:05 MST"),
			},
		},
	}

	client.PostMessage(
		"#bulk-tests",
		slack.MsgOptionAttachments(attachment),
	)

}
