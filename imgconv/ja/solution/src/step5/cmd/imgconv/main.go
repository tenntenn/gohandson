package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

var (
	resize string
)

func init() {
	// TODO: resizeというフラグを追加し、変数resizeに入れる。
	// デフォルト値は、""。
	// 説明は、"出力する画像サイズ（`幅[px|%]x高さ[px|%]`）"。
	flag.StringVar(&resize, "resize", "", "出力する画像サイズ（`幅[px|%]x高さ[px|%]`）")
	// TODO: フラグをパースする。
	flag.Parse()
}

func convert(dst, src string) error {

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("画像ファイルが開けませんでした。%s", src)
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。%s", dst)
	}
	defer df.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	// TODO: resizeで何か指定されていれば、
	// 標準出力に"リサイズをする予定"という文字列とともにresizeの中身を出力する
	if resize != "" {
		fmt.Println("リサイズする予定", resize)
	}

	switch strings.ToLower(path.Ext(dst)) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{jpeg.DefaultQuality})
	}

	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。%s", dst)
	}

	return nil
}

func run() error {
	// TODO: os.Argsではなく、flag.Args()を使ってコマンドライン引数を取得する。
	args := flag.Args()
	// TODO: フラグ（オプション）以外で、引数が2つ以上指定されているかチェックする。
	// 引数が２つ以上指定されていない場合は、"画像ファイルを指定してください。"というエラーを返す。
	if len(args) < 2 {
		return fmt.Errorf("画像ファイルを指定してください。")
	}

	return convert(args[1], args[0])
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
