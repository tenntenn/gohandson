# STEP0: 環境構築

ハンズオンを行うために、以下の環境を整える必要があります。
ハンズオンをスムーズに行うために、可能な場合は以下の準備をお願いします。

* Googleアカウントの準備（Gmailが使える状態）
* Google App Engine SDK for Goのインストール
* Python 2.7のインストール

上記のインストールが難しい場合は会場のネットワーク環境に負荷を与えないために、以下の準備をお願いします。

* Googleアカウントの準備（Gmailが使える状態）
* Google App Engine SDK for Goのダウンロード（[Mac](https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_darwin_amd64-1.9.62.zip)・[Windows](https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_windows_amd64-1.9.62.zip)）
* Python2.7のダウンロード（Macは不要・[Windowsの場合](https://www.python.org/ftp/python/2.7.4/python-2.7.4.msi)）

## Googleアカウントの準備

Google App Engineを用いるには、Googleアカウントが必要です。
お持ちでない方はGoogleのページから作成をお願いします。
Gmailが使える状態であれば、問題はありません。

## Google App Engine SDK for Goのインストール

### SDKのダウンロード

[ダウンロードページ](https://cloud.google.com/appengine/docs/standard/go/download)から
`Download and install the original App Engine SDK for Go.`をクリックし、使用しているOS用のファイルをダウンロードする。

ダウンロードしたファイルを適切な場所に解凍しておく。

### PATHを通す

解凍したディレクトリ以下の`go_appengine`へPATHを通します。

#### Macの場合
.bashrcや.zchrcなどに、以下のように記載してください。
なお、`DIRECTORY_PATH`はダウンロードして解凍したSDKの場所です。

```
export PATH=$PATH:DIRECTORY_PATH/go_appengine/
```

#### Windowsの場合

システム環境設定でPATHという環境変数に`DIRECTORY_PATH\go_appengine`を追加してください。
なお、`DIRECTORY_PATH`はダウンロードして解凍したSDKの場所です。

### 1.2.3. Python 2.7のインストール

#### Macの場合

Macの場合は最初からPython 2.7がインストールされています。
ターミナルでデフォルトのPythonのバージョンを確認してください。

```
/usr/bin/env/python -V
```

#### Windowsの場合

コマンドプロンプトで以下のコマンドを実行し、Python 2.7がインストールされていることが分からない場合は、以下の手順でPython 2.7をインストールしてください。

```
python -V
```

[PythonのWebサイト](https://www.python.org/download/releases/2.7.4)からPython 2.7をダウンロードし、インストールしてください。
