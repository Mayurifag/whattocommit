package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

var commitMessages []string

func readCommitMessages() error {
	fileContent, err := os.ReadFile("commit_messages.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	commitMessages = strings.Split(string(fileContent), "\n")
	// Last line is empty in commit_messages, so we can easily omit it
	commitMessages = commitMessages[:len(commitMessages)-1]
	return nil
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/all":
		ctx.SetContentType("text/plain")
		ctx.SetStatusCode(fasthttp.StatusOK)
		_, _ = ctx.WriteString(strings.Join(commitMessages, "\n"))
	case "/":
		ctx.SetContentType("text/plain")
		ctx.SetStatusCode(fasthttp.StatusOK)
		_, _ = ctx.WriteString(commitMessages[rand.Intn(len(commitMessages))])
	}
}

func main() {
	err := readCommitMessages()
	if err != nil {
		fmt.Println(err)
		return
	}

	server := fasthttp.Server{
		Handler: requestHandler,
	}

	if err := server.ListenAndServe(":8080"); err != nil {
		fmt.Println("Error in ListenAndServe:", err)
	}
}
