package main

import (
	"context"
	"log"
	"os"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"

	"sgb/handlers"
	"sgb/middlewares"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	bot, err := telego.NewBot(token)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// defer bot.Close(ctx)

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get updates: %v", err)
	}

	handler, err := telegohandler.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatalf("Failed to create bot handler: %v", err)
	}

	defer handler.Stop()

	handler.Use(middlewares.LogUserMessage)

	err = bot.SetMyCommands(ctx, &telego.SetMyCommandsParams{
		Commands: []telego.BotCommand{
			handlers.StartCommand,
			handlers.StopCommand,
			handlers.QuizCommand,
		},
	})
	if err != nil {
		log.Printf("Failed set a list of commands: %v", err)
	}

	handler.HandleMessage(handlers.HandleStartCommand, telegohandler.CommandEqual(handlers.StartCommand.Command))
	handler.HandleMessage(handlers.HandleStopCommand, telegohandler.CommandEqual(handlers.StopCommand.Command))
	handler.HandleMessage(handlers.HandleQuizCommand, telegohandler.CommandEqual(handlers.QuizCommand.Command))
	handler.HandleCallbackQuery(handlers.HandleNewQuizCallback, telegohandler.CallbackDataEqual("new_quiz"))

	log.Println("Bot is starting...")
	if err := handler.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}
