package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 20
	fmt.Println(b)
	//C = 20 에러 발생 - 상수는 대입문 좌변에 올 수 없음
	//fmt.Println(&C) 상수 주소 출력
	//상수 앞에 사용하면 상수의 메모리 주솟값을 접근할 수 없기 때문에
	//출력하면 에러 발생
	//상수는 값으로만 동작
	//변수 -> 값, 이름, 타입, 메모리 주소 4가지 속성
	//상수 -> 값, 이름, 타입 3가지 속성
}

//Go언어 상수로 사용될 수 있는 타입
//불리언, 룬, 정수, 실수, 복소수, 문자열
//const 		ConstValue int = 10
//상수 선언 키워드 상수명	   타입   값
