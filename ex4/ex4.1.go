package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)
}

// 변수 선언 -> var 변수명 타입 = 초기값
// 변수 선언은 컴퓨터에게 값을 저장할 공간을 마련하라고 명령을 내리는 것
// -> 메모리 할당
