package main

import "fmt"

func main() {
	//append()를 사용할 때 발생하는 예기치 못한 문제1
	slice1 := make([]int, 3, 5) //len: 3, cap: 5

	slice2 := append(slice1, 4, 5)
	//cap()함수를 이용해 슬라이스 capacity 값을 알 수 있다.
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 //slice2까지 바뀐다.

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500) //slice2까지 바뀐다.

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}

/*
append() 함수가 호출되면 먼저 슬라이스에 값을 추가할 수 있는 빈 공간이 있는지 확인
남은 빈 공간은 실제 배열 길이 cap에서 슬라이스 요소 개수 len을 뺀 값
*/
