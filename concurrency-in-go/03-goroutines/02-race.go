package main

import (
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			log.Println(salutation)
		}(salutation)
	}

	wg.Wait()
}
