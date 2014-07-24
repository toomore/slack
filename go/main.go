package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

type slackbot struct {
	team  string
	token string
}

func (bot *slackbot) send(channel, text string) {
	params := url.Values{}
	params.Add("token", bot.token)
	params.Add("channel", channel)
	url_path := fmt.Sprintf("https://%s.slack.com/services/hooks/slackbot?%s",
		bot.team, params.Encode())
	body := bytes.NewBufferString(text)
	http.Post(url_path, "text/plain", body)
}

//func main() {
//	team := flag.String("team", "pinkoi", "Team name")
//	token := flag.String("token", "...", "Slack token")
//	text := flag.String("text", "Hello Wrold! [from golang]", "Say something...")
//	channel := flag.String("chnl", "@toomore", "Send to who ...")
//	flag.Parse()
//	bot := slackbot{*team, *token}
//	bot.send(*channel, *text)
//}
