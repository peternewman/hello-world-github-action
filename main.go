package main

import (
	"bufio"
	"bytes"
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

func runCmd(cmd *exec.Cmd) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	fmt.Printf("CMD: %s\n", cmd)
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		log.Fatalf("%s - %s", cmd, err)
	}
	fmt.Println(out.String())
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
	fmt.Printf("payload: %s\n", payload)
	if *payload.Review.Body == "merge" {
		runCmd(exec.Command("git", "checkout", "master"))
		runCmd(exec.Command("git", "merge", *payload.PR.Head.Ref))
		runCmd(exec.Command("git", "push", "origin", "master"))
		fmt.Printf("done: merged\n")
	}
}
