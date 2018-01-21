package main

import (
	"fmt"
	"math"
)

func main() {
	tree := Tree{
		//Root: &Node{
		//	Key: 8,
		//	Left: &Node{
		//		Key: 5,
		//		Left: &Node{
		//			Key: 3,
		//			Left: &Node{
		//				Key: 2,
		//			},
		//			Right: &Node{
		//				Key: 4,
		//			},
		//		},
		//	},
		//	Right: &Node{
		//		Key: 10,
		//		Right: &Node{
		//			Key: 12,
		//			Left: &Node{
		//				Key: 11,
		//			},
		//			Right: &Node{
		//				Key: 13,
		//			},
		//		},
		//	},
		//},
	}

	tree.Insert(8)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(13)

	fmt.Println(tree.String())

	fmt.Printf("searching for key: %d = %+v\n", 10, tree.Search(10))

	fmt.Printf("tree is balanced=%t\n", tree.Root.IsBalanced())
	fmt.Printf("node 10 is balanced=%t\n", tree.Root.Right.IsBalanced())
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func Traverse(n *Node, callback func(*Node) bool) {
	if n == nil {
		return
	}
	if callback(n) == true {
		return
	}

	Traverse(n.Left, callback)
	Traverse(n.Right, callback)
}
func height(n *Node) int {
	if n == nil {
		return 0
	}
	leftHeight := height(n.Left)
	rightHeight := height(n.Right)
	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

func (n *Node) IsBalanced() bool {
	if n == nil {
		return true
	}

	var leftHeight = height(n.Left)
	var rightHeight = height(n.Right)

	return math.Abs(float64(leftHeight-rightHeight)) <= 0 && n.Left.IsBalanced() && n.Right.IsBalanced()
}

type Tree struct {
	Root *Node
}

func (t *Tree) Traverse(callback func(*Node) bool) {
	Traverse(t.Root, callback)
}

func (t *Tree) String() string {
	var str = ""
	t.Traverse(func(n *Node) bool {
		str += fmt.Sprintf("k:%d ", n.Key)
		return false
	})
	return str
}

func (t *Tree) Search(key int) *Node {
	var foundNode *Node
	eq := func(n *Node) bool {
		if key == n.Key {
			foundNode = n
			return true
		}
		return false
	}

	t.Traverse(eq)

	return foundNode
}
func findInsertionPointer(key int, n *Node) **Node {
	if key < n.Key {
		if n.Left == nil {
			return &n.Left
		}
		return findInsertionPointer(key, n.Left)
	} else {
		if n.Right == nil {
			return &n.Right
		}
		return findInsertionPointer(key, n.Right)
	}
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &Node{Key: key}
		return
	}

	var insertionPointer = findInsertionPointer(key, t.Root)

	node := Node{Key: key}
	*insertionPointer = &node
}
