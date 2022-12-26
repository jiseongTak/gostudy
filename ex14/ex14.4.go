package main

import "fmt"

type Data2 struct {
	value int
	data  [200]int
}

func ChangeData2(arg *Data2) { //매개변수로 Data 포인터를 받는다.
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data2

	ChangeData2(&data) //인수로 data의 주소를 넘긴다.
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
	//메모리 주소는 8바이트 숫자값이기 때문에
	//1608바이트의 구조체 전부가 복사되는게 아니라 8바이트만 복사된다.
	/*
		기존방식
		var data Data
		var p *Data = &data
		구조체를 생성해 초기화하는 방식
		var p *Data = &Data{}
	*/
}

/*
인스턴스란 메모리에 할당된 데이터의 실체를 말한다.
예를 들어 다음 코드는 Data 타입값을 저장할 수 있는 메모리 공간을 할당한다.
var data Data 이렇게 할당된 메모리 공간의 실체를 인스턴스라고 부른다.
포인터를 이용해서 인스턴스에 접근할 수 있다.
구조체 포인터를 함수 매개변수로 받는다는 말은 구조체 인스턴스로
입력을 받겠다는 얘기와 같다.

p1 := &Data{} &를 사용하는 초기화
var p2 = new(Data) new()를 사용하는 초기화
new() 내장 함수는 인수로 타입을 받는다. 타입을 메모리에 할당하고
기본값으로 채워 그 주소를 반환한다.
new를 이용해서 내부 필드값을 원하는 값으로 초기화할 수는 없다.
p1 := &Data{3,4}처럼 사용자 초기화가 가능하다.

메모리는 무한한 자원이 아니기 때문에 가비지 컬렉터라는 메모리 청소부 기능을
사용한다. 가비지 컬렉터는 일정 간격으로 메모리에서 쓸모없어진 데이터를 청소한다.
사용되는 데이터인지 아닌지 어떻게 알까?
'아무도 찾지 않는 데이터는 쓸모없는 데이터이다'라고 볼 수 있다.
func TestFunc() {
	u := &User{}	//u 포인터 변수를 선언하고 인스턴스를 생성
	u.Age = 30
	fmt.Println(u)
}	//내부 변수 u는 사라진다. 더불어 인스턴스도 사라진다.
메모리에 User 데이터가 할당됐고 u 포인터 변수가 가리킨다. 이때 이 인스턴스는
u 포인터 변수로 사용되는 인스턴스이기 때문에 지워지면 안된다.
하지만 TestFunc()이 종료되면 함수 내부 변수 u는 사라져 User 인스턴스를 가리키는
포인터 변수가 없게 된다. 이제 User 인스턴스는 쓸모가 없게 되었다. 가비지 컬렉터가
다음번 청소를 할 때 쓸모 없어진 이 User 인스턴스를 지우게 된다.

정리하면,
인스턴스는 메모리에 생성된 데이터의 실체이다.
포인터를 이용해서 인스턴스를 가리키게 할 수 있다.
함수 호출 시 포인터 인수를 통해서 인스턴스를 입력받고 그 값을 변경할 수 있게 된다.
쓸모 없어진 인스턴스는 가비지 컬렉터가 자동으로 지워준다.
*/
