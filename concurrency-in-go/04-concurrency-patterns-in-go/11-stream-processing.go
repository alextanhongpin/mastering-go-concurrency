package main

import "log"

func main() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		o := add(multiply(v, 2), 3)
		log.Println(o)
	}
}
