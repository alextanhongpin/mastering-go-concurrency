package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	N_ITEMS        = 100
	MAX_BATCH_SIZE = 10
	MAX_DURATION   = 20
)

type Batch struct {
	MaxBatchSize int
	Queue        []int
	Duration     time.Duration
	Out          chan interface{}
}

func (b *Batch) Process() {
	if !b.HasItems() {
		return
	}
	// Process in batch
	for _, v := range b.Queue {
		b.Out <- v
	}
	// Empty queue
	b.Queue = []int{}
}

func (b *Batch) IsFull() bool {
	return len(b.Queue) == b.MaxBatchSize
}

func (b *Batch) HasItems() bool {
	return len(b.Queue) > 0
}
func (b *Batch) Size() int {
	return len(b.Queue)
}

func (b *Batch) Close() {
	close(b.Out)
}

func (b *Batch) Run(done, in <-chan interface{}) {
	go func() {
		defer b.Close()
		for {
			select {
			case <-done:
				log.Println("done, processing remaining", b.Size())
				b.Process()
			case <-time.After(b.Duration):
				log.Println("timeout exceeded, processing", b.Size())
				b.Process()
			case v, ok := <-in:
				if !ok {
					log.Println("channel closed, processing remaining", b.Size())
					b.Process()
					return
				}
				b.Queue = append(b.Queue, v.(int))
				if b.IsFull() {
					log.Println("threshold exceeded, processing", b.Size())
					b.Process()
					// We can add a delay here to allow some buffer time for processing
					log.Println("sleep 1s before next batch")
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()
}

func NewBatch(maxBatchSize, duration int) *Batch {
	return &Batch{
		MaxBatchSize: maxBatchSize,
		Queue:        []int{},
		Duration:     time.Duration(duration) * time.Millisecond,
		Out:          make(chan interface{}),
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	generator := func(done <-chan interface{}, limit int) <-chan interface{} {
		outStream := make(chan interface{})

		go func() {
			defer close(outStream)
			for i := 0; i < limit; i++ {
				time.Sleep(time.Duration(rand.Intn(10)+10) * time.Millisecond)
				outStream <- i
			}
		}()

		return outStream
	}

	done := make(chan interface{})
	defer close(done)

	// batch := []int{}
	// out := make(chan interface{})

	in := generator(done, N_ITEMS)

	batch := NewBatch(MAX_BATCH_SIZE, MAX_DURATION)
	batch.Run(done, in)

	var wg sync.WaitGroup
	wg.Add(N_ITEMS)
	go func() {
		var sum int
		for o := range batch.Out {
			defer wg.Done()
			log.Println("got", o)
			sum++
		}
		log.Printf("processed %d items\n", sum)
	}()
	wg.Wait()
	log.Println("done")
}
