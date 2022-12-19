package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

const (
	Red   int = iota //0
	Blue  int = iota //1
	Green int = iota //2
)

// 첫 번째 값과 똑같은 규칙이 계속 적용된다면 타입과 iota 생략 가능
const (
	C1 uint = iota + 1 // 1 = 0 +1
	C2                 // 2 = 1 + 1
	C3                 // 3 = 2 + 1
)

const (
	BitFlag1 uint = 1 << iota // 1 = 1 << 0
	BitFlag2                  // 2 = 1 << 1
	BitFlag3                  // 4 = 1 << 2
	BitFlag4                  // 8 = 1 << 3
)

func main() {
	PrintAnimal(Cow)
	PrintAnimal(Pig)
	PrintAnimal(7)
	PrintAnimal(0)

	fmt.Println(Red, Blue, Green)
	fmt.Println(C1, C2, C3)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3, BitFlag4)
}

//코드값을 통해서 숫자에 의미를 부여할 때 사용
