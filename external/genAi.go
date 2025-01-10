package external

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func QueryOpenAi(request openai.ChatCompletionRequest) openai.ChatCompletionResponse {

	// Create new client
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return resp
}

func GenAi() {

}
