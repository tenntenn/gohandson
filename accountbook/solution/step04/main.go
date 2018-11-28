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
		var item Item

		fmt.Print("品目>")
		fmt.Scan(&item.Category)

		fmt.Print("値段>")
		fmt.Scan(&item.Price)

		items = append(items, item)
	}

	fmt.Println("===========")
	for i := 0; i < n; i++ {
		fmt.Printf("%s:%d円\n", items[i].Category, items[i].Price)
	}
	fmt.Println("===========")
}
