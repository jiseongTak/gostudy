package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(30))
	var b int = 20
	fmt.Println(myInt(b).add(50))
}

/*
모든 로컬 타입이 리시버 타입으로 가능하기 때문에 별칭 타입도 리시버가 될 수 있고
메서드를 가질 수 있다. int와 같은 내장 타입들도 별칭 타입을 활용해서 메서드를 가질 수가 있다.

*/
