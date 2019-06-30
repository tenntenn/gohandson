// STEP04: 外部パッケージを使ってみよう
package greeting

import (
	// fmtパッケージのインポート
	"fmt"
	// timeパッケージのインポート
	"time"

	// TODO: textパッケージのインポート
)

// デフォルトの言語
var lang = text.DefaultLang()

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
		// text.GoodMorningを使う
		fmt.Println(text.GoodMorning(lang))
	case h >= 10 && h <= 16:
		// TODO: text.Helloを使う
	default:
		// text.GoodEveningを使う
		fmt.Println(text.GoodEvening(lang))
	}
}
