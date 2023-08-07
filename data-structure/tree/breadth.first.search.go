package tree

import "fmt"

/*
Reference: https://faun.pub/2-different-ways-to-implement-bfs-in-golang-8399f5d2452d

BFS: Breadth first search (aka Level Order Traversal)
TO traverse a tree, each level as arow from left to right
                    1
			--------|-------
			|              |
			2              3
		--------      ---------
		|      |      |        |
		4		5     6        7

BFS : 1,2,3,4,5,6,7

Solutions:

1. Implementation using Queue:
Iterate each element
Enqueue left & right node elements of that node
Dequeue the traversed node element into a visited list

2. Implementation using Map:
Store level in Key and Value as Slice of tree Nodes
Iterate map by each key as level

While storing values in Queue or Map, order needs to be maintained
InOrder: left to right
*/

type QueueNode struct {
	Node  *BstNode
	Level int
}

func BreadthFirstSearchViaQueue(root *BstNode) (visitedNode []interface{}, err error) {
	if root == nil {
		err = fmt.Errorf("Tree is empty!!")
		return
	}

	// create visited nodes slice
	visitedNode = make([]interface{}, 1)

	// Add first node to Queue
	q := []QueueNode{
		{Node: root, Level: 0},
	}

	// Iterate tree
	for len(q) > 0 {
		currNode, currLevel := q[0].Node, q[0].Level
		visitedNode = append(visitedNode, currNode.Data)

		// dequeuete traversed node
		q = q[1:]

		// Enqueue left node of current node
		if currNode.Left != nil {
			q = append(q, QueueNode{Node: currNode.Left, Level: currLevel + 1})
		}
		// Enqueue right node of current node
		if currNode.Right != nil {
			q = append(q, QueueNode{Node: currNode.Right, Level: currLevel + 1})
		}
	}
	return
}

/*
Store level in key and value(order left to right) as array having all elements of that level
*/
func BreadthFirstSearchViaMap(root *bstNode) (visitedNode []interface{}, err error) {

	if root == nil {
		err = fmt.Errorf("Tree is Empty!!")
		return
	}

	visitedNode = make([]interface{}, 0)
	level := 0
	hashMap := make(map[int][]*bstNode)

	// add root node to map
	hashMap[level] = []*bstNode{root}

	for {
		levelSlice, ok := hashMap[level]

		// exit condition
		if !ok {
			break
		}

		// iterate each level slice of map values
		for _, element := range levelSlice {
			visitedNode = append(visitedNode, element.data)
			// initialise hash map level
			if element.left != nil || element.right != nil {
				_, ok := hashMap[level+1]
				if !ok {
					hashMap[level+1] = []*bstNode{} //make([]*bstNode, 0)
				}
			}
			if element.left != nil {
				hashMap[level+1] = append(hashMap[level+1], element.left)
			}
			if element.right != nil {
				hashMap[level+1] = append(hashMap[level+1], element.right)
			}
		}
		level++
	}
	return
}
