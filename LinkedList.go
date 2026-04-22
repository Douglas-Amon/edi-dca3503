package main

import "fmt"

type Node struct {
	value int   // Guarda o valor do nó
	next  *Node // Aponta para o próximo nó na lista
}

type LinkedList struct {
	head *Node // Aponta para o primeiro nó da lista
	size int   // Quantidade de elementos na lista
}

func (ll *LinkedList) Add(value int) {
	newNode := &Node{value: value, next: nil}
	if ll.head == nil {
		ll.head = newNode
		ll.size++
		return
	}
	aux := ll.head
	for aux.next != nil {
		aux = aux.next
	}
	aux.next = newNode
	ll.size++
}

func (ll *LinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= ll.size {
		newNode := &Node{value: value, next: nil}
		if index == 0 {
			// inserir no ínicio
			newNode.next = ll.head
			ll.head = newNode
			ll.size++
			return nil
		}
		if index == ll.size {
			// inserir no final
			ll.Add(value)
			return nil
		}
		// inserir no meio
		aux := ll.head
		for i := 0; i < index-1; i++ {
			aux = aux.next
		}
		newNode.next = aux.next
		aux.next = newNode
		ll.size++
		return nil

	}
	return fmt.Errorf("Índice inválido: %d", index)
}

func (ll *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < ll.size {
		if index == 0 {
			// remover o primeiro elemento
			ll.head = ll.head.next
			ll.size--
			return nil
		}
		// remover o elemento do meio ou do final
		anterior := ll.head
		for i := 0; i < index-1; i++ {
			anterior = anterior.next
		}
		// pular o nó a ser removido
		anterior.next = anterior.next.next
		ll.size--
		return nil
	}
	return fmt.Errorf("Índice inválido: %d", index)
}

func (ll *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < ll.size {
		aux := ll.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.value, nil
	}
	return -1, fmt.Errorf("Índice inválido: %d", index)
}

func (ll *LinkedList) Set(value int, index int) error {
	if index >= 0 && index < ll.size {
		aux := ll.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		aux.value = value
		return nil
	}
	return fmt.Errorf("Índice inválido: %d", index)
}

func (ll *LinkedList) Size() int {
	return ll.size
}

// Função auxiliar para imprimir a lista
func (ll *LinkedList) Print() {
	aux := ll.head
	fmt.Print("Lista: ")
	for aux != nil {
		fmt.Printf("%d -> ", aux.value)
		aux = aux.next
	}
	fmt.Println("nil")
}
