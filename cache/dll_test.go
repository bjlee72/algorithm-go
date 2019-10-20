package cache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testValue struct {
	value int
}

func (v testValue) String() string {
	return fmt.Sprintf("%d", v.value)
}

func TestAddToFront(t *testing.T) {
	dll := DoublyLinkedList{}

	dll.AddToFront(&Node{value: testValue{1}})

	status := dll.String()
	assert.Equal(t, "1", status)

	dll.AddToFront(&Node{value: testValue{2}})

	status = dll.String()
	assert.Equal(t, "2->1", status)

	dll.AddToFront(&Node{value: testValue{3}})

	status = dll.String()
	assert.Equal(t, "3->2->1", status)
}

func TestRemoveFromTail(t *testing.T) {
	dll := DoublyLinkedList{}
	dll.AddToFront(&Node{value: testValue{4}})
	dll.AddToFront(&Node{value: testValue{3}})
	dll.AddToFront(&Node{value: testValue{2}})
	dll.AddToFront(&Node{value: testValue{1}})

	node := dll.RemoveFromTail()
	assert.Equal(t, "4", node.Value().String())

	node = dll.RemoveFromTail()
	assert.Equal(t, "3", node.Value().String())

	node = dll.RemoveFromTail()
	assert.Equal(t, "2", node.Value().String())

	node = dll.RemoveFromTail()
	assert.Equal(t, "1", node.Value().String())
}

func TestMoveToFront(t *testing.T) {
	dll := DoublyLinkedList{}

	node4 := &Node{value: testValue{4}}
	node3 := &Node{value: testValue{3}}
	node2 := &Node{value: testValue{2}}
	node1 := &Node{value: testValue{1}}

	dll.AddToFront(node4)
	dll.AddToFront(node3)
	dll.AddToFront(node2)
	dll.AddToFront(node1)

	dll.MoveToFront(node3)
	assert.Equal(t, "3->1->2->4", dll.String())

	dll.MoveToFront(node2)
	assert.Equal(t, "2->3->1->4", dll.String())

	dll.MoveToFront(node1)
	assert.Equal(t, "1->2->3->4", dll.String())

	dll.MoveToFront(node4)
	assert.Equal(t, "4->1->2->3", dll.String())
}
