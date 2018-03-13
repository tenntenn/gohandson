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
		{{.}}
	</body>
</html>`))

func index(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	if msg == "" {
		msg = "NO MESSAGE"
	}
	if err := indexTmpl.Execute(w, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
