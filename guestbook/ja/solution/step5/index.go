package main

import (
	"html/template"
	"net/http"
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
	</body>
</html>`))

func index(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
