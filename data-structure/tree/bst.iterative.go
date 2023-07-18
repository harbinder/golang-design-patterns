package tree

import "fmt"

type BstNode struct {
	Data  int
	Left  *BstNode
	Right *BstNode
}

type Bst struct {
	Root      *BstNode
	NodeCount int
}

func (t *Bst) CreateNode(data int) *BstNode {
	t.NodeCount++
	return &BstNode{Data: data}
}

func (t *Bst) InsertNode(data int) {
	if t.Root == nil {
		newNode := t.CreateNode(data)
		t.Root = newNode
		return
	}

	current := t.Root
	for {
		// move to left subtree
		if current.Data > data {
			if current.Left == nil {
				newNode := t.CreateNode(data)
				current.Left = newNode
				return
			}
			current = current.Left
		}
		// move to right subtree
		if current.Data <= data {
			if current.Right == nil {
				newNode := t.CreateNode(data)
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

func (t *Bst) Find(data int) (found bool) {
	current := t.Root
	for {
		if current.Data == data {
			fmt.Printf("\nNode found : %v", data)
			found = true
			return
		}
		if current.Data > data {
			if current.Left == nil {
				fmt.Printf("\nNode not found : %v", data)
				return
			}
			current = current.Left
		}
		if current.Data < data {
			if current.Right == nil {
				fmt.Printf("\nNode not found : %v", data)
				return
			}
			current = current.Right
		}
	}

}

func TreeBstIterativeExample() {
	data := []int{5, 3, 2, 4, 1, 7, 6, 8, 9}
	tree := &Bst{}
	for _, v := range data {
		tree.InsertNode(v)
	}
	fmt.Println("Total Nodes : ", tree.NodeCount)

	tree.Find(1)
	tree.Find(9)
	tree.Find(0)
}
