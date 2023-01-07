package main

import "fmt"

func PrintVal(v interface{}) { //빈 인터페이스를 인수로 받는다.
	//빈 인터페이스이기 때문에 모든 타입을 인수로 쓸 수 있다.
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}
