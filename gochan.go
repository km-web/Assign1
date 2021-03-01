package main

import (
	"fmt"
)

func auto(bikes chan string) {
	fmt.Println("pulsar" + <-bikes)
}

func main() {
	fmt.Println("top model")
	bikes := make(chan string)

	go auto(bikes)
	bikes <- "150"
	fmt.Println("end")

}
