package main

import "github.com/containous/yaegi/interp"

const src = `package foo
func Bar(s string) string { return s +"-Foo" }`

func main() {
	i := interp.New(interp.Options{})
	i.Eval(src) // load "foo" package

	v, _ := i.Eval("foo.Bar")                  // access foo.Bar symbol
	bar := v.Interface().(func(string) string) // type assert value

	r := bar("Kung")
	println(r)
}
