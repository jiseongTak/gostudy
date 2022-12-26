package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

// 포인터를 사용하지 않는 예
func main() {
	var data Data

	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
	//Data 구조체는 1608 바이트
	//ChangeData()함수를 한 번 호출할 때마다 1608바이트가 복사된다.
	//포인터를 이용해서 이를 해결해줄 수 있다.
}

/*
변수 대입이나 함수 인수 전달은 항상 값을 복사하기 때문에 많은 메모리 공간을
사용하는 문제와 큰 메모리 공간을 복사할 때 발생하는 성능 문제를 안고 있다.
또한 다른 공간으로 복사되기 때문에 변경 사항이 적용되지도 않는다.
*/
