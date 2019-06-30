// STEP06: テストのパッケージを分けよう
package greeting_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/tenntenn/gohandson/greeting"
	"golang.org/x/text/language"
)

// Greeting.Doメソッドのテスト
func TestGreeting_Do(t *testing.T) {
	// 言語を日本語にしておき、関数実行時に元に戻す
	defer greeting.ExportSetLang(language.Japanese)()

	// greeting.Greeting型の値を作る
	g := greeting.Greeting{
		Clock: greeting.ClockFunc(func() time.Time {
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
