package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Check if the user provided a prompt as an argument
	if len(os.Args) <= 1 {
		fmt.Println("Usage: gptbash.exe \"<Description of what you are trying to do>\"")
		os.Exit(1)
	}

	// Create a new OpenAI client with your API key
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	//prompt := flag.String("prompt", "", "The command to execute")

	var prompt string

	forceful := flag.Bool("f", false, "Forceful mode: bypass confirmation prompt and execute the command immediately")
	flag.Parse()

	if len(flag.Args()) > 0 {
		prompt = strings.Join(flag.Args(), " ")
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		prompt, _ = reader.ReadString('\n')
		prompt = strings.TrimSpace(prompt)
	}

	fmt.Printf("Prompt: %s", prompt)

	// Send the prompt to the OpenAI API and get the response message
	response, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model: "text-davinci-003",
			Prompt: fmt.Sprintf(`This is a conversation between a Command Tool Helper called GPTBash. This tool converts an instruction into a command. So it only outputs a command with no extra text.
								 For example a conversation goes like this.
								 Prompt: list all the file name in current directories with their sizes. 
								 Command: ls -a
								 Prompt: %s.`, prompt),
			MaxTokens:   64,
			Temperature: 0.3,
		},
	)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}

	// Get the generated command from the response message

	response_text := response.Choices[0].Text
	command := strings.Split(response_text, "Command: ")[1]
	answer := "y"
	fmt.Printf("OpenAI response returned is : `%s` without quotes.\n", command)
	if !*forceful {
		fmt.Print("Do you want to execute this command? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))
		// Execute the command using Bash or Windows Terminal
	}
	var shell string
	if answer == "y" {
		{
			if os.Getenv("SHELL") == "/bin/bash" {
				shell = "bash"
			} else {
				shell = "powershell"
			}
			cmd := exec.Command(shell, "-c", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println("Command execution aborted.")
		os.Exit(1)
	}
}
