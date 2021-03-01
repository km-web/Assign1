package main

import (
	"fmt"
	"math/rand"
)

func random(x <-chan int) {
	for {
		<-x
		fmt.Println(rand.Intn(100))
	}
}
func main() {
	x := make(chan int)
	go random(x)
	x <- 1
}
