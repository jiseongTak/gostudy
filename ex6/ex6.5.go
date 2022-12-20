package main

import "fmt"

func main() {
	var x int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, x+1)
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, x+2)
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, x+3)

	var y int8 = -128
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
	fmt.Printf("y\t= %4d, %08b\n", y, y)
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, y-1)
}

/*
정수가 정수 타입의 범위를 벗어난 경우 값이 비정상으로
변화하는 현상을 오버플로라 한다
반대는 언더플로
가장 작은 값에서 -1을 했을 때 가장 큰 값으로 바뀜
*/
