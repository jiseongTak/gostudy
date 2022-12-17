package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}

/*
	\n	줄바꿈
	\t	탭
	\\	\자체를 출력
	\"	"를 출력
*/
