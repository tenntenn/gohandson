package typing

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	// 乱数の初期化
	rand.Seed(time.Now().Unix())
}

type TypingGame struct {
	wordList []string
	order    []int
}

func New(r io.Reader) (*TypingGame, error) {

	// 問題
	wordList, err := readWordList(r)
	if err != nil {
		return nil, err
	}

	// 出題順
	order := rand.Perm(len(wordList))

	return &TypingGame{
		wordList: wordList,
		order:    order,
	}, nil
}

func (g *TypingGame) Run() {

	// カウントダウン
	for i := 3; i >= 0; i-- {
		fmt.Println(i)
		<-time.After(time.Second)
	}

	var record int
	for i, word := range g.wordList {
		fmt.Printf("%3d問目 %s\n", i+1, word)
		fmt.Print(">")
		select {
		case ans := <-g.inputAnswer():
			if ans == word {
				record++
				fmt.Println("正解！！")
			} else {
				fmt.Println("不正解！！")
			}
		case <-time.After(10 * time.Second):
			fmt.Println("タイムアップ！")
		}
	}

	fmt.Println("正解数：", record)
}

func (g *TypingGame) inputAnswer() <-chan string {
	ch := make(chan string)

	go func() {
		var ans string
		fmt.Scan(&ans)
		ch <- ans
	}()

	return ch
}

func readWordList(r io.Reader) ([]string, error) {
	var wordList []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordList, nil
}
