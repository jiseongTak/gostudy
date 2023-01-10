package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")

			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료됐습니다.")
}

/*
continue는 이후 코드 블록을 수행하지 않고 곧바로 후처리를 하고
조건문 검사부터 다시하고, break는 for문에서 곧바로 빠져 나온다.
*/
