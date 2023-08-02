package creational

import (
	"fmt"
	"sync"
)

/*
There are 3 ways yto create an object once
1. Using init()
2. Using sync.Mutex struct - Lock/Unlock Methods
3. Using sync.Once struct - Do Method
*/
type single struct {
}

var singleInstance *single

var lock = &sync.Mutex{}
var once sync.Once

var flagSyncOnce = true

func getSingleton() *single {
	if singleInstance == nil {
		if flagSyncOnce {
			once.Do(
				func() {
					singleInstance = &single{}
				},
			)
		} else {
			lock.Lock()
			defer lock.Unlock()
			if singleInstance == nil {
				singleInstance = &single{}
			} else {
				fmt.Println("Concurrent Call - Instance already created!")
			}
		}
	} else {
		fmt.Println("Async Call - Instance already created!")
	}
	return singleInstance
}

func ExecuteSingleton() {
	fmt.Println("Creating Singleton Instance:")
	for i := 0; i < 10; i++ {
		go getSingleton()
	}
	fmt.Scanln()
}
