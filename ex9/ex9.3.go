package main

import "fmt"

func main() {
	age := 33

	if age >= 10 && age <= 15 {
		fmt.Println("You ar young")
	} else if age > 30 || age < 20 {
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best age of your life")
	}
}