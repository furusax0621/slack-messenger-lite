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

	blocks := []slack.Block{
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("```%s```", content), false, false),
			nil,
			nil,
		),
	}

	err = slack.PostWebhookContext(ctx, url, &slack.WebhookMessage{
		Blocks: &slack.Blocks{BlockSet: blocks},
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
