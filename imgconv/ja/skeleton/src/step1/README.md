# STEP 1: go installしてみよう

コマンドラインツールは、`cmd`ディレクトリを作り、その下にコマンド名のディレクトリを作ります。
そして、その下に`main`パッケージの`go`ファイルを置きます。

```
$GOPATH
└── src
    └── step1
        └── cmd
            └── imgconv
                └── main.go
```

`GOPATH`は、Javaのクラスパスのようなもので、`import`する際に設定された`GOPATH`以下からパッケージを探します。
`GOPATH`以下には、`src`、`bin`、`pkg`などがあり、`src`にはソースコードが、`bin`にはコンパイル済のバイナリが、`pkg`にはコンパイル済のパッケージが入っています。

`go get`を行うことで、指定したリポジトリからソースコードを取得し、`GOPATH`以下へパッケージやバイナリをインストールします。
すでにソースコードがある場合は、`go install`でインストールすることができます。

さて、`GOPATH`がこのハンズオンの`src`よりひとつ上のディレクトリに設定してあった場合、以下のように実行すると`Step1`のバイナリがインストールできます。

```
$ go install step1/cmd/imgconv
$ $GOPATH/bin/imgconv
hello
```
