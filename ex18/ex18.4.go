package main

import "fmt"

func changeArray(array2 [5]int) { //배열을 받아서 세 번째 값 변경
	array2[2] = 200
}

func changeSlice(slice2 []int) { //슬라이스를 받아서 세 번째 값 변경
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}

/*
Go 언어에서는 모든 값의 대입은 복사로 일어난다. 함수에 인수로 전달될 때나
다른 변수에 대입할 때나 값의 이동은 복사로 일어난다. 복사는 타입의 값이 복사된다.
포인터는 포인터의 값인 메모리 주소가 복사되고 구조체가 복사될 때는 구조체의 모든 필드가 복사
배열은 배열의 모든 값이 복사된다.
array 와 array2는 메모리 공간이 다른, 즉 완전히 다른 배열이기 때문에 array2의 세 번째 값을
200으로 변경해도 array 배열은 변경되지 않는다.
slice 타입은 []int, []int 내부는 포인터, len, cap 세 개의 필드를 갖는 구조체이다.
포인터는 메모리 주소로 8바이트 길이이고, len과 cap은 각각 int 타입으로 역시 8바이트
슬라이스의 크기는 24바이트, 구조체의 각 필드값이 복사되기 때문에 포인터의 메모리 주솟값도 복사되고
len, cap값도 복사된다.
*/
