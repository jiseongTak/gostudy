package main

import "fmt"

// 이중 배열 순회
func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	} //중괄호가 다른 줄에서 닫히면 마지막에 콤마 찍어야함

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	// 2*5*8 = 80바이트 할당
	// 40/40 나눠서 할당이 아니라 80만큼 덩어리로 메모리 할당
	// type이 결정되면 size 결정 -> 메모리 할당 -> 복사되는 사이즈 결정
}
