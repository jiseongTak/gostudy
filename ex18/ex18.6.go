package main

import "fmt"

func main() {
	//append()를 사용할 때 발생하는 예기치 못한 문제2
	slice1 := []int{1, 2, 3} //len: 3, cap: 3

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
append() 함수가 호출되면 빈공간이 충분한지 확인하고 만약 빈공간이 충분하지 않으면
새로운 더 큰 배열을 마련한다. 일반적으로 기존 배열의 2배 크기로 마련한다.
그런 뒤 기존 배열의 요소를 모두 새로운 배열에 복사한다. 그리고 새로운 배열의
맨 뒤에 새 값을 추가한다.
*/
