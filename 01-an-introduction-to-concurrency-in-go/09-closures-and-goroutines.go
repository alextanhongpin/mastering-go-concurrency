package main

import (
	"fmt"
	"strings"
)

func shortenString(message string) func() string {
	return func() string {
		messageSlice := strings.Split(message, " ")
		wordLength := len(messageSlice)
		if wordLength < 1 {
			return "Nothing Length"
		} else {
			messageSlice = messageSlice[:(wordLength - 1)]
			message = strings.Join(messageSlice, " ")
			return message
		}
	}
}
func main() {

	myString := shortenString("Welcome to concurrency in go")
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
}
