package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}

/*
조건문이 true 이면 무한루프가 된다.
for true {
	코드 블록
}
true 생략 가능
for {
	코드블록
}
*/
