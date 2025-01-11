package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/teilomillet/gollm"
	"log"
	"os"
	"strings"
	"time"
)

type GptResponse struct {
	Summary string
	Details string
}

// cleanJSONResponse removes markdown code block delimiters and trims whitespace
func cleanJSONResponse(response string) string {
	response = strings.TrimSpace(response)
	response = strings.TrimPrefix(response, "```json")
	response = strings.TrimSuffix(response, "```")
	return strings.TrimSpace(response)
}

func QueryOpenAi(request string) string {
	// Load API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalf("OPENAI_API_KEY environment variable is not set")
	}

	// Create the LLM instance
	llm, err := gollm.NewLLM(
		gollm.SetProvider("openai"),
		gollm.SetModel("gpt-3.5-turbo-0125"),
		gollm.SetAPIKey(apiKey),
		gollm.SetMaxTokens(1000),
		gollm.SetMaxRetries(3),
		gollm.SetRetryDelay(time.Second*2),
		gollm.SetLogLevel(gollm.LogLevelInfo),
	)
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}

	ctx := context.Background()

	// TODO: Externalize from this file. We want to have multiple templates for choice.
	// Create our prompt template to format our expectations and response
	template := gollm.NewPromptTemplate(
		"GeneralTemplate",
		"A template for answering questions",
		"Provide a comprehensive summary of {{.Topic}}.",
		gollm.WithPromptOptions(
			gollm.WithDirectives(
				"Use clear and concise language",
				"Provide specific examples where appropriate",
				"There is no maximum length",
			),
			gollm.WithOutput("Output in json format with the following structure:\n"+"{\n  \"summary\": \"this is a summary\",\n\"details\": \"details of the topic\"\n\n}"),
		),
	)

	// Trim research from our input. Temporary for POC
	topic := strings.Replace(request, "Research", "", -1)
	data := map[string]interface{}{
		"Topic": topic,
	}

	// Build prompt object
	prompt, err := template.Execute(data)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	// Generate a response from our LLM
	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate text: %v", err)
	}
	fmt.Printf("Response:\n%s\n", response)

	// Clean the JSON response
	cleanedJSON := cleanJSONResponse(response)

	var result GptResponse
	err = json.Unmarshal([]byte(cleanedJSON), &result)
	if err != nil {
		log.Printf("Warning: Failed to parse analysis JSON for topic '%s': %v", topic, err)
		log.Printf("Raw response: %s", cleanedJSON)
	}

	return response
}
