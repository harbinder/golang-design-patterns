package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	workFunc  func(interface{}) interface{}
	workerNum int
	batchSize int
}

func InitWork() (cw *Worker) {
	cw = &Worker{
		workerNum: 10,
		batchSize: 1,
	}

	// Define work
	cw.workFunc = func(num interface{}) (primeFlag interface{}) {
		primeFlag = true
		for n := num.(int) - 1; n > 1; n-- {
			if num.(int)%n == 0 {
				primeFlag = false
				break
			}
		}
		// if every prime number calculation takes 200ms
		// then 100 numbers will take 20000ms ~ 20s
		// So if we add a context with timeout < 20s, it should exit all goroutines gracefully
		time.Sleep(time.Millisecond * 200)
		return
	}

	return
}

/*
Print prime numbers 1-100

- Calculate prime in 1 goroutine
- Print in another

*/

func (w *Worker) StartWork() (response interface{}) {

	// change timeout to test scenarios
	// 1. time > 3sec : goroutine exit via  closing of channels
	// 2. time < 2 sec : goroutine exit via context cancel trigger
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelCtx()

	//fanOutChan := workFanOut(ctx)
	//fanInChan := workFanIn(ctx, fanOutChan)
	//workOutput(ctx, fanInChan)

	w.workOutput(ctx, w.workFanIn(ctx, w.workFanOut(ctx)))

	return
}

func (w *Worker) workFanOut(ctx context.Context) (foChan []chan int) {
	batchSize := 100 / w.workerNum
	foChan = make([]chan int, w.workerNum)

	var lowerNumber, upperNumber int
	for wc := 1; wc <= w.workerNum; wc++ {
		lowerNumber = (wc-1)*batchSize + 1
		upperNumber = wc * batchSize
		/* Note: cant use append to add fanIn channels to fanOut channel slice
		as it will double the slice length to 20 when appending last channel
		This will make the for loop in func isPrimeFanIn() iterate 20 times and will not exit on 10th iteration and
		goroutine to close fanIn Channel will not execute, as the loop will not exit after 10 iterations
		*/
		//foChan = append(foChan, Work(lowerNumber, upperNumber))
		foChan[wc-1] = w.Work(ctx, lowerNumber, upperNumber)
	}
	return
}

func (w *Worker) Work(ctx context.Context, lowerNumber, upperNumber int) chan int {
	ch := make(chan int)

	go func(ctx context.Context) {
		defer close(ch)
		for n := lowerNumber; n <= upperNumber; n++ {
			//  check prime, push to print channel if its prime and  then close the channel
			if w.workFunc(n).(bool) {
				ch <- n
				//fmt.Println("Send Prime No: from channel- fanOut: ", n)
			}

		}
	}(ctx)

	return ch

}

func (w *Worker) workFanIn(ctx context.Context, foChanSlice []chan int) (fiChan chan int) {
	fiChan = make(chan int)
	fmt.Println("foChanSlice: count - ", len(foChanSlice))

	wg := &sync.WaitGroup{}
	for _, ch := range foChanSlice {
		wg.Add(1)
		ch := ch
		go func(ch chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Cancel context timeout!! - isPrimeFanIn")
					return
				case data, ok := <-ch:
					if !ok {
						fmt.Println("Close channel - fanOut")
						return
					}
					//fmt.Println("Receive Prime No: from channel - fanOut: ", data)
					fiChan <- data
				}
			}
			/*
				for data := range ch {
					fmt.Println("Receive Prime No: from channel - fanOut: ", data)
					fiChan <- data
				}
			*/
			//fmt.Println("loop done:", data)
		}(ch)
	}

	go func() {
		//fmt.Println("Goroutine-FanIn Channel")
		defer close(fiChan)
		wg.Wait()
		fmt.Println("Close channel: fanIn")
	}()

	return
}

func (w *Worker) workOutput(ctx context.Context, respChan chan int) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancel context timeout!! - printPrime")
				return
			case data, ok := <-respChan:
				if !ok {
					fmt.Println("Close channel - print")
					return
				}
				fmt.Println("Print Prime No: received from channel - fanIn ", data)
			}
		}
		/*
			for n := range respChan {
				fmt.Println("Print Prime No: received from channel - fanIn ", n)
			}
		*/
	}()

	wg.Wait()
}

func calculatePrime(num int) (primeFlag bool) {
	primeFlag = true
	for n := num - 1; n > 1; n-- {
		if num%n == 0 {
			primeFlag = false
			break
		}
	}
	// if every prime number calculation takes 200ms
	// then 100 numbers will take 20000ms ~ 20s
	// So if we add a context with timeout < 20s, it should exit all goroutines gracefully
	time.Sleep(time.Millisecond * 200)
	return
}

func checkPrimeRecursive(num int, divisor int) bool {
	if num <= 2 {
		return true
	}

	if num%divisor == 0 {
		return false
	}
	if divisor == num {
		return true
	}
	return checkPrimeRecursive(num, divisor+1)

}
