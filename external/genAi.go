package external

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func QueryOpenAi(request openai.ChatCompletionRequest) openai.ChatCompletionResponse {

	// Create new client
	client := openai.NewClient("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")

	resp, err := client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return resp
}

func GenAi() {

}
