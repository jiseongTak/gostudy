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
	go square5(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

func square5(wg *sync.WaitGroup, ch chan int) {
	//일정 시간 간격 주기로 신호를 보내주는 채널을 생성해서 반환
	tick := time.Tick(time.Second) //1초 간격 시그널
	//현재 시간 이후로 일정 시간 경과 후에 신호를 보내주는 채널을 생성해서 반환
	terminate := time.After(time.Second * 10) //10초 이후 시그널

	for {
		select {
		case <-tick:
			fmt.Println("TICK")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}
