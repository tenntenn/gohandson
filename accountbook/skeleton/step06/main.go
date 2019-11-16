// STEP06: ブラッシュアップ

package main

import (
	"fmt"
	"os"
)

func main() {

	// AccountBookをNewAccountBookを使って作成
	ab := NewAccountBook("accountbook.txt")

LOOP: // 以下のループにラベル「LOOP」をつける
	for {

		// モードを選択して実行する
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]終了")
		fmt.Printf(">")
		fmt.Scan(&mode)

		// モードによって処理を変える
		switch mode {
		case 1: // 入力
			var n int
			fmt.Print("何件入力しますか>")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				if err := ab.AddItem(inputItem()); err != nil {
					// TODO: os.Stderrにエラーメッセージを出す
					break LOOP
				}
			}
		case 2: // 最新10件
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				// TODO: LOOPという名前のついたforから抜け出す
			}
			showItems(items)
		case 3: // 終了
			// TODO: 3のとき「終了します」と出力して終了する
		}
	}
}

// Itemを入力し返す
func inputItem() *Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return &item
}

// Itemの一覧を出力する
func showItems(items []*Item) {
	fmt.Println("===========")
	// itemsの要素を1つずつ取り出してitemに入れて繰り返す
	for _, item := range items {
		fmt.Printf("%s:%d円\n", item.Category, item.Price)
	}
	fmt.Println("===========")
}
