package main

import "fmt"

type User2 struct {
	Name string
	ID   string
	Age  int
}

type VIPUser2 struct {
	User2    //필드명 생략
	VIPLevel int
	Price    int
}

func main() {
	user := User2{"송하나", "hana", 23}
	vip := VIPUser2{
		User2{"화랑", "hwarang", 48},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n",
		vip.Name, //. 하나로 접근가능
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)

}
