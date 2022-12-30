package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str) //각 글자들이 배열의 요소로 바뀜

	fmt.Printf("len(str) = %d\n", len(str))
	fmt.Printf("len(runes) = %d\n", runes)
	fmt.Printf("len(runes) = %d\n", len(runes))
	fmt.Println(string(runes))
	fmt.Printf("len(runes) = %d\n", len(string(runes)))

	//len(str)은 문자열의 총 바이트 길이인 12를 반환
	//len(runes)은 배열의 요소 개수인 8을 반환
}

/*
string 타입과 []rune 타입은 모두 문자들의 집합을 나타내기 때문에
상호 타입 변환이 가능하다. 타입 변환을 할 경우 rune 배열의 각 요소에는 문자열의
각 글자가 대입된다. 이를 통해 문자열의 글자 수를 알 수 있다.

string 타입과 []byte 타입은 상호 타입 변환이 가능
[]byte 는 byte 즉 1바이트 부호 없는 정수 타입의 가변길이 배열
문자열이란 것도 결국 메모리에 있는 데이터이고, 메모리는 1바이트 단위로 저장되기 때문에
모든 문자열은 1바이트 배열로 변환 가능하다.
*/
