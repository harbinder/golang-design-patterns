package linklist

import (
	"fmt"
)

type NodeDouble struct {
	Data interface{}
	Next *NodeDouble
	Prev *NodeDouble
}

type ListDouble struct {
	Head *NodeDouble
	Tail *NodeDouble
	Len  int
}

func InitListDouble() (ld *ListDouble) {
	return &ListDouble{}
}

func (ld *ListDouble) Size() (l int) {
	return ld.Len
}

func (ld *ListDouble) AddFront(data interface{}) {
	fmt.Println("Add Front Node: ", data)
	node := &NodeDouble{Data: data}
	if ld.Head == nil { // || ld.Len == 0
		ld.Head = node
		ld.Tail = node
	} else {
		node.Next = ld.Head
		ld.Head.Prev = node
		ld.Head = node
	}
	ld.Len++
}

/*
 */
func (ld *ListDouble) AddBack(data interface{}) {
	fmt.Println("Add Back Node: ", data)
	node := &NodeDouble{Data: data}
	if ld.Len == 0 { // || ld.Tail == nil
		ld.Head = node
		ld.Tail = node
	} else {
		node.Prev = ld.Tail
		ld.Tail.Next = node
		ld.Tail = node
	}
	ld.Len++
}

func (ld *ListDouble) DeleteFront() {
	fmt.Println("Delete Front Node")
	if ld.Len == 0 {
		fmt.Printf("List is Empty !!")
		return
	}

	if ld.Len == 1 {
		ld.Head = nil
		ld.Tail = nil
	} else {
		firstNode := ld.Head
		ld.Head = firstNode.Next // point head to second node
		firstNode.Next = nil     // remove first node
		ld.Head.Prev = nil       // new first node
	}
	ld.Len--
}

func (ld *ListDouble) DeleteBack() {
	fmt.Println("Delete Back Node")
	if ld.Len == 0 {
		fmt.Printf("List is Empty !!")
		return
	}
	if ld.Len == 1 {
		ld.Head = nil
		ld.Tail = nil
	} else {
		lastNode := ld.Tail
		ld.Tail = lastNode.Prev
		lastNode.Prev = nil
		ld.Tail.Next = nil
		lastNode = nil
	}
	ld.Len--
}

func (ld *ListDouble) TraverseForward() {
	fmt.Print("Traverse Forward: ")

	if ld.Len == 0 { //|| ld.Head == nil {
		fmt.Printf("List is Empty !!")
		return
	}

	current := ld.Head
	for current != nil {
		fmt.Printf("<->%v", current.Data)
		current = current.Next
	}

}

func (ld *ListDouble) TraverseReverse() {
	fmt.Print("\nTraverse Reverse: ")

	if ld.Len == 0 { //|| ld.Tail == nil {
		fmt.Printf("List is Empty !!")
		return
	}

	current := ld.Tail
	for current != nil {
		fmt.Printf("<->%v", current.Data)
		current = current.Prev
	}

}

func DoublyListExample() {
	ld := InitListDouble()

	list := []int{1, 2, 3, 4, 5}
	//list := []string{"Father", "Mother", "Son", "Daughter"}

	fmt.Println("List Length: ", ld.Size())
	ld.TraverseForward()
	ld.TraverseReverse()
	fmt.Println()

	for _, v := range list {
		//ld.AddFront(v)
		ld.AddBack(v)
	}

	fmt.Println("List Length: ", ld.Size())
	ld.TraverseForward()
	ld.TraverseReverse()

	fmt.Println()

	ld.DeleteBack()
	fmt.Println("List Length: ", ld.Size())
	ld.TraverseForward()
	ld.DeleteFront()
	fmt.Println("List Length: ", ld.Size())
	ld.TraverseForward()
	ld.TraverseReverse()
}
