package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	//함수 타입 변수 사용 func(int, int) int
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("+")

	var result = operator(3, 4)
	fmt.Println(result)
}

/*
별칭으로 함수 정의 줄여쓰기
type opFunc func(int, int) int
func getOperator(op string) opFunc

함수 정의에서 매개변수명은 생략가능
func(int, int) int / func(a int, b int) int
*/
