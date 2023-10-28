package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

var commitMessages []string

func readCommitMessages() error {
	file, err := os.Open("commit_messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			commitMessages = append(commitMessages, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	return nil
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain")
	ctx.SetStatusCode(fasthttp.StatusOK)

	switch string(ctx.Path()) {
	case "/all":
		_, _ = ctx.WriteString(strings.Join(commitMessages, "\n"))
	case "/number":
		_, _ = ctx.WriteString(fmt.Sprint(len(commitMessages)))
	case "/":
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
