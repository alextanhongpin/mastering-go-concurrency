package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
)

func writeFile(i int) {
	rwLock.RLock()
	ioutil.WriteFile("test.txt", []byte(strconv.FormatInt(int64(i), 10)), 0x777)

	rwLock.RUnlock()
	writer <- true
}

var writer chan bool
var rwLock sync.RWMutex

func main() {
	writer = make(chan bool)
	for i := 0; i < 10; i++ {
		go writeFile(i)
	}
	<-writer
	fmt.Println("Done")
}
