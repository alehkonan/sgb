package handlers

import (
	"fmt"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var StartCommand = telego.BotCommand{
	Command:     "start",
	Description: "Start the subscription for the new words",
}

// HandleStart handles the /start command
func HandleStartCommand(ctx *telegohandler.Context, msg telego.Message) error {
	reply := tu.Message(
		tu.ID(msg.Chat.ID),
		fmt.Sprintf("Hello, %s! 👋\nYou have successfully started the subscription!", msg.From.FirstName),
	)

	_, err := ctx.Bot().SendMessage(ctx, reply)

	return err
}
