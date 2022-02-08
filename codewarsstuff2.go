package main

import "fmt"

// codewars practice
func main() {
	data := "iiiso"
	val := 0
	return_array := []int{}
	for i := 0; i < len(data); i++ {
		if data[i] == 'i' {
			val++
		} else if data[i] == 'd' {
			val--
		} else if data[i] == 's' {
			val = val * val
		} else if data[i] == 'o' {
			return_array = append(return_array, val)
		}
	}
	fmt.Println(return_array)
}
