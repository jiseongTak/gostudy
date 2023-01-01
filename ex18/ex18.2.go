package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(slice2)
}

/*
make()를 이용한 초기화
var slice = make([]int, 3)
-> 길이 3짜리 int 슬라이스값 각 요솟값은 int 타입의 기본값인 0

append()로 요소 추가
*/
