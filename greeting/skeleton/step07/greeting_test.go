// STEP07: テストヘルパーを作ってみよう
package greeting_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/tenntenn/gohandson/greeting"
	"golang.org/x/text/language"
)

// "YYYY/MM/DD hh:mm:ss" 形式の時刻を返すようなgreeting.Clockを作る
// 引数にtesting.Tと文字列で表した時刻を取得する
func mockClock( /* TODO: 引数を決める */ ) greeting.Clock {
	// TODO: ヘルパーであることを明示する

	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		// TODO: エラーが発生した場合はテスト中断させエラーにする
	}

	// TODO: nowを返す関数を作り、greeting.ClockFuncにキャストして返す
}

// Greeting.Doメソッドのテスト
func TestGreeting_Do(t *testing.T) {
	// 言語を日本語にしておく
	defer greeting.ExportSetLang(language.Japanese)()

	// greeting.Greeting型の値を作る
	g := greeting.Greeting{
		Clock: mockClock(t, "2018/08/31 06:00:00"),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
