package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker //기본값은 nil
	att.Attack()     //att가 nil이기 때문에 런 타임 에러가 발생
}

/*
인터페이스 기본값은 nil
기본값을 nil로 갖는 타입은 포인터, 인터페이스, 함수 타입, 슬라이스, 맵, 채널 등
*/
