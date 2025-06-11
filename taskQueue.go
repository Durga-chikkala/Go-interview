
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

// TaskQueue Pattern
type TaskQueue struct {
	queue []chan func()
	wg    *sync.WaitGroup
}

func NewTaskQueue(queues, queueSize int) *TaskQueue {
	tq := &TaskQueue{wg: &sync.WaitGroup{}}

	if queues == 0 {
		queues = 50
	}

	if queueSize == 0 {
		queueSize = 10
	}

	tq.queue = make([]chan func(), queues)

	for i := range tq.queue {
		tq.queue[i] = make(chan func(), queueSize)
	}

	return tq
}

func (t *TaskQueue) Add(key string, f func()) {

	t.queue[getInt(key)%len(t.queue)] <- f
}

func (t *TaskQueue) Close() {
	for _, q := range t.queue {
		close(q)
	}
}

func (t *TaskQueue) Run() {
	for _, queue := range t.queue {
		t.wg.Add(1)
		go func(q chan func()) {
			defer t.wg.Done()
			for fn := range queue {

				fn()
			}
		}(queue)
	}

}

func (t *TaskQueue) Wait() {
	t.wg.Wait()
}

func main() {
	tq := NewTaskQueue(3, 1)
	tq.Run()

	tq.Add("abc", func() {
		fmt.Println("1")
	})

	tq.Add("abc", func() {
		fmt.Println("2")
	})

	tq.Add("awc", func() {
		fmt.Println("3")
	})

	tq.Add("abc", func() {
		fmt.Println("4")
	})

	tq.Add("abc", func() {
		fmt.Println("5")
	})

	tq.Add("abc", func() {
		fmt.Println("6")
	})

	tq.Add("abc", func() {
		fmt.Println("7")
	})

	tq.Add("abc", func() {
		fmt.Println("8")
	})

	tq.Add("abc", func() {
		fmt.Println("9")
	})

	tq.Add("abc", func() {
		fmt.Println("10")
	})

	tq.Add("abhhtgrgc", func() {
		fmt.Println("11")
	})

	tq.Close()
	tq.Wait()

}

func getInt(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum
}
