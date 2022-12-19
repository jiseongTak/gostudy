package main

import "fmt"

func main() {
	var t = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	} //i: 인덱스, v: 원솟값

	for v := range t { //인덱스
		fmt.Println(v)
	}

	for _, v := range t { //빈칸 지시자 이용
		fmt.Println(v)
	}

	//배열은 연속된 메모리
	//요소위치 = 배열 시작 주소 + (인덱스 * 타입크기)
	//인덱스만 알면 빠르게 찾을 수 있다
	//Random Access할 때 배열이 가장 빠르다
}
