package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Category string
	Price    int
}

// 家計簿の処理を行う型
type AccountBook struct {
	fileName string
}

// 新しいAccountBookを作成する
func NewAccountBook(fileName string) *AccountBook {
	// AccountBookのポインタを返す
	return &AccountBook{fileName: fileName}
}

// ファイルに新しいItemを追加する
func (ab *AccountBook) AddItem(item *Item) error {

	// TODO: 追記用でファイルを開く
	if err != nil {
		return err
	}

	// 「品目 値段」の形式でファイルに出力する
	if _, err := /* TODO: fmt.Printlnを使って出力 */; err != nil {
		return err
	}

	// TODO: ファイルを閉じる

	return nil
}

// 最近追加したものを最大limit件だけItemを取得する
// エラーが発生したら第2戻り値で返す
func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {

	// 読み込み用でファイルを開く
	file, err := os.Open(ab.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close() // 関数終了時にCloseが呼び出される

	scanner := bufio.NewScanner(file)
	var items []*Item

	// 1行ずつ読み込む
	for scanner.Scan() {
		var item Item

		// 1行ずつパースする
		if err := ab.parseLine(scanner.Text(), &item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	// limit件より少ない場合は全件返す
	if len(items) < limit {
		return items, nil
	}

	// TODO: itemsの後方limit件だけを返す
}

// 1行ずつパースを行う
func (ab *AccountBook) parseLine(line string, /* TODO: 引数を追加する */) error {
	// 1行をスペースで分割する
	splited := strings.Split(line, " ")
	// 2つに分割できなかった場合はエラー
	if len(splited) != 2 {
		// エラーを生成して返す
		return errors.New("パースに失敗しました")
	}

	// 1つめが品目
	category := splited[0]

	// 2つめが値段
	// string型をint型に変換する
	price, err := strconv.Atoi(splited[1])
	if err != nil {
		return err
	}

	item.Category = category
	item.Price = price

	return nil
}
