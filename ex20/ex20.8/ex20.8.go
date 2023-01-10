package main

type Stringer interface {
	String() string
}

type Student struct {
}

func main() {
	//var stringer Stringer
	//stringer.(*Student) //컴파일 타임 에러 발생
}

/*
인터페이스 변수를 구체화된 타입으로 타입 변환하려면 해당 타입이
인터페이스 메서드 집합을 포함하고 있어야 한다. 그렇지 않을 경우 컴파일 타임 에러 발생
*/
