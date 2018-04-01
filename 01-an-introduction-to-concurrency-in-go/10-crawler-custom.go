package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://www.mastergoco.com/index1.html",
		"http://www.mastergoco.com/index2.html",
		"http://www.mastergoco.com/index3.html",
		"http://www.mastergoco.com/index4.html",
		"http://www.mastergoco.com/index5.html",
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fullText := ""
	for text := range processLinks(ctx, urls...) {
		fullText += text.(string)
	}
	// log.Println(fullText)
}

func fetch(url string) string {
	resp, _ := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("No HTML body")
		return ""
	}
	log.Println("success:", url)

	return string(body)
}

func processLinks(ctx context.Context, urls ...string) <-chan interface{} {
	outStream := make(chan interface{})

	go func() {
		defer close(outStream)
		for _, url := range urls {
			select {
			case <-ctx.Done():
				return
			case outStream <- fetch(url):
			}
		}
	}()

	return outStream
}
