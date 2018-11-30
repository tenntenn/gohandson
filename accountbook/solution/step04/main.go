// STEP04: 複数データの記録

package main

import "fmt"

type Item struct {
	Category string
	Price    int
}

func main() {

	// 入力するデータの件数を入力する
	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	// 複数のItem型の値を記録するために
	// itemsという名前のItem型のスライスの変数を定義
	// 長さ0で容量がnのスライスを作る
	items := make([]Item, 0, n)

	// iが0からitemsの容量-1の間繰り返す(n回繰り返す）
	// cap(items)はitemsの容量を返す
	for i := 0; i < cap(items); i++ {
		items = inputItem(items)
	}

	// 表示
	showItems(items)
}

// 入力を行う関数
// 追加を行うItemのスライスを受け取る
// 新しく入力したItemをスライスに追加して返す
func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	// スライスに新しく入力したitemを追加する
	items = append(items, item)

	return items
}

// 一覧の表示を行う関数
func showItems(items []Item) {
	fmt.Println("===========")

	// itemsの長さだけforを回す
	// len(items)はitemsの長さを返す
	for i := 0; i < len(items); i++ {
		// items[i]はitemsのi番目の要素(0からスタートする)
		fmt.Printf("%s:%d円\n", items[i].Category, items[i].Price)
	}

	fmt.Println("===========")
}
