package slack

import "testing"

func TestMain(m *testing.M) {
	message := Message()

	message.Add(
		Text("").
			AddField("*Question*", "*Answer*").
			AddField("How many fields can we have?", ""),
		// AddField("Why?", "Because it's a limitation of the Slack API").
		// AddField("How can I add more?", "You can make another text block").
		// AddField("Lorem", "Ipsum"),
	)

	err := Send("https://hooks.slack.com/services/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", message)
	if err != nil {
		panic(err)
	}
}
