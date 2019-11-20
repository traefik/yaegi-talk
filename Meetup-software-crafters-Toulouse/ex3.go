package main

import (
	"github.com/containous/yaegi/interp"
	//"github.com/containous/yaegi/stdlib"
)

const src = `
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}`

func main() {
	i := interp.New(interp.Options{})
	//i.Use(stdlib.Symbols)
	if _, err := i.Eval(src); err != nil {
		panic(err)
	}
}
