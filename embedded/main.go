package main

import "fmt"

// This is the inner struct that we're going to embed
type inner struct {
	val int
}

// This is the outer struct...
type outer1 struct {
	inner
	name string
}

// And here's a different struct embedding the same inner type
type outer2 struct {
	inner
	name string
}

func main() {
	o1 := outer1{inner{1}, "one"}
	o2 := outer2{inner{2}, "two"}
	fmt.Println(o1.val)
	fmt.Println(o2.name)
}
