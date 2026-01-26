package tp

import (
	"fmt"

	"github.com/alehkonan/sgb/packages/clients/tc"
	"github.com/alehkonan/sgb/packages/processors"
	"github.com/alehkonan/sgb/packages/storage"
)

type Processor struct {
	client  *tc.Client
	storage *storage.Storage
	offset  int
}

type Meta struct {
	ChatID   int
	Username string
}

// Creates new telegram processor
func New(client *tc.Client, storage storage.Storage) Processor {
	return Processor{
		client:  client,
		storage: &storage,
	}
}

func (p *Processor) Fetch(limit int) ([]processors.Event, error) {
	updates, err := p.client.Updates(p.offset, limit)
	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]processors.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, parseEvent(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Handle(event processors.Event) error {
	switch event.Type {
	case processors.Message:
		return p.processMessage(event)
	default:
		return ErrorUnknownEventType
	}
}

func (p *Processor) processMessage(event processors.Event) error {
	meta, err := parseMeta(event)
	if err != nil {
		return fmt.Errorf("can't process message, %w", err)
	}

	if err = p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return fmt.Errorf("can't process message, %w", err)
	}

	return nil
}

func parseMeta(event processors.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf(
			"can't parse meta from event: %w",
			ErrorUnknownMetaType,
		)
	}
	return res, nil
}

func parseEvent(update tc.Update) processors.Event {
	updateType := fetchType(update)

	res := processors.Event{
		Type: updateType,
		Text: fetchText(update),
	}

	if updateType == processors.Message {
		res.Meta = Meta{
			ChatID:   update.Message.Chat.ID,
			Username: update.Message.From.Username,
		}
	}

	return res
}

func fetchType(update tc.Update) processors.Type {
	if update.Message == nil {
		return processors.Unknown
	}

	return processors.Message
}

func fetchText(update tc.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}
