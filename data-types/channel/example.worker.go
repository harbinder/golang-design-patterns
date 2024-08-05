package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	workFunc  func(interface{}) interface{}
	workInput interface{}
	workerNum int
}

func InitWork() (cw *Worker) {
	cw = &Worker{}

	// Define work input
	cw.workInput = [][]int{
		{1, 10},
		{11, 20},
	}
	cw.workerNum = len(cw.workInput.([][]int))

	// Define work function
	cw.workFunc = workLogic //workLogicFunc()()

	return
}

func workLogicFunc(num interface{}) func() interface{} {

	return func() (primeFlag interface{}) {
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

}

func workLogic(num interface{}) (primeFlag interface{}) {
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

func (w *Worker) StartWork() (response interface{}) {

	// change timeout to test scenarios
	// 1. time > 3sec : goroutine exit via  closing of channels
	// 2. time < 2 sec : goroutine exit via context cancel trigger
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelCtx()

	//fanOutChan := workFanOut(ctx)
	//fanInChan := workFanIn(ctx, fanOutChan)
	//workOutput(ctx, fanInChan)

	//	w.workOutput(ctx, w.workFanIn(ctx, w.workFanOut(ctx, w.workInputGenerator(ctx))))

	w.workOutput(ctx, w.workFanIn(ctx, w.workFanOutNew(ctx, w.workInputGeneratorNew(ctx))))
	return
}

func (w *Worker) workInputGeneratorNew(ctx context.Context) (ipChanSlice []chan []int) {
	ipChanSlice = make([]chan []int, w.workerNum)
	// Initialise channels in the slice
	for wc := 0; wc < w.workerNum; wc++ {
		ipChanSlice[wc] = make(chan []int)
	}

	go func() {
		// defer close all channels of slice
		for wc := 0; wc < w.workerNum; wc++ {
			defer close(ipChanSlice[wc])
		}

		// iterate input values and send to each slice channel
		for idx, ip := range w.workInput.([][]int) {
			select {
			case ipChanSlice[idx] <- ip:
				fmt.Println("Send input to channel: ipChanSlice", idx, ip)
			case <-ctx.Done():
				fmt.Println("Cancel context timeout!! - workInputGenerator")
				return
			}
		}
		fmt.Println("Close channel: input Generator")

	}()
	return
}

func (w *Worker) workInputGenerator(ctx context.Context) (ipChan chan []int) {
	ipChan = make(chan []int)
	go func() {
		defer close(ipChan)
		for _, ip := range w.workInput.([][]int) {
			select {
			case ipChan <- ip:
				fmt.Println("Send input to channel: ipChan", ip)
			case <-ctx.Done():
				fmt.Println("Cancel context timeout!! - workInputGenerator")
				return
			}
		}
		fmt.Println("Close channel: input Generator")
	}()
	return
}

func (w *Worker) workFanOutNew(ctx context.Context, ipChanSlice []chan []int) (foChan []chan int) {
	foChan = make([]chan int, w.workerNum)
	for wc := 0; wc < w.workerNum; wc++ {
		/*
			Note: cant use append to add fanIn channels to fanOut channel slice
			as it will double the slice length to 20 when appending last channel
			This will make the for loop in func isPrimeFanIn() iterate 20 times and will not exit on 10th iteration and
			goroutine to close fanIn Channel will not execute, as the loop will not exit after 10 iterations
		*/
		//foChan = append(foChan, Work(lowerNumber, upperNumber))
		foChan[wc] = w.WorkNew(ctx, ipChanSlice[wc])
	}
	//}
	return
}

func (w *Worker) workFanOut(ctx context.Context, ipChan chan []int) (foChan []chan int) {
	//batchSize := 100 / w.workerNum
	foChan = make([]chan int, w.workerNum)

	var lowerNumber, upperNumber int
	for wc := 1; wc <= w.workerNum; wc++ {
		for ip := range ipChan {
			fmt.Println("input channel received : ", ip)
			lowerNumber = ip[0] //(wc-1)*batchSize + 1
			upperNumber = ip[1] //wc * batchSize

			/* Note: cant use append to add fanIn channels to fanOut channel slice
			as it will double the slice length to 20 when appending last channel
			This will make the for loop in func isPrimeFanIn() iterate 20 times and will not exit on 10th iteration and
			goroutine to close fanIn Channel will not execute, as the loop will not exit after 10 iterations
			*/
			//foChan = append(foChan, Work(lowerNumber, upperNumber))
			foChan[wc-1] = w.Work(ctx, lowerNumber, upperNumber)
		}
	}
	return
}

func (w *Worker) WorkNew(ctx context.Context, ipChan chan []int) chan int {
	//fmt.Println("Work args: ", lowerNumber, upperNumber)
	ch := make(chan int)

	go func(ctx context.Context) {
		defer close(ch)
		select {
		case <-ctx.Done():
			fmt.Println("Cancel context timeout - Worker")
			return
		case inputRange := <-ipChan:
			lowerNumber, upperNumber := inputRange[0], inputRange[1]
			fmt.Println("Work input range: ", lowerNumber, upperNumber)

			for n := lowerNumber; n <= upperNumber; n++ {
				if w.workFunc(n).(bool) { //workLogicFunc(n)().(bool) {
					ch <- n
					//fmt.Println("Send from channel - fanOut: ", n)
				}

			}
		}
	}(ctx)

	return ch

}
func (w *Worker) Work(ctx context.Context, lowerNumber, upperNumber int) chan int {
	fmt.Println("Work args: ", lowerNumber, upperNumber)
	ch := make(chan int)

	go func(ctx context.Context, lowerNumber, upperNumber int) {
		defer close(ch)
		for n := lowerNumber; n <= upperNumber; n++ {
			//n := n
			//  check prime, push to print channel if its prime and  then close the channel
			if w.workFunc(n).(bool) { //workLogicFunc(n)().(bool) {
				ch <- n
				//fmt.Println("Send Prime No: from channel- fanOut: ", n)
			}

		}
	}(ctx, lowerNumber, upperNumber)

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
					fmt.Println("Cancel context timeout - workFanIn")
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
				fmt.Println("Cancel context timeout!! - workOutput")
				return
			case data, ok := <-respChan:
				if !ok {
					fmt.Println("Close channel - workOutput")
					return
				}
				fmt.Println("Print output: received from channel - fanIn ", data)
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
