# STEP 5: `flag`パッケージを使おう

STEP 5では、`flag`パッケージを使いコマンドライン引数からフラグを取得してみます。

コマンドライン引数を`flag`パッケージを使ってパースすると、`string`型や`bool`型の値をフラグとして受け取ることができます。

`flag.StringVar`関数や`flag.IntVar`関数は、引数にその型の変数のポインタとデフォルトの値、使い方を渡します。そして、`flag.Parse`が呼ばれると、コマンドライン引数がパースされ、第1引数で渡したポインタの指す先に値が設定されます。

フラグのパースは、`init`関数の中で行われることが多いでしょう。`init`関数は、パッケージがインポートされた際に呼ばれる関数で、`main`パッケージの場合も`main`関数が実行される前に呼ばれます。なお、`init`関数は、1つのパッケージ、1つのファイル中にいくつも書くことができます。

`flag.Args`関数を使うと、フラグとしてパースされた部分以外のコマンドライン引数を取ることができます。
`os.Args`スライスと似たような値を返しますが、`flag.Args`関数が返すスライスは、`0`番目の要素にコマンド名は含まれません。

## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/skeleton
$ GOPATH=`pwd`
$ go install step5/cmd/imgconv
$ ./bin/imgconv -h
Usage of ./bin/imgconv:
  -clip 幅[px|%]x高さ[px|%]
        切り取る画像サイズ（幅[px|%]x高さ[px|%]）
$ go install tools/cmd/httpget
$ ./bin/httpget https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/hi.png > gopher.png
$ ./bin/imgconv -clip 10x10 gopher.png gopher.jpg
切り抜きを行う予定 10x10
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
