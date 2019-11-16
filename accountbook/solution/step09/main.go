// STEP09: 一覧ページの作成

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/tenntenn/sqlite"
)

func main() {

	// データベースへ接続
	// ドライバにはSQLiteを使って、
	// accountbook.dbというファイルでデータベース接続を行う
	db, err := sql.Open(sqlite.DriverName, "accountbook.db")
	if err != nil {
		log.Fatal(err)
	}

	// AccountBookをNewAccountBookを使って作成
	ab := NewAccountBook(db)

	// テーブルを作成
	if err := ab.CreateTable(); err != nil {
		log.Fatal(err)
	}

	// HandlersをNewHandlersを使って作成
	hs := NewHandlers(ab)

	// ハンドラの登録
	http.HandleFunc("/", hs.ListHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
