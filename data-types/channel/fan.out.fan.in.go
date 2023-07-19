package channel

import (
	"fmt"
	"sync"
)

/*
Fan-out, Fan-in

Following the pipeline pattern, what happens if one of your pipeline stages is more computationally expensive and requires a longer process time?

All upstream data will be stranded, and downstream stages will be left idle.

The most intuitive solution will be to increase the number of workers where work is most hefty.

This is where the Fan-out, Fan-in pattern comes into play.
*/

func fanOut(doneCh chan struct{}, inputCh chan int) []chan int {
	numWorkers := 5
	channels := make([]chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// TODO: rather than spawning N worker goroutines all together, we can use Semaphore pattern
		// to spawn in controlled way, so as not to choke the system/service/DB at the other end

		addResultCh := add(doneCh, inputCh) // this function present in pipeline.pattern.go file
		channels[i] = addResultCh
	}

	return channels
}

/*
1. The fanIn function takes in a slice of channels (The fanOut function produces them)
2. For each channel, we spawn a separate Goroutine to fetch the data from the channel and feed it to the finalCh.
3. At the end of the function, we spawn a separate Goroutine to wait for all Goroutines to finish and close the finalCh
4. We then return the merged channel — finalCh back to the main function
5. The chClosure acts as a closure to the ch in each iteration
*/
func fanIn(doneCh chan struct{}, resultChs ...chan int) chan int {
	finalCh := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range resultChs {
		wg.Add(1)
		chClosure := ch

		go func() {
			defer wg.Done()

			for data := range chClosure {
				select {
				case <-doneCh:
					return
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(finalCh)
	}()

	return finalCh
}

/*
1. We create a data stream inputCh using a generator
2. We spawn ten workers for our add function using fanOut
3. We merge all channels using the fanIn function
4. We then pass the addResultCh into the multiply stage for further processing
With the Fan-out, Fan-in pattern, we can increase the number of workers for a single stage of our pipeline,
thus increasing the throughput of our program.
*/
func FanOutFanInPattern() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator1(doneCh, input) // this function present in pipeline.pattern.go file

	// As more goroutines are required to process add() task,
	// we are using fanOut-fanIn pattern here
	channels := fanOut(doneCh, inputCh)
	addResultCh := fanIn(doneCh, channels...)

	resultCh := multiply(doneCh, addResultCh) // this function present in pipeline.pattern.go file

	for res := range resultCh {
		fmt.Println(res)
	}
}
