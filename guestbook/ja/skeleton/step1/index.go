package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	// TODO: "Hello, Google App Engine"とINFOログで出す。
	// TODO: "Hello, Google App Engine"とwに書き込む。
}
