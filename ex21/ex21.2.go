package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")
	//defer 는 역순으로 호출된다.

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}

/*
함수가 종료되기 직전에 실행해야 하는 코드가 있을 수 있다.
대표적으로 파일이나 소켓 핸들처럼 OS 내부 자원을 사용하는 경우
파일 작업 이후 반드시 파일 핸들을 반환해야 하기 때문에 함수 종료 전에 처리해야
하는 코드가 있을 때 defer를 사용해 실행할 수 있다.
*/
