package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

// VIPUser 다른 구조체를 포함
type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vipUser := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n",
		vipUser.UserInfo.Name,
		vipUser.UserInfo.ID,
		vipUser.UserInfo.Age,
		vipUser.VIPLevel,
		vipUser.Price,
	)
	fmt.Println(user, vipUser)
}
