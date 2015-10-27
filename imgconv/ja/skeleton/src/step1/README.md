# STEP 1: go installしてみよう

STEP 1では、簡単なコマンドを`go install`でビルドし、`GOPATH`以下にインストールする方法について説明します。

コマンドラインツールを作るには、`cmd`ディレクトリを作り、その下にコマンド名のディレクトリを作ります。
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

さて、`GOPATH`がこのハンズオンの`src`よりひとつ上のディレクトリに設定してあった場合、
以下のように実行すると`step1/cmd/imgconv/main.go`がビルドされ、
`GOPATH`以下の`bin`ディレクトリにバイナリがインストールされます。

```
$ go install step1/cmd/imgconv
$ $GOPATH/bin/imgconv
hello
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
