package main

import (
	"bufio" //io를 담당하는 패키지
	"fmt"
	"os" //표준 입출력 등을 가지고 있는 패키지
)

func main() {
	stdin := bufio.NewReader(os.Stdin) //표준 입력을 읽는 객체

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		//표준 입력 스트림 지우기
		stdin.ReadString('\n') //줄바꿈 문자가 나올 때까지 읽는다
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b) //다시 입력받기
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}

/*
stdin.ReadString('\n') 주석 처리 하면
hello 4
expected integer
expected integer
*/
