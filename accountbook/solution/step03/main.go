package main

import "fmt"

type Item struct {
	Category string
	Price    int
}

func main() {

	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	fmt.Printf("%sに%d円使いました\n", item.Category, item.Price)
}
