package processors

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Handler interface {
	Handle(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta any
}
