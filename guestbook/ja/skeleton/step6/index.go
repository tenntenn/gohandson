package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>ゲストブック</title>
	</head>
	<body>
	<form action="/post">
		<input type="text" name="name" placeholder="お名前">
		<input type="text" name="message" placeholder="メッセージ">
		<input type="submit">
	</form>
	<div class="messages">{{range .}}
		<div class="message">
			<h2 class="message-name">{{.Name}}</h2>
			<p class="message-text">{{.Text}}</p>
		</div>
	{{end}}</div>
	</body>
</html>`))

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var msgs []*Message
	// TODO: Message Kindから作成時間が新しい順に10件取得するためのクエリを作成
	if /* TODO: クエリに該当するEntityをすべて取得する */ ; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := indexTmpl.Execute(w, msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
