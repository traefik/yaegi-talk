package main

import "github.com/containous/yaegi/interp"

func main() {
	i := interp.New(interp.Options{})
	i.Eval(`println("Hello World")`)
}
