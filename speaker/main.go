// Simple introduction to interfaces adapted from
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import "fmt"

// A "Speaker" is anything that conforms to the Speaker interface, i.e. defines
// a `Speak` method
type Speaker interface {
	Speak() string
}

// define a couple animals, make them implement the Speak interface
type Dog struct {
}

func (d Dog) Speak() string {
	return "woof!\n"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "meow!\n"
}

// This function taks a Speaker, s, as an argument and prefixes it's output
// from calling `Speak()`
//
// This is why interfaces are cool - Anything can define a `Speak` method, and
// therefore can be passed to `animalSays`.
func animalSays(s Speaker) string {
	spoke := s.Speak()
	return "this animal says " + spoke
}

func main() {
	speakers := []Speaker{Dog{}, Cat{}}
	for _, speaker := range speakers {
		noise := animalSays(speaker)
		fmt.Println(noise)
	}
}
