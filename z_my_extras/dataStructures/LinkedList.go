package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) Append(value int) *LinkedList {
	newNode := &Node{Data: value, Next: nil}
	if list.Head != nil {
		list.Tail.Next = newNode // point the current tail's next to this newNode
		list.Tail = newNode      // now that we have a new tail, which is this newNode
		return list
	}
	list.Head = newNode
	list.Tail = newNode
	return list
}

func (list *LinkedList) Prepend(value int) *LinkedList {
	newNode := &Node{Data: value, Next: nil}
	if list.Head != nil {
		newNode.Next = list.Head
		list.Head = newNode
		return list
	}
	list.Head = newNode
	list.Tail = newNode
	return list
}

func (list *LinkedList) ToSlice() []int {
	lSlice := []int{}
	node := list.Head
	for ; node != nil; node = node.Next {
		lSlice = append(lSlice, node.Data)
	}
	return lSlice
}

func (list *LinkedList) Traverse() {
	node := list.Head
	for ; node != nil; node = node.Next {
		fmt.Println(node.Data)
	}
}
