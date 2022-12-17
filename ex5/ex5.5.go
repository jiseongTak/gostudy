package main

import "fmt"

func main() {
	var a int
	var b int

	//변수 앞에 &를 붙여서 변수의 메모리 주소를 입력으로 넘겨야함
	n, err := fmt.Scan(&a, &b) //n은 입력한 값 개수, err은 입력 시 발생한 에러
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

//Scan 한번에 여러 값을 입력받을 때는 변수 사이를 공란을 두어 구분
//엔터도 공란으로 인식
/*
	Scan()	 표준 입력에서 값을 입력받음
	Scanf()  표준 입력에서 서식 형태로 값을 입력받음
	Scanln() 표준 입력에서 한 줄을 읽어서 값을 입력받음
*/
