package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	file, err := os.Create("km.txt")
	file.WriteString("kathir batch11 trainee")
	file.Close()

	if err != nil {
		log.Fatal(err)
	}

}

func ReadFile() {

	file, err := ioutil.ReadFile("km.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}
func main() {
	CreateFile()
	ReadFile()

}
