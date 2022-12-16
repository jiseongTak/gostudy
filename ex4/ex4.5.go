package main

import "fmt"

func main() {
	var a int16 = 3456   //a는 int16타입 - 2바이트 정수
	var b int8 = int8(a) //int16 -> int8

	fmt.Println(a, b) //3456 -128
}

//0000 1101 1000 0000 a int16 = 3456
//--------- 1000 0000 b int8 = -128
