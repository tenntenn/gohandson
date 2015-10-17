package imgconv

import (
	"fmt"
	"image"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/image/draw"
)

var (
	SizeError = fmt.Errorf("画像サイズの形式が不正です。")
)

type Image struct {
	image.Image
}

func parseRelSize(base int, s string) (int, error) {
	sp := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsNumber(c)
	})

	if len(sp) <= 0 || len(sp) > 2 {
		return 0, SizeError
	}

	v, err := strconv.Atoi(sp[0])
	if err != nil {
		return 0, SizeError
	}

	if len(sp) == 1 {
		return v, nil
	}

	switch sp[1] {
	case "%":
		return int(float64(base) * float64(v) / 100), nil
	default:
		return v, nil
	}
}

func (img *Image) parseSize(s string) (r image.Rectangle, err error) {
	sp := strings.Split(s, "x")
	if len(sp) <= 0 || len(sp) > 2 {
		err = SizeError
		return
	}

	r.Max.X, err = parseRelSize(img.Bounds().Max.X, sp[0])
	if err != nil {
		err = SizeError
		return
	}

	if len(sp) == 1 {
		r.Max.Y, err = parseRelSize(img.Bounds().Max.Y, sp[0])
	} else {
		r.Max.Y, err = parseRelSize(img.Bounds().Max.Y, sp[1])
	}

	return
}

func (img *Image) Resize(s string) error {
	r, err := img.parseSize(s)
	if err != nil {
		return err
	}

	var dst draw.Image
	switch img.ColorModel() {
	default:
		dst = image.NewRGBA(r)
	}
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Src, nil)
	img.Image = dst

	return nil
}
