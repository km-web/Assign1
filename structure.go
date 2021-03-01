package main

import "fmt"

type user struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	var u user
	u.Name = "kathir"
	u.Age = 26
	u.Gender = "Male"
	fmt.Println(u)
}
