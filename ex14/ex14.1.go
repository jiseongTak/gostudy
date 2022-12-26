package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}

/*
포인터는 메모리 주소를 값으로 갖는 타입
int 타입 변수 a가 있을 때 a는 메모리에 저장되어 있고
속성으로 메모리 주소를 가지고 있다. 변수 a의 주소가 0x0100번지라고
했을 때 메모리 주솟값 또한 숫자값이기 때문에 다른 변수의 값으로 사용될 수 있다.
이렇게 메모리 주솟값을 변숫값으로 가질 수 있는 변수를 포인터 변수라 한다.
포인터를 이용하면 여러 포인터 변수가 하나의 메모리 공간을 가리킬 수도 있다.
포인터가 가리키고 있는 메모리 공간의 값을 읽을 수도 변경할 수도 있다.
var p *int
float64 타입을 가리키면 *float64, User 구조체를 가리키면 *User
var a int
var p *int
p = &a a의 메모리 주소를 포인터 변수 p에 대입한다
p를 이용해서 변수 a의 값을 변경할 수 있다. 포인터 변수 앞에 *를 붙이면
그 포인터 변수가 가리키는 메모리 공간에 접근할 수 있다.
*p = 20
p가 가리키는 메모리 공간의 값을 20으로 변경, p가 변수 a의 메모리 공간을 가리키기 때문에
a값이 20으로 변경된다.
*/
