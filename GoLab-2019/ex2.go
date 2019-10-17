package main

import "github.com/containous/yaegi/interp"

func main() {
	i := interp.New(interp.Options{})
	i.Eval(`a := 12 + 5`)
	i.Eval(`println(a)`)
}
