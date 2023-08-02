package tree

import "fmt"

/*
				5
		|--------------|
		3				7
	|--------|		|-------|
	2		4		6		8
|-------|				|-------|
1								9

Reference: https://faun.pub/implementing-recursive-and-iterative-dfs-on-a-binary-tree-golang-eda04949f4ee
Traversal: Depth First Search
Inorder :   1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
Preorder :  5 -> 3 -> 2 -> 1 -> 4 -> 7 -> 6 -> 8 -> 9
Postorder : 1 -> 2 -> 4 -> 3 -> 6 -> 9 -> 8 -> 7 -> 5
*/

type BstNode struct {
	Data  int
	Left  *BstNode
	Right *BstNode
}

type Bst struct {
	Root        *BstNode
	NodeCount   int
	visitedNode []int
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

func (t *Bst) Traversal(order string) {
	fmt.Println("\nTree Traversal : ")
	var visitedNode []*BstNode

	switch order {
	case InOrder:
		fmt.Println("InOrder - Using Stack")
		visitedNode = t.inOrder(t.Root)
		for _, node := range visitedNode {
			fmt.Printf("-->%v", node.Data)
		}

	case PreOrder:
		fmt.Println("PreOrder")
		//visitedNode = t.preOrder(t.root)
		fmt.Println(visitedNode)
	case PostOrder:
		fmt.Println("PostOrder")
		//visitedNode = t.postOrder(t.root)
		fmt.Println(visitedNode)
	default:
		fmt.Println("Order not implemeted yet : ", order)
	}
}

func (t *Bst) inOrder(root *BstNode) (visitedNode []*BstNode) {
	visitedNode = []*BstNode{} //make([]int, 0)

	// add root element node to stack
	stack := []*BstNode{root}

	for len(stack) > 0 {
		// pop node from top of stack
		popNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// append to visited nodes slice
		visitedNode = append(visitedNode, popNode)

		// push left/right nodes to stack
		if popNode.Right != nil {
			stack = append(stack, popNode.Right)
		}
		if popNode.Left != nil {
			stack = append(stack, popNode.Left)
		}
	}
	return
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

	tree.Traversal(InOrder)

	fmt.Println("\nBreadth First Search: Using Queue - (aka: Level Order Traversal)")
	visitedNode, err := BreadthFirstSearchViaQueue(tree.Root)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("LevelOrder Traversal : ", visitedNode)

	}

}
