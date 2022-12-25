package main

import "fmt"

type User3 struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser3 struct {
	User3 //내장된(embedded) 필드
	Price int
	Level int
}

func main() {
	user := User3{"송하나", "hana", 23, 10}
	vip := VIPUser3{
		User3{"화랑", "hwarang", 48, 10},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d 유저 레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,       //VIPUser3의 Level
		vip.User3.Level, //포함된 구조체명을 쓰고 접근
	)
}
