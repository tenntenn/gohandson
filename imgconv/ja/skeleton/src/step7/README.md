# STEP 7: 画像を縮小／拡大しよう

STEP 7では、さらにコマンドに機能を追加し、画像の縮小／拡大を行えるようにします。
拡大するサイズは、`resize`という名前でフラグとしてコマンドライン引数で指定します。
縮小／拡大を行うには、STEP 6で使用した`image/draw`パッケージの代わりに、`golang.org/x/image/draw`パッケージを使用します。

ここでは、`golang.org/x/image`パッケージの説明と外部パッケージの`go get`の方法について説明します。

## golang.org/x/imageパッケージ

`golang.org/x/image`パッケージは、標準パッケージの`image`パッケージより機能を増やしたパッケージです。
特に`draw`パッケージは、スケールなどの機能が追加されています。

`golang.org/x`以下にあるパッケージは、サブプロジェクトとしてGoチームによって保守されています。
`image`以外にも、`golang.org/x/net`など便利なパッケージがありますので、一度覗いてみると良いでしょう。

このステップでは、`golang.org/x/image/draw`パッケージを使用しますので、`go get`しておきましょう。
以下の通り、`go get`を実行すると、`src`と`pkg`以下に`golang.org/x/image/draw`がインストールされている事が分かります。

```
$ go get golang.org/x/image/draw
$ tree pkg
pkg
└── darwin_amd64
    └── golang.org
        └── x
            └── image
                ├── draw.a
                └── math
                    └── f64.a
$ ls src/golang.org/x/image/
AUTHORS         LICENSE         bmp             colornames      font            testdata        vp8l
CONTRIBUTING.md PATENTS         cmd             draw            math            tiff            webp
CONTRIBUTORS    README          codereview.cfg  example         riff            vp8
```

## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/skeleton
$ GOPATH=`pwd`
$ go install step7/cmd/imgconv
$ go install tools/cmd/httpget
$ ./bin/httpget https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/hi.png > gopher.png
$ ./bin/imgconv -resize 50%x50% gopher.png gopher2.png
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
