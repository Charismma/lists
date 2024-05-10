package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{value: data, next: nil}
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
func (list *LinkedList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println()
}
func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{value: data, next: nil}
		list.head = newNode
		return
	}
	newNode := &Node{value: data, next: list.head}
	list.head = newNode
}
func (list *LinkedList) insertAfterValue(afterValue, data int) {
	newNode := &Node{value: data, next: nil}
	current := list.head
	for current != nil {
		if current.value == afterValue {
			newNode.next = current.next
			current.next = newNode

		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, after value %d doesn't exist", afterValue)
	fmt.Println()
}
func (list *LinkedList) insertBeforeValue(beforeValue, data int) {
	if list.head == nil {
		return
	}
	if list.head.value == beforeValue {
		newNode := &Node{value: data, next: list.head}
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		if current.next.value == beforeValue {
			newNode := &Node{value: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, before value %d doesn't exist", beforeValue)
	fmt.Println()
}
func main() {
	spisok := LinkedList{head: nil}
	spisok.insertAtBack(10)
	spisok.insertAtBack(5)
	spisok.insertAtFront(44)
	spisok.print()
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
