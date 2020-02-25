# STEP 2: ボトルネックを探す

## トレースを行う

STEP 1で作成したプログラムはなんとなく遅い気がします。
STEP 2では、`runtime/trace`パッケージを用いてトレースし、ボトルネックとなる部分を探します。

まずは、`main`関数をトレースするために、`main`関数の中身を`_main`関数に移動させます。

そして、トレース対象の処理を開始する前に、`trace.Start`関数を呼び出し、終了後に`trace.Stop`を呼び出します。
なお、`defer`を用いることで、関数の終了時に呼び出すことができます。

```go
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()

	_main()
}
```

## ユーザアノテーション

Go1.11からユーザアノテーションという機能が入りました。
ユーザが任意のタイミングでトレース情報を出力することができる機能です。

ユーザアノテーションは、以下の単位で行うことができます。

* Task
* Region
* Log

1つのTaskに複数のRegionがあり、Regionの中でLogを出すというイメージです。

ここでは、Taskは「コーヒーを淹れること」、Regionは「お湯を沸かす」、「豆を挽く」、「コーヒー淹れる」の3つになります。

`Task`は次のように作成することができます。
作成した`Task`は、`End`メソッドを呼び出すまでの処理が対象となります。
通常は`defer`で呼び出すことで関数の終了時に呼び出します。

```go
ctx, task := trace.NewTask(context.Background(), "make coffee")
defer task.End()
```

`trace.NewTask`関数の第1引数で返される`context.Context`型の値には、Taskの情報が保持されています。

Regionは次のように、`trace.StartRegion`を呼び出すことで作ることができます。
`trace.StartRegion`の第1引数に`trace.NewTask`関数で返ってきた`context.Context`を渡すことで、TaskとRegionを紐付けることができます。

作成したRegionは、`End`メソッドを呼び出すまでの処理が対象となります。

```go
region := trace.StartRegion(ctx, "region_name")
defer region.End()
```

なお、次のように1行で書くこともできます。

```go
defer trace.StartRegion(ctx, "region_name").End()
```

## 実行とトレースデータの表示

`TODO`を埋めると次のコマンドで実行することができます。

```
$ go run main.go
```

`trace.out`というトレース情報を記録したファイルが出力されるため、次のコマンドで結果を表示します。

```
$ go tool trace trace.out
```

ブラウザが開くので、`User-defined tasks` -> `Count` -> `Task 1`の順番で開くと次のような結果が表示されれば成功です。

<img src="trace.png" width="500px">

図を見ると、`boil`、`grind`、`brew`が直列に処理されていることがわかります。
`boil`と`grind`は同時に行っても問題ないので、ここを改善すれば早くなりそうです。
