package main

import "fmt"

func getMyAge2() int {
	return 22
}

func main() {
	//초기문 사용
	switch age := getMyAge2(); true { //true 생략가능
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenage")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
