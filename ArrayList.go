package main

import "fmt"

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}

type ArrayList struct {
	valores   []int
	inseridos int
}

func (l *ArrayList) Init(size int) {
	l.valores = make([]int, size)
}

func (list *ArrayList) Add(value int) {
	if list.inseridos == len(list.valores) {
		list.doubleArray()
	}
	list.valores[list.inseridos] = value
	list.inseridos++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.inseridos {
		if list.inseridos == len(list.valores) {
			list.doubleArray()
		}
		for i := list.inseridos; i > index; i-- {
			list.valores[i] = list.valores[i-1]
		}
		list.valores[index] = value
		list.inseridos++
		return nil
	}
	return fmt.Errorf("Índice inválido: %d", index)
}

func (list *ArrayList) doubleArray() {
	newSize := len(list.valores) * 2
	if newSize == 0 {
		newSize = 1
	}
	newValores := make([]int, newSize)
	for i := 0; i < list.inseridos; i++ {
		newValores[i] = list.valores[i]
	}
	list.valores = newValores
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < list.inseridos {
		for i := index; i < list.inseridos-1; i++ {
			list.valores[i] = list.valores[i+1]
		}
		list.inseridos--
		return nil
	}
	return fmt.Errorf("Índice inválido: %d", index)
}

func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < list.inseridos {
		return list.valores[index], nil
	}
	return -1, fmt.Errorf("Índice inválido: %d", index)
}

func (list *ArrayList) Set(value int, index int) error {
	if index >= 0 && index < list.inseridos {
		list.valores[index] = value
		return nil
	}
	return fmt.Errorf("Índice inválido: %d", index)
}

func (list *ArrayList) Size() int {
	return list.inseridos
}

func (list *ArrayList) Reverse() {
	if list.inseridos == 0 {
		fmt.Println("A lista está vazia.")
		return
	}
	for i := 0; i < list.inseridos/2; i++ {
		list.valores[i] = list.valores[list.inseridos-1-i]
		list.valores[list.inseridos-1-i] = list.valores[i]
	}
}

func (list *ArrayList) Print() {
	fmt.Print("[")
	for i := 0; i < list.inseridos; i++ {
		fmt.Print(list.valores[i])
		if i < list.inseridos-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func main() {
	var lista ArrayList
	lista.Init(2)

	fmt.Println("Lista inicial:")
	lista.Print()

	fmt.Println("\nAdicionando elementos no fim...")
	lista.Add(10)
	lista.Add(20)
	lista.Add(30)
	lista.Print()

	fmt.Println("\nInserindo 15 na posição 1...")
	err := lista.AddOnIndex(15, 1)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	lista.Print()

	fmt.Println("\nRemovendo elemento da posição 2...")
	err = lista.RemoveOnIndex(2)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	lista.Print()

	fmt.Println("\nObtendo elemento da posição 1...")
	valor, err := lista.Get(1)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Valor encontrado:", valor)
	}

	fmt.Println("\nAlterando valor da posição 1 para 99...")
	err = lista.Set(99, 1)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	lista.Print()

	fmt.Println("\nTamanho da lista:")
	fmt.Println(lista.Size())

	fmt.Println("\nInvertendo a lista...")
	lista.Reverse()
	lista.Print()
}
