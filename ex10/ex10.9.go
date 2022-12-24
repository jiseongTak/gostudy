package main

import "fmt"

func main() {
	a := 3

	//Go 언어에서는 break를 사용하든 사용하지 않든 상관없이
	//case 하나를 실행 후 switch문을 빠져나간다
	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 2")
	}
}
