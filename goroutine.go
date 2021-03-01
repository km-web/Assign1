package main

import (
	"fmt"
	"time"
)

func call() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Printf("%d \n ", i)
	}
}

func main() {

	fmt.Println("kathir")
	go call()
	fmt.Println("kumar")
	go call()

	time.Sleep(time.Second)

}
