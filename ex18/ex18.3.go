package main

import (
	"fmt"
)

func main() {
	var slice []int

	//요소를 하나씩 추가
	for i := 0; i < 11; i++ {
		slice = append(slice, i)
	}

	//한 번에 여러 요소 추가
	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println(slice)
}

/*
reflect 패키지의 SliceHeader 구조체로 내부 구현을 살펴볼 수 있다.
type SliceHeader struct {
	Data uintptr	//실제 배열을 가리키는 포인터
	Len  int		//요소 개수
	Cap  int		//실제 배열의 길이
}

var slice2 = make([]int, 3, 5)
len: 3, cap: 5인 슬라이스
*/
