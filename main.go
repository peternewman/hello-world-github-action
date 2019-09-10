package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hello from `%s`\n", os.Getenv("GITHUB_EVENT_NAME"))
}
