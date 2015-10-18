package imgconv

import (
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/image/draw"
)

var (
	SizeError   = fmt.Errorf("指定したサイズの形式が不正です。")
	BoundsError = fmt.Errorf("指定した領域の形式が不正です。")
	UnkownUnit  = fmt.Errorf("不正な単位です")
)

type Image struct {
	image.Image
}

func parseRelSize(base int, s string) (int, error) {
	i := strings.IndexFunc(s, func(c rune) bool {
		return !unicode.IsNumber(c)
	})

	if i < 0 {
		return strconv.Atoi(s)
	}

	v, err := strconv.Atoi(s[:i])
	if err != nil {
		return 0, SizeError
	}

	switch s[i:] {
	case "%":
		return int(float64(base) * float64(v) / 100), nil
	case "px":
		return v, nil
	default:
		return 0, UnkownUnit
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

func (img *Image) parseBounds(s string) (r image.Rectangle, err error) {
	sp := strings.Split(s, "+")
	if len(sp) <= 0 || len(sp) > 3 {
		err = BoundsError
		return
	}

	r, err = img.parseSize(sp[0])
	if err != nil || len(sp) == 1 {
		return
	}

	var p image.Point
	p.X, err = parseRelSize(img.Bounds().Max.X, sp[1])
	if err != nil {
		return
	}

	if len(sp) == 3 {
		p.Y, err = parseRelSize(img.Bounds().Max.Y, sp[2])
	}

	r = r.Add(p)
	return
}

func newDrawImage(r image.Rectangle, m color.Model) draw.Image {
	switch m {
	case color.RGBA64Model:
		return image.NewRGBA64(r)
	case color.NRGBAModel:
		return image.NewNRGBA(r)
	case color.NRGBA64Model:
		return image.NewNRGBA64(r)
	case color.AlphaModel:
		return image.NewAlpha(r)
	case color.Alpha16Model:
		return image.NewAlpha16(r)
	case color.GrayModel:
		return image.NewGray(r)
	case color.Gray16Model:
		return image.NewGray16(r)
	default:
		return image.NewRGBA(r)
	}
}

func (img *Image) Resize(s string) error {
	r, err := img.parseSize(s)
	if err != nil {
		return err
	}

	dst := newDrawImage(r, img.ColorModel())
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Src, nil)
	img.Image = dst

	return nil
}

func (img *Image) Clip(s string) error {
	r, err := img.parseBounds(s)
	if err != nil {
		return err
	}

	dst := newDrawImage(r, img.ColorModel())
	draw.Draw(dst, dst.Bounds(), img, r.Min, draw.Src)

	img.Image = dst

	return nil
}
