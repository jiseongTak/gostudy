package main

import "fmt"

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}

func Add(a int, b int) int {
	return a + b
}

//첫 글자가 대문자인 함수는 패키지 외부로 공개되는 함수
//Go 언어에서 함수 코드 블록의 시작을 알리는 중괄호 {가 함수를 정의하는
//라인과 항상 같은 줄에 있어야 함
//함수를 호출할 때 입력하는 값을 argument, 인수
//함수가 외부로부터 입력받는 변수를 parameter, 매개변수

/*
func 		 Add	(a int, b int) int {
함수정의 키워드 함수명	매개변수		   반환타입
	return a + b
}
함수를 호출하며 입력한 값은 실제 함수에 어떻게 전달?
보낸 값을 그대로 사용하는 것이 아니라 값을 복사해 사용
인수는 매개변수로 복사된다.
*/