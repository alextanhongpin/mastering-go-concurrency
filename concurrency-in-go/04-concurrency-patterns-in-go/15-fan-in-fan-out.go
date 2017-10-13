package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	isPrime := func(value int) bool {
		for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
			if value%i == 0 {
				return false
			}
		}
		return value > 1
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	primeFinder := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	rand := func() interface{} { return rand.Intn(500000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFunc(done, rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Println(prime)
	}
}
