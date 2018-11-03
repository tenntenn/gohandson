# STEP 4: syncパッケージを使う

## チャネルを使わないゴールーチン間のやり取り

STEP 3では、ゴールーチン間でデータのやりとりを行う場合は、
データの競合を避けるためにチャネルを利用するという説明を行いました。

しかし、データの競合を避けるためには、チャネルを使わずロックをとって排他制御を行う方法もあります。
`sync`パッケージはゴールーチンを跨いだロックなどの便利な機能を提供するパッケージです。

例えば、次のように`sync.Mutex`を用いることでロックを取ることができます。

```go
var (
	count int
	mu    sync.Mutex
)

done := make(chan bool)
go func() {
	for i := 0; i < 10; i++ {
		mu.Lock()
		count++
		mu.Unlock()
		time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
	}
	done <- true
}()

go func() {
	for i := 0; i < 10; i++ {
		mu.Lock()
		count++
		mu.Unlock()
		time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
	}
	done <- true
}()

<-done
<-done
fmt.Println(count)
```

`sync.Mutex`は`Lock`メソッドでロックを取り、`Unlock`メソッドでロックを解除します。
すでにロックが掛かっているMutexに対して`Lock`メソッドを呼び出そうとすると、`Unlock`メソッドが呼び出されるまで処理がブロックされます。

`Unlock`メソッドは`defer`で呼び出すこともありますが、`for`や再帰呼び出しなどでデッドロックを起こす可能性があるため注意が必要です。

## ゴールーチンの待ち合わせ

複数のゴールーチンの処理を待って次の処理に移りたい場合があります。
例えば、お湯を沸かすことと豆を挽くことは並列で行っても問題ありませんが、
コーヒーを淹れるためには、お湯と挽いた豆が揃っている必要があります。

`sync.WaitGroup`はゴールーチンの待ち合わせを行う機能を提供しています。
使い方はとてもシンプルです。
`Wait`メソッドで複数のゴールーチンの処理を待ち合わせることができ、
`Add`メソッドで追加した数だけ`Done`メソッドが呼ばれるまで処理がブロックされます。

例えば、次の例では`wg.Add(1)`が2回実行されているため、`wg.Done()`が2回実行されるまで
`wg.Wait()`で処理をブロックします。

```go
var (
	count int
	mu    sync.Mutex
)

var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		count++
		mu.Unlock()
		time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
	}
	done <- true
}()

wg.Add(1)
go func() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		count++
		mu.Unlock()
		time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
	}
}()

// 2つのゴールーチンの処理が終わるまで待つ
wg.Wait() 
fmt.Println(count)
```

## プログラムの改造

チャネルを使わないようにコーヒーを淹れるプログラムを改造してみましょう。

`boil`関数、`grind`関数、`brew`関数の処理結果はチャネル経由ではなく戻り値で受け取ります。
受け取った戻り値を変数に足す必要がありますが、そのまま足すと競合が起きるためロックを取って足す必要があります。

`boil`関数と`grind`関数の処理は並列で実行しても問題ないため、それぞれゴールーチンで実行し、1つの`sync.WaitGroup`で待ち合わせすることにします。

`brew`関数の処理はゴールーチンで呼ばれますが、別の`sync.WaitGroup`でコーヒーが全て淹れ終わるまで待つことにしましょう。

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

実行時間はSTEP 3と大きく変わりません。
同じように`boil`と`grind`が並列に処理され、その後に`grind`が処理されています。
しかし、`boil`と`grind`のバーの長さが同じではなくなっています。

STEP 3では`boil`からのデータの送信をすべて待った後に、`grind`からのデータを受け取っていましたが、
今回は`sync.WaitGroup`で`boil`も`grind`も関係なく待っていたので、`grind`の方が`boil`の長さに引っ張られずに終了したためです。
