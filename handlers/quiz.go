package handlers

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var QuizCommand = telego.BotCommand{
	Command:     "quiz",
	Description: "Start a quiz",
}

// HandleStart handles the /quiz command
func HandleQuizCommand(ctx *telegohandler.Context, msg telego.Message) error {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("New Quiz").WithCallbackData("new_quiz"),
			tu.InlineKeyboardButton("Random action").WithCallbackData("random"),
		),
	)

	reply := tu.Message(tu.ID(msg.Chat.ID), "What do you want to do?").WithReplyMarkup(keyboard)

	_, err := ctx.Bot().SendMessage(ctx, reply.WithProtectContent())

	return err
}
