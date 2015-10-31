# STEP 2: コマンドライン引数を取ろう

コマンドラインツールは、コマンド実行時にいくつかの引数を取ることが多いでしょう。
STEP 2では、コマンドライン引数で受け取った文字列を標準出力に出力して見ます。

ここでは、コマンドライン引数の取り方と出力に使用する`fmt`パッケージについて説明します。

## コマンドライン引数
`os.Args`は、コマンドライン引数が入った`string`型のスライスです。
以下のように、コマンドを実行した場合、`["imgconv", "input.png", "output.png"]`のような値が入ります。

```
$ imgconv input.png output.png
```

## fmtパッケージ
`fmt.Println`や`fmt.Fprintf`などを提供するパッケージです。

`fmt.Fprintf`は、第1引数に`io.Writer`インタフェースを取ることができ、このインタフェースを実装した（`Writer`メソッドを持つ）型であれば、どのような型に対しても書式を指定して出力できます。

`fmt.Errorf`は、指定した書式でエラーメッセージを記述し、そのエラーメッセージを`Error`メソッドで返す`error`インタフェース型の値を返します。

## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/skeleton
$ GOPATH=`pwd`
$ go install step2/cmd/imgconv
$ ./bin/imgconv input.txt output.txt
./bin/imgconv
input.txt
output.txt
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
