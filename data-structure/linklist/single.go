package linklist

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct {
	Head *Node
	Len  int
}

func InitList() (ll *LinkList) {
	ll = &LinkList{}
	return
}

func (ll *LinkList) AddFront(data interface{}) {
	fmt.Println("\nAdd Front - ", data)
	node := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}
	ll.Len++
}
func (ll *LinkList) AddBack(data interface{}) {
	fmt.Println("\nAdd Back - ", data)
	node := &Node{Data: data}
	ll.Len++
	if ll.Head == nil { // if list is empty
		ll.Head = node
	} else { // travese to end of list
		current := ll.Head
		for {
			if current.Next != nil {
				current = current.Next
			} else {
				current.Next = node
				return
			}
		}
	}
}
func (ll *LinkList) DeleteFront() (err error) {
	fmt.Println("Delete Front")
	if ll.Head == nil {
		err = fmt.Errorf("List is empty !!")
		return
	}
	ll.Head = ll.Head.Next
	ll.Len--
	return
}

func (ll *LinkList) DeleteBack() (err error) {
	fmt.Println("Delete Back")
	if ll.Head == nil {
		err = fmt.Errorf("List is empty !!")
		return
	}
	current := ll.Head
	//prev := ll.Head
	for current.Next.Next != nil {
		//prev = current
		current = current.Next
	}
	current.Next = nil
	ll.Len--
	return
}
func (ll *LinkList) Traverse() (err error) {
	if ll.Head == nil {
		err = fmt.Errorf("List is empty !!")
		return
	}
	current := ll.Head
	fmt.Println()
	for current != nil {
		fmt.Printf("->%v", current.Data)
		current = current.Next
	}
	return
}

func (ll *LinkList) Reverse() (err error) {
	fmt.Println("Reverse List")
	if ll.Head == nil {
		err = fmt.Errorf("List is empty !!")
		return
	}

	prev := ll.Head
	current := ll.Head.Next
	next := current.Next
	for current.Next != nil {
		current.Next = prev
		prev = current
		current = next
		next = next.Next
	}
	current.Next = prev
	ll.Head.Next = nil
	ll.Head = current

	return
}

func LinklistExample() {
	ll := InitList()
	if err := ll.Traverse(); err != nil {
		fmt.Println(err)
	}
	ll.AddFront("a")
	ll.AddBack("b")
	ll.AddBack("c")
	ll.AddBack("d")

	if err := ll.Traverse(); err != nil {
		fmt.Println(err)
	}
	ll.DeleteBack()

	if err := ll.Traverse(); err != nil {
		fmt.Println(err)
	}
	ll.DeleteFront()

	if err := ll.Traverse(); err != nil {
		fmt.Println(err)
	}

	ll.AddFront("a")
	ll.AddBack("d")

	ll.Reverse()
	if err := ll.Traverse(); err != nil {
		fmt.Println(err)
	}

	ll.Reverse()
	if err := ll.Traverse(); err != nil {
		fmt.Println(err)
	}
}

/*
func Create(data interface{}) (ptr *Node) {
	fmt.Println("Create Linklist: ", data)
	ptr = &Node{Data: data}
	return
}
*/
/*
Add at 3 places
- Start
- End
- Anywhere
*/
/*
func Add(data interface{}, position int, head *Node) (newHead *Node, err error) {
	fmt.Println("\nAdd Node :", data)

	current := head
	currPos := 1
	for {
		switch position {
		case 0: // first position
			newNode := &Node{Data: data, Ptr: current}
			newHead = newNode
			return
		case 99: // last position
			if current.Ptr != nil {
				current = current.Ptr
			} else {
				newNode := &Node{Data: data}
				current.Ptr = newNode
				newHead = head
				return
			}
		default: // any other position
			if currPos == position {
				newNode := &Node{Data: data, Ptr: current.Ptr}
				current.Ptr = newNode
				newHead = head
				return
			} else {
				// move to next Node
				if current.Ptr == nil {
					err = fmt.Errorf("Position not valid!!")
					return
				}
				current = current.Ptr
			}

			currPos++
		}

	}

}
*/
/*
func Traverse(head *Node) {
	fmt.Println("Traverse Linklist")

	current := head
	for {
		if current.Ptr != nil {
			fmt.Printf("->%v", current.Data)
		} else {
			fmt.Printf("->%v", current.Data)
			break
		}
		current = current.Ptr
	}

}
*/
/*
func LinklistExample() {
	head := Create(1) // head(Node-1) -> [1,nil]
	Traverse(head)
	for i := 2; i <= 5; i++ {
		head, _ = Add(i, 99, head) // append at the end of the list
		//Traverse(head)
	}
	head, _ = Add(6, 0, head) // add at the start of the list
	Traverse(head)

	for i := 7; i <= 9; i++ {
		head, _ = Add(i, i-2, head) // append at different positions of the list
		//Traverse(head)
	}
	Traverse(head)

}
*/
