# 〜家計簿アプリを作ろう〜

## ハンズオンのやりかた

`skeleton`ディレクトリ以下に問題があり、10個のステップに分けられています。
STEP01からSTEP10までステップごとに進めていくことで、GoでWebアプリが作れるようになっています。

各ステップに、READMEが用意されていますので、まずは`README`を読みます。
`README`には、そのステップを理解するための解説が書かれています。

`README`を読んだら、ソースコードを開き`TODO`コメントが書かれている箇所をコメントに従って修正して行きます。
`TODO`コメントをすべて修正し終わったら、`README`に書かれた実行例に従ってプログラムをコンパイルして実行します。

途中でわからなくなった場合は、`solution`ディレクトリ以下に解答例を用意していますので、そちらをご覧ください。

`macOS`の動作結果をもとに解説しています。
`Windows`の方は、パスの区切り文字やコマンド等を適宜読み替えてください。

## 目次

* STEP01: Goに触れる
* STEP02: データの入力
* STEP03: データの記録
* STEP04: 複数データの記録
* STEP05: データベースへの記録
* STEP06: 品目ごとの集計
* STEP07: ブラッシュアップ
* STEP08: 一覧ページの作成
* STEP09: 入力ページの作成
* STEP10: 集計ページの作成

## ソースコードの取得

```
$ go env GOPATH
$ cd ↑のディレクトリに移動
$ mkdir -p src/github.com/tenntenn/
$ cd src/github.com/tenntenn
$ git clone https://github.com/tenntenn/gohandson.git
$ cd gohandson/accountbook
```
