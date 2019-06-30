// STEP03: パッケージを分けてみよう
package greeting

import (
	// fmtパッケージのインポート
	"fmt"
	// timeパッケージのインポート
	"time"
)

// Do関数の定義
// パッケージ外からアクセスできるように関数名を大文字から始める
//
// 以下のように時間よってメッセージを変える
// 04時00分 〜 09時59分: おはよう
// 10時00分 〜 16時59分: こんにちは
// 17時00分 〜 03時59分: こんばんは
func Do() {
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
