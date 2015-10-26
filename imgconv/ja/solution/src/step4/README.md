# STEP 4: 画像形式を変換しよう

## imageパッケージ
`image`パッケージは、画像を扱うパッケージです。
画像を表す`image.Image`インタフェースやそれを実装する具体的な型が定義されています。

`image/png`パッケージや`image/jpeg`パッケージでは、`png`や`jpeg`形式の画像を`io.Reader`から`image.Image`にデコードしたり、`image.Image`から`io.Writer`へエンコードする機能が提供されています。


## pathパッケージ
`path`パッケージは、パスに関する機能を提供しています。
例えば、`path.Ext`はファイル名から拡張子を取得でき、`path.Join`はOSごとの適切な区切り文字でパスを結合することができます。

## stringsパッケージ
`strings`パッケージは、文字列操作に関する処理を提供するパッケージです。
例えば、`strings.ToUpper`や`strings.ToLower`など大文字／小文字に変換する関数や、`strings.Join`や`strings.Split`などを文字列を結合／分割する関数が提供されています。

多くのパッケージで、引数に`io.Reader`をとっているため、`string`型から`io.Reader`を取得したい場合があります。
その場合には、`strings.NewReader`で`string`型をそのまま`io.Reader`に変換できることができます。

なお、`bytes`パッケージも`[]byte`向けに、`strings`と似たような機能を提供しています。
