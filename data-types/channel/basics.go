package channel

import (
	"fmt"
	"sync"
	"time"
)

/*
Channel is a type, used to share(send/receive) data (ANY data type) between go-routines
Channel can be closed
- Send/Receive ONLY
- Buffered/Unbuffered :
Sending on Full Channel will WAIT
Receiving from EMPTY channel will WAIT
Sending on Closed Channel  will PANIC
ch := make(chan bool)
<-chan (Receive ONLY)
chan<- (Send ONLY)

Select Statement
select {

}

Reference Link : https://golangbyexample.com/channel-golang/
A channel is internally represented by a hchan struct whose main elements are:

type hchan struct {
    qcount   uint           // total data in the queue
    dataqsiz uint           // size of the circular queue
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    elemsize uint16
    closed   uint32         // denotes weather channel is closed or not
    elemtype *_type         // element type
    sendx    uint           // send index
    recvx    uint           // receive index
    recvq    waitq          // list of recv waiters
    sendq    waitq          // list of send waiters
    lock     mutex
}
When using make, an instance of hchan struct is created and all the fields are initialized to their default values.
ch := make(chan {type}, capacity)

### 2 Operations on Channel ###
1. Send : The send operation is used to send data to the channel.
ch<-val
2. Receive: The receive operation is used to read data from the channel.
val := <-ch

### 2 Types of Channel ###
Unbuffered : capacity = zero
Buffered : capacity  > zero

### Operation on 2 Types ###
Unbuffered :
	Send: Blocks unless there is other goroutine to receive
	Receive: Blocks until there is other goroutine on the other side to send
Buffered :
	Send: Blocks if channel is Full
	Receive: Bocks if channel is Empty
PS: So sending and receiving in the SAME Goroutine is only possible for a BUFFERED channel.

### Channel Direction ###
1. Bi-directional
ch := make(chan {type})
2. Unidirectional
ONLY Send To:    ch := make(chan<- {type})
ONLY Receive From: ch := make(<-chan {type})

chan     :bidirectional channel (Both read and write)
chan <-  :only writing to channel
<- chan  :only reading from channel (input channel)

### Close a Channel ###
Close is an inbuilt function that can be used to close a channel.
Closing of a channel means that no more data can be send to the channel.
Channel is generally closed when all the data has been sent and there's no more data to be send.
1. PANIC: Sending on a close channel will cause a panic.
2. PANIC: Also closing an already closed channel will cause a panic

While receiving from a  channel we can also use an additional variable to determine if the channel has been closed.
val, NotClosed := <-ch
NotClosed : true/false

### Nil Channel ###
Some points to note about nil channel
1. SENDING to a  NIL channel blocks FOREVER
2. RECEIVING from NIL channel blocks FOREVER
3. CLOSING a NIL channel results in PANIC

Summary
Command	| Unbuffered Channel(Not Closed and not nil) |	Buffered Channel(Not Closed and not nil) |	Closed Channel |	Nil Channel
Send	  Block if there is is no corresponding receiver otherwise success |	Block if the channel is full otherwise success	| Panic |	Block forever
Receive	  Block if there is no corresponding sender otherwise success	Block if the channel is empty otherwise success	Receives the default value of data type from the channel if channel is empty else  receives the actual value |	Block forever
Close	  Success |	Success |	Panic |	Panic
Length	  0	 | Number of elements queued in the buffer of the channel	-0 if unbuffered channel-Number of elements queued in the buffer if buffered channel |	0
Capacity  0	| Size of the buffer of the channel	-0 if unbuffered channel-Size of the buffer if buffered channel |	0


Pointers:
1. Channel Must Always Be Closed by the Sender
2. We can use  Context package to close channels and return from gorouotines
3. We can use WaitGroup struct of sync package as well
*/

func Basic() {
	BasicsGoroutine()

	result := 0
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int, ch chan int, wg *sync.WaitGroup) {
			result += i * 2
			fmt.Printf("\nResult: %v - %v", i, result)
			//ch <- result
			wg.Done()
		}(i, ch, wg)
	}

	//time.Sleep(time.Second * 5)
	/*
		cnt := 0
		for v := range ch {
			cnt++
			fmt.Println()
			fmt.Printf("Result: %v - %v", v, result)
			if cnt == 5 {
				break
			}
		}
	*/
	wg.Wait()
	fmt.Println()
	fmt.Printf("Result: %v", result)

	checkClosedChan()

	closeChannelViaSender()

}

func checkClosedChan() {
	fmt.Println("\nCheck Closed Channel")
	ch := make(chan int, 1)
	ch <- 2
	val, open := <-ch
	fmt.Printf("Val: %d Open: %t\n", val, open)

	close(ch)
	val, open = <-ch
	// Read : default data type value, AND false as channel is closed
	fmt.Printf("Val: %d Close: %t\n", val, open)

	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	// what happens if you don't close a channel in the main function.
	// goroutine below will leak, it will wait on channel to receive value
	close(ch1)
	go func() {
		fmt.Printf("Range loop a channel\n")
		for val := range ch1 {
			fmt.Printf("Received Value: %v\n", val)
		}
		fmt.Printf("Channel Closed\n")
	}()

	time.Sleep(time.Second * 2)
}

func closeChannelViaSender() {
	fmt.Println("Example: Close channel via sender only")

	ch := make(chan int)
	go senderChannel(ch)
	var sum int
	for v := range ch {
		fmt.Println("Response: ", v)
		sum += v
	}
	fmt.Println("Total Response: ", sum)
	fmt.Println("Channel closed now")
}

func senderChannel(ch chan<- int) {
	wg := sync.WaitGroup{}
	// spawn 5 goroutines
	// wait for each to complete
	// send response of each to calling function
	// close the  channel in the end
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- i * 2
		}(i, &wg)
	}
	wg.Wait()
	close(ch)
}
