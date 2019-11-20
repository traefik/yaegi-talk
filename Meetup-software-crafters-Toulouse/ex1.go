package main

import (
	"fmt"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

func main() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Eval(`import "fmt"
	
	type Animal struct {
		Name string
		Age  uint
	}

	func (a Animal) String() string { return fmt.Sprintf("%v (%d)", a.Name, a.Age) }
	func (a Animal) Get() interface{} { return a }
	
	var Gopher fmt.Stringer = Animal{ "Gopher", 2 }
	`)
	v, _ := i.Eval("Gopher")
	f := v.Interface().(fmt.Stringer)
	fmt.Println(f)
	//i.Eval(`fmt.Println("hello")`)
}
