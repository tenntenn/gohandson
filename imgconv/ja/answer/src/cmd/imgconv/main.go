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

	"../../imgconv"
)

var (
	resize string
	clip   string
)

func init() {
	flag.StringVar(&resize, "resize", "", "出力する画像サイズ（`幅[px|%]x高さ[px|%]`）")
	flag.StringVar(&clip, "clip", "", "切り取る画像サイズ（`幅[px|%]x高さ[px|%]`）")
	flag.Parse()
}

func run() int {
	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "画像ファイルを指定してください。")
		return 1
	}

	src, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "画像ファイルが開けませんでした。%s", args[0])
		return 1
	}
	defer src.Close()

	dst, err := os.Create(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "画像ファイルを書き出せませんでした。%s", args[1])
		return 1
	}
	defer dst.Close()

	_img, _, err := image.Decode(src)
	img := &imgconv.Image{_img}

	if resize != "" {
		if err := img.Resize(resize); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return 1
		}
	}

	if clip != "" {
		if err := img.Clip(clip); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return 1
		}
	}

	switch strings.ToLower(path.Ext(args[1])) {
	case ".png":
		err = png.Encode(dst, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(dst, img, &jpeg.Options{jpeg.DefaultQuality})
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "画像ファイルを書き出せませんでした。%s", args[1])
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
