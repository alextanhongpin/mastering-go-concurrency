package main

import (
	"fmt"
)

type intInterface struct{}
type strInterface struct{}

func (number intInterface) Add(a, b int) int {
	return a + b
}
func (text strInterface) Add(a, b string) string {
	return a + b
}

func main() {

	number := intInterface{}
	fmt.Println(number.Add(1, 2))

	text := strInterface{}
	fmt.Println(text.Add("hello", "world"))
}
