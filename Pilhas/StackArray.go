package pilhas

import "fmt"

type StackArray struct {
	items []int
}

func (s *StackArray) IsEmpty() bool {
	return len(s.items) == 0 // a pilha está vazia se o comprimento do slice for zero
}

func (s *StackArray) Size() int {
	return len(s.items) // retorna tamanho da pilha
}

func (s *StackArray) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("A pilha está vazia") // retorna um erro se a pilha estiver vazia
	}
	return s.items[len(s.items)-1], nil // retorna o valor do topo da pilha
}

func (s *StackArray) Push(value int) {
	s.items = append(s.items, value) // adiciona o valor ao final da pilha
}

func (s *StackArray) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("A pilha está vazia") // retorna um erro se a pilha estiver vazia
	}
	topIndex := len(s.items) - 1 // índice do último elemento da pilha
	value := s.items[topIndex]   // obtém o último valor da pilha
	s.items = s.items[:topIndex] // remove o último valor da pilha
	return value, nil
}
