package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"

	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()
	log.Println(salutation)
}
