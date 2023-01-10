package main

import "fmt"

func main() {
	//모든 for문을 빠져나가는 방법 - 불리언 변수 사용(플래그flag 변수)
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
