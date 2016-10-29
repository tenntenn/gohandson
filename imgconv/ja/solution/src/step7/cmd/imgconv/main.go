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

	"step7/imgconv"
)

var (
	clip   string
	resize string
)

func init() {
	flag.StringVar(&clip, "clip", "", "切り取る画像サイズ（`幅[px|%]x高さ[px|%]`）")
	flag.StringVar(&resize, "resize", "", "出力する画像サイズ（`幅[px|%]x高さ[px|%]`）")
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

	_img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	img := imgconv.Image{_img}

	if clip != "" {
		if err := img.Clip(clip); err != nil {
			return fmt.Errorf("%s\n", err.Error())
		}
	}

	// TODO: resizeが指定されていれば、リサイズを行う。
	if resize != "" {
		if err := img.Resize(resize); err != nil {
			return fmt.Errorf("%s\n", err.Error())
		}
	}

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
	args := flag.Args()
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
