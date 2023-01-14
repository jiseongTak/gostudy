package main

import (
	"fmt"
	"sync"
	"time"
)

func square4(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
		//ch 채널을 먼저 시도하기 때문에 ch 채널에서
		//데이터를 읽을 수 있으면 계속 읽는다.
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square4(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}

/*
채널에서 데이터가 들어오기를 대기하는 상황에서 만약 데이터가 들어오지 않으면
다른 작업을 하거나, 아니면 여러 채널을 동시에 대기하고 싶을 때 select 문 사용
select {
case n := <-ch1:
	...		//ch1 채널에서 데이터를 빼낼 수 있을 때 실행
case n := <-ch2:
	...		//ch2 채널에서 데이터를 빼낼 수 있을 때 실행
case ...
}
여러 채널을 동시에 기다릴 수 있다. 만약 어떤 채널이라도 하나의 채널에서
데이터를 읽어오면 해당 구문을 실행하고 select 문이 종료
하나의 case만 처리되면 종료되기 때문에 반복해서 데이터를 처리하고 싶다면
for문과 함께 사용
*/
