package main

import (
	"gostudy/ex20/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	//sender := &koreaPost.PostSender{}
	//SendBook("어린왕자", sender) //타입이 다름
	//SendBook("그리스인 조르바", sender)
}
