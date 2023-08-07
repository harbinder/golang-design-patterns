package tree

/*

 */

type stackNode struct {
	node  *BstNode
	level int
}

func DepthFirstSearchViaStack(root *BstNode) (visitedNode []interface{}, err error) {
	visitedNode = []interface{}{}
	stack := []*BstNode{root}

	for len(stack) > 0 {
		//time.Sleep(time.Millisecond * 500)
		//fmt.Println("Start : stack length: ", len(stack))

		node := stack[len(stack)-1]
		// remove processed node from stack
		stack = stack[:len(stack)-1]

		// add popped node from stack to visited nodes list
		visitedNode = append(visitedNode, node)
		//fmt.Println("visited nodes : ", node.Data)

		// push left/right child of visited node to stack (right to left)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		//fmt.Println("End : stack length: ", len(stack))
	}

	return
}
