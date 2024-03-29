package channel

import (
	"context"
	"fmt"
	"time"
)

/*
Reference: https://betterprogramming.pub/writing-better-code-with-go-concurrency-patterns-9bc5f9f73519

#### Error handling in goroutines
1. We create a new struct type Result that couples both result and error
2. In the consumer function, we call Database and return both result and error in the Result struct via the resultCh
3. The main function loops through the resultCh and processes the result and error

The errors in a Goroutine should be coupled with its result type and passed back to the main Goroutine that knows the program’s complete state.

The bottom line is that an error should be considered a first-class citizen. It should be given equal attention as all other parameters in our program.

Again, the consumer closes the resultCh since it’s the sender of the channel.

1. The generator creates a channel and returns it immediately
2. Via a separate Goroutine, the generator feeds the input set into the channel
3. The consumer loops through the channel and processes the data concurrently
4. Note that the generator (the sender) is the one closing the inputCh
*/
type Generator struct {
	Input  []int
	Result Response
}

type Response struct {
	Output interface{}
	err    error
}

/*
sender will get some input and return a channel, over which input will be sent to consumer
*/
func (g *Generator) generate(ctx context.Context) (inputChan chan interface{}) {
	inputChan = make(chan interface{})
	/*
		Important Point:
		Input channel declared here will only be returned, if below code is executed separately in another goroutine
		Else it will wait on sending data to inputChannel  (inputChan <- ip)
	*/
	go func(ctx context.Context) {
		defer close(inputChan) // close input channel, when all input data sent

		for ip := range g.Input {
			select {
			case inputChan <- ip:
				time.Sleep(2 * time.Second) // wait for sometime after sending each input
				fmt.Println("Generator goroutine: sent input - ", ip)
			case <-ctx.Done():
				fmt.Println("Generator goroutine: received cancel event - ", ctx.Err())
				return
			}

		}
	}(ctx)

	return
}

func (g *Generator) consumer(ctx context.Context, inputChan, outputChan chan interface{}) {
	// never forget to close channel, that too via sender
	defer close(outputChan)

	// anonynous function to check even number input
	checkEven := func(val int) (isEven bool, err error) {
		if val%2 == 0 {
			isEven = true
		} else {
			err = fmt.Errorf("Error, input is Odd !!")
		}
		return
	}

	// iterate over input channel
	for ip := range inputChan {
		isEven, err := checkEven(ip.(int))
		op := map[string]interface{}{
			"input":  ip.(int),
			"output": isEven,
		}
		select {
		case outputChan <- Response{Output: op, err: err}:
			fmt.Println("Consumer goroutine: sends output - ", op)
		case <-ctx.Done():
			fmt.Println("Consumer goroutine: received cancel event - ", ctx.Err())
			return
		}
	}
	return
}

func GeneratorPattern() {
	// static inout data
	g := Generator{Input: []int{1, 2, 3, 4, 5}}
	outputChan := make(chan interface{})

	// initialise input channel, on which sender goroutine will push input data
	// WithTimeout() -> will trigger cancel event after speppcific time(6 seconds in this case)
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)

	//ctx, cancel := context.WithCancel(context.Background())
	// to trigger cancel event, before both goroutines send all input and output
	// mock some error and trigger manually
	defer cancel()

	inputChan := g.generate(ctx)

	// initialise output channel, on which consumer goroutine will  push output after processing
	go g.consumer(ctx, inputChan, outputChan)

	// listen to output channel
	// handle output & error accordingly
	for op := range outputChan {
		if opVal, ok := op.(Response); ok == true {
			op := opVal.Output.(map[string]interface{})
			if opVal.err != nil {
				fmt.Printf("\nFor input - %v, Output is : %v", op["input"], opVal.err)
			} else {
				fmt.Printf("\nFor input - %v, Output is : %v", op["input"], op["output"])
			}
		}

	}
}
