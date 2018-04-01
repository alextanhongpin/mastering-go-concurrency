// This program demonstrates how to pause all running goroutines when an error
// occured in one of them. Consider this as a global circuit-breaker

package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numWorkers := 5
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	pause := make(chan interface{}, numWorkers)

	counter := 0
	start := time.Now()
	for i := range worker(ctx, generator(ctx, 100), numWorkers, pause) {
		counter++
		log.Println(i)
	}

	log.Println(time.Since(start))
	log.Println("done", counter)
}

func generator(ctx context.Context, limit int) <-chan interface{} {
	outStream := make(chan interface{})

	go func() {
		defer close(outStream)
		for i := 0; i < limit; i++ {
			select {
			// The sequence of context done, does it matters?
			case outStream <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return outStream
}

func worker(ctx context.Context, inStream <-chan interface{}, numWorkers int, pause chan interface{}) <-chan interface{} {

	outStream := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	failIndex := rand.Intn(100)
	multiplex := func(index int, in <-chan interface{}) {
		for i := range in {
			// Assuming each operation takes roughly 10 ms
			if i == failIndex {
				log.Println("error occurred")
				// Fill up all the workers
				for j := 0; j < numWorkers; j++ {
					pause <- j
				}
			}
			time.Sleep(10 * time.Millisecond)
			select {
			case <-pause:
				log.Println("pausing worker", index)
				time.Sleep(1 * time.Second)
				log.Println("done pausing worker", index)
				outStream <- i
			case <-ctx.Done():
				return
			case outStream <- i:
			}
		}
		wg.Done()
	}

	for i := 0; i < numWorkers; i++ {
		go multiplex(i, inStream)
	}

	go func() {
		wg.Wait()
		close(outStream)
	}()

	return outStream
}
