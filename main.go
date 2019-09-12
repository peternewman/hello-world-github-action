package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type event struct {
	Review *struct {
		Body *string `json:"body"`
		User *struct {
			Login *string `json:"login"`
		} `json:"user"`
	} `json:"review"`
	PR *struct {
		Head *struct {
			Ref *string `json:"ref"`
		} `json:"head"`
	} `json:"pull_request"`
}

func (e *event) String() string {
	return fmt.Sprintf("review: `%s` by `%s`; head=`%s`",
		*e.Review.Body, *e.Review.User.Login, *e.PR.Head.Ref)
}

func (e *event) parseReader(r io.Reader) error {
	return json.NewDecoder(r).Decode(e)
}

func (e *event) parseFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return e.parseReader(bufio.NewReader(f))
}

func main() {
	evt := os.Getenv("GITHUB_EVENT_NAME")
	if evt != "pull_request_review" {
		log.Fatalf("Unsupported GitHub event: %s", evt)
	}
	path := os.Getenv("GITHUB_EVENT_PATH")
	payload := new(event)
	if err := payload.parseFile(path); err != nil {
		log.Fatalf("Failed to parse payload json: %s", err)
	}
	if payload.Review.Body == "merge" {
		if _, err := os.Exec("git", []string{"checkout", "master"}); err != nil {
			log.Fatalf("Failed to checkout: %s", err)
		}
		if _, err := os.Exec("git", []string{"merge", payload.PR.Head.Ref}); err != nil {
			log.Fatalf("Failed to merge: %s", err)
		}
	}
	fmt.Printf("success! payload: %s\n", payload)
}
