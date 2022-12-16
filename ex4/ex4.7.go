package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b //4266663.334329
	var d float32 = c * 3 //12799990.002987

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) //4266663
	fmt.Println(d) //12799989
}

//1234.523
//3456.123
//4.266663e+06
//1.2799989e+07
//자리수 넘어가면 오차가 발생할 수 있다

//float32 소수부 7자리
//float64 소수부 15자리
