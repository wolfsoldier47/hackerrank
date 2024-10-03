package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// making it first node
func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) insertBeforeValue(beforeValue, data int) {
	if list.head == nil {
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Cannot insert node, before value %d doesn't exist\n", beforeValue)

}

// insert value after the node
func (list *LinkedList) insertAfterValue(aftervalue, data int) {
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current != nil {
		if current.data == aftervalue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, after value %d doesn't exist\n", aftervalue)
	fmt.Println()
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) deleteFront() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Printf("Head node of the list has been deleted")
		return
	}
}

func (list *LinkedList) deleteLast() {
	if list.head == nil {
		fmt.Printf("Linked list is empty\n")
	}
	if list.head.next == nil {
		list.head = nil
		fmt.Printf("Last node of linked list has been deleted\n")
		return
	}
	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	fmt.Printf("Last node of linked list has been deleted")
}

func (list *LinkedList) deleteAfterValue(afterValue int) {
	current := list.head
	for current != nil && current.data != afterValue {
		current = current.next
	}
	if current == nil || current.next == nil {
		fmt.Printf("Node with after value %d doesn't exist\n", afterValue)
		return
	}
	fmt.Printf("Node after value %d has data %d and will be deleted", afterValue, current.next.data)

	current.next = current.next.next
}

func (list *LinkedList) deleteBeforeValue(beforeValue int) {
	current := list.head
	var previous *Node
	if current == nil || current.next == nil {
		fmt.Printf("Node with before value %d doesn't exist\n", beforeValue)
		return
	}
	for current.next != nil {
		if current.next.data == beforeValue {
			if previous == nil {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				list.head = current.next
			} else {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
	fmt.Printf("Node before value %d not found\n", beforeValue)
}

func (list *LinkedList) countNodes() (count int) {
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}

//find node at a given index

func (list *LinkedList) findNodeAt(index int) *Node {
	current := list.head
	count := list.countNodes()

	if index <= 0 || index > count {
		return nil
	}

	current = list.head
	for count = 1; count < index; count++ {
		current = current.next
	}
	return current
}

// Print the whole linked list
func (list *LinkedList) print() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	// Create an empty linked list
	myList := LinkedList{}

	// Insert some data at the beginning
	myList.insertAtFront(10)
	myList.insertAtFront(20)
	myList.insertAtFront(30)
	myList.insertAtBack(90)
	myList.insertAtBack(80)
	myList.insertAtBack(70)
	myList.insertAtFront(100)

	// Print the contents of the linked list
	fmt.Println("Linked List Contents:")
	myList.print()

	myList.insertAfterValue(20, 50)
	myList.print()
	// Count the nodes in the linked list
	count := myList.countNodes()
	fmt.Printf("Total number of nodes: %d\n", count)

	// Find and print a node at a specific index
	indexToFind := 4
	foundNode := myList.findNodeAt(indexToFind)
	if foundNode != nil {
		fmt.Printf("Node at index %d has data: %d\n", indexToFind, foundNode.data)
	} else {
		fmt.Printf("Node at index %d not found\n", indexToFind)
	}

	// Perform other operations as needed...

	// Delete the last node
	myList.deleteLast()

	// Print the updated linked list
	fmt.Println("Linked List After Deletion:")
	myList.print()
}
