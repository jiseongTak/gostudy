package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}

/*
함수 리터럴을 이용해서 원하는 함수를 그때그때 정의해서 함수 타입 변숫값으로 사용할 수 있다.
또 필요한 외부 변수를 내부 상태로 가져와서 편리하게 사용할 수 있다.
*/
