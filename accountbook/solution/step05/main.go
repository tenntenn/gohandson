package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID       int
	Category string
	Price    int
}

func main() {

	// データベースへ接続
	db, err := sql.Open("sqlite3", "accountbook.db")
	if err != nil {
		log.Fatal(err)
	}

	// テーブルを作成（なければ）
	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if err := inputItem(db); err != nil {
			log.Fatal(err)
		}
	}

	if err := showItems(db); err != nil {
		log.Fatal(err)
	}
}

// テーブルの作成
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

// 入力
func inputItem(db *sql.DB) error {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	const sqlStr = `INSERT INTO items(category, price) VALUES (?,?);`
	_, err := db.Exec(sqlStr, item.Category, item.Price)
	if err != nil {
		return err
	}

	return nil
}

// 一覧の表示
func showItems(db *sql.DB) error {
	const sqlStr = `SELECT * FROM items`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	fmt.Println("===========")
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Category, &item.Price); err != nil {
			return err
		}
		fmt.Printf("[%04d] %s:%d円\n", item.ID, item.Category, item.Price)
	}
	fmt.Println("===========")

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
