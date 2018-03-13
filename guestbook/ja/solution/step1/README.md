# STEP1: Hello, World

## HTTPハンドラ

GoではHTTPリクエストを処理するハンドラは次のように記述します。

```
func index(w http.ResponseWriter, r *http.Request) {
    // リクエストを処理する
}
```

第1引数はレスポンスを書き込むWriterで第2引数はリクエストです。

レスポンスを返すには、`ResponseWriter`に対して書き込みます。
書き込むには`fmt.Fprintln`などが使えます。

```
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Google App Engine")
}
```

## ログ

GAEでは、ログを書き込むために、`log`パッケージを利用します。
InfoやDebugなどのログレベルがあり、呼び出す関数を変えるとログレベルが変わります。

GAEでは、コンテキストと呼ばれる

## 開発サーバで確認
