package main

import "net/http"
import "fmt"

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request)
	if request.Method == "POST" {
		fmt.Println("IN POST")
	}
}

func main() {
	http.HandleFunc("/slack", index)
	http.ListenAndServe(":9001", nil)
}
