package main

import (
	"fmt"
	"sync"
	"time"
)

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //2.데이터를 계속 기다림
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	//ch 채널 인스턴스로부터 데이터가 들어오길 기다렸다가 데이터가 들어오면
	//데이터를 빼내서 n 변수에 값을 복사하고 for 본문을 실행
	wg.Done() //3.실행되지 않음
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square2(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 //1.데이터를 넣음
	}
	wg.Wait() //4.작업 완료를 기다림
}

/*
wg.Wait() 메서드로 작업이 완료되기를 기다린다. 하지만 for range 구문은
채널에 데이터가 들어고기를 계속 기다리기 때문에 절대 4 가 실행되지 않고 모든
고루틴이 멈추게 되어 deadlock이 표시된다.
채널을 다 사용하면 close(ch)를 호출해 채널을 닫고 채널이 닫혔음을 알려줘야 한다.
*/
