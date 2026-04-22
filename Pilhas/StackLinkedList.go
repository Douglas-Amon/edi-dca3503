package pilhas

import "fmt"

type Node struct {
	value int   // valor armazenado no nó
	next  *Node // ponteiro para o próximo nó
}

type StackLinkedList struct {
	top  *Node // ponteiro para o topo da pilha
	size int   // tamanho da pilha
}

func (sll *StackLinkedList) IsEmpty() bool {
	return sll.top == nil // a pilha está vazia se o topo for nulo
}

func (sll *StackLinkedList) Size() int {
	return sll.size // retorna o tamanho da pilha
}

func (sll *StackLinkedList) Top() (int, error) {
	if sll.IsEmpty() {
		return 0, fmt.Errorf("A pilha está vazia") // retorna um erro se a pilha estiver vazia
	}
	return sll.top.value, nil // retorna o valor do topo da pilha
}

func (sll *StackLinkedList) Push(value int) {
	newNode := &Node{value: value, next: sll.top}
	sll.top = newNode // o novo nó se torna o topo da pilha
	sll.size++        // incrementa o tamanho da pilha
}

func (sll *StackLinkedList) Pop() (int, error) {
	if sll.IsEmpty() {
		return 0, fmt.Errorf("A pilha está vazia") // retorna um erro se a pilha estiver vazia
	}
	value := sll.top.value // obtém o valor do topo da pilha
	sll.top = sll.top.next // o próximo nó se torna o novo topo da pilha
	sll.size--             // decrementa o tamanho da pilha

	return value, nil
}
