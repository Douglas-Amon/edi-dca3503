package main

import "fmt"

func CountingSort(v []int) []int{
  //contar maior e menor
  maior := v[0]
  menor := v[0]
  for i:=1; i < len(v); i++{
    if v[i] > maior {
      maior = v[i]
    }
    if v[i] < menor {
      menor = v[i]
    }
  }

  c := make([]int, maior-menor+1)
  //contar ocorrencias
  for i:=0; i < len(v); i++{
    c[v[i]-menor]++
  }
  //soma acumulativa
  for i:=0; i < len(c); i++{
    c[i]+=c[i-1]
  }
  //criar vetor ordenado e reconstruir
  ord := make([]int, len(v))
  for i:=0; i < len(v); i++{
    indiceOrd := c[v[i]-menor]-1
    c[v[i]-menor]--
    ord[indiceOrd] = v[i]
  }
  return ord
}
