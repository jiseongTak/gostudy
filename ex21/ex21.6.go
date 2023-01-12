package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}

/*
함수 리터럴 외부 변수를 내부 상태로 가져오는 것을 캡쳐라고 한다.
캡쳐는 값 복사가 아닌 참조 형태로 가져오게 되니 주의해야 한다.

i 변수를 캡쳐할 때 캡쳐하는 순간의 i값이 복사되는게 아니라 i 변수가 참조로
캡쳐되기 때문에 3,3,3이 호출
for문이 진행될 때마다 i값이 증가하고 최종적으로 i=3에서 for문이 종료되기 때문에
함수 리터럴이 호출되는 시점의 캡쳐된 i값은 3이 된다.

v 변수를 선언해서 i값을 복사한다. 함수 리터럴에서 v 변수를 캡쳐한다.
v 변수는 for문 내부에서 선언됐기 때문에 매 루프마다 새로운 v 변수가 생성된다.
따라서 f 슬라이스의 각 함수 리터럴 요소는 서로 다른 v 변수를 캡쳐하게 된다.

참조로 변수를 가져온다 ->
변수의 주소를 포인터 값으로 복사한다고 보면 된다.함수 리터럴 외부의 변수 i를 캡쳐할 때
변수 i의 주솟값을 포인터 형태로 함수 리터럴 내부 상태로 가져와 나중에 캡쳐된 내부 상태를 사용할 때
메모리 주솟값을 통해 외부 변수 i에 접근하게 된다.
*/
