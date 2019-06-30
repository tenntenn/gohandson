// STEP04: 外部パッケージを使ってみよう

// mainパッケージの定義
package main

// greetingパッケージのインポート
import "github.com/tenntenn/gohandson/greeting"

// main関数から実行される
func main() {
	// greeting.Do関数を呼び出す
	greeting.Do()
}
