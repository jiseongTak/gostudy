package main

import (
	"fmt"
	"unsafe"
)

type Student2 struct {
	Age   int32   // 4
	Score float64 // 8
}

func main() {
	var a int = 8
	student := Student2{23, 77.2}

	fmt.Println(unsafe.Sizeof(student), unsafe.Sizeof(a))
	//unsafe.Sizeof() -> 변수의 메모리 공간 크기
}

/*
메모리 정렬(Memory Alignment)
컴퓨터가 데이터에 효과적으로 접근하고자 메모리를 일정 크기 간격으로 정렬하는 것
레지스터 크기가 4바이트인 컴퓨터 -> 32비트 컴퓨터
레지스터 크기가 8바이트인 컴퓨터 -> 64비트 컴퓨터
한 번 연산에 8바이트 크기를 연산할 수 있다. 따라서 데이터가 레지스터 크기와
똑같은 크기로 정렬되어 있으면 더욱 효율적으로 데이터를 읽어올 수 있다

메모리 정렬을 위해서 필드 사이에 공간을 띄우는것을 memory padding이라 한다
*/
