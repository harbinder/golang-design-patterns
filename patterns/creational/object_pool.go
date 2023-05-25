package creational

import (
	"fmt"
	"strconv"
	"sync"
)

/*
The Object Pool Design Pattern is a creational design pattern in which a pool of objects is initialized and
created beforehand and kept in a pool.
As and when needed, a client can request an object from the pool, use it, and return it to the pool.
The object in the pool is never destroyed.

When to Use:
1. When the cost to create the object of the class is high and
the number of such objects that will be needed at a particular time is not much.
	-Let’s take the example of DB connections.
	Each of the connection object creation is cost is high as there is network calls involved and also at a time
	not more than a certain connection might be needed. The object pool design pattern is perfectly suitable for such cases.

2. When the pool object is the immutable object
	-Again take the example of DB connection again. A DB connection is an immutable object.
	Almost none of its property needs to be changed

3. For performance reasons. It will boost the application performance significantly since the pool is already created
*/

type iPoolObject interface {
	getID() string
}

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

func (p *pool) borrowObj() (po iPoolObject, err error) {
	// acquire lock on pool
	p.mulock.Lock()
	defer p.mulock.Unlock()

	// check if no object idle in pool
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free/idle. Please request in someitme!")
	}

	// move object from idle to active list
	po = p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, po)

	return //return the pool object
}

func (p *pool) returnObj(po iPoolObject) (err error) {
	// acquire lock on pool
	p.mulock.Lock()
	defer p.mulock.Unlock()

	// check if
	if len(p.idle) == p.capacity {
		return fmt.Errorf("Pool already full with idle objects, unable to return object to pool!")
	}

	// move object from active to idle list
	if err = p.removeActiveObj(po); err != nil {
		return
	}
	p.idle = append(p.idle, po)

	return
}

func (p *pool) removeActiveObj(targetObj iPoolObject) (err error) {
	/*
		// Cannot create new slice
			var newActivePool []iPoolObject
			for obj := 0; obj < p.capacity; obj++ {
				if p.active[obj].getID() == targetObj.getID() {
					newActivePool = p.active[0 : obj-1]
					newActivePool = append(newActivePool, p.active[obj+1:]...)
					p.active = newActivePool
				}
			}
	*/
	currActiveLen := len(p.active)
	for i, obj := range p.active {
		if targetObj.getID() == obj.getID() {
			// need to check this logic ?
			p.active[i] = p.active[currActiveLen-1] // copy last element of slice to this location
			p.active = p.active[:currActiveLen-1]   // reslice again and remove the last duplicate value
			return
		}
	}
	return
}

func InitPool(poolObjs []iPoolObject) (p *pool, err error) {
	if len(poolObjs) == 0 {
		return nil, fmt.Errorf("No pool objects to be initialised!")
	}

	p = &pool{
		idle:     poolObjs,
		active:   make([]iPoolObject, 0),
		capacity: len(poolObjs),
		mulock:   &sync.Mutex{}, // new(sync.Mutex)
	}
	return
}

/*
Create connection entity
This could be DB connections or logger connections or any worker pool  connections
*/
type connectionEntity struct {
	Id string
}

func (ce connectionEntity) getID() (id string) {
	return
}
func ExecuteObjectPool() {
	TOTAL_OBJ := 5
	//poolObjs := make([]iPoolObject, TOTAL_OBJ)
	var poolObjs []iPoolObject
	for i := 1; i <= TOTAL_OBJ; i++ {
		poolObjs = append(poolObjs, connectionEntity{Id: strconv.Itoa(i)})
	}

	var (
		p   *pool
		err error
	)

	if p, err = InitPool(poolObjs); err != nil {
		err = fmt.Errorf("Unable to initialise object pool!")
		fmt.Println(err)
	}

	fmt.Println("Initialised Object pool")
	fmt.Printf("Object pool  : %+v", p)
	fmt.Println()

	//  borrow object
	var obj iPoolObject
	for i := 1; i <= 2; i++ {
		fmt.Println("Borrow object from pool")
		obj, _ = p.borrowObj()
		fmt.Println()
		fmt.Printf("Object pool : %+v", p)
		fmt.Println()
	}

	//  return object
	fmt.Println("Return object to pool")
	p.returnObj(obj)
	fmt.Println()
	fmt.Printf("Object pool : %+v", p)

}
