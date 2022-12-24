package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}

/*
Violet 같은 새로운 색깔을 추가하면 colorToString() 함수도 수정해줘야 한다
이런 경우를 커플링 됐다고 하거나 결합되어 있다고 말함
열거값이 수정될 때 연관된 모든 switch case문도 수정해야 한다
하나의 열거값에 연관된 switch case는 최대한 줄이는게 좋다
*/
