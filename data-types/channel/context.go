package channel

/*
context Module:
Context is an argument passed to your functions and Goroutines and
allows you to stop them promptly should you not require them anymore.

#### What Is the Context Module?
Typical usage of the context module is when a client terminates the connection with a server.
What if the termination occurs while the server is in the middle of some heavy lifting work or database query?
The context module allows these processes to be stopped instantly as soon as they are not further in need.

The usage of the context module boils down to three primary parts
1. Listening to a cancellation event
2. Emitting a cancellation event
3. Passing request scope data

##### Usage-1: Listening to Cancellation Event

Context is an interface type.

type Context interface {
    Done() <- chan struct{}
    Err() error

    Deadline() (deadline time.Time, ok bool)
    Value(key interface{}) interface{}
}

The Done() :
function returns a channel that receives an empty struct when a context is cancelled.
The Err() :
function returns a non-nil error in the event of cancellation otherwise, it returns a nil value.

Example:

select {
case <-ctx.Done():
	fmt.Println(ctx.Err())
	return
default:
	// do some work
}

##### Usage-2: Emitting a Cancellation Event
There are 3 functions in context module, used to trigger cancel event
The context module provides three functions that return a CancelFunc .
Calling the cancelFunc emits an empty struct to the ctx.Done() channel
and notifies downstream functions that are listening to it.

1. func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
2. func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
3. func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

All these have 1 argument as PARENT context.
There is a Context TREE with parent and child context, which can be used to cancel specific
sub-tree or route and even all context via root/parent context.

## How to create the PARENT/ROOT context
1. func Background() Context
2. func TODO() Context

These functions output an empty context that does nothing AT ALL. It cannot be cancelled nor carry a value.
Their primary purpose is to serve as a root context that will later be passed to any of the WithX functions to create a cancellable context.

Example:
rootCtx := context.Background()

child1Ctx, cancelFunc1 := context.WithCancel(rootCtx)
child2Ctx, cancelFunc2 := context.WithCancel(rootCtx)

child3Ctx, cancelFunc3 := context.WithCancel(child1Ctx)

It has 2 sub-trees :
rootCtx -> child1Ctx -> child3Ctx
rootCtx -> child2Ctx

When we call cancelFunc1, we will cancel child1Ctx and child3Ctx, while leaving child2Ctx unaffected.


####### Example-1: WithCancel

If the databaseQuery returns an error, the cancel function will be invoked. operation1 will then be notified via ctx.Done() and exits gracefully.

You can find the cancellation reasons in ctx.Err().

func handler() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)

    go operation1(ctx)

    data, err := databaseQuery()
    if err != nil {
        cancel()
    }
}

func operation1(ctx context.Context) {
    for {
        select {
            case <- ctx.Done():
                fmt.Println(ctx.Err().Error())
                return

            default:
                fmt.Println("Do something")
        }
    }
}

#### Example-2: WithTimeout
func handler() {
    ctx := context.Background()

    ctx, cancel := context.WithTimeout(
        ctx,
        3*time.Second,
    )
    defer cancel()

    dataChan := make(chan string)
    go databaseQuery(dataChan)

    select {
        case <- dataChan:
            fmt.Println("Query succeeds, do something")

        case <- ctx.Done():
            fmt.Println("Timeout exceeded, returning")
            return
    }
}

##### Usage-3: Passing Request Scope Data

As we usually pass the ctx variable across functions, request scope data can tag along this
variable using the WithValue function.

Considering an application that involves multiple function calls, we can pass a traceID to
these functions for monitoring and logging via the ctx variable.

Example:
const traceIDKey = "trace_id"

func main() {
    ctx := context.Background()
    ctx = context.WithValue(
        ctx,
        traceIDKey,
        "random_id123",
    )

    function1(ctx)
}

func function1(ctx context.Context) {
    log.Println(
        "Entered function 1 with traceID: ",
        ctx.Value(traceIDKey),
    )

}

Caveats & Best Practices:
1. Always Defer Cancel Function

	When you spawn a new cancellable context via the WithCancel function, the module will

	Spawn a new Goroutine in the background to propagate the cancellation event to all children if the cancel function is invoked
	Keep track of all the children contexts in the struct of the parent context
	If a function returns without cancelling the context, the Goroutine and the child contexts will remain in the memory indefinitely causing a memory leak.

	This also applies to WithTimeout and WithDeadline except, these functions automatically cancel the context when the deadline is exceeded.

	However, it’s still a best practice to defer the cancellation for any of the WithX functions.
*/
