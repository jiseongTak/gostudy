package main

import (
	"fmt"
	"sort"
)

// int 슬라이스 정렬
func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)

	//Float64s() 함수를 이용하면 float64 슬라이스를 정렬할 수 있다.
}
