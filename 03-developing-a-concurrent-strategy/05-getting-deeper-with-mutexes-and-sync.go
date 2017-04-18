package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeStruct struct {
	totalChanges int
	currentTime  time.Time
	rwLock       sync.RWMutex
}

var TimeElement TimeStruct

func updateTime() {
	TimeElement.rwLock.Lock()
	defer TimeElement.rwLock.Unlock()
	TimeElement.currentTime = time.Now()
	TimeElement.totalChanges++
}

func main() {
	var wg sync.WaitGroup

	TimeElement.totalChanges = 0
	TimeElement.currentTime = time.Now()
	timer := time.NewTicker(1 * time.Second)
	writeTimer := time.NewTicker(10 * time.Second)
	endTimer := make(chan bool)

	wg.Add(1)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println(TimeElement.totalChanges, TimeElement.currentTime.String())
			case <-writeTimer.C:
				updateTime()
			case <-endTimer:
				timer.Stop()
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println(TimeElement.currentTime.String())
}
