package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/teilomillet/gollm"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
	"time"
)

type LLMResponse struct {
	Summary string
	Details string
}

type PromptTemplate struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Template    string   `yaml:"template"`
	Directives  []string `yaml:"directives"`
}

type PromptTemplates []PromptTemplate

func loadPromptTemplates() (PromptTemplate, error) {
	file, err := os.Open("templates.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode yaml to PromptTemplate struct
	var promptTemplates PromptTemplate
	if err := yaml.NewDecoder(file).Decode(&promptTemplates); err != nil {
		log.Fatal(err)
		return promptTemplates, err
	}
	return promptTemplates, nil
}

// cleanJSONResponse removes markdown code block delimiters and trims whitespace
func cleanJSONResponse(response string) string {
	response = strings.TrimSpace(response)
	response = strings.TrimPrefix(response, "```json")
	response = strings.TrimSuffix(response, "```")
	return strings.TrimSpace(response)
}

func QueryLLM(request string) string {

	// Load API key from environment variable
	os.Setenv("OPENAI_API_KEY", "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")
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalf("OPENAI_API_KEY environment variable is not set")
	}

	// todo: Right now loading a single default template. We want choice eventually
	template, err := loadPromptTemplates()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Response:\n%s\n", err)
	}
	fmt.Printf("Response:\n%s\n", template)

	// todo: load values from config
	// Create the LLM instance.
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

	// Create our prompt template to format our expectations and response
	promptTemp := gollm.NewPromptTemplate(
		template.Name,
		template.Description,
		template.Template,
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
	prompt, err := promptTemp.Execute(data)
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

	var result LLMResponse
	err = json.Unmarshal([]byte(cleanedJSON), &result)
	if err != nil {
		log.Printf("Warning: Failed to parse analysis JSON for topic '%s': %v", topic, err)
		log.Printf("Raw response: %s", cleanedJSON)
	}

	return response
}
