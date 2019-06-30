package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart"
)

// HTTPハンドラを集めた型
type Handlers struct {
	ab *AccountBook
}

// Handlersを作成する
func NewHandlers(ab *AccountBook) *Handlers {
	return &Handlers{ab: ab}
}

// ListHandlerで仕様するテンプレート
var listTmpl = template.Must(template.New("list").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8"/>
		<title>家計簿</title>
	</head>
	<body>
		<h1>家計簿</h1>
		<h2>入力</h2>
		<form method="post" action="/save">
			<label for="category">品目</label>
			<input name="category" type="text">
			<label for="price">値段</label>
			<input name="price" type="number">
			<input type="submit" value="保存">
		</form>

		<h2>最新{{len .}}件(<a href="/summary">集計</a>)</h2>
		{{- if . -}}
		<table border="1">
			<tr><th>品目</th><th>値段</th></tr>
			{{- range .}}
			<tr><td>{{.Category}}</td><td>{{.Price}}円</td></tr>
			{{- end}}
		</table>
		{{- else}}
			データがありません
		{{- end}}
	</body>
</html>
`))

// 最新の入力データを表示するハンドラ
func (hs *Handlers) ListHandler(w http.ResponseWriter, r *http.Request) {
	// 最新の10件を取得する
	items, err := hs.ab.GetItems(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 取得したitemsをテンプレートに埋め込む
	if err := listTmpl.Execute(w, items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 保存
func (hs *Handlers) SaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
		return
	}

	category := r.FormValue("category")
	if category == "" {
		http.Error(w, "品目が指定されていません", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item := &Item{
		Category: category,
		Price:    price,
	}

	if err := hs.ab.AddItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// SummaryHandlerで仕様するテンプレート
var summaryTmpl = template.Must(template.New("summary").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8"/>
		<title>家計簿 集計</title>
	</head>
	<body>
		<h1>集計</h1>
		{{- if . -}}
		<img src="/chart?w=400&h=300
			{{- /* 値 */ -}}
			&v={{- range . -}}{{.Sum}},{{- end -}}
			{{- /* ラベル */ -}}
			&l={{- range . -}}{{.Category}},{{- end -}}
		">
		<table border="1">
			<tr><th>品目</th><th>合計</th><th>平均</th></tr>
			{{- range .}}
			<tr><td>{{.Category}}</td><td>{{.Sum}}円</td><td>{{.Avg}}円</tr>
			{{- end}}
		</table>
		{{- else}}
			データがありません
		{{- end}}

		<div><a href="/">一覧に戻る</a></div>
	</body>
</html>`))

// 集計を表示するハンドラ
func (hs *Handlers) SummaryHandler(w http.ResponseWriter, r *http.Request) {
	summaries, err := hs.ab.GetSummaries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 取得した集計結果をテンプレートに埋め込む
	if err := summaryTmpl.Execute(w, summaries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// グラフを生成するハンドラ
func (hs *Handlers) ChartHandler(w http.ResponseWriter, r *http.Request) {

	// 幅を取得する
	width, err := strconv.Atoi(r.FormValue("w"))
	if err != nil {
		// デフォルトの幅
		width = 200
	}

	// 高さを取得する
	height, err := strconv.Atoi(r.FormValue("h"))
	if err != nil {
		// デフォルトの高さ
		height = 200
	}

	// 値を入れるスライス
	var vs []float64
	// カンマ区切り（末尾のカンマは無視）を分解し、float64に変換する
	for _, s := range strings.Split(strings.TrimRight(r.FormValue("v"), ","), ",") {
		// 64ビットの浮動小数点数として文字列をパースする
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		vs = append(vs, v)
	}

	// ラベルを入れるスライス
	var ls []string
	// カンマ区切り（末尾のカンマは無視）を分解し、float64に変換する
	for _, l := range strings.Split(strings.TrimRight(r.FormValue("l"), ","), ",") {
		ls = append(ls, l)
	}

	// 値とラベルの数が一致しないとエラー
	if len(vs) != len(ls) {
		http.Error(w, "値とラベルの数が一致しません", http.StatusBadRequest)
		return
	}

	pie := chart.PieChart{
		Width:  width,
		Height: height,
	}

	for i := range vs {
		v := chart.Value{
			Value: vs[i],
			Label: ls[i],
		}
		pie.Values = append(pie.Values, v)
	}

	w.Header().Set("Content-Type", chart.ContentTypePNG)
	if err := pie.Render(chart.PNG, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
