package main

import (
	"fmt"
	"time"
)

func thinkAboutKeys(bc chan int) {
	i := 0
	max := 10

	for {
		if i >= max {
			bc <- 1
		}
		fmt.Println("Still thinking")
		time.Sleep(1 * time.Second)
		i++
	}
}

func main() {
	fmt.Println("Where did I leave my keys?")

	blockChannel := make(chan int)
	go thinkAboutKeys(blockChannel)
	<-blockChannel

	fmt.Println("OK I found them")
}
