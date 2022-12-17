package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)
	//서식에 맞춰서 두 숫자 사이에 공란이 와야 한다
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
