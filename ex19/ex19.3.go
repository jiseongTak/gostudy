package main

import "fmt"

type account2 struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account2) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account2) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account2) withdrawReturnValue(amount int) account2 {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account2 = &account2{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account2 = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}

/*
언제 포인터 메서드를 만들고 값 타입 메서드를 만들어야 할까
포인터 메서드는 메서드 내부에서 리시버의 값을 변경시킬 수 있다.
하지만 값 타입 메서드는 호출하는 쪽과 메서드 내부의 값은 별도 인스턴스로 독립되기
때문에 메서드 내부에서 리시버의 값을 변경시킬 수 없다.
포인터 메서드는 인스턴스 중심이고 값 타입 메서드는 값 중심이 된다.
B.4절, B.5절
*/
