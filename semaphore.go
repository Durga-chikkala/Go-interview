// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore is a counting semaphore with a fixed capacity.
type Semaphore struct {
	ch chan struct{}
}

func New(n int) *Semaphore {

	return &Semaphore{ch: make(chan struct{}, n)}

}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}

func main() {
	sem := New(2)

	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			fmt.Printf("Goroutine %d acquired semaphore\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d released semaphore\n", id)
		}(i)
	}

	wg.Wait()

}
