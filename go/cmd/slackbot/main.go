package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/toomore/slack/go/slackbot"
)

var (
	token   = flag.String("token", os.Getenv("SLACK_TOKEN"), "Slack token.")
	team    = flag.String("team", os.Getenv("SLACK_TEAM"), "Slack team")
	channel = flag.String("chan", "", "Channel name, # or @")
	msg     = flag.String("msg", "", "Message")
)

func main() {
	flag.Parse()
	log.Println(">>> channel, msg: ", *channel, *msg)
	bot := slackbot.NewBot(*team, *token)
	if resp, err := bot.Send(*channel, *msg); err != nil {
		log.Println(">>> Fail", err)
	} else {
		text, _ := ioutil.ReadAll(resp.Body)
		log.Printf(">>> %s", text)
	}
}
