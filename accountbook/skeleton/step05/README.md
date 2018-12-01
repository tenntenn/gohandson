# STEP05: データベースへの記録

## 新しく学ぶこと

* サードパーティパッケージの使い方
* エラー処理
* database/sqlパッケージの使い方
 * テーブルの作成
 * INSERT
 * SELECT
* defer
* fmt.Printfの%04d

## SQLiteのライブラリのための準備

参考: https://github.com/mattn/go-sqlite3#compilation

### macOS

homebrewでSQLiteをいれる。

```
$ brew install sqlite3
```

### Windows

Cのコンパイラが必要なため、gccを入れる。

https://sourceforge.net/projects/tdm-gcc/

## 動かし方

```
$ go build -v -o step05
$ ./step05
```
