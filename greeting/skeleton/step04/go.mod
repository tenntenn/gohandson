// 以下の手順で作成
//   $ export GO111MODULE=on
//   $ go mod init github.com/tenntenn/gohandson/greeting
module github.com/tenntenn/gohandson/greeting

go 1.12

// 依存するモジュールを記述する
require (
	github.com/tenntenn/greeting/v2 v2.1.0
	golang.org/x/text v0.3.2 // indirect
)
