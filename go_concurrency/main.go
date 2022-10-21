package main

import (
	"fmt"
)

func someFunx(num int, channel chan int) {
	channel <- num
}
func main() {

	myChannel := make(chan string)
	go func() {
		myChannel <- "ok"
	}()

	msg := <-myChannel

	intChan := make(chan int)

	fmt.Println(msg)
	go someFunx(5, intChan)
	go someFunx(4, intChan)
	go someFunx(2, intChan)
	go someFunx(1, intChan)

	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

	fmt.Println("hi")
}

/*
Fork join model
*/
