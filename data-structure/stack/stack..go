package stack

import (
	"fmt"
)

/*
LIFO : Last Out First In
Operations:
Push - Append to the end of slice
Pop - Remove from end of the slice
Len -
*/

type Stack struct {
	stackSlice []interface{}
	isEmpty    bool
}

func (s *Stack) Pop() (ele interface{}, err error) {
	if len(s.stackSlice) == 0 {
		err = fmt.Errorf("Stack is Empty!!")
		return
	}

	ele = s.stackSlice[len(s.stackSlice)-1]
	s.stackSlice = s.stackSlice[0 : len(s.stackSlice)-1]
	fmt.Println("Pop: ", ele)
	return
}

func (s *Stack) Push(ele interface{}) (err error) {
	fmt.Println("Push: ", ele)
	s.stackSlice = append(s.stackSlice, ele)
	return
}

func (s *Stack) Print() (err error) {
	fmt.Println("Print Stack:")
	for ele := len(s.stackSlice) - 1; ele >= 0; ele-- {
		fmt.Println(s.stackSlice[ele])
	}
	return
}

func StackExample() {
	stack := Stack{}
	if _, er := stack.Pop(); er != nil {
		fmt.Println(er)
	}
	for i := 0; i < 6; i++ {
		stack.Push(i)
	}
	stack.Print()
	stack.Pop()
	stack.Push(8)
	stack.Print()
}
