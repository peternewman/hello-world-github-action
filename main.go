package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type event struct {
	Review *struct {
		Body *string `json:"body"`
		User *struct {
			Login *string `json:"login"`
		} `json:"user"`
	} `json:"review"`
	Head *struct {
		Ref *string `json:"ref"`
	} `json:"head"`
}

func (e *event) String() string {
	return fmt.Sprintf("review: `%s` by `%s`; head=`%s`",
		e.Review.Body, e.Review.User, e.Head.Ref)
}

func main() {
	evt := os.Getenv("GITHUB_EVENT_NAME")
	if evt != "pull_request_review" {
		log.Fatalf("Unsupported GitHub event: %s", evt)
	}
	path := os.Getenv("GITHUB_EVENT_PATH")
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Event payload not found at `%s`", path)
	}
	payload := new(event)
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(payload); err != nil {
		log.Fatalf("Failed to parse payload json: %s", err)
	}
	fmt.Printf("success! payload: %s\n", payload)
}
