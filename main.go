package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := os.Getenv("SLACK_WEBHOOK_URL")
	if url == "" {
		log.Printf("SLACK_WEBHOOK_URL is missing")
		os.Exit(1)
	}

	content, err := getContent()
	if err != nil {
		log.Printf("unknown error: %s", err.Error())
	}

	err = slack.PostWebhookContext(ctx, url, &slack.WebhookMessage{
		Text: fmt.Sprintf("```%s```", content),
	})
	if err != nil {
		log.Printf("unknown error: %s", err.Error())
	}
}

func getContent() (string, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
