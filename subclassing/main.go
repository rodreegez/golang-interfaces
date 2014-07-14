package main

import "fmt"

type thinger interface {
	thing() string
}

type foo struct {
}

func (f *foo) thing() string {
	return "foo is doing the thing"
}

type bar struct {
}

func (b *bar) thing() string {
	return "bar is doing the thing"
}

func main() {
	t1 := foo{}
	t2 := bar{}
	thingers := []thinger{&t1, &t2}
	for i := range thingers {
		fmt.Println(thingers[i].thing())
	}
}
