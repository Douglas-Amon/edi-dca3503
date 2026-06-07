package main

import "fmt"

func InsertionSort(v []int){
  for i := 1; i< len(v); i++{
    temp := v[i]
    j := i
    for ; j>0 && v[j-1] > temp; j--{
      v[j] = v[j-1]
    }
    v[j] = temp
  }
}

func main() {
	arr := []int{20, 1, 3, 2, 12, 4, 7, 9}

	fmt.Println("Antes:", arr)

	InsertionSort(arr)

	fmt.Println("Depois:", arr)
}
