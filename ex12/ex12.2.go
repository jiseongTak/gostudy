package main

import "fmt"

func main() {
	//배열 변수 선언 및 초기화
	//var nums [5]int 0 0 0 0 0
	//days := [3]string{"monday", "tuesday", "wednesday"}
	//var temps [5]float64 = [5]float{24.3, 26.7} 나머지는 기본값으로 채움
	//24.3 26.7 0.0 0.0 0.0
	//var s = [5]int{1:10, 3:30} 0 10 0 30 0
	//x := [...]int{10, 20, 30} 초기값 개수만큼 길이 길이가 고정
	//[]int{10, 20, 30} 와 다르다 (slice, 동적배열)
	//배열 선언시 개수는 항상 상수

	const Y int = 3
	//x := 5

	//a := [x]int{1, 2, 3, 4, 5} 에러
	b := [Y]int{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Println(b[i])
	}
}
