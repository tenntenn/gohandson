# STEP04: 外部パッケージを使ってみよう

## 新しく学ぶこと

* サードパーティパッケージの使い方
* Go Modules

## 外部パッケージを取得する

```sh
$ export GO111MODULE=on # 1度だけ
$ go get github.com/tenntenn/greeting/v2/text
```

## 動かし方

```sh
$ export GOBIN=`pwd`/_bin # 1度だけ
$ go install github.com/tenntenn/gohandson/greeting/cmd/greeting
$ _bin/step04
```
