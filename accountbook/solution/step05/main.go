package main

import (
	"database/sql"
	"fmt"
	"log"

	// SQLiteのドライバを使うためにインポートするが直接は使わない
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	// IDはデータベースに記録した際に振られるID
	ID       int
	Category string
	Price    int
}

func main() {

	// データベースへ接続
	// ドライバにはSQLiteを使って、
	// accountbook.dbというファイルでデータベース接続を行う
	db, err := sql.Open("sqlite3", "accountbook.db")
	if err != nil {
		log.Fatal(err)
	}

	// テーブルを作成（なければ）する
	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	// 入力
	for i := 0; i < n; i++ {
		if err := inputItem(db); err != nil {
			log.Fatal(err)
		}
	}

	// 一覧の出力
	if err := showItems(db); err != nil {
		log.Fatal(err)
	}
}

// テーブルの作成
// SQLのCREATE文を使ってテーブルを作成する
// エラーが発生した場合にはそのまま返す
func createTable(db *sql.DB) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS items(
		id        INTEGER PRIMARY KEY,
		category  TEXT NOT NULL,
		price     INTEGER NOT NULL
	);`

	_, err := db.Exec(sqlStr)
	if err != nil {
		return err
	}

	return nil
}

// 入力を行いデータベースに保存する
// エラーが発生した場合にはそのまま返す
func inputItem(db *sql.DB) error {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	// SQLのINSERTを使ってデータベースに保存する
	// ?の部分にcategoryやpriceの値が来る
	const sqlStr = `INSERT INTO items(category, price) VALUES (?,?);`
	_, err := db.Exec(sqlStr, item.Category, item.Price)
	if err != nil {
		return err
	}

	return nil
}

// 一覧の表示
func showItems(db *sql.DB) error {

	// SELECTでitemsテーブルのすべて行を取得する
	const sqlStr = `SELECT * FROM items`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	fmt.Println("===========")
	// 1つずつ取得した行をみる
	// rows.Nextはすべての行を取得し終わるとfalseを返す
	for rows.Next() {
		var item Item
		// rows.Scanで取得した行からデータを取り出し、itemの各フィールドに入れる
		err := rows.Scan(&item.ID, &item.Category, &item.Price)
		if err != nil {
			return err
		}
		fmt.Printf("[%04d] %s:%d円\n", item.ID, item.Category, item.Price)
	}
	fmt.Println("===========")

	err := rows.Err()
	if err != nil {
		return err
	}

	return nil
}
