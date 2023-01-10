package main

import "fmt"

func main() {
	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b > 9 {
				break //안쪽 for문 종료
			}
		}
		fmt.Println("---------")
		b = 1
		dan++
		if dan > 9 {
			break //바깥 for문 종료
		}
	}
	fmt.Println("for문이 종료됐습니다.")

}
