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
}

func (list *LinkedList) Append(value int) *LinkedList {
	if list.Head != nil {
		newNode := &Node{Data: value, Next: nil}
		tail := list.Head
		for ; tail.Next != nil; tail = tail.Next {
		}
		tail.Next = newNode
	} else {
		newNode := &Node{Data: value, Next: nil}
		list.Head = newNode
	}
	return list
}

func (list *LinkedList) Prepend(value int) *LinkedList {
	if list.Head != nil {
		newNode := &Node{Data: value, Next: list.Head}
		list.Head = newNode
	} else {
		newNode := &Node{Data: value, Next: nil}
		list.Head = newNode
	}
	return list
}

func (list *LinkedList) ToSlice() []int {
	lSlice := []int{}
	node := list.Head
	for ; node.Next != nil; node = node.Next {
		lSlice = append(lSlice, node.Data)
	}
	lSlice = append(lSlice, node.Data)
	return lSlice
}

func (list *LinkedList) Traverse() {
	node := list.Head
	for ; node.Next != nil; node = node.Next {
		fmt.Println(node.Data)
	}
	fmt.Println(node.Data)
}

func main() {
	list := &LinkedList{}
	list.Append(4).Append(5).Append(6).Append(7).Prepend(3)

	// val, _ := json.Marshal(list)
	// fmt.Printf("%s\n", val)

	fmt.Println(list.ToSlice())

	// list.Traverse()
}
