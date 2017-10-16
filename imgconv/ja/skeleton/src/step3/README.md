# STEP 3: ファイルを扱おう

コマンドラインツールでファイルを扱うことは少なくないでしょう。
STEP 3では、コマンドライン引数で指定されたテキストファイルを開き、
行ごとに加工して、コマンドライン引数でもう一つ指定されたテキストファイルへと書きだすコマンドを作ります。

ここでは、Goでファイルを読みこんだり、書き込む方法について説明します。

## ファイルを開いて閉じる
`os.Open`を使うと読み込み専用のファイルが開けます。
また、`os.Create`を使うと読み書き可能なファイルを作ることができます。
どちらの関数もファイルを開くことができると、`*os.File`型の値が返って来ます。
なお、ファイルが開けなかったりすると、`error`が返ってくるので、適切に処理をします。

開いたファイルは、必要がなくなったら`Close`メソッドを使って閉じます。
`defer`を使って、`defer f.Close()`のように関数の終わりに閉じる場合が多いでしょう。
なお、`defer`は関数の遅延実行を行う機能で、`defer`の後ろに書いた関数呼び出しは、
実行中の関数が`return`される直前に呼ばれます。
1つの関数内で複数`defer`を書いた場合は、最後に記述したものから実行されます。

## ファイル型
`os.File`型はファイルを表す型です。
そのポインタ型の`*os.File`型が`io.Writer`インタフェースと`io.Reader`インタフェースを実装しています。
多くの標準パッケージで、これらのインタフェースを引数に取ったり、戻りに返したりします。
とくに、`io`パッケージ、`bytes`パッケージ、`bufio`パッケージ、`encoding`パッケージなどで多用されているので、
一度ドキュメントを読むと良いでしょう。


## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/skeleton
$ GOPATH=`pwd`
$ go install step3/cmd/imgconv
$ echo foo > input.txt
$ echo bar >> input.txt
$ cat input.txt
foo
bar
$ ./bin/imgconv input.txt output.txt
$ cat output.txt
1:foo
2:bar
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
