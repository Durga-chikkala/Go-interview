// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func Even(n int, evenCh, oddCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= n; i += 2 {
		<-evenCh
		fmt.Println(i)

		if i+1 <= n {
			oddCh <- true
		}

	}
}

func Odd(n int, evenCh, oddCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-oddCh
		fmt.Println(i)

		if i+1 <= n {
			evenCh <- true
		}
	}
}

func main() {
	n := 11
	wg := &sync.WaitGroup{}
	wg.Add(2)

	evenCh, oddCh := make(chan bool), make(chan bool)

	go Odd(n, evenCh, oddCh, wg)
	go Even(n, evenCh, oddCh, wg)

	oddCh <- true

	wg.Wait()

}
