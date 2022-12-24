package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 2")
	}
	//case 마지막에 fallthrough 키워드를 사용하면 다음 case까지 같이 실행
}
