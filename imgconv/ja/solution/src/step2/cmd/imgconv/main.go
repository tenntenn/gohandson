package main

import (
	"fmt"
	"os"
)

func run() error {

	// TODO: 引数が足りない場合は、エラーを返す
	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません。")
	}

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	return nil
}

func main() {
	if err := run(); err != nil {
		// TODO: 標準エラー出力（os.Stderr）にエラーを出力する
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
