package main

import (
	"testing"
)

func TestParsePayload(t *testing.T) {
	p := new(event)
	err := p.parseFile("./test/payload.json")
	if err != nil {
		t.Fatalf("Got parsing error: %s", err)
	}
	if *p.Review.Body != "merge" {
		t.Fatalf("Wrong review body: %s", *p.Review.Body)
	}
	if *p.Review.User.Login != "g4s8" {
		t.Fatalf("Wrong user login: %s", *p.Review.User.Login)
	}
	if *p.PR.Head.Ref != "test" {
		t.Fatalf("Wrong head ref: %s", *p.PR.Head.Ref)
	}
}
