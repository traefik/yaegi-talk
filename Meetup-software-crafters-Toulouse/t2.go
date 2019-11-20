package main

import (
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

func main() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	r, err := i.Eval(`println(42)`)
}
