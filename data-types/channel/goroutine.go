package channel

import (
	"fmt"
	"runtime"
)

func BasicsGoroutine() {
	fmt.Println("\nruntime.NumCPU(): ", runtime.NumCPU())
}
