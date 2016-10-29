# STEP 8: 複数のファイルを処理しよう

STEP 7までは、ひとつの画像ファイルしか加工できませんでした。
STEP 8では、ディレクトリが指定された場合に、そのディレクトリ以下にある
すべての画像ファイルを指定した方法で加工します。

対象とする画像ファイルの形式は、`format`フラグとして指定もできますが、
省略した場合は、`png`と`jpeg`画像が対象となります。

出力するファイル名は、`"{{.Dir}}/{{.Name}}_small{{.Ext}}"`のようにテンプレートで指定することができるようにします。
`.Dir`は入力ファイルが入っているディレクトリで、`.Name`は拡張子を除いた名前、`.Ext`は拡張子を表します。
実際には、`"a/b/gopher.png"`というファイルを対象とした場合、上記の例だと`"a/b/gopher_small.png"`という名前で書き出されます。

ここでは、ディレクトリ内を再帰的に探索するために使用する`path/filepath`パッケージの説明や出力ファイル名を決定するためのテンプレートエンジンを提供する`text/template`パッケージについて説明します。

また、ファイル処理関係のエラーについて`os`パッケージで提供されている関数を使って適切に処理する方法について説明します。

## path/filepath.Walk

`path/filepath`パッケージは、現在のディレクトリを元にして絶対パスを求めたり、
`Walk`でディレクトリ内を再帰的に読み込むことができたりする機能を提供しています。

`filepath.Walk`関数は、指定したディレクトリを再帰的に潜っていきます。
ディレクトリ内のファイルごとに第2引数で渡した`filepath.WalkFunc`を実行します。
なお、関数リテラルから関数型の値へはキャストする必要がありませんので、通常は第2引数に直接関数リテラルを書きます。
関数リテラルとは、`func() { fmt.Println("hello") }`のような無名関数のことを表します。

## text/templateパッケージ

`text/template`パッケージはテンプレートエンジンを提供するパッケージです。
コマンドラインツールでも引数にテンプレートを取ることが多いかと思います。
なお、Webアプリでテンプレートエンジンを使う場合は、
HTMLのエスケープ処理などを行ってくれる`html/template`の方を使うと良いでしょう。

テンプレートに、データを埋め込むには、`*template.Template`型の`Execute`メソッドを使います。
テンプレート内では、埋め込んだ値は`{{.}}`でアクセスできます。
なお、埋め込んだ値のフィールドやメソッドにアクセスすることもできます。

```
type MyString string

func (s MyString) ToLower() string {
    return strings.ToLower(string(s))
}

// tは*template.Template型の値
t, _ := template.New("MyTemplate").Parse(`{{.}} {{.ToLower}}`)
// rには、"Hello, Gophers hello, gophers"と出力される
t.Execute(r, MyString("Hello, Gophers"))
```

なお上記の例では、`string`型に新しく`MyString`という型名をつけ、その型にメソッドを設けています。
このように、Goでは構造体以外の型にも、パッケージ内で`type`で宣言していればメソッドを設けることができます。

## osパッケージのエラー処理

`os`パッケージが提供する関数では、最後の戻り値に`error`型の値を返すことが多いです。
`os`パッケージでは、これらのエラーの種類を知るために、`os.IsNotExist`関数や`os.IsExist`関数などを提供しています。
ファイル関係のエラーを受け取った際に、これらの関数を使えば適切にエラー処理をすることができるでしょう。

## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/solution
$ GOPATH=`pwd`
$ go get golang.org/x/image/draw
$ go install step8/cmd/imgconv
$ go install tools/cmd/httpget
$ mkdir a
$ mkdir a/b
$ mkdir a/c
$ ./bin/httpget https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/hi.png > a/b/gopher.png
$ cp a/b/gopher.png a/c/
$ ./bin/imgconv -format png -resize 50%x50% a "{{.Dir}}/{{.Name}}_small{{.Ext}}"
$ tree a
a
├── b
│   ├── gopher.png
│   └── gopher_small.png
└── c
    ├── gopher.png
    └── gopher_small.png
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
