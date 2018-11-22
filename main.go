package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Serve Http:", err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path[1:] {
	case "ping":
		fmt.Fprint(w, "pong")
	case "hi":
		fmt.Fprint(w, "holla")
	default:
		fmt.Fprint(w, "Hi Nabil, Welcome to Devops CI CD Pipeline")
		fmt.Fprint(w, "My First Pipeline project is successful")
	}
}
