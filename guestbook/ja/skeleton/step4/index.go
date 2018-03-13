package main

import (
	"html/template"
	"net/http"
)

const limitMessages = 10

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>ゲストブック</title>
	</head>
	<body>
		{{.}}
	</body>
</html>`))

func index(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	if msg == "" {
		msg = "NO MESSAGE"
	}
	if /* TODO: テンプレートからHTMLを生成し、レスポンスとして返す */; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
