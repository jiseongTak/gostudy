package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	//컨텍스트에 값을 추가
	ctx := context.WithValue(context.Background(), "number", 9)

	go square6(&wg, ctx)

	wg.Wait()
}

func square6(wg *sync.WaitGroup, ctx context.Context) {
	if v := ctx.Value("number"); v != nil { //컨텍스트에서 값을 읽기
		//Value() 메서드의 반환 타입은 빈 인터페이스, int 타입으로 변환
		n := v.(int)
		fmt.Printf("Square: %d\n", n*n)
	}
	wg.Done()
}
