package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) insert(key int) {
	if root == nil {
		root = &Node{data: key}
	}
	if key < root.data {
		if root.left == nil {
			root.left = &Node{data: key}
		} else {
			root.left.insert(key)
		}
	}
	if key > root.data {
		if root.right == nil {
			root.right = &Node{data: key}
		} else {
			root.right.insert(key)
		}
	}

}

func (root *Node) searchItem(item int) bool {

	if root == nil {
		return false
	}
	if item < root.data {
		root.left.searchItem(item)
	}
	if item > root.data {
		root.right.searchItem(item)
	}
	return true
}

func (root *Node) preOrder() {
	if root == nil {
		return
	}
	fmt.Printf("%d -> ", root.data)
	root.left.preOrder()
	root.right.preOrder()
}

func (root *Node) inOrder() {
	if root == nil {
		return
	}
	root.left.preOrder()
	fmt.Printf("%d -> ", root.data)
	root.right.preOrder()
}

func (root *Node) postOrder() {
	if root == nil {
		return
	}
	root.left.preOrder()
	root.right.preOrder()
	fmt.Printf("%d -> ", root.data)
}

func (root *Node) countTotalNode() int {
	if root == nil {
		return 0
	}
	return 1 + root.left.countTotalNode() + root.right.countTotalNode()

}

func (root *Node) leaves() int {
	if root == nil {
		return 0
	} else if root.left == nil && root.right == nil {
		return 1
	}
	return root.left.leaves() + root.right.leaves()
}

func main() {
	tree := &Node{data: 15}
	tree.insert(17)
	tree.insert(10)
	tree.insert(20)

	fmt.Println("Pre Order: ")
	tree.preOrder()
	fmt.Println()

	fmt.Println("In Order: ")
	tree.inOrder()
	fmt.Println()

	fmt.Println("Post Order: ")
	tree.postOrder()
	fmt.Println()

	// searching in tree
	isFound := tree.searchItem(15)
	if isFound == true {
		fmt.Println("Item found in TREE")
	} else {
		fmt.Println("Item NOT found in TREE")
	}

	fmt.Println()

	totalNodes := tree.countTotalNode()
	fmt.Printf("Total nodes: %d", totalNodes)

	fmt.Println()
	fmt.Println()

	totalLeaves := tree.leaves()
	fmt.Printf("Total leaves: %d", totalLeaves)

}
