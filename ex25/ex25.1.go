package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch) //고루틴 생성
	//square() 함수는 main() 루틴이 아닌 새로운 고루틴에서 동시에 실행
	//square() 함수는 앞서 생성한 채널 인스턴스를 인수로 받아서 먼저 채널에서
	//데이터를 빼오려 시도한다. 현재 채널에는 아무런 데이터가 없기 때문에
	//데이터가 들어올 때까지 대기한다.
	ch <- 9   //채널에 데이터 넣음
	wg.Wait() //작업이 완료되길 기다림
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second) //1초 대기
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}

/*
채널이란 고루틴끼리 메시지를 전달할 수 있는 메시지 큐이다
메시지 큐에 메시지들은 차례대로 쌓이게 되고 메시지를 읽을 때는
맨 처음 온 메시지부터 차례대로 읽게 된다.

채널 인스턴스를 만들어야 함
var messages chan string = make(chan string)
messages <- "This is a message" 데이터 넣기
var msg string = <- messages 데이터 빼기
*/
