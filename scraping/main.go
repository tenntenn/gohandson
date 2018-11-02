package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args[0], os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(cmd string, args []string) error {
	fs := flag.NewFlagSet(cmd, flag.ExitOnError)
	dir := fs.String("o", ".", "output dir")
	allowHosts := fs.String("allow", "", "allow hosts")
	format := fs.String("format", "", "convert format")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() == 0 {
		return errors.New("第1引数は指定してください")
	}

	u, err := url.Parse(fs.Arg(0))
	if err != nil {
		return err
	}

	s := New(*dir)
	s.AllowHost = strings.Split(*allowHosts, ",")
	if IsAllowFormat(*format) {
		s.Format = ImageFormat(*format)
	}
	if err := s.Visit(u); err != nil {
		return err
	}

	return nil
}
