# STEP 1: ゴールーチンを使わずに処理する

## コーヒーを淹れるプログラムを作ろう

このハンズオンでは、題材としてコーヒーを淹れるプログラムを作ります。
実際にはコーヒーを淹れるわけではありませんが、一連の手順をいかに並行処理で効率化していくかということを学ぶことができます。

コーヒーを淹れるには次の手順が必要でしょう。

* お湯を沸かす
* 豆を挽く
* コーヒーを淹れる

## データ型を定義しよう

コーヒーを淹れるには、お湯と挽かれたコーヒー豆の粉が必要になります。
お湯を沸かすには水が必要になり、コーヒー豆の粉を手に入れるには、コーヒー豆が必要です。
つまり、各手順で次のデータの変換が行われます。

* お湯を沸かす: 水 -> お湯
* 豆を挽く: コーヒー豆 -> 挽かれたコーヒー豆の粉
* コーヒーを淹れる: お湯, 挽かれたコーヒー豆の粉 -> コーヒー

このプログラムには、水、お湯、豆、挽かれた豆、コーヒーの5つの種類のデータが存在することがわかります。
これらのデータを表すために、データ型を作成しましょう。
データ型は`type`を用いて定義することができます。

```go
type (
	Bean       int // 豆
	GroundBean int // 挽かれた豆
	Water      int // 水
	HotWater   int // お湯
	Coffee     int // コーヒー
)
```

また、利便性のために次のようにいくつか定数も用意しておきます。
こうすることで、`10 * GramBeans`のように記述することができます。

```go
const (
	GramBeans          Bean       = 1
	GramGroundBeans    GroundBean = 1
	MilliLiterWater    Water      = 1
	MilliLiterHotWater HotWater   = 1
	CupsCoffee         Coffee     = 1
)
```

つぎに、N杯のコーヒーを淹れるために必要な材料の分量を返すメソッドを用意します。
メソッドは、`Coffee`型のメソッドとして設けます。
こうすることで、2杯のコーヒーに必要な水の分量を`(2 * Cupscoffee).Water()`のように取得することができます。

```go
// 1カップのコーヒーを淹れるのに必要な水の量
func (cups Coffee) Water() Water {
        return Water(180*cups) / MilliLiterWater
}

// 1カップのコーヒーを淹れるのに必要なお湯の量
func (cups Coffee) HotWater() HotWater {
        return HotWater(180*cups) / MilliLiterHotWater
}

// 1カップのコーヒーを淹れるのに必要な豆の量
func (cups Coffee) Beans() Bean {
        return Bean(20*cups) / GramBeans
}

// 1カップのコーヒーを淹れるのに必要な粉の量
func (cups Coffee) GroundBeans() GroundBean {
        return GroundBean(20*cups) / GramGroundBeans
}
```

## お湯を沸かす

お湯を沸かす関数`boil`を作成します。
`boil`は一定時間立つと、引数で与えた分量と同じ量のお湯を返します。

```go
// お湯を沸かす
func boil(water Water) HotWater {
	time.Sleep(400 * time.Millisecond)
	return HotWater(water)
}
```

## コーヒー豆を挽く

コーヒー豆を挽く関数`grind`を作ります。
`grind`は引数にコーヒー豆を受け取り、一定時間後に挽いた豆を返します。

```go
// コーヒー豆を挽く
func grind(beans Bean) GroundBean {
	time.Sleep(200 * time.Millisecond)
	return GroundBean(beans)
}
```

## コーヒーを淹れる

コーヒーを淹れる関数`brew`を作ります。
`brew`は引数にお湯と挽いた豆を受け取り、一定時間後にコーヒーを返します。

```go
// コーヒーを淹れる
func brew(hotWater HotWater, groundBeans GroundBean) Coffee {
	time.Sleep(1 * time.Second)
	// 少ない方を優先する
	cups1 := Coffee(hotWater / (1 * CupsCoffee).HotWater())
	cups2 := Coffee(groundBeans / (1 * CupsCoffee).GroundBeans())
	if cups1 < cups2 {
		return cups1
	}
	return cups2
}
```

## 処理をまとめる

`main`関数から、`boil`関数、`grind`関数、`brew`関数を順に呼び、コーヒーを淹れます。
材料は次のように20杯分のコーヒーとしてます。

```go
// 作るコーヒーの数
const amountCoffee = 20 * CupsCoffee

// 材料
water := amountCoffee.Water()
beans := amountCoffee.Beans()
```

一度に沸かせるお湯の量や挽ける豆の量、淹れれるコーヒーの量は次のように決まっているものとします。

* 一度に沸かせるお湯の量: 600[ml]
* 一度に挽ける豆の量: 20[g]
* 一度に淹れれるコーヒー: 4杯

`boil`関数、`grind`関数、`brew`関数を複数回呼び出すことで、20杯のコーヒーを淹れることができます。

## 実行

`TODO`を埋めると次のコマンドで実行することができます。

```
$ go run main.go
```

次のように表示されれば成功です。

```
3600[ml] water
400[g] beans
3600[ml] hot water
400[g] ground beans
20 cup(s) coffee
```
