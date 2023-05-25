package channels

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GoChannel struct {
}

func (gc GoChannel) RoutineOne() {
	fmt.Println("RoutineOne")

	messageList := []int{1, 2, 3, 4}
	for key, msg := range messageList {
		fmt.Println(key, msg)
	}
}

func (gc GoChannel) CheckOsSignal() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	SigChan := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked

	go func() {
		signal.Notify(SigChan, os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		cancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("Start: Goroutine-1")
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		select {

		case <-ctx.Done():
			fmt.Println("End: Goroutine-1")
			return
			/*
				case <-SigChan:
					fmt.Println("Goroutine-1")
					wg.Done()
					return
			*/
		}
	}(wg)
	wg.Add(1)
	fmt.Println("Start: Goroutine-2")
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		select {

		case <-ctx.Done():
			fmt.Println("End: Goroutine-2")

			return
			/*
				case <-SigChan:
					fmt.Println("Goroutine-2")
					wg.Done()
					return
			*/
		}
	}(wg)

	fmt.Println("Waiting for goroutines to end.....")
	wg.Wait()
	fmt.Println("Shutdown Goroutines....")
}

func (gc GoChannel) ShutdownViaChannel() {
	SigChan := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(SigChan, os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	for {
		select {
		case <-SigChan:
			fmt.Println("Stop Working....")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Working....")
		}
	}
}

// 10
// G1-1
// G2-2
// .
// G5 -5
// .
// G(1-5) -
// R1 - PROCESSING
// R2
// R3
// R4
// .
// .
// .
// R100

// P-1
// G1-R1
// G2-R4
// P-2
// G1-R2
// G2-R3

// for 5 {
// 		SELECT+UPDATE
// 		finOneAndUpdate("CREATED")
// 	}
// 	G1
// 	R1 - PROCESSING

// }
