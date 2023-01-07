package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
	//Sprintf() 함수는 서식에 따라 문자열을 만들어서 반환
	//Printf() 함수가 서식에 따라 문자열을 터미널에 출력하는 함수라면
	//Sprintf() 는 화면에 출력하는 것이 아닌 string 타입으로 반환한다.
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())
}

/*
인터페이스를 이용하면 구체화된 객체가 아닌 인터페이스만 가지고 메서드를
호출할 수 있기 때문에 큰 코드 수정이 없이 필요에 따라 구체화된 객체를 바꿔서
사용할 수 있게 된다.
*/
