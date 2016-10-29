package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"step8/imgconv"
)

var (
	clip   string
	resize string
	format string
)

func init() {
	flag.StringVar(&clip, "clip", "", "切り取る画像サイズ（`幅[px|%]x高さ[px|%]`）")
	flag.StringVar(&resize, "resize", "", "出力する画像サイズ（`幅[px|%]x高さ[px|%]`）")
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

	if clip != "" {
		if err := img.Clip(clip); err != nil {
			return fmt.Errorf("%s\n", err.Error())
		}
	}

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

type file string

// TODO: 拡張子を返すメソッドを作る。
func (f file) Ext() string {
	return filepath.Ext(string(f))
}

// TODO: ディレクトリを返すメソッドを作る。
func (f file) Dir() string {
	return filepath.Dir(string(f))
}

func (f file) Name() string {
	return strings.Replace(filepath.Base(string(f)), f.Ext(), "", -1)
}

func run() error {
	args := flag.Args()
	if len(args) < 2 {
		return fmt.Errorf("画像ファイルを指定してください。")
	}

	// TODO: 入力ファイルのファイル情報を取得する。
	info, err := os.Stat(args[0])
	if /* TODO: 画像ファイルが存在しない場合  */ os.IsNotExist(err) {
		return fmt.Errorf("画像ファイルが存在しません。%s", args[0])
	}

	// ディレクトリの場合
	if info.IsDir() {
		t, err := template.New("dst").Parse(args[1])
		if err != nil {
			return err
		}

		return filepath.Walk(args[0], func( /* TODO: 引数を埋める */ p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// TODO: ディレクトリだったら何もしない。
			if info.IsDir() {
				return nil
			}

			ext := strings.ToLower(filepath.Ext(p))
			// TODO: 拡張子がformatで指定されたものでなければ何もしない。
			// formatで何もしてなければ、拡張子が".png"、".jpg"、".jpeg"以外は何もしない。
			if format != "" {
				if ext != "."+format {
					return nil
				}
			} else if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
				return nil
			}

			var buf bytes.Buffer
			// TODO: pをfile型に変換して、テンプレートに埋め込む。
			// file型に変換すると、DirやExt、Nameなどのメソッドが使用できる。
			// テンプレートの展開先は、bufにする。
			t.Execute(&buf, file(p))

			// TODO: ファイル1つを変換する。
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
