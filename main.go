package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/spf13/cobra"
)

func parseInstruction(apiKey, instruction, fileContent string) (string, error) {
	// Create the OpenAI client
	client := openai.NewClient(option.WithAPIKey(apiKey))

	// Create the prompt for the API
	prompt := fmt.Sprintf(`
The following is the content of a file:
---
%s
---
Instruction: "%s"

Perform the requested update and return the updated file content.
`, fileContent, instruction)

	// Call the OpenAI API
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You are a helpful assistant for text processing tasks. You are given a file and an instruction, and you need to update the file content based on the instruction. You should only return the updated file content, without any additional text or comments. Please don't encapsulate the response in Markdown or any other formatting."),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse instruction: %w", err)
	}

	return chatCompletion.Choices[0].Message.Content, nil
}

func updateFileWithInstruction(filePath, instruction, apiKey string) error {
	// Read the file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Parse and apply the instruction
	updatedContent, err := parseInstruction(apiKey, instruction, string(fileContent))
	if err != nil {
		return fmt.Errorf("failed to parse instruction: %w", err)
	}

	// Write the updated content back to the file
	err = os.WriteFile(fmt.Sprintf("%s.updated", filePath), []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated file: %w", err)
	}

	return nil
}

func main() {
	var (
		filePath    string
		instruction string
		apiKey      string
	)

	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "file-update",
		Short: "A tool to update file content using natural language instructions with OpenAI",
		Run: func(cmd *cobra.Command, args []string) {
			// Ensure API key is set
			if apiKey == "" {
				log.Fatal("OpenAI API key is not set. Use --api-key flag or set OPENAI_API_KEY environment variable.")
			}

			// Perform the file update
			err := updateFileWithInstruction(filePath, instruction, apiKey)
			if err != nil {
				log.Fatalf("Error updating file: %v", err)
			}

			fmt.Println("File updated successfully.")
		},
	}

	// Add flags
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file to update (required)")
	rootCmd.Flags().StringVarP(&instruction, "instruction", "i", "", "Natural language instruction for the update (required)")
	rootCmd.Flags().StringVarP(&apiKey, "api-key", "k", os.Getenv("OPENAI_API_KEY"), "OpenAI API key (defaults to OPENAI_API_KEY environment variable)")

	// Mark required flags
	rootCmd.MarkFlagRequired("file")
	rootCmd.MarkFlagRequired("instruction")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
