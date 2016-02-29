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
/path/to/gohandson/imgconv/ja/solution
$ GOPATH=`pwd`
$ go install step2/cmd/imgconv
$ ./bin/imgconv input.txt output.txt
./bin/imgconv
input.txt
output.txt
```

## 目次

* STEP 1: [go installしてみよう](../../../skeleton/src/step1)（[解答例](../step1)）
* STEP 2: [コマンドライン引数を取ろう](../../../skeleton/src/step2)（[解答例](../step2)）
* STEP 3: [ファイルを扱おう](../../../skeleton/src/step3)（[解答例](../step3)）
* STEP 4: [画像形式を変換しよう](../../../skeleton/src/step4)（[解答例](../step4)）
* STEP 5: [`flag`パッケージを使おう](../../../skeleton/src/step5)（[解答例](../step5)）
* STEP 6: [画像を切り抜こう](../../../skeleton/src/step6)（[解答例](../step6)）
* STEP 7: [画像を縮小／拡大しよう](../../../skeleton/src/step7)（[解答例](../step7)）
* STEP 8: [複数のファイルを処理しよう](../../../skeleton/src/step8)（[解答例](../step8)）
