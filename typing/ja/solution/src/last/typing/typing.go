package typing

import (
	"bufio"
	"io"
)

type TypingGame struct {
	wordlist bufio.Scanner
}

func New(r io.Reader) *TypingGame {
	return &TypingGame{
		wordlist: r,
	}
}

func (g *TypingGame) Run() error {
}
