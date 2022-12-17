package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, histroy int) {
	total := math + eng + histroy
	avg := total / 3
	fmt.Print(name, "님 평균 점수는 ", avg, "점 입니다\n")
}

func main() {
	PrintAvgScore("김일등", 80, 74, 95)
	PrintAvgScore("송이등", 88, 92, 53)
	PrintAvgScore("박삼등", 78, 73, 78)
}
