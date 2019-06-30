// STEP05: テストを書いてみよう
package greeting

import (
	"bytes"
	"testing"
	"time"

	"golang.org/x/text/language"
)

// Greeting.Doメソッドのテスト
func TestGreeting_Do(t *testing.T) {
	// パッケージのlangを入れ替える
	orgLang := lang
	lang = language.Japanese
	defer func() {
		// deferで元に戻す
		lang = orgLang
	}()

	// Greeting型の値を作る
	g := Greeting{
		Clock: ClockFunc(func() time.Time {
			// 2018/08/31 06:00:00を返すようにしておく
			return time.Date(2018, 8, 31, 06, 0, 0, 0, time.Local)
		}),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
