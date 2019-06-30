// STEP08: テーブル駆動テストを行おう
package greeting_test

import (
	"bytes"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/tenntenn/gohandson/greeting"
	"golang.org/x/text/language"
)

// "YYYY/MM/DD hh:mm:ss" 形式の時刻を返すようなgreeting.Clockを作る
// 引数にtesting.Tと文字列で表した時刻を取得する
func mockClock(t *testing.T, v string) greeting.Clock {
	// ヘルパーであることを明示する
	t.Helper()
	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		// エラーが発生した場合はテスト中断させエラーにする
		t.Fatal("unexpected error:", err)
	}

	// nowを返す関数を作り、greeting.ClockFuncにキャストして返す
	return greeting.ClockFunc(func() time.Time {
		return now
	})
}

// 設定されたエラーを返すWriter
type errorWriter struct {
	Err error
}

// フィールドにエラーが設定されていたらそのエラーを返す
func (w *errorWriter) Write(p []byte) (n int, err error) {
	return 0, w.Err
}

// Greeting.Doメソッドのテスト
func TestGreeting_Do(t *testing.T) {
	// 言語を日本語にしておく
	defer greeting.ExportSetLang(language.Japanese)()

	cases := map[string]struct {
		writer io.Writer
		clock  greeting.Clock

		msg       string
		expectErr bool
	}{
		"04時": {new(bytes.Buffer), mockClock(t, "2018/08/31 04:00:00"), "おはよう", false},
		"09時": {new(bytes.Buffer), mockClock(t, "2018/08/31 09:00:00"), "おはよう", false},
		"10時": {new(bytes.Buffer), mockClock(t, "2018/08/31 10:00:00"), "こんにちは", false},
		"16時": {new(bytes.Buffer), mockClock(t, "2018/08/31 16:00:00"), "こんにちは", false},
		"17時": {new(bytes.Buffer), mockClock(t, "2018/08/31 17:00:00"), "こんばんは", false},
		"03時": {new(bytes.Buffer), mockClock(t, "2018/08/31 03:00:00"), "こんばんは", false},
		"エラー": {&errorWriter{Err: errors.New("error")}, nil, "", true},
	}

	for n, tt := range cases {
		tt := tt
		t.Run(n, func(t *testing.T) {
			g := greeting.Greeting{
				Clock: tt.clock,
			}

			switch err := g.Do(tt.writer); true {
			// エラーを期待してるのにエラーが発生していない
			case err == nil && tt.expectErr:
				t.Error("expected error did not occur")
			// エラーは期待してないのにエラーが発生した
			case err != nil && !tt.expectErr:
				t.Error("unexpected error:", err)
			}

			// tc.writerが*bytes.Bufferだったら値もチェック
			if buff, ok := tt.writer.(*bytes.Buffer); ok {
				msg := buff.String()
				if msg != tt.msg {
					t.Errorf("greeting msg wont %s but got %s", tt.msg, msg)
				}
			}
		})
	}
}
