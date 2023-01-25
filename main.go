package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

aa:
	for {
		fmt.Println("Enter your prompt:(q if you want to exit!!!)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		prompt := scanner.Text()
		if prompt == "q" {
			break aa
		}

		err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
			Prompt: []string{
				prompt,
			},
			MaxTokens:   gpt3.IntPtr(1000),
			Temperature: gpt3.Float32Ptr(0),
		}, func(resp *gpt3.CompletionResponse) {
			fmt.Print(resp.Choices[0].Text)
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(13)
		}
		fmt.Printf("\n")
	}

}
//env GOOS=windows GOARCH=amd64 go build .
