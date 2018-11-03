# STEP 6: コンテキストとキャンセル処理

## キャンセル処理

とあるゴールーチンでエラーが発生した場合に、
他のゴールーチンの処理をキャンセルしたい場合があります。

Goでは、ゴールーチンのキャンセル処理のために`context.Context`を用います。
`context.WithCancel`関数でラップしたコンテキストは第2戻り値返されたキャンセル用の関数が呼び出されるか、親のコンテキストがキャンセルされるとキャンセルされます。
キャンセルされたことを知るには、`context.Context`の`Done`メソッドから返ってくるチャネルを用います。

次の例では、2つのゴールーチンを立ち上げ、キャンセルを`Done`メソッドのチャネルで伝えています。
`select`は複数のチャンネルへの送受信を待機することのできる構文で、どのケースのチャネルも反応しない場合は、`default`が実行されます。

```go
func main() {
	root := context.Background()
	ctx1, cancel := context.WithCancel(root)
	ctx2, _ := context.WithCancel(ctx1)

	var wg sync.WaitGroup
	wg.Add(2) // 2つのゴールーチンが終わるの待つため

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("cancel goroutine1")
				return
			default:
				fmt.Println("waint goroutine1")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("cancel goroutine2")
				return
			default:
				fmt.Println("waint goroutine2")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
```

## errgroup.Groupを使ったキャンセル処理

`errgroup.Group`は`errgroup.WithContext`を用いることで、エラーが起きた際に処理をキャンセルすることができます。

```go
func main() {
	root := context.Background()
	eg, ctx := errgroup.WithContext(root)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel goroutine1")
				return nil
			default:
				fmt.Println("waint goroutine1")
				time.Sleep(500 * time.Millisecond)
			}
		}
	})

	eg.Go(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("error")
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
```

## プログラムの改造

`errgroup.WithContext`を用いてエラーが発生した場合のキャンセル処理をハンドリングしましょう。

## 実行

`errgroup`パッケージは外部パッケージであるため、`go get`コマンドでインストールする必要があります。

```
$ go get -u golang.org/x/sync/errgroup
```

次のコマンドで実行することができます。

```
$ go run main.go
```

`boil`関数に渡す水の量を2倍にしたりしてエラーが発生するようにしてみましょう。
