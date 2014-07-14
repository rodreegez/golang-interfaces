package main

import "fmt"

type thinger interface {
	thing()
}

type comment struct {
	val string
}

type foo struct {
	comment
}

func (f *foo) thing() {
	fmt.Printf("foo is doing the thing: %v\n", f.val)
}

type bar struct {
	comment
}

func (b *bar) thing() {
	fmt.Printf("bar is doing the thing: %v\n", b.val)
}

func main() {
	t1 := foo{comment{"wee!!!"}}
	t2 := bar{comment{"argh!!"}}
	thingers := []thinger{&t1, &t2}
	for i := range thingers {
		thingers[i].thing()
	}
}
