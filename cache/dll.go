package cache

import "fmt"

// Value stored inside the Node.
type Value interface {
	String() string
}

// Node inside the doubly linked list.
type Node struct {
	prev  *Node
	next  *Node
	value Value
}

// Value returns the value of the value field (getter)
func (node *Node) Value() Value {
	return node.value
}

// SetValue sets the value field (setter)
func (node *Node) SetValue(value Value) {
	node.value = value
}

// String returns the string representation of the Node.
func (node *Node) String() string {
	if node.next == nil {
		return node.Value().String()
	}

	stringToNext := node.next.String()

	return fmt.Sprintf("%s->%s", node.Value().String(), stringToNext)
}

// DoublyLinkedList that keeps head and tail pointers.
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// String returns the string representation of the doubly linked list.
func (dll *DoublyLinkedList) String() string {
	if dll.head == nil {
		return "nil"
	}

	return dll.head.String()
}

// RemoveFromTail removes a node from the tail.
func (dll *DoublyLinkedList) RemoveFromTail() *Node {
	return dll.Remove(dll.tail)
}

// MoveToFront moves a node in the doubly linked list to the front.
func (dll *DoublyLinkedList) MoveToFront(node *Node) {
	dll.AddToFront(dll.Remove(node))
}

// AddToFront adds a Node to the front of the given doubly linked list.
func (dll *DoublyLinkedList) AddToFront(node *Node) {

	if dll.head == nil {
		dll.head = node
		dll.tail = node
		node.prev = nil
		node.next = nil
		return
	}

	node.prev = nil
	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}

// Remove removes a Node from the given doubly linked list.
func (dll *DoublyLinkedList) Remove(node *Node) *Node {
	if node == nil {
		// doesn't need to do anything
		return nil
	}

	if node.prev == nil && node.next == nil {
		dll.head = nil
		dll.tail = nil
		return node
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		dll.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		dll.tail = node.prev
	}

	// it's removed, thus the pointers should be initialized.
	node.prev = nil
	node.next = nil
	return node
}
