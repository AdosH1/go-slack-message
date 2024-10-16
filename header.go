package slack

type SlackHeader struct {
	Type string `json:"type"`
	Text HeaderBody `json:"text"`
}

type HeaderBody struct {
	Type string `json:"type"`
	Text string `json:"text"`	
}

func Header(text string) SlackHeader {
	return SlackHeader{
		Type: "header",
		Text: HeaderBody{
			Type: "plain_text",
			Text: text,
		},
	}
}