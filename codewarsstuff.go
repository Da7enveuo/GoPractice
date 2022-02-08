package main

import (
	"strings"
)

// random codewars challenge
func main() {
	var word string = "Hey fellow warriors"
	var spaces int = 0
	w := []string{}
	for key, value := range word {
		if key == ' ' {
			spaces++
		} else if spaces == 1 && key-strings.Index(word, string(value)) == 0 {

		}
	}
}
