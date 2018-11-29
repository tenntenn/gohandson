package main

import "fmt"

type Item struct {
	Category string
	Price    int
}

func main() {
	var items []Item

	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		items = inputItem(items)
	}

	// 表示
	showItems(items)
}

// 入力
func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	items = append(items, item)

	return items
}

// 一覧の表示
func showItems(items []Item) {
	fmt.Println("===========")
	for i := 0; i < len(items); i++ {
		fmt.Printf("%s:%d円\n", items[i].Category, items[i].Price)
	}
	fmt.Println("===========")
}
