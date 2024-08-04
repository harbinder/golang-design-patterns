package structs

import (
	"fmt"
	"sync"
	"time"
)

/*
Ref: https://eli.thegreenplace.net/2018/beware-of-copying-mutexes-in-go/

Embedding structs in structs could be tricky

Keep 2 things in mind while embedding wrt promoted fields:
1. type of promoted fields - value type or pointer type
2. Receiver type of methods - value or pointer type


*/

type Container struct {
	sync.Mutex                 // embedded struct (value type mutex struct)
	pointerLock *sync.Mutex    // embedded struct (pointer type mutex struct)
	counter     map[string]int // (pointer type)
}

/*
incValueReceiverNoLock: increments map key WITHOUT lock AND object is VALUE type receiver
*/
func (c Container) incValueReceiverNoLock(key string) {
	c.counter[key]++
}

/*
incPointerReceiverNoLock: increments map key WITHOUT lock AND object is POINTER type receiver
*/
func (c *Container) incPointerReceiverNoLock(key string) {
	c.counter[key]++
}

/*
incValueReceiverWithLock: increments map key WITH lock AND object is VALUE type receiver
*/
func (c Container) incValueReceiverWithLock(key string) {
	c.Lock()
	defer c.Unlock()
	c.counter[key]++
}

/*
incPointerReceiverWithLock: increments map key WITH lock AND object is POINTER type receiver
*/
func (c *Container) incPointerReceiverWithLock(key string) {
	c.Lock()
	defer c.Unlock()
	c.counter[key]++
}

/*
incValueReceiverWithPointerLock: increments map key WITH lock AND object is VALUE type receiver
*/
func (c Container) incValueReceiverWithPointerLock(key string) {
	c.pointerLock.Lock()
	defer c.pointerLock.Lock()
	c.counter[key]++
}

func EmbeddingExample() {
	c := Container{
		counter: map[string]int{
			"a": 0, "b": 0,
		},
		pointerLock: &sync.Mutex{},
	}

	doIncrementNoLock := func(key string) {
		for i := 0; i < 100000; i++ {
			c.incValueReceiverNoLock(key)
			c.incPointerReceiverNoLock(key)
		}
	}

	/*

	 */
	doIncrementNoLock("a") // This will work, but wrong counter value
	//go doIncrementNoLock("a") // This will give error -> fatal error: concurrent map writes

	fmt.Println(c.counter)

	doIncrementWithLock := func(key string) {
		for i := 0; i < 100000; i++ {
			c.incValueReceiverWithLock(key)
			c.incPointerReceiverWithLock(key)
		}
	}

	doIncrementWithLock("b") // This will work, but wrong counter value
	//go doIncrementWithLock("b")  // This will give error -> fatal error: concurrent map writes
	fmt.Println(c.counter)

	doIncrementWithLockAndPointerReceiver := func(key string) {
		for i := 0; i < 100000; i++ {
			//c.incValueReceiverWithLock(key)
			c.incPointerReceiverWithLock(key)
		}
	}

	go doIncrementWithLockAndPointerReceiver("b") // Only this will work PROPERLY with concurrent goroutines. Why?????
	time.Sleep(time.Second * 5)
	fmt.Println(c.counter)
	/*
		doIncrementNoLock() : With no lock, goroutines will not work at all.
		When we concurrently try to access map, it will return error as map is a pointer type value i.e. a single variable being accessed concurrently
		Irrespective of calling this method with same object (c) or a copy of it, counter variable will be pointing to same map

		doIncrementWithLock() : With  lock, goroutine SHOULD work.
		BUT it will NOT WORK with
		c.incValueReceiverWithLock() -> object c as Value receiver, BECAUSE sync.Mutex struct is of Value type and when a COPY of c object is created with Value receiver call,
		 it creates a copy of sync.Mutex struct as well. So actually we are not working with the same c.Lock() for each goroutine
		So ONLY   the last option works correctly.
		c.incPointerReceiverWithLock()

		NOTE: Moreover, we can also try with creating Named embedding of sync.Mutex along with Pointer type Value.


	*/

	doIncrementWithPointerLockAndValueReceiver := func(key string) {
		for i := 0; i < 100000; i++ {
			//c.incValueReceiverWithLock(key)
			c.incValueReceiverWithPointerLock(key)
			//c.incPointerReceiverWithLock(key)
		}
	}

	go doIncrementWithPointerLockAndValueReceiver("b")

	time.Sleep(time.Second * 5)
	fmt.Println(c.counter)
}
