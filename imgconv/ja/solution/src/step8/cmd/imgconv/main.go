package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"step8/imgconv"
)

var (
	resize string
	clip   string
	format string
)

func init() {
	flag.StringVar(&resize, "resize", "", "出力する画像サイズ（`幅[px|%]x高さ[px|%]`）")
	flag.StringVar(&clip, "clip", "", "切り取る画像サイズ（`幅[px|%]x高さ[px|%]`）")
	flag.StringVar(&format, "format", "", "ディレクトリ指定の場合に、変換する画像ファイルのフォーマット（`png|jpeg|jpg`）")
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
	img := &imgconv.Image{_img}

	if resize != "" {
		if err := img.Resize(resize); err != nil {
			return fmt.Errorf("%s\n", err.Error())
		}
	}

	if clip != "" {
		if err := img.Clip(clip); err != nil {
			return fmt.Errorf("%s\n", err.Error())
		}
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

type file string

func (f file) Ext() string {
	return path.Ext(string(f))
}

func (f file) Dir() string {
	return path.Dir(string(f))
}

func (f file) Name() string {
	return strings.Replace(path.Base(string(f)), f.Ext(), "", -1)
}

func run() error {
	args := flag.Args()
	if len(args) < 2 {
		return fmt.Errorf("画像ファイルを指定してください。")
	}

	info, err := os.Stat(args[0])
	if os.IsNotExist(err) {
		return fmt.Errorf("画像ファイルが存在しません。%s", args[0])
	}

	// ディレクトリの場合
	if info.IsDir() {
		t, err := template.New("dst").Parse(args[1])
		if err != nil {
			return err
		}

		return filepath.Walk(args[0], func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			ext := strings.ToLower(path.Ext(p))
			if format != "" {
				if ext != "."+format {
					return nil
				}
			} else if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
				return nil
			}

			var buf bytes.Buffer
			t.Execute(&buf, file(p))

			return convert(buf.String(), p)
		})
	}

	// ファイルの場合
	return convert(args[1], args[0])
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
