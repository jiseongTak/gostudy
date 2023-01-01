package main

import "fmt"

// 요소 삭제
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2 //삭제할 인덱스

	//for i := idx + 1; i < len(slice); i++ {
	//	slice[i-1] = slice[i]
	//}
	//fmt.Println(slice)
	//
	//slice2 := slice[:len(slice)-1]
	//fmt.Println(slice2)

	//append() 이용
	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice)
}
