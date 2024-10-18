package slack

type SlackImage struct {
	Type string `json:"type"`
	ImageUrl string `json:"image_url"`
	AltText string `json:"alt_text"`
}

func Image(url string, altText string) SlackImage {
	return SlackImage{
		Type: "image",
		ImageUrl: url,
		AltText: altText,
	}
}