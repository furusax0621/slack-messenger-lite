package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type content struct {
	Text string `json:"text"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := os.Getenv("SLACK_WEBHOOK_URL")
	if url == "" {
		log.Printf("SLACK_WEBHOOK_URL is missing")
		os.Exit(1)
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("unknown error: %s", err.Error())
		os.Exit(1)
	}

	body, err := json.Marshal(content{Text: string(data)})
	if err != nil {
		log.Printf("unknown error: %s", err.Error())
		os.Exit(1)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		log.Printf("unknown error: %s", err.Error())
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("unknown error: %s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, req.Body)
}
