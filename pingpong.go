package main

import (
	"fmt"
	"time"
)

func ping1(ping <-chan int) {
	for {
		<-ping
		fmt.Println("ping")
		time.Sleep(time.Second)
	}
}
func ping2(pong chan<- int) {
	for {
		fmt.Println("ping")
		time.Sleep(time.Second)
		pong <- 1
	}
}
func pong1(ping chan<- int) {
	for {
		fmt.Println("pong")
		time.Sleep(time.Second)
		ping <- 1
	}
}
func pong2(pong <-chan int) {
	for {
		<-pong
		fmt.Println("pong")
		time.Sleep(time.Second)

	}
}
func main() {
	ping := make(chan int)
	pong := make(chan int)

	go ping1(ping)
	go ping2(pong)
	go pong1(ping)
	go pong2(pong)

	ping <- 1

	for {

		time.Sleep(time.Second)
	}
}
