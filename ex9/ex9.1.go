package main

import "fmt"

func main() {
	light := "red"

	if light == "green" {
		fmt.Printf("길을 건넌다")
	} else {
		fmt.Println("대기한다")
	}
}
