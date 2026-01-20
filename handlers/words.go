package handlers

import (
	"fmt"
	"sgb/repository"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var WordsCommand = telego.BotCommand{
	Command:     "words",
	Description: "Get a list of available words",
}

// HandleStart handles the /start command
func HandleWordsCommand(ctx *telegohandler.Context, m telego.Message) error {
	repo := repository.New()
	words, err := repo.GetWords()
	if err != nil {
		return err
	}

	reply := tu.Message(tu.ID(m.Chat.ID), fmt.Sprintf("Words: %v", words))

	_, err = ctx.Bot().SendMessage(ctx, reply)

	return err
}
