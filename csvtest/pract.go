package main

import (
	"fmt"
	"math"
	"encoding/csv"
	"os"
	"strconv"
)
// Linear prediction model, poorly written

func main(){
	file, _ := os.Open("tester.csv")
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	matrix := [][]float64{}
	
	for _, v := range records{
		list_add := []float64{}
		//fmt.Println(float64(v[0]), " | ", v[1])
		a,_ := strconv.ParseFloat(v[0], 64)
		b,_ := strconv.ParseFloat(v[1], 64)
		c,_ := strconv.ParseFloat(v[2], 64)
		list_add = append(list_add, a)
		list_add = append(list_add, b)
		list_add = append(list_add, c)
		matrix = append(matrix, list_add)
	}
	var x []float64 = []float64{}
	var x_ []float64 = []float64{}
	var y []float64 = []float64{}

	var x_var []float64 = []float64{}
	var x_var_ []float64 = []float64{}
	var y_var []float64 = []float64{}
	
	var x_distance float64 = 0
	var x_2_distance float64 = 0
	var y_distance float64 = 0
	
	var x_total float64 = 0
	var x_2_total float64 = 0
	var y_total float64 = 0
	
	var y_dis float64 = 0
	var x_2_dis float64 = 0
	var x_dis float64 = 0

	
	//matrix := [][]float64{{9,1.453},{10,1.047},{12,.739},{13,.404}}
	//matrix = [][]float64{{1,1},{2,3},{4,5},{5,7}}

	for _, row := range matrix{
		fmt.Println("X: ",row[0]," | X2: ", row[1], " | Y: ", row[2])
		x = append(x, row[0])
		x_total = x_total + row[0]
		x_ = append(x_, row[1])
		x_2_total = x_2_total + row[1]
		y = append(y,row[2])
		y_total = y_total + row[2]
	}
	//fmt.Println(matrix)
	fmt.Println("X vars mean: ", float64(x_total/float64(len(matrix))))
	fmt.Println("X 2 vars mean: ", float64(x_2_total/float64(len(matrix))))
	fmt.Println("Y vars mean: ", float64(y_total/float64(len(matrix))))
	 
	for _, column := range matrix{
		x_var = append(x_var, column[0] - x_total/float64(len(matrix)))
		x_var_ = append(x_var_, column[1] - x_2_total/float64(len(matrix)))
		y_var = append(y_var, column[2] - y_total/float64(len(matrix)))
	}
	
	var covariance_numerator float64 = 0
	var x2y_dis float64 = 0
	var xy_dis float64 = 0
	
	for count, vari := range x_var{
		x_distance = x_distance + (vari * vari) 
		y_distance = y_distance + (y_var[count] * y_var[count])
		x_2_distance = x_2_distance + (x_var_[count] * x_var_[count])
		x_dis = x_dis + vari
		x_2_dis = x_2_dis + x_var_[count]
		//fmt.Println(x_dis)
		//fmt.Println(x_dis, " | ", vari, " + ", vari)
		y_dis = y_dis + y_var[count]
		//fmt.Println(y_dis, " | ", y_var[count]," + ", y_var[count])
		xy_dis = xy_dis + (vari * y_var[count])
		x2y_dis = x2y_dis + (x_var_[count] * y_var[count])
	}
	
	fmt.Println("Covariance Numerator (AKA Pop. Covariance): ",xy_dis/float64(len(x)))
	fmt.Println("Covariance Numerator (AKA Samp. Covariance): ",xy_dis/float64(len(x) -1 ))
	fmt.Println("Covariance #2 Numerator (AKA Pop. Covariance): ",x2y_dis/float64(len(x)))
	fmt.Println("Covariance #2 Numerator (AKA Samp. Covariance): ",x2y_dis/float64(len(x) -1))
	covariance_numerator = covariance_numerator + (x_distance * y_distance)
	x_v := float64(float64(x_dis)/float64(len(matrix)))
	x_y := float64(y_distance/float64(len(x)))
	x2_v := float64(float64(x_2_dis)/float64(len(matrix)))
	fmt.Printf("X standard deviation: %v\n", math.Sqrt((x_v)))
	fmt.Println("X 2 standard deviation: ", math.Sqrt((x2_v)))
	fmt.Println("Y standard deviation: ", math.Sqrt((x_y)))	
	
	new_x := 0.00
	new_x2 := 0.00
	new_y := 0.00
	list_vals := []float64{}
	list_vals2 := []float64{}
	for count, _ := range matrix{
		//fmt.Println(column[0], " - ",float64(x_total/float64(len(matrix))), " / ", math.Sqrt(math.Abs(x_v)) )
		new_x = (x_var[count])/math.Sqrt((x_v))
		new_x2 = (x_var_[count])/math.Sqrt((x2_v))
		//fmt.Println("New x ",new_x)
		//fmt.Println(column[1], " - ",float64(y_total/float64(len(matrix))), " / ", math.Sqrt(x_y) )
		new_y = (y_var[count])/math.Sqrt(x_y)
		//fmt.Println("New y ",new_y)
		list_vals = append(list_vals, new_x * new_y)
		list_vals2 = append(list_vals2, new_x2 * new_y)
		//fmt.Println(new_x, " * ", new_y, " = ", new_x * new_y)
	}
	tots := float64(0)
	tots2 := float64(0)
	for c,a := range list_vals{
		tots = float64(tots) + a
		tots2 = float64(tots2) + list_vals2[c]
	}

	corr := tots/float64(len(matrix))
	corr2 := tots2/float64(len(matrix))
	fmt.Println("Correlation Coefficent: ",corr)
	fmt.Println("Correlation Coefficent2 : ", corr2)
	if corr < -.8{
		fmt.Println("\nThe data is strongly correlated to the negative side.")
	}else if corr > .8{
		fmt.Println("\nThe data is strongly correlated to the positive side.")
	}else if corr > .5 && corr < .8{
		fmt.Println("\nThe data is leaning to the positive side.")
	}else if corr < -.8 && corr > -.5{
		fmt.Println("\nThe data is leaning to the negative side.")	
	}else if corr < .5 && corr > -.5{
		fmt.Println("\nThe data is not strongly correlated in either direction.")
	}
	if corr2 < -.8{
		fmt.Println("\nThe data is strongly correlated to the negative side.")
	}else if corr2 > .8{
		fmt.Println("\nThe data is strongly correlated to the positive side.")
	}else if corr2 > .5 && corr2 < .8{
		fmt.Println("\nThe data is leaning to the positive side.")
	}else if corr2 < -.8 && corr2 > -.5{
		fmt.Println("\nThe data is leaning to the negative side.")	
	}else if corr2 < .5 && corr2 > -.5{
		fmt.Println("\nThe data is not strongly correlated in either direction.")
	}
	//((x_total * x_distance) - (x_total * (x_total * y_total)))/
	// linear regression =   y = a + bX
	// a = ((sum y * sum x^2) - (sum x)( sum x*y))/ n(sum x^2) -(sum x)^2
	constant := ((y_total * x_distance) - (x_total * xy_dis)) / ((float64(len(matrix)) * x_distance) - x_distance)
	constant2 := ((x_2_total * x_2_distance)- (x_2_total * x2y_dis))/((float64(len(matrix)) * x_2_distance) - x_2_distance)
	// b = n(sum x * y) - (sum x)(sumy)/n(sum x^2) - (sum x^2)
	// X is the index of future variable to predict
	regression_coefficent := xy_dis / x_distance
	regression_coefficent2 := x2y_dis /x_2_distance
	fmt.Println("The Constant is: ", constant)
	fmt.Println("The second constant is: ", constant2)
	fmt.Println("The Regression Coefficent is: ", regression_coefficent)
	fmt.Println("The Regression Coefficent2 is: ", regression_coefficent2)
	var input float64
	fmt.Println("Please Enter the x variable to predict: ")
	fmt.Scanln(&input)
	fmt.Println("The Y variable should be: ", constant + (regression_coefficent * input))
	//B1 = b1 = Σ [ (xi – x)(yi – y) ] / Σ [ (xi – x)2]
}
