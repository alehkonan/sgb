package tp

import "errors"

var (
	ErrorUnknownEventType = errors.New("Unknown event type")
	ErrorUnknownMetaType  = errors.New("Unknown meta type")
)
