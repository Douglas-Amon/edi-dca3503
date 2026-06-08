package main

import "fmt"

func merge(v []int, e []int, d []int){
  ie := 0
  id := 0
  iv := 0

  for ie < len(e) && id < len(d){
    if e[ie] < d[id]{
      v[iv] = e[ie]
      ie++
    } else { 
      v[iv] = d[id]
      id++
    }
    iv++
  }
  for ie < len(e){
    v[iv] = e[ie]
    ie++
    iv++
  }
  for id < len(d){
    v[iv] = d[id]
    id++
    iv++
  }
}

func MergeSort(v []int){
  if len(v) > 1 {
    tamE := len(v)/2
    tamD := len(v) - tamE
    e := make([]int, tamE)
    d := make([]int, tamD)
    for i := 0; i < tamE; i++{
      e[i] = v[i]
    }
    for i := 0; i < tamD; i++{
      d[i] = v[i+tamE]
    }
    MergeSort(e)
    MergeSort(d)
    merge(v, e, d)
  }
}

func main() {

	// Vetor original
	v := []int{8, 3, 5, 1, 9, 2, 7, 4}

	// Imprime antes da ordenação
	fmt.Println("Vetor desordenado:")
	fmt.Println(v)

	// Ordena
	MergeSort(v)

	// Imprime depois da ordenação
	fmt.Println("\nVetor ordenado:")
	fmt.Println(v)
}
