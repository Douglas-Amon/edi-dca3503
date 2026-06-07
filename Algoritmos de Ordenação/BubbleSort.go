package main

import "fmt"

func BubbleSort(v []int){
  for varredura := 0; varredura < len(v)-1; varredura++{
    trocou :=  false
    for i := 0; i < len(v)-varredura-1; i++{
      if v[i] > v[i+1]{
        v[i], v[i+1] = v[i+1], v[i]
        trocou := true
      }
    }
    if !trocou{
      return 
    }
  }
}

func main() {
	arr := []int{20, 1, 5, 2, 15, 4, 7, 9}

	fmt.Println("Antes:", arr)

	BubbleSort(arr)

	fmt.Println("Depois:", arr)
}
