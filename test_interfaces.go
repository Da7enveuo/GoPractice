package main

import (
	"fmt"
)

// Just a basic implementation of using casting structs to interfaces and calling a few different functions to show some usage

type Test struct {
	hi  string
	bye string
}

type Tester interface {
	tt() string
	tv() string
	th() int
	td() int
}

func (t *Test) tt() string {
	return t.hi
}
func (t *Test) tv() string {
	return t.bye
}
func (t *Test) th() int {
	return len(t.hi)
}
func (t *Test) td() int {
	return len(t.bye)
}

func main() {
	f := &Test{"Whats up!", "Good Bye"}
	fmt.Println("This is the first word: ", f.tt())
	fmt.Println("This is the second word: ", f.tv())
	fmt.Println("This is the length of the first word: ", f.th())
	fmt.Println("This is the length of the second word: ", f.td())
}
