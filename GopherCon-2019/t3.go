package main

import "github.com/containous/yaegi/interp"

var src = `
package main

func send(c chan string) { c <- "ping" }

func main() {
		channel := make(chan string)
		go send(channel)
		println(<-channel)
}
`

func main() {
	i := interp.New(interp.Options{})
	i.Eval(src)
}
