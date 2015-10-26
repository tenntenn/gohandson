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
	// ErrInvalidSize は、指定したサイズが不正だった場合のエラーです。
	ErrInvalidSize = fmt.Errorf("指定したサイズの形式が不正です。")
	// ErrInvalidBounds は、指定した領域の形式が不正だった場合のエラーです。
	ErrInvalidBounds = fmt.Errorf("指定した領域の形式が不正です。")
	// ErrUnkownUnit は、想定外の不正な単位だった場合のエラーです。
	ErrUnkownUnit = fmt.Errorf("不正な単位です。")
)

// Image は、image.Image をラップした構造体です。
// ラップした画像に対して、切り抜き等の操作を提供します。
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
		return 0, ErrInvalidSize
	}

	switch s[i:] {
	case "%":
		return int(float64(base) * float64(v) / 100), nil
	case "px":
		return v, nil
	default:
		return 0, ErrUnkownUnit
	}
}

func (img *Image) parseSize(s string) (r image.Rectangle, err error) {
	sp := strings.Split(s, "x")
	if len(sp) <= 0 || len(sp) > 2 {
		err = ErrInvalidSize
		return
	}

	r.Max.X, err = parseRelSize(img.Bounds().Max.X, sp[0])
	if err != nil {
		err = ErrInvalidSize
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
		err = ErrInvalidBounds
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

// Resize は、画像のリサイズを行う。
// "10x20"のように、幅x高さを指定する。
// "10%X10%"のように、単位も指定できる。
// 使用できる単位は、"px"と"%"である。
// "%"を指定すると、元の画像の幅や高さを基準とする。
// 高さを省略すると、幅と同じになる。
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

// Clip は、画像の一部部分を矩形で切り抜く。
// 切り抜く領域は、幅x高さ+X座標+Y座標で指定し、
// (X座標, Y座標) - (X座標+幅, Y座標+高さ)の領域が切り抜かれる。
// 幅と高さは、Resizeで指定できるものと同じである。
// XY座標にも"px"や"%"の単位が使える。
// "%"を指定すると、元の画像の幅や高さを基準とする。
// XY座標は省略でき、省略すると0となる。
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
