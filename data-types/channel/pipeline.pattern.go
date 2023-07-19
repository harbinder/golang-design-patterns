package channel

import "fmt"

/*
Reference: https://betterprogramming.pub/writing-better-code-with-go-concurrency-patterns-9bc5f9f73519

A pipeline is a series of stages that takes in data, processes them, and passes them to another stage.

In the first example, we add 1 to the input and pass the result to the multiply stage for further processing.

# The benefit of a pipeline is evident

It separates the concerns of each stage in the pipeline. Each stage is responsible for one and only one thing.
The stages are modular and allow us to mix and match how stages are combined.
The stages in the example above run sequentially. Each can only begin after the previous stage has processed all the data.

Leveraging the Goroutine and channel, stages can run and process data concurrently.
*/

/*
1. We create a data stream using the generator function
2. We create a doneCh and pass to all Goroutines for explicit cancellation
3. We then chain the add and multiply stage together
4. Whenever the add function has done processing an input. It will immediately pass the result to the multiply stage for further processing

Each stage processes the data concurrently and immediately passes it to the next stage once it’s done.

Moreover, the multiply and the add stage can be mixed and matched to produce different results.


#### FAN-OUT FAN-IN : Pattern

Following the pipeline pattern, what happens if one of your pipeline stages is more computationally expensive and requires a longer process time?
All upstream data will be stranded, and downstream stages will be left idle.
The most intuitive solution will be to increase the number of workers where work is most hefty.
This is where the Fan-out, Fan-in pattern comes into play.

Fan-out, in a nutshell, refers to spawning more Goroutines in a specific stage to increase its throughput,
in other words, DEMULTIPLEXING.
*/

func PipelinePattern() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator1(doneCh, input)

	resultCh := multiply(doneCh, add(doneCh, inputCh))

	for res := range resultCh {
		fmt.Println(res)
	}
}

func generator1(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}
func add(doneCh chan struct{}, inputCh chan int) chan int {
	addRes := make(chan int)

	go func() {
		defer close(addRes)

		for data := range inputCh {
			result := data + 1

			select {
			case <-doneCh:
				return
			case addRes <- result:
			}
		}
	}()

	return addRes
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	multiplyRes := make(chan int)

	go func() {
		defer close(multiplyRes)

		for data := range inputCh {
			result := data * 2

			select {
			case <-doneCh:
				return
			case multiplyRes <- result:
			}
		}
	}()

	return multiplyRes
}
