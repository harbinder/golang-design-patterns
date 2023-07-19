package channel

import (
	"fmt"
	"sync"
	"time"
)

/*
As you spawn more Goroutines to process requests concurrently, this leaves us with another problem.

What happens if all your Goroutines access the same shared resources, say a remote cache?

Bombarding your cache with an unbounded number of concurrent requests is a surefire recipe to bring down your cache immediately.

This is where the Semaphore comes in handy.

Unlike mutex lock, which allows a single thread to access a resource at a time, Semaphore allows N threads to access a resource at a time.

Using the concept of a buffered channel, we can design a semaphore easily.
*/

// Mutex can be used only for 1 Thread
// WHEREAS Semaphore can be used for N THREADS using BUFFERED CHANNEL
/*
1. The NewSemaphore initiates a Semaphore by creating a buffered channel with the capacity of maxReq
2. When a Goroutine Acquire a semaphore, we send an empty struct to semaCh
3. When the buffered channel is full, call to Acquire will be blocked
4. When a Goroutine Release a semaphore, an empty struct will be sent out of the channel, creating space in the buffered channel for subsequent Acquire
*/
type Semaphore struct {
	// channel of empty struct
	semaChan chan struct{}
}

func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semaChan: make(chan struct{}, maxReq),
	}
}

// Acquire method similar to LOCK in sync.mutex
func (s *Semaphore) Acquire() {
	s.semaChan <- struct{}{}
}

// Release method similar to UNLOCK in sync.mutex
func (s *Semaphore) Release() {
	<-s.semaChan
}

/*
 1. We create a semaphore with the capacity of 3
 2. If semaphore Acquired in for loop, only specific goroutines will be spawned at a time
    Else if semaphore Acquired in goroutine, then all 10 goroutinrs will be spawned together
 3. Each Goroutine acquires a semaphore before processing or before being spawned, based on your choice in point 2 above.
 4. Since there are 10 tasks and the maximum number of concurrent tasks is 3,
    the total time needed to process all tasks will be 4 seconds (Each task takes 1 second)

Sample Response:
13:50:24 Running worker 2
13:50:24 Running worker 3
13:50:24 Running worker 1

13:50:25 Running worker 5
13:50:25 Running worker 6
13:50:25 Running worker 4

13:50:26 Running worker 8
13:50:26 Running worker 7
13:50:26 Running worker 9

13:50:27 Running worker 10
*/
func SemaphoreExample() {
	maxParallelReq := 3
	s := NewSemaphore(maxParallelReq)

	wg := sync.WaitGroup{}
	for req := 1; req <= 10; req++ {
		wg.Add(1)
		s.Acquire()
		go func(reqID int) {
			defer wg.Done()
			defer s.Release()

			// Do some work in goroutine
			msg := fmt.Sprintf(
				"%s Running worker %d",
				time.Now().Format("15:04:05"),
				reqID,
			)
			fmt.Println(msg)
			time.Sleep(time.Second * 1)

		}(req)
	}
	wg.Wait()
}
