package main

import "fmt"

func main() {
	root := newRoot()
	for _, fn := range []func(root *Node){
		runStackTraversal,
		runRecursionTraversal,
		runMorrisTraversal,
	} {
		fn(root)
		fmt.Println()
	}
}

func runMorrisTraversal(root *Node) {
	fmt.Println("morris")

	curr := root
	for curr != nil {
		if curr.Left != nil {
			node := getPredecessor(curr.Left)
			node.Right = curr
			curr, curr.Left = curr.Left, nil
		} else {
			fmt.Println("node:", curr.Val)
			curr = curr.Right
		}
	}
}

func getPredecessor(root *Node) *Node {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func runRecursionTraversal(root *Node) {
	fmt.Println("recursion")
	recursion(root)
}

func recursion(root *Node) {
	if root.Left != nil {
		recursion(root.Left)
	}
	if root.Right != nil {
		recursion(root.Right)
	}
	fmt.Println("node:", root.Val)
}

func runStackTraversal(root *Node) {
	fmt.Println("stack")

	s := new(stack)
	s.Push(root)
	for !s.Empty() {
		node := s.Pop()
		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
		fmt.Println("node:", node.Val)
	}
}

type stack []*Node

func (s *stack) Push(n *Node) {
	*s = append(*s, n)
}

func (s *stack) Pop() (n *Node) {
	n = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func newRoot() *Node {
	return &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 9,
				Right: &Node{
					Val:   10,
					Right: &Node{Val: 11},
				},
			},
			Right: &Node{Val: 12},
		},
		Right: &Node{
			Val:  3,
			Left: &Node{Val: 4},
			Right: &Node{
				Val: 5,
				Right: &Node{
					Val: 6,
					Right: &Node{
						Val:   7,
						Right: &Node{Val: 8},
					},
				},
			},
		},
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
