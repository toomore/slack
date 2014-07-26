package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type IncomingWebHooks struct {
	team  string
	token string
}

func (bot *IncomingWebHooks) send(channel, text string) {
	params := url.Values{}
	params.Add("token", bot.token)
	params.Add("channel", channel)
	url_path := fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?%s",
		bot.team, params.Encode())
	text_map := make(map[string]string)
	text_map["text"] = text
	text_json, _ := json.Marshal(text_map)
	body := bytes.NewBufferString(fmt.Sprintf("%s", text_json))
	http.Post(url_path, "text/plain", body)
}

func main() {
	bot := IncomingWebHooks{"pinkoi", "..."}
	bot.send("@toomore", "From IncomingWebHooks.")
}
