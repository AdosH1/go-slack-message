package slack

type SlackText struct {
	Type string `json:"type"`
	Text TextBody `json:"text"`
	Accessory Block `json:"accessory,omitempty"`
}

type TextBody struct {
	Type string `json:"type"`
	Text string `json:"text"`	
}

func Text(text string) *SlackText {
	return &SlackText{
		Type: "section",
		Text: TextBody{
			Type: "mrkdwn",
			Text: text,
		},
	}
}

func (t *SlackText) Add(b Block) *SlackText {
	t.Accessory = b
	return t
}
