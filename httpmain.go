package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/health", health)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Server has been started on port 50080")

	http.ListenAndServe(":50080", nil)

}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

//ping
//pong
