# STEP 6: 画像を切り抜こう

STEP 6では、`clip`というフラグをコマンドライン引数で取得し、
そのフラグで指定された領域で画像を切り抜きファイルに書き出すというプログラムを作ります。

ここでは、`image/draw`パッケージを使って画像の描画を行う方法について説明します。
また、文字列をパースして切り抜く領域を取得するために使用する`strconv`パッケージについても説明を行います。

また、ここで定義している`imgconv`パッケージにある`Image`構造体の定義には、埋込みという機能を使用しています。
非常に面白い機能ですので、時間に余裕のある方は、説明を読んでみると良いでしょう。
なお、埋込みについては、理解しなくても次のSTEPに進んでもらっても問題ありません。

## image/drawパッケージ

`image/draw`パッケージは、画像の描画を行う機能を提供するパッケージです。
`draw.Image`は`image.Image`に`Set`メソッドを追加したインタフェースで、画像に対して書き込みができるようになっています。

`draw.Draw`を使えば、指定した画像対して、別の画像を描画する事ができます。
もちろん、一部分を描画することもできます。

`image/draw`パッケージについては、[パッケージドキュメント](https://golang.org/pkg/image/draw/)に詳しく記載されています。

## strconvパッケージ

文字列から数値に変換するには、`strconv`パッケージの関数を使います。
`string`から`int`型に変換する(10進数として解釈)には、`strconv.Atoi`関数を使うのが手っ取り早いです。

`int`型以外に、または10進数以外として数値に変換したい場合は、`Parse`で始まる関数を使います。
なお、`int`と`int64`は別の型であり、別の型同士の演算は必ずキャストがいります。
キャストは、`int(f)`や`float64(n)`のようにできます。

`strconv`パッケージについては、[パッケージドキュメント](https://golang.org/pkg/strconv)に詳しく記載されています。
特に[サンプル](https://golang.org/pkg/strconv/#pkg-examples)を読むと理解が進むと思います。

## 埋込み

*※埋込みの説明は難しいので、理解できなくても構いません。時間が掛かりそうであれば、飛ばしてください。*

構造体には、匿名フィールドとして指定した型の値を埋め込むことができます。
匿名フィールドは、以下のように、名前を持たず型情報だけを指定します。

```
type MyImage struct {
    image.Image
}
```

埋め込まれた構造体は、埋め込んだ型の持つメソッドや構造体の場合はフィールドをあたかも自分のメソッドやフィールドのように呼び出すことができます。

```
// _imgはimage.Image型
img := &MyImage{_img} 

// 埋め込んだimage.Image型のメソッドを呼び出せる
fmt.Println(img.Bounds()) 
```

しかし、これはあくまで処理を埋込んだ値に委譲しているだけで、決して継承をしているわけではありません。
埋込みでJavaに出てくるような型階層を作ろうとしても失敗します。

Goのインタフェースは、インタフェースで指定したメソッドリストのメソッドをすべて実装している型であれば、明示的に実装することを記述する（Javaの`implements`など）ことなく、そのインタフェース型として振る舞うことができます。（[インタフェース](https://go-tour-jp.appspot.com/methods/4)については、Tour Of Go に出てきますので、覚えてない方はぜひ復習をお願いします。）
そして、埋め込んだ値の持つメソッドもインタフェースを実装するためのメソッドリストとしてカウントされます。
さらに、インタフェースも匿名フィールドとして埋め込むことができるため、上記の例だと`Image`構造体が`image.Image`インタフェースを実装していることになります。

```
// _imgはimage.Image型
img := &MyImage{_img} 

// 埋め込んだ_imgがimage.Imageを実装しているので、
// imgもimage.Imageとして扱える
var img2 image.Image = img
```

匿名フィールドには、以下のように型名でアクセスできます。
当然ながら同じ型の値を2つ以上埋め込むことはできません。
フィールドであることには変わらないので、埋め込んだ値を後で変更することも可能です。

```
// _imgはimage.Image型
img := &MyImage{_img} 

// _img2もimage.Image型
img.Image = _img2
```

## 実行例

```
$ pwd
/path/to/gohandson/imgconv/ja/solution
$ GOPATH=`pwd`
$ go install step6/cmd/imgconv
$ go install tools/cmd/httpget
$ ./bin/httpget https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/hi.png > gopher.png
$ ./bin/imgconv -clip 100x100+50%+50% gopher.png gopher2.png
```

## 目次

* STEP 1: [go installしてみよう](../step1)（[解答例](../../../solution/src/step1)）
* STEP 2: [コマンドライン引数を取ろう](../step2)（[解答例](../../../solution/src/step2)）
* STEP 3: [ファイルを扱おう](../step3)（[解答例](../../../solution/src/step3)）
* STEP 4: [画像形式を変換しよう](../step4)（[解答例](../../../solution/src/step4)）
* STEP 5: [`flag`パッケージを使おう](../step5)（[解答例](../../../solution/src/step5)）
* STEP 6: [画像を切り抜こう](../step6)（[解答例](../../../solution/src/step6)）
* STEP 7: [画像を縮小／拡大しよう](../step7)（[解答例](../../../solution/src/step7)）
* STEP 8: [複数のファイルを処理しよう](../step8)（[解答例](../../../solution/src/step8)）
