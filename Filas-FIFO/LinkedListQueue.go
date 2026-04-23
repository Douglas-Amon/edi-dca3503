package main

import "fmt"

type LinkedListQueue struct {
	size  int
	front *Node
	rear  *Node
}

type Node struct {
	value int
	next  *Node
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) Enqueue(value int) {
	newNode := &Node{value: value, next: nil}

	if q.rear == nil { // se a fila estiver vazia, o novo nó é tanto o front quanto o rear
		q.front = newNode
		q.rear = newNode
	} else { // caso contrário, o próximo do rear atual aponta para o novo nó
		q.rear.next = newNode
	}
	q.rear = newNode
	q.size++
}

func (q *LinkedListQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("fila vazia")
	}
	auxil := q.front.value // salva o valor de front
	q.front = q.front.next // coloca o front para o próximo

	if q.IsEmpty() { // se a fila ficou vazia, rear também deve ser nil
		q.rear = nil
	}

	q.size-- //decrementa o tamanho da fila
	return auxil, nil
}

func (q *LinkedListQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("fila vazia")
	}
	return q.front.value, nil
}
