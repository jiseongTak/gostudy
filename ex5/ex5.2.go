package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%015d, %05d\n", a, b)
	fmt.Printf("%-5d, %-5d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-5d\n", c, c)
}

/*
최소 출력 너비 지정: 서식 문자의 %와 타입을 나타내는 문자 사이에
숫자를 넣어서 너비를 지정 %5d는 최소 5칸을 사용해서 정숫값을 출력
공란 채우기: 너비 앞에 0을 붙이면 빈자리를 0으로 채움
왼쪽정렬하기, 마이너스 -를 붙이면 왼쪽을 기준선 삼아 출력
최소 너비보다 긴 값은 모두 지정한 최소 너비가 무시되어 출력
*/
