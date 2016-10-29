# STEP 4: 画像形式を変換しよう

STEP 4では、いよいよ画像を扱います。
ここでは、1つめのコマンドライン引数で指定された画像ファイルを開き、
2つめの引数で指定されたファイル名で画像を保存します。
このとき、拡張子を見て保存する画像形式を判断します。

ここでは画像を扱うために必要な`image`パッケージとパスを扱う`path`パッケージ、
文字列処理を行う`strings`パッケージについて説明を行います。

## imageパッケージ
`image`パッケージは、画像を扱うパッケージです。
画像を表す`image.Image`インタフェースやそれを実装する具体的な型が定義されています。

`image/png`パッケージや`image/jpeg`パッケージでは、`png`や`jpeg`形式の画像を`io.Reader`から`image.Image`にデコードしたり、`image.Image`から`io.Writer`へエンコードする機能が提供されています。


## path/filepathパッケージ
`path/filepath`パッケージは、パスに関する機能を提供しています。
例えば、`filepath.Ext`はファイル名から拡張子を取得でき、`filepath.Join`はOSごとの適切な区切り文字でパスを結合することができます。

## stringsパッケージ
`strings`パッケージは、文字列操作に関する処理を提供するパッケージです。
例えば、`strings.ToUpper`や`strings.ToLower`など大文字／小文字に変換する関数や、`strings.Join`や`strings.Split`などを文字列を結合／分割する関数が提供されています。

多くのパッケージで、引数に`io.Reader`をとっているため、`string`型から`io.Reader`を取得したい場合があります。
その場合には、`strings.NewReader`で`string`型をそのまま`io.Reader`に変換できることができます。

なお、`bytes`パッケージも`[]byte`向けに、`strings`と似たような機能を提供しています。

## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/skeleton
$ GOPATH=`pwd`
$ go install step4/cmd/imgconv
$ go install tools/cmd/httpget
$ ./bin/httpget https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/hi.png > gopher.png
$ ./bin/imgconv gopher.png gopher.jpg
```

## 目次

* STEP 1: [go installしてみよう](../step1)（[解答例](../../../solution/src/step1)）
* STEP 2: [コマンドライン引数を取ろう](../step2)（[解答例](../../../solution/src/step2)）
* STEP 3: [ファイルを扱おう](../step3)（[解答例](../../../solution/src/step3)）
* STEP 4: [画像形式を変換しよう](../step4)（[解答例](../../../solution/src/step4)）
* STEP 5: [`flag`パッケージを使おう](../step5)（[解答例](../../../solution/src/step5)）
* STEP 6: [画像を切り抜こう](../step6)（[解答例](../../../solution/src/step6)）
* STEP 7: [画像を縮小／拡大しよう](../step7)（[解答例](../../../solution/src/step7)）
* STEP 8: [複数のファイルを処理しよう](../step8)（[解答例](../../../solution/src/step8)）
