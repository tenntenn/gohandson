package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func run() error {
	res, err := http.Get(os.Args[1])
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, res.Body)
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
