package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	/*	house.Address = "서울시 관악구"
		house.Size = 28
		house.Price = 10.2
		house.Category = "아파트"*/

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입:", house.Category)
	/*
		초기화 안하면 기본값으로
		string 빈 문자열 ""
		int 0
		float64 0.0
	*/

	fmt.Printf("%v\n", house)

	//필드 초기화
	var house2 House = House{"서울시 강동구",
		28,
		9.80,
		"아파트", //여러 줄로 초기화할 때는 제일 마지막 값 뒤에 쉼표
	}
	fmt.Println(house2)

	//일부 필드값만 초기화
	var house3 House = House{Size: 28, Category: "아파트"}
	fmt.Println(house3)
}

/*
여러 필드를 묶어서 하나의 구조체를 만든다
다른 타입의 값들을 변수 하나로 묶어준다

type 타입명 struct {
	필드명 타입
	...
	필드명 타입
}
*/
