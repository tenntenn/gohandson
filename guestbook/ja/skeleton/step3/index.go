package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var msg string
	// TODO: msgという名前のリクエストパラメタを取得
	if msg == "" {
		msg = "NO MESSAGE"
	}
	fmt.Fprintln(w, msg)
}
