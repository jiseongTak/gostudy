package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
}

/*
for 초기문; 조건문; 후처리 {
	코드 블록 //조건문이 true인 경우 수행
}
초기문 생략
for ; 조건문; 후처리 {
	코드 블록
}
후처리 생략
for 초기문; 조건문; {
	코드 블록
}
조건문만 있는 경우
for ; 조건문; {
	코드 블록
}
for 조건문 {
	코드 블록
}
*/
