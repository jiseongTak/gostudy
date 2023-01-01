package main

import "fmt"

func main() {
	var slice []int //슬라이스 초기화하지 않으면 길이가 0

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 //패닉 발생
	fmt.Println(slice)
}
