package main

import "net/http"
import "fmt"

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request)
	if request.Method == "POST" {
        request.ParseForm()
		fmt.Println("IN POST")
        fmt.Println(request.Form)
        bot := slackbot{"pinkoi", "..."}
        bot.send(fmt.Sprintf("%s%s", "#", request.Form["channel_name"][0]), fmt.Sprintf("@%s: %s [from golang bot]", request.Form["user_name"][0], request.Form["text"][0]))
	}
}

func main() {
	http.HandleFunc("/slack", index)
	http.ListenAndServe(":9001", nil)
}
