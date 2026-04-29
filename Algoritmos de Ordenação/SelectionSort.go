package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		// procura o menor elemento no restante do array
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// troca o elemento atual com o menor encontrado
		arr[i] = arr[minIndex]
    arr[minIndex] = arr[i]
	}
}

func main() {
	arr := []int{6, 1, 3, 2, 12, 4, 9}

	fmt.Println("Antes:", arr)

	selectionSort(arr)

	fmt.Println("Depois:", arr)
}
