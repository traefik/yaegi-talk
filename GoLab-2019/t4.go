package main

import (
	"os"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

func main() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Repl(os.Stdin, os.Stdout)
}
