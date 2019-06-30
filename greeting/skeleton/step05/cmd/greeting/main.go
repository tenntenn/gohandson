// STEP05: テストを書いてみよう

// mainパッケージの定義
package main

import (
	"os"

	// greetingパッケージのインポート
	"github.com/tenntenn/gohandson/greeting"
)

// main関数から実行される
func main() {
	// Greeting型の変数を定義
	var g greeting.Greeting
	// 引数にos.Stdoutを渡してDoメソッドを呼び出す
	g.Do(os.Stdout)
}
