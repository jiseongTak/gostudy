package main

import "fmt"

// 슬라이스 복제
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	//for문 대신 append()를 이용
	slice3 := append([]int{}, slice1...)

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
