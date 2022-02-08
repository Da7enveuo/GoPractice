package main

import (
	"fmt"
)

func main() {
	peopleInLine := []int{25, 25, 50, 50, 100}
	total := 0
	msg := ""
	for i := 0; i < len(peopleInLine); i++ {
		if peopleInLine[i] == 25 {
			total += 25
		} else {
			remainder := peopleInLine[i] - 25
			fmt.Println(remainder)
			fmt.Println(total)
			if remainder > total {
				msg = "NO"
			} else {
				msg = "YES"
			}
		}
	}
	fmt.Println(msg)
}
