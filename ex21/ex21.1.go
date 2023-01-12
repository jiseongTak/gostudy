package main

import "fmt"

// ...키워드 사용
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
	//슬라이스 타입으로 처리
	//Println() 처럼 여러 타입 섞어 쓰는법 -> interface{} 빈 인터페이스 사용
}
