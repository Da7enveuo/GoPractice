package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

func main(){
	
	file, _ := os.Open("t.csv")
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	list_ := [][]float64{}
	
	for _, v := range records{
		list_add := []float64{}
		//fmt.Println(float64(v[0]), " | ", v[1])
		a,_ := strconv.ParseFloat(v[0], 64)
		b,_ := strconv.ParseFloat(v[1], 64)
		list_add = append(list_add, a)
		list_add = append(list_add, b)
		list_ = append(list_, list_add)
	}
	//fmt.Println(records)
	fmt.Printf("%T",list_)
}