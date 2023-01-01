package main

import "fmt"

// 요소 추가
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	idx := 2
	////맨 뒤에 요소 추가
	//slice = append(slice, 0)
	//for i := len(slice) - 2; i >= idx; i-- {
	//	slice[i+1] = slice[i]
	//}
	//slice[idx] = 100

	//append()이용
	//slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	//불필요한 메모리 사용이 없도로 코드 개선하기
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100

	fmt.Println(slice)
}
