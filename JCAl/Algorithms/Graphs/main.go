package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PrintNode(n *Node) {
	ok := hasLeftNode(n)
	if !ok {
		fmt.Println("No Left Node")
	} else {
		fmt.Println("Left Node:", n.Left.Value)
	}

	ok = hasRightNode(n)
	if !ok {
		fmt.Println("No Right Node")
	} else {
		fmt.Println("Right Node:", n.Right.Value)
	}
}

func hasLeftNode(n *Node) bool {
	ok := true
	if n.Left != nil {
		return ok
	}

	return !ok
}

func hasRightNode(n *Node) bool {
	ok := true
	if n.Right != nil {
		return ok
	}

	return !ok
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(N)

	var nodes = make([]Node, N)
	for i := 0; i < N; i++ {
		var val, li, ri int
		fmt.Scanf("%d %d %d", &val, &li, &ri)
		nodes[i].Value = val
		if li >= 0 {
			nodes[i].Left = &nodes[li]
		}

		if ri >= 0 {
			nodes[i].Right = &nodes[ri]
		}

		PrintNode(&nodes[i])
	}
}
