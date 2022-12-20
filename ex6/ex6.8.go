package main

import (
	"fmt"
	"math"
)

func equal2(a, b float64) bool {
	return math.Nextafter(a, b) == b //값 비교
}
func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f = %v\n", c, a+b, equal2(a+b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal2(a+b, c))
}

//어디까지나 오차를 무시하는 방법이라는 점
//그 오차가 매우 작을 뿐이지, 정확한 계산은 아니다
//금융 프로그램이라면 math/big 패키지에서 제공하는 Float 객체를 사용해야 한다
//math/big의 Float를 이용하면 정밀도를 직접 조정할 수 있기 때문에 정밀도를 높여서
//더 정확한 수치 계산을 할 수 있다
