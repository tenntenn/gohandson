# コマンドラインツールを作ろう

## はじめに

このハンズオンでは、画像変換を行うコマンドラインツールを作ります。
基本的な文法などは、このハンズオンでは扱いません。
そのため、ハンズオンを始める前に、必ず [A Tour of Go](https://go-tour-jp.appspot.com)を終わらせてください。
また、文法や周辺ツール、`GOPATH`などの詳しい説明は[公式ドキュメント](https://golang.org/doc/)の「Learning Go」の項目を読んでください。
特に他のオブジェクト指向言語などを学習されている方は、「FAQ」に目を通すとよいでしょう。

英語が辛い方は、有志によって[翻訳されたドキュメント](http://golang-jp.org/doc/)の「Goを学ぶ」を読んで下さい。
すべてが翻訳されているわけではありませんが、役に立つでしょう。
また、[はじめてのGo](http://gihyo.jp/dev/feature/01/go_4beginners)も日本語で書かれていて分かりやすいのでぜひ読んで下さい。

標準パッケージについては、[パッケージドキュメント](https://golang.org/pkg/)を見ると、使い方が説明されています。

このハンズオンを行うには、`GOPATH`をこの`README`があるディレクトリ以下の`skeleton`ディレクトリに設定してください。

MacやLinuxの場合：

```
$ GOPATH=`pwd`/skeleton
```

Windowsの場合：

```
TODO
```

## 目次

* STEP 1: [go installしてみよう](./skeleton/src/step1)（[解答例](./solution/src/step1)）
* STEP 2: [コマンドライン引数を取ろう](./skeleton/src/step2)（[解答例](./solution/src/step2)）
* STEP 3: [ファイルを扱おう](./skeleton/src/step3)（[解答例](./solution/src/step3)）
* STEP 4: [画像形式を変換しよう](./skeleton/src/step4)（[解答例](./solution/src/step4)）
* STEP 5: [`flag`パッケージを使おう](./skeleton/src/step5)（[解答例](./solution/src/step5)）
* STEP 6: [画像を切り取ろう](./skeleton/src/step6)（[解答例](./solution/src/step6)）
* STEP 7: [画像を縮小／拡大しよう](./skeleton/src/step7)（[解答例](./solution/src/step7)）
* STEP 8: [複数のファイルを処理しよう](./skeleton/src/step8)（[解答例](./solution/src/step8)）
