package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var msg string
	msg = r.FormValue("msg")
	if msg == "" {
		msg = "NO MESSAGE"
	}
	fmt.Fprintln(w, msg)
}
