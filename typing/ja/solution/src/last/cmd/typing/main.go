package main

import (
	"fmt"
	"os"

	"last/typing"
)

// step1: go install
// step2: 配列とスライス
/// 配列とスライスの違い
/// for range
// step3: 標準入力とbufio.Scanner
/// io.Readerとio.Writerについて
/// os.Stdinやos.Stdoutについて
/// bufio.Scannerについて
// step4: ファイル読み込み
/// os.Openとos.Fileについて
// step5: コマンドライン引数
/// os.Argsについて
// step6: timeパッケージ
/// time.Nowについて
// step7: ゴールーチンとチャネル その1
/// time.Tickについて
// step8: ゴールーチンとチャネル その2
/// time.Afterについて
/// ゴールーチンの作り方について

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	game, err := typing.New(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	game.Run()
}
