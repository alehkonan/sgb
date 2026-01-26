package tp

import (
	"log"
	"strings"
)

const (
	CmdStart = "/start"
	CmdHelp  = "/help"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command %s from %s", text, username)

	switch text {
	case CmdStart:
		return p.handleStartCmd(chatID)
	case CmdHelp:
		return p.handleHelpCmd(chatID)
	default:
		return p.client.SendMessage(chatID, msgUnknownCmd)
	}
}

func (p *Processor) handleStartCmd(chatID int) error {
	return p.client.SendMessage(chatID, msgHello)
}

func (p *Processor) handleHelpCmd(chatID int) error {
	return p.client.SendMessage(chatID, msgHelp)
}
