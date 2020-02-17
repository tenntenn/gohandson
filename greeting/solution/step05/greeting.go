// STEP05: テストを書いてみよう
package greeting

import (
	// fmtパッケージのインポート
	"fmt"
	"io"
	// timeパッケージのインポート
	"time"

	// textパッケージのインポート
	"github.com/tenntenn/greeting/v2/text"
)

// デフォルトの言語
var lang = text.DefaultLang()

// 時刻の取得を抽象化したインタフェース
type Clock interface {
	Now() time.Time
}

// 時刻を返すような関数をClockFunc型として定義
type ClockFunc func() time.Time

// 関数にClockインタフェースを実装させる
func (f ClockFunc) Now() time.Time {
	// レシーバは関数なのでそのまま呼び出す
	return f()
}

// 挨拶を行うための構造体型
type Greeting struct {
	// Clockインタフェースをフィールドに持つことで
	// 時刻の取得を抽象化する
	Clock Clock
}

// 現在時刻を取得する
// Clockフィールドがnilの場合はtime.Now()の値を使う
// nilじゃない場合はg.Clock.Now()の値を使う
func (g *Greeting) now() time.Time {
	if g.Clock == nil {
		// time.Now()の値を使う
		return time.Now()
	}
	// g.Clock.Now()の値を使う
	return g.Clock.Now()
}

// Do関数の定義
// パッケージ外からアクセスできるように関数名を大文字から始める
// 引数にio.Writerを取ることで出力先を自由に変えることができる
//
// 以下のように時間よってメッセージを変える
// 04時00分 〜 09時59分: おはよう
// 10時00分 〜 16時59分: こんにちは
// 17時00分 〜 03時59分: こんばんは
func (g *Greeting) Do(w io.Writer) error {
	h := g.now().Hour()
	var msg string
	switch {
	case h >= 4 && h <= 9:
		msg = text.GoodMorning(lang)
	case h >= 10 && h <= 16:
		msg = text.Hello(lang)
	default:
		msg = text.GoodEvening(lang)
	}

	_, err := fmt.Fprint(w, msg)
	if err != nil {
		return err
	}

	return nil
}
