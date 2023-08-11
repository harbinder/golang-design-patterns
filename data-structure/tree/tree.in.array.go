package tree

import "fmt"

/*
A Binary tree can be respresented by an array.
For any element at index = i
Left child of i = 2*i
Right child of i = 2*i+1
Parent of i = floor[i/2]

A Level-Order Traversal of a binary tree will result in an array follwoing the  rules mentioned above

*/

func TreeViaArrayExample() {
	data := []int{5, 3, 2, 4, 1, 7, 6, 8, 9}
	tree := &Bst{}
	for _, v := range data {
		tree.InsertNode(v)
	}
	fmt.Println("Total Nodes : ", tree.NodeCount)

	fmt.Println("\nBreadth First Search: Using Queue - (aka: Level Order Traversal)")
	visitedNode, err := BreadthFirstSearchViaQueue(tree.Root)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("LevelOrder Traversal : ", visitedNode)
	}

	fmt.Println("Array representation of Tree")
	for i, val := range visitedNode {
		parentIndex := i
		leftChildIndex := 2 * i
		rightChildIndex := 2*i + 1
		if len(visitedNode) > leftChildIndex {
			fmt.Printf("\nParent(%v): %v -> Left-Right : %v-%v",
				val,
				visitedNode[parentIndex],
				visitedNode[leftChildIndex],
				visitedNode[rightChildIndex],
			)
		}

	}
}
