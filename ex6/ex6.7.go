package main

import "fmt"

// 실수값을 정확히 표현할 수 없기 때문에 오차가 생길 수밖에 업다.
const epsilon = 0.000001 //매우 작은 값

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f = %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}

//위의 방법은 좋은 방법이 아니다
//문제는 도대체 얼만큼의 오차가 무시할만큼 작은 오차냐는 것이다
//float64의 경우 10e-308 ~ 10e308까지 매우 큰 값의 범위
//가장 간편하고 좋은 방법은 1비트 차이만큼 비교하는 것
/*
0.29999 9982 1186 0656 7382 8125
0	0111 1101	00110011001100110011001
0.3000 0001 1920 9289 5507 8125
0	0111 1101	00110011001100110011010

0.3은 float32 타입에서 위 두 값 중 하나로 표현해야 한다
가장 마지막 비트만 차이난다
0.3을 표현할 수 있는 값의 실수 타입 범위에서는 가장 작은 차이이다
가장 마지막 비트가 1비트만큼 차이 나는지 어떻게? ->
math 패키지 Nextafter() 함수
*/
