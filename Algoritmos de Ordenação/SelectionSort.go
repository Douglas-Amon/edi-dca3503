package main

import "fmt"

func SelectionSort(v []int){
	for i := 0 ; i < len(v)-1 ; i++{
		menor := i
		for j := i+1 ; j < len(v) ; j++{
			if v[j] < v[menor]{
				menor = j
			}
		}
	v[i], v[menor] = v[menor], v[i]
	}
}

func main() {
	arr := []int{6, 1, 3, 2, 12, 4, 9}

	fmt.Println("Antes:", arr)

	SelectionSort(arr)

	fmt.Println("Depois:", arr)
}
