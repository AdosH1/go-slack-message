package slack

type SlackText struct {
	Type      string     `json:"type"`
	Text      *TextBody  `json:"text,omitempty"`
	Fields    []TextBody `json:"fields,omitempty"`
	Accessory Block      `json:"accessory,omitempty"`
}

type TextBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func Text(text string) *SlackText {
	if text == "" {
		return &SlackText{
			Type: "section",
		}
	}
	return &SlackText{
		Type: "section",
		Text: &TextBody{
			Type: "mrkdwn",
			Text: text,
		},
	}
}

func (t *SlackText) AddAccessory(b Block) *SlackText {
	t.Accessory = b
	return t
}

func (t *SlackText) AddField(f1 string, f2 string) *SlackText {
	// Slack Block API does not allow empty strings for fields
	// so we chech for empty strings here before adding them
	if f1 != "" {
		field1 := TextBody{
			Type: "mrkdwn",
			Text: f1,
		}
		t.Fields = append(t.Fields, field1)
	}
	if f2 != "" {
		field2 := TextBody{
			Type: "mrkdwn",
			Text: f2,
		}
		t.Fields = append(t.Fields, field2)
	}
	return t
}
