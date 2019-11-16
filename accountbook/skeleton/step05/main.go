// STEP05: ファイルへの保存

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Category string
	Price    int
}

func main() {

	// TODO: "accountbook.txt"という名前のファイルを書き込み用で開く
	// 開く場合にエラーが発生した場合
	if err != nil {
		// エラーを出力して終了する
		log.Fatal(err)
	}

	// 入力するデータの件数を入力する
	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	// n回繰り返す
	for i := 0; i < n; i++ {
		if err := inputItem(file); err != nil {
			// エラーを出力して終了する
			log.Fatal(err)
		}
	}

	if err := file.Close(); err != nil {
		// エラーを出力して終了する
		log.Fatal(err)
	}

	// 表示
	if err := showItems(); err != nil {
		// エラーを出力して終了する
		log.Fatal(err)
	}
}

// 入力を行いファイルに保存する
// エラーが発生した場合にはそのまま返す
func inputItem(file *os.File) error {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	// ファイルに書き出す
	// 「品目 値段」のように書き出す
	line := fmt.Sprintf("%s %d\n", item.Category, item.Price)
	if _, err := file.WriteString(line); err != nil {
		// エラーが発生した場合はエラーを返す
		return err
	}

	// TODO: エラーがなかったことを表すnilを返す
}

// 一覧の表示を行う関数
func showItems() error {

	// "accountbook.txt"という名前のファイルを読み込み用で開く
	file, err := os.Open("accountbook.txt")
	// 開く場合にエラーが発生した場合
	if err != nil {
		return err
	}

	fmt.Println("===========")

	scanner := bufio.NewScanner(file)
	// 1行ずつ読み込む
	for scanner.Scan() {
		// TODO: 1行分を取り出す

		// 1行をスペースで分割する
		splited := strings.Split(line, " ")
		// 2つに分割できなかった場合はエラー
		if len(splited) != 2 {
			// TODO: 「パースに失敗しました」というエラーを生成して返す
		}

		// 1つめが品目
		category := splited[0]

		// 2つめが値段
		// TODO: string型をint型に変換する
		if err != nil {
			return err
		}

		fmt.Printf("%s:%d円\n", category, price)
	}

	// エラーが発生したかどうか調べる
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("===========")

	return nil
}
