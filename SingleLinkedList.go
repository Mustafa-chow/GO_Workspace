package main

import "fmt"

type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	data int
	next *Node
}

// This function will traverse the single linkedList and print the data
// (lList *LinkedList) is a pointer reciver so that in main method we can call directly
func (lList *LinkedList) printList() {
	currNode := lList.head
	for currNode != nil {
		fmt.Printf("%v ->", currNode.data)
		currNode = currNode.next
	}
	fmt.Println()
}

// add node in the end of the linked list
func (lList *LinkedList) append(node *Node) {
	if lList.length == 0 {
		lList.head = node
		lList.length++
	} else {
		currNode := lList.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = node
		lList.length++
	}
}

// time complexecity O(n2)
func (l *LinkedList) searchInList(item int) {
	tempNode := l.head
	isFound := false

	for isFound == false && tempNode != nil {
		if item == tempNode.data {
			isFound = true
		} else {
			tempNode = tempNode.next
		}
	}
	if isFound == true {
		fmt.Println("Item found in LinkedList")
	} else {
		fmt.Println("Item not found")
	}
}

// time complexecity O(n)
func (l *LinkedList) reverseLinkedList() {
	curr := l.head
	var prev *Node = nil
	var nxt *Node = nil

	for curr != nil {
		nxt = curr.next
		curr.next = prev
		prev = curr
		curr = nxt
	}
	l.head = prev
	l.printList()
}

// time complexecity O(n^2)
func (l *LinkedList) delete(item int) {
	curr := l.head
	tmp := l.head
	isFound := false
	for curr != nil && isFound == false {
		if curr.data == item {
			isFound = true
			for tmp.next != curr {
				tmp = tmp.next
			}
			tmp.next = curr.next
			curr = nil
		}
		if isFound == false {
			curr = curr.next
		}
	}
}

func main() {
	linkedList := LinkedList{nil, 0}
	node1 := &Node{data: 10, next: nil}
	node2 := &Node{data: 20, next: nil}
	node3 := &Node{data: 30, next: nil}
	node4 := &Node{data: 40, next: nil}
	node5 := &Node{data: 50, next: nil}

	linkedList.append(node1)
	linkedList.append(node2)
	linkedList.append(node3)
	linkedList.append(node4)
	linkedList.append(node5)

	fmt.Println("printing...")
	linkedList.printList()
	fmt.Println()

	var data1 int = 30
	linkedList.searchInList(data1)
	var data2 int = 70
	linkedList.searchInList(data2)
	fmt.Println()

	fmt.Println("Reverseing list...")
	linkedList.reverseLinkedList()
	fmt.Println()

	fmt.Println("Deleting...")
	linkedList.delete(data1)
	linkedList.printList()
	fmt.Println()

}
