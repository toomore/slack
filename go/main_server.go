package main

import "fmt"
import "net/http"
import "net/url"
import "strings"

func BotSend(form url.Values) {
	bot := slackbot{"pinkoi", "..."}
	channel := fmt.Sprintf("%s%s", "#", form["channel_name"][0])
	text := fmt.Sprintf("@%s: %s [from golang bot]", form["user_name"][0], form["text"][0][len("ok go "):])
	bot.send(channel, text)
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request)
	if request.Method == "POST" {
		request.ParseForm()
		fmt.Println(request.Form)
		if strings.HasPrefix(request.Form["text"][0], "ok go") {
			BotSend(request.Form)
		}
	}
}

func main() {
	http.HandleFunc("/slack", index)
	http.ListenAndServe(":9001", nil)
}
