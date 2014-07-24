package main

import "net/http"
import "fmt"
import "strings"

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request)
	if request.Method == "POST" {
		request.ParseForm()
		fmt.Println(request.Form)
		if strings.HasPrefix(request.Form["text"][0], "ok go") {
			bot := slackbot{"pinkoi", "..."}
			channel := fmt.Sprintf("%s%s", "#", request.Form["channel_name"][0])
			text := fmt.Sprintf("@%s: %s [from golang bot]", request.Form["user_name"][0], request.Form["text"][0][len("ok go "):])
			bot.send(channel, text)
		}
	}
}

func main() {
	http.HandleFunc("/slack", index)
	http.ListenAndServe(":9001", nil)
}
