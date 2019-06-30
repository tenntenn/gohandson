// STEP07: テストヘルパーを作ってみよう
package greeting

import "golang.org/x/text/language"

// パッケージ変数langを一時的に変更する関数
// greetingパッケージだがファイル名が_test.goで終わるため
// go testの際しかビルドされない
func ExportSetLang(l language.Tag) func() {
	orgLang := lang
	lang = l
	return func() {
		lang = orgLang
	}
}
