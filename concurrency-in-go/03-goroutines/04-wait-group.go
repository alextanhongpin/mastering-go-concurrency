package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine is sleeping")
		time.Sleep(1 * time.Second)
		fmt.Println("1st goroutine is done")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine is sleeping")
		time.Sleep(2 * time.Second)
		fmt.Println("2nd goroutine is done")
	}()

	wg.Wait()
	fmt.Println("all goroutines completed")

}
