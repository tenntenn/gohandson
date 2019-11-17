package main

import (
	"database/sql"
)

type Item struct {
	ID       int
	Category string
	Price    int
}

// 家計簿の処理を行う型
type AccountBook struct {
	db *sql.DB
}

// 新しいAccountBookを作成する
func NewAccountBook(db *sql.DB) *AccountBook {
	// AccountBookのポインタを返す
	return &AccountBook{db: db}
}

// テーブルがなかったら作成する
func (ab *AccountBook) CreateTable() error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS items(
		id        INTEGER PRIMARY KEY,
		category  TEXT NOT NULL,
		price     INTEGER NOT NULL
	);`

	_, err := ab.db.Exec(sqlStr)
	if err != nil {
		return err
	}

	return nil
}

// データベースに新しいItemを追加する
func (ab *AccountBook) AddItem(item *Item) error {
	const sqlStr = `INSERT INTO items(category, price) VALUES (?,?);`
	_, err := ab.db.Exec(sqlStr, item.Category, item.Price)
	if err != nil {
		return err
	}
	return nil
}

// 最近追加したものを最大limit件だけItemを取得する
// エラーが発生したら第2戻り値で返す
func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {
	// ORDER BY id DESCでidの降順（大きい順）=最近追加したものが先にくる
	// LIMITで件数を最大の取得する件数を絞る
	const sqlStr = `SELECT * FROM items ORDER BY id DESC LIMIT ?`
	rows, err := ab.db.Query(sqlStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	var items []*Item
	// 1つずつ取得した行をみる
	// rows.Nextはすべての行を取得し終わるとfalseを返す
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Category, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
