package middlewares

import (
	"fmt"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

func LogUserMessage(ctx *telegohandler.Context, update telego.Update) error {
	if update.Message != nil && update.Message.From != nil {
		fmt.Printf(
			"Incoming message from user %s %s\n- %s\n",
			update.Message.From.FirstName,
			update.Message.From.LastName,
			update.Message.Text,
		)
	}
	return ctx.Next(update)
}
