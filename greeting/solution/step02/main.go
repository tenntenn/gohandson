// STEP02: 標準パッケージを使ってみよう

// mainパッケージの定義
package main

import (
	// fmtパッケージのインポート
	"fmt"
	// timeパッケージのインポート
	"time"
)

// main関数から実行される
func main() {
	// greet関数を呼び出す
	greet()
}

// greet関数の定義
// 以下のように時間よってメッセージを変える
// 04時00分 〜 09時59分: おはよう
// 10時00分 〜 16時59分: こんにちは
// 17時00分 〜 03時59分: こんばんは
func greet() {
	// 現在時刻から何時かを取得
	h := time.Now().Hour()
	switch {
	case h >= 4 && h <= 9:
		fmt.Println("おはよう")
	case h >= 10 && h <= 16:
		fmt.Println("こんにちは")
	default:
		fmt.Println("こんばんは")
	}
}
