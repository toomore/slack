package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type slackbot struct {
	team  string
	token string
}

func (bot *slackbot) send(channel, text string) {
	var url_path string
	url_path = fmt.Sprintf("https://%s.slack.com/services/hooks/slackbot?token=%s&channel=%s", bot.team, bot.token, channel)

	fmt.Println(url_path)
	body := bytes.NewBufferString(text)
	http.Post(url_path, "text/plain", body)
}

func main() {
	bot := slackbot{"pinkoi", "..."}
	bot.send("@toomore", "Hello, Toomore [from golang]")
}
