package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, greeting("Code.education Rocks!"))
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func greeting(msg string) string {
	return "<b>" + msg + "</b>"
}
