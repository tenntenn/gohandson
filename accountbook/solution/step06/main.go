// STEP06: 品目ごとの集計

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

	// 品目ごとの集計結果を出力
	if err := showSummary(db); err != nil {
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

// 集計結果の表示
func showSummary(db *sql.DB) error {

	// GROUP BYで品目ごとにグループ化して金額の合計を出す
	const sqlStr = `
	SELECT
		category,
		COUNT(1) as count,
		SUM(price) as sum
	FROM
		items
	GROUP BY
		category`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	fmt.Println("===========")
	// タブ区切りで「品目 個数 合計 平均」を出力
	fmt.Printf("品目\t個数\t合計\t平均\n")
	// 1つずつ取得した行をみる
	// rows.Nextはすべての行を取得し終わるとfalseを返す
	for rows.Next() {
		var (
			category string
			sum      int
			count    int
		)
		// rows.Scanで取得した行からデータを取り出し変数に入れる
		err := rows.Scan(&category, &count, &sum)
		if err != nil {
			return err
		}
		avg := float64(sum) / float64(count)
		fmt.Printf("%s\t%d\t%d円\t%.2f円\n", category, count, sum, avg)
	}
	fmt.Println("===========")

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
