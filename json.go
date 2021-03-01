package main

import (
	"encoding/json"
	"fmt"
)

type batch11 struct {
	Age  int
	Name string
}

func main() {

	Age, _ := json.Marshal(26)
	fmt.Println(string(Age))

	Name, _ := json.Marshal("kathir")
	fmt.Println(string(Name))

}
