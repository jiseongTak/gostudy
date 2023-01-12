package main

import "fmt"

type opFunc func(a, b int) int

func getOperator2(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			//함수 리터럴을 사용해서 더하기 함수를 정의하고 반환
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator2("*")
	/*
		fn := func(a, b int) int {
			return a + b
		}
	*/

	result := fn(3, 4)
	result2 := func(a, b int) int {
		return a + b
	}(3, 4)

	fmt.Println(result)
	fmt.Println(result2)
}
