package main

import "fmt"

func main() {
	var category string
	var price int

	fmt.Print("品目>")
	fmt.Scan(&category)

	fmt.Print("値段>")
	fmt.Scan(&price)

	fmt.Printf("%sに%d円使いました\n", category, price)
}
