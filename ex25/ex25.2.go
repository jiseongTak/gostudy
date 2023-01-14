package main

import "fmt"

// 채널에서 데이터를 가져가지 않아서 프로그램이 멈추는 경우
func main() {
	ch := make(chan int)

	ch <- 9
	fmt.Println("Never print")
}

/*
일반적으로 채널을 생성하면 크기가 0인 채널(unbuffered channel)이 만들어짐
크기가 0이라는 뜻은 채널에 들어온 데이터를 담아둘 곳이 없다는 뜻

내부에 데이터를 보관할 수 있는 메모리 영역을 버퍼라고 부른다.
그래서 보관함을 가지고 있는 채널을 버퍼를 가진 채널이라고 한다.
var chan string messages = make(chan string, 2)
버퍼가 다 차면 버퍼가 없을 떄와 마찬가지로 보관함에 빈자리가 생길 때 까지 대기
그래서 데이터를 제때 빼가지 않으면, 버퍼가 없을 때처럼 고루틴이 멈추게 된다.
*/
