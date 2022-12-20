package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("x:%08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2)
}

/*
<< 왼쪽 시프트
전체 비트를 왼쪽으로 밀어냄
빈 자리는 0으로
*/
