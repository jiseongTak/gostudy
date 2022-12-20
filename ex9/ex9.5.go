package main

import "fmt"

/*
음식값이 오만 원이 넘고 친구 중에 부자가 있다면 신발끈을 묶는다.
부자가 없다면 돈을 나눠 낸다. 음식값이 3만 원 이상 오만 원 이하이고
같이 간 친구 수가 3명이 넘으면 신발끈을 묶는다.
3명 이하면 돈을 나눠 낸다. 3만 원 미만이면 내가 낸다.
*/

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈 묶기")
		} else {
			fmt.Println("나눠 낸다")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("신발끈 묶기")
		} else {
			fmt.Println("나눠 낸다")
		}
	} else {
		fmt.Println("내가 낸다")
	}
}

//중첩이 심할 경우 코드를 이해하기 힘들다.
//되도록 3중첩 이상은 하지 않도록 권장
