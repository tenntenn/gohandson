package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime/trace"
	"strings"
)

func main() {
	if err := runWithTrace(os.Args[0], os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func runWithTrace(cmd string, args []string) (rerr error) {
	f, err := os.Create("trace.out")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			if rerr != nil {
				rerr = err
			}
		}
	}()

	if err := trace.Start(f); err != nil {
		return err
	}
	defer trace.Stop()

	return run(cmd, args)
}

func run(cmd string, args []string) error {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()
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
