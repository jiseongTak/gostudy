package main

import "fmt"

func main() {
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
	//str[i]처럼 인덱스로 접근하면 요소의 타입은 uint8 즉 byte 이다.
	//3바이트 크기인 한글은 깨져서 나옴
}

/*
문자열 순회
1. 인덱스를 사용한 바이트 단위 순회
2. []rune으로 타입 변환 후 한 글자씩 순회
3. range 키워드를 이용한 한 글자씩 순회
*/
