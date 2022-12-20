package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}

	fmt.Println("cnt:", cnt)
}

/*
&& 연산은 좌변이 false이면 검사하지 않고 false 처리
|| 연산 좌변이 true이면 우변은 검사하지 않고 true 처리
이를 쇼트서킷(short-circuit)이라 한다
*/
