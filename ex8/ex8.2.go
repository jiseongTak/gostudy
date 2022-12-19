package main

import "fmt"

func main() {
	const pi1 float64 = 3.14159 //상수
	var pi2 float64 = 3.14159   //변수

	//pi1 = 3
	pi2 = 4

	fmt.Printf("원주율: %f\n", pi1)
	fmt.Printf("원주율: %f\n", pi2)
}

//변하면 안 되는 값에 사용
