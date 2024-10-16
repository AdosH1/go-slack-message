package models

type SlackText struct {
	Type string `json:"type"`
	Text TextBody `json:"text"`
}

type TextBody struct {
	Type string `json:"type"`
	Text string `json:"text"`	
}

func Text(text string) SlackText {
	return SlackText{
		Type: "section",
		Text: TextBody{
			Type: "mrkdwn",
			Text: text,
		},
	}
}