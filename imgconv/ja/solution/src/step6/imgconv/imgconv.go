package imgconv

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strconv"
	"strings"
	"unicode"
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
		// TODO: cが数字の場合はfalse、そうでない場合はtrueを返す。
		// なお、iにはここがtrueになった箇所（インデックス）が入る。
		// ヒント：unicodeパッケージのドキュメントを見てみよう。
		return !unicode.IsNumber(c)
	})

	// TODO: 数字のみだった場合は、単位なしの数値のみとし、
	// sをint型に変換して返す。
	// ヒント：stringsパッケージのドキュメントを見て、strings.IndexFuncの戻り値を調べよう。
	if i < 0 {
		return strconv.Atoi(s)
	}

	// TODO:sのうち、数字だけの部分をint型に変換する。
	v, err := strconv.Atoi(s[:i])
	if err != nil {
		return 0, ErrInvalidSize
	}

	switch s[i:] {
	// TODO: "%"が指定された場合は、baseを100%として値を計算する。
	case "%":
		return int(float64(base) * float64(v) / 100), nil
	case "px":
		return v, nil
	default:
		// TODO: "%"と"px"以外の単位が指定された場合は、ErrUnknownUnitエラーを返す。
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
	// TODO: sを"x"で分割し、spに入れる。
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

	// TODO: 高さが省略された場合は、高さは幅と同じにする。
	// そうでない場合は、"x"で分割した2番目の方をパースして高さとする。
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
	// TODO: "+"で1つ〜3つに分割できない場合はErrInvalidBoundsエラーを返す。
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

	// TODO: Y座標が指定されている場合はパースし、そうでない場合は0とする
	if len(sp) == 3 {
		p.Y, err = parseRelSize(img.Bounds().Max.Y, sp[2])
	}

	// TODO: 開始座標分だけrを並行移動させる。
	r = r.Add(p)

	return
}

func newDrawImage(r image.Rectangle, m color.Model) draw.Image {
	// TODO: 各カラーモデルごとに画像を初期化し返す。
	// なお、指定されたカラーモデルがimage/colorパッケージに定義されていない場合は、
	// RGBAの画像を作って返す。
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
	// TODO: 現在の画像をdstに描画する。
	// 描画する現在の画像の開始点は、rの左上である。
	draw.Draw(dst, dst.Bounds(), img, r.Min, draw.Src)

	// TODO: imgに埋め込まれているimage.Imageをdstで更新する。
	img.Image = dst

	return nil
}
