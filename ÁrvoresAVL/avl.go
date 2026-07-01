package main

import "fmt"

type BstNode struct{
  left *BstNode
  val int
  right *BstNode
  bf int
  height int
}

// Função auxiliar interna 
func createBSTNode(val int) *BstNode {
	return &BstNode{val: val, height: 0, bf: 0}
}

func (root *BstNode) Min() int {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current.val
}

func (root *BstNode) Max() int {
	current := root
	for current.right != nil {
		current = current.right
	}
	return current.val
}

//Rotação para direita
func (root *BstNode) RotRight() *BstNode {
  newRoot := root.left
  root.left = newRoot.right
  newRoot.right = root
  root.UpdateProperties()
  newRoot.UpdateProperties()
  return newRoot
}

//Rotação para esquerda
func (root *BstNode) RotLeft() *BstNode {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

//Atualiza as propriedades de fator de balanço e altura
func (root *BstNode) UpdateProperties() {
  if root.left == nil && root.right == nil{
    root.height = 0
    root.bf = 0
  } else {
    hleft := 0
	hright := 0
    if root.left != nil {
		hleft = root.left.height
	}
    if root.right != nil {
		hright = root.right.height
	}
    if hleft >= hright {
		root.height = hleft+1
	} else {
		root.height = hright+1
	}
    root.bf = hright - hleft
  }
}

//Caso 1: esquerda-esquerda
func (root *BstNode) RebalanceLeftLeft() *BstNode{
  return root.RotRight()
}

//Caso 2: esquerda-neutro
func (root *BstNode) RebalanceLeftNeutral() *BstNode{
  return root.RotRight()
}

//Caso 3: esquerda-direita
func (root *BstNode) RebalanceLeftRight() *BstNode{
  root.left = root.left.RotLeft()
  return root.RotRight()
}

//Caso 4: direita-direita
func (root *BstNode) RebalanceRightRight() *BstNode{
  return root.RotLeft()
}

//Caso 5: direita-neutro
func (root *BstNode) RebalanceRightNeutral() *BstNode{
  return root.RotLeft()
}

//Caso 6: direita-esquerda
func (root *BstNode) RebalanceRightLeft() *BstNode{
  root.right = root.right.RotRight()
  return root.RotLeft()
}

func (root *BstNode) Rebalance() *BstNode{
  if root.bf == -2 { //desbalanceamento à esquerda
     if root.left.bf == -1 {
       return root.RebalanceLeftLeft()
     } else if root.left.bf == 0 {
       return root.RebalanceLeftNeutral()
     } else if root.left.bf == 1 {
       return root.RebalanceLeftRight()
     }
  } else if root.bf == 2 { //desbalanceamento à direita
    if root.right.bf == 1 {
      return root.RebalanceRightRight()
    } else if root.right.bf == 0 {
      return root.RebalanceRightNeutral()
    } else if root.right.bf == -1 {
      return root.RebalanceRightLeft()
    }
  }
  return root
}

func (root *BstNode) Add(val int) *BstNode{
  if val <= root.val{
    if root.left != nil{
      root.left = root.left.Add(val)
    } else {
      root.left = createBSTNode(val)
    }
  } else { 
    if root.right != nil {
      root.right = root.right.Add(val)
    } else {
      root.right = createBSTNode(val)
    }
  }
  root.UpdateProperties()
  return root.Rebalance()
}
    
func (root *BstNode) Remove(val int) *BstNode{
  if val < root.val{
    root.left = root.left.Remove(val)
	return root
  } else if val > root.val{
    root.right = root.right.Remove(val)
	return root
  } else { // encontramos o nó
    if root.left == nil && root.right == nil{
      //caso 1: nó folha
      return nil // esse é o novo valor que o pai do nó folha apontará
    } else if root.left != nil && root.right == nil{
      //caso 2: há 1 filho a esquerda
      return root.left
    } else if root.left == nil && root.right != nil{
      //caso 2: há um filho a direita
      return root.right
    } else {
      // caso 3: 2 filhos
      min := root.right.Min()
      root.val = min
      root.right = root.right.Remove(min)
      return root
    }
  }
  root.UpdateProperties()
  return root.Rebalance()
}
