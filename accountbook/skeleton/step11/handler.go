package main

import (
	"html/template"
	"net/http"
	"strconv"
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
		<script src="https://www.gstatic.com/charts/loader.js"></script>
		<script>
			google.charts.load('current', {'packages':['corechart']});
			google.charts.setOnLoadCallback(drawChart);

			function drawChart() {
			var data = google.visualization.arrayToDataTable([
				['品目', '値段'],
				{{- range . -}}
				['{{js .Category}}', {{- /* TODO: 合計を埋め込む */ -}}],
				{{- end -}}
			]);
		
		var options = { title: '割合' };
		var chart = new google.visualization.PieChart(document.getElementById('piechart'));
		chart.draw(data, options);
		}
		</script>
	</head>
	<body>
		<h1>集計</h1>
		{{- if . -}}
		<div id="piechart" style="width:400px; height:300px;"></div>
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
	// TODO: 集計結果を取得し、summariesに入れる
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: 取得した集計結果をテンプレートに埋め込む
	if /* ここに書く */; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
