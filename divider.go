package slack

type SlackDivider struct {
	Type string `json:"type"`
}

func Divider() SlackDivider {
	return SlackDivider{
		Type: "divider",
	}
}
