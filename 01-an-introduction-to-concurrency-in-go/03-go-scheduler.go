package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Job struct {
	i, max int
	text   string
}

func outputText(j *Job) {
	fileName := j.text + ".text"
	fileContents := ""
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if err != nil {
		panic("Something went awry")
	}
}
func main() {

	hello := &Job{0, 3, "Hello"}
	world := &Job{0, 5, "World"}

	go outputText(hello)
	go outputText(world)
}
