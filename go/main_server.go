package main

import "fmt"
import "net/http"
import "net/url"
import "strings"

func BotSend(formdata *url.Values, msg string) {
	form := *formdata
	bot := slackbot{"...", "..."}
	channel := fmt.Sprintf("%s%s", "#", form.Get("channel_name"))
	text := fmt.Sprintf("@%s: %s [from golang bot]", form.Get("user_name"), msg)
	bot.send(channel, text)
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request)
	if request.Method == "POST" {
		request.ParseForm()
		fmt.Println(request.Form)
		switch {
		case strings.HasPrefix(request.Form.Get("text"), "ok go Hello"):
			BotSend(&request.Form, request.Form.Get("text")[len("ok go "):])

		case strings.HasPrefix(request.Form.Get("text"), "ok go time"):
			BotSend(&request.Form, request.Form.Get("text")[len("ok go "):])
		}
	}
}

func main() {
	http.HandleFunc("/slack", index)
	http.ListenAndServe(":9001", nil)
}
