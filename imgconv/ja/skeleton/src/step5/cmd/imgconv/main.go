package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var (
	clip string
)

func init() {
	// TODO: clipというフラグを追加し、変数clipに入れる。
	// デフォルト値は、""。
	// 説明は、"切り取る画像サイズ（`幅[px|%]x高さ[px|%]`）"。

	// TODO: フラグをパースする。
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

	// TODO: clipで何か指定されていれば、
	// 標準出力に"切り抜きを行う予定"という文字列とともにclipの中身を出力する

	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}

	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。%s", dst)
	}

	return nil
}

func run() error {
	// TODO: os.Argsではなく、flag.Args()を使ってコマンドライン引数を取得する。

	// TODO: フラグ（オプション）以外で、引数が2つ以上指定されているかチェックする。
	// 引数が２つ以上指定されていない場合は、"画像ファイルを指定してください。"というエラーを返す。

	return convert(args[1], args[0])
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
