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

func (bot *IncomingWebHooks) Send(channel, text string) {
	text_map := make(map[string]interface{})
	text_map["text"] = text
	text_map["channel"] = channel
	text_map["username"] = "PinkoiBot [go]"
	text_map["icon_emoji"] = ":mypinkoi:"

	bot.Post(text_map)
}

func (bot *IncomingWebHooks) SendAtta(channel, text, fallback, title, value string) {
	text_map := make(map[string]interface{})
	text_map["text"] = text
	text_map["channel"] = channel
	text_map["username"] = "PinkoiBot [go]"
	text_map["icon_emoji"] = ":mypinkoi:"

	atta := bot.RenderAtta(fallback, text, title, value)
	text_map["attachments"] = []map[string]interface{}{atta}
	bot.Post(text_map)
}

func (bot *IncomingWebHooks) Post(text_map map[string]interface{}) {
	text_json, _ := json.Marshal(text_map)
	body := bytes.NewBufferString(fmt.Sprintf("%s", text_json))

	params := url.Values{}
	params.Add("token", bot.token)
	url_path := fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?%s",
		bot.team, params.Encode())
	http.Post(url_path, "text/plain", body)
}

func (bot *IncomingWebHooks) RenderAtta(fallback, text, title, value string) map[string]interface{} {
	result := make(map[string]interface{})
	result["fallback"] = fallback
	result["pretext"] = fallback
	result["color"] = "#5060ef"
	fields := make(map[string]interface{})
	fields["title"] = title
	fields["value"] = value
	fields["short"] = false
	result["fields"] = []map[string]interface{}{fields}
	return result
}

func main() {
	bot := IncomingWebHooks{"pinkoi", "..."}
	//bot.Send("@toomore", "From IncomingWebHooks.")
	bot.SendAtta("@toomore", "text", "fallback", "title", "value")
	//j, _ := json.Marshal(bot.renderAtta("fallback", "text", "title", "value"))
	//fmt.Printf("%s", j)
}
