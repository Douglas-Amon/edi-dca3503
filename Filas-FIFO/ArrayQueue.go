package main

import (
	"errors"
	"fmt"
)

// circularidade da fila -> index := (q.front + i) % len(q.data)

type ArrayQueue struct {
	data  []int
	front int
	rear  int
	size  int
}

func (q *ArrayQueue) Init(capacity int) {
	q.data = make([]int, capacity)
	q.front = 0
	q.rear = 0
	q.size = 0
}

func (q *ArrayQueue) Enqueue(value int) {
	if q.size == len(q.data) {
		fmt.Println("erro: fila cheia")
		return
	}

	q.data[q.rear] = value
	q.rear = (q.rear + 1) % len(q.data)
	q.size++
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia")
	}

	value := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.size--

	return value, nil
}

func (q *ArrayQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia")
	}
	return q.data[q.front], nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) Print() {
	fmt.Print("Fila: ")
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % len(q.data)
		fmt.Print(q.data[index], " ")
	}
	fmt.Println()
}
