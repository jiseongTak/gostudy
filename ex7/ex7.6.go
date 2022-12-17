package main

import "fmt"

func printNo(n int) {
	if n == 0 {
		return
	} // 재귀 호출 탈출 조건
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	printNo(3)
}

//재귀 호출을 사용할 때는 항상 탈출 조건을 정해야 한다
//재귀 호출이 종료되는 시점을 명확히 하지 않으면 재귀 호출이
//무한히 반복되어 프로그램이 비정상 종료된다
/*
	3
	2
	1
	After 1
	After 2
	After 3
*/
