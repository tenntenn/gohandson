# STEP 2: コマンドライン引数を取ろう

## コマンドライン引数
`os.Args`は、コマンドライン引数が入った`string`型のスライスです。。
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
