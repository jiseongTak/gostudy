package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]
	slice2 := array[1:3:4]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}

/*
슬라이싱은 배열의 일부를 집어내는 기능을 말한다.
슬라이싱 기능을 사용하면 그 결과로 슬라이스를 반환한다.
array[startIdx:endindex]
배열의 시작인덱스부터 끝인덱스-1까지의 배열 일부를 나타내는 슬라이스 반환
슬라이싱하면 배열의 일부를 가리키는 슬라이스를 반환, 새로운 배열이 만들어지는게 아니라
배열의 일부를 포인터로 가리키는 슬라이스를 만든다.

슬라이싱 기능은 배열뿐 아니라 슬라이스도 가능
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice2[1:2]
처음부터 슬라이싱하면 시작인덱스 생략 가능
slice1[0:3] -> slice1[:3]
끝까지 슬라이싱한면 끝인덱스 생략 가능
slice1[2:len(slice1)] -> slice1[2:]
전체 슬라이싱
array[:] 배열 전체를 가리키는 슬라이스를 만들고 싶을 때 주로 사용
인덱스를 2개만 사용할 때 cap은 배열의 전체 길이에서 시작인덱스를 뺀 값
슬라이싱할 때 인덱스를 3개 사용해서 cap까지 조절할 수 있다.
slice[시작인덱스:끝인덱스:최대인덱스]
cap 값은 최대인덱스 - 시작인덱스
slice1 := []int{1,2,3,4,5}
slice2 := slice1[1:3:4]
*/
