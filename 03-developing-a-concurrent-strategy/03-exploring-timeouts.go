package main

import (
	"fmt"
	"time"
)

func main() {

	ourCh := make(chan string, 1)

	go func() {

	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Enough is enough")
		close(ourCh)
	}
}
