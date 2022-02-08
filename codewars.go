package main

import (
	"fmt"
	"strings"
)

// Code wars challenge

func main() {
	s := "The_Stealth_Warrior"
	//h := []string{s}
	return_string := []string{}
	wc := 0
	first := false
	for count, letter := range s {

		if string(letter) == "-" || string(letter) == "_" {
			fmt.Println("Hi")
			fmt.Println(string(s[count]))
			switch first {
			case false:
				return_string = append(return_string, s[count-wc:count])
			case true:
				f_l := strings.ToUpper(string(s[count+1]))
				return_string = append(return_string, f_l+s[count-wc+1:count])
				first = false
			}
			wc = 0
		} else if len(s) == count+1 {
			fmt.Println("Hi")
			return_string = append(return_string, s[count-wc+1:count+1])
		} else {
			fmt.Println("Hi")
			wc++
		}
	}
	fmt.Println(return_string)
}
