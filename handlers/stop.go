package handlers

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var StopCommand = telego.BotCommand{
	Command:     "stop",
	Description: "Stop the subscription for the new words",
}

// HandleStart handles the /start command
func HandleStopCommand(ctx *telegohandler.Context, msg telego.Message) error {
	reply := tu.Message(
		tu.ID(msg.Chat.ID),
		"You have successfully unsubscribed",
	)

	_, err := ctx.Bot().SendMessage(ctx, reply)

	return err
}
