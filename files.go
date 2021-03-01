package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Create("batch11.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("kathir batch11 trainee")
	file.Close()
	stream, err := ioutil.ReadFile("batch11.txt")
	if err != nil {
		log.Fatal(err)
	}
	output := string(stream)
	fmt.Println(output)
}
