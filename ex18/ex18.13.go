package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// Students []Student 의 별칭 타입 Students
type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	//구조체 슬라이스 정렬
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}

/*
Sort() 함수을 이용하기 위해서는 Len(), Less(), Swap() 세 메서드가 필요
이들 메서드만 구현하면 우리가 정의한 구조체도 정렬을 할 수 있다.

Students(s)는 []Student 타입인 s를 정렬 인터페이스를 포함한 타입인
Students 타입으로 변환하는 구문이다.
[]Student 타입은 정렬에 필요한 Len(), Less(), Swap() 메서드를 가지고 있지
않기 때문에 sort.Sort() 함수의 인수로 사용될 수 없습니다. 그래서 []Student의 별칭
타입을 만들어서 정렬 인터페이스를 포함하도록 했다. 처음부터 Students 타입으로 만들지 않은 이유는
실제 코딩하다 보면 슬라이스 타입으로 만들어서 사용하다가 정렬이 필요한 경우에도 별도 타입을 만들어서
변환하는 경우가 빈번하기 때문이다. 또 이 방식이 sort.Ints() 함수가 내부에서 동작하는 방식이기도 하다.
*/
