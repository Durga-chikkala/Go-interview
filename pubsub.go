// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

//pubsub in go

func publisher(n int, dataCh chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i <= n; i++ {
		dataCh <- i
		time.Sleep(1 * time.Second)
	}

	close(dataCh)

}

func Subscriber(num int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataCh {
		fmt.Println(data, num)
	}
}

func main() {
	dataCh := make(chan int)
	wg := &sync.WaitGroup{}
	num := 10

	wg.Add(1)
	go publisher(num, dataCh, wg)

	wg.Add(2)
	go Subscriber(1, dataCh, wg)
	go Subscriber(2, dataCh, wg)
	go Subscriber(4, dataCh, wg)
	go Subscriber(3, dataCh, wg)

	wg.Wait()
}
