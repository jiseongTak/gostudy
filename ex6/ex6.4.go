package main

import "fmt"

func main() {
	var (
		x int8  = 16
		y int8  = -128
		z int8  = -1
		w uint8 = 128
	)

	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", uint8(y), uint8(y>>2), y>>2)
	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), w>>2)
}

/*
>> 오른쪽 시프트
전체 비트를 왼쪽으로 밀어냄
부호 있는 정수이면 왼쪽 비트에 부호와 같은 값
부호 없는 정수이면 0으로
음수이면 최상위 비트가 1이므로 1로, 양수이면 0으로
*/
