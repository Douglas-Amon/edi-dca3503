package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandom(ini, fim int) int {
	return rand.Intn(fim-ini+1) + ini
}

func QuickSort(v []int, ini int, fim int){
  if ini < fim{
    indexPivot := Partition(v, ini, fim)
    QuickSort(v, ini, indexPivot-1)
    QuickSort(v, indexPivot+1, fim)
  }
}

func Partition(v []int, ini int, fim int) int{
  aleat := getRandom(ini, fim)
  v[aleat], v[fim] = v[fim], v[aleat]
  pIndex := ini
  pivot := v[fim]
  for i := ini; i < fim; i++{
    if v[i] <= pivot{
      v[pIndex], v[i] = v[i], v[pIndex]
      pIndex++
    }
  }
  v[pIndex], v[fim] = v[fim], v[pIndex]
  return pIndex
}

func main() {
    rand.Seed(time.Now().UnixNano())
    v := []int{8, 3, 1, 7, 0, 10, 2}

    fmt.Println("Antes:", v)

    QuickSort(v, 0, len(v)-1)

    fmt.Println("Depois:", v)
}
