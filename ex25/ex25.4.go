package main

import (
	"fmt"
	"sync"
	"time"
)

func square3(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //채널이 닫히면 종료
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	fmt.Println(&wg)
	go square3(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) //채널 닫음
	wg.Wait()
}
