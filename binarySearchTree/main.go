package main

import "fmt"

func main() {

	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(156)
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(130)
	tree.Insert(1110)
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(50))
	fmt.Println(tree)
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert will add a Node to the tree
func (n *Node) Insert(k int) {

	if n.Key < k {
		//Insert right

		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//Insert left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

//Search will find a node inside the method.
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true

}
