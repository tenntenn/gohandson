package main

import "net/http"

func init() {
	http.HandleFunc("/post", post)
	http.HandleFunc("/", index)
}
