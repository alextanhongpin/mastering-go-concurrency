package main

import (
	// "flag"
	// "fmt"
	// "math/rand"
	// "os"
	// "runtime"
	// "runtime/pprof"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// const ITERATIONS = 99999
// const STRINGLENGTH = 300

// var profile = flag.String("cpuprofile", "", "output pprof data to file")

// func generateString(length int, seed *rand.Rand, chHater chan string) string {
// 	bytes := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		bytes[i] = byte(rand.Int())
// 	}
// 	chHater <- string(bytes[:length])
// 	return string(bytes[:length])
// }

// func generateChannel() <-chan int {
// 	ch := make(chan int)
// 	return ch
// }

func main() {
	// goodbye := make(chan bool, ITERATIONS)
	// channelThatHatesLetters := make(chan string)

	// runtime.GOMAXPROCS(2)
	// flag.Parse()

	// if *profile != "" {
	// 	flag, err := os.Create(*profile)
	// 	if err != nil {
	// 		fmt.Println("Could not create profile", err)
	// 	}
	// 	pprof.StartCPUProfile(flag)
	// 	defer pprof.StopCPUProfile()
	// }
	// seed := rand.New(rand.NewSource(19))
	// initString := ""

	// for i := 0; i < ITERATIONS; i++ {
	// 	go func() {
	// 		initString = generateString(STRINGLENGTH, seed, channelThatHatesLetters)
	// 		goodbye <- true
	// 	}()
	// }

	// select {
	// case <-channelThatHatesLetters:
	// }
	// <-goodbye
	// fmt.Println(initString)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	// go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
}
