# STEP 1: go installしてみよう

STEP 1では、簡単なコマンドを`go install`でビルドし、`GOPATH`以下にインストールする方法について説明します。

コマンドラインツールを作るには、`cmd`ディレクトリを作り、その下にコマンド名のディレクトリを作ります。
そして、その下に`main`パッケージの`go`ファイルを置きます。

```
$GOPATH
└── src
    └── step1
        └── cmd
            └── typing
                └── main.go
```

`GOPATH`は、Javaのクラスパスのようなもので、`import`する際に設定された`GOPATH`以下からパッケージを探します。
`GOPATH`以下には、`src`、`bin`、`pkg`などがあり、`src`にはソースコードが、`bin`にはコンパイル済のバイナリが、`pkg`にはコンパイル済のパッケージが入っています。

`go get`を行うことで、指定したリポジトリからソースコードを取得し、`GOPATH`以下へパッケージやバイナリをインストールします。
すでにソースコードがある場合は、`go install`でインストールすることができます。

さて、`GOPATH`がこのハンズオンの`src`よりひとつ上のディレクトリに設定してあった場合、
以下のように実行すると`step1/cmd/typing/main.go`がビルドされ、
`GOPATH`以下の`bin`ディレクトリにバイナリがインストールされます。

```
$ go install step1/cmd/typing
$ $GOPATH/bin/typing
hello
```

## 目次

* STEP 1: [go installしてみよう](../step1)（[解答例](../../../solution/src/step1)）
