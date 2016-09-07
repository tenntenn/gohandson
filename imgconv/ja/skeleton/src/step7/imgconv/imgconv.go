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
	// ErrUnknownUnit は、想定外の不正な単位だった場合のエラーです。
	ErrUnknownUnit = fmt.Errorf("不正な単位です。")
)

// Image は、image.Image をラップした構造体です。
// ラップした画像に対して、切り抜き等の操作を提供します。
type Image struct {
	image.Image
}

// baseで指定した数値をもとに、sで記述された値をパースし返します。
// baseを使うのは、s内で単位として"%"が使われた場合のみです。
// 単位に"px"が使われた場合と単位がない場合は、
// 単位を省いた数字の部分を数値に変換して返します。
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
		return 0, ErrUnknownUnit
	}
}

// sで指定された画像の幅と高さをパース返します。
// 幅と高さは、"幅x高さ"のように指定されます。
// 高さを省略した場合は、幅と同じになります。
// 幅と高さには単位を指定することができ、使用できる単位は"px"と"%"です。
// "%"と指定した場合は、現在の画像の幅と高さをもとに計算します。
// 単位を指定していない場合は、"px"を指定した場合と同じです。
func (img *Image) parseSize(s string) (sz image.Point, err error) {
	sp := strings.Split(s, "x")
	if len(sp) <= 0 || len(sp) > 2 {
		err = ErrInvalidSize
		return
	}

	sz.X, err = parseRelSize(img.Bounds().Max.X, sp[0])
	if err != nil {
		err = ErrInvalidSize
		return
	}

	if len(sp) == 1 {
		sz.Y = sz.X
	} else {
		sz.Y, err = parseRelSize(img.Bounds().Max.Y, sp[1])
	}

	return
}

// sで指定された画像の幅と高さ、開始座標をパースして領域を返します。
// sには、"幅x高さ+X座標+Y座標"のように指定します。
// 高さを省略した場合は、幅と同じになります。
// 幅と高さ、XY座標には単位を指定することができ、使用できる単位は"px"と"%"です。
// "%"と指定した場合は、現在の画像の幅と高さをもとに計算します。
// 単位を指定していない場合は、"px"を指定した場合と同じです。
func (img *Image) parseBounds(s string) (r image.Rectangle, err error) {
	sp := strings.Split(s, "+")
	if len(sp) <= 0 || len(sp) > 3 {
		err = ErrInvalidBounds
		return
	}

	r.Max, err = img.parseSize(sp[0])
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

// Clip は、画像の一部部分を矩形で切り抜く。
// 切り抜く領域は、幅x高さ+X座標+Y座標で指定し、
// (X座標, Y座標) - (X座標+幅, Y座標+高さ)の領域が切り抜かれる。
// 幅と高さ、XY座標には"px"や"%"の単位が使える。
// "%"を指定すると、元の画像の幅や高さを基準とする。
// 高さは省略すると、幅と同じになる。
// また、XY座標も省略でき、省略するとそれぞれ0となる。
func (img *Image) Clip(s string) error {
	r, err := img.parseBounds(s)
	if err != nil {
		return err
	}

	dst := newDrawImage(image.Rectangle{image.ZP, r.Size()}, img.ColorModel())
	draw.Draw(dst, dst.Bounds(), img, r.Min, draw.Src)

	img.Image = dst

	return nil
}

// Resize は、画像のリサイズを行う。
// "10x20"のように、幅x高さを指定する。
// "10%X10%"のように、単位も指定できる。
// 使用できる単位は、"px"と"%"である。
// "%"を指定すると、元の画像の幅や高さを基準とする。
// 高さを省略すると、幅と同じになる。
func (img *Image) Resize(s string) error {
	sz, err := img.parseSize(s)
	if err != nil {
		return err
	}

	dst := newDrawImage(image.Rectangle{image.ZP, sz}, img.ColorModel())
	// TODO: draw.NearestNeighbor.Scaleを使って画像を縮小する。
	// なお、第6引数の*draw.Optionsはnilで構わない。

	img.Image = dst

	return nil
}
