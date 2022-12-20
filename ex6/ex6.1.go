package main

import "fmt"

/*
Go 언어에서 모든 연산자의 각 항의 타입은 항상 같아야 한다(시프트 연산 예외)
연산의 결과 타입도 인수 타입과 같다
*/
func main() {
	var (
		x int32   = 7
		y int32   = 3
		s float32 = 3.14
		t float32 = 5
	)

	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	fmt.Println("x % y =", x%y)

	fmt.Println("s * t =", s*t)
	fmt.Println("s / t =", s/t)
}

/*
+	덧셈		정수, 실수, 복소수, 문자열
-	뺄셈		정수, 실수, 복소수
*	곱셈		정수, 실수, 복소수
/	나눗셈	정수, 실수, 복소수
%	나머지	정수
*/
