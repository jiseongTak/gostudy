package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	//Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환
	//런타임 에러 발생
	c := reader.(Closer)
	c.Close()
}

func main() {
	//다른 인터페이스로 타입 변환
	file := &File{}
	ReadFile(file)
}

/*
reader 인터페이스 변수가 *File 타입을 가리키고 있고 *File 타입은 Close() 메서드를
포함하고 있지 않기 때문에 Closer 인터페이스로 사용할 수 없다.

타입 변환이 아예 불가능한 타입이라면 컴파일 타임 에러가 발생
문법적으로 문제 없지만, 실행 도중 타입 변환에 실패하는 경우에는 런 타임 에러 발생

타입 변환 성공 여부 반환(런 타임 에러 방지)
var a Interface
t, ok := a.(ConcreteType)

func ReadFile(reader Reader) {
	c, ok := reader.(Closer)
	if ok {
		c.Close()
	}
}

if c, ok := reader.(Closer); ok {
	...
}
*/
