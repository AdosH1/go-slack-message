package models

type SlackMessage struct {
	Blocks []Block `json:"blocks"`
}

type Block interface {}

func Message() *SlackMessage {
	return &SlackMessage{
		Blocks: []Block{},
	}
}

func (m *SlackMessage) Add(b Block) {
	m.Blocks = append(m.Blocks, b)
}