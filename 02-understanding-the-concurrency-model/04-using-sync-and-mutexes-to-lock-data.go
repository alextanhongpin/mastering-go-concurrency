package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex = &sync.Mutex{}

func test1() {
	current := 0
	iterations := 100

	wg := new(sync.WaitGroup)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			current++
			fmt.Println(current)
			wg.Done()
		}()
		wg.Wait()
	}
}

func test2() {
	runtime.GOMAXPROCS(2)
	current := 0
	iterations := 100

	wg := new(sync.WaitGroup)
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			mutex.Lock()
			current++
			mutex.Unlock()
			fmt.Println(current)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	// test1()
	test2()
}
