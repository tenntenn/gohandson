package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません。")
	}

	src, dst := os.Args[1], os.Args[2]

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("ファイルが開けませんでした。%s", src)
	}
	// TODO: 関数終了時にファイルを閉じる
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ファイルを書き出せませんでした。%s", dst)
	}
	// TODO: 関数終了時にファイルを閉じる
	defer df.Close()

	scanner := bufio.NewScanner(sf)
	// TODO: sfから1行ずつ読み込み、"行数:"を前に付けてdfに書き出す。
	for i := 1; scanner.Scan(); i++ {
		fmt.Fprintf(df, "%d:%s\n", i, scanner.Text())
	}

	// TODO: scannerから得られたエラーを返す
	return scanner.Err()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
