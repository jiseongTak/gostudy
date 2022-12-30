package main

import "fmt"

func main() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}

	//[]rune 으로 변환한 다음 순회하면 한 글자씩 순회할 수 있지만
	//[]rune 으로 변환되는 과정에서 별도의 배열을 할당하므로 불필요한 메모리를 사용
	//range 키워드를 사용해 순회하면 이를 방지할 수 있다.
}
