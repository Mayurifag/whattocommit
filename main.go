package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var commitMessages []string

func readCommitMessages() error {
	fileContent, err := os.ReadFile("commit_messages.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	commitMessages = strings.Split(string(fileContent), "\n")
	commitMessages = commitMessages[:len(commitMessages)-1]
	return nil
}

func main() {
	err := readCommitMessages()
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()

	router.GET("/all", func(c *gin.Context) {
		c.IndentedJSON(200, commitMessages)
	})

	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte(commitMessages[rand.Intn(len(commitMessages))]))
	})

	router.Run()
}
