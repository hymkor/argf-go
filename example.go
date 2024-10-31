//go:build ignore

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hymkor/argf-go"
)

func main() {
	in := argf.New(os.Args[1:])
	// or argf.New(flag.Args())
	if _, err := io.Copy(os.Stdout, in); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
