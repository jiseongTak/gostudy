package main

import "fmt"

const PI = 3.14              //타입 없는 상수
const FloatPI float64 = 3.14 //float64 타입 상수

func main() {
	var a int = PI * 100
	//var b int = FloatPI * 100 //타입 오류 발생

	fmt.Println(a)
	fmt.Println(FloatPI)
}

//상수 선언 시 타입을 명시하지 않을 수 있다
//타입 없는 상수는 변수에 복사될 때 타입이 정해지기 때문에
//여러 타입에 사용되는 상숫값을 사용할 때 편리하다

/*
컴퓨터에서 리터럴이란 고정된 값, 값 자체로 쓰인 문구로 볼 수 있다
	var str string = "Hello World"
	var i int = 0
	i = 30
위 코드같이 고정된 값 자체로 쓰인 문구가 리터럴
Go 언어에서 상수는 리터럴과 같이 취급
그래서 컴파일될 때 상수는 리터럴로 변환되어 실행파일에 쓰인다
상수 표현식 역시 컴파일 타임에 실제 결괏값 리터럴로 변환하기 때문에
상수 표현식 계산에 CPU 자원을 사용하지 않는다
const PI = 3.14
var a int = PI * 100
위 구문은 컴파일 타임에 아래와 같이 변환
var a int = 314
상수의 메모리 주솟값에 접근할 수 없는 이유 역시 컴파일 타임에 리터럴로
전환되어서 실행 파일에 값 형태로 쓰이기 때문
*/
