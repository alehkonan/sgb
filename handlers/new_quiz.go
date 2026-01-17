package handlers

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleNewQuizCallback(ctx *telegohandler.Context, callback telego.CallbackQuery) error {
	reply := tu.Message(
		tu.ID(callback.Message.GetChat().ID),
		"Unfortunately we don't have any quiz yet",
	)

	_, err := ctx.Bot().SendMessage(ctx, reply)

	return err
}
