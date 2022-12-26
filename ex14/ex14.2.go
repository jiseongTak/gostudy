package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1: %p\n", p1)
	fmt.Printf("p2: %p\n", p2)
	fmt.Printf("p3: %p\n", p3)
	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
}

//포인터 변숫값을 초기화하지 않으면 기본값은 nil
