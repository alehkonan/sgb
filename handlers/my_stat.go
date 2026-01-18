package handlers

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var MyStatCommand = telego.BotCommand{
	Command:     "my_stat",
	Description: "Show my statistics",
}

func MyStatCommandHandler(ctx *telegohandler.Context, m telego.Message) error {
	_, err := ctx.Bot().SendMessage(
		ctx.Context(),
		tu.Message(tu.ID(m.Chat.ID), "Here is your statistics"),
	)

	return err
}
