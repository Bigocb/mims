package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/teilomillet/gollm"
	"github.com/teilomillet/gollm/llm"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
	"time"
)

// Response object to help standardize output from various LLM tools
type Response struct {
	Summary string
	Details string
}

// PromptTemplate used to map from template.yaml
type PromptTemplate struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Template    string   `yaml:"template"`
	Directives  []string `yaml:"directives"`
}

// loadPromptTemplates loads yaml template(s) and maps them to PromptTemplate struct
func loadPromptTemplates() (PromptTemplate, error) {
	file, err := os.Open("./llm/templates/templates.yaml")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Decode yaml to PromptTemplate struct
	var promptTemplates PromptTemplate
	if err := yaml.NewDecoder(file).Decode(&promptTemplates); err != nil {
		log.Fatal(err)
		return promptTemplates, err
	}
	return promptTemplates, nil
}

// buildPrompt responsible for building the prompt to be sent to the LLM
func buildPrompt(request string) *llm.Prompt {

	topic := strings.Replace(request, "Research", "", -1)
	data := map[string]interface{}{
		"Topic": topic,
	}
	// todo: Right now loading a single default template. We want choice eventually
	template, err := loadPromptTemplates()
	if err != nil {
		fmt.Printf("Response:\n%s\n", err)
		log.Fatal(err)
	}

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

	// Build prompt object
	prompt, err := promptTemp.Execute(data)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	return prompt
}

// cleanJSONResponse removes markdown code block delimiters and trims whitespace. Useful for some AI responses                                          .
func cleanJSONResponse(response string) string {
	response = strings.TrimSpace(response)
	response = strings.TrimPrefix(response, "```json")
	response = strings.TrimSuffix(response, "```")
	return strings.TrimSpace(response)
}

// buildClient builds the LLM client needed based on general config or override
func buildClient() gollm.LLM {
	// Load API key from environment variable
	err := os.Setenv("OPENAI_API_KEY", "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")
	if err != nil {
		return nil
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalf("OPENAI_API_KEY environment variable is not set")
	}

	// todo: load values from config. Helper functions. Choose your model. Hmm, Ollamma?
	// Create the LLM instance.
	llmClient, err := gollm.NewLLM(
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

	return llmClient
}

// Query responsible for building the request and sending
func Query(request string) Response {

	ctx := context.Background()

	// Build prompt
	prompt := buildPrompt(request)

	llmClient := buildClient()
	// Generate a response from our LLM
	response, err := llmClient.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate text: %v", err)
	}

	// Clean the JSON response
	cleanedJSON := cleanJSONResponse(response)

	var result Response
	err = json.Unmarshal([]byte(cleanedJSON), &result)
	if err != nil {
		log.Printf("Warning: Failed to parse analysis JSON for topic : %v", err)
		log.Printf("Raw response: %s", cleanedJSON)
	}

	return result
}
