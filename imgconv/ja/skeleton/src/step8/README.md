# STEP 8: 複数のファイルを処理しよう

## osパッケージのエラー処理

`os`パッケージが提供する関数では、最後の戻り値に`error`型の値を返すことが多いです。
`os`パッケージでは、これらのエラーの種類を知るために、`os.IsNotExist`関数や`os.IsExist`関数などを提供しています。

## path/filepath

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
$ ./bin/httpget https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/01.png > a/b/gopher.png
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
