package main

import "fmt"

func main() {
	str1 := "가나다라마"
	str2 := "abcde"

	//len()을 통해 문자열 크기(문자열이 차지하는 메모리 크기) 확인
	fmt.Printf("len(str1) = %d\n", len(str1))
	fmt.Printf("len(str2) = %d\n", len(str2))
	//한글은 UTF-8에서 글자당 3바이트를 차지한다.
}
