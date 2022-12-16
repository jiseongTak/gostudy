package main

import "fmt"

func main() {
	a := 3 //int - 64bit -> int64
	var b float64 = 3.5

	var c int = int(b) //3.5 -> 3
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	var h int = int(b * 3)
	var i int = int(b) * 3
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(h, i)
}

//타입이 반드시 같아야 함
//타입변환을 해야한다
//실수 -> 정수 소수점 이하 숫자 없어짐
