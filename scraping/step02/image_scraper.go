package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// ImageFormat は変換する画像のフォーマットを表します。
type ImageFormat string

const (
	// FormatNotChange は変換しないことを表します。
	FormatNotChange ImageFormat = ""
	// FormatPNG はPNG形式に変換を表します。
	FormatPNG ImageFormat = "png"
	// FormatJPG はJPEG形式に変換を表します。
	FormatJPEG ImageFormat = "jpeg"
)

// Ext は対応する拡張子を取得します。
// 対応するものがない場合は空文字が返されます。
func (f ImageFormat) Ext() string {
	switch f {
	case FormatPNG:
		return ".png"
	case FormatJPEG:
		return ".jpg"
	}
	return ""
}

// 指定した文字列が対応している画像形式か取得します。
func IsAllowFormat(s string) bool {
	switch ImageFormat(s) {
	case FormatPNG, FormatJPEG:
		return true
	}
	return false
}

// ImageScraper はスクレイピングを行いimgタグの画像をダウンロードして保存します。
//
// AllowHostにはアクセス可能なホストを"host:port"形式で指定します。
// AllowHostが空の場合は特に制限を設けません。
//
// Formatを指定するとダウンロードした画像を指定したフォーマットで変換します。
//
// HTTPClientを指定するとHTMLや画像のダウンロードに指定したHTTPクライアントを使用します。
type ImageScraper struct {
	AllowHost  []string
	Format     ImageFormat
	HTTPClient *http.Client

	visited map[string]bool
	dir     string
}

// New は新しいImageScraperを作成します。
// dirはダウンロードした画像を保存するディレクトリです。
func New(dir string) *ImageScraper {
	return &ImageScraper{
		visited: map[string]bool{},
		dir:     dir,
	}
}

// Visit は指定したURLにアクセスし、imgタグの画像をダウンロードします。
// aタグがある場合は再帰的にダウンロードを行います。
func (s *ImageScraper) Visit(u *url.URL) error {

	// フラグメントを取り除く
	u = &(*u) // copy
	u.Fragment = ""

	urlStr := u.String()
	if s.visited[urlStr] {
		return nil
	}
	s.visited[urlStr] = true

	if !s.isAllowed(u) {
		return nil
	}

	// TODO
	if urlStr == "http://localhost:8080/blog/" {
		return nil
	}
	if urlStr == "http://localhost:8080/cmd/godoc" {
		return nil
	}
	if urlStr == "http://localhost:8080/doc/articles/error_handling.html" {
		return nil
	}
	if urlStr == "http://localhost:8080/doc/articles/defer_panic_recover.html" {
		return nil
	}

	fmt.Println("Visit", urlStr)

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return err
	}

	resp, err := s.httpClient().Do(req)
	if err != nil {
		return err
	}

	if !s.isAllowed(resp.Request.URL) {
		return nil
	}

	if respURL := resp.Request.URL.String(); urlStr != respURL {
		if s.visited[respURL] {
			return nil
		}
		s.visited[respURL] = true
	}

	switch {
	case resp.StatusCode == http.StatusNotFound:
		// 404は無視
		return nil
	case resp.StatusCode >= http.StatusBadRequest:
		return fmt.Errorf("%s へのリクエストでエラーが発生（ステータスコード:%d)", urlStr, resp.StatusCode)
	case resp.StatusCode != http.StatusOK: // リダイレクトなどを無視
		return nil
	}

	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	if err := s.parse(resp.Request.URL, bytes.NewReader(body)); err != nil {
		return err
	}

	return nil
}

// parse は指定したReaderからHTMLをパースし、トラバースを行います。
func (s *ImageScraper) parse(baseURL *url.URL, r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	if err := s.traverse(baseURL, doc); err != nil {
		return err
	}

	return nil
}

// traverse は指定したノードをトラバースします。
// aタグの場合は、さらにVisitし、imgタグの場合は画像をダウンロードします。
// 子ノードについても再帰的に処理を行います。
func (s *ImageScraper) traverse(baseURL *url.URL, n *html.Node) error {

	switch {
	case n.Type == html.ElementNode && n.Data == "a":
		if urlStr := s.attr(n, "href"); urlStr != "" {
			absURL, err := s.absoluteURL(baseURL, urlStr)
			if err != nil {
				return err
			}

			if err := s.Visit(absURL); err != nil {
				return err
			}
		}
	case n.Type == html.ElementNode && n.Data == "img":
		if src := s.attr(n, "src"); src != "" && !strings.HasPrefix(src, "base64") {
			absURL, err := s.absoluteURL(baseURL, src)
			if err != nil {
				return err
			}

			if err := s.downloadImage(absURL); err != nil {
				return err
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err := s.traverse(baseURL, c); err != nil {
			return err
		}
	}

	return nil
}

// downloadImage は指定したURLから画像をダウンロードします。
// 変換する必要がある場合は画像形式を変換して保存します。
func (s *ImageScraper) downloadImage(srcURL *url.URL) error {
	req, err := http.NewRequest(http.MethodGet, srcURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := s.httpClient().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// copy
	u := *srcURL
	u.RawQuery = ""
	u.Fragment = ""
	path := filepath.Join(s.dir, path.Base(u.String()))

	// 変換が必要な場合は変換する
	if s.Format != FormatNotChange {
		if err := s.convert(resp.Body, path); err != nil {
			return err
		}
		return nil
	}

	// 変換が必要ない
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}

	return nil
}

// convert 画像の変換を行い保存まで行います。
// pathで指定したパスと同じ場所に拡張子だけ変更して保存します。
func (s *ImageScraper) convert(r io.Reader, path string) error {
	img, format, err := image.Decode(r)
	if err != nil {
		return err
	}

	// 変換元がpngかjpeg以外は無視
	switch format {
	case "png", "jpeg":
	default:
		return nil
	}

	// 拡張子を変換する
	i := strings.LastIndex(path, ".")
	p := path[:i] + s.Format.Ext()

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	// フォーマットごとに変換する
	switch s.Format {
	case FormatPNG:
		if err := png.Encode(f, img); err != nil {
			return err
		}
	case FormatJPEG:
		opts := &jpeg.Options{Quality: jpeg.DefaultQuality}
		if err := jpeg.Encode(f, img, opts); err != nil {
			return err
		}
	}

	return nil
}

// httpClient は使用するHTTPクライアントを取得します。
func (s *ImageScraper) httpClient() *http.Client {
	if s.HTTPClient == nil {
		return http.DefaultClient
	}
	return s.HTTPClient
}

// isAllowed は指定したURLがアクセス可能なホストか調べます。
func (s *ImageScraper) isAllowed(u *url.URL) bool {
	if len(s.AllowHost) == 0 {
		return true
	}

	hp := net.JoinHostPort(u.Hostname(), u.Port())
	for _, h := range s.AllowHost {
		if h == hp {
			return true
		}
	}
	return false
}

// attr は指定したキーのHTML属性を取得します。
func (s *ImageScraper) attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// absoluteURL は指定したURLからの相対的なパスから絶対パスのURLを取得します。
func (s *ImageScraper) absoluteURL(baseURL *url.URL, ref string) (*url.URL, error) {
	refURL, err := url.Parse(ref)
	if err != nil {
		return nil, err
	}
	return baseURL.ResolveReference(refURL), nil
}
