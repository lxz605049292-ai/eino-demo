package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("you are an assistant"),
		schema.MessagesPlaceholder("chat_history", true),
		schema.UserMessage("question: {question}"),
	)
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "assistantor",
		"style":    "brave, positive",
		"question": "my code looks very negative",
		"chat_history": []*schema.Message{
			schema.UserMessage("hello"),
			schema.AssistantMessage("i am a assistantor for your write code", nil),
			schema.UserMessage("i think my code looks very negative"),
			schema.AssistantMessage("i think you can watch more beautiful code,then you can write code you looks beautiful", nil),
		},
	})
	if err != nil {
		log.Fatalf("template.Format failed: %v", err)
	}

	chatModel, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		BaseURL: "https://ark.cn-beijing.volces.com/api/v3",
		Model:   os.Getenv("ARK_ENDPOINT_ID"),
		APIKey:  os.Getenv("ARK_API_KEY"),
	})
	if err != nil {
		log.Fatalf("openai.NewChatModel failed: %v", err)
	}

	result, err := chatModel.Generate(context.Background(), messages)
	if err != nil {
		log.Fatalf("chatModel.Generate failed: %v", err)
	}
	fmt.Println(result.Content)
}
