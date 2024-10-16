package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	ext  string
	size int64
	list bool
	del  bool
	wLog io.Writer
}

func main() {
	var (
		f   = os.Stdout
		err error
	)
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	logFile := flag.String("log", "", "Log deletes to this file")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	del := flag.Bool("del", false, "Delete files")
	flag.Parse()
	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		wLog: f,
	}
	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}
}
