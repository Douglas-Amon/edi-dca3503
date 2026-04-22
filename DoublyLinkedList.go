package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dll *DoublyLinkedList) Add(value int) {
	newNode := &Node{value: value, next: nil, prev: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		dll.size++
		return
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
	dll.size++
}

func (dll *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index > dll.size {
		return fmt.Errorf("índice inválido: %d", index)
	}

	newNode := &Node{value: value, next: nil, prev: nil}

	if index == 0 {
		if dll.head == nil {
			dll.head = newNode
			dll.tail = newNode
		} else {
			newNode.next = dll.head
			dll.head.prev = newNode
			dll.head = newNode
		}
		dll.size++
		return nil
	}

	if index == dll.size {
		dll.Add(value)
		return nil
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	previous := current.prev

	newNode.prev = previous
	newNode.next = current

	previous.next = newNode
	current.prev = newNode

	dll.size++
	return nil
}

func (dll *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= dll.size {
		return fmt.Errorf("índice inválido: %d", index)
	}

	if dll.size == 1 {
		dll.head = nil
		dll.tail = nil
		dll.size--
		return nil
	}

	if index == 0 {
		dll.head = dll.head.next
		dll.head.prev = nil
		dll.size--
		return nil
	}

	if index == dll.size-1 {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		dll.size--
		return nil
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	previous := current.prev
	nextNode := current.next

	previous.next = nextNode
	nextNode.prev = previous

	dll.size--
	return nil
}

func (dll *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= dll.size {
		return -1, fmt.Errorf("índice inválido: %d", index)
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func (dll *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 || index >= dll.size {
		return fmt.Errorf("índice inválido: %d", index)
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = value
	return nil
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

func (dll *DoublyLinkedList) PrintForward() {
	current := dll.head
	fmt.Print("Lista: ")
	for current != nil {
		fmt.Printf("%d <-> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) PrintBackward() {
	current := dll.tail
	fmt.Print("Reversa: ")
	for current != nil {
		fmt.Printf("%d <-> ", current.value)
		current = current.prev
	}
	fmt.Println("nil")
}
