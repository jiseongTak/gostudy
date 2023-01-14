package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 작업 취소가 가능한 컨텍스트
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//취소 가능한 컨텍스트를 생성
	//상위 컨텍스트를 인수로 넣으면 그 컨텍스트를 감싼 새로운 컨텍스트를 만든다
	//상위 컨텍스트가 없으면 가장 기본적인 컨텍스트인 context.Background()를 넣어준다
	//첫 번째 반환값은 컨텍스트 객체, 두 번째는 취소함수
	//ctx, cancel := context.WithCancel(context.Background()) //컨텍스트 생성

	//3초 뒤 종료
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go PrintEverySecond(ctx, &wg)
	time.Sleep(5 * time.Second)
	cancel()
	/*
		main() 함수에서 5초 이후에 취소 함수를 호출해 작업 취소를 알린다
		그러면 컨텍스트의 Done() 채널에 시그널을 보내 작업자가 작업 취소를 할 수 있도록 알린다.
	*/

	wg.Wait()
}

func PrintEverySecond(ctx context.Context, wg *sync.WaitGroup) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

/*
컨텍스트는 작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는
작업 명세서 역할을 한다. 새로운 고루틴으로 작업을 시작할 때 일정 시간 동안만
작업을 지시하거나 외부에서 작업을 취소할 때 사용
작업 설정에 관한 데이터를 전달할 수도 있다.
*/
