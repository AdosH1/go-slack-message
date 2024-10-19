package slack

type SlackMessage struct {
	Blocks []Block `json:"blocks"`
}

type Block interface{}

func Message() SlackMessage {
	return SlackMessage{
		Blocks: []Block{},
	}
}

func (m *SlackMessage) Add(blocks ...Block) {
	for _, b := range blocks {
		m.Blocks = append(m.Blocks, b)
	}
}
