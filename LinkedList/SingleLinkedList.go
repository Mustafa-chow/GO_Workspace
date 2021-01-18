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
	fmt.Println("printing...")
	currNode := lList.head;
	for lList.length != 0 {
		fmt.Printf("%s -> ", currNode.data);
		currNode = currNode.next;
		lList.length--;
	}
}

// add node in the end of the linked list
func (lList *LinkedList) append(node *Node) {
	if lList.length == 0 {
		lList.head = node;
		lList.length++;
	} else {
		currNode := lList.head
		for currNode.next != nil {
			currNode = currNode.next;
		}
		currNode.next = node;
		lList.length++;
	}

}

func main() {
	linkedList := LinkedList{nil, 0};
	node1 := &Node{data: 10, next: nil};
	node2 := &Node{data: 20, next: nil};
	node3 := &Node{data: 30, next: nil};
	node4 := &Node{data: 40, next: nil};
	node5 := &Node{data: 50, next: nil};

	linkedList.append(node1);
	linkedList.append(node2);
	linkedList.append(node3);
	linkedList.append(node4);
	linkedList.append(node5);

	linkedList.printList();
}