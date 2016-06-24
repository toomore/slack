package slackbot

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// Bot struct
type Bot struct {
	params  url.Values
	baseurl string
}

// NewBot new one
func NewBot(team, token string) *Bot {
	return &Bot{
		params:  url.Values{"token": []string{token}},
		baseurl: fmt.Sprintf("https://%s.slack.com/services/hooks/slackbot", team),
	}
}

// Send message
func (b Bot) Send(channel, text string) {
	b.params.Add("channel", channel)
	http.Post(fmt.Sprintf("%s?%s", b.baseurl, b.params.Encode()), "text/plain", bytes.NewBufferString(text))
}
