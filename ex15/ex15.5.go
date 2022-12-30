package main

import "fmt"

func main() {
	str := "Hello World"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))
}

/*
Hello World 문자열은 문자들의 집합이고 각 문자들은 UTF-8 코드로 값을 가진다.
문자열은 코드값의 배열인 rune 배열로 나타낼 수 있다.
*/
