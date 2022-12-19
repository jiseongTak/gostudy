package main

import "fmt"

func ChangeArray(arr [5]int) {
	arr[3] = 3000
}

// 배열 복사
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
	fmt.Println()

	ChangeArray(a)
	fmt.Println(a[3])
	//ChangeArray 함수의 인수로 a값이 복사되기 때문에 arr와 a는
	//서로 다른 메모리 주소를 가진 다른 배열

	//Go 언어에서 대입 연산자는 우변의 값을 좌변의 메모리 공간에 복사
	//이때 복사되는 크기는 타입 크기와 같습니다.
	//요소를 하나씩 복사하는게 아니라 전체를 통으로 복사
	//복사하는 값의 크기 만큼(a의 크기 40바이트) 그대로 복사
	//Go 언어에서 모든 연산자의 각 항의 타입이 같아야 한다. 대입 연산자도
	//타입의 크기만큼 할당
}
