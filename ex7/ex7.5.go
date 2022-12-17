package main

import "fmt"

// Divide2 출력값에 이름 지정
// return에 값을 안적어도 된다.
func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide2(9, 3)
	fmt.Println(c, success)
	d, success := Divide2(9, 0)
	fmt.Println(d, success)
}

//반환할 변수에 이름을 지정할 경우 모든 반환 변수에 이름을 지정해야 한다
//모두 지정하거나, 모두 지정하지 않거나
