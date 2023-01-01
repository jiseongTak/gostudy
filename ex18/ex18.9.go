package main

import "fmt"

// copy()를 이용한 복제
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10) //len:3, cap:10 슬라이스
	slice3 := make([]int, 10)    //len:10, cap:10 슬라이스

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
}

/*
func copy(dst, src []Type) int
첫 번째 인수로 복사한 결과를 저장하는 슬라이스 변수
두 번째 인수로 복사 대상이 되는 슬라이스 변수
반환값은 실제로 복사된 요소 개수
실제 복사되는 요소 개수는 목적직의 슬라이스 길이와 대상의 슬라이스 길이 중 작은 개수만큼
cap 개수는 영향을 주지 않는다.
*/
