package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"jervis/data"
	"jervis/external"
	"os"
	"strconv"
	"strings"
	"time"
)

// todo: right now the when the job runs it creates a content.db in the directory its run in.
// need a more centralized approach. Options. Create the db in a .mim folder in the users home directory

// UserContext for the start of the conversation. Currently loaded from a config file.
//type UserContext struct {
//	Role    string
//	Content string
//}
//
//// todo: helpers?
//func loadContext() []openai.ChatCompletionMessage {
//
//	// Open local file for user context
//	data, err := os.ReadFile("context.json")
//	if err != nil {
//		fmt.Println("Error reading file:", err)
//		return nil
//	}
//
//	// Process config json and load to struct
//	var jsonData []UserContext
//	err = json.Unmarshal(data, &jsonData)
//	if err != nil {
//		fmt.Println("Error unmarshling file:", err)
//		return nil
//	}
//
//	// Beginning of conversation set context about our user
//	var messages []openai.ChatCompletionMessage
//	for _, value := range jsonData {
//		message := openai.ChatCompletionMessage{
//			Role:    strings.Trim(value.Role, "\""),
//			Content: value.Content,
//		}
//		messages = append(messages, message)
//	}
//	return messages
//}

func main() {

	fmt.Print("> ")
	message := []openai.ChatCompletionMessage{
		openai.ChatCompletionMessage{
			Role:    "system",
			Content: "You are a helpful AI, named mim",
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: message,
	}

	s := bufio.NewScanner(os.Stdin)
	var searchResult []data.JervisData
	var err error
	for s.Scan() {

		requestText := s.Text()
		fmt.Print("> ")

		if strings.Contains(s.Text(), "Search:") {
			updMessage := strings.ReplaceAll(s.Text(), "Search:", "")

			searchResult, err = data.Search(updMessage)
			if err != nil {
				fmt.Println("Error searching chats about:", err)
			}
			fmt.Println("Here are the results: ")
			fmt.Println(searchResult)
			fmt.Println(">")
			continue
		}

		if strings.Contains(s.Text(), "Add Context:") {
			out, err := json.Marshal(searchResult)
			if err != nil {
				fmt.Println("Error marshalling search result:", err)
			}
			var updateText string
			updateText = "These are previous responses on this topic. Please use these as context when responding. " + string(out)
			requestText = updateText
		}

		if strings.Contains(s.Text(), "Research:") {

			req.Messages = append(req.Messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: requestText,
			})

			var resp openai.ChatCompletionResponse
			resp = external.QueryOpenAi(req)
			// todo: Better response formatting
			fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
			req.Messages = append(req.Messages, resp.Choices[0].Message)
			fmt.Print("> ")
		}

		if strings.Contains(s.Text(), "End:") {
			break
		}
		if strings.Contains(s.Text(), "Save:") {
			timeNow := time.Now().Unix()
			message := req.Messages[len(req.Messages)-1]

			messageObject := data.JervisData{
				Key:     strconv.FormatInt(timeNow, 10),
				Message: message.Content,
			}

			if err := data.Put(&messageObject); err != nil {
				fmt.Print("bad put")
			}
		}

	}
}
