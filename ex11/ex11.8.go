package main

import "fmt"

func main() {
	//레이블을 이용
	//break 할 때 앞서 정의한 레이블을 적어주면 그 레이블에서
	//가장 먼저 속한 for문까지 종료
	a := 1
	b := 1

OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
