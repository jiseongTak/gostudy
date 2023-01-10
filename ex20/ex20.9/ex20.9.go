package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	//런타임 에러 발생: *Student 타입은 Stringer 인터페이스로 쓰일 수 있지만
	//stringer 값이 *Student 타입이 아니기 때문에 에러가 발생
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	actor := &Actor{}
	ConvertType(actor)
	//panic: interface conversion: main.Stringer is *main.Actor, not *main.Student
}

/*
*Student 타입과 *Actor 타입은 모두 String() 메서드를 가지고 있기 때문에
Stringer 인터페이스로 사용할 수 있습니다. 하지만 ConvertType() 함수 인수인
stringer 인터페이스 변수는 *Actor 타입 인스턴스를 가리키고 있기 때문에
*Student 타입으로 변환을 시도하면 런 타임 에러가 발생
*/
