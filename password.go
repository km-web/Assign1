package main

import (
	"fmt"
	"strings"
	"unicode"
)

func verifyPassword(password string) error {

	const minPassLength = 8
	const maxPassLength = 12
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {

		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):

			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}

	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

func main() {
	password := "kathir"
	err := verifyPassword(password)
	fmt.Println(password, " ", err)
}
