package main

import "fmt"

var g int = 10 //패키지 전역 변수

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	//m = s + 20
}

//2의 보수
//0000 0000 0000 1111 15
//1111 1111 1111 0000 비트반전
//1111 1111 1111 0001 +1 -> -15

//0000 0000 0000 1111 +15
//1111 1111 1111 0001 -15
//0000 0000 0000 0000 둘이 더하면 0