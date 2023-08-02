package tree

/*
  Node: memory location having DATA with left & right pointers
  Edge: connection/pointer from one NODE to other
  Tree: having NODES & EDGES
  SubTree: part of tree, which itself is a tree
  Left SubTree:
  Right SubTree:
	- Root:
	- Parent:
	- Child:
	- Siblings: nodes at same level
	- Leaf: having no child
	- Ancestors: nodes in the path towards root node
	- Descendents:
	- Depth: count of Edges when moving from node towards the root
	- Height: MAX. count of EDGES when moving from ROOT to LEAF NODE
	- Level: starts from 0
	Number of nodes @ level: 2^l
	root: l=0, total nodes: 2^0 = 1
		  l=1, total nodes: 2^1 = 2
		  l=2, total nodes: 2^2 = 4
	leaf: l=h, total nodes: 2^h
	total nodes in perfect tree:
		N = 2^0 + 2^1 + 2^2 ...... 2^h
		N = 2^(h+1) - 1
		2^(h+1) = N - 1
		h+1 = log(N-1)
		h = log(N)

	- Binary Tree (BT): in which every NODE has 0/1/2 CHILDREN
	- Strict BT: BT in which every NODE has exactly 2 CHILDREN
	- Complete BT: Strict BT  in which any NODE doesn't  have right NODE without left NODE
	- Perfect BT: Complete BT in which each NODE has left and right NODE except LEAF NODES
	- Balanced BT:
	- Binary Search Tree: in which every NODE has DATA > any node in its Left SubTree and Data < any node in its Right SubTree

	Time Complexity:
	  Operation       Array(unsorted)			Array(sorted)	LinkedList		BinartTree
	1. Search				O(n)					O(log n)		O(n)		O(log n)
	2. Insertion	O(1)-static,O(n)-dynamic 		O(n)			O(n)		O(log n)+
	3. Deletion     		O(n)					O(n)			O(n)		O(log n)+
*/

import "fmt"

const (
	PreOrder  = "PreOrder"
	InOrder   = "InOrder"
	PostOrder = "PostOrder"
)

type bstNode struct {
	data  int
	left  *bstNode
	right *bstNode
}

type bst struct {
	root        *bstNode
	visitedNode []int //optional: to store traversed nodes
}

func (t *bst) insert(data int) {
	fmt.Println("Insert node: ", data)
	t.insertRec(t.root, data)
}

func (t *bst) insertRec(node *bstNode, data int) *bstNode {
	// insert root node
	if t.root == nil {
		t.root = &bstNode{data: data}
		return t.root
	}
	// Order of this check is important!!
	// when we reach at leaf node, add new node
	// Means, when there is nil at either left or right node, insert the new node and return its address
	if node == nil {
		return &bstNode{data: data}
	}

	// traverse left, till exact node found
	if data < node.data {
		node.left = t.insertRec(node.left, data)
	}
	// traverse right, till exact node found
	if data > node.data {
		node.right = t.insertRec(node.right, data)
	}
	return node // important to return current node, to maintain the tree structure
}

func (t *bst) find(data int) {
	fmt.Println("\nFind data : ", data)
	node := t.findRecursive(t.root, data)
	if node == nil {
		fmt.Println("Node not found : ", data)
	} else {
		fmt.Println("Node found : ", node.data)
	}
	return
}

func (t *bst) findRecursive(node *bstNode, data int) *bstNode {
	// reached leaf node and data not found in any node
	if node == nil {
		return nil
	}
	// data found in current node being traversed
	if data == node.data {
		return node
	}
	// traverse left, if data less than current node data
	if data < node.data {
		return t.findRecursive(node.left, data)
	}
	// traverse right, if data less than current node data
	if data > node.data {
		return t.findRecursive(node.right, data)
	}
	return nil // should never reach this return statement
}

func (t *bst) traverse(order string) {
	fmt.Println("Tree Traversal : ")

	t.visitedNode = []int{} //make([]int, 0)

	switch order {
	case InOrder:
		fmt.Println("InOrder")
		t.inOrder(t.root)
	case PreOrder:
		fmt.Println("PreOrder")
		t.preOrder(t.root)
	case PostOrder:
		fmt.Println("PostOrder")
		t.postOrder(t.root)
	default:
		fmt.Println("Order not implemeted yet : ", order)
	}

}

func (t *bst) inOrder(node *bstNode) {
	if node != nil {
		t.inOrder(node.left)
		t.visitedNode = append(t.visitedNode, node.data)
		//fmt.Printf("%v->", node.data)
		t.inOrder(node.right)
	}
}

func (t *bst) preOrder(node *bstNode) {
	if node != nil {
		t.visitedNode = append(t.visitedNode, node.data)
		//fmt.Printf("%v->", node.data)
		t.inOrder(node.left)
		t.inOrder(node.right)
	}
}

func (t *bst) postOrder(node *bstNode) {
	if node != nil {
		t.inOrder(node.left)
		t.inOrder(node.right)
		t.visitedNode = append(t.visitedNode, node.data)
		//fmt.Printf("%v->", node.data)
	}
}

/*
				5
		|--------------|
		3				7
	|--------|		|-------|
	2		4		6		8

|-----|					|-------|
1								9

Inorder :   1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
Preorder :  5 -> 3 -> 2 -> 1 -> 4 -> 7 -> 6 -> 8 -> 9
Postorder : 1 -> 2 -> 4 -> 3 -> 6 -> 9 -> 8 -> 7 -> 5
*/
func TreeBstExample() {
	data := []int{5, 3, 2, 4, 1, 7, 6, 8, 9}
	tree := &bst{}
	for _, v := range data {
		tree.insert(v)
	}
	tree.traverse(InOrder)
	fmt.Println(tree.visitedNode)
	tree.traverse(PreOrder)
	fmt.Println(tree.visitedNode)
	tree.traverse(PostOrder)
	fmt.Println(tree.visitedNode)

	tree.find(1)
	tree.find(9)
	tree.find(0)

	fmt.Println("\nBreadth First Search: Using Map - (aka: Level Order Traversal)")
	visitedNode, err := BreadthFirstSearchViaMap(tree.root)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("LevelOrder Traversal : ", visitedNode)

	}

}
