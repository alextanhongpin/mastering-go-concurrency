package main

import "log"

func main() {
	// Write a program that finds a prime number from 50.000 and return them
	// through channels. Spawn it for multiple channels

	// numCPUs := runtime.NumCPU()

	onlyEven := func(
		done <-chan interface{},
		takeStream <-chan int,
	) <-chan int {
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for v := range takeStream {
				select {
				case <-done:
					return
				case valueStream <- v * 2:
				}
			}
		}()
		return valueStream
	}

	gen := func(
		done <-chan interface{},
		limit int,
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for i := 0; i < limit; i++ {
				select {
				case <-done:
					return
				case valueStream <- i:
				}
			}
		}()
		return valueStream
	}

	toInt := func(
		done <-chan interface{},
		takeStream <-chan interface{},
	) <-chan int {
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for v := range takeStream {
				select {
				case <-done:
					return
				case valueStream <- v.(int):
				}
			}
		}()
		return valueStream
	}

	done := make(chan interface{})

	for v := range onlyEven(done, toInt(done, gen(done, 100))) {
		log.Printf("%#v\n", v)
	}
}
