import "fmt"

package main

type BstNode struct{
  left *BstNode
  val int
  right *BstNode
  bf int
  height int
}

//Rotação para direita
func (root *BstNode) RotRight() *BstNode {
  newRoot := root.left
  root.left = newRoot.right
  root.UpdateProperties()
  newRoot.UpdateProperties()
  return newRoot
}

//Rotação para esquerda
func (root *BstNode) RotLeft *BstNode {
  newRoot := root.right
  root.right = newRoot.left
  root.UpdateProperties()
  newRoot.UpdateProperties()
  return newRoot
}

//Atualiza as propriedades de fator de balanço e altura == 
