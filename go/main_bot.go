package main

import "flag"

func main() {
	team := flag.String("team", "pinkoi", "Team name")
	token := flag.String("token", "...", "Slack token")
	text := flag.String("text", "Hello Wrold! [from golang]", "Say something...")
	channel := flag.String("chnl", "@toomore", "Send to who ...")
	flag.Parse()
	bot := slackbot{*team, *token}
	bot.send(*channel, *text)
}
