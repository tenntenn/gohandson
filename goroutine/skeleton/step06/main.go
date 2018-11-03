package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type (
	Bean       int
	GroundBean int
	Water      int
	HotWater   int
	Coffee     int
)

const (
	GramBeans          Bean       = 1
	GramGroundBeans    GroundBean = 1
	MilliLiterWater    Water      = 1
	MilliLiterHotWater HotWater   = 1
	CupsCoffee         Coffee     = 1
)

func (w Water) String() string {
	return fmt.Sprintf("%d[ml] water", int(w))
}

func (hw HotWater) String() string {
	return fmt.Sprintf("%d[ml] hot water", int(hw))
}

func (b Bean) String() string {
	return fmt.Sprintf("%d[g] beans", int(b))
}

func (gb GroundBean) String() string {
	return fmt.Sprintf("%d[g] ground beans", int(gb))
}

func (cups Coffee) String() string {
	return fmt.Sprintf("%d cup(s) coffee", int(cups))
}

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

func _main() {
	// 作るコーヒーの数
	const amountCoffee = 20 * CupsCoffee

	taskCtx, task := trace.NewTask(context.Background(), "make coffe")
	defer task.End()

	// 材料
	water := amountCoffee.Water()
	beans := amountCoffee.Beans()

	fmt.Println(water)
	fmt.Println(beans)

	// TODO: taskCtxをベースにしてerrgroup.WithContextで
	// errgroup.Groupとコンテキストを作成する。

	// お湯を沸かす
	var hotWater HotWater
	var hwmu sync.Mutex
	for water > 0 {
		water -= 600 * MilliLiterWater
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				trace.Log(ctx, "boil error", ctx.Err().Error())
				return ctx.Err()
			default:
			}
			hw, err := boil(ctx, 600*MilliLiterWater)
			if err != nil {
				return err
			}
			hwmu.Lock()
			defer hwmu.Unlock()
			hotWater += hw
			return nil
		})
	}

	// 豆を挽く
	var groundBeans GroundBean
	var gbmu sync.Mutex
	for beans > 0 {
		beans -= 20 * GramBeans
		eg.Go(func() error {
			// TODO: キャンセルを検出する

			gb, err := grind(ctx, 20*GramBeans)
			if err != nil {
				return err
			}
			gbmu.Lock()
			defer gbmu.Unlock()
			groundBeans += gb
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	fmt.Println(hotWater)
	fmt.Println(groundBeans)

	// コーヒーを淹れる
	eg, ctx = errgroup.WithContext(taskCtx)
	var coffee Coffee
	var cfmu sync.Mutex
	cups := 4 * CupsCoffee
	for hotWater >= cups.HotWater() && groundBeans >= cups.GroundBeans() {
		hotWater -= cups.HotWater()
		groundBeans -= cups.GroundBeans()
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				trace.Log(ctx, "brew error", ctx.Err().Error())
				return ctx.Err()
			default:
			}

			cf, err := brew(ctx, cups.HotWater(), cups.GroundBeans())
			if err != nil {
				return err
			}
			cfmu.Lock()
			defer cfmu.Unlock()
			coffee += cf
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	fmt.Println(coffee)
}

// お湯を沸かす
func boil(ctx context.Context, water Water) (HotWater, error) {
	defer trace.StartRegion(ctx, "boil").End()
	if water > 600*MilliLiterWater {
		return 0, errors.New("1度に沸かすことのできるお湯は600[ml]までです")
	}
	time.Sleep(400 * time.Millisecond)
	return HotWater(water), nil
}

// コーヒー豆を挽く
func grind(ctx context.Context, beans Bean) (GroundBean, error) {
	defer trace.StartRegion(ctx, "grind").End()
	if beans > 20*GramBeans {
		return 0, errors.New("1度に挽くことのできる豆は20[g]までです")
	}
	time.Sleep(200 * time.Millisecond)
	return GroundBean(beans), nil
}

// コーヒーを淹れる
func brew(ctx context.Context, hotWater HotWater, groundBeans GroundBean) (Coffee, error) {
	defer trace.StartRegion(ctx, "brew").End()

	if hotWater < (1 * CupsCoffee).HotWater() {
		return 0, errors.New("お湯が足りません")
	}

	if groundBeans < (1 * CupsCoffee).GroundBeans() {
		return 0, errors.New("粉が足りません")
	}

	time.Sleep(1 * time.Second)
	// 少ない方を優先する
	cups1 := Coffee(hotWater / (1 * CupsCoffee).HotWater())
	cups2 := Coffee(groundBeans / (1 * CupsCoffee).GroundBeans())
	if cups1 < cups2 {
		return cups1, nil
	}
	return cups2, nil
}
