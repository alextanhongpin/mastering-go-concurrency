package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/
// Solution 3: Internet Cafe
func main() {
	rand.Seed(time.Now().UnixNano())

	numTourists := 25
	maxOnline := 8
	var wg sync.WaitGroup
	wg.Add(numTourists)

	done := make(chan interface{}, maxOnline)
	defer close(done)

	kv := make(map[int]int)
	var mu sync.Mutex
	online := func(done chan interface{}, wg *sync.WaitGroup, i int) {
		mu.Lock()
		kv[i] = i
		mu.Unlock()
		if len(kv) == maxOnline {
			for j := 0; j < numTourists; j++ {
				mu.Lock()
				v, ok := kv[j+1]
				mu.Unlock()
				if ok && v != i {
					log.Printf("Tourist %d waiting for turn.\n", j)
				}
			}
		}

		done <- i
		log.Printf("Tourist %d is online.\n", i)
		duration := 5 + rand.Intn(10)
		time.Sleep(time.Duration(duration) * time.Second)
		log.Printf("Tourist %d is done, having spent %d seconds online.\n", i, duration)
		<-done
		wg.Done()

	}

	for i := 0; i < numTourists; i++ {
		go online(done, &wg, i+1)
	}

	wg.Wait()

	log.Println("The place is empty, let's close up and go to the beach!")
}
