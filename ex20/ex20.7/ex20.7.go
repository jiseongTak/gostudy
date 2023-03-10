package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student) //인터페이스 변수를 다른 구체화된 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15}

	PrintAge(s)
}
