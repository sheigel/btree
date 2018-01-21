package main

import (
	"fmt"
	"math"
	"encoding/json"
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
	tree.printString()
	tree.Insert(5)
	tree.printString()
	tree.Insert(3)
	tree.printString()
	tree.Insert(2)
	tree.printString()
	tree.Insert(4)
	tree.printString()
	tree.Insert(10)
	tree.printString()
	tree.Insert(12)
	tree.printString()
	tree.Insert(11)
	tree.printString()
	tree.Insert(13)
	tree.printString()

	fmt.Printf("searching for key: %d = %+v\n", 10, tree.Search(10))

	fmt.Printf("tree is balanced=%t\n", tree.Root.IsBalanced())
	fmt.Printf("node 10 is balanced=%t\n", tree.Root.Right.IsBalanced())
}

type Node struct {
	Key    int   `json:"k"`
	Left   *Node `json:"l"`
	Right  *Node `json:"r"`
	parent *Node
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

	return math.Abs(float64(leftHeight-rightHeight)) <= 1 && n.Left.IsBalanced() && n.Right.IsBalanced()
}

type Tree struct {
	Root *Node
}

func (t *Tree) printString() {
	res1B, _ := json.MarshalIndent(t.Root, "", "       ")
	fmt.Println("tree", string(res1B))
}
func (t *Tree) Traverse(callback func(*Node) bool) {
	Traverse(t.Root, callback)
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
func Insert(key int, n *Node) *Node {
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{
				Key:    key,
				parent: n,
			}
			return n.Left
		}
		return Insert(key, n.Left)
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Key:    key,
				parent: n,
			}
			return n.Right
		}
		return Insert(key, n.Right)
	}
}

func findUnbalancePath(n *Node) (*Node, *Node, *Node) {
	var (
		a *Node
		b *Node
		c *Node
	)
	for ; n != nil; n = n.parent {
		c = b
		b = a
		a = n
		if !n.IsBalanced() {
			return a, b, c
		}
	}
	return nil, nil, nil
}

func rightRotate(to *Node) {
	var parent = to.parent
	var from = to.Left
	if parent != nil {
		if parent.Left == to {
			parent.Left = from
		}
		if parent.Right == to {
			parent.Right = from
		}
	}
	to.Left = from.Right
	from.Right = to
	from.parent = to.parent
	to.parent = from
}
func leftRotate(to *Node) {
	var parent = to.parent
	var from = to.Right
	if parent != nil {
		if parent.Left == to {
			parent.Left = from
		}
		if parent.Right == to {
			parent.Right = from
		}
	}
	to.Right = from.Left
	from.Left = to
	from.parent = to.parent
	to.parent = from
}

func BalanceFrom(n *Node) {
	a, b, c := findUnbalancePath(n)
	if a == nil {
		return
	}
	if a.Left == b && b.Left == c {
		rightRotate(a)
	} else if a.Left == b && b.Right == c {
		leftRotate(b)
		rightRotate(a)
	} else if a.Right == b && b.Right == c {
		leftRotate(a)
	} else if a.Right == b && b.Left == c {
		rightRotate(b)
		leftRotate(a)
	}
	m, _ := json.Marshal([] int{a.Key, b.Key, c.Key})
	fmt.Println("key=", n.Key, " unbalancePath ", string(m))
}

func findNewRoot(candidate *Node) *Node {
	if candidate.parent != nil {
		return findNewRoot(candidate.parent)
	}
	return candidate
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &Node{Key: key}
		return
	}

	var insertedNode = Insert(key, t.Root)

	BalanceFrom(insertedNode)
	t.Root = findNewRoot(t.Root)
}
