package main

import "github.com/containous/yaegi/interp"

const src = `
package main

import "github.com/containous/yaegi/interp"

func main() {
	i := interp.New(interp.Options{})
	i.Use(interp.Symbols)
	if _, err := i.Eval("println(21 * 2)"); err != nil {
		panic(err)
	}
}`

func main() {
	i := interp.New(interp.Options{})
	i.Use(interp.Symbols)
	if _, err := i.Eval(src); err != nil {
		panic(err)
	}
}
